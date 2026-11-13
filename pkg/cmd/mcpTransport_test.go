package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mcpServer answers initialize with version and tools/call with reply, as a
// server-sent events stream. Requests after initialize are recorded.
type mcpServer struct {
	version   string
	initError bool
	splitData bool // split the tools/call reply across data: lines
	// bigFirstLine sends a first data: line too long to stay put in the
	// scanner's buffer.
	bigFirstLine bool
	redirect     bool // answer tools/call with a redirect
	holdOpen     chan struct{}

	mu           sync.Mutex
	afterInitHdr []http.Header
}

func (m *mcpServer) headersAfterInit() []http.Header {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.afterInitHdr
}

func (m *mcpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     any    `json:"id"`
		Method string `json:"method"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if req.Method != "initialize" {
		m.mu.Lock()
		m.afterInitHdr = append(m.afterInitHdr, r.Header.Clone())
		m.mu.Unlock()
	}

	switch req.Method {
	case "notifications/initialized":
		w.WriteHeader(http.StatusAccepted)
		return
	case "initialize":
		if m.initError {
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%v,"error":{"code":-32602,"message":"Unsupported protocol version"}}`, req.ID)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%v,"result":{"protocolVersion":%q}}`, req.ID, m.version)
		return
	}

	if m.redirect {
		http.Redirect(w, r, "https://elsewhere.example.com/mcp", http.StatusTemporaryRedirect)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	if m.bigFirstLine {
		// The second line arrives in a later read, after the scanner has had
		// to move the first one within its buffer.
		long := strings.Repeat("x", 40*1024)
		fmt.Fprintf(w, "data: {\"jsonrpc\":\"2.0\",\"id\":%v,\"result\":{\"content\":[{\"type\":\"text\",\"text\":\"%s\"}\n", req.ID, long)
		w.(http.Flusher).Flush()
		time.Sleep(50 * time.Millisecond)
		fmt.Fprint(w, "data: ]}}\n\n")
		return
	}
	reply := fmt.Sprintf(`{"jsonrpc":"2.0","id":%v,"result":{"content":[{"type":"text","text":"{\"ok\":true}"}]}}`, req.ID)
	if m.splitData {
		half := strings.Index(reply, `"result"`)
		fmt.Fprintf(w, "event: message\ndata: %s\ndata: %s\n\n", reply[:half], reply[half:])
	} else {
		fmt.Fprintf(w, "event: message\ndata: %s\n\n", reply)
	}
	w.(http.Flusher).Flush()
	if m.holdOpen != nil {
		// Keep the stream open after the reply, as a server may.
		select {
		case <-m.holdOpen:
		case <-r.Context().Done():
		}
	}
}

func callTool(t *testing.T, server *mcpServer) (*http.Response, error) {
	t.Helper()
	httpServer := httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	if server.holdOpen != nil {
		t.Cleanup(func() { close(server.holdOpen) })
	}

	transport := &mcpTransport{mcpURL: httpServer.URL + "/mcp", operation: "GetUsers", path: "/api/v1/users", base: http.DefaultTransport}
	req, err := http.NewRequest(http.MethodGet, "https://mybiz.kinde.com/api/v1/users?page_size=1", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer token")

	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := transport.RoundTrip(req)
		done <- result{resp, err}
	}()
	select {
	case r := <-done:
		return r.resp, r.err
	case <-time.After(5 * time.Second):
		t.Fatal("the transport didn't return after the reply arrived")
		return nil, nil
	}
}

func TestMCPTransport_SendsTheAgreedProtocolVersion(t *testing.T) {
	server := &mcpServer{version: "2025-03-26"}

	resp, err := callTool(t, server)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	headers := server.headersAfterInit()
	require.Len(t, headers, 2, "the initialized notification and the tool call")
	for _, header := range headers {
		assert.Equal(t, "2025-03-26", header.Get("MCP-Protocol-Version"), "the version the server agreed, not the one asked for")
	}
}

func TestMCPTransport_JoinsDataLinesOfAnEvent(t *testing.T) {
	resp, err := callTool(t, &mcpServer{version: mcpProtocolVersion, splitData: true})

	require.NoError(t, err)
	body, _ := io.ReadAll(resp.Body)
	assert.JSONEq(t, `{"ok":true}`, string(body))
}

func TestMCPTransport_ReturnsWhenTheReplyArrivesEvenIfTheStreamStaysOpen(t *testing.T) {
	resp, err := callTool(t, &mcpServer{version: mcpProtocolVersion, holdOpen: make(chan struct{})})

	require.NoError(t, err)
	body, _ := io.ReadAll(resp.Body)
	assert.JSONEq(t, `{"ok":true}`, string(body))
}

func TestMCPTransport_KeepsLongDataLinesIntact(t *testing.T) {
	resp, err := callTool(t, &mcpServer{version: mcpProtocolVersion, bigFirstLine: true})

	require.NoError(t, err)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, strings.Repeat("x", 40*1024), string(body))
}

func TestMCPTransport_RedirectIsAnErrorNotSuccess(t *testing.T) {
	_, err := callTool(t, &mcpServer{version: mcpProtocolVersion, redirect: true})

	assert.EqualError(t, err, "the MCP server redirected the request (307 Temporary Redirect) instead of answering it")
}

func TestMCPTransport_ReportsAFailedInitialize(t *testing.T) {
	server := &mcpServer{initError: true}

	_, err := callTool(t, server)

	assert.EqualError(t, err, "the MCP server refused to start a session: Unsupported protocol version")
	assert.Empty(t, server.headersAfterInit(), "nothing is sent after a failed initialize")
}
