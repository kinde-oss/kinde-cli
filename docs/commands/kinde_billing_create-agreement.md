## kinde billing create-agreement

Create billing agreement

### Synopsis

Creates a new billing agreement based on the plan code passed, and cancels the
customer's existing agreements. Existing agreements are cancelled even if they
belong to a different billing group. The new plan must be the same customer type
as the customer (organization or user).

Changing plan starts a new subscription. One-time purchases stay on the previous
agreement and are no longer current. Buying the same item on the new plan is a
new purchase. This does not refund what was already paid.

`plan_version` selects a published version of `plan_code`. When it is omitted,
the latest version is assigned. Posting the customer's current plan_code does
not change their version. `plan_version` is required to move them onto a
different published version.

Required scope: create:billing_agreements

API: POST /api/v1/billing/agreements

```
kinde billing create-agreement [flags]
```

### Options

```
      --body string          Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --customer-id string   The ID of the billing customer to create a new agreement for (required)
  -h, --help                 help for create-agreement
      --is-invoice-now       Generate a final invoice for any un-invoiced metered usage
      --is-prorate           Generate a proration invoice item that credits remaining unused features
      --plan-code string     The code of the billing plan the new agreement will be based on (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde billing](kinde_billing.md)	 - Manage billing agreements, entitlements and usage

