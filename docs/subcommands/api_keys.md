# `kinde manage api_keys`

The `kinde manage api_keys` CLI provides various subcommands to manage API key operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

API keys are used for authentication and authorization in Kinde applications. This command group allows you to create, manage, and verify API keys for your Kinde environment.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create](#create)
  - [delete](#delete)
  - [get](#get)
  - [get_all](#get_all)
  - [rotate](#rotate)
  - [verify](#verify)
- [Common Usage Patterns](#common-usage-patterns)
  - [Creating API Keys for Different Scopes](#creating-api-keys-for-different-scopes)
  - [Managing API Keys](#managing-api-keys)
- [Security Best Practices](#security-best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create`

Create a new API key.

**Usage:**
```bash
kinde manage api_keys create [flags]
```

**Flags:**
- `--api_id string` - API ID for the key
- `--name string` - Name for the API key
- `--org_code string` - Organization code
- `--user_id string` - User ID

**Example:**
```bash
kinde manage api_keys create --name "My API Key" --api_id "api_123"
```

### `delete`

Delete an existing API key.

**Usage:**
```bash
kinde manage api_keys delete [flags]
```

**Flags:**
- `--key_id string` - ID of the API key to delete

**Example:**
```bash
kinde manage api_keys delete --key_id "key_123"
```

### `get`

Retrieve details of a specific API key.

**Usage:**
```bash
kinde manage api_keys get [flags]
```

**Flags:**
- `--key_id string` - ID of the API key to retrieve

**Example:**
```bash
kinde manage api_keys get --key_id "key_123"
```

### `get_all`

Retrieve all API keys with optional filtering and pagination.

**Usage:**
```bash
kinde manage api_keys get_all [flags]
```

**Flags:**
- `--org_code string` - Filter by organization code
- `--page_size int` - Number of results to return per page
- `--starting_after string` - Token for pagination (start after this key)
- `--user_id string` - Filter by user ID

**Examples:**
```bash
# Get all API keys
kinde manage api_keys get_all

# Get API keys with pagination
kinde manage api_keys get_all --page_size 10

# Get API keys for a specific organization
kinde manage api_keys get_all --org_code "my-org"

# Get API keys for a specific user
kinde manage api_keys get_all --user_id "user_123"
```

### `rotate`

Rotate (regenerate) an existing API key.

**Usage:**
```bash
kinde manage api_keys rotate [flags]
```

**Flags:**
- `--key_id string` - ID of the API key to rotate

**Example:**
```bash
kinde manage api_keys rotate --key_id "key_123"
```

**Note:** Rotating an API key will generate a new key value, making the old key invalid. Update your applications with the new key value.

### `verify`

Verify the validity of an API key.

**Usage:**
```bash
kinde manage api_keys verify [flags]
```

**Flags:**
- `--api_key string` - The API key to verify

**Example:**
```bash
kinde manage api_keys verify --api_key "your_api_key_here"
```

## Common Usage Patterns

### Creating API Keys for Different Scopes

```bash
# Create an API key for a specific API
kinde manage api_keys create --name "My App API Key" --api_id "api_123"

# Create an API key for a specific organization
kinde manage api_keys create --name "Org API Key" --org_code "my-org"

# Create an API key for a specific user
kinde manage api_keys create --name "User API Key" --user_id "user_123"
```

### Managing API Keys

```bash
# List all API keys
kinde manage api_keys get_all

# Get details of a specific key
kinde manage api_keys get --key_id "key_123"

# Rotate a key for security
kinde manage api_keys rotate --key_id "key_123"

# Verify a key is still valid
kinde manage api_keys verify --api_key "your_key_here"

# Delete an unused key
kinde manage api_keys delete --key_id "key_123"
```

## Security Best Practices

1. **Regular Rotation**: Rotate API keys regularly for enhanced security
2. **Minimal Scope**: Create API keys with only the necessary permissions
3. **Secure Storage**: Store API keys securely and never commit them to version control
4. **Monitor Usage**: Regularly review and audit API key usage
5. **Immediate Revocation**: Delete unused or compromised API keys immediately

## Troubleshooting

- **Authentication Errors**: Ensure you're logged in with appropriate permissions
- **Key Not Found**: Verify the key ID exists using `get_all` command
- **Permission Denied**: Check that your M2M application has the required scopes
- **Invalid Key**: Use the `verify` command to check if a key is still valid

## Related Commands

- `kinde manage applications` - Manage applications that use API keys
- `kinde manage apis` - Manage APIs that API keys authenticate against
- `kinde manage organizations` - Manage organizations for organization-scoped keys
