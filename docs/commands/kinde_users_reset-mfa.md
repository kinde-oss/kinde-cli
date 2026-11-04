## kinde users reset-mfa

Reset specific environment MFA for a user

### Synopsis

Reset a specific environment MFA factor for a user.

Required scope: delete:user_mfa

API: DELETE /api/v1/users/{user_id}/mfa/{factor_id}

```
kinde users reset-mfa [flags]
```

### Options

```
      --factor-id string   The identifier for the MFA factor (required)
  -h, --help               help for reset-mfa
      --user-id string     The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

