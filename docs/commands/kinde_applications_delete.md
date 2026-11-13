## kinde applications delete

Delete application

### Synopsis

Delete a client / application.

Required scope: delete:applications

API: DELETE /api/v1/applications/{application_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde applications delete [flags]
```

### Options

```
      --application-id string   The identifier for the application (required)
  -h, --help                    help for delete
  -y, --yes                     Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

