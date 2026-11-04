## kinde organizations reset-user-mfa

Reset specific organization MFA for a user

### Synopsis

Reset a specific organization MFA factor for a user.

Required scope: delete:organization_user_mfa

API: DELETE /api/v1/organizations/{org_code}/users/{user_id}/mfa/{factor_id}

```
kinde organizations reset-user-mfa [flags]
```

### Options

```
      --factor-id string   The identifier for the MFA factor (required)
  -h, --help               help for reset-user-mfa
      --org-code string    The identifier for the organization (required)
      --user-id string     The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

