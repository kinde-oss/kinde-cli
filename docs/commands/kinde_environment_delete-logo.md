## kinde environment delete-logo

Delete logo

### Synopsis

Delete environment logo

Required scope: update:environments

API: DELETE /api/v1/environment/logos/{type}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde environment delete-logo [flags]
```

### Options

```
  -h, --help          help for delete-logo
      --type string   The type of logo to delete (one of: dark, light; required)
  -y, --yes           Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment](kinde_environment.md)	 - Manage the current environment

