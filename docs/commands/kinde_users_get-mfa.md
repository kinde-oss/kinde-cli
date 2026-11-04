## kinde users get-mfa

Get user's MFA configuration

### Synopsis

Get a user’s MFA configuration.

Required scope: read:user_mfa

API: GET /api/v1/users/{user_id}/mfa

```
kinde users get-mfa [flags]
```

### Options

```
  -h, --help             help for get-mfa
      --user-id string   The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

