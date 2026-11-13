## kinde account entitlements

Get entitlements

### Synopsis

Returns entitlements for self-serve billing visibility (for example a Plan
details screen), using the caller's user access token.

Prefers the user's active personal plan. If they do not have one, returns their
organization's plan entitlements when they have `org:read:billing` (same gate as
the self-serve portal organization Plan page).

Not intended as the universal B2B feature gate for every org member. For
server-side feature gating from a billing customer's plan, use the Management
API `GET /api/v1/billing/entitlements` with the relevant `customer_id` instead.

API: GET /account_api/v1/entitlements

```
kinde account entitlements [flags]
```

### Options

```
  -h, --help                    help for entitlements
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the entitlement to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user and see what they can access (Account API)

