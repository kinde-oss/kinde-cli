# Kinde CLI Manage Commands Reference

This document provides a comprehensive reference for all `kinde manage` subcommands and their available operations. Each subcommand provides access to specific management API operations in Kinde.

## Overview

The `kinde manage` command provides access to the Kinde Management API, allowing you to perform various operations such as creating, updating, and deleting resources like applications, organizations, users, roles, and more.

**Usage:** `kinde manage <subcommand> [command] [options]`

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Available Subcommands

### 1. API Keys Management (`api_keys`)

Manage API keys for authentication and authorization in your Kinde environment. API keys provide secure access to your applications and can be scoped to specific APIs and permissions.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/api_keys.md#create) | Create a new API key with specified scopes and associations |
| [`delete`](subcommands/api_keys.md#delete) | Delete an API key permanently |
| [`get`](subcommands/api_keys.md#get) | Retrieve detailed information about a specific API key |
| [`get_all`](subcommands/api_keys.md#get_all) | List all API keys with filtering by type, status, user, or organization |
| [`rotate`](subcommands/api_keys.md#rotate) | Generate a new key value while maintaining permissions and associations |
| [`verify`](subcommands/api_keys.md#verify) | Verify an API key's validity and retrieve its details (public endpoint) |

**Usage:** `kinde manage api_keys [command] [options]`

### 2. APIs Management (`apis`)

Manage your backend APIs and their configurations. APIs represent the services that your applications can access, including scopes, callback URLs, and logout redirect URLs.

| Command | Description |
|---------|-------------|
| [`add_apis`](subcommands/apis.md#add_apis) | Create a new API with specified name and audience |
| [`add_logout_redirect_urls`](subcommands/apis.md#add_logout_redirect_urls) | Add logout redirect URLs to an application |
| [`add_redirect_callback_urls`](subcommands/apis.md#add_redirect_callback_urls) | Add redirect callback URLs to an application |
| [`add_scope`](subcommands/apis.md#add_scope) | Create a new scope for an API with key and description |
| [`delete_api`](subcommands/apis.md#delete_api) | Delete an API and all its associated configurations |
| [`delete_callback_urls`](subcommands/apis.md#delete_callback_urls) | Remove specific callback URLs from an application |
| [`delete_logout_urls`](subcommands/apis.md#delete_logout_urls) | Remove specific logout URLs from an application |
| [`delete_scope`](subcommands/apis.md#delete_scope) | Delete a scope from an API |
| [`get_api`](subcommands/apis.md#get_api) | Retrieve detailed information about a specific API |
| [`get_apis`](subcommands/apis.md#get_apis) | List all APIs in your environment, sorted by name |
| [`get_callback_urls`](subcommands/apis.md#get_callback_urls) | Retrieve all callback URLs configured for an application |
| [`get_logout_urls`](subcommands/apis.md#get_logout_urls) | Retrieve all logout URLs configured for an application |
| [`get_scope`](subcommands/apis.md#get_scope) | Retrieve detailed information about a specific scope |
| [`get_scopes`](subcommands/apis.md#get_scopes) | List all scopes associated with an API |
| [`replace_logout_redirect_urls`](subcommands/apis.md#replace_logout_redirect_urls) | Replace all logout redirect URLs for an application |
| [`replace_redirect_callback_urls`](subcommands/apis.md#replace_redirect_callback_urls) | Replace all redirect callback URLs for an application |
| [`update_applications`](subcommands/apis.md#update_applications) | Update which applications are associated with an API |
| [`update_scope`](subcommands/apis.md#update_scope) | Update the description or other properties of an existing scope |

**Usage:** `kinde manage apis [command] [options]`

### 3. Applications Management (`applications`)

Manage your client applications that integrate with Kinde for authentication and authorization. Applications can be web apps, SPAs, native apps, or machine-to-machine applications.

| Command | Description |
|---------|-------------|
| [`add_application_scope`](subcommands/applications.md#add_application_scope) | Associate a scope with an application for API access |
| [`create`](subcommands/applications.md#create) | Create a new application with specified type and organization |
| [`delete`](subcommands/applications.md#delete) | Delete an application and all its configurations |
| [`delete_application_scope`](subcommands/applications.md#delete_application_scope) | Remove a scope association from an application |
| [`get`](subcommands/applications.md#get) | Retrieve detailed information about a specific application |
| [`get_all`](subcommands/applications.md#get_all) | List all applications with pagination support |
| [`get_connections`](subcommands/applications.md#get_connections) | Retrieve authentication connections configured for an application |
| [`get_property_values`](subcommands/applications.md#get_property_values) | Retrieve custom property values for an application |
| [`update`](subcommands/applications.md#update) | Update application settings including name, URIs, and language |
| [`update_property`](subcommands/applications.md#update_property) | Update a specific custom property for an application |
| [`update_tokens`](subcommands/applications.md#update_tokens) | Configure token lifetimes and authentication settings |

**Usage:** `kinde manage applications [command] [options]`

### 4. Billing Management (`billing`)

Manage billing agreements, usage tracking, and entitlements for your Kinde environment. These operations handle subscription management and usage-based billing.

| Command | Description |
|---------|-------------|
| [`create_agreement`](subcommands/billing.md#create_agreement) | Create a new billing agreement with plan and customer details |
| [`create_meter_usage_record`](subcommands/billing.md#create_meter_usage_record) | Record usage for metered billing features |
| [`get_agreements`](subcommands/billing.md#get_agreements) | Retrieve all billing agreements and their details |
| [`get_entitlements`](subcommands/billing.md#get_entitlements) | Retrieve billing entitlements and feature access |

**Usage:** `kinde manage billing [command] [options]`

### 5. Business Management (`business`)

Manage business-level settings and configuration for your Kinde environment. These operations control organization-wide settings and preferences.

| Command | Description |
|---------|-------------|
| [`get`](subcommands/business.md#get) | Retrieve current business settings and configuration |
| [`update`](subcommands/business.md#update) | Update business settings and configuration |

**Usage:** `kinde manage business [command] [options]`

### 6. Categories Management (`categories`)

Manage property categories for organizing and classifying custom properties in your Kinde environment. Categories help structure and group related properties.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/categories.md#create) | Create a new property category |
| [`get_all`](subcommands/categories.md#get_all) | Retrieve all property categories |
| [`update`](subcommands/categories.md#update) | Update an existing property category |

**Usage:** `kinde manage categories [command] [options]`

### 7. Connected Apps Management (`connected_apps`)

Manage third-party application integrations and OAuth connections. These operations handle external service connections and token management for connected applications.

| Command | Description |
|---------|-------------|
| [`get_auth_url`](subcommands/connected_apps.md#get_auth_url) | Generate authorization URL for OAuth flow with connected app |
| [`get_token`](subcommands/connected_apps.md#get_token) | Exchange authorization code for access token |
| [`revoke_token`](subcommands/connected_apps.md#revoke_token) | Revoke access token and disconnect from external service |

**Usage:** `kinde manage connected_apps [command] [options]`

### 8. Connections Management (`connections`)

Manage authentication connections for social logins, SAML, and other identity providers. These operations configure how users can authenticate with your applications.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/connections.md#create) | Create a new authentication connection (OAuth, SAML, etc.) |
| [`delete`](subcommands/connections.md#delete) | Delete an authentication connection permanently |
| [`enable`](subcommands/connections.md#enable) | Enable a connection for use with an application |
| [`enable_org`](subcommands/connections.md#enable_org) | Enable a connection for use within an organization |
| [`get`](subcommands/connections.md#get) | Retrieve detailed information about a specific connection |
| [`get_all`](subcommands/connections.md#get_all) | List all connections with filtering and pagination |
| [`remove`](subcommands/connections.md#remove) | Disable a connection for an application |
| [`remove_org`](subcommands/connections.md#remove_org) | Disable a connection for an organization |
| [`replace`](subcommands/connections.md#replace) | Replace all connection settings and configuration |
| [`update`](subcommands/connections.md#update) | Update connection settings and configuration |

**Usage:** `kinde manage connections [command] [options]`

### 9. Environment Variables Management (`environment_variables`)

Manage configuration variables and secrets for your Kinde environment. These operations handle application configuration, API keys, and other environment-specific settings.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/environment_variables.md#create) | Create a new environment variable with optional encryption |
| [`delete`](subcommands/environment_variables.md#delete) | Delete an environment variable permanently |
| [`get`](subcommands/environment_variables.md#get) | Retrieve details of a specific environment variable |
| [`get_all`](subcommands/environment_variables.md#get_all) | List all environment variables in your environment |
| [`update`](subcommands/environment_variables.md#update) | Update an existing environment variable's value or settings |

**Usage:** `kinde manage environment_variables [command] [options]`

### 10. Environments Management (`environments`)

Manage environment-specific settings, feature flags, and branding. These operations control environment-level configurations and feature flag overrides.

| Command | Description |
|---------|-------------|
| [`add_logo`](subcommands/environments.md#add_logo) | Upload and set a logo for the environment |
| [`delete_feature_flag_override`](subcommands/environments.md#delete_feature_flag_override) | Remove a specific feature flag override |
| [`delete_feature_flag_overrides`](subcommands/environments.md#delete_feature_flag_overrides) | Remove all feature flag overrides for the environment |
| [`delete_logo`](subcommands/environments.md#delete_logo) | Remove the environment logo |
| [`get`](subcommands/environments.md#get) | Retrieve environment details and configuration |
| [`get_feature_flags`](subcommands/environments.md#get_feature_flags) | List all feature flags and their environment-level overrides |
| [`read_logo`](subcommands/environments.md#read_logo) | Retrieve the environment logo |
| [`update_feature_flag_override`](subcommands/environments.md#update_feature_flag_override) | Set or update a feature flag override for the environment |

**Usage:** `kinde manage environments [command] [options]`

### 11. Events Management (`events`)

Manage and retrieve system events, audit logs, and user activities. These operations provide access to event data for monitoring and compliance purposes.

| Command | Description |
|---------|-------------|
| [`get`](subcommands/events.md#get) | Retrieve detailed information about a specific event |
| [`get_types`](subcommands/events.md#get_types) | List all available event types in the system |

**Usage:** `kinde manage events [command] [options]`

### 12. Feature Flags Management (`feature_flags`)

Manage feature flags for controlling feature rollouts, A/B testing, and gradual feature releases. Feature flags enable dynamic feature control without code deployments.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/feature_flags.md#create) | Create a new feature flag with type, default value, and description |
| [`delete`](subcommands/feature_flags.md#delete) | Delete a feature flag and all its overrides |
| [`update`](subcommands/feature_flags.md#update) | Update feature flag settings including override levels and descriptions |

**Usage:** `kinde manage feature_flags [command] [options]`

### 13. Industries Management (`industries`)

Retrieve available industry classifications for categorizing your business or organization. These operations provide access to predefined industry types.

| Command | Description |
|---------|-------------|
| [`get_all`](subcommands/industries.md#get_all) | Retrieve all available industry classifications |

**Usage:** `kinde manage industries [command] [options]`

### 14. Organizations Management (`organizations`)

Manage organizations, their users, permissions, and settings. Organizations represent separate entities within your Kinde environment with their own users, roles, and configurations.

| Command | Description |
|---------|-------------|
| [`add_logo`](subcommands/organizations.md#add_logo) | Upload and set a logo for the organization |
| [`add_user_api_scope`](subcommands/organizations.md#add_user_api_scope) | Grant API scope access to a user within the organization |
| [`add_users`](subcommands/organizations.md#add_users) | Add users to the organization |
| [`create`](subcommands/organizations.md#create) | Create a new organization |
| [`create_user_permission`](subcommands/organizations.md#create_user_permission) | Create a custom permission for users in the organization |
| [`create_user_role`](subcommands/organizations.md#create_user_role) | Create a custom role for users in the organization |
| [`delete`](subcommands/organizations.md#delete) | Delete an organization and all its data |
| [`delete_feature_flag_override`](subcommands/organizations.md#delete_feature_flag_override) | Remove a specific feature flag override for the organization |
| [`delete_feature_flag_overrides`](subcommands/organizations.md#delete_feature_flag_overrides) | Remove all feature flag overrides for the organization |
| [`delete_handle`](subcommands/organizations.md#delete_handle) | Delete the organization handle |
| [`delete_logo`](subcommands/organizations.md#delete_logo) | Remove the organization logo |
| [`delete_user_api_scope`](subcommands/organizations.md#delete_user_api_scope) | Remove API scope access from a user |
| [`delete_user_permission`](subcommands/organizations.md#delete_user_permission) | Delete a custom permission from the organization |
| [`delete_user_role`](subcommands/organizations.md#delete_user_role) | Delete a custom role from the organization |
| [`get`](subcommands/organizations.md#get) | Retrieve detailed information about a specific organization |
| [`get_all`](subcommands/organizations.md#get_all) | List all organizations with pagination and filtering |
| [`get_connections`](subcommands/organizations.md#get_connections) | Retrieve authentication connections available to the organization |
| [`get_feature_flags`](subcommands/organizations.md#get_feature_flags) | List feature flags and their organization-level overrides |
| [`get_property_values`](subcommands/organizations.md#get_property_values) | Retrieve custom property values for the organization |
| [`get_user_mfa`](subcommands/organizations.md#get_user_mfa) | Get MFA settings for users in the organization |
| [`get_user_permissions`](subcommands/organizations.md#get_user_permissions) | List permissions available to users in the organization |
| [`get_user_roles`](subcommands/organizations.md#get_user_roles) | List roles available to users in the organization |
| [`get_users`](subcommands/organizations.md#get_users) | List all users in the organization |
| [`read_logo`](subcommands/organizations.md#read_logo) | Retrieve the organization logo |
| [`remove_user`](subcommands/organizations.md#remove_user) | Remove a user from the organization |
| [`replace_mfa`](subcommands/organizations.md#replace_mfa) | Replace MFA settings for the organization |
| [`reset_user_mfa`](subcommands/organizations.md#reset_user_mfa) | Reset MFA for a specific user in the organization |
| [`reset_user_mfa_all`](subcommands/organizations.md#reset_user_mfa_all) | Reset MFA for all users in the organization |
| [`update`](subcommands/organizations.md#update) | Update organization settings and configuration |
| [`update_feature_flag_override`](subcommands/organizations.md#update_feature_flag_override) | Set or update a feature flag override for the organization |
| [`update_properties`](subcommands/organizations.md#update_properties) | Update multiple custom properties for the organization |
| [`update_property`](subcommands/organizations.md#update_property) | Update a specific custom property for the organization |
| [`update_sessions`](subcommands/organizations.md#update_sessions) | Update session settings for the organization |
| [`update_users`](subcommands/organizations.md#update_users) | Update user settings within the organization |

**Usage:** `kinde manage organizations [command] [options]`

### 15. Permissions Management (`permissions`)

Manage permissions that define what actions users and applications can perform. Permissions are granular access controls that can be assigned to roles and users.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/permissions.md#create) | Create a new permission with name and description |
| [`delete`](subcommands/permissions.md#delete) | Delete a permission permanently |
| [`get_all`](subcommands/permissions.md#get_all) | List all permissions in your environment |
| [`update`](subcommands/permissions.md#update) | Update permission name or description |

**Usage:** `kinde manage permissions [command] [options]`

### 16. Properties Management (`properties`)

Manage custom properties for users, organizations, and applications. Properties allow you to store additional metadata and configuration data.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/properties.md#create) | Create a new custom property with type and validation rules |
| [`delete`](subcommands/properties.md#delete) | Delete a custom property permanently |
| [`get_all`](subcommands/properties.md#get_all) | List all custom properties in your environment |
| [`update`](subcommands/properties.md#update) | Update property settings including type and validation rules |

**Usage:** `kinde manage properties [command] [options]`

### 17. Roles Management (`roles`)

Manage roles that group permissions and can be assigned to users. Roles provide a way to organize and manage access control efficiently.

| Command | Description |
|---------|-------------|
| [`add_scope`](subcommands/roles.md#add_scope) | Add an API scope to a role |
| [`create`](subcommands/roles.md#create) | Create a new role with name and description |
| [`delete`](subcommands/roles.md#delete) | Delete a role permanently |
| [`delete_scope`](subcommands/roles.md#delete_scope) | Remove an API scope from a role |
| [`get`](subcommands/roles.md#get) | Retrieve detailed information about a specific role |
| [`get_all`](subcommands/roles.md#get_all) | List all roles in your environment |
| [`get_permissions`](subcommands/roles.md#get_permissions) | List all permissions assigned to a role |
| [`get_scopes`](subcommands/roles.md#get_scopes) | List all API scopes assigned to a role |
| [`remove_permission`](subcommands/roles.md#remove_permission) | Remove a permission from a role |
| [`update`](subcommands/roles.md#update) | Update role name or description |
| [`update_permissions`](subcommands/roles.md#update_permissions) | Update the permissions assigned to a role |

**Usage:** `kinde manage roles [command] [options]`

### 18. Subscribers Management (`subscribers`)

Manage subscribers for webhooks and notifications. Subscribers receive event notifications when specific actions occur in your Kinde environment.

| Command | Description |
|---------|-------------|
| `create` | Create a new subscriber for webhook notifications |
| `get` | Retrieve detailed information about a specific subscriber |
| `get_all` | List all subscribers in your environment |

**Usage:** `kinde manage subscribers [command] [options]`

### 19. Timezones Management (`timezones`)

Retrieve available timezone information for configuring user and organization settings. These operations provide access to timezone data for localization.

| Command | Description |
|---------|-------------|
| `get_all` | Retrieve all available timezones |

**Usage:** `kinde manage timezones [command] [options]`

### 20. Users Management (`users`)

Manage user accounts, identities, authentication, and user-specific settings. These operations handle the complete user lifecycle from creation to deletion, including authentication management.

| Command | Description |
|---------|-------------|
| [`create`](subcommands/users.md#create) | Create a new user with profile information and organization |
| [`create_identity`](subcommands/users.md#create_identity) | Add a new identity (email, phone, etc.) to a user |
| [`delete`](subcommands/users.md#delete) | Delete a user and optionally their profile data |
| [`delete_identity`](subcommands/users.md#delete_identity) | Remove an identity from a user |
| [`delete_sessions`](subcommands/users.md#delete_sessions) | Force logout by deleting all active sessions for a user |
| [`get_all`](subcommands/users.md#get_all) | List all users with filtering by email, phone, organization, etc. |
| [`get_data`](subcommands/users.md#get_data) | Retrieve comprehensive user data including expanded relationships |
| [`get_identities`](subcommands/users.md#get_identities) | List all identities associated with a user |
| [`get_mfa`](subcommands/users.md#get_mfa) | Retrieve multi-factor authentication settings for a user |
| [`get_property_values`](subcommands/users.md#get_property_values) | Get custom property values for a user |
| [`get_sessions`](subcommands/users.md#get_sessions) | List all active sessions for a user |
| [`refresh_claims`](subcommands/users.md#refresh_claims) | Refresh authentication claims and tokens for a user |
| [`replace_mfa`](subcommands/users.md#replace_mfa) | Replace MFA settings with new policy for a user |
| [`reset_mfa`](subcommands/users.md#reset_mfa) | Reset MFA for a specific user and factor |
| [`reset_mfa_all`](subcommands/users.md#reset_mfa_all) | Reset MFA for all users in the system |
| [`search`](subcommands/users.md#search) | Search for users using query strings with pagination |
| [`set_password`](subcommands/users.md#set_password) | Set or update a user's password with optional temporary flag |
| [`update`](subcommands/users.md#update) | Update user profile, suspension status, and other settings |
| [`update_feature_flag_override`](subcommands/users.md#update_feature_flag_override) | Set user-specific feature flag overrides |
| [`update_identity`](subcommands/users.md#update_identity) | Update identity settings including primary status |
| [`update_properties`](subcommands/users.md#update_properties) | Update multiple custom properties for a user |
| [`update_property`](subcommands/users.md#update_property) | Update a specific custom property for a user |

**Usage:** `kinde manage users [command] [options]`

### 21. Webhooks Management (`webhooks`)

Manage webhooks for receiving real-time notifications about events in your Kinde environment. Webhooks allow you to integrate with external systems and respond to user actions.

| Command | Description |
|---------|-------------|
| `create` | Create a new webhook endpoint for event notifications |
| `delete` | Delete a webhook endpoint permanently |
| `get_all` | List all webhook endpoints in your environment |
| `update` | Update webhook endpoint settings and configuration |

**Usage:** `kinde manage webhooks [command] [options]`

## Getting Detailed Help

For detailed help on any specific command, use:

```bash
kinde manage <subcommand> [command] --help
```

For example:
```bash
kinde manage users create --help
kinde manage applications get_all --help
kinde manage organizations get_users --help
```

## Authentication Requirements

To use the management API commands, you need to:

1. Create an M2M (Machine-to-Machine) application in your Kinde environment
2. Grant it access to the management API
3. Enable appropriate scopes to grant the specific level of access required
4. Configure the CLI with your API credentials

## Configuration

Configuration is stored in `$HOME/.kinde/config.json` by default for *nix-based systems, `<user profile>\.config\kinde\config.json` for Windows.

Example configuration:
```json
{
  "current": "mybusiness.kinde.com",
  "environments": {
    "mybusiness.kinde.com": {
      "domain_name": "app.kinde.com",
      "client_id": "xxx",
      "client_secret": "yyy"
    }
  }
}
```

## Examples

Here are some common usage examples:

```bash
# List all users
kinde manage users get_all --page_size 10

# Create a new user
kinde manage users create

# List all applications
kinde manage applications get_all

# Get organization details
kinde manage organizations get --organization_code "my-org"

# Create a new role
kinde manage roles create

# List all permissions
kinde manage permissions get_all

# Create a feature flag
kinde manage feature_flags create --name "new-feature" --key "new_feature_key" --type "bool"
```

## Notes

- All commands support the `--domain` flag to specify which Kinde environment to work against
- Management API scopes need to be granted to the application you are using to run these commands
- The CLI validates tokens against public JWKS and requires internet connectivity
- Check your configuration file for errors if commands fail
- Ensure your API credentials are valid and have the necessary permissions
