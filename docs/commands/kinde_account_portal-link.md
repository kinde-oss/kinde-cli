## kinde account portal-link

Get self-serve portal link

### Synopsis

Returns a link to the self-serve portal for the authenticated user. The user can
use this link to manage their account, update their profile, and view their
entitlements.

The link is single use and expires 10 minutes after it is generated. Redirect
the user to it rather than storing it or sending it on.

API: GET /account_api/v1/portal_link

```
kinde account portal-link [flags]
```

### Options

```
  -h, --help                help for portal-link
      --return-url string   The URL to redirect the user to after they have completed their actions in the portal
      --subnav string       The area of the portal you want the user to land on (one of: profile, organization_details, organization_payment_details, organization_plan_selection, payment_details, plan_details, plan_selection)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user and see what they can access (Account API)

