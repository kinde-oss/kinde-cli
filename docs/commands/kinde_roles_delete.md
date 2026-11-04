## kinde roles delete

Delete role

### Synopsis

Delete role

Required scope: delete:roles

API: DELETE /api/v1/roles/{role_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde roles delete [flags]
```

### Options

```
  -h, --help             help for delete
      --role-id string   The identifier for the role (required)
  -y, --yes              Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

