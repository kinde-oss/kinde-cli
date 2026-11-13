## kinde environment delete-feature-flag-overrides

Delete Environment Feature Flag Overrides

### Synopsis

Delete all environment feature flag overrides.

Required scope: delete:environment_feature_flags

API: DELETE /api/v1/environment/feature_flags

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde environment delete-feature-flag-overrides [flags]
```

### Options

```
  -h, --help   help for delete-feature-flag-overrides
  -y, --yes    Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment](kinde_environment.md)	 - Manage the current environment

