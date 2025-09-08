package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kinde-oss/kinde-go/frameworks/cli"
	"github.com/kinde-oss/kinde-go/jwt"
	"github.com/kinde-oss/kinde-go/oauth2/authorization_code"
	"github.com/kinde-oss/kinde-go/oauth2/client_credentials"
)

type (
	IEnvironment interface {
		NewDeviceAuthorizationFlow() (authorization_code.IDeviceAuthorizationFlow, error)
		NewClientCredentialsFlow() (client_credentials.IClientCredentialsFlow, error)
	}

	Environment struct {
		DomainName   string `json:"domain_name"`
		ClientID     string `json:"client_id,omitempty"`
		ClientSecret string `json:"client_secret,omitempty"`
	}
)

func (env *Environment) getCliSession() (cli.ICliSession, error) {

	chainFileName, err := env.keychainFolderName(env.DomainName)
	if err != nil {
		return nil, err
	}
	cliSession, err := cli.NewCliSession(fmt.Sprintf("%v_%v", CLI_NAME, env.DomainName), cli.WithFileDir(chainFileName))
	return cliSession, err

}

func (env *Environment) NewClientCredentialsFlow() (client_credentials.IClientCredentialsFlow, error) {
	kindeDomain := fmt.Sprintf("https://%s", env.DomainName)

	cliSession, err := env.getCliSession()

	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	// Try to get client secret from session first, fallback to config if not found
	clientSecret := env.ClientSecret
	if secretFromSession, err := env.GetClientSecretFromSession(); err == nil && secretFromSession != "" {
		clientSecret = secretFromSession
	}

	opts := []client_credentials.Option{
		client_credentials.WithSessionHooks(cliSession),
		client_credentials.WithAudience(fmt.Sprintf("%s/api", kindeDomain)),
		client_credentials.WithTokenValidation(
			true,
			jwt.WillValidateAlgorithm(),
			jwt.WillValidateIssuer(kindeDomain),
		),
	}

	deviceFlow, err := client_credentials.NewClientCredentialsFlow(
		kindeDomain,
		env.ClientID,
		clientSecret,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return deviceFlow, nil
}

// NewDeviceAuthorizationFlow creates a new Device Authorization Flow with the given options.
func (env *Environment) NewDeviceAuthorizationFlow() (authorization_code.IDeviceAuthorizationFlow, error) {

	kindeDomain := fmt.Sprintf("https://%s", env.DomainName)

	cliSession, err := env.getCliSession()

	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	opts := []authorization_code.Option{
		authorization_code.WithSessionHooks(cliSession),
		authorization_code.WithOffline(),
		authorization_code.WithTokenValidation(
			true,
			jwt.WillValidateAlgorithm(),
			jwt.WillValidateIssuer(kindeDomain),
		),
	}

	if env.ClientID != "" {
		opts = append(opts, authorization_code.WithClientID(env.ClientID))
	}

	deviceFlow, err := authorization_code.NewDeviceAuthorizationFlow(
		kindeDomain,
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return deviceFlow, nil
}

func (c *Environment) keychainFolderName(fileName string) (string, error) {

	fileName = fmt.Sprintf("kc_%s", normalizeServiceName(fileName))

	configLocation := os.Getenv("XDG_CONFIG_HOME")
	if configLocation == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configLocation = filepath.Join(home, ".config")
	}
	configLocation = filepath.Join(configLocation, CLI_NAME)

	// Ensure the base CLI_NAME directory exists
	if err := os.MkdirAll(configLocation, 0700); err != nil {
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

// RedactSecret returns a redacted version of the secret for display purposes
func RedactSecret(secret string) string {
	if secret == "" {
		return ""
	}
	if len(secret) <= 8 {
		return "****"
	}
	return secret[:4] + "****" + secret[len(secret)-4:]
}

// StoreClientSecretInSession stores the client secret in the CLI session
func (env *Environment) StoreClientSecretInSession(secret string) error {
	cliSession, err := env.getCliSession()
	if err != nil {
		return fmt.Errorf("failed to get CLI session: %w", err)
	}

	// Store the secret in the session using a key
	err = cliSession.SetKey("client_secret", []byte(secret))
	if err != nil {
		return fmt.Errorf("failed to store client secret in session: %w", err)
	}

	return nil
}

// GetClientSecretFromSession retrieves the client secret from the CLI session
func (env *Environment) GetClientSecretFromSession() (string, error) {
	cliSession, err := env.getCliSession()
	if err != nil {
		return "", fmt.Errorf("failed to get CLI session: %w", err)
	}

	secretBytes, err := cliSession.GetKey("client_secret")
	if err != nil {
		return "", fmt.Errorf("failed to retrieve client secret from session: %w", err)
	}

	return string(secretBytes), nil
}

// ClearClientSecretFromSession removes the client secret from the CLI session
func (env *Environment) ClearClientSecretFromSession() error {
	cliSession, err := env.getCliSession()
	if err != nil {
		return fmt.Errorf("failed to get CLI session: %w", err)
	}

	// Try to remove the secret from the session using DeleteKey
	err = cliSession.DeleteKey("client_secret")
	if err != nil {
		return fmt.Errorf("failed to remove client secret from session: %w", err)
	}

	return nil
}
