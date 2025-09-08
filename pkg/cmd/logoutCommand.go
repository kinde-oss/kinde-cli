package cmd

import (
	"fmt"

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
	config := config.FromContext[config.Config](cmd.Context())

	env := config.GetEnvironment()
	if env == nil {
		return fmt.Errorf("no environment configured")
	}

	// Clear client secret from session if it exists (for client credentials flow)
	if err := env.ClearClientSecretFromSession(); err != nil {
		// Log the error but don't fail the logout - the secret might not exist
		log.Debug().Err(err).Msg("Failed to clear client secret from session (may not exist)")
	}

	// Handle device authorization flow logout
	deviceFlow, err := env.NewDeviceAuthorizationFlow()
	if err != nil {
		return err
	}

	err = deviceFlow.Logout()
	if err != nil {
		// Log the error but don't fail the logout - device flow might not be active
		log.Debug().Err(err).Msg("Failed to logout from device flow (may not be active)")
	}

	log.Info().Msg("Successfully logged out")
	return nil
}
