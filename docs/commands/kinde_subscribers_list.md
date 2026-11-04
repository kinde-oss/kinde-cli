## kinde subscribers list

List Subscribers

### Synopsis

The returned list can be sorted by full name or email address in ascending or
descending order. The number of records to return at a time can also be
controlled using the `page_size` query string parameter.

Required scope: read:subscribers

API: GET /api/v1/subscribers

```
kinde subscribers list [flags]
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

* [kinde subscribers](kinde_subscribers.md)	 - Manage subscribers

