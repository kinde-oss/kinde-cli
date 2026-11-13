## kinde organizations update-users

Update Organization Users

### Synopsis

Update users that belong to an organization.

`roles` replaces the user's roles. Every key in it must be a business-wide role
key, or the key of one of this organization's own roles while
organization-managed roles are available on the business's plan. Any other key
fails the whole request with `INVALID_ROLES` instead of being skipped. Below
that plan the user's organization-managed role assignments are kept but not
issued in tokens; they cannot be named in `roles` until the plan is upgraded,
and replacing `roles` does not remove them.

Users in organizations managed by directory sync cannot be updated
(`ORGANIZATION_MANAGED_BY_DIRECTORY_SYNC`); user and role changes for those
organizations must be made in the upstream identity provider.

Required scope: update:organization_users

API: PATCH /api/v1/organizations/{org_code}/users

```
kinde organizations update-users [flags]
```

### Options

```
      --body string       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help              help for update-users
      --org-code string   The organization's code (required)
      --users string      Users to add, update or remove from the organization (JSON value)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

