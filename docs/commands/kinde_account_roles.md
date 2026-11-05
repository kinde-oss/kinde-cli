## kinde account roles

Get roles

### Synopsis

Returns all roles for the user

API: GET /account_api/v1/roles

```
kinde account roles [flags]
```

### Options

```
  -h, --help                    help for roles
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the role to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user and see what they can access (Account API)

