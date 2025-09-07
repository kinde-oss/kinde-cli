# `kinde manage organizations`

The `kinde manage organizations` CLI provides various subcommands to manage organization operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Organizations in Kinde represent separate entities within your Kinde environment with their own users, roles, permissions, and configurations. This command group allows you to create, configure, and manage organizations, including their users, roles, permissions, feature flags, and branding settings.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [add_logo](#add_logo)
  - [add_user_api_scope](#add_user_api_scope)
  - [add_users](#add_users)
  - [create](#create)
  - [create_user_permission](#create_user_permission)
  - [create_user_role](#create_user_role)
  - [delete](#delete)
  - [delete_feature_flag_override](#delete_feature_flag_override)
  - [delete_feature_flag_overrides](#delete_feature_flag_overrides)
  - [delete_handle](#delete_handle)
  - [delete_logo](#delete_logo)
  - [delete_user_api_scope](#delete_user_api_scope)
  - [delete_user_permission](#delete_user_permission)
  - [delete_user_role](#delete_user_role)
  - [get](#get)
  - [get_all](#get_all)
  - [get_connections](#get_connections)
  - [get_feature_flags](#get_feature_flags)
  - [get_property_values](#get_property_values)
  - [get_user_mfa](#get_user_mfa)
  - [get_user_permissions](#get_user_permissions)
  - [get_user_roles](#get_user_roles)
  - [get_users](#get_users)
  - [read_logo](#read_logo)
  - [remove_user](#remove_user)
  - [replace_mfa](#replace_mfa)
  - [reset_user_mfa](#reset_user_mfa)
  - [reset_user_mfa_all](#reset_user_mfa_all)
  - [update](#update)
  - [update_feature_flag_override](#update_feature_flag_override)
  - [update_properties](#update_properties)
  - [update_property](#update_property)
  - [update_sessions](#update_sessions)
  - [update_users](#update_users)
- [Common Usage Patterns](#common-usage-patterns)
  - [Creating and Configuring a New Organization](#creating-and-configuring-a-new-organization)
  - [Managing Organization Users](#managing-organization-users)
  - [Managing Organization Roles and Permissions](#managing-organization-roles-and-permissions)
  - [Managing Organization Feature Flags](#managing-organization-feature-flags)
  - [Organization Branding and Customization](#organization-branding-and-customization)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `add_logo`

Upload and set a logo for the organization.

**Usage:**
```bash
kinde manage organizations add_logo [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations add_logo --org_code "my-org"
```

### `add_user_api_scope`

Grant API scope access to a user within the organization.

**Usage:**
```bash
kinde manage organizations add_user_api_scope [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--scope_id string` - API scope ID to grant
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations add_user_api_scope --org_code "my-org" --user_id "user_123" --scope_id "scope_456"
```

### `add_users`

Add users to the organization.

**Usage:**
```bash
kinde manage organizations add_users [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations add_users --org_code "my-org"
```

### `create`

Create a new organization.

**Usage:**
```bash
kinde manage organizations create [flags]
```

**Flags:**
- `--background_color string` - Background color for the organization theme
- `--background_color_dark string` - Dark mode background color
- `--billing_email string` - Billing email address
- `--billing_plan_code string` - Billing plan code
- `--button_color string` - Button color for the organization theme
- `--button_color_dark string` - Dark mode button color
- `--button_text_color string` - Button text color
- `--button_text_color_dark string` - Dark mode button text color
- `--external_id string` - External ID for the organization
- `--handle string` - Organization handle/identifier
- `--is_allow_registrations` - Allow user registrations
- `--is_create_billing_customer` - Create billing customer
- `--link_color string` - Link color for the organization theme
- `--link_color_dark string` - Dark mode link color
- `--name string` - Organization name
- `--sender_email string` - Sender email address
- `--sender_name string` - Sender name
- `--theme_code string` - Theme code

**Example:**
```bash
kinde manage organizations create --name "My Organization" --handle "my-org" --is_allow_registrations
```

### `create_user_permission`

Create a custom permission for users in the organization.

**Usage:**
```bash
kinde manage organizations create_user_permission [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--permission_id string` - Permission ID
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations create_user_permission --org_code "my-org" --user_id "user_123" --permission_id "perm_456"
```

### `create_user_role`

Create a custom role for users in the organization.

**Usage:**
```bash
kinde manage organizations create_user_role [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--role_id string` - Role ID
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations create_user_role --org_code "my-org" --user_id "user_123" --role_id "role_456"
```

### `delete`

Delete an organization and all its data.

**Usage:**
```bash
kinde manage organizations delete [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations delete --org_code "my-org"
```

**Warning:** This action is irreversible and will delete all organization data.

### `delete_feature_flag_override`

Remove a specific feature flag override for the organization.

**Usage:**
```bash
kinde manage organizations delete_feature_flag_override [flags]
```

**Flags:**
- `--feature_flag_key string` - Feature flag key
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations delete_feature_flag_override --org_code "my-org" --feature_flag_key "beta_features"
```

### `delete_feature_flag_overrides`

Remove all feature flag overrides for the organization.

**Usage:**
```bash
kinde manage organizations delete_feature_flag_overrides [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations delete_feature_flag_overrides --org_code "my-org"
```

### `delete_handle`

Delete the organization handle.

**Usage:**
```bash
kinde manage organizations delete_handle [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations delete_handle --org_code "my-org"
```

### `delete_logo`

Remove the organization logo.

**Usage:**
```bash
kinde manage organizations delete_logo [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations delete_logo --org_code "my-org"
```

### `delete_user_api_scope`

Remove API scope access from a user.

**Usage:**
```bash
kinde manage organizations delete_user_api_scope [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--scope_id string` - API scope ID to remove
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations delete_user_api_scope --org_code "my-org" --user_id "user_123" --scope_id "scope_456"
```

### `delete_user_permission`

Delete a custom permission from the organization.

**Usage:**
```bash
kinde manage organizations delete_user_permission [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--permission_id string` - Permission ID
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations delete_user_permission --org_code "my-org" --user_id "user_123" --permission_id "perm_456"
```

### `delete_user_role`

Delete a custom role from the organization.

**Usage:**
```bash
kinde manage organizations delete_user_role [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--role_id string` - Role ID
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations delete_user_role --org_code "my-org" --user_id "user_123" --role_id "role_456"
```

### `get`

Retrieve detailed information about a specific organization.

**Usage:**
```bash
kinde manage organizations get [flags]
```

**Flags:**
- `--code string` - Organization code
- `--expand string` - Specify related resources to expand in the response

**Example:**
```bash
kinde manage organizations get --code "my-org" --expand "users,roles"
```

### `get_all`

List all organizations with pagination and filtering.

**Usage:**
```bash
kinde manage organizations get_all [flags]
```

**Flags:**
- `--next_token string` - Token for fetching the next page of results
- `--page_size int` - Number of results to return per page

**Examples:**
```bash
# Get all organizations
kinde manage organizations get_all

# Get organizations with pagination
kinde manage organizations get_all --page_size 10

# Get next page of results
kinde manage organizations get_all --next_token "token_123"
```

### `get_connections`

Retrieve authentication connections available to the organization.

**Usage:**
```bash
kinde manage organizations get_connections [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations get_connections --org_code "my-org"
```

### `get_feature_flags`

List feature flags and their organization-level overrides.

**Usage:**
```bash
kinde manage organizations get_feature_flags [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations get_feature_flags --org_code "my-org"
```

### `get_property_values`

Retrieve custom property values for the organization.

**Usage:**
```bash
kinde manage organizations get_property_values [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations get_property_values --org_code "my-org"
```

### `get_user_mfa`

Get MFA settings for users in the organization.

**Usage:**
```bash
kinde manage organizations get_user_mfa [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations get_user_mfa --org_code "my-org"
```

### `get_user_permissions`

List permissions available to users in the organization.

**Usage:**
```bash
kinde manage organizations get_user_permissions [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations get_user_permissions --org_code "my-org"
```

### `get_user_roles`

List roles available to users in the organization.

**Usage:**
```bash
kinde manage organizations get_user_roles [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations get_user_roles --org_code "my-org"
```

### `get_users`

List all users in the organization.

**Usage:**
```bash
kinde manage organizations get_users [flags]
```

**Flags:**
- `--next_token string` - Token for fetching the next page of results
- `--org_code string` - Organization code
- `--page_size int` - Number of results to return per page
- `--permissions string` - Filter by permissions
- `--roles string` - Filter by roles

**Examples:**
```bash
# Get all users in organization
kinde manage organizations get_users --org_code "my-org"

# Get users with pagination
kinde manage organizations get_users --org_code "my-org" --page_size 10

# Filter users by role
kinde manage organizations get_users --org_code "my-org" --roles "admin"
```

### `read_logo`

Retrieve the organization logo.

**Usage:**
```bash
kinde manage organizations read_logo [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations read_logo --org_code "my-org"
```

### `remove_user`

Remove a user from the organization.

**Usage:**
```bash
kinde manage organizations remove_user [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations remove_user --org_code "my-org" --user_id "user_123"
```

### `replace_mfa`

Replace MFA settings for the organization.

**Usage:**
```bash
kinde manage organizations replace_mfa [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--policy string` - MFA policy

**Example:**
```bash
kinde manage organizations replace_mfa --org_code "my-org" --policy "enforced"
```

### `reset_user_mfa`

Reset MFA for a specific user in the organization.

**Usage:**
```bash
kinde manage organizations reset_user_mfa [flags]
```

**Flags:**
- `--factor_id string` - MFA factor ID
- `--org_code string` - Organization code
- `--user_id string` - User ID

**Example:**
```bash
kinde manage organizations reset_user_mfa --org_code "my-org" --user_id "user_123" --factor_id "factor_456"
```

### `reset_user_mfa_all`

Reset MFA for all users in the organization.

**Usage:**
```bash
kinde manage organizations reset_user_mfa_all [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations reset_user_mfa_all --org_code "my-org"
```

### `update`

Update organization settings and configuration.

**Usage:**
```bash
kinde manage organizations update [flags]
```

**Flags:**
- `--background_color string` - Background color for the organization theme
- `--background_color_dark string` - Dark mode background color
- `--button_color string` - Button color for the organization theme
- `--button_color_dark string` - Dark mode button color
- `--button_text_color string` - Button text color
- `--button_text_color_dark string` - Dark mode button text color
- `--external_id string` - External ID for the organization
- `--handle string` - Organization handle/identifier
- `--is_allow_registrations` - Allow user registrations
- `--is_auto_join_domain_list` - Auto-join domain list
- `--is_enable_advanced_orgs` - Enable advanced organizations
- `--is_enforce_mfa` - Enforce MFA
- `--link_color string` - Link color for the organization theme
- `--link_color_dark string` - Dark mode link color
- `--name string` - Organization name
- `--org_code string` - Organization code
- `--sender_email string` - Sender email address
- `--sender_name string` - Sender name

**Example:**
```bash
kinde manage organizations update --org_code "my-org" --name "Updated Organization Name" --is_enforce_mfa
```

### `update_feature_flag_override`

Set or update a feature flag override for the organization.

**Usage:**
```bash
kinde manage organizations update_feature_flag_override [flags]
```

**Flags:**
- `--feature_flag_key string` - Feature flag key
- `--org_code string` - Organization code
- `--value string` - Override value

**Example:**
```bash
kinde manage organizations update_feature_flag_override --org_code "my-org" --feature_flag_key "beta_features" --value "true"
```

### `update_properties`

Update multiple custom properties for the organization.

**Usage:**
```bash
kinde manage organizations update_properties [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations update_properties --org_code "my-org"
```

### `update_property`

Update a specific custom property for the organization.

**Usage:**
```bash
kinde manage organizations update_property [flags]
```

**Flags:**
- `--org_code string` - Organization code
- `--property_key string` - Property key to update

**Example:**
```bash
kinde manage organizations update_property --org_code "my-org" --property_key "custom_setting"
```

### `update_sessions`

Update session settings for the organization.

**Usage:**
```bash
kinde manage organizations update_sessions [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations update_sessions --org_code "my-org"
```

### `update_users`

Update user settings for the organization.

**Usage:**
```bash
kinde manage organizations update_users [flags]
```

**Flags:**
- `--org_code string` - Organization code

**Example:**
```bash
kinde manage organizations update_users --org_code "my-org"
```

## Common Usage Patterns

### Creating and Configuring a New Organization

```bash
# Create a new organization with basic settings
kinde manage organizations create --name "My Company" --handle "my-company" --is_allow_registrations

# Create organization with custom branding
kinde manage organizations create --name "My Company" --handle "my-company" \
  --background_color "#ffffff" --button_color "#007bff" --button_text_color "#ffffff"

# Create organization with billing
kinde manage organizations create --name "My Company" --handle "my-company" \
  --billing_email "billing@mycompany.com" --billing_plan_code "premium" --is_create_billing_customer
```

### Managing Organization Users

```bash
# List all users in an organization
kinde manage organizations get_users --org_code "my-org"

# Add users to organization
kinde manage organizations add_users --org_code "my-org"

# Remove a user from organization
kinde manage organizations remove_user --org_code "my-org" --user_id "user_123"

# Filter users by role
kinde manage organizations get_users --org_code "my-org" --roles "admin"
```

### Managing Organization Roles and Permissions

```bash
# Get available roles and permissions
kinde manage organizations get_user_roles --org_code "my-org"
kinde manage organizations get_user_permissions --org_code "my-org"

# Assign role to user
kinde manage organizations create_user_role --org_code "my-org" --user_id "user_123" --role_id "role_456"

# Grant permission to user
kinde manage organizations create_user_permission --org_code "my-org" --user_id "user_123" --permission_id "perm_789"

# Remove role from user
kinde manage organizations delete_user_role --org_code "my-org" --user_id "user_123" --role_id "role_456"
```

### Managing Organization Feature Flags

```bash
# Get organization feature flags
kinde manage organizations get_feature_flags --org_code "my-org"

# Set feature flag override
kinde manage organizations update_feature_flag_override --org_code "my-org" --feature_flag_key "beta_features" --value "true"

# Remove specific feature flag override
kinde manage organizations delete_feature_flag_override --org_code "my-org" --feature_flag_key "beta_features"

# Remove all feature flag overrides
kinde manage organizations delete_feature_flag_overrides --org_code "my-org"
```

### Organization Branding and Customization

```bash
# Update organization branding
kinde manage organizations update --org_code "my-org" \
  --background_color "#f8f9fa" --button_color "#28a745" --button_text_color "#ffffff"

# Add organization logo
kinde manage organizations add_logo --org_code "my-org"

# Read organization logo
kinde manage organizations read_logo --org_code "my-org"

# Delete organization logo
kinde manage organizations delete_logo --org_code "my-org"
```

### MFA Management

```bash
# Get MFA settings
kinde manage organizations get_user_mfa --org_code "my-org"

# Replace MFA policy
kinde manage organizations replace_mfa --org_code "my-org" --policy "enforced"

# Reset MFA for specific user
kinde manage organizations reset_user_mfa --org_code "my-org" --user_id "user_123" --factor_id "factor_456"

# Reset MFA for all users
kinde manage organizations reset_user_mfa_all --org_code "my-org"
```

## Best Practices

1. **Organization Structure**: Plan your organization hierarchy before creating multiple organizations
2. **User Management**: Use roles and permissions to implement proper access control
3. **Feature Flags**: Use organization-level feature flags for gradual rollouts
4. **Branding**: Maintain consistent branding across all organization touchpoints
5. **Security**: Enable MFA and enforce security policies at the organization level
6. **Billing**: Set up proper billing configuration for paid organizations
7. **Documentation**: Keep track of organization codes and configurations

## Troubleshooting

- **Organization Not Found**: Verify the organization code exists using `get_all`
- **Permission Errors**: Ensure your M2M application has organization management scopes
- **User Issues**: Check that user IDs exist and are valid
- **Feature Flag Problems**: Verify feature flag keys exist using the feature flags management commands
- **Billing Issues**: Ensure billing configuration is properly set up

## Related Commands

- `kinde manage users` - Manage users that belong to organizations
- `kinde manage applications` - Manage applications that can be associated with organizations
- `kinde manage connections` - Manage authentication connections for organizations
- `kinde manage feature_flags` - Manage feature flags that can be overridden at organization level
- `kinde manage roles` - Manage roles that can be assigned to organization users
- `kinde manage permissions` - Manage permissions that can be granted to organization users
