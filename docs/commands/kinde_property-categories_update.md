## kinde property-categories update

Update Category

### Synopsis

Update category.

Required scope: update:property_categories

API: PUT /api/v1/property_categories/{category_id}

```
kinde property-categories update [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --category-id string   The unique identifier for the category (required)
  -h, --help                 help for update
      --name string          The name of the category
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde property-categories](kinde_property-categories.md)	 - Manage property categories

