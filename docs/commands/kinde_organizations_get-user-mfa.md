## kinde organizations get-user-mfa

Get an organization user's MFA configuration

### Synopsis

Get an organization user’s MFA configuration.

Required scope: read:organization_user_mfa

API: GET /api/v1/organizations/{org_code}/users/{user_id}/mfa

```
kinde organizations get-user-mfa [flags]
```

### Options

```
  -h, --help              help for get-user-mfa
      --org-code string   The identifier for the organization (required)
      --user-id string    The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

