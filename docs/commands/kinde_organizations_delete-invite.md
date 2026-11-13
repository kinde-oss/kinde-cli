## kinde organizations delete-invite

Delete organization invite

### Synopsis

Revoke (delete) an invitation. This will mark the invitation as revoked and
prevent it from being accepted.

Required scope: delete:organization_invites

API: DELETE /api/v1/organization/{org_code}/invites/{invite_code}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations delete-invite [flags]
```

### Options

```
  -h, --help                 help for delete-invite
      --invite-code string   The invitation's code (required)
      --org-code string      The organization's code (required)
  -y, --yes                  Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

