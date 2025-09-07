# `kinde manage categories`

The `kinde manage categories` CLI provides various subcommands to manage categories operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Categories commands allow you to manage categorization systems in your Kinde environment. These operations are typically used for organizing and classifying resources.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create](#create)
  - [get_all](#get_all)
  - [update](#update)
- [Common Usage Patterns](#common-usage-patterns)
  - [Managing Categories](#managing-categories)
  - [Complete Category Management Workflow](#complete-category-management-workflow)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create`

Create a new category for organizing and classifying resources in your Kinde environment.

**Usage:**
```bash
kinde manage categories create [flags]
```

**Flags:**
- `--context string` - Context or description for the category
- `--name string` - Name of the category

**Example:**
```bash
# Create a basic category
kinde manage categories create --name "User Types"

# Create a category with context
kinde manage categories create --name "Department" --context "Organizational departments"
```

### `get_all`

Retrieve all categories with optional pagination support.

**Usage:**
```bash
kinde manage categories get_all [flags]
```

**Flags:**
- `--ending_before string` - Pagination: get categories ending before this cursor
- `--page_size int` - Number of categories to return per page
- `--starting_after string` - Pagination: get categories starting after this cursor

**Example:**
```bash
# Get all categories
kinde manage categories get_all

# Get categories with pagination
kinde manage categories get_all --page_size 10 --starting_after "cursor_123"
```

### `update`

Update an existing category's name and properties.

**Usage:**
```bash
kinde manage categories update [flags]
```

**Flags:**
- `--category_id string` - ID of the category to update
- `--name string` - New name for the category

**Example:**
```bash
# Update a category name
kinde manage categories update --category_id "cat_123" --name "Updated Category Name"
```

## Common Usage Patterns

### Managing Categories

```bash
# List all categories
kinde manage categories get_all

# Create a new category
kinde manage categories create --name "User Types" --context "Classification for different user types"

# Update an existing category
kinde manage categories update --category_id "cat_123" --name "Updated Category Name"

# Get categories with pagination
kinde manage categories get_all --page_size 5 --starting_after "cursor_123"
```

### Complete Category Management Workflow

```bash
# 1. List existing categories
kinde manage categories get_all

# 2. Create new categories for your organization
kinde manage categories create --name "Departments" --context "Organizational departments"
kinde manage categories create --name "User Roles" --context "User role classifications"

# 3. Update a category if needed
kinde manage categories update --category_id "cat_123" --name "Updated Department Names"

# 4. Verify the changes
kinde manage categories get_all
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
