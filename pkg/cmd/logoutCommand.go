package cmd

import (
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

func (c *logoutCmd) runLogout(cmd *cobra.Command, args []string) error {
	log := log.Ctx(cmd.Context())
	config := config.FromContext(cmd.Context())

	deviceFlow, err := config.NewDeviceAuthorizationFlow()
	if err != nil {
		return err
	}

	err = deviceFlow.Logout()
	if err != nil {
	}

	log.Info().Msg("Successfully logged out")
	return nil
}
