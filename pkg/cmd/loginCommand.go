package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type loginCmd struct {
	cmd               *cobra.Command
	ios               *iostreams.IOStreams
	clientID          string
	clientSecretStdin bool
}

func newLoginCmd(ios *iostreams.IOStreams) *loginCmd {

	loginCmd := &loginCmd{ios: ios}

	loginCmd.cmd = &cobra.Command{
		Use:   "login",
		Args:  cobra.NoArgs,
		Short: "Log in with an M2M application to use the Management API",
		Long: fmt.Sprintf(`Log in with a machine-to-machine (M2M) application, which the resource commands
use to call the Kinde Management API.

Create the M2M application in Kinde, authorize it for the Kinde Management API,
and grant it the scopes you need. The client secret is asked for without being
shown, and is stored in your operating system's keychain.

In scripts, pass the secret on stdin with --client-secret-stdin, or skip
logging in and set %s and %s.

To log in as a user through the browser, use 'kinde account login'.`, config.EnvM2MClientID, config.EnvM2MClientSecret),
		Example: `  kinde login --domain mybusiness.kinde.com --client-id <client ID>
  echo "$CLIENT_SECRET" | kinde login --client-id <client ID> --client-secret-stdin`,
		RunE: loginCmd.runLogin,
	}

	loginCmd.cmd.Flags().StringVar(&loginCmd.clientID, "client-id", "", "Client ID of the M2M application (required)")
	loginCmd.cmd.Flags().BoolVar(&loginCmd.clientSecretStdin, "client-secret-stdin", false, "Read the client secret from stdin")
	return loginCmd
}

// runLogin checks the credentials first and only saves them once they work,
// so a failed login leaves any existing M2M login in place.
func (c *loginCmd) runLogin(cmd *cobra.Command, args []string) error {
	log := log.Ctx(cmd.Context())
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil || env.DomainName == "" {
		return errNoEnvironment
	}
	if c.clientID == "" {
		return errors.New("missing required flags: --client-id")
	}

	clientSecret, err := c.readClientSecret()
	if err != nil {
		return err
	}
	log.Debug().Str("client_id", c.clientID).Msg("Using client credentials flow")

	if err := env.VerifyClientCredentials(context.WithoutCancel(cmd.Context()), c.clientID, clientSecret); err != nil {
		return fmt.Errorf("failed to log in with the client ID and secret: %w", err)
	}

	if err := env.SaveM2MSecret(clientSecret); err != nil {
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

// readClientSecret reads the secret from stdin or a hidden prompt. There is
// deliberately no flag for it: flag values are kept in shell history and are
// visible to other users in the process list.
func (c *loginCmd) readClientSecret() (string, error) {
	var secret string
	switch {
	case c.clientSecretStdin:
		data, err := io.ReadAll(c.ios.In)
		if err != nil {
			return "", fmt.Errorf("failed to read the client secret from stdin: %w", err)
		}
		secret = string(data)
	case c.ios.CanPrompt():
		var err error
		if secret, err = c.ios.ReadSecret("Client secret: "); err != nil {
			return "", fmt.Errorf("failed to read the client secret: %w", err)
		}
	default:
		return "", fmt.Errorf("can't prompt for the client secret outside a terminal: pass it on stdin with --client-secret-stdin, or set %s and %s instead of logging in",
			config.EnvM2MClientID, config.EnvM2MClientSecret)
	}

	secret = strings.TrimSpace(secret)
	if secret == "" {
		return "", errors.New("the client secret is empty")
	}
	return secret, nil
}
