## kinde whoami

Show the logins for the domain

### Synopsis

Show the logins for the domain as JSON: the M2M application used by the
Management API commands, and the user logged in through the browser for the
'kinde account' commands. A login that isn't set up is null.

Exits with status 1 if neither login is set up.

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

