package cmd

import (
	"context"
	"testing"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
)

// testCLI runs the CLI in isolation: a temporary config directory, a
// file-based session store instead of the OS keychain, and captured output.
type testCLI struct {
	t         *testing.T
	configDir string
}

func newTestCLI(t *testing.T) *testCLI {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)
	t.Cleanup(config.UseFileSessionStore("test-password"))
	return &testCLI{t: t, configDir: dir}
}

// run executes the CLI with args and returns what it wrote to stdout and stderr.
func (c *testCLI) run(args ...string) (stdout, stderr string, err error) {
	c.t.Helper()
	ios, _, out, errOut := iostreams.Test()
	err = Execute(context.Background(), args, ios)
	return out.String(), errOut.String(), err
}
