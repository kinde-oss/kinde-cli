## kinde organizations list-properties

Get Organization Property Values

### Synopsis

Gets properties for an organization by org code.

Required scope: read:organization_properties

API: GET /api/v1/organizations/{org_code}/properties

```
kinde organizations list-properties [flags]
```

### Options

```
  -h, --help              help for list-properties
      --org-code string   The organization's code (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

