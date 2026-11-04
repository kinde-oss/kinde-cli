## kinde identities update

Update identity

### Synopsis

Update identity by ID.

Identities cannot be updated on a user managed by directory sync
(`USER_MANAGED_BY_DIRECTORY_SYNC`); user and role changes must be made in the
upstream identity provider.

Required scope: update:identities

API: PATCH /api/v1/identities/{identity_id}

```
kinde identities update [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                 help for update
      --identity-id string   The unique identifier for the identity (required)
      --is-primary           Whether the identity is the primary for it's type
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde identities](kinde_identities.md)	 - Manage user identities

