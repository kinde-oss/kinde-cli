package cmd

import (
	"context"
	"fmt"
	"reflect"

	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-go/kinde"
	"github.com/kinde-oss/kinde-go/kinde/management_api"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type manageCmd struct {
	cmd *cobra.Command
}

type commandOperationPair[string, T any] struct {
	commandName string
	operation   T
}

func newManageCmd(ctx context.Context) *manageCmd {

	manageCmd := &manageCmd{}

	manageCmd.cmd = &cobra.Command{
		Use:   "manage",
		Args:  cobra.NoArgs,
		Short: "Manage Kinde business",
	}

	ops := map[string][]commandOperationPair[string, management_api.OperationName]{
		"api": {
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
		"application": {
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
		"organization": {
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
		"user": {
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
		"role": {
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
		"permission": {
			{"create", management_api.CreatePermissionOperation},
			{"delete", management_api.DeletePermissionOperation},
			{"get_all", management_api.GetPermissionsOperation},
			{"update", management_api.UpdatePermissionsOperation},
		},
		"feature_flag": {
			{"create", management_api.CreateFeatureFlagOperation},
			{"delete", management_api.DeleteFeatureFlagOperation},
			{"get_all", management_api.GetFeatureFlagsOperation},
			{"update", management_api.UpdateFeatureFlagOperation},
		},
		"property": {
			{"create", management_api.CreatePropertyOperation},
			{"delete", management_api.DeletePropertyOperation},
			{"get_all", management_api.GetPropertiesOperation},
			{"update", management_api.UpdatePropertyOperation},
		},
		"connection": {
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
		"environment": {
			{"get", management_api.GetEnvironmentOperation},
			{"get_feature_flags", management_api.GetEnvironementFeatureFlagsOperation},
			{"update_feature_flag_override", management_api.UpdateEnvironementFeatureFlagOverrideOperation},
			{"delete_feature_flag_override", management_api.DeleteEnvironementFeatureFlagOverrideOperation},
			{"delete_feature_flag_overrides", management_api.DeleteEnvironementFeatureFlagOverridesOperation},
		},
		"environment_variable": {
			{"create", management_api.CreateEnvironmentVariableOperation},
			{"delete", management_api.DeleteEnvironmentVariableOperation},
			{"get", management_api.GetEnvironmentVariableOperation},
			{"get_all", management_api.GetEnvironmentVariablesOperation},
			{"update", management_api.UpdateEnvironmentVariableOperation},
		},
		"category": {
			{"create", management_api.CreateCategoryOperation},
			{"get_all", management_api.GetCategoriesOperation},
			{"update", management_api.UpdateCategoryOperation},
		},
		"subscriber": {
			{"create", management_api.CreateSubscriberOperation},
			{"get", management_api.GetSubscriberOperation},
			{"get_all", management_api.GetSubscribersOperation},
		},
		"webhook": {
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
		"event": {
			{"get", management_api.GetEventOperation},
			{"get_types", management_api.GetEventTypesOperation},
		},
		"portal": {
			{"get_link", management_api.GetPortalLinkOperation},
		},
		"industry": {
			{"get_all", management_api.GetIndustriesOperation},
		},
		"timezone": {
			{"get_all", management_api.GetTimezonesOperation},
		},
		"connected_app": {
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
	//log := log.Ctx(ctx)
	config := config.FromContext[config.Config](ctx)

	env := config.GetEnvironment()

	if env.DomainName == "" {
		return nil, fmt.Errorf("no environment configured. Please run 'kinde login'")
	}

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
			return callApiMethod(cmd, env, op)
		},
	}

	numArgs := apiMethod.Type.NumIn()
	var methodArgs []reflect.Value
	for i := range numArgs {
		if apiMethod.Type.In(i).String() != "context.Context" {
			argumentType := reflect.Zero(apiMethod.Type.In(i))
			methodArgs = append(methodArgs, argumentType)
			generateCommandFlags(argumentType.Type(), command)
		}
	}

	return command, nil
}

func callApiMethod(cmd *cobra.Command, env *config.Environment, op commandOperationPair[string, management_api.OperationName]) error {

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
		argInstance := mapFlagsToStruct(methodType.In(i), cmd.Flags())
		args = append(args, reflect.ValueOf(argInstance))
	}

	results := instanceMethod.Call(args[0:inParams])
	for _, result := range results {
		if result.IsNil() {
			continue
		}
		resultInstance := result.Interface()
		log.Info().Any("result", resultInstance).Msg("API call result")
	}

	return nil
}

func generateCommandFlags(t reflect.Type, command *cobra.Command) {
	// Ensure t is a struct type (dereference pointer if needed)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		switch field.Type {
		case reflect.TypeOf(management_api.OptNilBool{}), reflect.TypeOf(management_api.OptBool{}):
			if tag != "" {
				command.Flags().Bool(tag, false, "")
			}
		default:
			if tag != "" {
				command.Flags().String(tag, "", "")
			}
		}
	}
}

func mapFlagsToStruct(t reflect.Type, flagSet *pflag.FlagSet) any {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	instance := reflect.New(t).Interface()
	val := reflect.Indirect(reflect.ValueOf(instance))
	if t.Kind() != reflect.Struct {
		return instance
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		switch field.Type {
		case reflect.TypeOf(management_api.OptNilString{}):
			flag := flagSet.Lookup(tag)
			if flag != nil && flag.Changed {
				val.Field(i).Set(reflect.ValueOf(management_api.NewOptNilString(flag.Value.String())))
			}
		case reflect.TypeOf(management_api.OptNilBool{}):
			flag := flagSet.Lookup(tag)
			if flag != nil && flag.Changed {
				if flagValue, err := flagSet.GetBool(tag); err == nil {
					val.Field(i).Set(reflect.ValueOf(management_api.NewOptNilBool(flagValue)))
				}
			}
		default:
			if value, err := flagSet.GetString(tag); err == nil {
				val.Field(i).Set(reflect.ValueOf(management_api.NewOptNilString(value)))
			}
		}
	}
	return instance
}
