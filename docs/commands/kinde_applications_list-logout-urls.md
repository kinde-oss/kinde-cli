## kinde applications list-logout-urls

List logout URLs

### Synopsis

Returns an application's logout redirect URLs.

Required scope: read:application_logout_uris

API: GET /api/v1/applications/{app_id}/auth_logout_urls

```
kinde applications list-logout-urls [flags]
```

### Options

```
      --app-id string   The identifier for the application (required)
  -h, --help            help for list-logout-urls
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

