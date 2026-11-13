## kinde roles list-permissions

Get role permissions

### Synopsis

Get permissions for a role.

Required scope: read:role_permissions

API: GET /api/v1/roles/{role_id}/permissions

```
kinde roles list-permissions [flags]
```

### Options

```
  -h, --help                help for list-permissions
      --next-token string   A string to get the next page of results if there are more results
      --page-size int       Number of results per page. Defaults to 10 if parameter not sent
      --role-id string      The role's public id (required)
      --sort string         Field and order to sort the result by (one of: name_asc, name_desc, id_asc, id_desc)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

