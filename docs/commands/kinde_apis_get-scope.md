## kinde apis get-scope

Get API scope

### Synopsis

Retrieve API scope by API ID.

Required scope: read:api_scopes

API: GET /api/v1/apis/{api_id}/scopes/{scope_id}

```
kinde apis get-scope [flags]
```

### Options

```
      --apiid string      API ID (required)
  -h, --help              help for get-scope
      --scope-id string   Scope ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

