## kinde api-keys delete

Delete API key

### Synopsis

Delete an API key.

Required scope: delete:api_keys

API: DELETE /api/v1/api_keys/{key_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde api-keys delete [flags]
```

### Options

```
  -h, --help            help for delete
      --key-id string   The ID of the API key (required)
  -y, --yes             Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde api-keys](kinde_api-keys.md)	 - Manage API keys

