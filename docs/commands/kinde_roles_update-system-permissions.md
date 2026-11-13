## kinde roles update-system-permissions

Update role system permissions

### Synopsis

Update role system permissions.

System permissions control what organization users can do in the self-serve
portal.

Required scope: update:role_system_permissions

API: PATCH /api/v1/roles/{role_id}/system_permissions

```
kinde roles update-system-permissions [flags]
```

### Options

```
      --body string                 Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                        help for update-system-permissions
      --role-id string              The identifier for the role (required)
      --system-permissions string   System permissions to add or remove from the role (JSON value; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

