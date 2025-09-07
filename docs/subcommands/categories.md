# `kinde manage categories`

The `kinde manage categories` CLI provides various subcommands to manage categories operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Categories commands allow you to manage categorization systems in your Kinde environment. These operations are typically used for organizing and classifying resources.

## Available Commands

### `create`

Create a new category.

**Usage:**
```bash
kinde manage categories create [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage categories create
```

### `get_all`

Retrieve all categories.

**Usage:**
```bash
kinde manage categories get_all [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage categories get_all
```

### `update`

Update an existing category.

**Usage:**
```bash
kinde manage categories update [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage categories update
```

## Common Usage Patterns

### Managing Categories

```bash
# List all categories
kinde manage categories get_all

# Create a new category
kinde manage categories create

# Update an existing category
kinde manage categories update
```

## Best Practices

1. **Consistent Naming**: Use consistent naming conventions for categories
2. **Hierarchical Structure**: Organize categories in a logical hierarchy
3. **Regular Review**: Periodically review and clean up unused categories

## Troubleshooting

- **Permission Errors**: Ensure your M2M application has category management scopes
- **Category Not Found**: Verify the category exists using `get_all`

## Related Commands

- `kinde manage properties` - Manage properties that may be categorized
- `kinde manage organizations` - Manage organizations that may use categories
