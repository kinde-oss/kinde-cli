## kinde users update

Update user

### Synopsis

Update a user record.

When `given_name` or `family_name` is updated and the user is the owner of a
family billing customer, the change is also propagated to the corresponding
billing customer details.

Users managed by directory sync cannot be updated
(`USER_MANAGED_BY_DIRECTORY_SYNC`); user and role changes must be made in the
upstream identity provider.

Required scope: update:users

API: PATCH /api/v1/user

```
kinde users update [flags]
```

### Options

```
      --body string                   Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --family-name string            User's last name
      --given-name string             User's first name
  -h, --help                          help for update
      --id string                     The user's id (required)
      --is-password-reset-requested   Prompt the user to change their password on next sign in
      --is-suspended                  Whether the user is currently suspended or not
      --picture string                The user's profile picture
      --provided-id string            An external id to reference the user
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

