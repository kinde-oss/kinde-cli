# `kinde manage business`

The `kinde manage business` CLI provides various subcommands to manage business operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Business commands allow you to manage business settings and configuration in your Kinde environment. These operations are typically used for configuring business-level settings and preferences.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [get](#get)
  - [update](#update)
- [Common Usage Patterns](#common-usage-patterns)
  - [Managing Business Settings](#managing-business-settings)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `get`

Retrieve business settings and configuration.

**Usage:**
```bash
kinde manage business get [flags]
```

**Flags:**

**Example:**
```bash
kinde manage business get
```

### `update`

Update business settings and configuration.

**Usage:**
```bash
kinde manage business update [flags]
```

**Flags:**

**Example:**
```bash
kinde manage business update
```

## Common Usage Patterns

### Managing Business Settings

```bash
# Get current business settings
kinde manage business get

# Update business settings
kinde manage business update
```

## Best Practices

1. **Regular Review**: Periodically review business settings to ensure they're current
2. **Backup Settings**: Keep track of important business configuration changes
3. **Documentation**: Document any custom business settings for your team

## Troubleshooting

- **Permission Errors**: Ensure your M2M application has business management scopes
- **Configuration Issues**: Verify that business settings are properly configured

## Related Commands

- `kinde manage organizations` - Manage organizations within the business
- `kinde manage applications` - Manage applications for the business
