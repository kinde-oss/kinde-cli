## kinde organizations count-role-users

Count organization role users

### Synopsis

Get the number of users that have a given role within a specific organization.

Required scope: read:organization_user_roles

API: GET /api/v1/organizations/{org_code}/roles/{role_id}/users/count

```
kinde organizations count-role-users [flags]
```

### Options

```
  -h, --help              help for count-role-users
      --org-code string   The organization's code (required)
      --role-id string    The role's public id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

