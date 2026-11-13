## kinde organizations update-property

Update Organization Property value

### Synopsis

Update organization property value.

Required scope: update:organization_properties

API: PUT /api/v1/organizations/{org_code}/properties/{property_key}

```
kinde organizations update-property [flags]
```

### Options

```
  -h, --help                  help for update-property
      --org-code string       The identifier for the organization (required)
      --property-key string   The identifier for the property (required)
      --value string          The new property value. Send an empty string to unset the property, so it reads back as not set (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

