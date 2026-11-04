## kinde applications remove-connection

Remove connection

### Synopsis

Turn off an auth connection for an application

Required scope: delete:application_connections

API: DELETE /api/v1/applications/{application_id}/connections/{connection_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde applications remove-connection [flags]
```

### Options

```
      --application-id string   The identifier/client ID for the application (required)
      --connection-id string    The identifier for the connection (required)
  -h, --help                    help for remove-connection
  -y, --yes                     Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

