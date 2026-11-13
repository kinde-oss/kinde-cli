## kinde environment-variables delete

Delete environment variable

### Synopsis

Delete an environment variable you previously created. This feature is in beta
and admin UI is not yet available.

Required scope: delete:environment_variables

API: DELETE /api/v1/environment_variables/{variable_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde environment-variables delete [flags]
```

### Options

```
  -h, --help                 help for delete
      --variable-id string   The environment variable's ID (required)
  -y, --yes                  Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment-variables](kinde_environment-variables.md)	 - Manage environment variables

