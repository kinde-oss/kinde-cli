## kinde environment-variables get

Get environment variable

### Synopsis

Retrieve environment variable details by ID. This feature is in beta and admin
UI is not yet available.

Required scope: read:environment_variables

API: GET /api/v1/environment_variables/{variable_id}

```
kinde environment-variables get [flags]
```

### Options

```
  -h, --help                 help for get
      --variable-id string   The environment variable's ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment-variables](kinde_environment-variables.md)	 - Manage environment variables

