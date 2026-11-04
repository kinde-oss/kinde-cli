## kinde organizations list-feature-flags

List Organization Feature Flags

### Synopsis

Get all organization feature flags.

Required scope: read:organization_feature_flags

API: GET /api/v1/organizations/{org_code}/feature_flags

```
kinde organizations list-feature-flags [flags]
```

### Options

```
  -h, --help              help for list-feature-flags
      --org-code string   The identifier for the organization (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

