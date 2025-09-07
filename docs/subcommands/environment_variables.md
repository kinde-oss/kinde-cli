# `kinde manage environment_variables`

The `kinde manage environment_variables` CLI provides various subcommands to manage environment variables in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Environment variables commands allow you to manage configuration variables in your Kinde environment. These operations are used to store and manage application configuration, secrets, and other environment-specific settings.

## Available Commands

### `create`

Create a new environment variable.

**Usage:**
```bash
kinde manage environment_variables create [flags]
```

**Flags:**
- `--is_secret` - Mark the variable as secret (encrypted)
- `--key string` - Variable key/name
- `--value string` - Variable value
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Examples:**
```bash
# Create a regular environment variable
kinde manage environment_variables create --key "API_URL" --value "https://api.example.com"

# Create a secret environment variable
kinde manage environment_variables create --key "API_SECRET" --value "secret123" --is_secret
```

### `delete`

Delete an environment variable.

**Usage:**
```bash
kinde manage environment_variables delete [flags]
```

**Flags:**
- `--variable_id string` - Variable ID to delete
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage environment_variables delete --variable_id "var_123"
```

### `get`

Retrieve details of a specific environment variable.

**Usage:**
```bash
kinde manage environment_variables get [flags]
```

**Flags:**
- `--variable_id string` - Variable ID to retrieve
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage environment_variables get --variable_id "var_123"
```

### `get_all`

Retrieve all environment variables.

**Usage:**
```bash
kinde manage environment_variables get_all [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage environment_variables get_all
```

### `update`

Update an existing environment variable.

**Usage:**
```bash
kinde manage environment_variables update [flags]
```

**Flags:**
- `--is_secret` - Mark the variable as secret (encrypted)
- `--key string` - New variable key/name
- `--value string` - New variable value
- `--variable_id string` - Variable ID to update
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage environment_variables update --variable_id "var_123" --value "new_value"
```

## Common Usage Patterns

### Managing Environment Variables

```bash
# List all environment variables
kinde manage environment_variables get_all

# Create configuration variables
kinde manage environment_variables create --key "DATABASE_URL" --value "postgresql://localhost:5432/mydb"
kinde manage environment_variables create --key "REDIS_URL" --value "redis://localhost:6379"

# Create secret variables
kinde manage environment_variables create --key "JWT_SECRET" --value "my-secret-key" --is_secret
kinde manage environment_variables create --key "API_KEY" --value "sk-123456789" --is_secret

# Get specific variable details
kinde manage environment_variables get --variable_id "var_123"

# Update a variable
kinde manage environment_variables update --variable_id "var_123" --value "updated_value"

# Delete a variable
kinde manage environment_variables delete --variable_id "var_123"
```

## Best Practices

1. **Secret Management**: Use `--is_secret` for sensitive data like API keys, passwords, and tokens
2. **Naming Convention**: Use clear, consistent naming for variable keys
3. **Environment Separation**: Consider using different variables for different environments
4. **Regular Cleanup**: Remove unused environment variables to reduce clutter
5. **Documentation**: Document the purpose of each environment variable

## Security Considerations

- **Secret Variables**: Variables marked with `--is_secret` are encrypted and not visible in plain text
- **Access Control**: Ensure only authorized users can manage environment variables
- **Audit Trail**: Keep track of who creates, updates, or deletes environment variables

## Troubleshooting

- **Variable Not Found**: Verify the variable ID exists using `get_all`
- **Permission Errors**: Ensure your M2M application has environment variable management scopes
- **Value Issues**: Check that variable values are properly formatted

## Related Commands

- `kinde manage applications` - Applications that may use environment variables
- `kinde manage environments` - Environment-specific configurations
