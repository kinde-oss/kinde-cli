## kinde applications enable-connection

Enable connection

### Synopsis

Enable an auth connection for an application.

Required scope: create:application_connections

API: POST /api/v1/applications/{application_id}/connections/{connection_id}

```
kinde applications enable-connection [flags]
```

### Options

```
      --application-id string   The identifier/client ID for the application (required)
      --connection-id string    The identifier for the connection (required)
  -h, --help                    help for enable-connection
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

