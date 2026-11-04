## kinde applications list-access-roles

Get application access roles

### Synopsis

Gets the roles configured as allowed to access an application. These roles are
enforced only when role-based access control is enabled for the application.

Required scope: read:applications

API: GET /api/v1/applications/{application_id}/access_roles

```
kinde applications list-access-roles [flags]
```

### Options

```
      --application-id string   The identifier/client ID for the application (required)
      --ending-before string    The ID of the role to end before
  -h, --help                    help for list-access-roles
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the role to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

