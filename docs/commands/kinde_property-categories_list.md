## kinde property-categories list

List categories

### Synopsis

Returns a list of categories.

Required scope: read:property_categories

API: GET /api/v1/property_categories

```
kinde property-categories list [flags]
```

### Options

```
      --context string          Filter the results by User or Organization context (one of: usr, org)
      --ending-before string    The ID of the category to end before
  -h, --help                    help for list
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the category to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde property-categories](kinde_property-categories.md)	 - Manage property categories

