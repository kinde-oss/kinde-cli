## kinde users update-feature-flag-override

Update User Feature Flag Override

### Synopsis

Update user feature flag override.

Required scope: update:user_feature_flags

API: PATCH /api/v1/users/{user_id}/feature_flags/{feature_flag_key}

```
kinde users update-feature-flag-override [flags]
```

### Options

```
      --feature-flag-key string   The identifier for the feature flag (required)
  -h, --help                      help for update-feature-flag-override
      --user-id string            The identifier for the user (required)
      --value string              Override value (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

