## kinde account login

Log in as a user through the browser

### Synopsis

Log in as a user through the browser, using the device authorization flow.

Open the URL shown (or scan the QR code), sign in, and approve the login. Use
--client-id to sign in to a specific application; otherwise Kinde uses the
business's default application for device logins.

```
kinde account login [flags]
```

### Examples

```
  kinde account login --domain mybusiness.kinde.com
  kinde account login --client-id <application client ID>
```

### Options

```
      --client-id string   Client ID of the application to sign in to
  -h, --help               help for login
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user through the browser

