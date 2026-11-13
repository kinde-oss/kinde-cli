## kinde organizations enable-connection

Enable connection

### Synopsis

Enable an auth connection for an organization.

Required scope: create:organization_connections

API: POST /api/v1/organizations/{organization_code}/connections/{connection_id}

```
kinde organizations enable-connection [flags]
```

### Options

```
      --connection-id string       The identifier for the connection (required)
  -h, --help                       help for enable-connection
      --organization-code string   The unique code for the organization (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

