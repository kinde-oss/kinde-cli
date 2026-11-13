## kinde properties create

Create Property

### Synopsis

Create property.

Required scope: create:properties

API: POST /api/v1/properties

```
kinde properties create [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --category-id string   Which category the property belongs to (required)
      --context string       The context that the property applies to (one of: org, usr, app; required)
      --description string   Description of the property purpose
  -h, --help                 help for create
      --is-private           Whether the property can be included in id and access tokens (required)
      --key string           The property identifier to use in code (required)
      --name string          The name of the property (required)
      --type string          The property type (one of: single_line_text, multi_line_text; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde properties](kinde_properties.md)	 - Manage custom properties

