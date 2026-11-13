## kinde apis authorize-applications

Authorize API applications

### Synopsis

Authorize applications to be allowed to request access tokens for an API

Required scope: update:apis

API: PATCH /api/v1/apis/{api_id}/applications

```
kinde apis authorize-applications [flags]
```

### Options

```
      --apiid string          The API's ID (required)
      --applications string   JSON value; required
      --body string           Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                  help for authorize-applications
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

