## kinde feature-flags delete

Delete Feature Flag

### Synopsis

Delete feature flag

Required scope: delete:feature_flags

API: DELETE /api/v1/feature_flags/{feature_flag_key}

```
kinde feature-flags delete [flags]
```

### Options

```
      --feature-flag-key string   The identifier for the feature flag (required)
  -h, --help                      help for delete
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde feature-flags](kinde_feature-flags.md)	 - Manage feature flags

