## kinde organizations list

Get organizations

### Synopsis

Get a list of organizations.

Required scope: read:organizations

API: GET /api/v1/organizations

```
kinde organizations list [flags]
```

### Options

```
  -h, --help                help for list
      --next-token string   A string to get the next page of results if there are more results
      --page-size int       Number of results per page. Defaults to 10 if parameter not sent
      --sort string         Field and order to sort the result by (one of: name_asc, name_desc, email_asc, email_desc)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

