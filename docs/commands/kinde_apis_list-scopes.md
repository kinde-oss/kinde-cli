## kinde apis list-scopes

Get API scopes

### Synopsis

Retrieve API scopes by API ID.

Required scope: read:api_scopes

API: GET /api/v1/apis/{api_id}/scopes

```
kinde apis list-scopes [flags]
```

### Options

```
      --apiid string   API ID (required)
  -h, --help           help for list-scopes
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

