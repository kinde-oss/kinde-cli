## kinde organizations delete-logo

Delete organization logo

### Synopsis

Delete organization logo

Required scope: update:organizations

API: DELETE /api/v1/organizations/{org_code}/logos/{type}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations delete-logo [flags]
```

### Options

```
  -h, --help              help for delete-logo
      --org-code string   The organization's code (required)
      --type string       The type of logo to delete (one of: dark, light; required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

