## kinde organizations remove-connection

Remove connection

### Synopsis

Turn off an auth connection for an organization

Required scope: delete:organization_connections

API: DELETE /api/v1/organizations/{organization_code}/connections/{connection_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations remove-connection [flags]
```

### Options

```
      --connection-id string       The identifier for the connection (required)
  -h, --help                       help for remove-connection
      --organization-code string   The unique code for the organization (required)
  -y, --yes                        Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

