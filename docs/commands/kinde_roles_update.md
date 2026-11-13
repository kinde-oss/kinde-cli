## kinde roles update

Update role

### Synopsis

Update a role

Required scope: update:roles

API: PATCH /api/v1/roles/{role_id}

```
kinde roles update [flags]
```

### Options

```
      --assignment-permission-id string   The public ID of the permission required to assign this role to users
      --body string                       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --description string                The role's description
  -h, --help                              help for update
      --is-default-role                   Set role as default for new users
      --key string                        The role identifier to use in code
      --name string                       The role's name
      --role-id string                    The identifier for the role (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde roles](kinde_roles.md)	 - Manage roles, their permissions and scopes

