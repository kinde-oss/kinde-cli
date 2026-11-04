## kinde applications update-tokens

Update application tokens

### Synopsis

Configure tokens for an application.

Required scope: update:application_tokens

API: PATCH /api/v1/applications/{application_id}/tokens

```
kinde applications update-tokens [flags]
```

### Options

```
      --access-token-lifetime int            The lifetime of an access token in seconds
      --application-id string                The identifier/client ID for the application (required)
      --authenticated-session-lifetime int   The lifetime of an authenticated session in seconds
      --body string                          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                                 help for update-tokens
      --id-token-lifetime int                The lifetime of an ID token in seconds
      --is-hasura-mapping-enabled            Enable or disable Hasura mapping
      --refresh-token-lifetime int           The lifetime of a refresh token in seconds
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde applications](kinde_applications.md)	 - Manage applications, their connections and URLs

