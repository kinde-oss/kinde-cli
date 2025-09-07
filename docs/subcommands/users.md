# `kinde manage users`

The `kinde manage users` CLI provides various subcommands to manage user operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Users commands allow you to manage user accounts, identities, authentication, and user-specific settings in your Kinde environment. These operations are used for user lifecycle management, authentication configuration, and user data management.

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create](#create)
  - [create_identity](#create_identity)
  - [delete](#delete)
  - [delete_identity](#delete_identity)
  - [delete_sessions](#delete_sessions)
  - [get_all](#get_all)
  - [get_data](#get_data)
  - [get_identities](#get_identities)
  - [get_mfa](#get_mfa)
  - [get_property_values](#get_property_values)
  - [get_sessions](#get_sessions)
  - [refresh_claims](#refresh_claims)
  - [replace_mfa](#replace_mfa)
  - [reset_mfa](#reset_mfa)
  - [reset_mfa_all](#reset_mfa_all)
  - [search](#search)
  - [set_password](#set_password)
  - [update](#update)
  - [update_feature_flag_override](#update_feature_flag_override)
  - [update_identity](#update_identity)
  - [update_properties](#update_properties)
  - [update_property](#update_property)
- [Common Usage Patterns](#common-usage-patterns)
  - [User Lifecycle Management](#user-lifecycle-management)
  - [Identity Management](#identity-management)
  - [User Search and Filtering](#user-search-and-filtering)
  - [MFA Management](#mfa-management)
  - [Session Management](#session-management)
  - [Password Management](#password-management)
  - [Property Management](#property-management)
  - [Feature Flag Management](#feature-flag-management)
- [Best Practices](#best-practices)
- [Security Considerations](#security-considerations)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create`

Create a new user.

**Usage:**
```bash
kinde manage users create [flags]
```

**Flags:**
- `--organization_code string` - Organization code for the user
- `--profile.family_name string` - User's family name
- `--profile.given_name string` - User's given name
- `--profile.picture string` - User's profile picture URL
- `--provided_id string` - External user ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users create --profile.given_name "John" --profile.family_name "Doe" --organization_code "my-org"
```

### `create_identity`

Add a new identity to a user.

**Usage:**
```bash
kinde manage users create_identity [flags]
```

**Flags:**
- `--connection_id string` - Connection ID for the identity
- `--phone_country_id string` - Phone country ID
- `--user_id string` - User ID
- `--value string` - Identity value (email, phone, etc.)
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users create_identity --user_id "user_123" --connection_id "conn_456" --value "user@example.com"
```

### `delete`

Delete a user.

**Usage:**
```bash
kinde manage users delete [flags]
```

**Flags:**
- `--id string` - User ID to delete
- `--is_delete_profile` - Delete user profile as well
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users delete --id "user_123" --is_delete_profile
```

### `delete_identity`

Remove an identity from a user.

**Usage:**
```bash
kinde manage users delete_identity [flags]
```

**Flags:**
- `--identity_id string` - Identity ID to delete
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users delete_identity --identity_id "identity_123"
```

### `delete_sessions`

Delete all sessions for a user.

**Usage:**
```bash
kinde manage users delete_sessions [flags]
```

**Flags:**
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users delete_sessions --user_id "user_123"
```

### `get_all`

Retrieve all users with optional filtering and pagination.

**Usage:**
```bash
kinde manage users get_all [flags]
```

**Flags:**
- `--email string` - Filter by user email address
- `--expand string` - Specify related resources to expand in the response
- `--has_organization` - Filter to include only users with an organization
- `--next_token string` - Token for fetching the next page of results
- `--page_size int` - Number of results to return per page
- `--phone string` - Filter by user phone number
- `--user_id string` - Filter by user ID
- `--username string` - Filter by username
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Examples:**
```bash
# Get all users
kinde manage users get_all

# Get users with pagination
kinde manage users get_all --page_size 10

# Filter by email
kinde manage users get_all --email "user@example.com"

# Filter by organization
kinde manage users get_all --has_organization

# Get next page
kinde manage users get_all --next_token "token_123"
```

### `get_data`

Get detailed data for a user.

**Usage:**
```bash
kinde manage users get_data [flags]
```

**Flags:**
- `--expand string` - Specify related resources to expand in the response
- `--id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users get_data --id "user_123" --expand "identities,roles"
```

### `get_identities`

List all identities associated with a user.

**Usage:**
```bash
kinde manage users get_identities [flags]
```

**Flags:**
- `--ending_before string` - Token for pagination (end before this identity)
- `--starting_after string` - Token for pagination (start after this identity)
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users get_identities --user_id "user_123"
```

### `get_mfa`

Get multi-factor authentication settings for a user.

**Usage:**
```bash
kinde manage users get_mfa [flags]
```

**Flags:**
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users get_mfa --user_id "user_123"
```

### `get_property_values`

Get values of specific properties for a user.

**Usage:**
```bash
kinde manage users get_property_values [flags]
```

**Flags:**
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users get_property_values --user_id "user_123"
```

### `get_sessions`

List active sessions for a user.

**Usage:**
```bash
kinde manage users get_sessions [flags]
```

**Flags:**
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users get_sessions --user_id "user_123"
```

### `refresh_claims`

Refresh authentication claims for a user.

**Usage:**
```bash
kinde manage users refresh_claims [flags]
```

**Flags:**
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users refresh_claims --user_id "user_123"
```

### `replace_mfa`

Replace multi-factor authentication settings for a user.

**Usage:**
```bash
kinde manage users replace_mfa [flags]
```

**Flags:**
- `--policy string` - MFA policy
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users replace_mfa --policy "enforced"
```

### `reset_mfa`

Reset multi-factor authentication for a user.

**Usage:**
```bash
kinde manage users reset_mfa [flags]
```

**Flags:**
- `--factor_id string` - MFA factor ID
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users reset_mfa --user_id "user_123" --factor_id "factor_456"
```

### `reset_mfa_all`

Reset multi-factor authentication for all users.

**Usage:**
```bash
kinde manage users reset_mfa_all [flags]
```

**Flags:**
- `--user_id string` - User ID (optional, for specific user)
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users reset_mfa_all
```

### `search`

Search for users by criteria.

**Usage:**
```bash
kinde manage users search [flags]
```

**Flags:**
- `--ending_before string` - Token for pagination (end before this user)
- `--expand string` - Specify related resources to expand in the response
- `--page_size int` - Number of results to return per page
- `--query string` - Search query
- `--starting_after string` - Token for pagination (start after this user)
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users search --query "john doe" --page_size 10
```

### `set_password`

Set or update a user's password.

**Usage:**
```bash
kinde manage users set_password [flags]
```

**Flags:**
- `--hashed_password string` - Hashed password
- `--is_temporary_password` - Mark as temporary password
- `--salt string` - Password salt
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users set_password --user_id "user_123" --is_temporary_password
```

### `update`

Update user details.

**Usage:**
```bash
kinde manage users update [flags]
```

**Flags:**
- `--family_name string` - User's family name
- `--given_name string` - User's given name
- `--id string` - User ID
- `--is_password_reset_requested` - Request password reset
- `--is_suspended` - Suspend the user
- `--picture string` - User's profile picture URL
- `--provided_id string` - External user ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users update --id "user_123" --given_name "Jane" --family_name "Smith"
```

### `update_feature_flag_override`

Update feature flag overrides for a user.

**Usage:**
```bash
kinde manage users update_feature_flag_override [flags]
```

**Flags:**
- `--feature_flag_key string` - Feature flag key
- `--user_id string` - User ID
- `--value string` - Override value
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users update_feature_flag_override --user_id "user_123" --feature_flag_key "dark_mode" --value "true"
```

### `update_identity`

Update an identity for a user.

**Usage:**
```bash
kinde manage users update_identity [flags]
```

**Flags:**
- `--identity_id string` - Identity ID to update
- `--is_primary` - Mark as primary identity
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users update_identity --identity_id "identity_123" --is_primary
```

### `update_properties`

Update multiple properties for a user.

**Usage:**
```bash
kinde manage users update_properties [flags]
```

**Flags:**
- `--user_id string` - User ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users update_properties --user_id "user_123"
```

### `update_property`

Update a specific property for a user.

**Usage:**
```bash
kinde manage users update_property [flags]
```

**Flags:**
- `--property_key string` - Property key to update
- `--user_id string` - User ID
- `--value string` - New property value
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage users update_property --user_id "user_123" --property_key "department" --value "Engineering"
```

## Common Usage Patterns

### User Lifecycle Management

```bash
# Create a new user
kinde manage users create --profile.given_name "John" --profile.family_name "Doe" --organization_code "my-org"

# Get user details
kinde manage users get_data --id "user_123"

# Update user information
kinde manage users update --id "user_123" --given_name "Jane" --family_name "Smith"

# Suspend a user
kinde manage users update --id "user_123" --is_suspended

# Delete a user
kinde manage users delete --id "user_123" --is_delete_profile
```

### Identity Management

```bash
# Add an identity to a user
kinde manage users create_identity --user_id "user_123" --connection_id "conn_456" --value "user@example.com"

# List user identities
kinde manage users get_identities --user_id "user_123"

# Update an identity
kinde manage users update_identity --identity_id "identity_123" --is_primary

# Remove an identity
kinde manage users delete_identity --identity_id "identity_123"
```

### User Search and Filtering

```bash
# Search for users
kinde manage users search --query "john" --page_size 10

# Get all users with pagination
kinde manage users get_all --page_size 20

# Filter users by email
kinde manage users get_all --email "user@example.com"

# Filter users by organization
kinde manage users get_all --has_organization
```

### MFA Management

```bash
# Get user MFA settings
kinde manage users get_mfa --user_id "user_123"

# Reset MFA for a specific user
kinde manage users reset_mfa --user_id "user_123" --factor_id "factor_456"

# Reset MFA for all users
kinde manage users reset_mfa_all

# Replace MFA policy
kinde manage users replace_mfa --policy "enforced"
```

### Session Management

```bash
# Get user sessions
kinde manage users get_sessions --user_id "user_123"

# Delete all user sessions (force logout)
kinde manage users delete_sessions --user_id "user_123"

# Refresh user claims
kinde manage users refresh_claims --user_id "user_123"
```

### Password Management

```bash
# Set a temporary password
kinde manage users set_password --user_id "user_123" --is_temporary_password

# Request password reset
kinde manage users update --id "user_123" --is_password_reset_requested
```

### Property Management

```bash
# Get user properties
kinde manage users get_property_values --user_id "user_123"

# Update a specific property
kinde manage users update_property --user_id "user_123" --property_key "department" --value "Engineering"

# Update multiple properties
kinde manage users update_properties --user_id "user_123"
```

### Feature Flag Management

```bash
# Update user feature flag override
kinde manage users update_feature_flag_override --user_id "user_123" --feature_flag_key "beta_features" --value "true"
```

## Best Practices

1. **User Creation**: Always provide meaningful profile information when creating users
2. **Identity Management**: Use primary identities for main user accounts
3. **Security**: Regularly review and manage user sessions and MFA settings
4. **Data Privacy**: Be careful when deleting users and consider data retention policies
5. **Search Optimization**: Use appropriate filters and pagination for large user bases
6. **Property Management**: Use consistent property keys and values across your application

## Security Considerations

- **Password Security**: Use strong password policies and consider temporary passwords for new users
- **Session Management**: Regularly audit and clean up inactive sessions
- **MFA**: Encourage and properly configure multi-factor authentication
- **User Suspension**: Use suspension instead of deletion when users need temporary access removal
- **Identity Verification**: Verify identities before making them primary

## Troubleshooting

- **User Not Found**: Verify the user ID exists using `get_all` or `search`
- **Identity Issues**: Check that connection IDs and identity values are correct
- **Permission Errors**: Ensure your M2M application has user management scopes
- **Session Issues**: Use `delete_sessions` to force logout problematic sessions
- **MFA Problems**: Use `reset_mfa` to resolve MFA configuration issues

## Related Commands

- `kinde manage organizations` - Manage organizations that users belong to
- `kinde manage connections` - Manage authentication connections for user identities
- `kinde manage roles` - Manage roles that can be assigned to users
- `kinde manage permissions` - Manage permissions that users can have
- `kinde manage feature_flags` - Manage feature flags that users can override
