## kinde roles list-users

List role users

### Synopsis

Get users that have a given role, across all organizations. Each user entry
includes the organization codes where they hold that role.

Required scope: read:organization_user_roles

API: GET /api/v1/roles/{role_id}/users

```
kinde roles list-users [flags]
```

### Options

```
  -h, --help                help for list-users
      --next-token string   A string to get the next page of results if there are more results
      --page-size int       Number of results per page. Defaults to 10 if parameter not sent
      --role-id string      The role's public id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

