## kinde organizations list-users

Get organization users

### Synopsis

Get user details for all members of an organization.

Required scope: read:organization_users

API: GET /api/v1/organizations/{org_code}/users

```
kinde organizations list-users [flags]
```

### Options

```
  -h, --help                 help for list-users
      --next-token string    A string to get the next page of results if there are more results
      --org-code string      The organization's code (required)
      --page-size int        Number of results per page. Defaults to 10 if parameter not sent
      --permissions string   Filter by user permissions comma separated (where all match)
      --roles string         Filter by user roles comma separated (where all match)
      --sort string          Field and order to sort the result by (one of: name_asc, name_desc, email_asc, email_desc)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

