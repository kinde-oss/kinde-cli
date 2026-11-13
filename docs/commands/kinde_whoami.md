## kinde whoami

Show the logins for the domain

### Synopsis

Show the logins for the domain as JSON:

  m2m             the M2M application ('kinde login --client-id')
  browser         you, signed in through the browser ('kinde login')
  user            a user of your business, for the 'kinde account' commands
                  ('kinde account login')
  management_api  which login the Management API commands use: "m2m" when
                  it's set up, otherwise "browser"

A login that isn't set up is null. Exits with status 1 if none is set up.

```
kinde whoami [flags]
```

### Options

```
  -h, --help   help for whoami
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde](kinde.md)	 - Kinde CLI

