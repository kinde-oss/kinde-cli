package cmd

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// useTempConfig points the CLI at an empty config directory so tests never
// read or write the real ~/.config/kinde.
func useTempConfig(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)
	return dir
}

func TestExecute_ReturnsErrors(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"unknown command", []string{"bogus"}},
		{"unknown flag", []string{"--bogus"}},
		{"manage without login", []string{"manage", "business", "get"}},
		{"whoami without login", []string{"whoami"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useTempConfig(t)
			err := Execute(context.Background(), tt.args)
			assert.Error(t, err)
		})
	}
}

func TestExecute_Succeeds(t *testing.T) {
	useTempConfig(t)
	err := Execute(context.Background(), []string{"--help"})
	assert.NoError(t, err)
}

func TestExecute_InvalidConfigIsReportedAndKept(t *testing.T) {
	dir := useTempConfig(t)
	configFile := filepath.Join(dir, "kinde", "config.json")
	invalid := []byte(`{"current":"mybiz.kinde.com",}`)
	require.NoError(t, os.MkdirAll(filepath.Dir(configFile), 0700))
	require.NoError(t, os.WriteFile(configFile, invalid, 0600))

	err := Execute(context.Background(), []string{"--help"})

	assert.ErrorContains(t, err, configFile)
	contents, readErr := os.ReadFile(configFile)
	require.NoError(t, readErr)
	assert.Equal(t, invalid, contents, "an unreadable config file must not be overwritten")
}
