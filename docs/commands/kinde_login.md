## kinde login

Log in with an M2M application to use the Management API

### Synopsis

Log in with a machine-to-machine (M2M) application, which the resource commands
use to call the Kinde Management API.

Create the M2M application in Kinde, authorize it for the Kinde Management API,
and grant it the scopes you need. The client secret is asked for without being
shown, and is stored in your operating system's keychain.

In scripts, pass the secret on stdin with --client-secret-stdin, or skip
logging in and set KINDE_MANAGEMENT_CLIENT_ID and KINDE_MANAGEMENT_CLIENT_SECRET.

To log in as a user through the browser, use 'kinde account login'.

```
kinde login [flags]
```

### Examples

```
  kinde login --domain mybusiness.kinde.com --client-id <client ID>
  echo "$CLIENT_SECRET" | kinde login --client-id <client ID> --client-secret-stdin
```

### Options

```
      --client-id string      Client ID of the M2M application (required)
      --client-secret-stdin   Read the client secret from stdin
  -h, --help                  help for login
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde](kinde.md)	 - Kinde CLI

