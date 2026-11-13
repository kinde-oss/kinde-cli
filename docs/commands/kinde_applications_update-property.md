## kinde applications update-property

Update property

### Synopsis

Update application property value.

Required scope: update:application_properties

API: PUT /api/v1/applications/{application_id}/properties/{property_key}

```
kinde applications update-property [flags]
```

### Options

```
      --application-id string   The application's ID / client ID (required)
      --body string             Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                    help for update-property
      --property-key string     The property's key (required)
      --value string            The new value for the property. Send an empty string to unset the property, so it reads back as not set (JSON value; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

