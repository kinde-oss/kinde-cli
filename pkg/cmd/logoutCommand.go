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
		Short: "Logout and clear authentication tokens",
		RunE:  logoutCmd.runLogout,
	}

	return logoutCmd
}

// runLogout clears both logins for the current domain. Both session stores
// are always cleared, so it also removes anything left by an interrupted login.
func (c *logoutCmd) runLogout(cmd *cobra.Command, args []string) error {
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil {
		return errNoEnvironment
	}

	if err := errors.Join(env.ClearM2MLogin(), env.ClearUserLogin()); err != nil {
		return err
	}

	if env.M2M != nil || env.User != nil {
		if err := cfg.SetEnvironment(func(e *config.Environment) {
			e.M2M = nil
			e.User = nil
		}); err != nil {
			return err
		}
	}

	fmt.Fprintf(c.ios.ErrOut, "Logged out of %s\n", env.DomainName)
	return nil
}
