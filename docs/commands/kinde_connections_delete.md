## kinde connections delete

Delete Connection

### Synopsis

Delete connection.

Required scope: delete:connections

API: DELETE /api/v1/connections/{connection_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde connections delete [flags]
```

### Options

```
      --connection-id string   The identifier for the connection (required)
  -h, --help                   help for delete
  -y, --yes                    Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde connections](kinde_connections.md)	 - Manage authentication connections

