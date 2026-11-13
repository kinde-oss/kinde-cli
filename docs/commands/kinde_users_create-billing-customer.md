## kinde users create-billing-customer

Create user billing customer

### Synopsis

Creates a billing customer for a user in an organization, and assigns a
published user billing plan.

This mirrors the admin "Start billing customer" action on the user billing page.

At most one billing customer is created per user in an organization. If one
already exists, the request fails with `BILLING_CUSTOMER_EXISTS`. Concurrent
creates that both pass the existence check before either finishes may both
return success with the same customer and agreement ids (onboarding reuses the
row created under the advisory lock).

Required scope: create:user_billing_customers

API: POST /api/v1/users/{user_id}/billing_customer

```
kinde users create-billing-customer [flags]
```

### Options

```
      --billing-email string       The email address used for billing purposes for the user (required)
      --billing-plan-code string   Code of a published user billing plan to assign to the new billing customer
      --body string                Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
  -h, --help                       help for create-billing-customer
      --org-code string            The organization code the user belongs to. The billing customer is linked to this organization (required)
      --user-id string             The user's ID (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

