## kinde organizations delete-feature-flag-overrides

Delete Organization Feature Flag Overrides

### Synopsis

Delete all organization feature flag overrides.

Required scope: delete:organization_feature_flags

API: DELETE /api/v1/organizations/{org_code}/feature_flags

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations delete-feature-flag-overrides [flags]
```

### Options

```
  -h, --help              help for delete-feature-flag-overrides
      --org-code string   The identifier for the organization (required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

