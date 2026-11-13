## kinde organizations create-invite

Create organization invite

### Synopsis

Create a new invitation for an organization. An invitation email will be sent to
the provided email address if `send_email` is set to `true`.

Invitations cannot be created for organizations that are managed by directory
sync; user and role changes for those organizations must be made in the upstream
identity provider.

Roles that require an explicit assignment permission cannot be granted to an
invitee unless the caller (or the user the token represents) holds that
permission. On Kinde-hosted plans, roles outside `owner`/`admin` additionally
require the `extended_roles` entitlement.

Per-organization rate limits apply: a maximum number of invitations may be
created per rolling 24 hour window, and a maximum number of active
(non-accepted, non-revoked) invitations may exist at any time. Requests that
exceed either limit are rejected.

Required scope: create:organization_invites

API: POST /api/v1/organization/{org_code}/invites

```
kinde organizations create-invite [flags]
```

### Options

```
      --body string         Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --email string        The email address of the user to invite. Maximum 254 characters (required)
      --first-name string   The first name of the user to invite. Maximum 64 characters
  -h, --help                help for create-invite
      --last-name string    The last name of the user to invite. Maximum 64 characters
      --org-code string     The organization's code (required)
      --roles strings       Array of role keys to assign to the user (comma-separated list; required)
      --send-email          Whether to send an invitation email to the user. Defaults to false
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

