package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type whoAmICmd struct {
	cmd *cobra.Command
	ios *iostreams.IOStreams
}

func newWhoAmI(ios *iostreams.IOStreams) *whoAmICmd {
	whoAmICmd := &whoAmICmd{ios: ios}

	whoAmICmd.cmd = &cobra.Command{
		Use:   "whoami",
		Args:  nil,
		Short: "Show the user logged in through the browser",
		RunE:  whoAmICmd.runWhoAmI,
	}

	return whoAmICmd
}

func (c *whoAmICmd) runWhoAmI(cmd *cobra.Command, args []string) error {
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil {
		return errNoEnvironment
	}

	deviceFlow, err := env.UserDeviceAuthorizationFlow()
	if errors.Is(err, config.ErrNoUserLogin) {
		return fmt.Errorf("you are not logged in through the browser. Run 'kinde account login' first")
	}
	if err != nil {
		return err
	}

	if isAuthenticated, err := deviceFlow.IsAuthenticated(cmd.Context()); !isAuthenticated {
		if err != nil {
			return fmt.Errorf("failed to check authentication status: %w", err)
		}
		return fmt.Errorf("your browser login is no longer valid. Run 'kinde account login' again")
	}

	token, err := deviceFlow.GetToken(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	// The token's claims are the data, so they go to stdout as JSON.
	claims, err := json.MarshalIndent(token.GetClaims(), "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode token claims: %w", err)
	}
	fmt.Fprintf(c.ios.Out, "%s\n", claims)
	return nil
}
