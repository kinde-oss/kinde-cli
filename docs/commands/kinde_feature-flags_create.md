## kinde feature-flags create

Create Feature Flag

### Synopsis

Create feature flag.

Required scope: create:feature_flags

API: POST /api/v1/feature_flags

```
kinde feature-flags create [flags]
```

### Options

```
      --allow-override-level string   Allow the flag to be overridden at a different level (one of: env, org, usr)
      --body string                   Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --default-value string          Default value for the flag used by environments and organizations (required)
      --description string            Description of the flag purpose
  -h, --help                          help for create
      --key string                    The flag identifier to use in code (required)
      --name string                   The name of the flag (required)
      --type string                   The variable type (one of: str, int, bool; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde feature-flags](kinde_feature-flags.md)	 - Manage feature flags

