## kinde environment-variables create

Create environment variable

### Synopsis

Create a new environment variable. This feature is in beta and admin UI is not
yet available.

Required scope: create:environment_variables

API: POST /api/v1/environment_variables

```
kinde environment-variables create [flags]
```

### Options

```
      --body string    Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help           help for create
      --is-secret      Whether the environment variable is sensitive. Secrets are not-readable by you or your team after creation
      --key string     The name of the environment variable (max 128 characters) (required)
      --value string   The value of the new environment variable (max 2048 characters) (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment-variables](kinde_environment-variables.md)	 - Manage environment variables

