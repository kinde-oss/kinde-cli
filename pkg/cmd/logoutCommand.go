package cmd

import (
	"errors"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type logoutCmd struct {
	cmd         *cobra.Command
	kindeDomain string
}

func newLogoutCmd() *logoutCmd {
	logoutCmd := &logoutCmd{}

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
	log := log.Ctx(cmd.Context())
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil {
		return errNoEnvironment
	}

	if err := errors.Join(env.ClearM2MLogin(), env.ClearUserLogin()); err != nil {
		return err
	}

	if err := cfg.SetEnvironment(func(e *config.Environment) {
		e.M2M = nil
		e.User = nil
	}); err != nil {
		return err
	}

	log.Info().Str("domain", env.DomainName).Msg("Logged out")
	return nil
}
