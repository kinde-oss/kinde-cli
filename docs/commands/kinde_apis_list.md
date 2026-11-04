## kinde apis list

Get APIs

### Synopsis

Returns a list of your APIs. The APIs are returned sorted by name.

Required scope: read:apis

API: GET /api/v1/apis

```
kinde apis list [flags]
```

### Options

```
      --expand string   Additional data to include in the response. Allowed value: "scopes" (one of: scopes)
  -h, --help            help for list
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

