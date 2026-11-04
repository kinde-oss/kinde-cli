## kinde api-keys delete

Delete API key

### Synopsis

Delete an API key.

Required scope: delete:api_keys

API: DELETE /api/v1/api_keys/{key_id}

```
kinde api-keys delete [flags]
```

### Options

```
  -h, --help            help for delete
      --key-id string   The ID of the API key (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde api-keys](kinde_api-keys.md)	 - Manage API keys

