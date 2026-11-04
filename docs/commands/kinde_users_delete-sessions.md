## kinde users delete-sessions

Delete user sessions

### Synopsis

Invalidate user sessions.

Required scope: delete:user_sessions

API: DELETE /api/v1/users/{user_id}/sessions

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde users delete-sessions [flags]
```

### Options

```
  -h, --help             help for delete-sessions
      --user-id string   The identifier for the user (required)
  -y, --yes              Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

