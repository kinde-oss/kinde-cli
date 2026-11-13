package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"reflect"
	"runtime"
	"time"

	"github.com/kinde-oss/kinde-cli/pkg/agentauth"
	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// openBrowser opens url in the user's browser. Tests replace it.
var openBrowser = func(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}

// newHTTPClient returns the client used for sign-in and MCP requests. Tests
// replace it.
var newHTTPClient = func() *http.Client {
	return &http.Client{Timeout: 60 * time.Second}
}

// signInNotEnabledHint explains how to turn on the browser login.
const signInNotEnabledHint = "An owner can turn it on in Kinde under Settings > APIs > Kinde Management API > Enable OAuth sign-in"

// runBrowserLogin signs a Kinde team member in through the browser with
// Kinde's AI agent sign-in, and saves the login for the domain.
func runBrowserLogin(cmd *cobra.Command, ios *iostreams.IOStreams, cfg *config.Config, env *config.Environment) error {
	if !ios.CanPrompt() {
		return fmt.Errorf("the browser login needs a terminal. In scripts, log in with an M2M application (--client-id with --client-secret-stdin), or set %s and %s",
			config.EnvM2MClientID, config.EnvM2MClientSecret)
	}
	ctx := cmd.Context()
	httpClient := newHTTPClient()

	server, err := agentauth.Discover(ctx, httpClient, "https://"+env.DomainName+"/mcp")
	if errors.Is(err, agentauth.ErrSignInNotEnabled) {
		log.Ctx(ctx).Debug().Err(err).Msg("Agent sign-in discovery failed")
		return fmt.Errorf("the browser login isn't turned on for %s. %s, or log in with an M2M application: run 'kinde login --client-id <client ID>'",
			env.DomainName, signInNotEnabledHint)
	}
	if err != nil {
		return err
	}

	// Reuse the client registered at the last login while it's still valid,
	// rather than registering a new one every time.
	session := &agentauth.Session{Server: *server}
	if previous, err := loadBrowserSession(env); err == nil && previous.Server.Issuer == server.Issuer && previous.Registration.Usable(time.Now()) {
		session.Registration = previous.Registration
	} else {
		registration, err := agentauth.Register(ctx, httpClient, server)
		if err != nil {
			return err
		}
		session.Registration = *registration
	}

	token, err := agentauth.Authorize(ctx, httpClient, server, &session.Registration, func(signInURL string) {
		fmt.Fprintf(ios.ErrOut, "Opening the browser to sign in to %s. If it doesn't open, visit:\n\n  %s\n\n", env.DomainName, signInURL)
		if err := openBrowser(signInURL); err != nil {
			log.Ctx(ctx).Debug().Err(err).Msg("Could not open the browser")
		}
		fmt.Fprintln(ios.ErrOut, "Waiting for you to sign in...")
	})
	if err != nil {
		return err
	}
	session.Token = token

	if err := saveBrowserSession(env, session); err != nil {
		return err
	}
	if err := cfg.SetEnvironment(func(e *config.Environment) {
		e.Browser = &config.BrowserLogin{ClientID: session.Registration.ClientID, Issuer: server.Issuer}
	}); err != nil {
		return err
	}
	cfg.SetDefaultDomain(env.DomainName)

	fmt.Fprintf(ios.ErrOut, "Logged in to %s through the browser\n", env.DomainName)
	if m2m := m2mLoginStatus(env); m2m != nil {
		fmt.Fprintf(ios.ErrOut, "Note: Management API commands keep using the M2M application %s while it's logged in. To use the browser login, run 'kinde logout' and then 'kinde login' again\n", m2m.ClientID)
		return nil
	}
	fmt.Fprint(ios.ErrOut, browserLoginLimits)
	return nil
}

// browserLoginLimits tells the user, once they're logged in, what the browser
// login can't do compared with an M2M application.
const browserLoginLimits = `
With a browser login:
  - secrets in responses, such as client secrets, are hidden
  - a few commands aren't available, such as billing, organization roles and
    permissions, and logos
  - it needs you to sign in, so it isn't suited to CI
For those, log in with an M2M application: kinde login --client-id <client ID>
`

// callWithBrowserLogin sends a Management API request through the domain's
// MCP server with the browser login.
func callWithBrowserLogin(cmd *cobra.Command, ios *iostreams.IOStreams, env *config.Environment, op resourceCommand, methodArgs []reflect.Value) error {
	session, err := loadBrowserSession(env)
	if err != nil {
		return fmt.Errorf("%w: run 'kinde login' again", err)
	}

	// Ask once the request is valid and there's a login to send it with.
	if isDestructive(op) {
		if err := confirmDestructive(cmd, ios, env, op); err != nil {
			return err
		}
	}

	httpClient := newHTTPClient()
	accessToken := func(ctx context.Context) (string, error) {
		changed, err := agentauth.Refresh(ctx, httpClient, session)
		if err != nil {
			log.Ctx(ctx).Debug().Err(err).Msg("Could not refresh the browser login")
			if agentauth.IsInvalidGrant(err) || session.Token == nil || session.Token.RefreshToken == "" {
				return "", &userFacingError{fmt.Sprintf("your browser login for %s has expired or been revoked: run 'kinde login' again", env.DomainName)}
			}
			return "", &userFacingError{fmt.Sprintf("failed to refresh your browser login: %v", err)}
		}
		if changed {
			if err := saveBrowserSession(env, session); err != nil {
				log.Ctx(ctx).Warn().Err(err).Msg("Could not save the refreshed browser login")
			}
		}
		return session.Token.AccessToken, nil
	}

	transport := &mcpTransport{
		mcpURL:    session.Server.Resource,
		operation: op.operation,
		path:      docsFor(op.operation).Path,
		base:      httpClient.Transport,
	}
	if transport.base == nil {
		transport.base = http.DefaultTransport
	}
	client := *httpClient
	client.Transport = transport

	err = callOperation(cmd, "https://"+env.DomainName, &client, accessToken, op.operation, methodArgs, loginBrowser)
	if err == nil && transport.redacted {
		fmt.Fprintln(ios.ErrOut, "Note: secrets are hidden with a browser login. "+m2mHint+" to see them")
	}
	return err
}

func loadBrowserSession(env *config.Environment) (*agentauth.Session, error) {
	data, err := env.BrowserSession()
	if err != nil {
		return nil, err
	}
	var session agentauth.Session
	if err := json.Unmarshal(data, &session); err != nil || session.Token == nil {
		return nil, fmt.Errorf("%w: the saved login is unreadable", config.ErrNoBrowserLogin)
	}
	return &session, nil
}

func saveBrowserSession(env *config.Environment, session *agentauth.Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return env.SaveBrowserSession(data)
}

// loginEnvironment returns the domain to log in to. In a terminal it asks
// for one when none was given, so a first login needs no flags.
func loginEnvironment(ios *iostreams.IOStreams, cfg *config.Config) (*config.Environment, error) {
	if env := cfg.GetEnvironment(); env != nil && env.DomainName != "" {
		return env, nil
	}
	if !ios.CanPrompt() {
		return nil, errNoEnvironment
	}
	domain, err := ios.Prompt("Kinde domain (e.g. mybusiness.kinde.com): ")
	if err != nil {
		return nil, err
	}
	if err := cfg.UseDomain(domain); err != nil {
		return nil, err
	}
	return cfg.GetEnvironment(), nil
}
