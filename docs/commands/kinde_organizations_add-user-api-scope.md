## kinde organizations add-user-api-scope

Add scope to organization user api

### Synopsis

Add a scope to an organization user api.

Required scope: create:organization_user_api_scopes

API: POST /api/v1/organizations/{org_code}/users/{user_id}/apis/{api_id}/scopes/{scope_id}

```
kinde organizations add-user-api-scope [flags]
```

### Options

```
      --apiid string      API ID (required)
  -h, --help              help for add-user-api-scope
      --org-code string   The identifier for the organization (required)
      --scope-id string   Scope ID (required)
      --user-id string    User ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

