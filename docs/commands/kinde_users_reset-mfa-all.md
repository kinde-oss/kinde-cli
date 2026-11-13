## kinde users reset-mfa-all

Reset all environment MFA for a user

### Synopsis

Reset all environment MFA factors for a user.

Required scope: delete:user_mfa

API: DELETE /api/v1/users/{user_id}/mfa

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde users reset-mfa-all [flags]
```

### Options

```
  -h, --help             help for reset-mfa-all
      --user-id string   The identifier for the user (required)
  -y, --yes              Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

