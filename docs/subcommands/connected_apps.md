# `kinde manage connected_apps`

The `kinde manage connected_apps` CLI provides various subcommands to manage connected applications operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Connected apps commands allow you to manage third-party application integrations in your Kinde environment. These operations are typically used for OAuth integrations and external service connections.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [get_auth_url](#get_auth_url)
  - [get_token](#get_token)
  - [revoke_token](#revoke_token)
- [Common Usage Patterns](#common-usage-patterns)
  - [Managing Connected Apps](#managing-connected-apps)
  - [Complete OAuth Flow](#complete-oauth-flow)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `get_auth_url`

Get the authorization URL for a connected app to initiate OAuth flow.

**Usage:**
```bash
kinde manage connected_apps get_auth_url [flags]
```

**Flags:**
- `--key_code_ref string` - Key code reference for the connected app
- `--org_code string` - Organization code for the authorization
- `--override_callback_url string` - Override the default callback URL
- `--user_id string` - User ID for the authorization

**Example:**
```bash
# Get auth URL for a specific user and organization
kinde manage connected_apps get_auth_url --user_id "user_123" --org_code "my-org" --key_code_ref "app_key_456"

# Get auth URL with custom callback
kinde manage connected_apps get_auth_url --key_code_ref "app_key_456" --override_callback_url "https://myapp.com/callback"
```

### `get_token`

Get an access token for a connected app after successful authorization.

**Usage:**
```bash
kinde manage connected_apps get_token [flags]
```

**Flags:**
- `--session_id string` - Session ID from the authorization flow

**Example:**
```bash
# Get token using session ID from authorization
kinde manage connected_apps get_token --session_id "session_123456"
```

### `revoke_token`

Revoke an access token for a connected app to invalidate it.

**Usage:**
```bash
kinde manage connected_apps revoke_token [flags]
```

**Flags:**
- `--session_id string` - Session ID associated with the token to revoke

**Example:**
```bash
# Revoke token using session ID
kinde manage connected_apps revoke_token --session_id "session_123456"
```

## Common Usage Patterns

### Managing Connected Apps

```bash
# Get authorization URL for OAuth flow
kinde manage connected_apps get_auth_url --key_code_ref "app_key_456" --user_id "user_123" --org_code "my-org"

# Get access token after authorization
kinde manage connected_apps get_token --session_id "session_123456"

# Revoke access token when no longer needed
kinde manage connected_apps revoke_token --session_id "session_123456"
```

### Complete OAuth Flow

```bash
# 1. Get authorization URL for user
kinde manage connected_apps get_auth_url --key_code_ref "app_key_456" --user_id "user_123" --org_code "my-org"

# 2. After user authorizes, get the session ID and exchange for token
kinde manage connected_apps get_token --session_id "session_123456"

# 3. When done, revoke the token
kinde manage connected_apps revoke_token --session_id "session_123456"
```

## Best Practices

1. **Token Security**: Store tokens securely and rotate them regularly
2. **Scope Management**: Request only the minimum required scopes
3. **Error Handling**: Implement proper error handling for token operations
4. **Monitoring**: Monitor token usage and expiration

## Troubleshooting

- **Authorization Issues**: Verify the connected app is properly configured
- **Token Errors**: Check token validity and expiration
- **Permission Errors**: Ensure your M2M application has connected app scopes

## Related Commands

- `kinde manage applications` - Manage applications that may connect to external services
- `kinde manage api_keys` - Manage API keys for external integrations
