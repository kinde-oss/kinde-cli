## kinde organizations add-user-role

Add Organization User Role

### Synopsis

Add role to an organization user.

Roles cannot be added to users in organizations managed by directory sync
(`ORGANIZATION_MANAGED_BY_DIRECTORY_SYNC`); user and role changes for those
organizations must be made in the upstream identity provider.

Required scope: create:organization_user_roles

API: POST /api/v1/organizations/{org_code}/users/{user_id}/roles

```
kinde organizations add-user-role [flags]
```

### Options

```
      --body string       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help              help for add-user-role
      --org-code string   The organization's code (required)
      --role-id string    The role id
      --user-id string    The user's id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

