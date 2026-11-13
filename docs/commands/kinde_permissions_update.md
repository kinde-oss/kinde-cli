## kinde permissions update

Update Permission

### Synopsis

Update permission

Required scope: update:permissions

API: PATCH /api/v1/permissions/{permission_id}

```
kinde permissions update [flags]
```

### Options

```
      --body string            Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --description string     The permission's description
  -h, --help                   help for update
      --key string             The permission identifier to use in code
      --name string            The permission's name
      --permission-id string   The identifier for the permission (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde permissions](kinde_permissions.md)	 - Manage permissions

