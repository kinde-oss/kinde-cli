## kinde apis create

Create API

### Synopsis

Register a new API. For more information read [Register and manage
APIs](https://docs.kinde.com/developer-tools/your-apis/register-manage-apis/).

Required scope: create:apis

API: POST /api/v1/apis

```
kinde apis create [flags]
```

### Options

```
      --audience string   A unique identifier for the API - commonly the URL (required)
      --body string       Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help              help for create
      --name string       The name of the API. (1-64 characters) (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde apis](kinde_apis.md)	 - Manage APIs and their scopes

