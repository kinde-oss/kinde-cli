## kinde api-keys create

Create API key

### Synopsis

Create a new API key.

Required scope: create:api_keys

API: POST /api/v1/api_keys

```
kinde api-keys create [flags]
```

### Options

```
      --api-id string       The ID of the API this key is associated with (required)
      --body string         Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                help for create
      --name string         The name of the API key (required)
      --org-code string     The organization code to associate with this API key (for organization-level keys)
      --scope-ids strings   Array of scope IDs to associate with this API key (comma-separated list)
      --type string         The entity type that will use this API key (one of: user, organization, environment; required)
      --user-id string      The ID of the user to associate with this API key (for user-level keys)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde api-keys](kinde_api-keys.md)	 - Manage API keys

