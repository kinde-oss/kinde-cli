## kinde roles remove-permission

Remove role permission

### Synopsis

Remove a permission from a role.

Required scope: delete:role_permissions

API: DELETE /api/v1/roles/{role_id}/permissions/{permission_id}

```
kinde roles remove-permission [flags]
```

### Options

```
  -h, --help                   help for remove-permission
      --permission-id string   The permission's public id (required)
      --role-id string         The role's public id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

