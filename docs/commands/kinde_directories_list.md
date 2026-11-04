## kinde directories list

Get SCIM directories

### Synopsis

Returns a list of SCIM directories for your organization.

Required scope: read:scim_directories

API: GET /api/v1/directories

```
kinde directories list [flags]
```

### Options

```
  -h, --help                       help for list
      --organization-code string   Filter by organization code to get directories for a specific organization
      --page-size int              Number of results per page. Defaults to 50 if parameter not sent
      --starting-after string      The ID of the directory to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde directories](kinde_directories.md)	 - Manage SCIM directories

