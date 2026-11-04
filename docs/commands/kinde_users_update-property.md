## kinde users update-property

Update Property value

### Synopsis

Update property value.

Required scope: update:user_properties

API: PUT /api/v1/users/{user_id}/properties/{property_key}

```
kinde users update-property [flags]
```

### Options

```
  -h, --help                  help for update-property
      --property-key string   The identifier for the property (required)
      --user-id string        The identifier for the user (required)
      --value string          The new property value. Send an empty string to unset the property, so it reads back as not set (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde users](kinde_users.md)	 - Manage users, their identities, sessions and MFA

