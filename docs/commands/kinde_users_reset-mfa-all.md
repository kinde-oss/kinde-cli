## kinde users reset-mfa-all

Reset all environment MFA for a user

### Synopsis

Reset all environment MFA factors for a user.

Required scope: delete:user_mfa

API: DELETE /api/v1/users/{user_id}/mfa

```
kinde users reset-mfa-all [flags]
```

### Options

```
  -h, --help             help for reset-mfa-all
      --user-id string   The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

