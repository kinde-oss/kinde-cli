## kinde organizations replace-mfa

Replace Organization MFA Configuration

### Synopsis

Replace Organization MFA Configuration.

Required scope: update:organization_mfa

API: PUT /api/v1/organizations/{org_code}/mfa

```
kinde organizations replace-mfa [flags]
```

### Options

```
      --body string                 Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --enabled-factors strings     The MFA methods to enable (comma-separated list; one of: mfa:email, mfa:sms, mfa:authenticator_app; required)
  -h, --help                        help for replace-mfa
      --is-recovery-codes-enabled   Determines whether recovery codes are shown to users during MFA setup for this specific organization
      --org-code string             The identifier for the organization (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

