package cmd

import (
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type whoAmICmd struct {
	cmd         *cobra.Command
	kindeDomain string
}

func newWhoAmI() *whoAmICmd {
	whoAmICmd := &whoAmICmd{}

	whoAmICmd.cmd = &cobra.Command{
		Use:   "whoami",
		Args:  nil,
		Short: "Show current logged in user",
		RunE:  whoAmICmd.runWhoAmI,
	}

	whoAmICmd.cmd.Flags().StringVar(&whoAmICmd.kindeDomain, "kinde-domain", "app.kinde.com", "Uses the kinde domain to connect to")

	return whoAmICmd
}

func (c *whoAmICmd) runWhoAmI(cmd *cobra.Command, args []string) error {
	log := log.Ctx(cmd.Context())
	config := config.FromContext(cmd.Context())

	deviceFlow, err := config.NewDeviceAuthorizationFlow()
	if err != nil {
		return err
	}

	if !deviceFlow.IsAuthenticated() {
		return fmt.Errorf("you are not logged in. Please run 'login' command first")
	}

	token, err := deviceFlow.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	log.Info().Msgf("Authenticated as %v", token.GetSubject())

	return nil
}
