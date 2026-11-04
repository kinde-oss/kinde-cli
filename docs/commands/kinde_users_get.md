## kinde users get

Get user

### Synopsis

Retrieve a user record.

Required scope: read:users

API: GET /api/v1/user

```
kinde users get [flags]
```

### Options

```
      --expand string   Additional data to include in the response. One or more of (comma-separated): "organizations", "identities", "billing"
  -h, --help            help for get
      --id string       The user's id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

