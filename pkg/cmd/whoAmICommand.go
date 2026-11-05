package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/rs/zerolog/log"
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
		Args:  cobra.NoArgs,
		Short: "Show the logins for the domain",
		Long: `Show the logins for the domain as JSON: the M2M application used by the
Management API commands, and the user logged in through the browser for the
'kinde account' commands. A login that isn't set up is null.

Exits with status 1 if neither login is set up.`,
		RunE: whoAmICmd.runWhoAmI,
	}

	return whoAmICmd
}

type (
	whoAmIStatus struct {
		Domain string      `json:"domain"`
		M2M    *m2mStatus  `json:"m2m"`
		User   *userStatus `json:"user"`
	}

	m2mStatus struct {
		ClientID string `json:"client_id"`
		// Source is "login" for 'kinde login', or "environment" when the
		// KINDE_MANAGEMENT_CLIENT_* variables are set, which take precedence.
		Source string `json:"source"`
	}

	userStatus struct {
		ID        string     `json:"id"`
		ClientID  string     `json:"client_id,omitempty"`
		OrgCode   string     `json:"org_code,omitempty"`
		ExpiresAt *time.Time `json:"expires_at,omitempty"`
	}
)

func (c *whoAmICmd) runWhoAmI(cmd *cobra.Command, args []string) error {
	cfg := config.FromContext[config.Config](cmd.Context())

	env := cfg.GetEnvironment()
	if env == nil {
		return errNoEnvironment
	}

	status := whoAmIStatus{Domain: env.DomainName, M2M: m2mLoginStatus(env)}
	user, err := c.userLoginStatus(cmd, env)
	if err != nil {
		return err
	}
	status.User = user

	if status.M2M == nil && status.User == nil {
		return fmt.Errorf("not logged in to %s: run 'kinde login' for the Management API, or 'kinde account login' to log in as a user", env.DomainName)
	}

	out, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		return err
	}
	fmt.Fprintf(c.ios.Out, "%s\n", out)
	return nil
}

// m2mLoginStatus reports the M2M application the resource commands will use.
// It doesn't contact Kinde; a command reports bad credentials when it runs.
func m2mLoginStatus(env *config.Environment) *m2mStatus {
	if id := os.Getenv(config.EnvM2MClientID); id != "" && os.Getenv(config.EnvM2MClientSecret) != "" {
		return &m2mStatus{ClientID: id, Source: "environment"}
	}
	if env.M2M != nil {
		return &m2mStatus{ClientID: env.M2M.ClientID, Source: "login"}
	}
	return nil
}

// userLoginStatus reports the user logged in through the browser, or nil if
// there's no valid browser login.
func (c *whoAmICmd) userLoginStatus(cmd *cobra.Command, env *config.Environment) (*userStatus, error) {
	if env.User == nil {
		return nil, nil
	}
	flow, err := env.UserDeviceAuthorizationFlow()
	if err != nil {
		return nil, err
	}
	token, err := flow.GetToken(cmd.Context())
	if err != nil {
		log.Ctx(cmd.Context()).Debug().Err(err).Msg("Browser login is not valid")
		fmt.Fprintln(c.ios.ErrOut, "The browser login is no longer valid: run 'kinde account login' again")
		return nil, nil
	}

	user := &userStatus{ID: token.GetSubject(), ClientID: env.User.ClientID, OrgCode: token.GetOrganizationCode()}
	if exp, ok := token.GetExpiration(); ok {
		expiresAt := time.Unix(exp, 0).UTC()
		user.ExpiresAt = &expiresAt
	}
	return user, nil
}
