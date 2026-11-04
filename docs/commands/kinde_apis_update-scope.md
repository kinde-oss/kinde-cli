## kinde apis update-scope

Update API scope

### Synopsis

Update an API scope.

Required scope: update:api_scopes

API: PATCH /api/v1/apis/{api_id}/scopes/{scope_id}

```
kinde apis update-scope [flags]
```

### Options

```
      --apiid string         API ID (required)
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --description string   Description of the api scope purpose
  -h, --help                 help for update-scope
      --scope-id string      Scope ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

