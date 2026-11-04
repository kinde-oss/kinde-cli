## kinde environment delete-feature-flag-override

Delete Environment Feature Flag Override

### Synopsis

Delete environment feature flag override.

Required scope: delete:environment_feature_flags

API: DELETE /api/v1/environment/feature_flags/{feature_flag_key}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde environment delete-feature-flag-override [flags]
```

### Options

```
      --feature-flag-key string   The identifier for the feature flag (required)
  -h, --help                      help for delete-feature-flag-override
  -y, --yes                       Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment](kinde_environment.md)	 - Manage the current environment

