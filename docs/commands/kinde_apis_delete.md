## kinde apis delete

Delete API

### Synopsis

Delete an API you previously created.

Required scope: delete:apis

API: DELETE /api/v1/apis/{api_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde apis delete [flags]
```

### Options

```
      --apiid string   The API's ID (required)
  -h, --help           help for delete
  -y, --yes            Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

