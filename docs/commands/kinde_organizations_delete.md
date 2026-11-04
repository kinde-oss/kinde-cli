## kinde organizations delete

Delete Organization

### Synopsis

Delete an organization.

Deletion is refused when a billing customer linked to the organization has
agreements. Billing customers with no agreements are removed with the
organization.

Required scope: delete:organizations

API: DELETE /api/v1/organization/{org_code}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde organizations delete [flags]
```

### Options

```
  -h, --help              help for delete
      --org-code string   The identifier for the organization (required)
  -y, --yes               Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

