## kinde permissions create

Create Permission

### Synopsis

Create a new permission.

Required scope: create:permissions

API: POST /api/v1/permissions

```
kinde permissions create [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --description string   The permission's description
  -h, --help                 help for create
      --key string           The permission identifier to use in code
      --name string          The permission's name
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde permissions](kinde_permissions.md)	 - Manage permissions

