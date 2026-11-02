package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientSecretSessionRoundTrip(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)
	t.Cleanup(UseFileSessionStore("test-password"))

	env := &Environment{DomainName: "mybiz.kinde.com"}

	require.NoError(t, env.StoreClientSecretInSession("s3cret"))

	secret, err := env.GetClientSecretFromSession()
	require.NoError(t, err)
	assert.Equal(t, "s3cret", secret)

	// The secret must live in the test's directory, not the OS keychain.
	storeDir, err := env.keychainFolderName(env.DomainName)
	require.NoError(t, err)
	entries, err := os.ReadDir(storeDir)
	require.NoError(t, err)
	assert.NotEmpty(t, entries, "expected the session store to write files under %s", storeDir)

	// kinde-go's DeleteKey reports an error with the file backend because it
	// doesn't recognise the backend's "file not found" for a legacy key that
	// was never written, even though it removes the secret itself. Check the
	// secret is gone rather than relying on the returned error.
	_ = env.ClearClientSecretFromSession()
	_, err = env.GetClientSecretFromSession()
	assert.Error(t, err, "secret should be gone after clearing")
}
