# Environments Management

## Overview

Environments represent your Kinde domain instances where you can configure environment-specific settings, feature flags, and branding. This command group allows you to manage environment-level configurations, including logos, feature flag overrides, and environment details. These settings apply globally to all applications and users within the environment.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [add_logo](#add_logo)
  - [delete_feature_flag_override](#delete_feature_flag_override)
  - [delete_feature_flag_overrides](#delete_feature_flag_overrides)
  - [delete_logo](#delete_logo)
  - [get](#get)
  - [get_feature_flags](#get_feature_flags)
  - [read_logo](#read_logo)
  - [update_feature_flag_override](#update_feature_flag_override)
- [Common Usage Patterns](#common-usage-patterns)
  - [Environment Branding](#environment-branding)
  - [Feature Flag Management](#feature-flag-management)
  - [Environment Configuration](#environment-configuration)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `add_logo`

Upload and set a logo for the environment. This logo will be displayed in authentication flows and other environment-specific UI elements.

**Usage:**
```bash
kinde manage environments add_logo [flags]
```

**Flags:**
- `--logo string` - Path to the logo file or logo data
- `--type string` - Type of logo (e.g., "png", "jpg", "svg")

**Example:**
```bash
kinde manage environments add_logo --logo "/path/to/logo.png" --type "png"
```

### `delete_feature_flag_override`

Remove a specific feature flag override for the environment, reverting it to the default value.

**Usage:**
```bash
kinde manage environments delete_feature_flag_override [flags]
```

**Flags:**
- `--feature_flag_key string` - Key of the feature flag to remove override for

**Example:**
```bash
kinde manage environments delete_feature_flag_override --feature_flag_key "new_dashboard"
```

### `delete_feature_flag_overrides`

Remove all feature flag overrides for the environment, reverting all flags to their default values.

**Usage:**
```bash
kinde manage environments delete_feature_flag_overrides [flags]
```

**Flags:**
No additional flags required.

**Example:**
```bash
kinde manage environments delete_feature_flag_overrides
```

### `delete_logo`

Remove the environment logo, reverting to the default Kinde branding.

**Usage:**
```bash
kinde manage environments delete_logo [flags]
```

**Flags:**
- `--type string` - Type of logo to delete (e.g., "png", "jpg", "svg")

**Example:**
```bash
kinde manage environments delete_logo --type "png"
```

### `get`

Retrieve detailed information about the current environment configuration.

**Usage:**
```bash
kinde manage environments get [flags]
```

**Flags:**
No additional flags required.

**Example:**
```bash
kinde manage environments get
```

### `get_feature_flags`

List all feature flags and their environment-level overrides for the current environment.

**Usage:**
```bash
kinde manage environments get_feature_flags [flags]
```

**Flags:**
No additional flags required.

**Example:**
```bash
kinde manage environments get_feature_flags
```

### `read_logo`

Retrieve the current environment logo data.

**Usage:**
```bash
kinde manage environments read_logo [flags]
```

**Flags:**
No additional flags required.

**Example:**
```bash
kinde manage environments read_logo
```

### `update_feature_flag_override`

Set or update a feature flag override for the environment, allowing you to customize feature behavior at the environment level.

**Usage:**
```bash
kinde manage environments update_feature_flag_override [flags]
```

**Flags:**
- `--feature_flag_key string` - Key of the feature flag to override
- `--value string` - New value for the feature flag override

**Example:**
```bash
kinde manage environments update_feature_flag_override --feature_flag_key "new_dashboard" --value "true"
```

## Common Usage Patterns

### Environment Branding

Setting up custom branding for your environment:

```bash
# Upload a logo for your environment
kinde manage environments add_logo --logo "/path/to/company-logo.png" --type "png"

# Check the current logo
kinde manage environments read_logo

# Remove the logo if needed
kinde manage environments delete_logo --type "png"
```

### Feature Flag Management

Managing feature flags at the environment level:

```bash
# View all feature flags and their current values
kinde manage environments get_feature_flags

# Enable a feature for the entire environment
kinde manage environments update_feature_flag_override --feature_flag_key "new_ui" --value "true"

# Disable a feature for the entire environment
kinde manage environments update_feature_flag_override --feature_flag_key "beta_features" --value "false"

# Remove a specific feature flag override
kinde manage environments delete_feature_flag_override --feature_flag_key "new_ui"

# Reset all feature flag overrides to defaults
kinde manage environments delete_feature_flag_overrides
```

### Environment Configuration

Getting information about your environment:

```bash
# Get detailed environment information
kinde manage environments get

# Check current feature flag configuration
kinde manage environments get_feature_flags

# Verify logo is set correctly
kinde manage environments read_logo
```

## Best Practices

### Environment Branding

1. **Logo Requirements**
   - Use high-quality images with appropriate dimensions
   - Ensure logos are optimized for web display
   - Test logos across different authentication flows
   - Consider accessibility and contrast requirements

2. **Brand Consistency**
   - Maintain consistent branding across all environments
   - Use appropriate logo formats (PNG, JPG, SVG)
   - Test logo display in different contexts

### Feature Flag Management

1. **Environment-Level Overrides**
   - Use environment overrides for environment-specific configurations
   - Document why certain overrides are set
   - Regularly review and clean up unused overrides

2. **Feature Rollout Strategy**
   - Use environment overrides for gradual feature rollouts
   - Test features in staging environments before production
   - Monitor feature flag usage and impact

3. **Override Hierarchy**
   - Understand the hierarchy: default → environment → organization → user
   - Use environment overrides for global environment settings
   - Avoid conflicts with organization or user-level overrides

### Environment Configuration

1. **Regular Audits**
   - Periodically review environment configuration
   - Check for unused feature flag overrides
   - Verify branding consistency

2. **Documentation**
   - Document environment-specific configurations
   - Maintain records of feature flag overrides and their purposes
   - Keep track of branding changes and updates

## Troubleshooting

### Common Issues

**Logo Upload Fails**
- Ensure the logo file exists and is accessible
- Check that the logo format is supported
- Verify file size is within limits
- Ensure you have the necessary permissions to upload logos

**Feature Flag Override Not Working**
- Verify the feature flag key is correct and exists
- Check that the value format matches the feature flag type
- Ensure there are no conflicting overrides at higher levels
- Confirm the feature flag is properly configured

**Environment Information Not Found**
- Verify you're using the correct domain/environment
- Check that the environment exists and is accessible
- Ensure you have the necessary permissions to access environment data

**Logo Not Displaying**
- Verify the logo was uploaded successfully
- Check that the logo type matches what was uploaded
- Ensure the logo is being referenced correctly in applications
- Test logo display in different browsers and contexts

### Getting Help

For additional help with environments:

```bash
# Get help for the environments command group
kinde manage environments --help

# Get help for a specific command
kinde manage environments add_logo --help
kinde manage environments update_feature_flag_override --help
```

## Related Commands

- **Feature Flags Management** (`kinde manage feature_flags`) - Create and manage feature flags
- **Organizations Management** (`kinde manage organizations`) - Set organization-level feature flag overrides
- **Users Management** (`kinde manage users`) - Set user-level feature flag overrides
- **Applications Management** (`kinde manage applications`) - Configure application-specific settings

### Workflow Integration

Environments work in conjunction with:

1. **Feature Flags** - Environment-level overrides for feature behavior
2. **Organizations** - Organization-specific overrides that inherit from environment
3. **Users** - User-specific overrides that inherit from organization and environment
4. **Applications** - Applications that use environment branding and feature flags

Example workflow:
```bash
# 1. Create a feature flag
kinde manage feature_flags create --name "New Dashboard" --key "new_dashboard" --type "bool" --default_value "false"

# 2. Enable it for the entire environment
kinde manage environments update_feature_flag_override --feature_flag_key "new_dashboard" --value "true"

# 3. Check the current environment configuration
kinde manage environments get_feature_flags

# 4. Set up environment branding
kinde manage environments add_logo --logo "/path/to/logo.png" --type "png"

# 5. Verify the setup
kinde manage environments get
```

### Integration with Applications

Environment settings affect:
- **Authentication Flows** - Custom logos and branding
- **Feature Availability** - Environment-level feature flag overrides
- **User Experience** - Consistent branding across all applications
- **Configuration Management** - Centralized environment settings

### Multi-Environment Management

When working with multiple environments:

```bash
# Development environment
kinde manage environments update_feature_flag_override --domain "dev.mycompany.kinde.com" --feature_flag_key "debug_mode" --value "true"

# Staging environment
kinde manage environments update_feature_flag_override --domain "staging.mycompany.kinde.com" --feature_flag_key "debug_mode" --value "false"

# Production environment
kinde manage environments update_feature_flag_override --domain "mycompany.kinde.com" --feature_flag_key "debug_mode" --value "false"
```
