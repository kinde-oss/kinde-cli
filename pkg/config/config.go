package config

import (
	"context"
	"encoding/json"
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

type (
	contextKey string

	// used for testing
	IConfig interface {
		Validate() (error, bool)
		SwitchEnvironment(val string) error
		PersistConfig() error
		GetEnvironment() *Environment
		SetEnvironment(f func(e *Environment)) error
	}
	Config struct {
		CurrentEnvironment string                 `json:"current"`
		Environments       map[string]Environment `json:"environments"`
	}
)

func (i *Config) SetEnvironment(f func(e *Environment)) error {
	env := i.GetEnvironment()
	f(env)
	i.Environments[i.CurrentEnvironment] = *env
	return nil
}

// GetEnvironment implements IConfig.
func (i *Config) GetEnvironment() *Environment {
	if env, ok := i.Environments[i.CurrentEnvironment]; !ok {
		return nil
	} else {
		return &env
	}
}

func (c *Config) Validate() (error, bool) {
	if c.CurrentEnvironment == "" {
		return fmt.Errorf("no environment configured, please use `kinde login`"), false
	}
	return nil, true
}

// NewConfig initializes a new Config instance, reading from the config file if it exists.
func NewConfig() (IConfig, error) {
	config := &Config{
		Environments: make(map[string]Environment),
	}
	if isRead, err := config.readConfig(); isRead && err == nil {
		config.SwitchEnvironment(config.CurrentEnvironment)
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

// SwitchEnvironment sets the current environment to the specified value, creating it if it doesn't exist.
func (c *Config) SwitchEnvironment(val string) error {

	if val == "" {
		return fmt.Errorf("environment name cannot be empty")
	}

	c.CurrentEnvironment = strings.ToLower(val)

	if _, ok := c.Environments[c.CurrentEnvironment]; !ok {
		c.Environments[c.CurrentEnvironment] = Environment{
			DomainName: c.CurrentEnvironment,
		}
	}

	c.PersistConfig()

	return nil
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

// PersistConfig writes the current configuration to the config file.
func (c *Config) PersistConfig() error {
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
		return false, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return true, nil
}
