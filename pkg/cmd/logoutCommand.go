package cmd

import (
	"errors"
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type logoutCmd struct {
	cmd *cobra.Command
	ios *iostreams.IOStreams
}

func newLogoutCmd(ios *iostreams.IOStreams) *logoutCmd {
	logoutCmd := &logoutCmd{ios: ios}

	logoutCmd.cmd = &cobra.Command{
		Use:   "logout",
		Short: "Log out of every login for the domain",
		Long: `Log out of every login for the domain: the M2M application, the browser
login and the user login.

This removes them from this computer. The access key Kinde created for your
browser login stays until an owner turns off AI agent sign-in.`,
		RunE: logoutCmd.runLogout,
	}

	return logoutCmd
}

// runLogout clears every login for the current domain. The session stores
// are always cleared, so it also removes anything left by an interrupted login.
func (c *logoutCmd) runLogout(cmd *cobra.Command, args []string) error {
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil {
		return errNoEnvironment
	}

	if err := errors.Join(env.ClearM2MLogin(), env.ClearBrowserLogin(), env.ClearUserLogin()); err != nil {
		return err
	}

	if env.M2M != nil || env.Browser != nil || env.User != nil {
		if err := cfg.SetEnvironment(func(e *config.Environment) {
			e.M2M = nil
			e.Browser = nil
			e.User = nil
		}); err != nil {
			return err
		}
	}

	fmt.Fprintf(c.ios.ErrOut, "Logged out of %s\n", env.DomainName)
	return nil
}
