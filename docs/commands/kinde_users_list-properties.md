## kinde users list-properties

Get property values

### Synopsis

Gets properties for an user by ID.

Required scope: read:user_properties

API: GET /api/v1/users/{user_id}/properties

```
kinde users list-properties [flags]
```

### Options

```
  -h, --help             help for list-properties
      --user-id string   The user's ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

