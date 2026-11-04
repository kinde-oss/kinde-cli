## kinde environment add-logo

Add logo

### Synopsis

Add environment logo

Required scope: update:environments

API: PUT /api/v1/environment/logos/{type}

```
kinde environment add-logo [flags]
```

### Options

```
  -h, --help          help for add-logo
      --logo string   The logo file to upload (path to a file; required)
      --type string   The type of logo to add (one of: dark, light; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment](kinde_environment.md)	 - Manage the current environment

