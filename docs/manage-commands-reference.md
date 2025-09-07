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

Manage API keys operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `delete` - Manage delete operation
- `get` - Manage get operation
- `get_all` - Manage get_all operation
- `rotate` - Manage rotate operation
- `verify` - Manage verify operation

**Usage:** `kinde manage api_keys [command] [options]`

### 2. APIs Management (`apis`)

Manage APIs operations in Kinde.

**Available Commands:**
- `add_apis` - Manage add_apis operation
- `add_logout_redirect_urls` - Manage add_logout_redirect_urls operation
- `add_redirect_callback_urls` - Manage add_redirect_callback_urls operation
- `add_scope` - Manage add_scope operation
- `delete_api` - Manage delete_api operation
- `delete_callback_urls` - Manage delete_callback_urls operation
- `delete_logout_urls` - Manage delete_logout_urls operation
- `delete_scope` - Manage delete_scope operation
- `get_api` - Manage get_api operation
- `get_apis` - Manage get_apis operation
- `get_callback_urls` - Manage get_callback_urls operation
- `get_logout_urls` - Manage get_logout_urls operation
- `get_scope` - Manage get_scope operation
- `get_scopes` - Manage get_scopes operation
- `replace_logout_redirect_urls` - Manage replace_logout_redirect_urls operation
- `replace_redirect_callback_urls` - Manage replace_redirect_callback_urls operation
- `update_applications` - Manage update_applications operation
- `update_scope` - Manage update_scope operation

**Usage:** `kinde manage apis [command] [options]`

### 3. Applications Management (`applications`)

Manage applications operations in Kinde.

**Available Commands:**
- `add_application_scope` - Manage add_application_scope operation
- `create` - Manage create operation
- `delete` - Manage delete operation
- `delete_application_scope` - Manage delete_application_scope operation
- `get` - Manage get operation
- `get_all` - Manage get_all operation
- `get_connections` - Manage get_connections operation
- `get_property_values` - Manage get_property_values operation
- `update` - Manage update operation
- `update_property` - Manage update_property operation
- `update_tokens` - Manage update_tokens operation

**Usage:** `kinde manage applications [command] [options]`

### 4. Billing Management (`billing`)

Manage billing operations in Kinde.

**Available Commands:**
- `create_agreement` - Manage create_agreement operation
- `create_meter_usage_record` - Manage create_meter_usage_record operation
- `get_agreements` - Manage get_agreements operation
- `get_entitlements` - Manage get_entitlements operation

**Usage:** `kinde manage billing [command] [options]`

### 5. Business Management (`business`)

Manage business operations in Kinde.

**Available Commands:**
- `get` - Manage get operation
- `update` - Manage update operation

**Usage:** `kinde manage business [command] [options]`

### 6. Categories Management (`categories`)

Manage categories operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `get_all` - Manage get_all operation
- `update` - Manage update operation

**Usage:** `kinde manage categories [command] [options]`

### 7. Connected Apps Management (`connected_apps`)

Manage connected applications operations in Kinde.

**Available Commands:**
- `get_auth_url` - Manage get_auth_url operation
- `get_token` - Manage get_token operation
- `revoke_token` - Manage revoke_token operation

**Usage:** `kinde manage connected_apps [command] [options]`

### 8. Connections Management (`connections`)

Manage connections operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `delete` - Manage delete operation
- `enable` - Manage enable operation
- `enable_org` - Manage enable_org operation
- `get` - Manage get operation
- `get_all` - Manage get_all operation
- `remove` - Manage remove operation
- `remove_org` - Manage remove_org operation
- `replace` - Manage replace operation
- `update` - Manage update operation

**Usage:** `kinde manage connections [command] [options]`

### 9. Environment Variables Management (`environment_variables`)

Manage environment variables operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `delete` - Manage delete operation
- `get` - Manage get operation
- `get_all` - Manage get_all operation
- `update` - Manage update operation

**Usage:** `kinde manage environment_variables [command] [options]`

### 10. Environments Management (`environments`)

Manage environments operations in Kinde.

**Available Commands:**
- `add_logo` - Manage add_logo operation
- `delete_feature_flag_override` - Manage delete_feature_flag_override operation
- `delete_feature_flag_overrides` - Manage delete_feature_flag_overrides operation
- `delete_logo` - Manage delete_logo operation
- `get` - Manage get operation
- `get_feature_flags` - Manage get_feature_flags operation
- `read_logo` - Manage read_logo operation
- `update_feature_flag_override` - Manage update_feature_flag_override operation

**Usage:** `kinde manage environments [command] [options]`

### 11. Events Management (`events`)

Manage events operations in Kinde.

**Available Commands:**
- `get` - Manage get operation
- `get_types` - Manage get_types operation

**Usage:** `kinde manage events [command] [options]`

### 12. Feature Flags Management (`feature_flags`)

Manage feature flags operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `delete` - Manage delete operation
- `update` - Manage update operation

**Usage:** `kinde manage feature_flags [command] [options]`

### 13. Industries Management (`industries`)

Manage industries operations in Kinde.

**Available Commands:**
- `get_all` - Manage get_all operation

**Usage:** `kinde manage industries [command] [options]`

### 14. Organizations Management (`organizations`)

Manage organizations operations in Kinde.

**Available Commands:**
- `add_logo` - Manage add_logo operation
- `add_user_api_scope` - Manage add_user_api_scope operation
- `add_users` - Manage add_users operation
- `create` - Manage create operation
- `create_user_permission` - Manage create_user_permission operation
- `create_user_role` - Manage create_user_role operation
- `delete` - Manage delete operation
- `delete_feature_flag_override` - Manage delete_feature_flag_override operation
- `delete_feature_flag_overrides` - Manage delete_feature_flag_overrides operation
- `delete_handle` - Manage delete_handle operation
- `delete_logo` - Manage delete_logo operation
- `delete_user_api_scope` - Manage delete_user_api_scope operation
- `delete_user_permission` - Manage delete_user_permission operation
- `delete_user_role` - Manage delete_user_role operation
- `get` - Manage get operation
- `get_all` - Manage get_all operation
- `get_connections` - Manage get_connections operation
- `get_feature_flags` - Manage get_feature_flags operation
- `get_property_values` - Manage get_property_values operation
- `get_user_mfa` - Manage get_user_mfa operation
- `get_user_permissions` - Manage get_user_permissions operation
- `get_user_roles` - Manage get_user_roles operation
- `get_users` - Manage get_users operation
- `read_logo` - Manage read_logo operation
- `remove_user` - Manage remove_user operation
- `replace_mfa` - Manage replace_mfa operation
- `reset_user_mfa` - Manage reset_user_mfa operation
- `reset_user_mfa_all` - Manage reset_user_mfa_all operation
- `update` - Manage update operation
- `update_feature_flag_override` - Manage update_feature_flag_override operation
- `update_properties` - Manage update_properties operation
- `update_property` - Manage update_property operation
- `update_sessions` - Manage update_sessions operation
- `update_users` - Manage update_users operation

**Usage:** `kinde manage organizations [command] [options]`

### 15. Permissions Management (`permissions`)

Manage permissions operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `delete` - Manage delete operation
- `get_all` - Manage get_all operation
- `update` - Manage update operation

**Usage:** `kinde manage permissions [command] [options]`

### 16. Properties Management (`properties`)

Manage properties operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `delete` - Manage delete operation
- `get_all` - Manage get_all operation
- `update` - Manage update operation

**Usage:** `kinde manage properties [command] [options]`

### 17. Roles Management (`roles`)

Manage roles operations in Kinde.

**Available Commands:**
- `add_scope` - Manage add_scope operation
- `create` - Manage create operation
- `delete` - Manage delete operation
- `delete_scope` - Manage delete_scope operation
- `get` - Manage get operation
- `get_all` - Manage get_all operation
- `get_permissions` - Manage get_permissions operation
- `get_scopes` - Manage get_scopes operation
- `remove_permission` - Manage remove_permission operation
- `update` - Manage update operation
- `update_permissions` - Manage update_permissions operation

**Usage:** `kinde manage roles [command] [options]`

### 18. Subscribers Management (`subscribers`)

Manage subscribers operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `get` - Manage get operation
- `get_all` - Manage get_all operation

**Usage:** `kinde manage subscribers [command] [options]`

### 19. Timezones Management (`timezones`)

Manage timezones operations in Kinde.

**Available Commands:**
- `get_all` - Manage get_all operation

**Usage:** `kinde manage timezones [command] [options]`

### 20. Users Management (`users`)

Manage users operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `create_identity` - Manage create_identity operation
- `delete` - Manage delete operation
- `delete_identity` - Manage delete_identity operation
- `delete_sessions` - Manage delete_sessions operation
- `get_all` - Manage get_all operation
- `get_data` - Manage get_data operation
- `get_identities` - Manage get_identities operation
- `get_mfa` - Manage get_mfa operation
- `get_property_values` - Manage get_property_values operation
- `get_sessions` - Manage get_sessions operation
- `refresh_claims` - Manage refresh_claims operation
- `replace_mfa` - Manage replace_mfa operation
- `reset_mfa` - Manage reset_mfa operation
- `reset_mfa_all` - Manage reset_mfa_all operation
- `search` - Manage search operation
- `set_password` - Manage set_password operation
- `update` - Manage update operation
- `update_feature_flag_override` - Manage update_feature_flag_override operation
- `update_identity` - Manage update_identity operation
- `update_properties` - Manage update_properties operation
- `update_property` - Manage update_property operation

**Usage:** `kinde manage users [command] [options]`

### 21. Webhooks Management (`webhooks`)

Manage webhooks operations in Kinde.

**Available Commands:**
- `create` - Manage create operation
- `delete` - Manage delete operation
- `get_all` - Manage get_all operation
- `update` - Manage update operation

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
