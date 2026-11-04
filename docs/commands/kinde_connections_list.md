## kinde connections list

Get connections

### Synopsis

Returns a list of authentication connections. Optionally you can filter this by
a home realm domain.

Required scope: read:connections

API: GET /api/v1/connections

```
kinde connections list [flags]
```

### Options

```
      --ending-before string       The ID of the connection to end before
  -h, --help                       help for list
      --home-realm-domain string   Filter the results by the home realm domain
      --page-size int              Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string      The ID of the connection to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde connections](kinde_connections.md)	 - Manage authentication connections

