## kinde applications list-properties

Get property values

### Synopsis

Gets properties for an application by client ID.

Required scope: read:application_properties

API: GET /api/v1/applications/{application_id}/properties

```
kinde applications list-properties [flags]
```

### Options

```
      --application-id string   The application's ID / client ID (required)
  -h, --help                    help for list-properties
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

