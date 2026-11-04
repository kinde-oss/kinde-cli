## kinde connected-apps get-auth-url

Get Connected App URL

### Synopsis

Get a URL that authenticates and authorizes a user to a third-party connected
app.

Required scope: read:connected_apps

API: GET /api/v1/connected_apps/auth_url

```
kinde connected-apps get-auth-url [flags]
```

### Options

```
  -h, --help                           help for get-auth-url
      --key-code-ref string            The unique key code reference of the connected app to authenticate against (required)
      --org-code string                The code of the Kinde organization that needs to authenticate to the third-party connected app
      --override-callback-url string   A URL that overrides the default callback URL setup in your connected app configuration
      --user-id string                 The id of the user that needs to authenticate to the third-party connected app
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde connected-apps](kinde_connected-apps.md)	 - Manage connected apps

