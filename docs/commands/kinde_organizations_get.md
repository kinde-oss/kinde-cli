## kinde organizations get

Get organization

### Synopsis

Retrieve organization details by code.

Required scope: read:organizations

API: GET /api/v1/organization

```
kinde organizations get [flags]
```

### Options

```
      --code string     The organization's code (required)
      --expand string   Additional data to include in the response. Allowed value: "billing"
  -h, --help            help for get
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

