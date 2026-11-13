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
	"golang.org/x/oauth2"
)

// newAccountCmd groups the commands that act as a user of the business,
// rather than as an M2M application or a Kinde team member.
func newAccountCmd(ios *iostreams.IOStreams) (*cobra.Command, error) {
	accountCmd := &cobra.Command{
		Use:   "account",
		Args:  cobra.NoArgs,
		Short: "Log in as a user and see what they can access (Account API)",
		Long: `Log in as a user of your Kinde business, and use the Account API to see what
that user can access: their permissions, roles, feature flags, entitlements
and properties.

This is useful for testing how your application will see a user. A user login
can't call the Management API; use 'kinde login' for that.`,
		Example: `  kinde account login --domain mybusiness.kinde.com --client-id <device app client ID>
  kinde account permissions
  kinde account feature-flags | jq '.data.feature_flags'`,
	}
	apiCmds, err := newAccountAPICommands()
	if err != nil {
		return nil, err
	}
	accountCmd.AddCommand(newAccountLoginCmd(ios).cmd, newAccountTokenCmd())
	accountCmd.AddCommand(apiCmds...)
	return accountCmd, nil
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
		Short: "Log in as a user of your business, for the Account API",
		Long: `Log in as a user of your business, using the device authorization flow, so
the 'kinde account' commands can show what that user can access.

Open the URL shown (or scan the QR code), sign in, and approve the login.

The login needs a Device and IoT application in your business. Pass its client
ID with --client-id, or set one as the default for device logins in Kinde and
leave --client-id out.

This isn't how you log in to run Management API commands: use 'kinde login'
for that.`,
		Example: `  kinde account login --domain mybusiness.kinde.com
  kinde account login --client-id <application client ID>`,
		RunE: c.run,
	}
	c.cmd.Flags().StringVar(&c.clientID, "client-id", "", "Client ID of a Device and IoT application (default: the business's default device application)")
	return c
}

func (c *accountLoginCmd) run(cmd *cobra.Command, args []string) error {
	log := log.Ctx(cmd.Context())
	cfg := config.FromContext[config.Config](cmd.Context())

	env, err := loginEnvironment(c.ios, cfg)
	if err != nil {
		return err
	}
	log.Debug().Str("client_id", c.clientID).Msg("Using device authorization flow")

	deviceFlow, err := env.NewDeviceAuthorizationFlow(c.clientID)
	if err != nil {
		return err
	}

	deviceAuth, err := deviceFlow.StartDeviceAuth(cmd.Context())
	var oauthErr *oauth2.RetrieveError
	if c.clientID == "" && errors.As(err, &oauthErr) && (oauthErr.ErrorCode == "invalid_request" || oauthErr.ErrorCode == "invalid_client") {
		log.Debug().Err(err).Msg("Device login without a client ID was refused")
		return fmt.Errorf("%s has no default application for device logins: pass --client-id with a Device and IoT application's client ID, or set one as the default for device logins in Kinde", env.DomainName)
	}
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
