## kinde environment-variables update

Update environment variable

### Synopsis

Update an environment variable you previously created. This feature is in beta
and admin UI is not yet available.

Required scope: update:environment_variables

API: PATCH /api/v1/environment_variables/{variable_id}

```
kinde environment-variables update [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                 help for update
      --is-secret            Whether the environment variable is sensitive. Secret variables are not-readable by you or your team after creation
      --key string           The key to update
      --value string         The new value for the environment variable
      --variable-id string   The environment variable's ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment-variables](kinde_environment-variables.md)	 - Manage environment variables

