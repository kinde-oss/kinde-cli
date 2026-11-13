## kinde users list-sessions

Get user sessions

### Synopsis

Retrieve the list of active sessions for a specific user.

Required scope: read:user_sessions

API: GET /api/v1/users/{user_id}/sessions

```
kinde users list-sessions [flags]
```

### Options

```
  -h, --help             help for list-sessions
      --user-id string   The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

