// Package agentauth signs a Kinde team member in through the browser using
// Kinde's "AI agent sign-in" for the Management API, the same flow MCP clients
// such as Claude use. The resulting token is for the domain's MCP server
// (https://<domain>/mcp), which calls the Management API on the user's behalf.
//
// The flow follows the MCP authorization spec: discover the sign-in server
// from the MCP server's 401 challenge, register with Dynamic Client
// Registration (RFC 7591), then sign in with the authorization code flow and
// PKCE, receiving the code on a loopback redirect (RFC 8252).
package agentauth

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// ClientName is how the CLI names itself when it registers.
const ClientName = "Kinde CLI"

// registeredRedirectURI is registered without a port; the port is chosen when
// signing in (RFC 8252 section 7.3). Kinde refuses loopback URIs with a port.
const registeredRedirectURI = "http://127.0.0.1/callback"

// scope asks for a refresh token. Kinde's agent sign-in doesn't need openid.
const scope = "offline"

// loginTimeout is how long Authorize waits for the browser sign-in.
const loginTimeout = 5 * time.Minute

// ErrSignInNotEnabled means the domain doesn't offer agent sign-in, usually
// because "AI agent sign-in" is turned off for the Kinde Management API.
var ErrSignInNotEnabled = errors.New("browser sign-in isn't turned on")

// Server describes the MCP server and the Kinde sign-in server for it.
type Server struct {
	Resource              string `json:"resource"`
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	RegistrationEndpoint  string `json:"registration_endpoint"`
}

// Registration is the client the CLI registered with the sign-in server.
type Registration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"`
	// ExpiresAt is when the registration expires, in Unix seconds, or 0 if
	// it doesn't.
	ExpiresAt int64 `json:"client_secret_expires_at,omitempty"`
}

// Usable reports whether the registration can still be used for a while.
func (r *Registration) Usable(now time.Time) bool {
	if r == nil || r.ClientID == "" {
		return false
	}
	return r.ExpiresAt == 0 || time.Unix(r.ExpiresAt, 0).After(now.Add(24*time.Hour))
}

// Session is everything needed to call the MCP server and keep the login
// alive. It's kept in the OS keychain.
type Session struct {
	// MCPURL is the MCP server the login is for. It can differ from
	// Server.Resource, which may name a parent of it.
	MCPURL       string        `json:"mcp_url"`
	Server       Server        `json:"server"`
	Registration Registration  `json:"registration"`
	Token        *oauth2.Token `json:"token"`
}

// Discover finds the sign-in server for an MCP server by calling it without
// a token and following its 401 challenge (RFC 9728, then RFC 8414).
func Discover(ctx context.Context, client *http.Client, mcpURL string) (*Server, error) {
	client = secureClient(client)
	if !secureURL(mcpURL) {
		return nil, fmt.Errorf("invalid MCP server %q: it must use HTTPS", mcpURL)
	}
	resp, err := postJSON(ctx, client, mcpURL, initializeRequest, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to reach %s: %w", mcpURL, err)
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		return nil, fmt.Errorf("%w (%s answered %s instead of asking to sign in)", ErrSignInNotEnabled, mcpURL, resp.Status)
	}

	metadataURL := resourceMetadataURL(resp.Header.Values("WWW-Authenticate"))
	if metadataURL == "" {
		u, _ := url.Parse(mcpURL)
		metadataURL = fmt.Sprintf("%s://%s/.well-known/oauth-protected-resource%s", u.Scheme, u.Host, u.Path)
	}
	// The metadata names the sign-in server, so it must come over TLS, or
	// anyone on the network could send the login somewhere else (RFC 9728).
	if !secureURL(metadataURL) {
		return nil, fmt.Errorf("the protected resource metadata %q doesn't use HTTPS", metadataURL)
	}
	var resource struct {
		Resource             string   `json:"resource"`
		AuthorizationServers []string `json:"authorization_servers"`
	}
	if err := getJSON(ctx, client, metadataURL, &resource); err != nil {
		return nil, fmt.Errorf("%w (no protected resource metadata: %v)", ErrSignInNotEnabled, err)
	}
	if !covers(resource.Resource, mcpURL) {
		return nil, fmt.Errorf("the protected resource metadata is for %q, not %s", resource.Resource, mcpURL)
	}
	if len(resource.AuthorizationServers) == 0 {
		return nil, fmt.Errorf("%w (no sign-in server listed for %s)", ErrSignInNotEnabled, mcpURL)
	}

	issuer := resource.AuthorizationServers[0]
	if !secureURL(issuer) {
		return nil, fmt.Errorf("invalid sign-in server %q: it must use HTTPS", issuer)
	}
	u, _ := url.Parse(issuer)
	// RFC 8414 path insertion: the well-known segment goes before the path.
	discoveryURL := fmt.Sprintf("%s://%s/.well-known/oauth-authorization-server%s", u.Scheme, u.Host, strings.TrimSuffix(u.Path, "/"))
	var metadata struct {
		Issuer                        string   `json:"issuer"`
		AuthorizationEndpoint         string   `json:"authorization_endpoint"`
		TokenEndpoint                 string   `json:"token_endpoint"`
		RegistrationEndpoint          string   `json:"registration_endpoint"`
		CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported"`
	}
	if err := getJSON(ctx, client, discoveryURL, &metadata); err != nil {
		return nil, fmt.Errorf("failed to read the sign-in server's metadata: %w", err)
	}
	switch {
	case metadata.Issuer != issuer:
		return nil, fmt.Errorf("the sign-in server's metadata is for %q, not %q", metadata.Issuer, issuer)
	case !slices.Contains(metadata.CodeChallengeMethodsSupported, "S256"):
		return nil, errors.New("the sign-in server doesn't support PKCE with S256")
	case metadata.AuthorizationEndpoint == "" || metadata.TokenEndpoint == "":
		return nil, errors.New("the sign-in server's metadata has no authorization or token endpoint")
	}
	// The code, PKCE verifier and refresh token go to these endpoints, so
	// they must be protected in transit.
	for _, endpoint := range []string{metadata.AuthorizationEndpoint, metadata.TokenEndpoint, metadata.RegistrationEndpoint} {
		if endpoint != "" && !secureURL(endpoint) {
			return nil, fmt.Errorf("the sign-in server's endpoint %q doesn't use HTTPS", endpoint)
		}
	}

	return &Server{
		Resource:              resource.Resource,
		Issuer:                metadata.Issuer,
		AuthorizationEndpoint: metadata.AuthorizationEndpoint,
		TokenEndpoint:         metadata.TokenEndpoint,
		RegistrationEndpoint:  metadata.RegistrationEndpoint,
	}, nil
}

// Register registers the CLI as a public client (RFC 7591).
func Register(ctx context.Context, client *http.Client, server *Server) (*Registration, error) {
	client = secureClient(client)
	if server.RegistrationEndpoint == "" {
		return nil, errors.New("the sign-in server doesn't accept client registration")
	}
	resp, err := postJSON(ctx, client, server.RegistrationEndpoint, map[string]any{
		"client_name":                ClientName,
		"redirect_uris":              []string{registeredRedirectURI},
		"grant_types":                []string{"authorization_code", "refresh_token"},
		"response_types":             []string{"code"},
		"token_endpoint_auth_method": "none",
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to register with the sign-in server: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to register with the sign-in server: %s", oauthError(resp.Status, body))
	}
	var registration Registration
	if err := json.Unmarshal(body, &registration); err != nil || registration.ClientID == "" {
		return nil, fmt.Errorf("failed to register with the sign-in server: unexpected response %q", truncate(body))
	}
	return &registration, nil
}

// Authorize signs the user in through the browser and returns their token.
// open is called with the sign-in URL; it should show the URL and try to
// open the browser.
func Authorize(ctx context.Context, client *http.Client, server *Server, registration *Registration, open func(url string)) (*oauth2.Token, error) {
	client = secureClient(client)
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen for the sign-in redirect: %w", err)
	}
	defer listener.Close()
	redirectURI := fmt.Sprintf("http://%s/callback", listener.Addr())

	config := oauthConfig(server, registration, redirectURI)
	verifier := oauth2.GenerateVerifier()
	state, err := randomString()
	if err != nil {
		return nil, err
	}

	type callback struct {
		code string
		err  error
	}
	results := make(chan callback, 1)
	httpServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/callback" {
				http.NotFound(w, r)
				return
			}
			query := r.URL.Query()
			var result callback
			switch {
			case query.Get("state") != state:
				result.err = errors.New("the sign-in response didn't match this login; try again")
			case query.Get("error") != "":
				result.err = fmt.Errorf("sign-in failed: %s", strings.TrimSpace(query.Get("error")+": "+query.Get("error_description")))
			case query.Get("code") == "":
				result.err = errors.New("sign-in failed: no authorization code was returned")
			default:
				result.code = query.Get("code")
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if result.err != nil {
				fmt.Fprint(w, page("Sign-in failed", "Go back to the terminal for details."))
			} else {
				fmt.Fprint(w, page("You're signed in", "You can close this tab and go back to the terminal."))
			}
			select {
			case results <- result:
			default:
			}
		}),
	}
	go httpServer.Serve(listener)
	defer httpServer.Close()

	open(config.AuthCodeURL(state,
		oauth2.S256ChallengeOption(verifier),
		oauth2.SetAuthURLParam("resource", server.Resource),
	))

	ctx, cancel := context.WithTimeout(ctx, loginTimeout)
	defer cancel()
	var result callback
	select {
	case result = <-results:
	case <-ctx.Done():
		return nil, errors.New("timed out waiting for you to sign in in the browser")
	}
	if result.err != nil {
		return nil, result.err
	}

	token, err := config.Exchange(context.WithValue(ctx, oauth2.HTTPClient, client), result.code,
		oauth2.VerifierOption(verifier),
		oauth2.SetAuthURLParam("resource", server.Resource),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to complete the sign-in: %w", err)
	}
	return token, nil
}

// Refresh returns a valid access token for the session, refreshing it if it
// has expired. It reports whether the token changed, so the caller can save it.
func Refresh(ctx context.Context, client *http.Client, session *Session) (changed bool, err error) {
	if session.Token == nil {
		return false, errors.New("no token")
	}
	if session.Token.Valid() {
		return false, nil
	}
	if session.Token.RefreshToken == "" {
		return false, errors.New("the login has expired")
	}
	config := oauthConfig(&session.Server, &session.Registration, "")
	ctx = context.WithValue(ctx, oauth2.HTTPClient, secureClient(client))
	token, err := config.TokenSource(ctx, session.Token).Token()
	if err != nil {
		return false, err
	}
	// Kinde may or may not rotate the refresh token; keep the old one if not.
	if token.RefreshToken == "" {
		token.RefreshToken = session.Token.RefreshToken
	}
	session.Token = token
	return true, nil
}

// IsInvalidGrant reports whether err means the refresh token or client is no
// longer accepted, so the user has to sign in again.
func IsInvalidGrant(err error) bool {
	var retrieveErr *oauth2.RetrieveError
	if errors.As(err, &retrieveErr) {
		switch retrieveErr.ErrorCode {
		case "invalid_grant", "invalid_client", "unauthorized_client":
			return true
		}
	}
	return false
}

func oauthConfig(server *Server, registration *Registration, redirectURI string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     registration.ClientID,
		ClientSecret: registration.ClientSecret,
		RedirectURL:  redirectURI,
		Scopes:       []string{scope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  server.AuthorizationEndpoint,
			TokenURL: server.TokenEndpoint,
			// Public clients send client_id (and any secret) in the body.
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
}

// initializeRequest is a minimal MCP initialize request, used to get the
// server's 401 challenge.
var initializeRequest = map[string]any{
	"jsonrpc": "2.0",
	"id":      1,
	"method":  "initialize",
	"params": map[string]any{
		"protocolVersion": "2025-06-18",
		"capabilities":    map[string]any{},
		"clientInfo":      map[string]any{"name": "kinde-cli", "version": "1"},
	},
}

// resourceMetadataURL reads resource_metadata from a Bearer challenge.
func resourceMetadataURL(challenges []string) string {
	pattern := regexp.MustCompile(`(?i)resource_metadata\s*=\s*"([^"]*)"`)
	for _, challenge := range challenges {
		if match := pattern.FindStringSubmatch(challenge); match != nil {
			return match[1]
		}
	}
	return ""
}

// covers reports whether the resource identifier applies to target: same
// origin, and target's path is the resource's path or under it.
func covers(resource, target string) bool {
	r, err1 := url.Parse(resource)
	t, err2 := url.Parse(target)
	if err1 != nil || err2 != nil || resource == "" {
		return false
	}
	if !strings.EqualFold(r.Scheme, t.Scheme) || !strings.EqualFold(r.Host, t.Host) {
		return false
	}
	rp, tp := strings.TrimSuffix(r.Path, "/"), strings.TrimSuffix(t.Path, "/")
	return tp == rp || strings.HasPrefix(tp, rp+"/")
}

// secureClient returns a copy of client that refuses to follow a redirect to
// an address secureURL rejects. Token requests are POSTs whose body holds the
// code, PKCE verifier or refresh token, and Go follows 307 and 308 redirects
// with the body.
func secureClient(client *http.Client) *http.Client {
	secured := *client
	previous := client.CheckRedirect
	secured.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if !secureURL(req.URL.String()) {
			return fmt.Errorf("refusing to follow a redirect to %s: it doesn't use HTTPS", req.URL.Redacted())
		}
		if previous != nil {
			return previous(req, via)
		}
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		return nil
	}
	return &secured
}

// secureURL reports whether credentials can be sent to rawURL: HTTPS, or
// plain HTTP to this machine.
func secureURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" {
		return false
	}
	return u.Scheme == "https" || u.Scheme == "http" && isLoopback(u.Hostname())
}

func isLoopback(host string) bool {
	ip := net.ParseIP(host)
	return host == "localhost" || ip != nil && ip.IsLoopback()
}

func getJSON(ctx context.Context, client *http.Client, url string, v any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s returned %s", url, resp.Status)
	}
	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("%s didn't return JSON", url)
	}
	return nil
}

func postJSON(ctx context.Context, client *http.Client, url string, body any, header http.Header) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	for key, values := range header {
		req.Header[key] = values
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	return client.Do(req)
}

// oauthError formats an OAuth error response for a message.
func oauthError(status string, body []byte) string {
	var parsed struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
	if json.Unmarshal(body, &parsed) == nil && parsed.Error != "" {
		return strings.TrimSuffix(parsed.Error+": "+parsed.ErrorDescription, ": ")
	}
	return status
}

func truncate(body []byte) string {
	if len(body) > 200 {
		return string(body[:200]) + "..."
	}
	return string(body)
}

func randomString() (string, error) {
	b := make([]byte, 24)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func page(title, message string) string {
	return fmt.Sprintf(`<!doctype html><html><head><meta charset="utf-8"><title>%[1]s</title></head>
<body style="font-family:system-ui,sans-serif;max-width:32rem;margin:4rem auto;padding:0 1rem">
<h1>%[1]s</h1><p>%[2]s</p></body></html>`, title, message)
}
