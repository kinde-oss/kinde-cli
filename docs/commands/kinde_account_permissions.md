## kinde account permissions

Get permissions

### Synopsis

Returns all the permissions the user has

API: GET /account_api/v1/permissions

```
kinde account permissions [flags]
```

### Options

```
  -h, --help                    help for permissions
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the permission to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user and see what they can access (Account API)

