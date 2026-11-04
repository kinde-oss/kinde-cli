## kinde applications list-connections

Get connections

### Synopsis

Gets all connections for an application.

Required scope: read:application_connections

API: GET /api/v1/applications/{application_id}/connections

```
kinde applications list-connections [flags]
```

### Options

```
      --application-id string   The identifier/client ID for the application (required)
  -h, --help                    help for list-connections
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

