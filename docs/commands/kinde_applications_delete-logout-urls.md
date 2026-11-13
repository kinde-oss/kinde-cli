## kinde applications delete-logout-urls

Delete Logout URLs

### Synopsis

Delete logout URLs.

Required scope: delete:application_logout_uris

API: DELETE /api/v1/applications/{app_id}/auth_logout_urls

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde applications delete-logout-urls [flags]
```

### Options

```
      --app-id string   The identifier for the application (required)
  -h, --help            help for delete-logout-urls
      --urls string     Urls to delete, comma separated and url encoded (required)
  -y, --yes             Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

