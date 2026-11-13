## kinde applications create

Create application

### Synopsis

Create a new client.

Required scope: create:applications

API: POST /api/v1/applications

```
kinde applications create [flags]
```

### Options

```
      --body string       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help              help for create
      --name string       The application's name (required)
      --org-code string   Scope an M2M application to an org (Plus plan required)
      --type string       The application's type (one of: reg, spa, m2m, device; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

