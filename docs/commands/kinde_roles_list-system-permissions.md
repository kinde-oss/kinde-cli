## kinde roles list-system-permissions

List system permissions

### Synopsis

The returned list can be sorted by system permission name or ID in ascending or
descending order. The number of records to return at a time can also be
controlled using the `page_size` query string parameter.

System permissions control what organization users can do in the self-serve
portal.

Required scope: read:system_permissions

API: GET /api/v1/system_permissions

```
kinde roles list-system-permissions [flags]
```

### Options

```
  -h, --help                help for list-system-permissions
      --next-token string   A string to get the next page of results if there are more results
      --page-size int       Number of results per page. Defaults to 10 if parameter not sent
      --sort string         Field and order to sort the result by (one of: name_asc, name_desc, id_asc, id_desc)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

