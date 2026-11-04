package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/keyring"
	"github.com/kinde-oss/kinde-go/frameworks/cli"
	"github.com/kinde-oss/kinde-go/jwt"
	"github.com/kinde-oss/kinde-go/oauth2/authorization_code"
	"github.com/kinde-oss/kinde-go/oauth2/client_credentials"
)

// Identity is one way of signing in to a Kinde domain. Each identity keeps its
// tokens and secrets in its own session store, so the two logins never mix.
type Identity string

const (
	// IdentityM2M is a machine-to-machine application, used for the Management API.
	IdentityM2M Identity = "m2m"
	// IdentityUser is a user signed in through the browser, used for the Account API.
	IdentityUser Identity = "user"
)

const clientSecretKey = "client_secret"

var (
	ErrNoM2MLogin  = errors.New("no M2M login for this domain")
	ErrNoUserLogin = errors.New("no browser login for this domain")
)

type (
	Environment struct {
		DomainName string     `json:"domain_name"`
		M2M        *M2MLogin  `json:"m2m,omitempty"`
		User       *UserLogin `json:"user,omitempty"`
	}

	// M2MLogin records an M2M login. The client secret is kept in the
	// session store, never in the config file.
	M2MLogin struct {
		ClientID string `json:"client_id"`
	}

	// UserLogin records a browser login. ClientID is empty when Kinde's
	// default application for the device flow is used.
	UserLogin struct {
		ClientID string `json:"client_id,omitempty"`
	}
)

// sessionStoreOptions are extra options applied to every session store.
// Tests set them with UseFileSessionStore so secrets never reach the OS keychain.
var sessionStoreOptions []cli.Option

// UseFileSessionStore makes session stores use a password-encrypted file in
// the config directory instead of the OS keychain. It returns a function that
// restores the default. Intended for tests.
func UseFileSessionStore(password string) (restore func()) {
	previous := sessionStoreOptions
	sessionStoreOptions = []cli.Option{
		cli.WithAllowedBackends([]keyring.BackendType{keyring.FileBackend}),
		cli.WithFilePasswordFunc(keyring.FixedStringPrompt(password)),
	}
	return func() { sessionStoreOptions = previous }
}

func (env *Environment) getCliSession(identity Identity) (cli.ICliSession, error) {
	storeName := fmt.Sprintf("%v_%v", env.DomainName, identity)

	chainFileName, err := env.keychainFolderName(storeName)
	if err != nil {
		return nil, err
	}
	opts := append([]cli.Option{cli.WithFileDir(chainFileName)}, sessionStoreOptions...)
	cliSession, err := cli.NewCliSession(fmt.Sprintf("%v_%v", CLI_NAME, storeName), opts...)
	return cliSession, err
}

func (env *Environment) kindeDomain() string {
	return fmt.Sprintf("https://%s", env.DomainName)
}

func (env *Environment) newClientCredentialsFlow(clientID, clientSecret string, extra ...client_credentials.Option) (client_credentials.IClientCredentialsFlow, error) {
	opts := append([]client_credentials.Option{
		client_credentials.WithAudience(fmt.Sprintf("%s/api", env.kindeDomain())),
		client_credentials.WithTokenValidation(
			true,
			jwt.WillValidateAlgorithm(),
			jwt.WillValidateIssuer(env.kindeDomain()),
		),
	}, extra...)

	return client_credentials.NewClientCredentialsFlow(env.kindeDomain(), clientID, clientSecret, opts...)
}

// VerifyClientCredentials requests a token with the given credentials without
// saving anything, so failed logins leave any existing login untouched.
func (env *Environment) VerifyClientCredentials(ctx context.Context, clientID, clientSecret string) error {
	flow, err := env.newClientCredentialsFlow(clientID, clientSecret)
	if err != nil {
		return err
	}
	token, err := flow.GetToken(ctx)
	if err != nil {
		return err
	}
	if !token.IsValid() {
		return errors.New("received an invalid token")
	}
	return nil
}

// SaveM2MSecret stores the client secret in the M2M session store and clears
// any token cached for a previous M2M login. The caller records the client ID
// in the config.
func (env *Environment) SaveM2MSecret(clientSecret string) error {
	session, err := env.getCliSession(IdentityM2M)
	if err != nil {
		return fmt.Errorf("failed to open M2M session store: %w", err)
	}
	if err := clearToken(session); err != nil {
		return err
	}
	if err := session.SetKey(clientSecretKey, []byte(clientSecret)); err != nil {
		return fmt.Errorf("failed to store client secret: %w", err)
	}
	return nil
}

// m2mCredentialsFromEnv returns M2M credentials set through environment
// variables. ok is false when neither variable is set.
func m2mCredentialsFromEnv() (clientID, clientSecret string, ok bool, err error) {
	clientID = os.Getenv(EnvM2MClientID)
	clientSecret = os.Getenv(EnvM2MClientSecret)
	switch {
	case clientID == "" && clientSecret == "":
		return "", "", false, nil
	case clientID == "" || clientSecret == "":
		return "", "", false, fmt.Errorf("%s and %s must both be set", EnvM2MClientID, EnvM2MClientSecret)
	}
	return clientID, clientSecret, true, nil
}

// M2MClientCredentialsFlow returns a flow for the M2M credentials in the
// environment if set, otherwise for the saved M2M login. Environment
// credentials never touch the session store, so they work on machines without
// a keychain, such as CI runners.
func (env *Environment) M2MClientCredentialsFlow() (client_credentials.IClientCredentialsFlow, error) {
	clientID, clientSecret, ok, err := m2mCredentialsFromEnv()
	if err != nil {
		return nil, err
	}
	if ok {
		return env.newClientCredentialsFlow(clientID, clientSecret)
	}

	if env.M2M == nil {
		return nil, ErrNoM2MLogin
	}
	session, err := env.getCliSession(IdentityM2M)
	if err != nil {
		return nil, fmt.Errorf("failed to open M2M session store: %w", err)
	}
	secret, err := session.GetKey(clientSecretKey)
	if err != nil {
		return nil, fmt.Errorf("%w: client secret not found in session store", ErrNoM2MLogin)
	}
	return env.newClientCredentialsFlow(env.M2M.ClientID, string(secret), client_credentials.WithSessionHooks(session))
}

// ClearM2MLogin removes the stored client secret and cached token.
func (env *Environment) ClearM2MLogin() error {
	session, err := env.getCliSession(IdentityM2M)
	if err != nil {
		return fmt.Errorf("failed to open M2M session store: %w", err)
	}
	return errors.Join(clearToken(session), deleteKey(session, clientSecretKey))
}

// NewDeviceAuthorizationFlow returns a browser (device) login flow backed by
// the user session store. An empty clientID uses Kinde's default application.
func (env *Environment) NewDeviceAuthorizationFlow(clientID string) (authorization_code.IDeviceAuthorizationFlow, error) {
	session, err := env.getCliSession(IdentityUser)
	if err != nil {
		return nil, fmt.Errorf("failed to open user session store: %w", err)
	}

	opts := []authorization_code.Option{
		authorization_code.WithSessionHooks(session),
		authorization_code.WithOffline(),
		authorization_code.WithTokenValidation(
			true,
			jwt.WillValidateAlgorithm(),
			jwt.WillValidateIssuer(env.kindeDomain()),
		),
	}

	if clientID != "" {
		opts = append(opts, authorization_code.WithClientID(clientID))
	}

	return authorization_code.NewDeviceAuthorizationFlow(env.kindeDomain(), opts...)
}

// UserDeviceAuthorizationFlow returns a flow for the saved browser login.
func (env *Environment) UserDeviceAuthorizationFlow() (authorization_code.IDeviceAuthorizationFlow, error) {
	if env.User == nil {
		return nil, ErrNoUserLogin
	}
	return env.NewDeviceAuthorizationFlow(env.User.ClientID)
}

// ClearUserLogin removes the browser login's cached tokens.
func (env *Environment) ClearUserLogin() error {
	session, err := env.getCliSession(IdentityUser)
	if err != nil {
		return fmt.Errorf("failed to open user session store: %w", err)
	}
	return clearToken(session)
}

// clearToken removes the cached OAuth token from a session store.
func clearToken(session cli.ICliSession) error {
	if err := session.SetRawToken(nil); err != nil {
		// kinde-go reports an error with the file backend even when the token
		// was removed, so only fail if it is still there.
		if _, getErr := session.GetRawToken(); getErr == nil {
			return fmt.Errorf("failed to clear cached token: %w", err)
		}
	}
	return nil
}

// deleteKey removes a key from a session store.
func deleteKey(session cli.ICliSession, key string) error {
	if err := session.DeleteKey(key); err != nil {
		// See clearToken: only fail if the key is still there.
		if _, getErr := session.GetKey(key); getErr == nil {
			return fmt.Errorf("failed to remove %s: %w", key, err)
		}
	}
	return nil
}

func (c *Environment) keychainFolderName(fileName string) (string, error) {

	fileName = fmt.Sprintf("kc_%s", normalizeServiceName(fileName))

	configLocation, err := Dir()
	if err != nil {
		return "", err
	}
	configLocation = filepath.Join(configLocation, fileName)

	// Ensure per‐domain dir exists too
	if err := os.MkdirAll(configLocation, 0700); err != nil {
		return "", err
	}

	return configLocation, nil
}

func normalizeServiceName(name string) string {
	// Replace special characters and spaces that could cause issues in keychain
	normalized := strings.ReplaceAll(name, "/", "_")
	normalized = strings.ReplaceAll(normalized, ":", "_")
	normalized = strings.ReplaceAll(normalized, ".", "_")
	normalized = strings.ReplaceAll(normalized, " ", "_")
	return normalized
}
