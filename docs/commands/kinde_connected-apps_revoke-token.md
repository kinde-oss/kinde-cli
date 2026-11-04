## kinde connected-apps revoke-token

Revoke Connected App Token

### Synopsis

Revoke the tokens linked to the connected app session.

Required scope: create:connected_apps

API: POST /api/v1/connected_apps/revoke

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde connected-apps revoke-token [flags]
```

### Options

```
  -h, --help                help for revoke-token
      --session-id string   The unique sesssion id representing the login session of a user (required)
  -y, --yes                 Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde connected-apps](kinde_connected-apps.md)	 - Manage connected apps

