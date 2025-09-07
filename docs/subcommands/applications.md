# `kinde manage applications`

The `kinde manage applications` CLI provides various subcommands to manage application operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Applications in Kinde represent your client applications that integrate with Kinde for authentication and authorization. This command group allows you to create, configure, and manage applications, including their scopes, properties, and token settings.

## Available Commands

### `add_application_scope`

Add a scope to an application.

**Usage:**
```bash
kinde manage applications add_application_scope [flags]
```

**Flags:**
- `--apiid string` - API ID
- `--application_id string` - Application ID
- `--scope_id string` - Scope ID to add
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications add_application_scope --application_id "app_123" --apiid "api_456" --scope_id "scope_789"
```

### `create`

Create a new application.

**Usage:**
```bash
kinde manage applications create [flags]
```

**Flags:**
- `--name string` - Name of the application
- `--org_code string` - Organization code
- `--type string` - Application type
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications create --name "My Application" --type "web" --org_code "my-org"
```

### `delete`

Delete an application.

**Usage:**
```bash
kinde manage applications delete [flags]
```

**Flags:**
- `--application_id string` - Application ID to delete
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications delete --application_id "app_123"
```

### `delete_application_scope`

Remove a scope from an application.

**Usage:**
```bash
kinde manage applications delete_application_scope [flags]
```

**Flags:**
- `--apiid string` - API ID
- `--application_id string` - Application ID
- `--scope_id string` - Scope ID to remove
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications delete_application_scope --application_id "app_123" --apiid "api_456" --scope_id "scope_789"
```

### `get`

Retrieve details of a specific application.

**Usage:**
```bash
kinde manage applications get [flags]
```

**Flags:**
- `--application_id string` - Application ID to retrieve
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications get --application_id "app_123"
```

### `get_all`

Retrieve all applications with optional pagination.

**Usage:**
```bash
kinde manage applications get_all [flags]
```

**Flags:**
- `--next_token string` - Token for pagination
- `--page_size int` - Number of results to return per page
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Examples:**
```bash
# Get all applications
kinde manage applications get_all

# Get applications with pagination
kinde manage applications get_all --page_size 10

# Get next page of results
kinde manage applications get_all --next_token "token_123"
```

### `get_connections`

Retrieve connections for an application.

**Usage:**
```bash
kinde manage applications get_connections [flags]
```

**Flags:**
- `--application_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications get_connections --application_id "app_123"
```

### `get_property_values`

Retrieve property values for an application.

**Usage:**
```bash
kinde manage applications get_property_values [flags]
```

**Flags:**
- `--application_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications get_property_values --application_id "app_123"
```

### `update`

Update an application's basic settings.

**Usage:**
```bash
kinde manage applications update [flags]
```

**Flags:**
- `--application_id string` - Application ID to update
- `--homepage_uri string` - Homepage URI
- `--language_key string` - Language key
- `--login_uri string` - Login URI
- `--name string` - Application name
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications update --application_id "app_123" --name "Updated App Name" --homepage_uri "https://myapp.com"
```

### `update_property`

Update a specific property for an application.

**Usage:**
```bash
kinde manage applications update_property [flags]
```

**Flags:**
- `--application_id string` - Application ID
- `--property_key string` - Property key to update
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications update_property --application_id "app_123" --property_key "custom_setting"
```

### `update_tokens`

Update token settings for an application.

**Usage:**
```bash
kinde manage applications update_tokens [flags]
```

**Flags:**
- `--access_token_lifetime int` - Access token lifetime in seconds
- `--application_id string` - Application ID
- `--authenticated_session_lifetime int` - Authenticated session lifetime in seconds
- `--id_token_lifetime int` - ID token lifetime in seconds
- `--is_hasura_mapping_enabled` - Enable Hasura mapping
- `--refresh_token_lifetime int` - Refresh token lifetime in seconds
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage applications update_tokens --application_id "app_123" --access_token_lifetime 3600 --refresh_token_lifetime 86400
```

## Common Usage Patterns

### Creating and Configuring a New Application

```bash
# Create a new application
kinde manage applications create --name "My Web App" --type "web"

# Get the application ID from the response, then update settings
kinde manage applications update --application_id "app_123" --homepage_uri "https://myapp.com" --login_uri "https://myapp.com/login"

# Add scopes to the application
kinde manage applications add_application_scope --application_id "app_123" --apiid "api_456" --scope_id "scope_789"

# Configure token lifetimes
kinde manage applications update_tokens --application_id "app_123" --access_token_lifetime 3600 --refresh_token_lifetime 86400
```

### Managing Application Scopes

```bash
# List all applications to find the one you want
kinde manage applications get_all

# Get application details
kinde manage applications get --application_id "app_123"

# Add a scope to an application
kinde manage applications add_application_scope --application_id "app_123" --apiid "api_456" --scope_id "scope_789"

# Remove a scope from an application
kinde manage applications delete_application_scope --application_id "app_123" --apiid "api_456" --scope_id "scope_789"
```

### Application Management

```bash
# List all applications
kinde manage applications get_all --page_size 20

# Get specific application details
kinde manage applications get --application_id "app_123"

# Update application settings
kinde manage applications update --application_id "app_123" --name "New App Name"

# Get application connections
kinde manage applications get_connections --application_id "app_123"

# Get application properties
kinde manage applications get_property_values --application_id "app_123"

# Update a specific property
kinde manage applications update_property --application_id "app_123" --property_key "custom_setting"

# Delete an application
kinde manage applications delete --application_id "app_123"
```

## Application Types

Common application types include:
- `web` - Web applications
- `spa` - Single Page Applications
- `native` - Native mobile/desktop applications
- `machine_to_machine` - Machine-to-machine applications

## Token Configuration Best Practices

1. **Access Token Lifetime**: Set appropriate lifetimes based on your security requirements
2. **Refresh Token Lifetime**: Use longer lifetimes for better user experience
3. **ID Token Lifetime**: Typically shorter than access tokens
4. **Session Lifetime**: Balance security with user experience

## Troubleshooting

- **Application Not Found**: Verify the application ID exists using `get_all`
- **Scope Issues**: Check that the API and scope IDs are correct
- **Permission Errors**: Ensure your M2M application has the required scopes
- **Token Issues**: Verify token lifetimes are set appropriately

## Related Commands

- `kinde manage apis` - Manage APIs that applications can access
- `kinde manage api_keys` - Manage API keys for applications
- `kinde manage connections` - Manage authentication connections
- `kinde manage properties` - Manage custom properties
