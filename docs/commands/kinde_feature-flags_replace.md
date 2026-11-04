## kinde feature-flags replace

Replace Feature Flag

### Synopsis

Update feature flag.

While the flag has overrides in this environment, a change to its type or
override level is refused with `DELETE_NOT_ALLOWED`. Remove the overrides first.

Required scope: update:feature_flags

API: PUT /api/v1/feature_flags/{feature_flag_key}

```
kinde feature-flags replace [flags]
```

### Options

```
      --allow-override-level string   Allow the flag to be overridden at a different level (one of: env, org; required)
      --default-value string          Default value for the flag used by environments and organizations (required)
      --description string            Description of the flag purpose (required)
      --feature-flag-key string       The key identifier for the feature flag (required)
  -h, --help                          help for replace
      --name string                   The name of the flag (required)
      --type string                   The variable type (one of: str, int, bool; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde feature-flags](kinde_feature-flags.md)	 - Manage feature flags

