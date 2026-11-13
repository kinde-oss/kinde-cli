## kinde users create-identity

Create identity

### Synopsis

Creates an identity for a user.

Identities cannot be added to a user managed by directory sync
(`USER_MANAGED_BY_DIRECTORY_SYNC`); user and role changes must be made in the
upstream identity provider.

Adding an `enterprise` identity pre-provisions the user for that connection and
replaces all of their other identities. A user with an enterprise identity
cannot be given any other identity (`CANNOT_ADD_ENTERPRISE_USER_IDENTITY`).

Required scope: create:user_identities

API: POST /api/v1/users/{user_id}/identities

```
kinde users create-identity [flags]
```

### Options

```
      --body string               Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --connection-id string      The social or enterprise connection ID, only required when identity type is 'social' or 'enterprise'
  -h, --help                      help for create-identity
      --phone-country-id string   The country code for the phone number, only required when identity type is 'phone'
      --type string               The identity type (one of: email, username, phone, enterprise, social)
      --user-id string            The user's ID (required)
      --value string              The identity value, which depends on the type
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

