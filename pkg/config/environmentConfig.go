package config

import (
	"fmt"

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

func (env *Environment) getSession() (authorization_code.ISessionHooks, error) {
	return cli.NewCliSession(fmt.Sprintf("%v_%v", CLI_NAME, env.DomainName))

}

func (env *Environment) NewClientCredentialsFlow() (client_credentials.IClientCredentialsFlow, error) {
	kindeDomain := fmt.Sprintf("https://%s", env.DomainName)

	cliSession, err := env.getSession()

	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	opts := []client_credentials.Option{
		client_credentials.WithSessionHooks(cliSession),
		client_credentials.WithTokenValidation(
			true,
			jwt.WillValidateAlgorithm(),
			jwt.WillValidateIssuer(kindeDomain),
		),
	}

	deviceFlow, err := client_credentials.NewClientCredentialsFlow(
		kindeDomain,
		env.ClientID,
		env.ClientSecret,
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

	cliSession, err := env.getSession()

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
