## kinde mfa replace

Replace MFA Configuration

### Synopsis

Replace MFA Configuration.

Required scope: update:mfa

API: PUT /api/v1/mfa

```
kinde mfa replace [flags]
```

### Options

```
      --body string                 Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --enabled-factors strings     The MFA methods to enable (comma-separated list; one of: mfa:email, mfa:sms, mfa:authenticator_app; required)
  -h, --help                        help for replace
      --is-recovery-codes-enabled   Determines whether recovery codes are shown to users during MFA setup for the environment
      --policy string               Specifies whether MFA is required, optional, or not enforced (one of: required, off, optional; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde mfa](kinde_mfa.md)	 - Manage the environment's MFA configuration

