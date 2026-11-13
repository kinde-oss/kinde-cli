## kinde users list

Get users

### Synopsis

The returned list can be sorted by full name or email address in ascending or
descending order. The number of records to return at a time can also be
controlled using the `page_size` query string parameter.

Required scope: read:users

API: GET /api/v1/users

```
kinde users list [flags]
```

### Options

```
      --active-since string   Filter the results to only include users who have been active since this date. Date should be in ISO 8601 format (RFC 3339 timestamp, e.g. 2026-01-31T09:00:00Z)
      --email string          Filter the results by email address. The query string should be comma separated and url encoded
      --expand string         Additional data to include in the response. One or more of (comma-separated): "organizations", "identities", "billing"
      --has-organization      Filter the results by if the user has at least one organization assigned
  -h, --help                  help for list
      --next-token string     A string to get the next page of results if there are more results
      --page-size int         Number of results per page. Defaults to 10 if parameter not sent
      --phone string          Filter the results by phone. The query string should be comma separated and url encoded
      --user-id string        Filter the results by User ID. The query string should be comma separated and url encoded
      --username string       Filter the results by username. The query string should be comma separated and url encoded
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

