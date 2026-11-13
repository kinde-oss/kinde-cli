## kinde business update

Update business

### Synopsis

Update your business details.

Required scope: update:businesses

API: PATCH /api/v1/business

```
kinde business update [flags]
```

### Options

```
      --body string              Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --business-name string     The name of the business
      --email string             The email address of the business
  -h, --help                     help for update
      --industry-key string      The key of the industry of your business. Can be retrieved from the /industries endpoint
      --is-click-wrap            Whether the business is using clickwrap agreements
      --is-show-kinde-branding   Whether the business is showing Kinde branding. Requires a paid plan
      --kinde-perk-code string   The Kinde perk code for the business
      --phone string             The phone number of the business
      --privacy-url string       The URL to the business's privacy policy
      --terms-url string         The URL to the business's terms of service
      --timezone-key string      The key of the timezone of your business. Can be retrieved from the /timezones endpoint
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde business](kinde_business.md)	 - Manage business details

