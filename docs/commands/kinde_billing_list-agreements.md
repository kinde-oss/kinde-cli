## kinde billing list-agreements

Get billing agreements

### Synopsis

Returns all the agreements a billing customer currently has access to

Required scope: read:billing_agreements

API: GET /api/v1/billing/agreements

```
kinde billing list-agreements [flags]
```

### Options

```
      --customer-id string      The ID of the billing customer to retrieve agreements for (required)
      --ending-before string    The ID of the billing agreement to end before
      --feature-code string     The feature code to filter by agreements only containing that feature
  -h, --help                    help for list-agreements
      --page-size int           Number of results per page. Defaults to 10 if parameter not sent
      --starting-after string   The ID of the billing agreement to start after
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde billing](kinde_billing.md)	 - Manage billing agreements, entitlements and usage

