## kinde connections create

Create Connection

### Synopsis

Create Connection.

Required scope: create:connections

API: POST /api/v1/connections

```
kinde connections create [flags]
```

### Options

```
      --body string                    Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --display-name string            The public facing name of the connection
      --enabled-applications strings   Client IDs of applications in which this connection is to be enabled (comma-separated list; required)
  -h, --help                           help for create
      --name string                    The internal name of the connection
      --options string                 JSON value
      --organization-code string       Enterprise connections only - the code for organization that manages this connection
      --strategy string                The identity provider identifier for the connection (one of: oauth2:apple, oauth2:azure_ad, oauth2:bitbucket, oauth2:discord, oauth2:facebook, oauth2:github, oauth2:gitlab, oauth2:google, oauth2:linkedin, oauth2:microsoft, oauth2:patreon, oauth2:slack, oauth2:stripe, oauth2:twitch, oauth2:twitter, oauth2:xero, saml:custom, saml:cloudflare, saml:okta, saml:microsoft, saml:google, wsfed:azure_ad)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde connections](kinde_connections.md)	 - Manage authentication connections

