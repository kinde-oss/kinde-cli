## kinde api-keys verify

Verify API key

### Synopsis

Verify an API key (public endpoint, no authentication required).

API: POST /api/v1/api_keys/verify

```
kinde api-keys verify [flags]
```

### Options

```
      --api-key string   The API key to verify (required)
      --body string      Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help             help for verify
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde api-keys](kinde_api-keys.md)	 - Manage API keys

