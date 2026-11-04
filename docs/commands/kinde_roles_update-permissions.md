## kinde roles update-permissions

Update role permissions

### Synopsis

Update role permissions.

Required scope: update:role_permissions

API: PATCH /api/v1/roles/{role_id}/permissions

```
kinde roles update-permissions [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                 help for update-permissions
      --permissions string   Permissions to add or remove from the role (JSON value; required)
      --role-id string       The identifier for the role (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

