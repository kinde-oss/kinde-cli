## kinde organizations add-users

Add Organization Users

### Synopsis

Add existing users to an organization.

Every key in `roles` must be a business-wide role key, or the key of one of this
organization's own roles while organization-managed roles are available on the
business's plan. Any other key fails the whole request with `INVALID_ROLES`
instead of being skipped.

Users cannot be added to organizations that are managed by directory sync
(`ORGANIZATION_MANAGED_BY_DIRECTORY_SYNC`); user and role changes for those
organizations must be made in the upstream identity provider.

Required scope: create:organization_users

API: POST /api/v1/organizations/{org_code}/users

```
kinde organizations add-users [flags]
```

### Options

```
      --body string       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help              help for add-users
      --org-code string   The organization's code (required)
      --users string      Users to be added to the organization (JSON value)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

