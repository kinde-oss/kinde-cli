## kinde login

Log in to use the Management API, through the browser or with an M2M application

### Synopsis

Log in to a Kinde business so the Management API commands can run.

With no flags, the browser opens and you sign in with your Kinde account. What
you can do follows your team role. This uses Kinde's AI agent sign-in, which an
owner turns on in Kinde under Settings > APIs > Kinde Management API, and sends
commands through the business's MCP server. A few commands aren't available
there, and secrets such as client secrets are hidden: log in with an M2M
application for those.

With --client-id, you log in with a machine-to-machine (M2M) application
instead. Create it in Kinde, authorize it for the Kinde Management API, and
grant it the scopes you need. The client secret is asked for without being
shown, and is stored in your operating system's keychain. In scripts, pass the
secret on stdin with --client-secret-stdin, or skip logging in and set KINDE_MANAGEMENT_CLIENT_ID
and KINDE_MANAGEMENT_CLIENT_SECRET.

When both logins are set up, Management API commands use the M2M application.

To log in as a user of your business for the Account API, use
'kinde account login'.

```
kinde login [flags]
```

### Examples

```
  kinde login --domain mybusiness.kinde.com
  kinde login --domain mybusiness.kinde.com --client-id <client ID>
  echo "$CLIENT_SECRET" | kinde login --client-id <client ID> --client-secret-stdin
```

### Options

```
      --client-id string      Client ID of an M2M application, to log in with it instead of the browser
      --client-secret-stdin   Read the M2M application's client secret from stdin
  -h, --help                  help for login
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde](kinde.md)	 - Kinde CLI

