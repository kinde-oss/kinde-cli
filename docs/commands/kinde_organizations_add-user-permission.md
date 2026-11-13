## kinde organizations add-user-permission

Add Organization User Permission

### Synopsis

Add permission to an organization user.

Required scope: create:organization_user_permissions

API: POST /api/v1/organizations/{org_code}/users/{user_id}/permissions

```
kinde organizations add-user-permission [flags]
```

### Options

```
      --body string            Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                   help for add-user-permission
      --org-code string        The organization's code (required)
      --permission-id string   The permission id
      --user-id string         The user's id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

