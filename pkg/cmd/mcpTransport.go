package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// mcpTransport sends the Management API requests built by the SDK client to
// the domain's MCP server instead, as a call to the tool of the same name.
// The browser login's token only works there; the MCP server makes the REST
// call with the user's access and returns the REST response as text.
//
// Sitting under the SDK client keeps everything above it unchanged: flags,
// --body, confirmations, JSON output and error messages.
type mcpTransport struct {
	mcpURL    string
	operation string // the tool to call; tool names match the SDK's operation names
	path      string // the operation's path template, e.g. /api/v1/users/{id}
	base      http.RoundTripper

	nextID      int
	sessionID   string
	initialized bool
	// redacted is set when the response hid a secret.
	redacted bool
}

// userFacingError is an error meant for the user as it is, without the SDK
// client's "send request" wrapping.
type userFacingError struct{ msg string }

func (e *userFacingError) Error() string { return e.msg }

var (
	// mcpAPIError matches how the MCP server reports a failed REST call.
	mcpAPIError = regexp.MustCompile(`(?s)API error \(HTTP (\d{3})\):\s*(.*)$`)
	// redactedSecret is how the MCP server hides secrets in responses.
	redactedSecret = []byte(`"<redacted>"`)
)

func (t *mcpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	args, err := t.toolArguments(req)
	if err != nil {
		return nil, err
	}
	auth := req.Header.Get("Authorization")

	if !t.initialized {
		if resp := t.initialize(req.Context(), auth); resp != nil {
			return resp, nil
		}
	}

	status, _, body, err := t.post(req.Context(), auth, t.request("tools/call", map[string]any{"name": t.operation, "arguments": args}))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return response(req, status, body), nil
	}

	message, err := t.reply(body)
	if err != nil {
		return nil, err
	}
	if message.Error != nil {
		return nil, t.toolError(req.Context(), auth, message.Error.Message)
	}

	var text strings.Builder
	for _, content := range message.Result.Content {
		text.WriteString(content.Text)
	}
	if message.Result.IsError {
		if match := mcpAPIError.FindStringSubmatch(text.String()); match != nil {
			status, _ := strconv.Atoi(match[1])
			return response(req, status, []byte(match[2])), nil
		}
		return response(req, http.StatusBadGateway, []byte(text.String())), nil
	}

	result := []byte(text.String())
	if bytes.Contains(result, redactedSecret) {
		t.redacted = true
	}
	return response(req, http.StatusOK, result), nil
}

// toolArguments flattens the REST request into tool arguments: path and query
// parameters plus the top-level fields of the JSON body, as the MCP tools take them.
func (t *mcpTransport) toolArguments(req *http.Request) (map[string]any, error) {
	args := map[string]any{}

	if req.Body != nil && req.Body != http.NoBody {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return nil, err
		}
		mediaType, _, _ := mime.ParseMediaType(req.Header.Get("Content-Type"))
		switch {
		case len(bytes.TrimSpace(data)) == 0:
		case mediaType == "application/json":
			decoder := json.NewDecoder(bytes.NewReader(data))
			decoder.UseNumber()
			var body map[string]any
			if err := decoder.Decode(&body); err != nil {
				return nil, &userFacingError{fmt.Sprintf("'%s' can't be sent with a browser login: its request body isn't a JSON object. %s", t.operation, m2mHint)}
			}
			for key, value := range body {
				args[key] = value
			}
		default:
			return nil, &userFacingError{fmt.Sprintf("uploading files isn't available with a browser login. %s", m2mHint)}
		}
	}

	for key, values := range req.URL.Query() {
		args[key] = strings.Join(values, ",")
	}

	template := strings.Split(strings.Trim(t.path, "/"), "/")
	actual := strings.Split(strings.Trim(req.URL.EscapedPath(), "/"), "/")
	if len(template) != len(actual) {
		return nil, fmt.Errorf("request path %s doesn't match %s", req.URL.EscapedPath(), t.path)
	}
	for i, segment := range template {
		if name, ok := strings.CutPrefix(segment, "{"); ok {
			value, err := url.PathUnescape(actual[i])
			if err != nil {
				return nil, err
			}
			args[strings.TrimSuffix(name, "}")] = value
		}
	}
	return args, nil
}

// initialize starts the MCP session. It returns a response to pass back when
// the MCP server refuses, e.g. 401 when the login has been revoked.
func (t *mcpTransport) initialize(ctx context.Context, auth string) *http.Response {
	status, header, body, err := t.post(ctx, auth, t.request("initialize", map[string]any{
		"protocolVersion": "2025-06-18",
		"capabilities":    map[string]any{},
		"clientInfo":      map[string]any{"name": "kinde-cli", "version": userAgent()},
	}))
	if err != nil {
		return response(nil, http.StatusBadGateway, []byte(err.Error()))
	}
	if status != http.StatusOK {
		return response(nil, status, body)
	}
	t.sessionID = header.Get("Mcp-Session-Id")
	t.initialized = true
	// A notification has no reply worth checking.
	t.post(ctx, auth, map[string]any{"jsonrpc": "2.0", "method": "notifications/initialized"})
	return nil
}

// toolError explains a JSON-RPC error from a tool call. The commonest cause
// is a Management API operation the MCP server doesn't offer.
func (t *mcpTransport) toolError(ctx context.Context, auth, message string) error {
	status, _, body, err := t.post(ctx, auth, t.request("tools/list", map[string]any{}))
	if err == nil && status == http.StatusOK {
		if reply, err := t.reply(body); err == nil && reply.Result.Tools != nil && !hasTool(reply.Result.Tools, t.operation) {
			return &userFacingError{fmt.Sprintf("this command isn't available with a browser login. %s", m2mHint)}
		}
	}
	return &userFacingError{fmt.Sprintf("the MCP server rejected the request: %s", message)}
}

const m2mHint = "Log in with an M2M application to use it: run 'kinde login --client-id <client ID>'"

func hasTool(tools []struct {
	Name string `json:"name"`
}, name string) bool {
	for _, tool := range tools {
		if tool.Name == name {
			return true
		}
	}
	return false
}

func (t *mcpTransport) request(method string, params any) map[string]any {
	t.nextID++
	return map[string]any{"jsonrpc": "2.0", "id": t.nextID, "method": method, "params": params}
}

func (t *mcpTransport) post(ctx context.Context, auth string, message any) (status int, header http.Header, body []byte, err error) {
	data, err := json.Marshal(message)
	if err != nil {
		return 0, nil, nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.mcpURL, bytes.NewReader(data))
	if err != nil {
		return 0, nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	req.Header.Set("Authorization", auth)
	if t.sessionID != "" {
		req.Header.Set("Mcp-Session-Id", t.sessionID)
	}
	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return resp.StatusCode, resp.Header, body, err
}

type rpcReply struct {
	ID     json.RawMessage `json:"id"`
	Result struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
		IsError bool `json:"isError"`
		Tools   []struct {
			Name string `json:"name"`
		} `json:"tools"`
	} `json:"result"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// reply returns the reply to the last request, from a JSON or server-sent
// events response.
func (t *mcpTransport) reply(body []byte) (*rpcReply, error) {
	var payloads [][]byte
	if trimmed := bytes.TrimSpace(body); bytes.HasPrefix(trimmed, []byte("{")) {
		payloads = append(payloads, trimmed)
	} else {
		for _, line := range bytes.Split(body, []byte("\n")) {
			if data, ok := bytes.CutPrefix(bytes.TrimRight(line, "\r"), []byte("data:")); ok && len(bytes.TrimSpace(data)) > 0 {
				payloads = append(payloads, data)
			}
		}
	}
	want := strconv.Itoa(t.nextID)
	for _, payload := range payloads {
		var reply rpcReply
		if json.Unmarshal(payload, &reply) == nil && string(reply.ID) == want {
			return &reply, nil
		}
	}
	return nil, errors.New("the MCP server sent no reply")
}

// response builds the HTTP response the SDK client sees.
func response(req *http.Request, status int, body []byte) *http.Response {
	return &http.Response{
		StatusCode:    status,
		Status:        fmt.Sprintf("%d %s", status, http.StatusText(status)),
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        http.Header{"Content-Type": {"application/json"}},
		Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Request:       req,
	}
}
