## kinde organizations remove-user

Remove Organization User

### Synopsis

Remove user from an organization.

Users cannot be removed from organizations that are managed by directory sync
(`ORGANIZATION_MANAGED_BY_DIRECTORY_SYNC`); user and role changes for those
organizations must be made in the upstream identity provider.

Required scope: delete:organization_users

API: DELETE /api/v1/organizations/{org_code}/users/{user_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations remove-user [flags]
```

### Options

```
  -h, --help              help for remove-user
      --org-code string   The organization's code (required)
      --user-id string    The user's id (required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

