## kinde users list-identities

Get identities

### Synopsis

Gets a list of identities for an user by ID.

Required scope: read:user_identities

API: GET /api/v1/users/{user_id}/identities

```
kinde users list-identities [flags]
```

### Options

```
      --ending-before string    The ID of the identity to end before
  -h, --help                    help for list-identities
      --starting-after string   The ID of the identity to start after
      --user-id string          The user's ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

