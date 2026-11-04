## kinde directories update

Update SCIM directory

### Synopsis

Update SCIM directory configuration.

Required scope: update:scim_directories

API: PATCH /api/v1/directories/{directory_id}

```
kinde directories update [flags]
```

### Options

```
      --body string             Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --directory-id string     The directory's ID (required)
      --directory-name string   A descriptive name for the SCIM directory
  -h, --help                    help for update
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde directories](kinde_directories.md)	 - Manage SCIM directories

