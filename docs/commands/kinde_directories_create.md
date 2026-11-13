## kinde directories create

Create SCIM directory

### Synopsis

Create a new SCIM directory for user and group synchronization.

Required scope: create:scim_directories

API: POST /api/v1/directories

```
kinde directories create [flags]
```

### Options

```
      --body string                       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --directory-name string             An optional descriptive name for the SCIM directory
      --enterprise-connection-id string   The enterprise connection ID to associate with this directory for SCIM-provisioned users
  -h, --help                              help for create
      --org-code string                   The organization code to create the SCIM directory for (required)
      --provider-code string              The SCIM provider code to use for this directory (one of: entra_id_azure_ad, okta, google_workspace, custom_scim_v2, cyberark, jumpcloud, onelogin, pingfederate, rippling)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde directories](kinde_directories.md)	 - Manage SCIM directories

