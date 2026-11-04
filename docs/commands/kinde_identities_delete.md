## kinde identities delete

Delete identity

### Synopsis

Delete identity by ID.

Identities cannot be deleted from a user managed by directory sync
(`USER_MANAGED_BY_DIRECTORY_SYNC`); user and role changes must be made in the
upstream identity provider.

Required scope: delete:identities

API: DELETE /api/v1/identities/{identity_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde identities delete [flags]
```

### Options

```
  -h, --help                 help for delete
      --identity-id string   The unique identifier for the identity (required)
  -y, --yes                  Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde identities](kinde_identities.md)	 - Manage user identities

