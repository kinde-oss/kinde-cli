# Roles Management

## Overview

Roles provide a way to group permissions and can be assigned to users to create a flexible and manageable access control system. Roles help organize permissions into logical groups, making it easier to manage user access and maintain security policies. This command group allows you to create, manage, and configure roles with their associated permissions and API scopes.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [add_scope](#add_scope)
  - [create](#create)
  - [delete](#delete)
  - [delete_scope](#delete_scope)
  - [get](#get)
  - [get_all](#get_all)
  - [get_permissions](#get_permissions)
  - [get_scopes](#get_scopes)
  - [remove_permission](#remove_permission)
  - [update](#update)
  - [update_permissions](#update_permissions)
- [Common Usage Patterns](#common-usage-patterns)
  - [Creating a Role System](#creating-a-role-system)
  - [Managing Role Permissions](#managing-role-permissions)
  - [Managing Role Scopes](#managing-role-scopes)
- [Role Design Principles](#role-design-principles)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `add_scope`

Add an API scope to a role, granting the role access to specific API endpoints and functionality.

**Usage:**
```bash
kinde manage roles add_scope [flags]
```

**Flags:**
- `--role_id string` - ID of the role to add the scope to
- `--scope_id string` - ID of the API scope to add to the role

**Example:**
```bash
kinde manage roles add_scope --role_id "role_123456789" --scope_id "scope_987654321"
```

### `create`

Create a new role with specified name, key, and description. Roles can be marked as default roles for new users.

**Usage:**
```bash
kinde manage roles create [flags]
```

**Flags:**
- `--description string` - Description of the role and its purpose
- `--is_default_role` - Mark this role as a default role for new users
- `--key string` - Unique key identifier for the role
- `--name string` - Display name for the role

**Example:**
```bash
kinde manage roles create --name "Administrator" --key "admin" --description "Full administrative access to the system"
```

### `delete`

Delete a role permanently from your Kinde environment. This will remove the role from all users who have it assigned.

**Usage:**
```bash
kinde manage roles delete [flags]
```

**Flags:**
- `--role_id string` - ID of the role to delete

**Example:**
```bash
kinde manage roles delete --role_id "role_123456789"
```

### `delete_scope`

Remove an API scope from a role, revoking access to specific API endpoints.

**Usage:**
```bash
kinde manage roles delete_scope [flags]
```

**Flags:**
- `--role_id string` - ID of the role to remove the scope from
- `--scope_id string` - ID of the API scope to remove from the role

**Example:**
```bash
kinde manage roles delete_scope --role_id "role_123456789" --scope_id "scope_987654321"
```

### `get`

Retrieve detailed information about a specific role.

**Usage:**
```bash
kinde manage roles get [flags]
```

**Flags:**
- `--role_id string` - ID of the role to retrieve

**Example:**
```bash
kinde manage roles get --role_id "role_123456789"
```

### `get_all`

List all roles in your Kinde environment with pagination support.

**Usage:**
```bash
kinde manage roles get_all [flags]
```

**Flags:**
- `--next_token string` - Token for pagination to get next page of results
- `--page_size int` - Number of roles to return per page

**Example:**
```bash
kinde manage roles get_all --page_size 20
```

### `get_permissions`

List all permissions assigned to a specific role.

**Usage:**
```bash
kinde manage roles get_permissions [flags]
```

**Flags:**
- `--next_token string` - Token for pagination to get next page of results
- `--page_size int` - Number of permissions to return per page
- `--role_id string` - ID of the role to get permissions for

**Example:**
```bash
kinde manage roles get_permissions --role_id "role_123456789" --page_size 50
```

### `get_scopes`

List all API scopes assigned to a specific role.

**Usage:**
```bash
kinde manage roles get_scopes [flags]
```

**Flags:**
- `--role_id string` - ID of the role to get scopes for

**Example:**
```bash
kinde manage roles get_scopes --role_id "role_123456789"
```

### `remove_permission`

Remove a permission from a role, revoking that specific access right.

**Usage:**
```bash
kinde manage roles remove_permission [flags]
```

**Flags:**
- `--permission_id string` - ID of the permission to remove from the role
- `--role_id string` - ID of the role to remove the permission from

**Example:**
```bash
kinde manage roles remove_permission --role_id "role_123456789" --permission_id "perm_987654321"
```

### `update`

Update an existing role's name, key, description, or default role status.

**Usage:**
```bash
kinde manage roles update [flags]
```

**Flags:**
- `--description string` - New description for the role
- `--is_default_role` - Update the default role status
- `--key string` - New key identifier for the role
- `--name string` - New display name for the role
- `--role_id string` - ID of the role to update

**Example:**
```bash
kinde manage roles update --role_id "role_123456789" --name "Senior Administrator" --description "Updated role description"
```

### `update_permissions`

Update the permissions assigned to a role. This command allows you to modify the complete set of permissions for a role.

**Usage:**
```bash
kinde manage roles update_permissions [flags]
```

**Flags:**
- `--role_id string` - ID of the role to update permissions for

**Example:**
```bash
kinde manage roles update_permissions --role_id "role_123456789"
```

## Common Usage Patterns

### Creating a Role System

When setting up a new application, you'll typically create a hierarchy of roles:

```bash
# Create basic roles for different access levels
kinde manage roles create --name "Administrator" --key "admin" --description "Full system access" --is_default_role
kinde manage roles create --name "Manager" --key "manager" --description "Management-level access"
kinde manage roles create --name "User" --key "user" --description "Standard user access"
kinde manage roles create --name "Read Only" --key "readonly" --description "Read-only access"

# Create specialized roles
kinde manage roles create --name "Support Agent" --key "support" --description "Customer support access"
kinde manage roles create --name "Developer" --key "developer" --description "Development and testing access"
```

### Managing Role Permissions

Assigning and managing permissions for roles:

```bash
# First, get available permissions
kinde manage permissions get_all

# Get current permissions for a role
kinde manage roles get_permissions --role_id "role_123456789"

# Add permissions to a role (using update_permissions)
kinde manage roles update_permissions --role_id "role_123456789"

# Remove a specific permission from a role
kinde manage roles remove_permission --role_id "role_123456789" --permission_id "perm_987654321"
```

### Managing Role Scopes

Managing API scopes for roles:

```bash
# Get available API scopes
kinde manage apis get_scopes --api_id "api_123"

# Get current scopes for a role
kinde manage roles get_scopes --role_id "role_123456789"

# Add a scope to a role
kinde manage roles add_scope --role_id "role_123456789" --scope_id "scope_987654321"

# Remove a scope from a role
kinde manage roles delete_scope --role_id "role_123456789" --scope_id "scope_987654321"
```

## Role Design Principles

### Role Hierarchy

Design roles with a clear hierarchy and purpose:

1. **Administrator** - Full system access
2. **Manager** - Department or team management
3. **User** - Standard operational access
4. **Read Only** - View-only access
5. **Guest** - Limited access for external users

### Permission Organization

Organize permissions logically within roles:

```bash
# Administrator role - all permissions
kinde manage roles update_permissions --role_id "admin_role_id"

# Manager role - management permissions
kinde manage roles update_permissions --role_id "manager_role_id"

# User role - operational permissions
kinde manage roles update_permissions --role_id "user_role_id"

# Read-only role - view permissions only
kinde manage roles update_permissions --role_id "readonly_role_id"
```

### Default Roles

Use default roles for new user onboarding:

```bash
# Set a role as default for new users
kinde manage roles create --name "Standard User" --key "standard_user" --description "Default role for new users" --is_default_role

# Update an existing role to be default
kinde manage roles update --role_id "role_123456789" --is_default_role
```

## Best Practices

### Role Design and Naming

1. **Use Consistent Naming Conventions**
   - Use clear, descriptive role names
   - Follow a consistent naming pattern (e.g., "Admin", "Manager", "User")
   - Use lowercase keys with underscores (e.g., "admin", "user_manager")

2. **Plan Role Hierarchy**
   - Design roles that can be combined effectively
   - Consider role inheritance and relationships
   - Avoid overly granular roles that are hard to manage

3. **Document Role Purposes**
   - Always provide clear descriptions for roles
   - Document what each role is intended for
   - Keep descriptions up to date as requirements change

### Permission Management

1. **Principle of Least Privilege**
   - Grant only the minimum permissions necessary
   - Regularly review and audit role permissions
   - Remove unused or excessive permissions

2. **Permission Grouping**
   - Group related permissions together
   - Use consistent permission naming patterns
   - Consider permission dependencies and prerequisites

3. **Regular Audits**
   - Periodically review all role permissions
   - Check for unused or obsolete permissions
   - Ensure permissions align with current business requirements

### Role Lifecycle Management

1. **Role Creation**
   - Create roles with clear purposes and boundaries
   - Test roles with sample users before full deployment
   - Document role creation and assignment processes

2. **Role Updates**
   - Update roles carefully to avoid breaking existing access
   - Communicate role changes to affected users
   - Test role changes in non-production environments

3. **Role Deletion**
   - Ensure no users are assigned to roles before deletion
   - Document role deletion and migration processes
   - Consider role deprecation before deletion

## Troubleshooting

### Common Issues

**Role Creation Fails**
- Ensure the role key is unique and follows naming conventions
- Check that required fields (name, key) are provided
- Verify you have the necessary permissions to create roles

**Permission Assignment Issues**
- Verify permission IDs exist and are valid
- Check that permissions are compatible with the role
- Ensure you have the necessary permissions to modify roles

**Role Not Found**
- Double-check the role ID is correct
- Ensure the role exists in your current environment
- Verify you're using the correct domain/environment

**Scope Assignment Issues**
- Verify scope IDs exist and are valid
- Check that scopes are compatible with the role
- Ensure API scopes are properly configured

### Getting Help

For additional help with roles:

```bash
# Get help for the roles command group
kinde manage roles --help

# Get help for a specific command
kinde manage roles create --help
kinde manage roles update_permissions --help
```

## Related Commands

- **Permissions Management** (`kinde manage permissions`) - Create and manage permissions
- **Users Management** (`kinde manage users`) - Assign roles to users
- **Organizations Management** (`kinde manage organizations`) - Manage organization-specific roles
- **APIs Management** (`kinde manage apis`) - Manage API scopes for roles

### Workflow Integration

Roles are typically used in conjunction with:

1. **Permissions** - Define what actions roles can perform
2. **Users** - Assign roles to specific users
3. **Organizations** - Create organization-specific roles
4. **APIs** - Grant API access through scopes

Example workflow:
```bash
# 1. Create permissions
kinde manage permissions create --name "Read Users" --key "users:read" --description "Read user information"
kinde manage permissions create --name "Create Users" --key "users:create" --description "Create new users"

# 2. Create a role
kinde manage roles create --name "User Manager" --key "user_manager" --description "Manage user accounts"

# 3. Assign permissions to the role
kinde manage roles update_permissions --role_id "role_123" --permission_ids "perm_456,perm_789"

# 4. Assign role to user
kinde manage users update --user_id "user_123" --roles "role_123"

# 5. Add API scopes to the role
kinde manage roles add_scope --role_id "role_123" --scope_id "scope_456"
```

### Integration with Applications

Roles enable applications to:
- Implement role-based access control (RBAC)
- Provide different user experiences based on roles
- Enforce security policies at the application level
- Support multi-tenant access patterns
- Integrate with external identity providers
