## kinde billing list-entitlements

Get billing entitlements

### Synopsis

Returns all the entitlements a billing customer currently has access to

Required scope: read:billing_entitlements

API: GET /api/v1/billing/entitlements

```
kinde billing list-entitlements [flags]
```

### Options

```
      --customer-id string      The ID of the billing customer to retrieve entitlements for (required)
      --ending-before string    The ID of the billing entitlement to end before
      --expand string           Additional plan data to include in the response. Allowed value: "plans" (one of: plans)
  -h, --help                    help for list-entitlements
      --max-value string        When the maximum limit of an entitlement is null, this value is returned as the maximum limit
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the billing entitlement to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde billing](kinde_billing.md)	 - Manage billing agreements, entitlements and usage

