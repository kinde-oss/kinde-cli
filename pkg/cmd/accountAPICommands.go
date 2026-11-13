package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-go/oauth2/authorization_code"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// accountOperations are the Account API operations under 'kinde account'.
// They read what the logged-in user can see about themselves. The spec's
// OAuth operations (introspect, revoke) are for applications, not people.
var accountOperations = []struct{ name, operationID string }{
	{"entitlement", "GetEntitlement"},
	{"entitlements", "GetEntitlements"},
	{"feature-flags", "GetFeatureFlags"},
	{"permissions", "GetUserPermissions"},
	{"portal-link", "GetPortalLink"},
	{"profile", "getUserProfileV2"},
	{"properties", "GetUserProperties"},
	{"roles", "GetUserRoles"},
}

func accountDocsFor(operationID string) operationDoc {
	return accountAPIDocs[strings.ToLower(operationID)]
}

// newAccountAPICommands builds a command per Account API operation, with a
// flag per parameter in the spec.
func newAccountAPICommands() ([]*cobra.Command, error) {
	var commands []*cobra.Command
	for _, op := range accountOperations {
		doc := accountDocsFor(op.operationID)
		if doc.Path == "" {
			return nil, fmt.Errorf("account operation %s is missing from the Account API spec", op.operationID)
		}
		commands = append(commands, newAccountAPICommand(op.name, doc))
	}
	return commands, nil
}

func newAccountAPICommand(name string, doc operationDoc) *cobra.Command {
	fields := accountFields(doc)
	command := &cobra.Command{
		Use:   name,
		Args:  cobra.NoArgs,
		Short: doc.Summary,
		Long:  doc.long(),
		RunE: func(cmd *cobra.Command, args []string) error {
			requestURL, err := accountRequestURL(doc, fields, cmd.Flags())
			if err != nil {
				return err
			}

			env := config.FromContext[config.Config](cmd.Context()).GetEnvironment()
			if env == nil {
				return errNoEnvironment
			}
			flow, err := userLoginFlow(env)
			if err != nil {
				return err
			}
			httpClient, err := flow.GetClient(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to create HTTP client: %w", err)
			}
			return callAccountAPI(cmd, httpClient, "https://"+env.DomainName+requestURL)
		},
	}
	for _, field := range fields {
		usage := field.usage(doc.flag(field.flag))
		switch field.kind {
		case kindInt:
			command.Flags().Int64(field.flag, 0, usage)
		case kindBool:
			command.Flags().Bool(field.flag, false, usage)
		default:
			command.Flags().String(field.flag, "", usage)
		}
	}
	return command
}

// accountField is a requestField for one of the operation's parameters.
type accountField struct {
	requestField
	param paramDoc
}

func accountFields(doc operationDoc) []accountField {
	var fields []accountField
	for _, param := range doc.Params {
		field := requestField{
			flag:     strings.ReplaceAll(param.Name, "_", "-"),
			required: param.Required || param.In == "path",
			enum:     param.Enum,
		}
		switch param.Type {
		case "integer":
			field.kind = kindInt
		case "boolean":
			field.kind = kindBool
		default:
			field.kind = kindString
		}
		fields = append(fields, accountField{requestField: field, param: param})
	}
	return fields
}

// accountRequestURL builds the request's path and query string from the
// flags, checking required flags and allowed values first.
func accountRequestURL(doc operationDoc, fields []accountField, flags *pflag.FlagSet) (string, error) {
	path := doc.Path
	query := url.Values{}

	var missing []string
	for _, field := range fields {
		if !flags.Changed(field.flag) {
			if field.required {
				missing = append(missing, "--"+field.flag)
			}
			continue
		}
		value := flags.Lookup(field.flag).Value.String()
		if err := field.checkEnum(value); err != nil {
			return "", fmt.Errorf("invalid value for --%s: %w", field.flag, err)
		}

		if field.param.In != "path" {
			query.Set(field.param.Name, value)
			continue
		}
		placeholder := "{" + field.param.Name + "}"
		if strings.Contains(path, placeholder) {
			path = strings.ReplaceAll(path, placeholder, url.PathEscape(value))
		} else {
			// The spec declares GetEntitlement's key as a path parameter
			// without a placeholder; the API (and kinde-go) expect it last.
			path += "/" + url.PathEscape(value)
		}
	}
	if len(missing) > 0 {
		return "", fmt.Errorf("missing required flags: %s", strings.Join(missing, ", "))
	}

	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	return path, nil
}

// userLoginFlow returns the user login for the domain, explaining how to
// log in when there isn't one.
func userLoginFlow(env *config.Environment) (authorization_code.IDeviceAuthorizationFlow, error) {
	flow, err := env.UserDeviceAuthorizationFlow()
	if errors.Is(err, config.ErrNoUserLogin) {
		return nil, fmt.Errorf("%w (%s). Account commands act as a user: run 'kinde account login' first", err, env.DomainName)
	}
	return flow, err
}

// newAccountTokenCmd prints the user login's access token, for calling the
// Account API (or your own APIs) with other tools.
func newAccountTokenCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "token",
		Args:  cobra.NoArgs,
		Short: "Print the access token of the user login",
		Long: `Print the access token of the user login ('kinde account login').

Use it to call the Account API, or your own APIs, with other tools. Treat it like
a password: anyone with the token can act as you until it expires.`,
		Example: `  curl -H "Authorization: Bearer $(kinde account token)" https://mybusiness.kinde.com/account_api/v1/permissions`,
		RunE: func(cmd *cobra.Command, args []string) error {
			env := config.FromContext[config.Config](cmd.Context()).GetEnvironment()
			if env == nil {
				return errNoEnvironment
			}
			flow, err := userLoginFlow(env)
			if err != nil {
				return err
			}
			token, err := flow.GetToken(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to get token: %w", err)
			}
			accessToken, ok := token.GetAccessToken()
			if !ok {
				return errors.New("the user login has no access token: run 'kinde account login' again")
			}
			fmt.Fprintln(cmd.OutOrStdout(), accessToken)
			return nil
		},
	}
}
