## kinde applications delete-callback-urls

Delete Callback URLs

### Synopsis

Delete callback URLs.

Required scope: delete:application_redirect_uris

API: DELETE /api/v1/applications/{app_id}/auth_redirect_urls

```
kinde applications delete-callback-urls [flags]
```

### Options

```
      --app-id string   The identifier for the application (required)
  -h, --help            help for delete-callback-urls
      --urls string     Urls to delete, comma separated and url encoded (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

