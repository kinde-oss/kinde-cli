## kinde organizations update-properties

Update Organization Property values

### Synopsis

Update organization property values.

Required scope: update:organization_properties

API: PATCH /api/v1/organizations/{org_code}/properties

```
kinde organizations update-properties [flags]
```

### Options

```
      --body string       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help              help for update-properties
      --org-code string   The identifier for the organization (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

