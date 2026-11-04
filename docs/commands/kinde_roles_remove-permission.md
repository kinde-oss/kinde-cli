## kinde roles remove-permission

Remove role permission

### Synopsis

Remove a permission from a role.

Required scope: delete:role_permissions

API: DELETE /api/v1/roles/{role_id}/permissions/{permission_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde roles remove-permission [flags]
```

### Options

```
  -h, --help                   help for remove-permission
      --permission-id string   The permission's public id (required)
      --role-id string         The role's public id (required)
  -y, --yes                    Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

