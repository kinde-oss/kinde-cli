## kinde apis add-application-scope

Add scope to API application

### Synopsis

Add a scope to an API application.

Required scope: create:api_application_scopes

API: POST /api/v1/apis/{api_id}/applications/{application_id}/scopes/{scope_id}

```
kinde apis add-application-scope [flags]
```

### Options

```
      --apiid string            API ID (required)
      --application-id string   Application ID (required)
  -h, --help                    help for add-application-scope
      --scope-id string         Scope ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

