package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/kinde-oss/kinde-cli/pkg/release"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// accessTokenFunc returns an access token for the Management API.
type accessTokenFunc func(ctx context.Context) (string, error)

// callOperation calls a Management API operation and prints the response
// body as JSON on stdout.
//
// The SDK returns error responses (400, 403, 429...) as ordinary values, not
// errors, so the response is recorded at the HTTP level and any status of 400
// or more becomes an error. Printing the recorded body, rather than the SDK's
// decoded value, also keeps fields the SDK doesn't know about yet.
func callOperation(cmd *cobra.Command, serverURL string, httpClient *http.Client, accessToken accessTokenFunc, operation string, methodArgs []reflect.Value) error {
	ctx := cmd.Context()
	log := log.Ctx(ctx)

	recorder := &recordingTransport{base: httpClient.Transport}
	if recorder.base == nil {
		recorder.base = http.DefaultTransport
	}
	recordingClient := *httpClient
	recordingClient.Transport = recorder

	client, err := management_api.NewClient(serverURL, bearerToken(accessToken), management_api.WithClient(&recordingClient))
	if err != nil {
		return fmt.Errorf("failed to create management API client: %w", err)
	}

	method := reflect.ValueOf(client).MethodByName(operation)
	results := method.Call(append([]reflect.Value{reflect.ValueOf(ctx)}, methodArgs...))
	callErr, _ := results[len(results)-1].Interface().(error)

	response := recorder.last
	if response == nil {
		// The request was never sent, e.g. no access token could be obtained.
		return callErr
	}
	if response.status >= 400 {
		return newAPIError(response, docsFor(operation).Scopes)
	}
	if callErr != nil {
		// The API accepted the request; the SDK just couldn't decode the
		// response (often a value newer than the SDK). The body is still valid.
		log.Debug().Err(callErr).Msg("SDK could not decode the response")
	}
	return writeJSON(cmd.OutOrStdout(), response.body)
}

// writeJSON writes an API response body, indented when it's JSON.
func writeJSON(out io.Writer, body []byte) error {
	body = bytes.TrimSpace(body)
	if len(body) == 0 {
		return nil
	}
	var indented bytes.Buffer
	if err := json.Indent(&indented, body, "", "  "); err == nil {
		body = indented.Bytes()
	}
	_, err := fmt.Fprintf(out, "%s\n", body)
	return err
}

// bearerToken implements management_api.SecuritySource.
type bearerToken accessTokenFunc

func (f bearerToken) KindeBearerAuth(ctx context.Context, _ management_api.OperationName) (management_api.KindeBearerAuth, error) {
	token, err := f(ctx)
	return management_api.KindeBearerAuth{Token: token}, err
}

type recordedResponse struct {
	status int
	header http.Header
	body   []byte
}

// recordingTransport keeps the status and body of the last response, and
// identifies the CLI in the User-Agent header.
type recordingTransport struct {
	base http.RoundTripper
	last *recordedResponse
}

func (t *recordingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("User-Agent", userAgent())

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	t.last = &recordedResponse{status: resp.StatusCode, header: resp.Header, body: body}
	return resp, nil
}

func userAgent() string {
	version := release.Version
	if version == "" {
		version = "dev"
	}
	return "kinde-cli/" + version
}

// apiError is an error response from the Management API.
type apiError struct {
	status     int
	messages   []string
	scopes     []string // the scopes the operation needs, from the API docs
	retryAfter string
}

func newAPIError(response *recordedResponse, scopes []string) *apiError {
	return &apiError{
		status:     response.status,
		messages:   errorMessages(response.body),
		scopes:     scopes,
		retryAfter: response.header.Get("Retry-After"),
	}
}

func (e *apiError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "the API returned %d %s", e.status, http.StatusText(e.status))
	if len(e.messages) > 0 {
		b.WriteString(": " + strings.Join(e.messages, "; "))
	}

	switch e.status {
	case http.StatusUnauthorized:
		b.WriteString(". The M2M application's credentials may have changed: run 'kinde login' again")
	case http.StatusForbidden:
		if len(e.scopes) > 0 {
			fmt.Fprintf(&b, ". Check that your M2M application has the %s scope for the Kinde Management API", strings.Join(e.scopes, ", "))
		}
	case http.StatusTooManyRequests:
		if e.retryAfter != "" {
			fmt.Fprintf(&b, ". Try again in %s seconds", e.retryAfter)
		} else {
			b.WriteString(". Try again shortly")
		}
	}
	return b.String()
}

// maxRawErrorBody is how much of an unrecognized error body is shown.
const maxRawErrorBody = 300

// errorMessages extracts readable messages from an error response body.
// Kinde normally sends {"errors": [{"code": ..., "message": ...}]}.
func errorMessages(body []byte) []string {
	var parsed struct {
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`
		Code             string `json:"code"`
		Message          string `json:"message"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
	body = bytes.TrimSpace(body)
	if len(body) == 0 {
		return nil
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		text := string(body)
		if len(text) > maxRawErrorBody {
			text = text[:maxRawErrorBody] + "..."
		}
		return []string{text}
	}

	var messages []string
	add := func(code, message string) {
		switch {
		case code != "" && message != "":
			messages = append(messages, code+": "+message)
		case code+message != "":
			messages = append(messages, code+message)
		}
	}
	for _, e := range parsed.Errors {
		add(e.Code, e.Message)
	}
	add(parsed.Code, parsed.Message)
	add(parsed.Error, parsed.ErrorDescription)
	return messages
}
