## kinde applications remove-access-role

Remove application access role

### Synopsis

Remove a role from the set that is allowed to access an application. Removing a
role does not change whether role-based access control is enabled for the
application.

Required scope: update:applications

API: DELETE /api/v1/applications/{application_id}/access_roles/{role_id}

```
kinde applications remove-access-role [flags]
```

### Options

```
      --application-id string   The identifier/client ID for the application (required)
  -h, --help                    help for remove-access-role
      --role-id string          The identifier for the role (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

