## kinde organizations delete-user-permission

Delete Organization User Permission

### Synopsis

Delete permission for an organization user.

Required scope: delete:organization_user_permissions

API: DELETE /api/v1/organizations/{org_code}/users/{user_id}/permissions/{permission_id}

```
kinde organizations delete-user-permission [flags]
```

### Options

```
  -h, --help                   help for delete-user-permission
      --org-code string        The organization's code (required)
      --permission-id string   The permission id (required)
      --user-id string         The user's id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

