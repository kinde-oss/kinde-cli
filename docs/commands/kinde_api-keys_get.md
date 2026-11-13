## kinde api-keys get

Get API key

### Synopsis

Retrieve API key details by ID.

Required scope: read:api_keys

API: GET /api/v1/api_keys/{key_id}

```
kinde api-keys get [flags]
```

### Options

```
  -h, --help            help for get
      --key-id string   The ID of the API key (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde api-keys](kinde_api-keys.md)	 - Manage API keys

