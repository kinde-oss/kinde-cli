package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func useTestSessionStore(t *testing.T) {
	t.Helper()
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Cleanup(UseFileSessionStore("test-password"))
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
	assert.Error(t, err, "a browser login token must not be visible to the M2M login")
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

func TestEnvironmentJSONHasNoSecrets(t *testing.T) {
	env := Environment{
		DomainName: "mybiz.kinde.com",
		M2M:        &M2MLogin{ClientID: "m2m-app"},
		User:       &UserLogin{ClientID: "user-app"},
	}
	buf, err := json.Marshal(env)
	require.NoError(t, err)
	assert.JSONEq(t,
		`{"domain_name":"mybiz.kinde.com","m2m":{"client_id":"m2m-app"},"user":{"client_id":"user-app"}}`,
		string(buf))
}
