# `kinde manage apis`

The `kinde manage apis` CLI provides various subcommands to manage API operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

APIs in Kinde represent the backend services that your applications can access. This command group allows you to create, configure, and manage APIs, including their scopes, callback URLs, and logout redirect URLs.

## Available Commands

### `add_apis`

Add a new API to your Kinde environment.

**Usage:**
```bash
kinde manage apis add_apis [flags]
```

**Flags:**
- `--audience string` - The audience identifier for the API
- `--name string` - Name of the API
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis add_apis --name "My API" --audience "https://api.example.com"
```

### `add_logout_redirect_urls`

Add logout redirect URLs to an application.

**Usage:**
```bash
kinde manage apis add_logout_redirect_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis add_logout_redirect_urls --app_id "app_123"
```

### `add_redirect_callback_urls`

Add redirect callback URLs to an application.

**Usage:**
```bash
kinde manage apis add_redirect_callback_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis add_redirect_callback_urls --app_id "app_123"
```

### `add_scope`

Add a new scope to an API.

**Usage:**
```bash
kinde manage apis add_scope [flags]
```

**Flags:**
- `--apiid string` - API ID
- `--description string` - Description of the scope
- `--key string` - Scope key/identifier
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis add_scope --apiid "api_123" --key "read:users" --description "Read user data"
```

### `delete_api`

Delete an API from your Kinde environment.

**Usage:**
```bash
kinde manage apis delete_api [flags]
```

**Flags:**
- `--apiid string` - API ID to delete
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis delete_api --apiid "api_123"
```

### `delete_callback_urls`

Delete callback URLs from an application.

**Usage:**
```bash
kinde manage apis delete_callback_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `--urls string` - URLs to delete (comma-separated)
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis delete_callback_urls --app_id "app_123" --urls "https://app.com/callback"
```

### `delete_logout_urls`

Delete logout URLs from an application.

**Usage:**
```bash
kinde manage apis delete_logout_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `--urls string` - URLs to delete (comma-separated)
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis delete_logout_urls --app_id "app_123" --urls "https://app.com/logout"
```

### `delete_scope`

Delete a scope from an API.

**Usage:**
```bash
kinde manage apis delete_scope [flags]
```

**Flags:**
- `--apiid string` - API ID
- `--scope_id string` - Scope ID to delete
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis delete_scope --apiid "api_123" --scope_id "scope_456"
```

### `get_api`

Retrieve details of a specific API.

**Usage:**
```bash
kinde manage apis get_api [flags]
```

**Flags:**
- `--apiid string` - API ID to retrieve
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis get_api --apiid "api_123"
```

### `get_apis`

Retrieve all APIs in your Kinde environment.

**Usage:**
```bash
kinde manage apis get_apis [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis get_apis
```

### `get_callback_urls`

Retrieve callback URLs for an application.

**Usage:**
```bash
kinde manage apis get_callback_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis get_callback_urls --app_id "app_123"
```

### `get_logout_urls`

Retrieve logout URLs for an application.

**Usage:**
```bash
kinde manage apis get_logout_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis get_logout_urls --app_id "app_123"
```

### `get_scope`

Retrieve details of a specific scope.

**Usage:**
```bash
kinde manage apis get_scope [flags]
```

**Flags:**
- `--apiid string` - API ID
- `--scope_id string` - Scope ID to retrieve
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis get_scope --apiid "api_123" --scope_id "scope_456"
```

### `get_scopes`

Retrieve all scopes for an API.

**Usage:**
```bash
kinde manage apis get_scopes [flags]
```

**Flags:**
- `--apiid string` - API ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis get_scopes --apiid "api_123"
```

### `replace_logout_redirect_urls`

Replace all logout redirect URLs for an application.

**Usage:**
```bash
kinde manage apis replace_logout_redirect_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis replace_logout_redirect_urls --app_id "app_123"
```

### `replace_redirect_callback_urls`

Replace all redirect callback URLs for an application.

**Usage:**
```bash
kinde manage apis replace_redirect_callback_urls [flags]
```

**Flags:**
- `--app_id string` - Application ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis replace_redirect_callback_urls --app_id "app_123"
```

### `update_applications`

Update applications associated with an API.

**Usage:**
```bash
kinde manage apis update_applications [flags]
```

**Flags:**
- `--apiid string` - API ID
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis update_applications --apiid "api_123"
```

### `update_scope`

Update an existing scope.

**Usage:**
```bash
kinde manage apis update_scope [flags]
```

**Flags:**
- `--apiid string` - API ID
- `--description string` - New description for the scope
- `--scope_id string` - Scope ID to update
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage apis update_scope --apiid "api_123" --scope_id "scope_456" --description "Updated scope description"
```

## Common Usage Patterns

### Setting Up a New API

```bash
# Create a new API
kinde manage apis add_apis --name "My Backend API" --audience "https://api.myapp.com"

# Add scopes to the API
kinde manage apis add_scope --apiid "api_123" --key "read:users" --description "Read user data"
kinde manage apis add_scope --apiid "api_123" --key "write:users" --description "Write user data"

# List all scopes for the API
kinde manage apis get_scopes --apiid "api_123"
```

### Managing Application URLs

```bash
# Get current callback URLs
kinde manage apis get_callback_urls --app_id "app_123"

# Add new callback URLs
kinde manage apis add_redirect_callback_urls --app_id "app_123"

# Replace all callback URLs
kinde manage apis replace_redirect_callback_urls --app_id "app_123"

# Manage logout URLs
kinde manage apis get_logout_urls --app_id "app_123"
kinde manage apis add_logout_redirect_urls --app_id "app_123"
```

### API Management

```bash
# List all APIs
kinde manage apis get_apis

# Get specific API details
kinde manage apis get_api --apiid "api_123"

# Update API applications
kinde manage apis update_applications --apiid "api_123"

# Delete an API
kinde manage apis delete_api --apiid "api_123"
```

## Best Practices

1. **API Design**: Use clear, descriptive names and audiences for your APIs
2. **Scope Management**: Create granular scopes that follow the principle of least privilege
3. **URL Management**: Regularly review and update callback and logout URLs
4. **Documentation**: Keep track of API IDs and scope IDs for easy reference
5. **Security**: Regularly audit API permissions and remove unused scopes

## Troubleshooting

- **API Not Found**: Verify the API ID exists using `get_apis`
- **Scope Issues**: Check scope IDs using `get_scopes` for the specific API
- **URL Problems**: Ensure URLs are properly formatted and accessible
- **Permission Errors**: Verify your M2M application has the required scopes

## Related Commands

- `kinde manage applications` - Manage applications that use APIs
- `kinde manage api_keys` - Manage API keys for authentication
- `kinde manage permissions` - Manage permissions that may use API scopes
