## kinde organizations list-invites

Get organization invites

### Synopsis

Get a list of invitations for an organization. By default, only pending
(non-revoked, non-accepted) invitations are returned.

Required scope: read:organization_invites

API: GET /api/v1/organization/{org_code}/invites

```
kinde organizations list-invites [flags]
```

### Options

```
  -h, --help                help for list-invites
      --include-accepted    Include accepted invitations in the results
      --include-revoked     Include revoked invitations in the results
      --next-token string   A string to get the next page of results if there are more results
      --org-code string     The organization's code (required)
      --page-size int       Number of results per page. Defaults to 10 if parameter not sent
      --sort string         Field and order to sort the result by (one of: created_on_asc, created_on_desc, email_asc, email_desc, name_asc, name_desc)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

