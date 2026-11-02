package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const otherDomain = "other.kinde.com"

// newLoggedInCLI returns a CLI whose saved default is testDomain, with a saved
// M2M login. Commands that end up using otherDomain fail with ErrNoM2MLogin
// before any network call, which shows which domain a command used.
func newLoggedInCLI(t *testing.T) *testCLI {
	t.Helper()
	cli := newTestCLI(t)
	env := config.Environment{DomainName: testDomain, M2M: &config.M2MLogin{ClientID: "abc"}}
	cli.writeConfig(env)
	require.NoError(t, env.SaveM2MSecret("s3cret"))
	return cli
}

func TestDomainFlag_AppliesToThisRunOnly(t *testing.T) {
	cli := newLoggedInCLI(t)

	_, _, err := cli.run("--domain", otherDomain, "manage", "business", "get")

	assert.ErrorIs(t, err, config.ErrNoM2MLogin, "the command should run against --domain")
	saved := cli.readFullConfig()
	assert.Equal(t, testDomain, saved.CurrentEnvironment, "--domain must not change the saved default")
	assert.NotContains(t, saved.Environments, otherDomain)
}

func TestKindeDomainEnv_AppliesToThisRunOnly(t *testing.T) {
	cli := newLoggedInCLI(t)
	t.Setenv(config.EnvDomain, "https://Other.kinde.com/")

	_, _, err := cli.run("manage", "business", "get")

	assert.ErrorIs(t, err, config.ErrNoM2MLogin, "the command should run against KINDE_DOMAIN")
	assert.Equal(t, testDomain, cli.readFullConfig().CurrentEnvironment)
}

func TestDomainFlag_TakesPrecedenceOverEnv(t *testing.T) {
	cli := newLoggedInCLI(t)
	t.Setenv(config.EnvDomain, testDomain)

	_, _, err := cli.run("--domain", otherDomain, "manage", "business", "get")

	assert.ErrorIs(t, err, config.ErrNoM2MLogin, "--domain should win over KINDE_DOMAIN")
}

func TestDomainFlag_RejectsInvalidDomain(t *testing.T) {
	_, _, err := newTestCLI(t).run("--domain", "mybiz.kinde.com/path", "whoami")
	assert.ErrorContains(t, err, "invalid Kinde domain")
}

func TestReadOnlyCommands_DoNotWriteConfig(t *testing.T) {
	cli := newTestCLI(t)

	_, _, _ = cli.run("--help")
	_, _, _ = cli.run("--domain", otherDomain, "whoami")

	_, err := os.Stat(filepath.Join(cli.configDir, "kinde", "config.json"))
	assert.True(t, os.IsNotExist(err), "config file should not be created by read-only commands")
}

func TestFailedLogin_KeepsSavedDefault(t *testing.T) {
	cli := newLoggedInCLI(t)

	_, _, err := cli.run("--domain", otherDomain, "login", "--client_secret", "s3cret")

	require.Error(t, err)
	assert.Equal(t, testDomain, cli.readFullConfig().CurrentEnvironment)
}

func TestManage_EnvCredentialsMustBothBeSet(t *testing.T) {
	cli := newTestCLI(t)
	t.Setenv(config.EnvDomain, testDomain)
	t.Setenv(config.EnvM2MClientID, "abc")

	_, _, err := cli.run("manage", "business", "get")

	assert.ErrorContains(t, err, config.EnvM2MClientSecret)
}
