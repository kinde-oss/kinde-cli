## kinde organizations list-user-roles

List Organization User Roles

### Synopsis

Get roles for an organization user.

Required scope: read:organization_user_roles

API: GET /api/v1/organizations/{org_code}/users/{user_id}/roles

```
kinde organizations list-user-roles [flags]
```

### Options

```
  -h, --help              help for list-user-roles
      --org-code string   The organization's code (required)
      --user-id string    The user's id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

