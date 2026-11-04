package cmd

import (
	"context"
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/mdp/qrterminal/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newAccountCmd groups the commands that act as a user signed in through the
// browser, rather than as an M2M application.
func newAccountCmd(ios *iostreams.IOStreams) *cobra.Command {
	accountCmd := &cobra.Command{
		Use:   "account",
		Args:  cobra.NoArgs,
		Short: "Log in as a user through the browser",
		Long: `Log in as a user of your Kinde business through the browser, for example to
test your application's sign-in.

A user login can't call the Management API; use 'kinde login' with an M2M
application for that.`,
	}
	accountCmd.AddCommand(newAccountLoginCmd(ios).cmd)
	return accountCmd
}

type accountLoginCmd struct {
	cmd      *cobra.Command
	ios      *iostreams.IOStreams
	clientID string
}

func newAccountLoginCmd(ios *iostreams.IOStreams) *accountLoginCmd {
	c := &accountLoginCmd{ios: ios}
	c.cmd = &cobra.Command{
		Use:   "login",
		Args:  cobra.NoArgs,
		Short: "Log in as a user through the browser",
		Long: `Log in as a user through the browser, using the device authorization flow.

Open the URL shown (or scan the QR code), sign in, and approve the login. Use
--client-id to sign in to a specific application; otherwise Kinde uses the
business's default application for device logins.`,
		Example: `  kinde account login --domain mybusiness.kinde.com
  kinde account login --client-id <application client ID>`,
		RunE: c.run,
	}
	c.cmd.Flags().StringVar(&c.clientID, "client-id", "", "Client ID of the application to sign in to")
	return c
}

func (c *accountLoginCmd) run(cmd *cobra.Command, args []string) error {
	log := log.Ctx(cmd.Context())
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil || env.DomainName == "" {
		return errNoEnvironment
	}
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
