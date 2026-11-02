package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
			stdout, _, err := newTestCLI(t).run(tt.args...)
			assert.Error(t, err)
			assert.Empty(t, stdout, "failures must not write to stdout")
		})
	}
}

func TestExecute_HelpGoesToStdout(t *testing.T) {
	stdout, stderr, err := newTestCLI(t).run("--help")
	assert.NoError(t, err)
	assert.Contains(t, stdout, "Usage:")
	assert.Empty(t, stderr)
}

func TestExecute_InvalidConfigIsReportedAndKept(t *testing.T) {
	cli := newTestCLI(t)
	configFile := filepath.Join(cli.configDir, "kinde", "config.json")
	invalid := []byte(`{"current":"mybiz.kinde.com",}`)
	require.NoError(t, os.MkdirAll(filepath.Dir(configFile), 0700))
	require.NoError(t, os.WriteFile(configFile, invalid, 0600))

	_, _, err := cli.run("--help")

	assert.ErrorContains(t, err, configFile)
	contents, readErr := os.ReadFile(configFile)
	require.NoError(t, readErr)
	assert.Equal(t, invalid, contents, "an unreadable config file must not be overwritten")
}
