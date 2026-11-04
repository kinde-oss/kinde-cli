## kinde environment update-feature-flag-override

Update Environment Feature Flag Override

### Synopsis

Update environment feature flag override.

Required scope: update:environment_feature_flags

API: PATCH /api/v1/environment/feature_flags/{feature_flag_key}

```
kinde environment update-feature-flag-override [flags]
```

### Options

```
      --body string               Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --feature-flag-key string   The identifier for the feature flag (required)
  -h, --help                      help for update-feature-flag-override
      --value string              The flag override value (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment](kinde_environment.md)	 - Manage the current environment

