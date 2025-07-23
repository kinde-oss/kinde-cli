package cmd

import (
	"context"
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type loginCmd struct {
	cmd *cobra.Command
}

func newLoginCmd() *loginCmd {

	loginCmd := &loginCmd{}

	loginCmd.cmd = &cobra.Command{
		Use:   "login",
		Args:  nil,
		Short: "Version",
		RunE:  loginCmd.runLogin,
	}

	return loginCmd
}

func (c *loginCmd) runLogin(cmd *cobra.Command, args []string) error {

	log := log.Ctx(cmd.Context())
	config := config.FromContext(cmd.Context())

	env := config.GetEnvironment()

	if env.DomainName == "" {
		return fmt.Errorf("no environment configured. Please run 'kinde login'")
	}

	deviceFlow, err := config.NewDeviceAuthorizationFlow()
	if err != nil {
		return err
	}

	deviceAuth, err := deviceFlow.StartDeviceAuth(c.cmd.Context())
	if err != nil {
		return err
	}

	log.Info().Msgf("Please open the following URL in your browser: %v", deviceAuth.VerificationURIComplete)
	log.Info().Msg("Waiting for user to authorize...")

	err = deviceFlow.ExchangeDeviceAccessToken(context.WithoutCancel(cmd.Context()), deviceAuth)
	if err != nil {
		return err
	}
	token, err := deviceFlow.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	log.Info().Msgf("Authenticated as %v", token.GetSubject())

	return nil
}
