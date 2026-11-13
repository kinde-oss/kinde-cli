## kinde organizations get-passkey

Get organization passkey settings

### Synopsis

Retrieve passkey settings for an organization, including whether the
organization overrides the environment default.

Required scope: read:organization_passkey

API: GET /api/v1/organizations/{org_code}/passkey

```
kinde organizations get-passkey [flags]
```

### Options

```
  -h, --help              help for get-passkey
      --org-code string   The organization's code (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

