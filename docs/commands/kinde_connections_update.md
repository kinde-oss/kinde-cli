## kinde connections update

Update Connection

### Synopsis

Update Connection.

Required scope: update:connections

API: PATCH /api/v1/connections/{connection_id}

```
kinde connections update [flags]
```

### Options

```
      --body string                    Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --connection-id string           The unique identifier for the connection (required)
      --display-name string            The public facing name of the connection
      --enabled-applications strings   Client IDs of applications in which this connection is to be enabled (comma-separated list; required)
  -h, --help                           help for update
      --name string                    The internal name of the connection
      --options string                 JSON value
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde connections](kinde_connections.md)	 - Manage authentication connections

