## kinde account

Log in as a user and see what they can access (Account API)

### Synopsis

Log in as a user of your Kinde business through the browser, and use the
Account API to see what that user can access: their permissions, roles,
feature flags, entitlements and properties.

This is useful for testing how your application will see a user. A user login
can't call the Management API; use 'kinde login' with an M2M application for
that.

### Examples

```
  kinde account login
  kinde account permissions
  kinde account feature-flags | jq '.data.feature_flags'
```

### Options

```
  -h, --help   help for account
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde](kinde.md)	 - Kinde CLI
* [kinde account entitlement](kinde_account_entitlement.md)	 - Get entitlement
* [kinde account entitlements](kinde_account_entitlements.md)	 - Get entitlements
* [kinde account feature-flags](kinde_account_feature-flags.md)	 - Get feature flags
* [kinde account login](kinde_account_login.md)	 - Log in as a user through the browser
* [kinde account permissions](kinde_account_permissions.md)	 - Get permissions
* [kinde account portal-link](kinde_account_portal-link.md)	 - Get self-serve portal link
* [kinde account profile](kinde_account_profile.md)	 - Get user profile
* [kinde account properties](kinde_account_properties.md)	 - Get properties
* [kinde account roles](kinde_account_roles.md)	 - Get roles
* [kinde account token](kinde_account_token.md)	 - Print the access token of the browser login

