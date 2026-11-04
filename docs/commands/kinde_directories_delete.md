## kinde directories delete

Delete SCIM directory

### Synopsis

Delete a SCIM directory and all associated data.

Required scope: delete:scim_directories

API: DELETE /api/v1/directories/{directory_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde directories delete [flags]
```

### Options

```
      --directory-id string   The directory's ID (required)
  -h, --help                  help for delete
  -y, --yes                   Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde directories](kinde_directories.md)	 - Manage SCIM directories

