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
		Short: "Login to Kinde",
		RunE:  loginCmd.runLogin,
	}

	loginCmd.cmd.PersistentFlags().Func("client_id", "Client ID to use", func(val string) error {
		return config.FromContext[config.Config](loginCmd.cmd.Context()).
			SetEnvironment(func(env *config.Environment) {
				env.ClientID = val
			})
	})
	loginCmd.cmd.PersistentFlags().Func("client_secret", "Client secret to use, CLI will switch to client_credentials", func(val string) error {
		return config.FromContext[config.Config](loginCmd.cmd.Context()).
			SetEnvironment(func(env *config.Environment) {
				env.ClientSecret = val
			})
	})
	return loginCmd
}

func (c *loginCmd) runLogin(cmd *cobra.Command, args []string) error {

	log := log.Ctx(cmd.Context())
	config := config.FromContext[config.Config](cmd.Context())

	env := config.GetEnvironment()

	if env.DomainName == "" {
		return fmt.Errorf("no environment configured. Please run 'kinde login'")
	}

	if env.ClientSecret != "" {
		log.Debug().Str("client_id", env.ClientID).Msg("Using client credentials flow")
		clientCredentialsFlow, err := env.NewClientCredentialsFlow()
		if err != nil {
			return err
		}
		token, err := clientCredentialsFlow.GetToken(context.WithoutCancel(cmd.Context()))
		if err != nil {
			return fmt.Errorf("failed to get token: %w", err)
		}
		if token.IsValid() {
			log.Info().Msgf("Authenticated using client_credentials")
		} else {
			return fmt.Errorf("failed to authenticate using client_credentials")
		}

	} else {
		log.Debug().Str("client_id", env.ClientID).Msg("Using device authorization flow")
		deviceFlow, err := env.NewDeviceAuthorizationFlow()
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
		token, err := deviceFlow.GetToken(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get token: %w", err)
		}
		log.Info().Msgf("Authenticated as %v", token.GetSubject())
	}

	return nil
}
