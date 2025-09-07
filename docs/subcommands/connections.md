# `kinde manage connections`

The `kinde manage connections` CLI provides various subcommands to manage authentication connections in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Connections commands allow you to manage authentication connections (like social logins, SAML, etc.) in your Kinde environment. These operations are used to configure how users can authenticate with your applications.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create](#create)
  - [delete](#delete)
  - [enable](#enable)
  - [enable_org](#enable_org)
  - [get](#get)
  - [get_all](#get_all)
  - [remove](#remove)
  - [remove_org](#remove_org)
  - [replace](#replace)
  - [update](#update)
- [Common Usage Patterns](#common-usage-patterns)
  - [Setting Up Authentication Connections](#setting-up-authentication-connections)
  - [Managing Connection Access](#managing-connection-access)
  - [Connection Management](#connection-management)
- [Connection Types](#connection-types)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create`

Create a new authentication connection.

**Usage:**
```bash
kinde manage connections create [flags]
```

**Flags:**
- `--display_name string` - Display name for the connection
- `--name string` - Name of the connection
- `--options.type string` - Connection type
- `--organization_code string` - Organization code

**Example:**
```bash
kinde manage connections create --name "Google OAuth" --display_name "Google" --options.type "oauth2"
```

### `delete`

Delete an authentication connection.

**Usage:**
```bash
kinde manage connections delete [flags]
```

**Flags:**
- `--connection_id string` - Connection ID to delete

**Example:**
```bash
kinde manage connections delete --connection_id "conn_123"
```

### `enable`

Enable a connection for an application.

**Usage:**
```bash
kinde manage connections enable [flags]
```

**Flags:**
- `--application_id string` - Application ID
- `--connection_id string` - Connection ID to enable

**Example:**
```bash
kinde manage connections enable --application_id "app_123" --connection_id "conn_456"
```

### `enable_org`

Enable a connection for an organization.

**Usage:**
```bash
kinde manage connections enable_org [flags]
```

**Flags:**
- `--connection_id string` - Connection ID to enable
- `--organization_code string` - Organization code

**Example:**
```bash
kinde manage connections enable_org --connection_id "conn_123" --organization_code "my-org"
```

### `get`

Retrieve details of a specific connection.

**Usage:**
```bash
kinde manage connections get [flags]
```

**Flags:**
- `--connection_id string` - Connection ID to retrieve

**Example:**
```bash
kinde manage connections get --connection_id "conn_123"
```

### `get_all`

Retrieve all connections with optional filtering and pagination.

**Usage:**
```bash
kinde manage connections get_all [flags]
```

**Flags:**
- `--ending_before string` - Token for pagination (end before this connection)
- `--home_realm_domain string` - Filter by home realm domain
- `--page_size int` - Number of results to return per page
- `--starting_after string` - Token for pagination (start after this connection)

**Examples:**
```bash
# Get all connections
kinde manage connections get_all

# Get connections with pagination
kinde manage connections get_all --page_size 10

# Filter by home realm domain
kinde manage connections get_all --home_realm_domain "example.com"
```

### `remove`

Remove a connection from an application.

**Usage:**
```bash
kinde manage connections remove [flags]
```

**Flags:**
- `--application_id string` - Application ID
- `--connection_id string` - Connection ID to remove

**Example:**
```bash
kinde manage connections remove --application_id "app_123" --connection_id "conn_456"
```

### `remove_org`

Remove a connection from an organization.

**Usage:**
```bash
kinde manage connections remove_org [flags]
```

**Flags:**
- `--connection_id string` - Connection ID to remove
- `--organization_code string` - Organization code

**Example:**
```bash
kinde manage connections remove_org --connection_id "conn_123" --organization_code "my-org"
```

### `replace`

Replace connection settings.

**Usage:**
```bash
kinde manage connections replace [flags]
```

**Flags:**
- `--connection_id string` - Connection ID to replace
- `--display_name string` - New display name
- `--name string` - New name
- `--options.type string` - New connection type

**Example:**
```bash
kinde manage connections replace --connection_id "conn_123" --name "Updated Google OAuth" --display_name "Google Login"
```

### `update`

Update connection settings.

**Usage:**
```bash
kinde manage connections update [flags]
```

**Flags:**
- `--connection_id string` - Connection ID to update
- `--display_name string` - New display name
- `--name string` - New name
- `--options.type string` - New connection type

**Example:**
```bash
kinde manage connections update --connection_id "conn_123" --display_name "Updated Display Name"
```

## Common Usage Patterns

### Setting Up Authentication Connections

```bash
# Create a new OAuth connection
kinde manage connections create --name "Google OAuth" --display_name "Google" --options.type "oauth2"

# Create a SAML connection
kinde manage connections create --name "SAML SSO" --display_name "Corporate SSO" --options.type "saml"

# List all connections
kinde manage connections get_all

# Get specific connection details
kinde manage connections get --connection_id "conn_123"
```

### Managing Connection Access

```bash
# Enable connection for an application
kinde manage connections enable --application_id "app_123" --connection_id "conn_456"

# Enable connection for an organization
kinde manage connections enable_org --connection_id "conn_123" --organization_code "my-org"

# Remove connection from application
kinde manage connections remove --application_id "app_123" --connection_id "conn_456"

# Remove connection from organization
kinde manage connections remove_org --connection_id "conn_123" --organization_code "my-org"
```

### Connection Management

```bash
# Update connection settings
kinde manage connections update --connection_id "conn_123" --display_name "New Display Name"

# Replace connection configuration
kinde manage connections replace --connection_id "conn_123" --name "New Name" --options.type "oauth2"

# Delete a connection
kinde manage connections delete --connection_id "conn_123"
```

## Connection Types

Common connection types include:
- `oauth2` - OAuth 2.0 connections (Google, Facebook, etc.)
- `saml` - SAML SSO connections
- `openid_connect` - OpenID Connect connections
- `password` - Username/password connections

## Best Practices

1. **Naming Convention**: Use clear, descriptive names for connections
2. **Display Names**: Use user-friendly display names for the login UI
3. **Organization Scope**: Consider which organizations should have access to each connection
4. **Testing**: Test connections thoroughly before enabling for production applications
5. **Documentation**: Keep track of connection configurations and their purposes

## Troubleshooting

- **Connection Not Found**: Verify the connection ID exists using `get_all`
- **Application Issues**: Check that the application ID is correct
- **Organization Issues**: Verify the organization code is valid
- **Permission Errors**: Ensure your M2M application has connection management scopes

## Related Commands

- `kinde manage applications` - Manage applications that use connections
- `kinde manage organizations` - Manage organizations that may have connection access
- `kinde manage users` - Manage users who authenticate through connections
