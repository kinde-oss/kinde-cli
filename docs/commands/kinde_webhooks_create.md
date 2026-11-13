## kinde webhooks create

Create a Webhook

### Synopsis

Create a webhook

Required scope: create:webhooks

API: POST /api/v1/webhooks

```
kinde webhooks create [flags]
```

### Options

```
      --body string           Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --description string    The webhook description
      --endpoint string       The webhook endpoint url (required)
      --event-types strings   Array of event type keys (comma-separated list; required)
  -h, --help                  help for create
      --name string           The webhook name (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde webhooks](kinde_webhooks.md)	 - Manage webhooks

