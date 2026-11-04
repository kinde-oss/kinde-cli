## kinde organizations add-logo

Add organization logo

### Synopsis

Add organization logo

Required scope: update:organizations

API: POST /api/v1/organizations/{org_code}/logos/{type}

```
kinde organizations add-logo [flags]
```

### Options

```
  -h, --help              help for add-logo
      --logo string       The logo file to upload (path to a file; required)
      --org-code string   The organization's code (required)
      --type string       The type of logo to add (one of: dark, light; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

