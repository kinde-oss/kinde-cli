## kinde api-keys rotate

Rotate API key

### Synopsis

Rotate an API key to generate a new key while maintaining the same permissions
and associations.

Required scope: update:api_keys

API: PUT /api/v1/api_keys/{key_id}

```
kinde api-keys rotate [flags]
```

### Options

```
  -h, --help            help for rotate
      --key-id string   The ID of the API key to rotate (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde api-keys](kinde_api-keys.md)	 - Manage API keys

