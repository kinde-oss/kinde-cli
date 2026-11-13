package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/kinde-oss/kinde-cli/pkg/agentauth"
	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

const fakeIssuer = "https://app.kinde.com/oauth2/agents1"

// fakeKinde plays both the business's MCP server (mybiz.kinde.com/mcp) and
// Kinde's agent sign-in server (app.kinde.com), like the real thing.
type fakeKinde struct {
	*httptest.Server
	t       *testing.T
	enabled bool

	mu        sync.Mutex
	toolCalls []map[string]any
	// tools answers tool calls by name: the text to return and whether it's
	// an error. Unlisted tools don't exist.
	tools     map[string]func(args map[string]any) (string, bool)
	refreshes int
}

func newFakeKinde(t *testing.T) *fakeKinde {
	f := &fakeKinde{t: t, enabled: true, tools: map[string]func(map[string]any) (string, bool){}}
	mux := http.NewServeMux()
	f.Server = httptest.NewServer(mux)
	t.Cleanup(f.Close)

	mux.HandleFunc("POST /mcp", f.handleMCP)
	mux.HandleFunc("GET /.well-known/oauth-protected-resource/mcp", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"resource": "https://" + testDomain + "/mcp", "authorization_servers": []string{fakeIssuer}})
	})
	mux.HandleFunc("GET /.well-known/oauth-authorization-server/oauth2/agents1", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"issuer":                           fakeIssuer,
			"authorization_endpoint":           fakeIssuer + "/authorize",
			"token_endpoint":                   fakeIssuer + "/token",
			"registration_endpoint":            fakeIssuer + "/register",
			"code_challenge_methods_supported": []string{"S256"},
		})
	})
	mux.HandleFunc("POST /oauth2/agents1/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{"client_id": "cli-client"})
	})
	mux.HandleFunc("GET /oauth2/agents1/authorize", func(w http.ResponseWriter, r *http.Request) {
		redirect, _ := url.Parse(r.URL.Query().Get("redirect_uri"))
		redirect.RawQuery = url.Values{"code": {"code-1"}, "state": {r.URL.Query().Get("state")}}.Encode()
		http.Redirect(w, r, redirect.String(), http.StatusFound)
	})
	mux.HandleFunc("POST /oauth2/agents1/token", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.Header().Set("Content-Type", "application/json")
		token := "access-1"
		if r.PostForm.Get("grant_type") == "refresh_token" {
			f.mu.Lock()
			f.refreshes++
			f.mu.Unlock()
			token = "access-2"
		}
		json.NewEncoder(w).Encode(map[string]any{"access_token": token, "refresh_token": "refresh-1", "token_type": "bearer", "expires_in": 3600})
	})
	return f
}

func (f *fakeKinde) handleMCP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if !f.enabled {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if auth != "Bearer access-1" && auth != "Bearer access-2" {
		w.Header().Set("WWW-Authenticate", `Bearer resource_metadata="https://`+testDomain+`/.well-known/oauth-protected-resource/mcp"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var req struct {
		ID     any            `json:"id"`
		Method string         `json:"method"`
		Params map[string]any `json:"params"`
	}
	require.NoError(f.t, json.NewDecoder(r.Body).Decode(&req))

	reply := map[string]any{"jsonrpc": "2.0", "id": req.ID}
	switch req.Method {
	case "notifications/initialized":
		w.WriteHeader(http.StatusAccepted)
		return
	case "initialize":
		w.Header().Set("Mcp-Session-Id", "session-1")
		reply["result"] = map[string]any{"protocolVersion": "2025-06-18"}
	case "tools/list":
		var tools []map[string]any
		for name := range f.tools {
			tools = append(tools, map[string]any{"name": name})
		}
		reply["result"] = map[string]any{"tools": tools}
	case "tools/call":
		assert.Equal(f.t, "session-1", r.Header.Get("Mcp-Session-Id"))
		name := req.Params["name"].(string)
		args, _ := req.Params["arguments"].(map[string]any)
		f.mu.Lock()
		f.toolCalls = append(f.toolCalls, map[string]any{"name": name, "arguments": args})
		f.mu.Unlock()
		tool, ok := f.tools[name]
		if !ok {
			reply["error"] = map[string]any{"code": -32602, "message": "Tool " + name + " not found"}
			break
		}
		text, isError := tool(args)
		reply["result"] = map[string]any{"content": []map[string]any{{"type": "text", "text": text}}, "isError": isError}
	}
	// Answer as server-sent events, as streamable HTTP servers may.
	data, _ := json.Marshal(reply)
	w.Header().Set("Content-Type", "text/event-stream")
	fmt.Fprintf(w, "event: message\ndata: %s\n\n", data)
}

// lastToolCall returns the last tool call the MCP server received.
func (f *fakeKinde) lastToolCall() map[string]any {
	f.mu.Lock()
	defer f.mu.Unlock()
	require.NotEmpty(f.t, f.toolCalls)
	return f.toolCalls[len(f.toolCalls)-1]
}

// routeKindeTo sends requests for *.kinde.com to the fake, and makes the
// "browser" follow the sign-in URL like a real one would.
func routeKindeTo(t *testing.T, f *fakeKinde) {
	target, _ := url.Parse(f.URL)
	client := &http.Client{Transport: rewriteKinde{target}}
	previousClient, previousBrowser := newHTTPClient, openBrowser
	newHTTPClient = func() *http.Client { return &http.Client{Transport: rewriteKinde{target}} }
	openBrowser = func(signInURL string) error {
		go func() {
			if resp, err := client.Get(signInURL); err == nil {
				resp.Body.Close()
			}
		}()
		return nil
	}
	t.Cleanup(func() { newHTTPClient, openBrowser = previousClient, previousBrowser })
}

type rewriteKinde struct{ target *url.URL }

func (r rewriteKinde) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.HasSuffix(req.URL.Hostname(), "kinde.com") {
		req = req.Clone(req.Context())
		req.URL.Scheme, req.URL.Host, req.Host = r.target.Scheme, r.target.Host, ""
	}
	return http.DefaultTransport.RoundTrip(req)
}

func inTerminal(stdin string) func(*iostreams.IOStreams, *bytes.Buffer) {
	return func(ios *iostreams.IOStreams, in *bytes.Buffer) {
		ios.SetTerminalInput(func() (string, error) { return "", nil })
		in.WriteString(stdin)
	}
}

// loginThroughBrowser runs 'kinde login' against the fake, asking for the domain.
func loginThroughBrowser(t *testing.T, cli *testCLI) {
	t.Helper()
	_, stderr, err := cli.runWith(inTerminal(testDomain+"\n"), "login")
	require.NoError(t, err, stderr)
}

func TestBrowserLogin_AsksForTheDomainAndSavesTheLogin(t *testing.T) {
	kinde := newFakeKinde(t)
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)

	_, stderr, err := cli.runWith(inTerminal(testDomain+"\n"), "login")

	require.NoError(t, err, stderr)
	assert.Contains(t, stderr, "Kinde domain (e.g. mybusiness.kinde.com): ")
	assert.Contains(t, stderr, "Opening the browser to sign in to "+testDomain)
	assert.Contains(t, stderr, "Logged in to "+testDomain+" through the browser")
	assert.Contains(t, stderr, browserLoginLimits, "the browser login's limits are explained up front")
	assert.Equal(t, testDomain, cli.readFullConfig().CurrentEnvironment)
	assert.Equal(t, &config.BrowserLogin{ClientID: "cli-client", Issuer: fakeIssuer}, cli.readConfig().Browser)
	assert.NotContains(t, readConfigFile(t, cli), "access-1", "tokens must stay out of the config file")
}

func TestBrowserLogin_NotTurnedOn(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.enabled = false
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{})

	_, _, err := cli.runWith(inTerminal(""), "login")

	assert.ErrorContains(t, err, "the browser login isn't turned on for "+testDomain)
	assert.ErrorContains(t, err, "kinde login --client-id")
	assert.Nil(t, cli.readConfig().Browser)
}

func TestBrowserLogin_CommandsGoThroughMCP(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.tools["GetUsers"] = func(args map[string]any) (string, bool) {
		return `{"code":"OK","users":[{"id":"kp_1","email":"sally@example.com"}]}`, false
	}
	kinde.tools["GetUserData"] = func(args map[string]any) (string, bool) {
		return `{"id":"` + args["id"].(string) + `"}`, false
	}
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)
	loginThroughBrowser(t, cli)

	stdout, stderr, err := cli.run("users", "list", "--page-size", "5", "--expand", "organizations")
	require.NoError(t, err, stderr)
	assert.JSONEq(t, `{"code":"OK","users":[{"id":"kp_1","email":"sally@example.com"}]}`, stdout)
	assert.Equal(t, map[string]any{"name": "GetUsers", "arguments": map[string]any{"page_size": "5", "expand": "organizations"}}, kinde.lastToolCall())

	_, _, err = cli.run("users", "get", "--id", "kp_1")
	require.NoError(t, err)
	assert.Equal(t, map[string]any{"name": "GetUserData", "arguments": map[string]any{"id": "kp_1"}}, kinde.lastToolCall())
}

func TestBrowserLogin_SendsPathParametersAndBody(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.tools["CreateUserIdentity"] = func(args map[string]any) (string, bool) {
		return `{"code":"IDENTITY_CREATED"}`, false
	}
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)
	loginThroughBrowser(t, cli)

	_, stderr, err := cli.run("users", "create-identity", "--user-id", "kp_1", "--type", "email", "--value", "sally@example.com")

	require.NoError(t, err, stderr)
	assert.Equal(t, map[string]any{"name": "CreateUserIdentity", "arguments": map[string]any{
		"user_id": "kp_1", "type": "email", "value": "sally@example.com",
	}}, kinde.lastToolCall())
}

func TestBrowserLogin_APIErrorsKeepTheirStatus(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.tools["GetUserData"] = func(map[string]any) (string, bool) {
		return `API request failed: API error (HTTP 403): {"errors": [{"code": "FORBIDDEN", "message": "Missing scope"}]}`, true
	}
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)
	loginThroughBrowser(t, cli)

	stdout, _, err := cli.run("users", "get", "--id", "kp_1")

	assert.EqualError(t, err, "the API returned 403 Forbidden: FORBIDDEN: Missing scope. This needs the read:users scope: an owner can raise your team role's access level in Kinde under Settings > APIs > Kinde Management API")
	assert.Empty(t, stdout)
}

func TestBrowserLogin_CommandsMCPDoesNotOfferPointToM2M(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.tools["GetUsers"] = func(map[string]any) (string, bool) { return `{}`, false }
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)
	loginThroughBrowser(t, cli)

	_, _, err := cli.run("billing", "list-agreements", "--customer-id", "c_1")

	assert.EqualError(t, err, "this command isn't available with a browser login. Log in with an M2M application to use it: run 'kinde login --client-id <client ID>'")
}

func TestBrowserLogin_SaysWhenSecretsAreHidden(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.tools["GetApplication"] = func(map[string]any) (string, bool) {
		return `{"application":{"id":"app_1","client_secret":"<redacted>"}}`, false
	}
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)
	loginThroughBrowser(t, cli)

	stdout, stderr, err := cli.run("applications", "get", "--application-id", "app_1")

	require.NoError(t, err)
	assert.Contains(t, stdout, `"<redacted>"`)
	assert.Contains(t, stderr, "Note: secrets are hidden with a browser login")
}

func TestBrowserLogin_RefreshesAnExpiredToken(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.tools["GetUsers"] = func(map[string]any) (string, bool) { return `{}`, false }
	routeKindeTo(t, kinde)
	cli := newTestCLI(t)
	loginThroughBrowser(t, cli)

	env := &config.Environment{DomainName: testDomain, Browser: cli.readConfig().Browser}
	session, err := loadBrowserSession(env)
	require.NoError(t, err)
	session.Token.Expiry = time.Now().Add(-time.Minute)
	require.NoError(t, saveBrowserSession(env, session))

	_, stderr, err := cli.run("users", "list")

	require.NoError(t, err, stderr)
	assert.Equal(t, 1, kinde.refreshes)
	session, err = loadBrowserSession(env)
	require.NoError(t, err)
	assert.Equal(t, "access-2", session.Token.AccessToken, "the refreshed token is saved")
}

func TestBrowserLogin_ExpiredLoginAsksToLogInAgain(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{Browser: &config.BrowserLogin{ClientID: "cli-client", Issuer: fakeIssuer}})
	env := &config.Environment{DomainName: testDomain, Browser: &config.BrowserLogin{ClientID: "cli-client"}}
	require.NoError(t, saveBrowserSession(env, &agentauth.Session{
		Server: agentauth.Server{Resource: "https://" + testDomain + "/mcp"},
		Token:  &oauth2.Token{AccessToken: "old", Expiry: time.Now().Add(-time.Hour)},
	}))

	_, _, err := cli.run("users", "list")

	assert.EqualError(t, err, "your browser login for "+testDomain+" has expired or been revoked: run 'kinde login' again")
}

func TestManagementCommands_PreferM2MOverBrowserLogin(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{
		M2M:     &config.M2MLogin{ClientID: "m2m-app"},
		Browser: &config.BrowserLogin{ClientID: "cli-client", Issuer: fakeIssuer},
	})

	status, _, err := runWhoAmI(t, cli)

	require.NoError(t, err)
	assert.Equal(t, ptr("m2m"), status.ManagementAPI)
}

func TestManagementCommands_NeedALogin(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{})

	_, _, err := cli.run("users", "list")

	assert.ErrorIs(t, err, config.ErrNoM2MLogin)
	assert.ErrorContains(t, err, "run 'kinde login' to sign in through the browser")
}

func TestNoDomain_ErrorExplainsHowToSetOne(t *testing.T) {
	cli := newTestCLI(t)

	for _, args := range [][]string{{"account", "login"}, {"login"}, {"users", "list"}} {
		_, _, err := cli.run(args...)
		assert.EqualError(t, err, "no Kinde domain configured: pass --domain <your Kinde domain> or set KINDE_DOMAIN", args)
	}
}

func readConfigFile(t *testing.T, cli *testCLI) string {
	t.Helper()
	cfg := cli.readFullConfig()
	data, err := json.Marshal(cfg)
	require.NoError(t, err)
	return string(data)
}
