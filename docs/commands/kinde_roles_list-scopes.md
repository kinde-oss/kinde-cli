## kinde roles list-scopes

Get role scopes

### Synopsis

Get scopes for a role.

Required scope: read:role_permissions

API: GET /api/v1/roles/{role_id}/scopes

```
kinde roles list-scopes [flags]
```

### Options

```
  -h, --help             help for list-scopes
      --role-id string   The role id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

