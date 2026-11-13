package cmd

import (
	"bytes"
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
	// Ignore any Kinde settings in the developer's own shell.
	for _, name := range []string{config.EnvDomain, config.EnvM2MClientID, config.EnvM2MClientSecret} {
		t.Setenv(name, "")
	}
	t.Cleanup(config.UseFileSessionStore("test-password"))
	return &testCLI{t: t, configDir: dir}
}

// run executes the CLI with args and returns what it wrote to stdout and stderr.
func (c *testCLI) run(args ...string) (stdout, stderr string, err error) {
	c.t.Helper()
	return c.runWith(nil, args...)
}

// runWith is run with a chance to set up the streams first, e.g. to provide
// stdin or make them behave like a terminal.
func (c *testCLI) runWith(setup func(ios *iostreams.IOStreams, in *bytes.Buffer), args ...string) (stdout, stderr string, err error) {
	c.t.Helper()
	ios, in, out, errOut := iostreams.Test()
	if setup != nil {
		setup(ios, in)
	}
	err = Execute(context.Background(), args, ios)
	return out.String(), errOut.String(), err
}
