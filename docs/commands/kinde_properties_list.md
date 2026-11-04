## kinde properties list

List properties

### Synopsis

Returns a list of properties

Required scope: read:properties

API: GET /api/v1/properties

```
kinde properties list [flags]
```

### Options

```
      --context string          Filter results by user, organization or application context (one of: usr, org, app)
      --ending-before string    The ID of the property to end before
  -h, --help                    help for list
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the property to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde properties](kinde_properties.md)	 - Manage custom properties

