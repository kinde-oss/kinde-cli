## kinde apis delete-scope

Delete API scope

### Synopsis

Delete an API scope you previously created.

Required scope: delete:api_scopes

API: DELETE /api/v1/apis/{api_id}/scopes/{scope_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde apis delete-scope [flags]
```

### Options

```
      --apiid string      API ID (required)
  -h, --help              help for delete-scope
      --scope-id string   Scope ID (required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

