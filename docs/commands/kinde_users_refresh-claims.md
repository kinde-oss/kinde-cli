## kinde users refresh-claims

Refresh User Claims and Invalidate Cache

### Synopsis

Refreshes the user's claims and invalidates the current cache.

Required scope: update:user_refresh_claims

API: POST /api/v1/users/{user_id}/refresh_claims

```
kinde users refresh-claims [flags]
```

### Options

```
  -h, --help             help for refresh-claims
      --user-id string   The id of the user whose claims needs to be updated (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

