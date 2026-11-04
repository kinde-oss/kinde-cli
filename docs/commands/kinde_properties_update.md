## kinde properties update

Update Property

### Synopsis

Update property.

Required scope: update:properties

API: PUT /api/v1/properties/{property_id}

```
kinde properties update [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --category-id string   Which category the property belongs to (required)
      --description string   Description of the property purpose
  -h, --help                 help for update
      --is-private           Whether the property can be included in id and access tokens (required)
      --name string          The name of the property (required)
      --property-id string   The unique identifier for the property (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde properties](kinde_properties.md)	 - Manage custom properties

