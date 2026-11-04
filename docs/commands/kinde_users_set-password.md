## kinde users set-password

Set User password

### Synopsis

Set user password.

Passwords cannot be set on a user managed by directory sync
(`USER_MANAGED_BY_DIRECTORY_SYNC`); user and role changes must be made in the
upstream identity provider.

Required scope: update:user_passwords

API: PUT /api/v1/users/{user_id}/password

```
kinde users set-password [flags]
```

### Options

```
      --body string              Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --hashed-password string   The hashed password. For aspnet-identity-v2, provide the base64-encoded AspNetUsers.PasswordHash value as-is (required)
      --hashing-method string    The hashing method or algorithm used to encrypt the user’s password. Default is bcrypt (one of: bcrypt, crypt, md5, sha256, wordpress, pbkdf2, firebase-scrypt, aspnet-identity-v2)
  -h, --help                     help for set-password
      --is-temporary-password    The user will be prompted to set a new password after entering this one
      --iterations int           The iteration count (factor) used to derive the hash
      --mem-cost int             The scrypt memory cost from the Firebase project's password hash parameters (1-20)
      --rounds int               The scrypt rounds from the Firebase project's password hash parameters (1-16)
      --salt string              Extra characters added to passwords to make them stronger
      --salt-position string     Position of salt in password string. Not required for bcrypt (one of: prefix, suffix)
      --salt-separator string    The base64-encoded salt separator from the Firebase project's password hash parameters
      --signer-key string        The base64-encoded signer key from the Firebase project's password hash parameters
      --user-id string           The identifier for the user (required)
      --variant string           The hashing variant
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

