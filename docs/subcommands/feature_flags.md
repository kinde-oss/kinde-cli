# `kinde manage feature_flags`

The `kinde manage feature_flags` CLI provides various subcommands to manage feature flags in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Feature flags commands allow you to create, update, and manage feature flags in your Kinde environment. Feature flags enable you to control feature rollouts, A/B testing, and gradual feature releases.

## Available Commands

### `create`

Create a new feature flag.

**Usage:**
```bash
kinde manage feature_flags create [flags]
```

**Flags:**
- `--default_value string` - Default value for the feature flag
- `--description string` - Description of the feature flag
- `--key string` - Unique key for the feature flag
- `--name string` - Name of the feature flag
- `--type string` - Type of the feature flag (bool, str, int)
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Examples:**
```bash
# Create a boolean feature flag
kinde manage feature_flags create --name "New Feature" --key "new_feature" --type "bool" --default_value "false"

# Create a string feature flag
kinde manage feature_flags create --name "API Version" --key "api_version" --type "str" --default_value "v1"

# Create an integer feature flag
kinde manage feature_flags create --name "Max Users" --key "max_users" --type "int" --default_value "100"
```

### `delete`

Delete a feature flag.

**Usage:**
```bash
kinde manage feature_flags delete [flags]
```

**Flags:**
- `--feature_flag_key string` - Feature flag key to delete
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage feature_flags delete --feature_flag_key "new_feature"
```

### `update`

Update an existing feature flag.

**Usage:**
```bash
kinde manage feature_flags update [flags]
```

**Flags:**
- `--allow_override_level string` - Override level (env, org, user)
- `--default_value string` - New default value
- `--description string` - New description
- `--feature_flag_key string` - Feature flag key to update
- `--name string` - New name
- `--type string` - New type
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage feature_flags update --feature_flag_key "new_feature" --default_value "true" --description "Updated description"
```

## Common Usage Patterns

### Creating Feature Flags

```bash
# Create boolean feature flags
kinde manage feature_flags create --name "Enable Dark Mode" --key "dark_mode" --type "bool" --default_value "false"
kinde manage feature_flags create --name "Show Beta Features" --key "beta_features" --type "bool" --default_value "false"

# Create string feature flags
kinde manage feature_flags create --name "Theme Color" --key "theme_color" --type "str" --default_value "blue"
kinde manage feature_flags create --name "API Endpoint" --key "api_endpoint" --type "str" --default_value "https://api.example.com"

# Create integer feature flags
kinde manage feature_flags create --name "Page Size" --key "page_size" --type "int" --default_value "10"
kinde manage feature_flags create --name "Cache TTL" --key "cache_ttl" --type "int" --default_value "3600"
```

### Managing Feature Flags

```bash
# Update feature flag settings
kinde manage feature_flags update --feature_flag_key "dark_mode" --default_value "true"

# Update with override levels
kinde manage feature_flags update --feature_flag_key "beta_features" --allow_override_level "org"

# Update description
kinde manage feature_flags update --feature_flag_key "theme_color" --description "Primary theme color for the application"

# Delete unused feature flags
kinde manage feature_flags delete --feature_flag_key "old_feature"
```

## Feature Flag Types

- **bool**: Boolean values (true/false)
- **str**: String values
- **int**: Integer values

## Override Levels

- **env**: Environment-level overrides
- **org**: Organization-level overrides
- **user**: User-level overrides

## Best Practices

1. **Naming Convention**: Use clear, descriptive names and keys for feature flags
2. **Default Values**: Set sensible default values for all feature flags
3. **Documentation**: Provide clear descriptions for each feature flag
4. **Gradual Rollout**: Use feature flags for gradual feature rollouts
5. **Cleanup**: Remove unused feature flags to reduce complexity

## Common Use Cases

- **Feature Toggles**: Enable/disable features without code deployment
- **A/B Testing**: Test different versions of features
- **Gradual Rollouts**: Roll out features to a subset of users
- **Configuration Management**: Manage application configuration dynamically
- **Emergency Kill Switches**: Quickly disable features if issues arise

## Troubleshooting

- **Feature Flag Not Found**: Verify the feature flag key exists
- **Type Mismatch**: Ensure the type matches the expected value type
- **Permission Errors**: Check that your M2M application has feature flag management scopes

## Related Commands

- `kinde manage environments` - Environment-level feature flag overrides
- `kinde manage organizations` - Organization-level feature flag overrides
- `kinde manage users` - User-level feature flag overrides
