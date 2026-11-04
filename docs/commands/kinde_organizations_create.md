## kinde organizations create

Create organization

### Synopsis

Create a new organization. To learn more read about [multi tenancy using
organizations](https://docs.kinde.com/build/organizations/multi-tenancy-using-organizations/)

Required scope: create:organizations

API: POST /api/v1/organization

```
kinde organizations create [flags]
```

### Options

```
      --background-color string         The organization's brand settings - background color
      --background-color-dark string    The organization's brand settings - dark mode background color
      --billing-email string            The email address used for billing purposes for the organization. Required when is_create_billing_customer is true
      --billing-plan-code string        Code of a published organization billing plan to assign to the new billing customer
      --body string                     Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --button-color string             The organization's brand settings - button color
      --button-color-dark string        The organization's brand settings - dark mode button color
      --button-text-color string        The organization's brand settings - button text color
      --button-text-color-dark string   The organization's brand settings - dark mode button text color
      --external-id string              The organization's external identifier - commonly used when migrating from or mapping to other systems
      --feature-flags string            The organization's feature flag settings (JSON value)
      --handle string                   A unique handle for the organization - can be used for dynamic callback urls
  -h, --help                            help for create
      --is-allow-registrations          Deprecated - Use 'is_auto_membership_enabled' instead
      --is-auto-membership-enabled      If users become members of this organization when the org code is supplied during authentication
      --is-create-billing-customer      If an organization billing customer is also created for this organization
      --link-color string               The organization's brand settings - link color
      --link-color-dark string          The organization's brand settings - dark mode link color
      --name string                     The organization's name (required)
      --sender-email string             The email address that will be used in emails. Requires custom SMTP to be set up
      --sender-name string              The name of the organization that will be used in emails
      --theme-code string               The organization's brand settings - theme/mode 'light' | 'dark' | 'user_preference'
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

