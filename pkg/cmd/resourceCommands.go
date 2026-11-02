package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime"
	"net/textproto"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	visitor "github.com/kinde-oss/kinde-cli/pkg/reflectVisitor"
	"github.com/kinde-oss/kinde-go/kinde"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/ogen-go/ogen/http"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
	apiMethod, found := reflect.TypeOf(&management_api.Client{}).MethodByName(op.operation)
	if !found {
		return nil, fmt.Errorf("operation %s not found in management API", op.operation)
	}

	command := &cobra.Command{
		Use:   op.name,
		Args:  cobra.NoArgs,
		Short: fmt.Sprintf("Management API operation %s", op.operation),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read the environment when the command runs, after --domain is applied.
			env := config.FromContext[config.Config](cmd.Context()).GetEnvironment()
			if env == nil {
				return errNoEnvironment
			}
			return callApiMethod(cmd, env, op)
		},
	}

	numArgs := apiMethod.Type.NumIn()
	var methodArgs []reflect.Value

	for i := 1; i < numArgs; i++ {
		argumentType := reflect.Zero(apiMethod.Type.In(i))
		methodArgs = append(methodArgs, argumentType)
		generateCommandFlags(argumentType.Type(), command)
	}

	return command, nil
}

func generateCommandFlags(t reflect.Type, command *cobra.Command) {

	visitor.NewVisitor(t).
		Visit(nil, func(p visitor.Walker[reflect.Type]) []visitor.Walker[reflect.Type] {
			visits := []visitor.Walker[reflect.Type]{}
			for i := range p.T1.NumField() {
				field := p.T1.Field(i)
				flagName := flagNameFor(field, p.T2)

				switch field.Type {
				case reflect.TypeOf(management_api.OptNilBool{}), reflect.TypeOf(management_api.OptBool{}):
					command.Flags().Bool(flagName, false, "")
				case reflect.TypeOf(management_api.OptNilInt{}), reflect.TypeOf(management_api.OptInt{}):
					command.Flags().Int(flagName, 0, "")
				case reflect.TypeOf(management_api.OptString{}), reflect.TypeOf(management_api.OptNilString{}), reflect.TypeOf(""):
					command.Flags().String(flagName, "", "")
				case reflect.TypeOf(http.MultipartFile{}):
					command.Flags().String(flagName, "", "")
				default:
					resolvedFieldType := field.Type.String()
					if field.Type.Kind() == reflect.String {
						command.Flags().String(flagName, "", "")
					} else if strings.HasPrefix(resolvedFieldType, "management_api.Opt") {
						log.Trace().Msgf("%s - mapping type for field %s: %s", t.Name(), field.Name, field.Type.String())
						var i reflect.Value
						if p.OptionalSetter.Value != nil && p.OptionalSetter.Value.IsValid() {
							i = p.OptionalSetter.Value.FieldByName(field.Name)
						}
						visits = append(visits, visitor.Walker[reflect.Type]{T1: field.Type, T2: flagName, OptionalSetter: visitor.OptSetter{Value: &i}})
					} else {
						log.Trace().Msgf("%s - skipping unsupported type for field %s: %s", t.Name(), field.Name, field.Type.String())
					}
				}
			}
			return visits
		})

}

func callApiMethod(cmd *cobra.Command, env *config.Environment, op resourceCommand) error {
	log := log.Ctx(cmd.Context())
	ctx := cmd.Context()

	clientCredentials, err := env.M2MClientCredentialsFlow()
	if errors.Is(err, config.ErrNoM2MLogin) {
		return fmt.Errorf("%w (%s). Management commands need an M2M application: run 'kinde login --client_id <id> --client_secret <secret>', or set %s and %s",
			err, env.DomainName, config.EnvM2MClientID, config.EnvM2MClientSecret)
	}
	if err != nil {
		return fmt.Errorf("failed to create client credentials flow: %w", err)
	}

	kindeDomainUrl := fmt.Sprintf("https://%s", env.DomainName)
	managementApi, err := kinde.NewManagementAPI(ctx, kindeDomainUrl, clientCredentials)
	if err != nil {
		return fmt.Errorf("failed to create management API client: %w", err)
	}

	instanceMethod := reflect.ValueOf(managementApi).MethodByName(op.operation)

	methodType := instanceMethod.Type()
	inParams := methodType.NumIn()
	args := []reflect.Value{}
	for i := range inParams {
		if methodType.In(i) == reflect.TypeOf((*context.Context)(nil)).Elem() {
			args = append(args, reflect.ValueOf(ctx))
			continue
		}
		argInstance := mapFlagsToInstance(methodType.In(i), cmd.Flags())

		if methodType.In(i).Kind() != reflect.Ptr && methodType.In(i).Kind() != reflect.Interface {
			argInstance = reflect.ValueOf(argInstance).Elem().Interface()
		}
		args = append(args, reflect.ValueOf(argInstance))

	}

	results := instanceMethod.Call(args[0:inParams])
	for _, result := range results {
		if !result.IsValid() || (result.Kind() == reflect.Ptr || result.Kind() == reflect.Interface) && result.IsNil() {
			continue
		}
		resultInstance := result.Interface()
		if error, ok := resultInstance.(error); ok {
			return error
		}
		log.Info().Any("result", resultInstance).Msg("API call result")
	}

	return nil
}

func mapFlagsToInstance(t reflect.Type, flagSet *pflag.FlagSet) any {

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	instance := reflect.New(t).Interface()
	val := reflect.Indirect(reflect.ValueOf(instance))
	if t.Kind() != reflect.Struct {
		return instance
	}

	visitor.NewVisitor(t).
		Visit(&val, func(p visitor.Walker[reflect.Type]) []visitor.Walker[reflect.Type] {
			additionalVisits := []visitor.Walker[reflect.Type]{}
			for i := range p.T1.NumField() {
				field := p.T1.Field(i)
				flagName := flagNameFor(field, p.T2)

				flag := flagSet.Lookup(flagName)
				if flag == nil {
					possibleName := strings.ReplaceAll(toSnakeCase(field.Name), "_", "-")
					flag = flagSet.Lookup(possibleName)
				}

				log.Trace().Msgf("Mapping field %s with flag %s, val type %s", field.Name, flagName, val.Type().String())
				instanceField := p.OptionalSetter.Value.FieldByName(field.Name)

				set := func(v any) {
					instanceField.Set(reflect.ValueOf(v))
					if p.OptionalSetter.IsSet.IsValid() {
						p.OptionalSetter.IsSet.SetBool(true)
					}

				}

				switch field.Type {
				case reflect.TypeOf(""):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetString(flag.Name); err == nil {
							set(flagValue)
						}
					}
				case reflect.TypeOf(management_api.OptString{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetString(flag.Name); err == nil {
							set(management_api.NewOptString(flagValue))
						}
					}
				case reflect.TypeOf(management_api.OptNilString{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetString(flag.Name); err == nil {
							set(management_api.NewOptNilString(flagValue))
						}
					}
				case reflect.TypeOf(management_api.OptBool{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetBool(flag.Name); err == nil {
							set(management_api.NewOptBool(flagValue))
						}
					}
				case reflect.TypeOf(management_api.OptNilBool{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetBool(flag.Name); err == nil {
							set(management_api.NewOptNilBool(flagValue))
						}
					}
				case reflect.TypeOf(management_api.OptInt{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetInt(flag.Name); err == nil {
							set(management_api.NewOptInt(flagValue))
						}
					}
				case reflect.TypeOf(management_api.OptNilInt{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetInt(flag.Name); err == nil {
							set(management_api.NewOptNilInt(flagValue))
						}
					}
				case reflect.TypeOf(http.MultipartFile{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetString(flag.Name); err == nil {
							fileName := filepath.Base(flagValue)
							mimeHeader := textproto.MIMEHeader{}
							contentType := mime.TypeByExtension(filepath.Ext(fileName))
							if contentType == "" {
								contentType = "application/octet-stream"
							}
							mimeHeader.Set("Content-Type", contentType)
							data, err := os.ReadFile(flagValue)
							if err != nil {
								log.Error().Err(err).Msgf("Failed to read file: %s", flagValue)
								break
							}
							v := http.MultipartFile{
								Name:   fileName,
								Header: mimeHeader,
								File:   bytes.NewReader(data),
							}

							set(v)
						}
					}

				default:
					if field.Type.Kind() == reflect.String {
						if flag != nil && flag.Changed {
							if flagValue, err := flagSet.GetString(flag.Name); err == nil {
								v := reflect.ValueOf(flagValue).Convert(field.Type).Interface() // Convert to the actual field type from string
								set(v)
							}
						}
					} else if strings.HasPrefix(field.Type.String(), "management_api.Opt") {
						subProperty := instanceField
						log.Trace().Msgf("%s - mapping type for field %s: %s: %s", t.Name(), field.Name, field.Type.String(), subProperty.Type().String())
						additionalVisits = append(additionalVisits, visitor.Walker[reflect.Type]{T1: field.Type, T2: flagName, OptionalSetter: visitor.OptSetter{Value: &subProperty}})
					} else {
						log.Trace().Msgf("%s - skipping unsupported type for field %s: %s", t.Name(), field.Name, field.Type.String())
					}
				}
			}
			return additionalVisits
		})

	log.Debug().Any("post_mapping", instance).Msgf("Mapped instance for type %s", t.Name())
	return instance
}

// flagNameFor returns the kebab-case flag name for a request field, prefixed
// with its parent's flag name for nested fields ("parent.child-name").
func flagNameFor(field reflect.StructField, prefix string) string {
	name := field.Tag.Get("json")
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
