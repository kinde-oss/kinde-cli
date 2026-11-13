## kinde billing create-meter-usage-record

Create meter usage record

### Synopsis

Create a new meter usage record against a billing agreement and feature.

Metered one-time features can receive usage after the customer has completed
their first purchase of that feature.

`meter_type_code` chooses how this record updates the running total. The caller
is responsible for sending the correct type and value. `meter_usage_timestamp`
is stored on the record (and, for plan-metered entitlements, sent to the billing
provider). It does not choose a billing cycle, does not insert the record in
timestamp order, and does not change the running-total arithmetic.

Running-total rules:

- `absolute` (default): replaces the running total with `meter_value`.
- `delta`: adds `meter_value` to the last calculated running total.

Which running total is updated:

- Plan-metered entitlements, and one-time features that expire at the end of the
  current billing cycle: the current billing cycle total. A new cycle starts at
  0.
- One-time features that never expire: a lifetime total for that agreement and
  feature. Every accepted record updates it, including records whose
  `meter_usage_timestamp` is before the first paid purchase. This total does not
  reset when the billing cycle rolls.

The response includes the created record `id` and a grant snapshot
(`units_granted`, `current_usage`, and `expiration_policy_code` when the grant
is a one-time purchase) so the caller can tell whether usage has exceeded the
grant without a follow-up GET. `current_usage` is that grant-relative running
total after this record. Grant fields are omitted when there is no current
grant. `units_granted` is also omitted when the grant is unlimited; an unlimited
one-time grant still returns `expiration_policy_code`.

Required scope: create:meter_usage

API: POST /api/v1/billing/meter_usage

```
kinde billing create-meter-usage-record [flags]
```

### Options

```
      --billing-feature-code string    The code of the feature within the agreement against which to record usage (required)
      --body string                    Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --customer-agreement-id string   The billing agreement against which to record usage (required)
  -h, --help                           help for create-meter-usage-record
      --meter-type-code string         How this record updates the running total (one of: absolute, delta)
      --meter-usage-timestamp string   When this event happened (RFC 3339 timestamp, e.g. 2026-01-31T09:00:00Z)
      --meter-value string             The value of usage to record (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde billing](kinde_billing.md)	 - Manage billing agreements, entitlements and usage

