package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDomain = "mybiz.kinde.com"

// writeConfig saves a config file with testDomain selected and the given environment.
func (c *testCLI) writeConfig(env config.Environment) {
	c.t.Helper()
	env.DomainName = testDomain
	buf, err := json.Marshal(config.Config{
		CurrentEnvironment: testDomain,
		Environments:       map[string]config.Environment{testDomain: env},
	})
	require.NoError(c.t, err)
	path := filepath.Join(c.configDir, "kinde", "config.json")
	require.NoError(c.t, os.MkdirAll(filepath.Dir(path), 0700))
	require.NoError(c.t, os.WriteFile(path, buf, 0600))
}

// readConfig returns the saved environment for testDomain.
func (c *testCLI) readConfig() config.Environment {
	c.t.Helper()
	return c.readFullConfig().Environments[testDomain]
}

// readFullConfig returns the whole saved config file.
func (c *testCLI) readFullConfig() config.Config {
	c.t.Helper()
	buf, err := os.ReadFile(filepath.Join(c.configDir, "kinde", "config.json"))
	require.NoError(c.t, err)
	var cfg config.Config
	require.NoError(c.t, json.Unmarshal(buf, &cfg))
	return cfg
}

func TestLogin_ClientSecretRequiresClientID(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{})

	_, _, err := cli.run("login", "--client_secret", "s3cret")

	assert.ErrorContains(t, err, "--client_id is required")
	assert.Nil(t, cli.readConfig().M2M, "nothing should be saved")
}

func TestLogin_ClientIDBeforeDomainDoesNotCrash(t *testing.T) {
	cli := newTestCLI(t)

	// Previously --client_id wrote to the config while flags were parsed and
	// crashed when no domain was selected yet.
	assert.NotPanics(t, func() {
		_, _, _ = cli.run("login", "--client_secret", "s3cret", "--client_id", "abc")
	})
}

func TestLogin_SecretIsNeverWrittenToConfig(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{})

	_, _, _ = cli.run("login", "--client_secret", "s3cret")

	buf, err := os.ReadFile(filepath.Join(cli.configDir, "kinde", "config.json"))
	require.NoError(t, err)
	assert.NotContains(t, string(buf), "s3cret")
}

func TestManage_RequiresM2MLogin(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{User: &config.UserLogin{}})

	_, _, err := cli.run("business", "get")

	assert.ErrorIs(t, err, config.ErrNoM2MLogin, "a browser login must not be used for management commands")
	assert.ErrorContains(t, err, "kinde login --client_id")
}

func TestWhoami_RequiresUserLogin(t *testing.T) {
	cli := newTestCLI(t)
	cli.writeConfig(config.Environment{M2M: &config.M2MLogin{ClientID: "abc"}})

	_, _, err := cli.run("whoami")

	assert.ErrorContains(t, err, "not logged in through the browser")
}

func TestLogout_ClearsBothLogins(t *testing.T) {
	cli := newTestCLI(t)
	env := config.Environment{
		DomainName: testDomain,
		M2M:        &config.M2MLogin{ClientID: "abc"},
		User:       &config.UserLogin{},
	}
	cli.writeConfig(env)
	require.NoError(t, env.SaveM2MSecret("s3cret"))

	stdout, stderr, err := cli.run("logout")
	require.NoError(t, err)
	assert.Empty(t, stdout, "status messages belong on stderr")
	assert.Equal(t, "Logged out of "+testDomain+"\n", stderr)

	saved := cli.readConfig()
	assert.Nil(t, saved.M2M)
	assert.Nil(t, saved.User)

	saved.M2M = &config.M2MLogin{ClientID: "abc"}
	_, err = saved.M2MClientCredentialsFlow()
	assert.ErrorIs(t, err, config.ErrNoM2MLogin, "client secret should be removed from the session store")
}
