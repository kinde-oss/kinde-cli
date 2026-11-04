## kinde apis delete-scope

Delete API scope

### Synopsis

Delete an API scope you previously created.

Required scope: delete:api_scopes

API: DELETE /api/v1/apis/{api_id}/scopes/{scope_id}

```
kinde apis delete-scope [flags]
```

### Options

```
      --apiid string      API ID (required)
  -h, --help              help for delete-scope
      --scope-id string   Scope ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

