## kinde property-categories create

Create Category

### Synopsis

Create category.

Required scope: create:property_categories

API: POST /api/v1/property_categories

```
kinde property-categories create [flags]
```

### Options

```
      --body string      Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --context string   The context that the category applies to (one of: org, usr, app; required)
  -h, --help             help for create
      --name string      The name of the category (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde property-categories](kinde_property-categories.md)	 - Manage property categories

