## kinde organizations update-feature-flag-override

Update Organization Feature Flag Override

### Synopsis

Update organization feature flag override.

Required scope: update:organization_feature_flags

API: PATCH /api/v1/organizations/{org_code}/feature_flags/{feature_flag_key}

```
kinde organizations update-feature-flag-override [flags]
```

### Options

```
      --feature-flag-key string   The identifier for the feature flag (required)
  -h, --help                      help for update-feature-flag-override
      --org-code string           The identifier for the organization (required)
      --value string              Override value (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

