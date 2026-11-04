package cmd

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/spf13/cobra"
)

const resourceGroupID = "resources"

// newResourceCommands builds one top-level command per entry in
// resourceGroups, each with a subcommand per Management API operation.
func newResourceCommands() ([]*cobra.Command, error) {
	var groups []*cobra.Command

	for _, group := range resourceGroups {
		groupCmd := &cobra.Command{
			Use:     group.name,
			Args:    cobra.NoArgs,
			Short:   group.short,
			GroupID: resourceGroupID,
		}

		for _, op := range group.commands {
			command, err := buildCobraCommand(op)
			if err != nil {
				return nil, fmt.Errorf("failed to build command %s %s: %w", group.name, op.name, err)
			}
			groupCmd.AddCommand(command)
		}

		groups = append(groups, groupCmd)
	}

	return groups, nil
}

func buildCobraCommand(op resourceCommand) (*cobra.Command, error) {
	plan, err := planForOperation(op.operation)
	if err != nil {
		return nil, err
	}

	doc := docsFor(op.operation)
	short := doc.Summary
	if short == "" {
		short = fmt.Sprintf("Management API operation %s", op.operation)
	}

	command := &cobra.Command{
		Use:   op.name,
		Args:  cobra.NoArgs,
		Short: short,
		Long:  doc.long(),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read the environment when the command runs, after --domain is applied.
			env := config.FromContext[config.Config](cmd.Context()).GetEnvironment()
			if env == nil {
				return errNoEnvironment
			}
			methodArgs, err := plan.buildArgs(cmd)
			if err != nil {
				return err
			}
			return callApiMethod(cmd, env, op, methodArgs)
		},
	}
	plan.registerFlags(command.Flags(), doc.flag)

	return command, nil
}

// planForOperation builds the request plan for a Management API client method.
func planForOperation(operation management_api.OperationName) (*requestPlan, error) {
	method, found := reflect.TypeOf(&management_api.Client{}).MethodByName(operation)
	if !found {
		return nil, fmt.Errorf("operation %s not found in management API", operation)
	}
	// Skip the receiver and the context.
	var args []reflect.Type
	for i := 2; i < method.Type.NumIn(); i++ {
		args = append(args, method.Type.In(i))
	}
	return newRequestPlan(args)
}

func callApiMethod(cmd *cobra.Command, env *config.Environment, op resourceCommand, methodArgs []reflect.Value) error {
	clientCredentials, err := env.M2MClientCredentialsFlow()
	if errors.Is(err, config.ErrNoM2MLogin) {
		return fmt.Errorf("%w (%s). Management commands need an M2M application: run 'kinde login --client_id <id> --client_secret <secret>', or set %s and %s",
			err, env.DomainName, config.EnvM2MClientID, config.EnvM2MClientSecret)
	}
	if err != nil {
		return fmt.Errorf("failed to create client credentials flow: %w", err)
	}

	httpClient, err := clientCredentials.GetClient(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to create HTTP client: %w", err)
	}
	accessToken := func(ctx context.Context) (string, error) {
		token, err := clientCredentials.GetToken(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to get an access token: %w", err)
		}
		raw := token.GetRawToken()
		if raw == nil {
			return "", errors.New("failed to get an access token: the token is empty")
		}
		return raw.AccessToken, nil
	}

	return callOperation(cmd, "https://"+env.DomainName, httpClient, accessToken, op.operation, methodArgs)
}

// flagNameFor returns the kebab-case flag name for a request field, prefixed
// with its parent's flag name for nested fields ("parent.child-name").
func flagNameFor(field reflect.StructField, prefix string) string {
	name, _, _ := strings.Cut(field.Tag.Get("json"), ",")
	if name == "" {
		name = toSnakeCase(field.Name)
	}
	name = strings.ReplaceAll(name, "_", "-")
	if prefix != "" {
		name = prefix + "." + name
	}
	return name
}

func toSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// 1. Insert underscore before each uppercase letter that is preceded by a lowercase
	re1 := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	s = re1.ReplaceAllString(s, "${1}_${2}")

	// 2. Handle consecutive capitals (e.g., “HTTPServer” -> “http_server”)
	re2 := regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)
	s = re2.ReplaceAllString(s, "${1}_${2}")

	// 3. Convert the whole string to lower case
	return strings.ToLower(s)
}
