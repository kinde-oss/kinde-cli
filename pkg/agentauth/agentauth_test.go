package agentauth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

// fakeKinde is an MCP server plus the sign-in server for it, behaving like
// Kinde's agent sign-in.
type fakeKinde struct {
	*httptest.Server
	t             *testing.T
	enabled       bool
	tokenEndpoint string // overrides the advertised token endpoint
	metadataURL   string // overrides the metadata URL in the 401 challenge
	redirectToken string // makes the token endpoint redirect here
	challenge     string // code_challenge from the last authorize request
	refreshed     int
	registrations []map[string]any
}

func newFakeKinde(t *testing.T) *fakeKinde {
	f := &fakeKinde{t: t, enabled: true}
	mux := http.NewServeMux()
	f.Server = httptest.NewServer(mux)
	t.Cleanup(f.Close)
	issuer := f.URL + "/oauth2/agents/app1"

	mux.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		if !f.enabled {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		metadataURL := f.URL + "/.well-known/oauth-protected-resource/mcp"
		if f.metadataURL != "" {
			metadataURL = f.metadataURL
		}
		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Bearer resource_metadata="%s"`, metadataURL))
		w.WriteHeader(http.StatusUnauthorized)
	})
	mux.HandleFunc("/.well-known/oauth-protected-resource/mcp", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"resource": f.URL + "/mcp", "authorization_servers": []string{issuer}})
	})
	mux.HandleFunc("/.well-known/oauth-authorization-server/oauth2/agents/app1", func(w http.ResponseWriter, r *http.Request) {
		tokenEndpoint := issuer + "/token"
		if f.tokenEndpoint != "" {
			tokenEndpoint = f.tokenEndpoint
		}
		json.NewEncoder(w).Encode(map[string]any{
			"issuer":                           issuer,
			"authorization_endpoint":           issuer + "/authorize",
			"token_endpoint":                   tokenEndpoint,
			"registration_endpoint":            issuer + "/register",
			"code_challenge_methods_supported": []string{"S256"},
		})
	})
	mux.HandleFunc("/oauth2/agents/app1/register", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		f.registrations = append(f.registrations, body)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{"client_id": "cli-client", "client_secret_expires_at": time.Now().Add(30 * 24 * time.Hour).Unix()})
	})
	mux.HandleFunc("/oauth2/agents/app1/authorize", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		assert.Equal(t, "cli-client", q.Get("client_id"))
		assert.Equal(t, "S256", q.Get("code_challenge_method"))
		assert.Equal(t, "offline", q.Get("scope"))
		assert.Equal(t, f.URL+"/mcp", q.Get("resource"))
		f.challenge = q.Get("code_challenge")
		redirect, _ := url.Parse(q.Get("redirect_uri"))
		redirect.RawQuery = url.Values{"code": {"the-code"}, "state": {q.Get("state")}}.Encode()
		http.Redirect(w, r, redirect.String(), http.StatusFound)
	})
	mux.HandleFunc("/oauth2/agents/app1/token", func(w http.ResponseWriter, r *http.Request) {
		if f.redirectToken != "" {
			http.Redirect(w, r, f.redirectToken, http.StatusTemporaryRedirect)
			return
		}
		require.NoError(t, r.ParseForm())
		assert.Equal(t, "cli-client", r.PostForm.Get("client_id"))
		w.Header().Set("Content-Type", "application/json")
		switch r.PostForm.Get("grant_type") {
		case "authorization_code":
			sum := sha256.Sum256([]byte(r.PostForm.Get("code_verifier")))
			if base64.RawURLEncoding.EncodeToString(sum[:]) != f.challenge {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid_grant"})
				return
			}
			assert.Equal(t, "the-code", r.PostForm.Get("code"))
			json.NewEncoder(w).Encode(map[string]any{"access_token": "access-1", "refresh_token": "refresh-1", "token_type": "bearer", "expires_in": 3600})
		case "refresh_token":
			f.refreshed++
			if r.PostForm.Get("refresh_token") != "refresh-1" {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid_grant"})
				return
			}
			// Kinde doesn't rotate the refresh token.
			json.NewEncoder(w).Encode(map[string]any{"access_token": "access-2", "token_type": "bearer", "expires_in": 3600})
		}
	})
	return f
}

// browser follows the sign-in URL like a browser would, including the
// redirect back to the CLI.
func browser(t *testing.T) func(string) {
	return func(signInURL string) {
		go func() {
			resp, err := http.Get(signInURL)
			if assert.NoError(t, err) {
				resp.Body.Close()
			}
		}()
	}
}

func TestDiscover(t *testing.T) {
	kinde := newFakeKinde(t)

	server, err := Discover(context.Background(), http.DefaultClient, kinde.URL+"/mcp")

	require.NoError(t, err)
	assert.Equal(t, &Server{
		Resource:              kinde.URL + "/mcp",
		Issuer:                kinde.URL + "/oauth2/agents/app1",
		AuthorizationEndpoint: kinde.URL + "/oauth2/agents/app1/authorize",
		TokenEndpoint:         kinde.URL + "/oauth2/agents/app1/token",
		RegistrationEndpoint:  kinde.URL + "/oauth2/agents/app1/register",
	}, server)
}

func TestDiscover_SignInTurnedOff(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.enabled = false

	_, err := Discover(context.Background(), http.DefaultClient, kinde.URL+"/mcp")

	assert.ErrorIs(t, err, ErrSignInNotEnabled)
}

func TestDiscover_RejectsEndpointsWithoutHTTPS(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.tokenEndpoint = "http://kinde.example.com/token"

	_, err := Discover(context.Background(), http.DefaultClient, kinde.URL+"/mcp")

	assert.EqualError(t, err, `the sign-in server's endpoint "http://kinde.example.com/token" doesn't use HTTPS`)
}

func TestDiscover_RejectsMetadataWithoutHTTPS(t *testing.T) {
	kinde := newFakeKinde(t)
	kinde.metadataURL = "http://kinde.example.com/.well-known/oauth-protected-resource/mcp"

	_, err := Discover(context.Background(), http.DefaultClient, kinde.URL+"/mcp")

	assert.EqualError(t, err, `the protected resource metadata "http://kinde.example.com/.well-known/oauth-protected-resource/mcp" doesn't use HTTPS`)
}

func TestRefresh_RefusesARedirectToHTTP(t *testing.T) {
	kinde := newFakeKinde(t)
	server, err := Discover(context.Background(), http.DefaultClient, kinde.URL+"/mcp")
	require.NoError(t, err)
	kinde.redirectToken = "http://kinde.example.com/token"

	session := &Session{Server: *server, Registration: Registration{ClientID: "cli-client"},
		Token: &oauth2.Token{RefreshToken: "refresh-1", Expiry: time.Now().Add(-time.Minute)}}
	_, err = Refresh(context.Background(), http.DefaultClient, session)

	assert.ErrorContains(t, err, "refusing to follow a redirect to http://kinde.example.com/token: it doesn't use HTTPS")
}

func TestSecureURL(t *testing.T) {
	assert.True(t, secureURL("https://app.kinde.com/oauth2/x/token"))
	assert.True(t, secureURL("http://127.0.0.1:8080/token"), "this machine is allowed for tests")
	assert.True(t, secureURL("http://localhost/token"))
	assert.False(t, secureURL("http://app.kinde.com/token"))
	assert.False(t, secureURL("ftp://app.kinde.com/token"))
	assert.False(t, secureURL("not a url"))
}

func TestRegister_UsesAPortlessLoopbackRedirect(t *testing.T) {
	kinde := newFakeKinde(t)
	server, err := Discover(context.Background(), http.DefaultClient, kinde.URL+"/mcp")
	require.NoError(t, err)

	registration, err := Register(context.Background(), http.DefaultClient, server)

	require.NoError(t, err)
	assert.Equal(t, "cli-client", registration.ClientID)
	assert.True(t, registration.Usable(time.Now()))
	require.Len(t, kinde.registrations, 1)
	assert.Equal(t, []any{"http://127.0.0.1/callback"}, kinde.registrations[0]["redirect_uris"])
	assert.Equal(t, "none", kinde.registrations[0]["token_endpoint_auth_method"])
}

func TestAuthorizeAndRefresh(t *testing.T) {
	kinde := newFakeKinde(t)
	ctx := context.Background()
	server, err := Discover(ctx, http.DefaultClient, kinde.URL+"/mcp")
	require.NoError(t, err)
	registration, err := Register(ctx, http.DefaultClient, server)
	require.NoError(t, err)

	token, err := Authorize(ctx, http.DefaultClient, server, registration, browser(t))
	require.NoError(t, err)
	assert.Equal(t, "access-1", token.AccessToken)
	assert.Equal(t, "refresh-1", token.RefreshToken)

	session := &Session{Server: *server, Registration: *registration, Token: token}
	changed, err := Refresh(ctx, http.DefaultClient, session)
	require.NoError(t, err)
	assert.False(t, changed, "a valid token isn't refreshed")

	session.Token.Expiry = time.Now().Add(-time.Minute)
	changed, err = Refresh(ctx, http.DefaultClient, session)
	require.NoError(t, err)
	assert.True(t, changed)
	assert.Equal(t, "access-2", session.Token.AccessToken)
	assert.Equal(t, "refresh-1", session.Token.RefreshToken, "a refresh token that wasn't rotated is kept")

	session.Token = &oauth2.Token{RefreshToken: "revoked", Expiry: time.Now().Add(-time.Minute)}
	_, err = Refresh(ctx, http.DefaultClient, session)
	assert.True(t, IsInvalidGrant(err), "%v", err)
}

func TestAuthorize_ReportsSignInErrors(t *testing.T) {
	kinde := newFakeKinde(t)
	server := &Server{Resource: kinde.URL + "/mcp", AuthorizationEndpoint: kinde.URL + "/denied", TokenEndpoint: kinde.URL + "/token"}

	deny := func(signInURL string) {
		u, _ := url.Parse(signInURL)
		q := u.Query()
		redirect := q.Get("redirect_uri") + "?" + url.Values{"error": {"access_denied"}, "error_description": {"The user denied the request"}, "state": {q.Get("state")}}.Encode()
		go http.Get(redirect)
	}
	_, err := Authorize(context.Background(), http.DefaultClient, server, &Registration{ClientID: "c"}, deny)

	assert.EqualError(t, err, "sign-in failed: access_denied: The user denied the request")
}

func TestCovers(t *testing.T) {
	assert.True(t, covers("https://a.kinde.com/mcp", "https://a.kinde.com/mcp"))
	assert.True(t, covers("https://a.kinde.com/", "https://a.kinde.com/mcp"))
	assert.False(t, covers("https://a.kinde.com/mcp", "https://b.kinde.com/mcp"))
	assert.False(t, covers("https://a.kinde.com/mcp", "https://a.kinde.com/mcpx"))
	assert.False(t, covers("", "https://a.kinde.com/mcp"))
}

func TestResourceMetadataURL(t *testing.T) {
	assert.Equal(t, "https://x/prm", resourceMetadataURL([]string{`Bearer error="invalid_token", resource_metadata="https://x/prm"`}))
	assert.Equal(t, "", resourceMetadataURL([]string{`Bearer realm="x"`}))
}
