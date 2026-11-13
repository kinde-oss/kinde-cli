## kinde organizations update-passkey

Update organization passkey settings

### Synopsis

Update passkey settings for an organization. Set
`is_override_environment_passkey_settings` to `false` to revert to the
environment default without providing a policy.

Required scope: update:organization_passkey

API: PUT /api/v1/organizations/{org_code}/passkey

```
kinde organizations update-passkey [flags]
```

### Options

```
      --body string                                Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                                       help for update-passkey
      --is-override-environment-passkey-settings   Whether to use a custom passkey policy for this organization. Set to false to revert to the environment default
      --org-code string                            The organization's code (required)
      --policy string                              Passkey policy when overriding the environment default (one of: off, optional, mandatory)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

