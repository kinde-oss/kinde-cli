## kinde organizations update-sessions

Update organization session configuration

### Synopsis

Update the organization's session configuration.

Required scope: update:organizations

API: PATCH /api/v1/organizations/{org_code}/sessions

```
kinde organizations update-sessions [flags]
```

### Options

```
      --authenticated-session-lifetime int          Authenticated session lifetime in seconds
      --body string                                 Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                                        help for update-sessions
      --is-use-org-authenticated-session-lifetime   Whether to apply the organization's authenticated session lifetime override
      --is-use-org-sso-session-policy               Whether to use the organization's SSO session policy override
      --org-code string                             The organization's code (required)
      --sso-session-persistence-mode string         Determines if the session should be persistent or not (one of: persistent, non_persistent)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

