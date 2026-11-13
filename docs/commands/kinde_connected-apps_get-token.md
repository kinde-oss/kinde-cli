## kinde connected-apps get-token

Get Connected App Token

### Synopsis

Get an access token that can be used to call the third-party provider linked to
the connected app.

Required scope: read:connected_apps

API: GET /api/v1/connected_apps/token

```
kinde connected-apps get-token [flags]
```

### Options

```
  -h, --help                help for get-token
      --session-id string   The unique sesssion id representing the login session of a user (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde connected-apps](kinde_connected-apps.md)	 - Manage connected apps

