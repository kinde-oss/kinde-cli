# Permissions Management

## Overview

Permissions define what actions users and applications can perform in your Kinde environment. They are granular access controls that can be assigned to roles and users to create a flexible authorization system. This command group allows you to create, manage, and organize permissions for your applications and users.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create](#create)
  - [delete](#delete)
  - [get_all](#get_all)
  - [update](#update)
- [Common Usage Patterns](#common-usage-patterns)
  - [Creating a Permission System](#creating-a-permission-system)
  - [Managing Permissions](#managing-permissions)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create`

Create a new permission with specified name, key, and description.

**Usage:**
```bash
kinde manage permissions create [flags]
```

**Flags:**
- `--description string` - Description of the permission
- `--key string` - Unique key identifier for the permission
- `--name string` - Display name for the permission

**Example:**
```bash
kinde manage permissions create --name "Read Users" --key "users:read" --description "Permission to read user information"
```

### `delete`

Delete a permission permanently from your Kinde environment.

**Usage:**
```bash
kinde manage permissions delete [flags]
```

**Flags:**
- `--permission_id string` - ID of the permission to delete

**Example:**
```bash
kinde manage permissions delete --permission_id "perm_123456789"
```

### `get_all`

List all permissions in your Kinde environment with pagination support.

**Usage:**
```bash
kinde manage permissions get_all [flags]
```

**Flags:**
- `--next_token string` - Token for pagination to get next page of results
- `--page_size int` - Number of permissions to return per page

**Example:**
```bash
kinde manage permissions get_all --page_size 20
```

### `update`

Update an existing permission's name, key, or description.

**Usage:**
```bash
kinde manage permissions update [flags]
```

**Flags:**
- `--description string` - New description for the permission
- `--key string` - New key identifier for the permission
- `--name string` - New display name for the permission
- `--permission_id string` - ID of the permission to update

**Example:**
```bash
kinde manage permissions update --permission_id "perm_123456789" --name "Read User Data" --description "Updated permission to read user data"
```

## Common Usage Patterns

### Creating a Permission System

When setting up a new application, you'll typically create a set of permissions that define what actions can be performed:

```bash
# Create basic CRUD permissions for users
kinde manage permissions create --name "Create Users" --key "users:create" --description "Permission to create new users"
kinde manage permissions create --name "Read Users" --key "users:read" --description "Permission to read user information"
kinde manage permissions create --name "Update Users" --key "users:update" --description "Permission to update user information"
kinde manage permissions create --name "Delete Users" --key "users:delete" --description "Permission to delete users"

# Create permissions for different resources
kinde manage permissions create --name "Manage Organizations" --key "orgs:manage" --description "Permission to manage organizations"
kinde manage permissions create --name "View Analytics" --key "analytics:view" --description "Permission to view analytics data"
```

### Managing Permissions

Regular maintenance of your permission system:

```bash
# List all permissions to review current setup
kinde manage permissions get_all --page_size 50

# Update a permission that needs clarification
kinde manage permissions update --permission_id "perm_123" --description "Updated description for better clarity"

# Remove unused permissions
kinde manage permissions delete --permission_id "perm_unused_456"
```

## Best Practices

### Permission Naming and Organization

1. **Use Consistent Naming Conventions**
   - Follow a pattern like `resource:action` (e.g., `users:read`, `orders:create`)
   - Use lowercase with colons or underscores for separation
   - Be descriptive but concise

2. **Group Related Permissions**
   - Create permissions for related resources together
   - Use consistent prefixes for related functionality
   - Document permission relationships

3. **Plan Permission Hierarchy**
   - Consider which permissions are prerequisites for others
   - Design permissions that can be combined effectively
   - Avoid overly granular permissions that are hard to manage

### Security Considerations

1. **Principle of Least Privilege**
   - Create permissions that grant only the minimum necessary access
   - Avoid overly broad permissions like "admin" or "all"
   - Regularly review and audit permission usage

2. **Permission Documentation**
   - Always provide clear descriptions for permissions
   - Document what each permission allows users to do
   - Keep descriptions up to date as functionality changes

3. **Regular Audits**
   - Periodically review all permissions in your system
   - Remove unused or obsolete permissions
   - Ensure permissions align with current business requirements

## Troubleshooting

### Common Issues

**Permission Creation Fails**
- Ensure the permission key is unique and follows naming conventions
- Check that required fields (name, key) are provided
- Verify you have the necessary permissions to create permissions

**Permission Not Found**
- Double-check the permission ID is correct
- Ensure the permission exists in your current environment
- Verify you're using the correct domain/environment

**Pagination Issues**
- Use the `next_token` from previous responses for pagination
- Adjust `page_size` if you're getting too many or too few results
- Handle pagination properly in automated scripts

### Getting Help

For additional help with permissions:

```bash
# Get help for the permissions command group
kinde manage permissions --help

# Get help for a specific command
kinde manage permissions create --help
kinde manage permissions update --help
```

## Related Commands

- **Roles Management** (`kinde manage roles`) - Assign permissions to roles
- **Users Management** (`kinde manage users`) - Manage user accounts and their permissions
- **Organizations Management** (`kinde manage organizations`) - Manage organization-level permissions
- **Applications Management** (`kinde manage applications`) - Configure application permissions

### Workflow Integration

Permissions are typically used in conjunction with:

1. **Roles** - Group permissions into logical roles
2. **Users** - Assign roles and permissions to specific users
3. **Applications** - Configure which permissions applications can request
4. **Organizations** - Set organization-specific permission overrides

Example workflow:
```bash
# 1. Create permissions
kinde manage permissions create --name "Read Data" --key "data:read" --description "Read access to data"

# 2. Create a role with the permission
kinde manage roles create --name "Data Reader" --description "Role for reading data"
kinde manage roles update_permissions --role_id "role_123" --permission_ids "perm_456"

# 3. Assign role to user
kinde manage users update --user_id "user_789" --roles "role_123"
```
