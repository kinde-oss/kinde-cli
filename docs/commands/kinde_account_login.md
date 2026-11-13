## kinde account login

Log in as a user of your business, for the Account API

### Synopsis

Log in as a user of your business, using the device authorization flow, so
the 'kinde account' commands can show what that user can access.

Open the URL shown (or scan the QR code), sign in, and approve the login.

The login needs a Device and IoT application in your business. Pass its client
ID with --client-id, or set one as the default for device logins in Kinde and
leave --client-id out.

This isn't how you log in to run Management API commands: use 'kinde login'
for that.

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
      --client-id string   Client ID of a Device and IoT application (default: the business's default device application)
  -h, --help               help for login
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user and see what they can access (Account API)

