## kinde users delete-sessions

Delete user sessions

### Synopsis

Invalidate user sessions.

Required scope: delete:user_sessions

API: DELETE /api/v1/users/{user_id}/sessions

```
kinde users delete-sessions [flags]
```

### Options

```
  -h, --help             help for delete-sessions
      --user-id string   The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

