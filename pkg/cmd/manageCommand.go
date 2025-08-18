package cmd

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	visitor "github.com/kinde-oss/kinde-cli/pkg/reflectVisitor"
	"github.com/kinde-oss/kinde-go/kinde"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type (
	manageCmd struct {
		cmd *cobra.Command
	}

	commandOperationPair[string, T any] struct {
		commandName string
		operation   T
	}
)

func newManageCmd(ctx context.Context) *manageCmd {

	manageCmd := &manageCmd{}

	manageCmd.cmd = &cobra.Command{
		Use:   "manage",
		Args:  cobra.NoArgs,
		Short: "Kinde management API commands",
		Long: `Manage Kinde resources using the management API. 
This command allows you to perform various operations such as creating, updating, and deleting resources like applications, organizations, users, roles, and more.
Management API scopes need to be granted to the application you are using to run this command.`,
	}

	ops := map[string][]commandOperationPair[string, management_api.OperationName]{
		"apis": {
			{"add_application_scope", management_api.AddAPIApplicationScopeOperation},
			{"add_scope", management_api.AddAPIScopeOperation},
			{"add_apis", management_api.AddAPIsOperation},
			{"add_logo", management_api.AddLogoOperation},
			{"add_logout_redirect_urls", management_api.AddLogoutRedirectURLsOperation},
			{"add_redirect_callback_urls", management_api.AddRedirectCallbackURLsOperation},
			{"delete_api", management_api.DeleteAPIOperation},
			{"delete_application_scope", management_api.DeleteAPIAppliationScopeOperation},
			{"delete_scope", management_api.DeleteAPIScopeOperation},
			{"delete_callback_urls", management_api.DeleteCallbackURLsOperation},
			{"delete_logout_urls", management_api.DeleteLogoutURLsOperation},
			{"get_api", management_api.GetAPIOperation},
			{"get_scope", management_api.GetAPIScopeOperation},
			{"get_scopes", management_api.GetAPIScopesOperation},
			{"get_apis", management_api.GetAPIsOperation},
			{"get_callback_urls", management_api.GetCallbackURLsOperation},
			{"get_logout_urls", management_api.GetLogoutURLsOperation},
			{"update_applications", management_api.UpdateAPIApplicationsOperation},
			{"update_scope", management_api.UpdateAPIScopeOperation},
			{"replace_logout_redirect_urls", management_api.ReplaceLogoutRedirectURLsOperation},
			{"replace_redirect_callback_urls", management_api.ReplaceRedirectCallbackURLsOperation},
			{"read_logo", management_api.ReadLogoOperation},
			{"delete_logo", management_api.DeleteLogoOperation},
			{"token_introspection", management_api.TokenIntrospectionOperation},
			{"token_revocation", management_api.TokenRevocationOperation},
		},
		"applications": {
			{"create", management_api.CreateApplicationOperation},
			{"delete", management_api.DeleteApplicationOperation},
			{"get", management_api.GetApplicationOperation},
			{"get_connections", management_api.GetApplicationConnectionsOperation},
			{"get_property_values", management_api.GetApplicationPropertyValuesOperation},
			{"get_all", management_api.GetApplicationsOperation},
			{"update", management_api.UpdateApplicationOperation},
			{"update_tokens", management_api.UpdateApplicationTokensOperation},
			{"update_property", management_api.UpdateApplicationsPropertyOperation},
		},
		"business": {
			{"get", management_api.GetBusinessOperation},
			{"update", management_api.UpdateBusinessOperation},
		},
		"organizations": {
			{"create", management_api.CreateOrganizationOperation},
			{"delete", management_api.DeleteOrganizationOperation},
			{"get", management_api.GetOrganizationOperation},
			{"get_connections", management_api.GetOrganizationConnectionsOperation},
			{"get_feature_flags", management_api.GetOrganizationFeatureFlagsOperation},
			{"get_property_values", management_api.GetOrganizationPropertyValuesOperation},
			{"get_user_permissions", management_api.GetOrganizationUserPermissionsOperation},
			{"get_user_roles", management_api.GetOrganizationUserRolesOperation},
			{"get_users", management_api.GetOrganizationUsersOperation},
			{"get_all", management_api.GetOrganizationsOperation},
			{"update", management_api.UpdateOrganizationOperation},
			{"update_properties", management_api.UpdateOrganizationPropertiesOperation},
			{"update_property", management_api.UpdateOrganizationPropertyOperation},
			{"update_sessions", management_api.UpdateOrganizationSessionsOperation},
			{"update_users", management_api.UpdateOrganizationUsersOperation},
			{"add_logo", management_api.AddOrganizationLogoOperation},
			{"delete_logo", management_api.DeleteOrganizationLogoOperation},
			{"read_logo", management_api.ReadOrganizationLogoOperation},
			{"add_user_api_scope", management_api.AddOrganizationUserAPIScopeOperation},
			{"delete_user_api_scope", management_api.DeleteOrganizationUserAPIScopeOperation},
			{"add_users", management_api.AddOrganizationUsersOperation},
			{"remove_user", management_api.RemoveOrganizationUserOperation},
			{"create_user_permission", management_api.CreateOrganizationUserPermissionOperation},
			{"delete_user_permission", management_api.DeleteOrganizationUserPermissionOperation},
			{"create_user_role", management_api.CreateOrganizationUserRoleOperation},
			{"delete_user_role", management_api.DeleteOrganizationUserRoleOperation},
			{"get_user_mfa", management_api.GetOrgUserMFAOperation},
			{"reset_user_mfa", management_api.ResetOrgUserMFAOperation},
			{"reset_user_mfa_all", management_api.ResetOrgUserMFAAllOperation},
			{"delete_handle", management_api.DeleteOrganizationHandleOperation},
			{"update_feature_flag_override", management_api.UpdateOrganizationFeatureFlagOverrideOperation},
			{"delete_feature_flag_override", management_api.DeleteOrganizationFeatureFlagOverrideOperation},
			{"delete_feature_flag_overrides", management_api.DeleteOrganizationFeatureFlagOverridesOperation},
			{"replace_mfa", management_api.ReplaceOrganizationMFAOperation},
		},
		"users": {
			{"create", management_api.CreateUserOperation},
			{"delete", management_api.DeleteUserOperation},
			{"get_data", management_api.GetUserDataOperation},
			{"get_identities", management_api.GetUserIdentitiesOperation},
			{"get_permissions", management_api.GetUserPermissionsOperation},
			{"get_profile_v2", management_api.GetUserProfileV2Operation},
			{"get_properties", management_api.GetUserPropertiesOperation},
			{"get_property_values", management_api.GetUserPropertyValuesOperation},
			{"get_roles", management_api.GetUserRolesOperation},
			{"get_sessions", management_api.GetUserSessionsOperation},
			{"get_all", management_api.GetUsersOperation},
			{"get_mfa", management_api.GetUsersMFAOperation},
			{"search", management_api.SearchUsersOperation},
			{"set_password", management_api.SetUserPasswordOperation},
			{"update", management_api.UpdateUserOperation},
			{"update_feature_flag_override", management_api.UpdateUserFeatureFlagOverrideOperation},
			{"update_properties", management_api.UpdateUserPropertiesOperation},
			{"update_property", management_api.UpdateUserPropertyOperation},
			{"refresh_claims", management_api.RefreshUserClaimsOperation},
			{"reset_mfa", management_api.ResetUsersMFAOperation},
			{"reset_mfa_all", management_api.ResetUsersMFAAllOperation},
			{"create_identity", management_api.CreateUserIdentityOperation},
			{"delete_identity", management_api.DeleteIdentityOperation},
			{"delete_sessions", management_api.DeleteUserSessionsOperation},
			{"replace_mfa", management_api.ReplaceMFAOperation},
			{"update_identity", management_api.UpdateIdentityOperation},
		},
		"roles": {
			{"add_scope", management_api.AddRoleScopeOperation},
			{"create", management_api.CreateRoleOperation},
			{"delete", management_api.DeleteRoleOperation},
			{"delete_scope", management_api.DeleteRoleScopeOperation},
			{"get", management_api.GetRoleOperation},
			{"get_permissions", management_api.GetRolePermissionsOperation},
			{"get_scopes", management_api.GetRoleScopesOperation},
			{"get_all", management_api.GetRolesOperation},
			{"remove_permission", management_api.RemoveRolePermissionOperation},
			{"update_permissions", management_api.UpdateRolePermissionsOperation},
			{"update", management_api.UpdateRolesOperation},
		},
		"permissions": {
			{"create", management_api.CreatePermissionOperation},
			{"delete", management_api.DeletePermissionOperation},
			{"get_all", management_api.GetPermissionsOperation},
			{"update", management_api.UpdatePermissionsOperation},
		},
		"feature_flags": {
			{"create", management_api.CreateFeatureFlagOperation},
			{"delete", management_api.DeleteFeatureFlagOperation},
			{"update", management_api.UpdateFeatureFlagOperation},
		},
		"properties": {
			{"create", management_api.CreatePropertyOperation},
			{"delete", management_api.DeletePropertyOperation},
			{"get_all", management_api.GetPropertiesOperation},
			{"update", management_api.UpdatePropertyOperation},
		},
		"connections": {
			{"create", management_api.CreateConnectionOperation},
			{"delete", management_api.DeleteConnectionOperation},
			{"enable", management_api.EnableConnectionOperation},
			{"get", management_api.GetConnectionOperation},
			{"get_all", management_api.GetConnectionsOperation},
			{"remove", management_api.RemoveConnectionOperation},
			{"replace", management_api.ReplaceConnectionOperation},
			{"update", management_api.UpdateConnectionOperation},
			{"enable_org", management_api.EnableOrgConnectionOperation},
			{"remove_org", management_api.RemoveOrgConnectionOperation},
		},
		"environments": {
			{"get", management_api.GetEnvironmentOperation},
			{"get_feature_flags", management_api.GetEnvironementFeatureFlagsOperation},
			{"update_feature_flag_override", management_api.UpdateEnvironementFeatureFlagOverrideOperation},
			{"delete_feature_flag_override", management_api.DeleteEnvironementFeatureFlagOverrideOperation},
			{"delete_feature_flag_overrides", management_api.DeleteEnvironementFeatureFlagOverridesOperation},
		},
		"environment_variables": {
			{"create", management_api.CreateEnvironmentVariableOperation},
			{"delete", management_api.DeleteEnvironmentVariableOperation},
			{"get", management_api.GetEnvironmentVariableOperation},
			{"get_all", management_api.GetEnvironmentVariablesOperation},
			{"update", management_api.UpdateEnvironmentVariableOperation},
		},
		"categories": {
			{"create", management_api.CreateCategoryOperation},
			{"get_all", management_api.GetCategoriesOperation},
			{"update", management_api.UpdateCategoryOperation},
		},
		"subscribers": {
			{"create", management_api.CreateSubscriberOperation},
			{"get", management_api.GetSubscriberOperation},
			{"get_all", management_api.GetSubscribersOperation},
		},
		"webhooks": {
			{"create", management_api.CreateWebHookOperation},
			{"delete", management_api.DeleteWebHookOperation},
			{"get_all", management_api.GetWebHooksOperation},
			{"update", management_api.UpdateWebHookOperation},
		},
		"billing": {
			{"create_agreement", management_api.CreateBillingAgreementOperation},
			{"get_agreements", management_api.GetBillingAgreementsOperation},
			{"get_entitlements", management_api.GetBillingEntitlementsOperation},
			{"get_entitlement", management_api.GetEntitlementOperation},
			{"create_meter_usage_record", management_api.CreateMeterUsageRecordOperation},
		},
		"events": {
			{"get", management_api.GetEventOperation},
			{"get_types", management_api.GetEventTypesOperation},
		},
		"portal": {
			{"get_link", management_api.GetPortalLinkOperation},
		},
		"industries": {
			{"get_all", management_api.GetIndustriesOperation},
		},
		"timezones": {
			{"get_all", management_api.GetTimezonesOperation},
		},
		"connected_apps": {
			{"get_auth_url", management_api.GetConnectedAppAuthUrlOperation},
			{"get_token", management_api.GetConnectedAppTokenOperation},
			{"revoke_token", management_api.RevokeConnectedAppTokenOperation},
		},
	}

	for group, operations := range ops {
		groupCmd := &cobra.Command{
			Use:   group,
			Args:  cobra.NoArgs,
			Short: fmt.Sprintf("Manage %s operations", group),
			Long:  fmt.Sprintf("Manage %s operations in Kinde", group),
		}

		for _, op := range operations {
			generatedcommand, err := buildCobraCommand(ctx, op)
			if err != nil {
				log.Error().Err(err).Msg("Failed to build command")
				return nil
			}
			groupCmd.AddCommand(generatedcommand)
		}

		manageCmd.cmd.AddCommand(groupCmd)
	}

	return manageCmd
}

func buildCobraCommand(ctx context.Context, op commandOperationPair[string, management_api.OperationName]) (*cobra.Command, error) {
	config := config.FromContext[config.Config](ctx)

	env := config.GetEnvironment()

	apiMethod, found := reflect.TypeOf(&management_api.Client{}).MethodByName(op.operation)
	if !found {
		return nil, fmt.Errorf("operation %s not found in management API", op.operation)
	}

	command := &cobra.Command{
		Use:   op.commandName,
		Args:  cobra.NoArgs,
		Short: fmt.Sprintf("Manage %s operation", op.commandName),
		Long:  fmt.Sprintf("Manage %s operation in Kinde", op.commandName),
		RunE: func(cmd *cobra.Command, args []string) error {
			if env == nil {
				return fmt.Errorf("no environment configured. Please run 'kinde login'")
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

func callApiMethod(cmd *cobra.Command, env *config.Environment, op commandOperationPair[string, management_api.OperationName]) error {
	log := log.Ctx(cmd.Context())
	ctx := cmd.Context()

	clientCredentials, err := env.NewClientCredentialsFlow()
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
		log.Info().Any("result", resultInstance).Msg("API call result")
	}

	return nil
}

func generateCommandFlags(t reflect.Type, command *cobra.Command) {

	visitor.NewVisitor(t).
		Visit(nil, func(p visitor.Walker[reflect.Type]) []visitor.Walker[reflect.Type] {
			visits := []visitor.Walker[reflect.Type]{}
			for i := range p.T1.NumField() {
				field := p.T1.Field(i)
				flagName := field.Tag.Get("json")

				if flagName == "" {
					flagName = toSnakeCase(field.Name)
				}
				if p.T2 != "" {
					flagName = fmt.Sprintf("%v.%v", p.T2, flagName)
				}

				switch field.Type {
				case reflect.TypeOf(management_api.OptNilBool{}), reflect.TypeOf(management_api.OptBool{}):
					command.Flags().Bool(flagName, false, "")
				case reflect.TypeOf(management_api.OptNilInt{}), reflect.TypeOf(management_api.OptInt{}):
					command.Flags().Int(flagName, 0, "")
				case reflect.TypeOf(management_api.OptString{}), reflect.TypeOf(management_api.OptNilString{}), reflect.TypeOf(""):
					command.Flags().String(flagName, "", "")
				default:
					if strings.HasPrefix(field.Type.String(), "management_api.Opt") {
						log.Debug().Msgf("%s - mapping type for field %s: %s", t.Name(), field.Name, field.Type.String())
						var i reflect.Value
						if p.OptionalSetter.Value != nil && p.OptionalSetter.Value.IsValid() {
							i = p.OptionalSetter.Value.FieldByName(field.Name)
						}
						visits = append(visits, visitor.Walker[reflect.Type]{T1: field.Type, T2: flagName, OptionalSetter: visitor.OptSetter{Value: &i}})
					} else {
						log.Debug().Msgf("%s - skipping unsupported type for field %s: %s", t.Name(), field.Name, field.Type.String())
					}
				}
			}
			return visits
		})

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
				flagName := field.Tag.Get("json")

				if flagName == "" {
					flagName = toSnakeCase(field.Name)
				}
				if p.T2 != "" {
					flagName = fmt.Sprintf("%v.%v", p.T2, flagName)
				}

				flag := flagSet.Lookup(flagName)
				if flag == nil {
					possibleName := toSnakeCase(field.Name)
					flag = flagSet.Lookup(possibleName)
				}

				log.Debug().Msgf("Mapping field %s with flag %s, val type %s", field.Name, flagName, val.Type().String())
				switch field.Type {
				case reflect.TypeOf(management_api.OptString{}):
					if flag != nil && flag.Changed {
						if flagValue, err := flagSet.GetString(flag.Name); err == nil {
							instanceField := p.OptionalSetter.Value.FieldByName(field.Name)
							instanceField.Set(reflect.ValueOf(management_api.NewOptString(flagValue)))
							p.OptionalSetter.IsSet.SetBool(true)
						}
					}
				default:
					if strings.HasPrefix(field.Type.String(), "management_api.Opt") {
						subProperty := p.OptionalSetter.Value.FieldByName(field.Name)
						log.Debug().Msgf("%s - mapping type for field %s: %s: %s", t.Name(), field.Name, field.Type.String(), subProperty.Type().String())
						additionalVisits = append(additionalVisits, visitor.Walker[reflect.Type]{T1: field.Type, T2: flagName, OptionalSetter: visitor.OptSetter{Value: &subProperty}})
					} else {
						log.Debug().Msgf("%s - skipping unsupported type for field %s: %s", t.Name(), field.Name, field.Type.String())
					}
				}
			}
			return additionalVisits
		})

	return instance
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
