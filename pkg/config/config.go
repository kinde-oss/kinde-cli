package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	CLI_NAME                    = "kinde"
	USER_AGENT                  = "Kinde CLI"
	ConfigContextKey contextKey = "config"
)

// Environment variables the CLI reads. The names match Kinde's Management API
// SDK (@kinde/management-api-js) so the same values work for both.
const (
	EnvDomain          = "KINDE_DOMAIN"
	EnvM2MClientID     = "KINDE_MANAGEMENT_CLIENT_ID"
	EnvM2MClientSecret = "KINDE_MANAGEMENT_CLIENT_SECRET"
)

type (
	contextKey string

	// used for testing
	IConfig interface {
		Validate() (error, bool)
		UseDomain(domain string) error
		SetDefaultDomain(domain string)
		SaveIfChanged() error
		GetEnvironment() *Environment
		SetEnvironment(f func(e *Environment)) error
	}
	Config struct {
		// CurrentEnvironment is the saved default domain.
		CurrentEnvironment string                 `json:"current"`
		Environments       map[string]Environment `json:"environments"`

		// active is the domain used for this run when it differs from the
		// default (set from --domain or KINDE_DOMAIN). It is never saved.
		active  string
		changed bool
	}
)

func (i *Config) activeDomain() string {
	if i.active != "" {
		return i.active
	}
	return i.CurrentEnvironment
}

// SetEnvironment changes the active domain's saved settings.
func (i *Config) SetEnvironment(f func(e *Environment)) error {
	env := i.GetEnvironment()
	if env == nil {
		return errors.New("no Kinde domain selected")
	}
	f(env)
	if i.Environments == nil {
		i.Environments = make(map[string]Environment)
	}
	i.Environments[env.DomainName] = *env
	i.changed = true
	return nil
}

// GetEnvironment returns a copy of the active domain's settings, or nil if no
// domain is selected.
func (i *Config) GetEnvironment() *Environment {
	domain := i.activeDomain()
	if domain == "" {
		return nil
	}
	if env, ok := i.Environments[domain]; ok {
		return &env
	}
	return &Environment{DomainName: domain}
}

func (c *Config) Validate() (error, bool) {
	return nil, true
}

// NewConfig initializes a new Config instance, reading from the config file if it exists.
func NewConfig() (IConfig, error) {
	config := &Config{
		Environments: make(map[string]Environment),
	}
	if _, err := config.readConfig(); err != nil {
		return nil, err
	}
	return config, nil
}

// Ctx returns a new context with the Config instance stored in it.
func Ctx(ctx context.Context, config IConfig) context.Context {
	return context.WithValue(ctx, ConfigContextKey, config)
}

// NewContext retrieves the Config instance from the context.
func FromContext[T any](ctx context.Context) *T {
	if config, ok := ctx.Value(ConfigContextKey).(*T); ok {
		return config
	}
	return nil
}

// UseDomain selects the domain for this run only; it does not change the
// saved default.
func (c *Config) UseDomain(domain string) error {
	normalized, err := NormalizeDomain(domain)
	if err != nil {
		return err
	}
	c.active = normalized
	return nil
}

// SetDefaultDomain makes domain the saved default for future runs.
func (c *Config) SetDefaultDomain(domain string) {
	if c.CurrentEnvironment != domain {
		c.CurrentEnvironment = domain
		c.changed = true
	}
}

// NormalizeDomain accepts a Kinde domain with or without a scheme or trailing
// slash ("https://mybiz.kinde.com/") and returns the bare host in lower case.
func NormalizeDomain(domain string) (string, error) {
	d := strings.ToLower(strings.TrimSpace(domain))
	d = strings.TrimPrefix(d, "https://")
	d = strings.TrimPrefix(d, "http://")
	d = strings.TrimSuffix(d, "/")
	if d == "" || strings.ContainsAny(d, "/ ") {
		return "", fmt.Errorf("invalid Kinde domain %q: expected something like mybiz.kinde.com", domain)
	}
	return d, nil
}

func (c *Config) detectConfigFileName() (string, error) {
	configLocation := os.Getenv("XDG_CONFIG_HOME")
	if configLocation == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configLocation = filepath.Join(home, ".config")
	}
	configLocation = filepath.Join(configLocation, CLI_NAME)

	err := os.MkdirAll(configLocation, 0700)
	if err != nil {
		return "", err
	}

	configLocation = filepath.Join(configLocation, "config.json")

	return configLocation, nil
}

// SaveIfChanged writes the config file if a command changed the settings, so
// read-only commands never touch it.
func (c *Config) SaveIfChanged() error {
	if !c.changed {
		return nil
	}

	configFileLocation, err := c.detectConfigFileName()
	if err != nil {
		return fmt.Errorf("failed to detect config file name: %w", err)
	}

	buf, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(configFileLocation, buf, 0600)
	if err != nil {
		return fmt.Errorf("failed to write config file to %s: %w", configFileLocation, err)
	}

	return nil
}

func (c *Config) readConfig() (bool, error) {
	configFileLocation, err := c.detectConfigFileName()

	//check file exists
	if _, err := os.Stat(configFileLocation); os.IsNotExist(err) {
		return false, nil // config file does not exist, return false
	}

	if err != nil {
		return false, fmt.Errorf("failed to detect config file name: %w", err)
	}

	buf, err := os.ReadFile(configFileLocation)
	if err != nil {
		return false, fmt.Errorf("failed to read config file %s: %w", configFileLocation, err)
	}

	err = json.Unmarshal(buf, c)
	if err != nil {
		return false, fmt.Errorf("config file %s is not valid JSON: %w", configFileLocation, err)
	}

	return true, nil
}
