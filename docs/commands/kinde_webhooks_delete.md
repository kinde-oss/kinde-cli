## kinde webhooks delete

Delete Webhook

### Synopsis

Delete webhook

Required scope: delete:webhooks

API: DELETE /api/v1/webhooks/{webhook_id}

This can't be undone, so you'll be asked to confirm. Pass --yes to skip the
question, e.g. in scripts.

```
kinde webhooks delete [flags]
```

### Options

```
  -h, --help                help for delete
      --webhook-id string   The webhook id (required)
  -y, --yes                 Don't ask for confirmation; required when not running in a terminal
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde webhooks](kinde_webhooks.md)	 - Manage webhooks

