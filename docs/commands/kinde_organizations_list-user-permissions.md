## kinde organizations list-user-permissions

List Organization User Permissions

### Synopsis

Get permissions for an organization user.

Required scope: read:organization_user_permissions

API: GET /api/v1/organizations/{org_code}/users/{user_id}/permissions

```
kinde organizations list-user-permissions [flags]
```

### Options

```
      --expand string     Additional data to include in the response. Allowed value: "roles"
  -h, --help              help for list-user-permissions
      --org-code string   The organization's code (required)
      --user-id string    The user's id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

