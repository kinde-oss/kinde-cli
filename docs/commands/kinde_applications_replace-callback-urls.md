## kinde applications replace-callback-urls

Replace Redirect Callback URLs

### Synopsis

Replace all redirect callback URLs.

Required scope: update:application_redirect_uris

API: PUT /api/v1/applications/{app_id}/auth_redirect_urls

```
kinde applications replace-callback-urls [flags]
```

### Options

```
      --app-id string   The identifier for the application (required)
      --body string     Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help            help for replace-callback-urls
      --urls strings    Array of callback urls (comma-separated list; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

