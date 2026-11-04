## kinde roles add-scope

Add role scope

### Synopsis

Add scope to role.

Required scope: update:role_permissions

API: POST /api/v1/roles/{role_id}/scopes

```
kinde roles add-scope [flags]
```

### Options

```
      --body string       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help              help for add-scope
      --role-id string    The role id (required)
      --scope-id string   The scope identifier (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

