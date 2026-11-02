package cmd

import "github.com/kinde-oss/kinde-go/kinde/management_api"

type (
	// resourceCommand maps a CLI command name to a Management API operation.
	resourceCommand struct {
		name      string
		operation management_api.OperationName
	}

	// resourceGroup is a top-level command such as "users", holding the
	// commands for one area of the Management API.
	resourceGroup struct {
		name     string
		short    string
		commands []resourceCommand
	}
)

// resourceGroups lists every Management API operation the CLI exposes.
//
// Naming rules:
//   - Groups follow the categories in Kinde's API docs, in kebab-case.
//   - Commands use the verb from the API's own summary ("Replace Feature Flag"
//     becomes replace), except that getting a collection is "list".
//   - Actions on a sub-resource are "<verb>-<sub-resource>", e.g. create-identity.
var resourceGroups = []resourceGroup{
	{"api-keys", "Manage API keys", []resourceCommand{
		{"create", management_api.CreateApiKeyOperation},
		{"delete", management_api.DeleteApiKeyOperation},
		{"get", management_api.GetApiKeyOperation},
		{"list", management_api.GetApiKeysOperation},
		{"rotate", management_api.RotateApiKeyOperation},
		{"verify", management_api.VerifyApiKeyOperation},
	}},
	{"apis", "Manage APIs and their scopes", []resourceCommand{
		{"add-application-scope", management_api.AddAPIApplicationScopeOperation},
		{"authorize-applications", management_api.UpdateAPIApplicationsOperation},
		{"create", management_api.AddAPIsOperation},
		{"create-scope", management_api.AddAPIScopeOperation},
		{"delete", management_api.DeleteAPIOperation},
		{"delete-application-scope", management_api.DeleteAPIApplicationScopeOperation},
		{"delete-scope", management_api.DeleteAPIScopeOperation},
		{"get", management_api.GetAPIOperation},
		{"get-scope", management_api.GetAPIScopeOperation},
		{"list", management_api.GetAPIsOperation},
		{"list-scopes", management_api.GetAPIScopesOperation},
		{"update-scope", management_api.UpdateAPIScopeOperation},
	}},
	{"applications", "Manage applications, their connections and URLs", []resourceCommand{
		{"add-access-role", management_api.AddApplicationAccessRoleOperation},
		{"add-callback-urls", management_api.AddRedirectCallbackURLsOperation},
		{"add-logout-urls", management_api.AddLogoutRedirectURLsOperation},
		{"create", management_api.CreateApplicationOperation},
		{"delete", management_api.DeleteApplicationOperation},
		{"delete-callback-urls", management_api.DeleteCallbackURLsOperation},
		{"delete-logout-urls", management_api.DeleteLogoutURLsOperation},
		{"enable-connection", management_api.EnableConnectionOperation},
		{"get", management_api.GetApplicationOperation},
		{"list", management_api.GetApplicationsOperation},
		{"list-access-roles", management_api.GetApplicationAccessRolesOperation},
		{"list-callback-urls", management_api.GetCallbackURLsOperation},
		{"list-connections", management_api.GetApplicationConnectionsOperation},
		{"list-logout-urls", management_api.GetLogoutURLsOperation},
		{"list-properties", management_api.GetApplicationPropertyValuesOperation},
		{"remove-access-role", management_api.RemoveApplicationAccessRoleOperation},
		{"remove-connection", management_api.RemoveConnectionOperation},
		{"replace-callback-urls", management_api.ReplaceRedirectCallbackURLsOperation},
		{"replace-logout-urls", management_api.ReplaceLogoutRedirectURLsOperation},
		{"update", management_api.UpdateApplicationOperation},
		{"update-property", management_api.UpdateApplicationsPropertyOperation},
		{"update-tokens", management_api.UpdateApplicationTokensOperation},
	}},
	{"billing", "Manage billing agreements, entitlements and usage", []resourceCommand{
		{"create-agreement", management_api.CreateBillingAgreementOperation},
		{"create-meter-usage-record", management_api.CreateMeterUsageRecordOperation},
		{"list-agreements", management_api.GetBillingAgreementsOperation},
		{"list-entitlements", management_api.GetBillingEntitlementsOperation},
	}},
	{"business", "Manage business details", []resourceCommand{
		{"get", management_api.GetBusinessOperation},
		{"update", management_api.UpdateBusinessOperation},
	}},
	{"connected-apps", "Manage connected apps", []resourceCommand{
		{"get-auth-url", management_api.GetConnectedAppAuthUrlOperation},
		{"get-token", management_api.GetConnectedAppTokenOperation},
		{"revoke-token", management_api.RevokeConnectedAppTokenOperation},
	}},
	{"connections", "Manage authentication connections", []resourceCommand{
		{"create", management_api.CreateConnectionOperation},
		{"delete", management_api.DeleteConnectionOperation},
		{"get", management_api.GetConnectionOperation},
		{"list", management_api.GetConnectionsOperation},
		{"replace", management_api.ReplaceConnectionOperation},
		{"update", management_api.UpdateConnectionOperation},
	}},
	{"directories", "Manage SCIM directories", []resourceCommand{
		{"create", management_api.CreateDirectoryOperation},
		{"delete", management_api.DeleteDirectoryOperation},
		{"get", management_api.GetDirectoryOperation},
		{"list", management_api.GetDirectoriesOperation},
		{"update", management_api.UpdateDirectoryOperation},
	}},
	{"environment", "Manage the current environment", []resourceCommand{
		{"add-logo", management_api.AddLogoOperation},
		{"delete-feature-flag-override", management_api.DeleteEnvironementFeatureFlagOverrideOperation},
		{"delete-feature-flag-overrides", management_api.DeleteEnvironementFeatureFlagOverridesOperation},
		{"delete-logo", management_api.DeleteLogoOperation},
		{"get", management_api.GetEnvironmentOperation},
		{"get-passkey", management_api.GetPasskeyOperation},
		{"list-feature-flags", management_api.GetEnvironementFeatureFlagsOperation},
		{"list-logos", management_api.ReadLogoOperation},
		{"update-feature-flag-override", management_api.UpdateEnvironementFeatureFlagOverrideOperation},
		{"update-passkey", management_api.UpdatePasskeyOperation},
	}},
	{"environment-variables", "Manage environment variables", []resourceCommand{
		{"create", management_api.CreateEnvironmentVariableOperation},
		{"delete", management_api.DeleteEnvironmentVariableOperation},
		{"get", management_api.GetEnvironmentVariableOperation},
		{"list", management_api.GetEnvironmentVariablesOperation},
		{"update", management_api.UpdateEnvironmentVariableOperation},
	}},
	{"events", "Get webhook events and event types", []resourceCommand{
		{"get", management_api.GetEventOperation},
		{"list-types", management_api.GetEventTypesOperation},
	}},
	{"feature-flags", "Manage feature flags", []resourceCommand{
		{"create", management_api.CreateFeatureFlagOperation},
		{"delete", management_api.DeleteFeatureFlagOperation},
		{"replace", management_api.UpdateFeatureFlagOperation},
	}},
	{"identities", "Manage user identities", []resourceCommand{
		{"delete", management_api.DeleteIdentityOperation},
		{"get", management_api.GetIdentityOperation},
		{"update", management_api.UpdateIdentityOperation},
	}},
	{"industries", "List industries", []resourceCommand{
		{"list", management_api.GetIndustriesOperation},
	}},
	{"mfa", "Manage the environment's MFA configuration", []resourceCommand{
		{"replace", management_api.ReplaceMFAOperation},
	}},
	{"organizations", "Manage organizations, their users and settings", []resourceCommand{
		{"add-logo", management_api.AddOrganizationLogoOperation},
		{"add-user-api-scope", management_api.AddOrganizationUserAPIScopeOperation},
		{"add-user-permission", management_api.CreateOrganizationUserPermissionOperation},
		{"add-user-role", management_api.CreateOrganizationUserRoleOperation},
		{"add-users", management_api.AddOrganizationUsersOperation},
		{"count-role-active-users", management_api.GetOrganizationRoleActiveUsersCountOperation},
		{"count-role-users", management_api.GetOrganizationRoleUsersCountOperation},
		{"create", management_api.CreateOrganizationOperation},
		{"create-invite", management_api.CreateOrganizationInviteOperation},
		{"delete", management_api.DeleteOrganizationOperation},
		{"delete-feature-flag-override", management_api.DeleteOrganizationFeatureFlagOverrideOperation},
		{"delete-feature-flag-overrides", management_api.DeleteOrganizationFeatureFlagOverridesOperation},
		{"delete-handle", management_api.DeleteOrganizationHandleOperation},
		{"delete-invite", management_api.DeleteOrganizationInviteOperation},
		{"delete-logo", management_api.DeleteOrganizationLogoOperation},
		{"delete-user-api-scope", management_api.DeleteOrganizationUserAPIScopeOperation},
		{"delete-user-permission", management_api.DeleteOrganizationUserPermissionOperation},
		{"delete-user-role", management_api.DeleteOrganizationUserRoleOperation},
		{"enable-connection", management_api.EnableOrgConnectionOperation},
		{"get", management_api.GetOrganizationOperation},
		{"get-invite", management_api.GetOrganizationInviteOperation},
		{"get-passkey", management_api.GetOrganizationPasskeyOperation},
		{"get-user-mfa", management_api.GetOrgUserMFAOperation},
		{"list", management_api.GetOrganizationsOperation},
		{"list-connections", management_api.GetOrganizationConnectionsOperation},
		{"list-feature-flags", management_api.GetOrganizationFeatureFlagsOperation},
		{"list-invites", management_api.GetOrganizationInvitesOperation},
		{"list-logos", management_api.ReadOrganizationLogoOperation},
		{"list-properties", management_api.GetOrganizationPropertyValuesOperation},
		{"list-role-users", management_api.GetOrganizationRoleUsersOperation},
		{"list-user-permissions", management_api.GetOrganizationUserPermissionsOperation},
		{"list-user-roles", management_api.GetOrganizationUserRolesOperation},
		{"list-users", management_api.GetOrganizationUsersOperation},
		{"remove-connection", management_api.RemoveOrgConnectionOperation},
		{"remove-user", management_api.RemoveOrganizationUserOperation},
		{"replace-mfa", management_api.ReplaceOrganizationMFAOperation},
		{"reset-user-mfa", management_api.ResetOrgUserMFAOperation},
		{"reset-user-mfa-all", management_api.ResetOrgUserMFAAllOperation},
		{"update", management_api.UpdateOrganizationOperation},
		{"update-feature-flag-override", management_api.UpdateOrganizationFeatureFlagOverrideOperation},
		{"update-passkey", management_api.UpdateOrganizationPasskeyOperation},
		{"update-properties", management_api.UpdateOrganizationPropertiesOperation},
		{"update-property", management_api.UpdateOrganizationPropertyOperation},
		{"update-sessions", management_api.UpdateOrganizationSessionsOperation},
		{"update-users", management_api.UpdateOrganizationUsersOperation},
	}},
	{"permissions", "Manage permissions", []resourceCommand{
		{"create", management_api.CreatePermissionOperation},
		{"delete", management_api.DeletePermissionOperation},
		{"list", management_api.GetPermissionsOperation},
		{"update", management_api.UpdatePermissionsOperation},
	}},
	{"properties", "Manage custom properties", []resourceCommand{
		{"create", management_api.CreatePropertyOperation},
		{"delete", management_api.DeletePropertyOperation},
		{"list", management_api.GetPropertiesOperation},
		{"update", management_api.UpdatePropertyOperation},
	}},
	{"property-categories", "Manage property categories", []resourceCommand{
		{"create", management_api.CreateCategoryOperation},
		{"list", management_api.GetCategoriesOperation},
		{"update", management_api.UpdateCategoryOperation},
	}},
	{"roles", "Manage roles, their permissions and scopes", []resourceCommand{
		{"add-scope", management_api.AddRoleScopeOperation},
		{"create", management_api.CreateRoleOperation},
		{"delete", management_api.DeleteRoleOperation},
		{"delete-scope", management_api.DeleteRoleScopeOperation},
		{"get", management_api.GetRoleOperation},
		{"list", management_api.GetRolesOperation},
		{"list-permissions", management_api.GetRolePermissionsOperation},
		{"list-role-system-permissions", management_api.GetRoleSystemPermissionsOperation},
		{"list-scopes", management_api.GetRoleScopesOperation},
		{"list-system-permissions", management_api.GetSystemPermissionsOperation},
		{"list-users", management_api.GetRoleUsersOperation},
		{"remove-permission", management_api.RemoveRolePermissionOperation},
		{"update", management_api.UpdateRolesOperation},
		{"update-permissions", management_api.UpdateRolePermissionsOperation},
		{"update-system-permissions", management_api.UpdateRoleSystemPermissionsOperation},
	}},
	{"subscribers", "Manage subscribers", []resourceCommand{
		{"create", management_api.CreateSubscriberOperation},
		{"get", management_api.GetSubscriberOperation},
		{"list", management_api.GetSubscribersOperation},
	}},
	{"timezones", "List timezones", []resourceCommand{
		{"list", management_api.GetTimezonesOperation},
	}},
	{"users", "Manage users, their identities, sessions and MFA", []resourceCommand{
		{"create", management_api.CreateUserOperation},
		{"create-billing-customer", management_api.CreateUserBillingCustomerOperation},
		{"create-identity", management_api.CreateUserIdentityOperation},
		{"delete", management_api.DeleteUserOperation},
		{"delete-sessions", management_api.DeleteUserSessionsOperation},
		{"get", management_api.GetUserDataOperation},
		{"get-mfa", management_api.GetUsersMFAOperation},
		{"list", management_api.GetUsersOperation},
		{"list-identities", management_api.GetUserIdentitiesOperation},
		{"list-properties", management_api.GetUserPropertyValuesOperation},
		{"list-sessions", management_api.GetUserSessionsOperation},
		{"refresh-claims", management_api.RefreshUserClaimsOperation},
		{"reset-mfa", management_api.ResetUsersMFAOperation},
		{"reset-mfa-all", management_api.ResetUsersMFAAllOperation},
		{"search", management_api.SearchUsersOperation},
		{"set-password", management_api.SetUserPasswordOperation},
		{"update", management_api.UpdateUserOperation},
		{"update-feature-flag-override", management_api.UpdateUserFeatureFlagOverrideOperation},
		{"update-properties", management_api.UpdateUserPropertiesOperation},
		{"update-property", management_api.UpdateUserPropertyOperation},
	}},
	{"webhooks", "Manage webhooks", []resourceCommand{
		{"create", management_api.CreateWebHookOperation},
		{"delete", management_api.DeleteWebHookOperation},
		{"list", management_api.GetWebHooksOperation},
		{"update", management_api.UpdateWebHookOperation},
	}},
}
