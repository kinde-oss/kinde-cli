## kinde users update-properties

Update Property values

### Synopsis

Update property values.

Required scope: update:user_properties

API: PATCH /api/v1/users/{user_id}/properties

```
kinde users update-properties [flags]
```

### Options

```
      --body string      Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help             help for update-properties
      --user-id string   The identifier for the user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

