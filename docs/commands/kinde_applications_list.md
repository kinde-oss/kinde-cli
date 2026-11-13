## kinde applications list

Get applications

### Synopsis

Get a list of applications / clients.

Required scope: read:applications

API: GET /api/v1/applications

```
kinde applications list [flags]
```

### Options

```
  -h, --help                help for list
      --next-token string   A string to get the next page of results if there are more results
      --page-size int       Number of results per page. Defaults to 10 if parameter not sent
      --sort string         Field and order to sort the result by (one of: name_asc, name_desc)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

