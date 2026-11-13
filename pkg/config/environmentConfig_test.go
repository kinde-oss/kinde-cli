package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func useTestSessionStore(t *testing.T) {
	t.Helper()
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv(EnvM2MClientID, "")
	t.Setenv(EnvM2MClientSecret, "")
	t.Cleanup(UseFileSessionStore("test-password"))
}

func TestM2MCredentialsFromEnv_SkipSessionStore(t *testing.T) {
	useTestSessionStore(t)
	t.Setenv(EnvM2MClientID, "abc")
	t.Setenv(EnvM2MClientSecret, "s3cret")
	env := &Environment{DomainName: "mybiz.kinde.com"}

	_, err := env.M2MClientCredentialsFlow()
	require.NoError(t, err, "environment credentials should work without a saved login")

	storeDir := filepath.Join(os.Getenv("XDG_CONFIG_HOME"), CLI_NAME, "kc_mybiz_kinde_com_m2m")
	_, err = os.Stat(storeDir)
	assert.True(t, os.IsNotExist(err), "environment credentials must not open the session store")
}

func TestIdentitiesHaveSeparateSessionStores(t *testing.T) {
	useTestSessionStore(t)
	env := &Environment{DomainName: "mybiz.kinde.com"}

	userSession, err := env.getCliSession(IdentityUser)
	require.NoError(t, err)
	require.NoError(t, userSession.SetRawToken(&oauth2.Token{AccessToken: "user-token"}))

	m2mSession, err := env.getCliSession(IdentityM2M)
	require.NoError(t, err)
	_, err = m2mSession.GetRawToken()
	assert.Error(t, err, "a user login token must not be visible to the M2M login")
}

func TestM2MSecretSaveAndClear(t *testing.T) {
	useTestSessionStore(t)
	env := &Environment{DomainName: "mybiz.kinde.com", M2M: &M2MLogin{ClientID: "abc"}}

	_, err := env.M2MClientCredentialsFlow()
	assert.ErrorIs(t, err, ErrNoM2MLogin, "no secret saved yet")

	require.NoError(t, env.SaveM2MSecret("s3cret"))
	_, err = env.M2MClientCredentialsFlow()
	require.NoError(t, err)

	require.NoError(t, env.ClearM2MLogin())
	_, err = env.M2MClientCredentialsFlow()
	assert.ErrorIs(t, err, ErrNoM2MLogin, "secret should be gone after clearing")
}

func TestSaveM2MSecretClearsPreviousToken(t *testing.T) {
	useTestSessionStore(t)
	env := &Environment{DomainName: "mybiz.kinde.com"}

	session, err := env.getCliSession(IdentityM2M)
	require.NoError(t, err)
	require.NoError(t, session.SetRawToken(&oauth2.Token{AccessToken: "old-app-token"}))

	require.NoError(t, env.SaveM2MSecret("new-secret"))

	_, err = session.GetRawToken()
	assert.Error(t, err, "a token cached for a previous M2M app must not be reused")
}

func TestUserFlowRequiresUserLogin(t *testing.T) {
	useTestSessionStore(t)
	env := &Environment{DomainName: "mybiz.kinde.com", M2M: &M2MLogin{ClientID: "abc"}}

	_, err := env.UserDeviceAuthorizationFlow()
	assert.ErrorIs(t, err, ErrNoUserLogin)
}

func TestBrowserSessionSaveAndClear(t *testing.T) {
	useTestSessionStore(t)
	env := &Environment{DomainName: "mybiz.kinde.com"}

	_, err := env.BrowserSession()
	assert.ErrorIs(t, err, ErrNoBrowserLogin, "no browser login in the config")

	env.Browser = &BrowserLogin{ClientID: "cli", Issuer: "https://app.kinde.com/oauth2/x"}
	_, err = env.BrowserSession()
	assert.ErrorIs(t, err, ErrNoBrowserLogin, "nothing saved yet")

	require.NoError(t, env.SaveBrowserSession([]byte(`{"token":"t"}`)))
	session, err := env.BrowserSession()
	require.NoError(t, err)
	assert.Equal(t, `{"token":"t"}`, string(session))

	m2mSession, err := env.getCliSession(IdentityM2M)
	require.NoError(t, err)
	_, err = m2mSession.GetKey(browserSessionKey)
	assert.Error(t, err, "the browser login must not be visible to the M2M login")

	require.NoError(t, env.ClearBrowserLogin())
	_, err = env.BrowserSession()
	assert.ErrorIs(t, err, ErrNoBrowserLogin, "session should be gone after clearing")
}

func TestEnvironmentJSONHasNoSecrets(t *testing.T) {
	env := Environment{
		DomainName: "mybiz.kinde.com",
		M2M:        &M2MLogin{ClientID: "m2m-app"},
		Browser:    &BrowserLogin{ClientID: "cli", Issuer: "https://app.kinde.com/oauth2/x"},
		User:       &UserLogin{ClientID: "user-app"},
	}
	buf, err := json.Marshal(env)
	require.NoError(t, err)
	assert.JSONEq(t,
		`{"domain_name":"mybiz.kinde.com","m2m":{"client_id":"m2m-app"},"browser":{"client_id":"cli","issuer":"https://app.kinde.com/oauth2/x"},"user":{"client_id":"user-app"}}`,
		string(buf))
}

// storeDir is where the file-based session store keeps an identity's keys.
func storeDir(identity Identity) string {
	return filepath.Join(os.Getenv("XDG_CONFIG_HOME"), CLI_NAME, "kc_mybiz_kinde_com_"+string(identity))
}

// corruptStore overwrites the stored data so it can be found but not read.
// The chunk count is left alone: kinde-go treats an unreadable count as an
// unchunked item, which is then missing, so that case reads as not found.
func corruptStore(t *testing.T, identity Identity) {
	t.Helper()
	entries, err := os.ReadDir(storeDir(identity))
	require.NoError(t, err)
	corrupted := 0
	for _, entry := range entries {
		if strings.Contains(entry.Name(), "chunk_count") {
			continue
		}
		require.NoError(t, os.WriteFile(filepath.Join(storeDir(identity), entry.Name()), []byte("not a keyring item"), 0600))
		corrupted++
	}
	require.NotZero(t, corrupted)
}

func TestUnreadableKeyring_IsNotReportedAsLoggedOut(t *testing.T) {
	useTestSessionStore(t)
	env := &Environment{
		DomainName: "mybiz.kinde.com",
		M2M:        &M2MLogin{ClientID: "abc"},
		Browser:    &BrowserLogin{ClientID: "cli"},
	}
	require.NoError(t, env.SaveM2MSecret("s3cret"))
	require.NoError(t, env.SaveBrowserSession([]byte(`{"token":"t"}`)))
	corruptStore(t, IdentityM2M)
	corruptStore(t, IdentityBrowser)

	_, err := env.M2MClientCredentialsFlow()
	assert.Error(t, err)
	assert.NotErrorIs(t, err, ErrNoM2MLogin, "a keyring read failure must not make commands fall back to another login")

	_, err = env.BrowserSession()
	assert.Error(t, err)
	assert.NotErrorIs(t, err, ErrNoBrowserLogin)
}

func TestClearLogins_FailWhenTheSecretCannotBeRemoved(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root can remove files from a read-only directory")
	}
	useTestSessionStore(t)
	env := &Environment{DomainName: "mybiz.kinde.com"}
	require.NoError(t, env.SaveM2MSecret("s3cret"))
	require.NoError(t, env.SaveBrowserSession([]byte(`{"token":"t"}`)))

	// Locked: the removal fails, and so does reading the item back, so it
	// can't be confirmed gone.
	for _, identity := range []Identity{IdentityM2M, IdentityBrowser} {
		require.NoError(t, os.Chmod(storeDir(identity), 0))
		t.Cleanup(func() { os.Chmod(storeDir(identity), 0700) })
	}

	assert.Error(t, env.ClearM2MLogin(), "logout must not report success while the secret may still be stored")
	assert.Error(t, env.ClearBrowserLogin())
}

func TestClearLogins_SucceedWhenNothingIsStored(t *testing.T) {
	useTestSessionStore(t)
	env := &Environment{DomainName: "mybiz.kinde.com"}

	assert.NoError(t, env.ClearM2MLogin())
	assert.NoError(t, env.ClearBrowserLogin())
	assert.NoError(t, env.ClearUserLogin())
}
