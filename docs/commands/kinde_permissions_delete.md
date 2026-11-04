## kinde permissions delete

Delete Permission

### Synopsis

Delete permission

Required scope: delete:permissions

API: DELETE /api/v1/permissions/{permission_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde permissions delete [flags]
```

### Options

```
  -h, --help                   help for delete
      --permission-id string   The identifier for the permission (required)
  -y, --yes                    Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde permissions](kinde_permissions.md)	 - Manage permissions

