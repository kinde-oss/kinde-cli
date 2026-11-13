## kinde organizations get-invite

Get organization invite

### Synopsis

Get details of a specific invitation by its code.

Required scope: read:organization_invites

API: GET /api/v1/organization/{org_code}/invites/{invite_code}

```
kinde organizations get-invite [flags]
```

### Options

```
  -h, --help                 help for get-invite
      --invite-code string   The invitation's code (required)
      --org-code string      The organization's code (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

