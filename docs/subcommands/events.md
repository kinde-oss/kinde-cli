# `kinde manage events`

The `kinde manage events` CLI provides various subcommands to manage events in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Events commands allow you to retrieve and manage events in your Kinde environment. These operations are used to access audit logs, user activities, and system events.

## Available Commands

### `get`

Retrieve details of a specific event.

**Usage:**
```bash
kinde manage events get [flags]
```

**Flags:**
- `--event_id string` - Event ID to retrieve
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage events get --event_id "event_123"
```

### `get_types`

Retrieve all available event types.

**Usage:**
```bash
kinde manage events get_types [flags]
```

**Flags:**
- `-h, --help` - Show help information

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against

**Example:**
```bash
kinde manage events get_types
```

## Common Usage Patterns

### Event Management

```bash
# Get all available event types
kinde manage events get_types

# Get specific event details
kinde manage events get --event_id "event_123"
```

## Best Practices

1. **Event Monitoring**: Regularly review events for security and compliance
2. **Event Types**: Understand the different event types available in your system
3. **Retention**: Be aware of event retention policies
4. **Filtering**: Use appropriate filters when retrieving events

## Troubleshooting

- **Event Not Found**: Verify the event ID exists and is accessible
- **Permission Errors**: Ensure your M2M application has event access scopes

## Related Commands

- `kinde manage users` - User-related events
- `kinde manage organizations` - Organization-related events
