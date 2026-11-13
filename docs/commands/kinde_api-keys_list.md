## kinde api-keys list

Get API keys

### Synopsis

Returns a list of API keys.

Required scope: read:api_keys

API: GET /api/v1/api_keys

```
kinde api-keys list [flags]
```

### Options

```
  -h, --help                    help for list
      --key-type string         Filter by API key type (organization or user) (one of: organization, user)
      --org-code string         Filter by organization code to get API keys associated with a specific organization
      --page-size int           Number of results per page. Defaults to 50 if parameter not sent
      --starting-after string   The ID of the API key to start after
      --status string           Filter by API key status (active, inactive, revoked) (one of: active, inactive, revoked)
      --user-id string          Filter by user ID to get API keys associated with a specific user
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde api-keys](kinde_api-keys.md)	 - Manage API keys

