## kinde users create

Create user

### Synopsis

Creates a user record and optionally zero or more identities for the user. An
example identity could be the email address of the user.

Users cannot be created in an organization managed by directory sync
(`ORGANIZATION_MANAGED_BY_DIRECTORY_SYNC`); user and role changes for those
organizations must be made in the upstream identity provider. If
`organization_code` is omitted, this also applies to the default registration
organization.

Required scope: create:users

API: POST /api/v1/user

```
kinde users create [flags]
```

### Options

```
      --body string                  Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                         help for create
      --identities string            Array of identities to assign to the created user (JSON value)
      --organization-code string     The unique code associated with the organization you want the user to join
      --profile.family-name string   User's last name
      --profile.given-name string    User's first name
      --profile.picture string       The user's profile picture
      --provided-id string           An external id to reference the user
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

