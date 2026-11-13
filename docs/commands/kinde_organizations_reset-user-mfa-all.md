## kinde organizations reset-user-mfa-all

Reset all organization MFA for a user

### Synopsis

Reset all organization MFA factors for a user.

Required scope: delete:organization_user_mfa

API: DELETE /api/v1/organizations/{org_code}/users/{user_id}/mfa

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations reset-user-mfa-all [flags]
```

### Options

```
  -h, --help              help for reset-user-mfa-all
      --org-code string   The identifier for the organization (required)
      --user-id string    The identifier for the user (required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

