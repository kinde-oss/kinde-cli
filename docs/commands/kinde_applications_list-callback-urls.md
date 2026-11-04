## kinde applications list-callback-urls

List Callback URLs

### Synopsis

Returns an application's redirect callback URLs.

Required scope: read:applications_redirect_uris

API: GET /api/v1/applications/{app_id}/auth_redirect_urls

```
kinde applications list-callback-urls [flags]
```

### Options

```
      --app-id string   The identifier for the application (required)
  -h, --help            help for list-callback-urls
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

