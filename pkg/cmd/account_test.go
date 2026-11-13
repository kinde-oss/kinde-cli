package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountOperations_AreAllInTheSpec(t *testing.T) {
	commands, err := newAccountAPICommands()
	require.NoError(t, err)
	assert.Len(t, commands, len(accountOperations))
}

// accountURL builds the request URL for an account command with flags.
func accountURL(t *testing.T, operationID string, flags ...string) (string, error) {
	t.Helper()
	doc := accountDocsFor(operationID)
	command := newAccountAPICommand("test", doc)
	require.NoError(t, command.ParseFlags(flags))
	return accountRequestURL(doc, accountFields(doc), command.Flags())
}

func TestAccountRequestURL(t *testing.T) {
	got, err := accountURL(t, "GetUserPermissions", "--page-size", "5", "--starting-after", "perm_1")
	require.NoError(t, err)
	assert.Equal(t, "/account_api/v1/permissions?page_size=5&starting_after=perm_1", got)

	got, err = accountURL(t, "GetUserRoles")
	require.NoError(t, err)
	assert.Equal(t, "/account_api/v1/roles", got)
}

func TestAccountRequestURL_EntitlementKeyIsInThePath(t *testing.T) {
	got, err := accountURL(t, "GetEntitlement", "--key", "seats/extra")
	require.NoError(t, err)
	assert.Equal(t, "/account_api/v1/entitlement/seats%2Fextra", got)

	_, err = accountURL(t, "GetEntitlement")
	assert.EqualError(t, err, "missing required flags: --key")
}

func TestAccountRequestURL_ChecksAllowedValues(t *testing.T) {
	_, err := accountURL(t, "GetPortalLink", "--subnav", "billing")
	assert.ErrorContains(t, err, `invalid value for --subnav: "billing" is not one of: profile`)

	got, err := accountURL(t, "GetPortalLink", "--subnav", "plan_details", "--return-url", "https://app.example.com/?a=b")
	require.NoError(t, err)
	assert.Equal(t, "/account_api/v1/portal_link?return_url=https%3A%2F%2Fapp.example.com%2F%3Fa%3Db&subnav=plan_details", got)
}

func callFakeAccountAPI(t *testing.T, status int, body string) (*http.Request, string, error) {
	t.Helper()
	var received *http.Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = r
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	var out bytes.Buffer
	cmd := &cobra.Command{}
	cmd.SetOut(&out)
	cmd.SetContext(context.Background())
	err := callAccountAPI(cmd, server.Client(), server.URL+"/account_api/v1/roles")
	return received, out.String(), err
}

func TestCallAccountAPI_PrintsResponse(t *testing.T) {
	req, out, err := callFakeAccountAPI(t, http.StatusOK, `{"data":{"roles":[{"key":"admin"}]}}`)
	require.NoError(t, err)

	assert.Equal(t, "/account_api/v1/roles", req.URL.Path)
	assert.Equal(t, "kinde-cli/dev", req.Header.Get("User-Agent"))
	assert.JSONEq(t, `{"data":{"roles":[{"key":"admin"}]}}`, out)
}

func TestCallAccountAPI_UnauthorizedPointsToAccountLogin(t *testing.T) {
	_, out, err := callFakeAccountAPI(t, http.StatusUnauthorized, "")
	assert.EqualError(t, err, "the API returned 401 Unauthorized. Your login may have expired or been revoked: run 'kinde account login' again")
	assert.Empty(t, out)
}

func TestAccountCommands_NeedAUserLogin(t *testing.T) {
	for _, args := range [][]string{{"account", "roles"}, {"account", "token"}} {
		cli := newTestCLI(t)
		cli.writeConfig(config.Environment{M2M: &config.M2MLogin{ClientID: "abc"}})

		_, _, err := cli.run(args...)

		assert.ErrorIs(t, err, config.ErrNoUserLogin, args)
		assert.ErrorContains(t, err, "run 'kinde account login'", "an M2M login must not be used for account commands")
	}
}

func TestAccountCommand_InvalidFlagsFailBeforeLoginCheck(t *testing.T) {
	_, _, err := newTestCLI(t).run("account", "entitlement", "--domain", testDomain)
	assert.EqualError(t, err, "missing required flags: --key")
}

func runWhoAmI(t *testing.T, cli *testCLI) (whoAmIStatus, string, error) {
	t.Helper()
	stdout, stderr, err := cli.run("whoami")
	var status whoAmIStatus
	if err == nil {
		require.NoError(t, json.Unmarshal([]byte(stdout), &status), stdout)
	}
	return status, stderr, err
}

func TestWhoami_ShowsM2MLogin(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{M2M: &config.M2MLogin{ClientID: "abc"}})

	status, stderr, err := runWhoAmI(t, cli)

	require.NoError(t, err)
	assert.Equal(t, whoAmIStatus{Domain: testDomain, ManagementAPI: ptr("m2m"), M2M: &m2mStatus{ClientID: "abc", Source: "login"}}, status)
	assert.Empty(t, stderr)
}

func TestWhoami_EnvironmentVariablesTakePrecedence(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{M2M: &config.M2MLogin{ClientID: "abc"}})
	t.Setenv(config.EnvM2MClientID, "from-env")
	t.Setenv(config.EnvM2MClientSecret, "s3cret")

	status, _, err := runWhoAmI(t, cli)

	require.NoError(t, err)
	assert.Equal(t, &m2mStatus{ClientID: "from-env", Source: "environment"}, status.M2M)
}

func TestWhoami_InvalidUserLoginIsReported(t *testing.T) {
	cli := newTestCLI(t)
	// A user login in the config, but no token in the session store.
	cli.writeConfig(config.Environment{M2M: &config.M2MLogin{ClientID: "abc"}, User: &config.UserLogin{}})

	status, stderr, err := runWhoAmI(t, cli)

	require.NoError(t, err)
	assert.Nil(t, status.User)
	assert.Contains(t, stderr, "run 'kinde account login' again")
}

func TestWhoami_FailsWhenNotLoggedIn(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{})

	_, _, err := runWhoAmI(t, cli)

	assert.ErrorContains(t, err, "not logged in to "+testDomain)
}
