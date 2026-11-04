## kinde users search

Search users

### Synopsis

Search for users based on the provided query string. Set query to '*' to filter
by other parameters only. The number of records to return at a time can be
controlled using the `page_size` query string parameter.

Required scope: read:users

API: GET /api/v1/search/users

```
kinde users search [flags]
```

### Options

```
      --api-scopes string       Search the users by api scopes
      --ending-before string    The ID of the user to end before
      --expand string           Additional data to include in the response
  -h, --help                    help for search
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --properties string       JSON value
      --query string            Search the users by email or name. Use '*' to search all
      --starting-after string   The ID of the user to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

