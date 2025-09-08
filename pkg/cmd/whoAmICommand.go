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

	return whoAmICmd
}

func (c *whoAmICmd) runWhoAmI(cmd *cobra.Command, args []string) error {
	log := log.Ctx(cmd.Context())
	config := config.FromContext[config.Config](cmd.Context())

	deviceFlow, err := config.GetEnvironment().NewDeviceAuthorizationFlow()
	if err != nil {
		return err
	}

	log.Info().Str("domain", config.GetEnvironment().DomainName).Msg("Checking authentication status")

	if isAuthenticated, err := deviceFlow.IsAuthenticated(cmd.Context()); !isAuthenticated {
		if err != nil {
			log.Error().Err(err).Msg("Failed to check authentication status")
			return nil
		} else {
			return fmt.Errorf("You are not logged in. Please run 'login' command first")
		}
	}

	token, err := deviceFlow.GetToken(cmd.Context())
	if err != nil {
		return fmt.Errorf("Failed to get token: %w", err)
	}

	log.Info().Any("token_claims", token.GetClaims()).Str("client_id", config.GetEnvironment().ClientID).Msgf("Authenticated")

	return nil
}
