package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/mdp/qrterminal/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type loginCmd struct {
	cmd          *cobra.Command
	ios          *iostreams.IOStreams
	clientID     string
	clientSecret string
}

func newLoginCmd(ios *iostreams.IOStreams) *loginCmd {

	loginCmd := &loginCmd{ios: ios}

	loginCmd.cmd = &cobra.Command{
		Use:   "login",
		Args:  nil,
		Short: "Login to Kinde",
		RunE:  loginCmd.runLogin,
	}

	loginCmd.cmd.Flags().StringVar(&loginCmd.clientID, "client_id", "", "Client ID of the application to log in with")
	loginCmd.cmd.Flags().StringVar(&loginCmd.clientSecret, "client_secret", "", "Client secret of an M2M application; logs in with client credentials instead of the browser")
	return loginCmd
}

func (c *loginCmd) runLogin(cmd *cobra.Command, args []string) error {
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil || env.DomainName == "" {
		return errNoEnvironment
	}

	if c.clientSecret != "" {
		return c.loginM2M(cmd, cfg, env)
	}
	return c.loginUser(cmd, cfg, env)
}

// loginM2M checks the credentials first and only saves them once they work,
// so a failed login leaves any existing M2M login in place.
func (c *loginCmd) loginM2M(cmd *cobra.Command, cfg *config.Config, env *config.Environment) error {
	log := log.Ctx(cmd.Context())

	if c.clientID == "" {
		return errors.New("--client_id is required when using --client_secret")
	}
	log.Debug().Str("client_id", c.clientID).Msg("Using client credentials flow")

	if err := env.VerifyClientCredentials(context.WithoutCancel(cmd.Context()), c.clientID, c.clientSecret); err != nil {
		return fmt.Errorf("failed to authenticate using client_credentials: %w", err)
	}

	if err := env.SaveM2MSecret(c.clientSecret); err != nil {
		return err
	}
	if err := cfg.SetEnvironment(func(e *config.Environment) {
		e.M2M = &config.M2MLogin{ClientID: c.clientID}
	}); err != nil {
		return err
	}
	cfg.SetDefaultDomain(env.DomainName)

	fmt.Fprintf(c.ios.ErrOut, "Logged in to %s with M2M application %s\n", env.DomainName, c.clientID)
	return nil
}

func (c *loginCmd) loginUser(cmd *cobra.Command, cfg *config.Config, env *config.Environment) error {
	log := log.Ctx(cmd.Context())
	log.Debug().Str("client_id", c.clientID).Msg("Using device authorization flow")

	deviceFlow, err := env.NewDeviceAuthorizationFlow(c.clientID)
	if err != nil {
		return err
	}

	deviceAuth, err := deviceFlow.StartDeviceAuth(cmd.Context())
	if err != nil {
		return err
	}

	// Prompts go to stderr, keeping stdout for data.
	fmt.Fprintf(c.ios.ErrOut, "Open this URL in your browser to log in:\n\n  %s\n\n", deviceAuth.VerificationURIComplete)
	qrterminal.GenerateHalfBlock(deviceAuth.VerificationURIComplete, qrterminal.M, c.ios.ErrOut)
	fmt.Fprintln(c.ios.ErrOut)
	fmt.Fprintln(c.ios.ErrOut, "Waiting for you to approve the login...")

	err = deviceFlow.ExchangeDeviceAccessToken(context.WithoutCancel(cmd.Context()), deviceAuth)
	if err != nil {
		return err
	}
	token, err := deviceFlow.GetToken(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	if err := cfg.SetEnvironment(func(e *config.Environment) {
		e.User = &config.UserLogin{ClientID: c.clientID}
	}); err != nil {
		return err
	}
	cfg.SetDefaultDomain(env.DomainName)

	fmt.Fprintf(c.ios.ErrOut, "Logged in to %s as %s\n", env.DomainName, token.GetSubject())
	return nil
}
