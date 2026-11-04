## kinde webhooks update

Update a Webhook

### Synopsis

Update a webhook

Required scope: update:webhooks

API: PATCH /api/v1/webhooks/{webhook_id}

```
kinde webhooks update [flags]
```

### Options

```
      --body string           Request body as JSON: inline, @file.json, or - for stdin. Other flags override its fields
      --description string    The webhook description
      --event-types strings   Array of event type keys (comma-separated list; required)
  -h, --help                  help for update
      --name string           The webhook name
      --webhook-id string     The webhook id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde webhooks](kinde_webhooks.md)	 - Manage webhooks

