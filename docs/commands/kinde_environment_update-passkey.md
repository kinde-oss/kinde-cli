## kinde environment update-passkey

Update environment passkey settings

### Synopsis

Set the passkey policy for the current environment. Policies other than `off`
require the `passkeys` entitlement.

Required scope: update:passkey

API: PUT /api/v1/passkey

```
kinde environment update-passkey [flags]
```

### Options

```
      --body string     Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help            help for update-passkey
      --policy string   Environment passkey policy (one of: off, optional, mandatory; required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde environment](kinde_environment.md)	 - Manage the current environment

