## kinde account token

Print the access token of the user login

### Synopsis

Print the access token of the user login ('kinde account login').

Use it to call the Account API, or your own APIs, with other tools. Treat it like
a password: anyone with the token can act as you until it expires.

```
kinde account token [flags]
```

### Examples

```
  curl -H "Authorization: Bearer $(kinde account token)" https://mybusiness.kinde.com/account_api/v1/permissions
```

### Options

```
  -h, --help   help for token
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user and see what they can access (Account API)

