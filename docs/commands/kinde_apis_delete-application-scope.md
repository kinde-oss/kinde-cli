## kinde apis delete-application-scope

Delete API application scope

### Synopsis

Delete an API application scope you previously created.

Required scope: delete:api_application_scopes

API: DELETE /api/v1/apis/{api_id}/applications/{application_id}/scopes/{scope_id}

```
kinde apis delete-application-scope [flags]
```

### Options

```
      --apiid string            API ID (required)
      --application-id string   Application ID (required)
  -h, --help                    help for delete-application-scope
      --scope-id string         Scope ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

