## kinde organizations delete-handle

Delete organization handle

### Synopsis

Delete organization handle

Required scope: delete:organization_handles

API: DELETE /api/v1/organization/{org_code}/handle

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations delete-handle [flags]
```

### Options

```
  -h, --help              help for delete-handle
      --org-code string   The organization's code (required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

