package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
		Long: `Show the logins for the domain as JSON:

  m2m             the M2M application ('kinde login --client-id')
  browser         you, signed in through the browser ('kinde login')
  user            a user of your business, for the 'kinde account' commands
                  ('kinde account login')
  management_api  which login the Management API commands use: "m2m" when
                  it's set up, otherwise "browser"

A login that isn't set up is null. Exits with status 1 if none is set up.`,
		RunE: whoAmICmd.runWhoAmI,
	}

	return whoAmICmd
}

type (
	whoAmIStatus struct {
		Domain        string         `json:"domain"`
		ManagementAPI *string        `json:"management_api"`
		M2M           *m2mStatus     `json:"m2m"`
		Browser       *browserStatus `json:"browser"`
		User          *userStatus    `json:"user"`
	}

	browserStatus struct {
		// UserID is your Kinde team member ID, from the login's token.
		UserID   string `json:"user_id,omitempty"`
		ClientID string `json:"client_id"`
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

	status := whoAmIStatus{Domain: env.DomainName, M2M: m2mLoginStatus(env), Browser: c.browserLoginStatus(env)}
	user, err := c.userLoginStatus(cmd, env)
	if err != nil {
		return err
	}
	status.User = user

	switch {
	case status.M2M != nil:
		status.ManagementAPI = ptr("m2m")
	case status.Browser != nil:
		status.ManagementAPI = ptr("browser")
	}

	if status.M2M == nil && status.Browser == nil && status.User == nil {
		return fmt.Errorf("not logged in to %s: run 'kinde login' for the Management API, or 'kinde account login' to log in as a user of your business", env.DomainName)
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

// browserLoginStatus reports the browser login, or nil if there isn't one.
// It doesn't contact Kinde; a command reports an expired login when it runs.
func (c *whoAmICmd) browserLoginStatus(env *config.Environment) *browserStatus {
	session, err := loadBrowserSession(env)
	if err != nil {
		return nil
	}
	return &browserStatus{UserID: tokenSubject(session.Token.AccessToken), ClientID: session.Registration.ClientID}
}

// tokenSubject reads the sub claim of a JWT for display, without verifying it.
func tokenSubject(token string) string {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return ""
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}
	var claims struct {
		Sub string `json:"sub"`
	}
	json.Unmarshal(payload, &claims)
	return claims.Sub
}

func ptr[T any](v T) *T { return &v }

// userLoginStatus reports the user logged in with 'kinde account login', or
// nil if there's no valid user login.
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
		log.Ctx(cmd.Context()).Debug().Err(err).Msg("User login is not valid")
		fmt.Fprintln(c.ios.ErrOut, "The user login is no longer valid: run 'kinde account login' again")
		return nil, nil
	}

	user := &userStatus{ID: token.GetSubject(), ClientID: env.User.ClientID, OrgCode: token.GetOrganizationCode()}
	if exp, ok := token.GetExpiration(); ok {
		expiresAt := time.Unix(exp, 0).UTC()
		user.ExpiresAt = &expiresAt
	}
	return user, nil
}
