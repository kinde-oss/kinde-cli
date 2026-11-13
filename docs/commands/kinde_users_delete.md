## kinde users delete

Delete user

### Synopsis

Delete a user record.

Users managed by directory sync cannot be deleted
(`USER_MANAGED_BY_DIRECTORY_SYNC`); user and role changes must be made in the
upstream identity provider.

Required scope: delete:users

API: DELETE /api/v1/user

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde users delete [flags]
```

### Options

```
  -h, --help                help for delete
      --id string           The user's id (required)
      --is-delete-profile   Delete all data and remove the user's profile from all of Kinde, including the subscriber list
  -y, --yes                 Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

