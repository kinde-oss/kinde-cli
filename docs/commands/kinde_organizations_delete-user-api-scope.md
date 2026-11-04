## kinde organizations delete-user-api-scope

Delete scope from organization user API

### Synopsis

Delete a scope from an organization user api you previously created.

Required scope: delete:organization_user_api_scopes

API: DELETE /api/v1/organizations/{org_code}/users/{user_id}/apis/{api_id}/scopes/{scope_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations delete-user-api-scope [flags]
```

### Options

```
      --apiid string      API ID (required)
  -h, --help              help for delete-user-api-scope
      --org-code string   The identifier for the organization (required)
      --scope-id string   Scope ID (required)
      --user-id string    User ID (required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

