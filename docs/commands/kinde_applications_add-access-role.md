## kinde applications add-access-role

Add application access role

### Synopsis

Add a role to the set configured as allowed to access an application. Adding a
role does not enable role-based access control for the application.

Required scope: update:applications

API: POST /api/v1/applications/{application_id}/access_roles/{role_id}

```
kinde applications add-access-role [flags]
```

### Options

```
      --application-id string   The identifier/client ID for the application (required)
  -h, --help                    help for add-access-role
      --role-id string          The identifier for the role (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

