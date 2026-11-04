## kinde directories get

Get SCIM directory

### Synopsis

Retrieve SCIM directory details by ID.

Required scope: read:scim_directories

API: GET /api/v1/directories/{directory_id}

```
kinde directories get [flags]
```

### Options

```
      --directory-id string   The directory's ID (required)
  -h, --help                  help for get
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde directories](kinde_directories.md)	 - Manage SCIM directories

