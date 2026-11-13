package cmd

import (
	"bufio"
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

// mcpProtocolVersion is the MCP version the CLI asks for when it starts a
// session.
const mcpProtocolVersion = "2025-06-18"

// maxMCPReply limits how much of an MCP reply is read.
const maxMCPReply = 32 << 20

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

	nextID          int
	sessionID       string
	protocolVersion string // agreed when the session started
	initialized     bool
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
		refused, err := t.initialize(req.Context(), auth)
		if err != nil {
			return nil, err
		}
		if refused != nil {
			return response(req, refused.status, refused.body), nil
		}
	}

	reply, refused, err := t.call(req.Context(), auth, "tools/call", map[string]any{"name": t.operation, "arguments": args})
	if err != nil {
		return nil, err
	}
	if refused != nil {
		return response(req, refused.status, refused.body), nil
	}
	if reply.Error != nil {
		return nil, t.toolError(req.Context(), auth, reply.Error.Message)
	}

	var text strings.Builder
	for _, content := range reply.Result.Content {
		text.WriteString(content.Text)
	}
	if reply.Result.IsError {
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

// initialize starts the MCP session and agrees the protocol version. refused
// is set when the MCP server answers with an HTTP error, e.g. 401 when the
// login has been revoked, to pass back to the SDK client.
func (t *mcpTransport) initialize(ctx context.Context, auth string) (refused *httpFailure, err error) {
	reply, refused, err := t.call(ctx, auth, "initialize", map[string]any{
		"protocolVersion": mcpProtocolVersion,
		"capabilities":    map[string]any{},
		"clientInfo":      map[string]any{"name": "kinde-cli", "version": userAgent()},
	})
	if err != nil || refused != nil {
		return refused, err
	}
	if reply.Error != nil {
		return nil, &userFacingError{fmt.Sprintf("the MCP server refused to start a session: %s", reply.Error.Message)}
	}
	version := reply.Result.ProtocolVersion
	if version == "" {
		return nil, &userFacingError{"the MCP server didn't agree a protocol version"}
	}
	t.protocolVersion = version
	t.initialized = true

	if err := t.notify(ctx, auth, "notifications/initialized"); err != nil {
		return nil, err
	}
	return nil, nil
}

// toolError explains a JSON-RPC error from a tool call. The commonest cause
// is a Management API operation the MCP server doesn't offer.
func (t *mcpTransport) toolError(ctx context.Context, auth, message string) error {
	reply, refused, err := t.call(ctx, auth, "tools/list", map[string]any{})
	if err == nil && refused == nil && reply.Error == nil && reply.Result.Tools != nil && !hasTool(reply.Result.Tools, t.operation) {
		return &userFacingError{fmt.Sprintf("this command isn't available with a browser login. %s", m2mHint)}
	}
	return &userFacingError{fmt.Sprintf("the MCP server rejected the request: %s", message)}
}

const m2mHint = "Log in with an M2M application to use it: run 'kinde login --client-id <client ID>'"

func hasTool(tools []mcpTool, name string) bool {
	for _, tool := range tools {
		if tool.Name == name {
			return true
		}
	}
	return false
}

type (
	// httpFailure is an HTTP error from the MCP server.
	httpFailure struct {
		status int
		body   []byte
	}

	mcpTool struct {
		Name string `json:"name"`
	}

	rpcReply struct {
		ID     json.RawMessage `json:"id"`
		Result struct {
			ProtocolVersion string `json:"protocolVersion"`
			Content         []struct {
				Text string `json:"text"`
			} `json:"content"`
			IsError bool      `json:"isError"`
			Tools   []mcpTool `json:"tools"`
		} `json:"result"`
		Error *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
)

// call sends a JSON-RPC request and returns its reply. refused is set
// instead when the MCP server answers with an HTTP error.
func (t *mcpTransport) call(ctx context.Context, auth, method string, params any) (*rpcReply, *httpFailure, error) {
	t.nextID++
	id := strconv.Itoa(t.nextID)
	resp, err := t.send(ctx, auth, map[string]any{"jsonrpc": "2.0", "id": t.nextID, "method": method, "params": params})
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, maxMCPReply))
		return nil, &httpFailure{status: resp.StatusCode, body: body}, nil
	}
	reply, err := readReply(resp, id)
	return reply, nil, err
}

// notify sends a JSON-RPC notification, which has no reply.
func (t *mcpTransport) notify(ctx context.Context, auth, method string) error {
	resp, err := t.send(ctx, auth, map[string]any{"jsonrpc": "2.0", "method": method})
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode >= 300 {
		return &userFacingError{fmt.Sprintf("the MCP server refused %s: %s", method, resp.Status)}
	}
	return nil
}

func (t *mcpTransport) send(ctx context.Context, auth string, message any) (*http.Response, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.mcpURL, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	req.Header.Set("Authorization", auth)
	if t.sessionID != "" {
		req.Header.Set("Mcp-Session-Id", t.sessionID)
	}
	if t.protocolVersion != "" {
		req.Header.Set("MCP-Protocol-Version", t.protocolVersion)
	}
	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if id := resp.Header.Get("Mcp-Session-Id"); id != "" {
		t.sessionID = id
	}
	return resp, nil
}

// readReply reads the reply with the given ID from a JSON response, or from a
// server-sent events stream as events arrive, so a stream left open after the
// reply doesn't hold the command up.
func readReply(resp *http.Response, id string) (*rpcReply, error) {
	mediaType, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if mediaType != "text/event-stream" {
		body, err := io.ReadAll(io.LimitReader(resp.Body, maxMCPReply))
		if err != nil {
			return nil, err
		}
		if reply, ok := matchReply(body, id); ok {
			return reply, nil
		}
		return nil, errors.New("the MCP server sent no reply")
	}

	scanner := bufio.NewScanner(io.LimitReader(resp.Body, maxMCPReply))
	scanner.Buffer(make([]byte, 64*1024), maxMCPReply)
	var data [][]byte
	for {
		more := scanner.Scan()
		line := scanner.Bytes()
		// A blank line, or the end of the stream, ends an event. Its data
		// lines are joined with newlines before decoding.
		if !more || len(line) == 0 {
			if len(data) > 0 {
				if reply, ok := matchReply(bytes.Join(data, []byte("\n")), id); ok {
					return reply, nil
				}
				data = nil
			}
			if !more {
				break
			}
			continue
		}
		if value, ok := bytes.CutPrefix(line, []byte("data:")); ok {
			data = append(data, bytes.TrimPrefix(value, []byte(" ")))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read the MCP server's reply: %w", err)
	}
	return nil, errors.New("the MCP server sent no reply")
}

// matchReply decodes a JSON-RPC message, returning it if it answers id.
func matchReply(data []byte, id string) (*rpcReply, bool) {
	var reply rpcReply
	if json.Unmarshal(data, &reply) != nil || string(reply.ID) != id {
		return nil, false
	}
	return &reply, true
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
