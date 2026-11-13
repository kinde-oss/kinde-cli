## kinde organizations list-connections

Get connections

### Synopsis

Gets all connections for an organization.

Required scope: read:organization_connections

API: GET /api/v1/organizations/{organization_code}/connections

```
kinde organizations list-connections [flags]
```

### Options

```
  -h, --help                       help for list-connections
      --organization-code string   The organization code (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

