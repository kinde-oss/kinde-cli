## kinde apis create-scope

Create API scope

### Synopsis

Create a new API scope.

Required scope: create:api_scopes

API: POST /api/v1/apis/{api_id}/scopes

```
kinde apis create-scope [flags]
```

### Options

```
      --apiid string         API ID (required)
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --description string   Description of the api scope purpose
  -h, --help                 help for create-scope
      --key string           The key reference for the scope (1-64 characters, no white space) (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

