## kinde properties delete

Delete Property

### Synopsis

Delete property.

Required scope: delete:properties

API: DELETE /api/v1/properties/{property_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde properties delete [flags]
```

### Options

```
  -h, --help                 help for delete
      --property-id string   The unique identifier for the property (required)
  -y, --yes                  Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde properties](kinde_properties.md)	 - Manage custom properties

