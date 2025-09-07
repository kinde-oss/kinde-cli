# `kinde manage connected_apps`

The `kinde manage connected_apps` CLI provides various subcommands to manage connected applications operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Connected apps commands allow you to manage third-party application integrations in your Kinde environment. These operations are typically used for OAuth integrations and external service connections.

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [get_auth_url](#get_auth_url)
  - [get_token](#get_token)
  - [revoke_token](#revoke_token)
- [Common Usage Patterns](#common-usage-patterns)
  - [Managing Connected Apps](#managing-connected-apps)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `get_auth_url`

Get the authorization URL for a connected app.

**Usage:**
```bash
kinde manage connected_apps get_auth_url [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage connected_apps get_auth_url
```

### `get_token`

Get an access token for a connected app.

**Usage:**
```bash
kinde manage connected_apps get_token [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage connected_apps get_token
```

### `revoke_token`

Revoke an access token for a connected app.

**Usage:**
```bash
kinde manage connected_apps revoke_token [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage connected_apps revoke_token
```

## Common Usage Patterns

### Managing Connected Apps

```bash
# Get authorization URL for OAuth flow
kinde manage connected_apps get_auth_url

# Get access token after authorization
kinde manage connected_apps get_token

# Revoke access token when no longer needed
kinde manage connected_apps revoke_token
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
