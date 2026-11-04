## kinde organizations delete-feature-flag-override

Delete Organization Feature Flag Override

### Synopsis

Delete organization feature flag override.

Required scope: delete:organization_feature_flags

API: DELETE /api/v1/organizations/{org_code}/feature_flags/{feature_flag_key}

```
kinde organizations delete-feature-flag-override [flags]
```

### Options

```
      --feature-flag-key string   The identifier for the feature flag (required)
  -h, --help                      help for delete-feature-flag-override
      --org-code string           The identifier for the organization (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

