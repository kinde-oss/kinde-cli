## kinde organizations update

Update Organization

### Synopsis

Update an organization.

When the organization name is updated and the organization is a billing
customer, the change is also propagated to the corresponding billing customer
details.

Required scope: update:organizations

API: PATCH /api/v1/organization/{org_code}

```
kinde organizations update [flags]
```

### Options

```
      --allowed-domains strings         Domains allowed for self-sign up to this environment (comma-separated list)
      --background-color string         The organization's brand settings - background color
      --background-color-dark string    The organization's brand settings - dark mode background color
      --body string                     Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --button-color string             The organization's brand settings - button color
      --button-color-dark string        The organization's brand settings - dark mode button color
      --button-text-color string        The organization's brand settings - button text color
      --button-text-color-dark string   The organization's brand settings - dark mode button text color
      --expand string                   Additional data to include in the response. Allowed value: "billing" (one of: billing)
      --external-id string              The organization's ID
      --handle string                   The organization's handle
  -h, --help                            help for update
      --is-allow-registrations          Deprecated - Use 'is_auto_membership_enabled' instead
      --is-auto-join-domain-list        Users can sign up to this organization
      --is-auto-membership-enabled      If users become members of this organization when the org code is supplied during authentication
      --is-enable-advanced-orgs         Activate advanced organization features
      --is-enforce-mfa                  Enforce MFA for all users in this organization
      --is-suspended                    Whether to suspend or unsuspend the organization
      --link-color string               The organization's brand settings - link color
      --link-color-dark string          The organization's brand settings - dark mode link color
      --name string                     The organization's name
      --org-code string                 The identifier for the organization (required)
      --sender-email string             The email address that will be used in emails. Requires custom SMTP to be set up
      --sender-name string              The name of the organization that will be used in emails
      --theme-code string               The organization's brand settings - theme/mode (one of: light, dark, user_preference)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

