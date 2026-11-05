## kinde account entitlement

Get entitlement

### Synopsis

Returns the entitlement matching the `key` feature key, for self-serve billing
visibility, using the caller's user access token.

Prefers the user's active personal plan. If they do not have one, returns the
matching organization entitlement when they have `org:read:billing` (same gate
as the self-serve portal organization Plan page).

Not intended as the universal B2B feature gate for every org member. For
server-side feature gating from a billing customer's plan, use the Management
API `GET /api/v1/billing/entitlements` with the relevant `customer_id` instead.

API: GET /account_api/v1/entitlement

```
kinde account entitlement [flags]
```

### Options

```
  -h, --help         help for entitlement
      --key string   The key of the feature (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde account](kinde_account.md)	 - Log in as a user and see what they can access (Account API)

