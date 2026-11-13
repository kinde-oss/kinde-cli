## kinde organizations list-role-users

List organization role users

### Synopsis

Get users that have a given role within a specific organization.

Required scope: read:organization_user_roles

API: GET /api/v1/organizations/{org_code}/roles/{role_id}/users

```
kinde organizations list-role-users [flags]
```

### Options

```
  -h, --help                help for list-role-users
      --next-token string   A string to get the next page of results if there are more results
      --org-code string     The organization's code (required)
      --page-size int       Number of results per page. Defaults to 10 if parameter not sent
      --role-id string      The role's public id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

