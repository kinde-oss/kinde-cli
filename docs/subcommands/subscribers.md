# Subscribers Management

## Overview

Subscribers manage webhook and notification recipients in your Kinde environment. They receive event notifications when specific actions occur in your system, enabling real-time integration with external systems and automated workflows. This command group allows you to create, manage, and monitor subscribers for webhook notifications.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create](#create)
  - [get](#get)
  - [get_all](#get_all)
- [Common Usage Patterns](#common-usage-patterns)
  - [Setting Up Webhook Subscribers](#setting-up-webhook-subscribers)
  - [Managing Subscriber Lifecycle](#managing-subscriber-lifecycle)
- [Webhook Integration](#webhook-integration)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create`

Create a new subscriber for webhook notifications. Subscribers receive event notifications when specific actions occur in your Kinde environment.

**Usage:**
```bash
kinde manage subscribers create [flags]
```

**Flags:**
- `--first_name string` - First name of the subscriber

**Example:**
```bash
kinde manage subscribers create --first_name "John"
```

### `get`

Retrieve detailed information about a specific subscriber.

**Usage:**
```bash
kinde manage subscribers get [flags]
```

**Flags:**
- `--subscriber_id string` - ID of the subscriber to retrieve

**Example:**
```bash
kinde manage subscribers get --subscriber_id "sub_123456789"
```

### `get_all`

List all subscribers in your Kinde environment with pagination support.

**Usage:**
```bash
kinde manage subscribers get_all [flags]
```

**Flags:**
- `--next_token string` - Token for pagination to get next page of results
- `--page_size int` - Number of subscribers to return per page

**Example:**
```bash
kinde manage subscribers get_all --page_size 20
```

## Common Usage Patterns

### Setting Up Webhook Subscribers

Creating subscribers for different notification scenarios:

```bash
# Create subscribers for different teams
kinde manage subscribers create --first_name "John"
kinde manage subscribers create --first_name "Sarah"
kinde manage subscribers create --first_name "Mike"

# List all subscribers to verify creation
kinde manage subscribers get_all --page_size 50

# Get details for a specific subscriber
kinde manage subscribers get --subscriber_id "sub_123456789"
```

### Managing Subscriber Lifecycle

Regular management and monitoring of subscribers:

```bash
# List all subscribers with pagination
kinde manage subscribers get_all --page_size 20

# Get next page of results
kinde manage subscribers get_all --page_size 20 --next_token "next_page_token"

# Retrieve specific subscriber details
kinde manage subscribers get --subscriber_id "sub_123456789"

# Monitor subscriber activity and status
kinde manage subscribers get_all
```

## Webhook Integration

### Subscriber Configuration

Subscribers are typically configured to receive notifications for specific events:

1. **User Events** - User registration, login, profile updates
2. **Organization Events** - Organization creation, updates, user additions
3. **Application Events** - Application configuration changes
4. **Security Events** - Authentication failures, suspicious activity

### Event Notification Flow

```bash
# 1. Create subscriber
kinde manage subscribers create --first_name "WebhookHandler"

# 2. Configure webhook endpoint (using webhooks management)
kinde manage webhooks create --name "User Events" --endpoint "https://api.example.com/webhooks/users" --description "User event notifications"

# 3. Monitor subscriber status
kinde manage subscribers get_all

# 4. Retrieve subscriber details for troubleshooting
kinde manage subscribers get --subscriber_id "sub_123456789"
```

## Best Practices

### Subscriber Management

1. **Naming Conventions**
   - Use descriptive names for subscribers
   - Include team or system information in names
   - Maintain consistent naming patterns

2. **Subscriber Organization**
   - Group subscribers by team or system
   - Document subscriber purposes and responsibilities
   - Keep subscriber information up to date

3. **Monitoring and Maintenance**
   - Regularly review subscriber lists
   - Monitor subscriber activity and health
   - Remove inactive or obsolete subscribers

### Webhook Security

1. **Endpoint Security**
   - Use HTTPS endpoints for webhook notifications
   - Implement proper authentication and authorization
   - Validate webhook signatures and payloads

2. **Error Handling**
   - Implement retry logic for failed webhook deliveries
   - Monitor webhook delivery success rates
   - Set up alerts for webhook failures

3. **Rate Limiting**
   - Implement rate limiting on webhook endpoints
   - Handle high-volume event scenarios
   - Monitor webhook performance and latency

## Troubleshooting

### Common Issues

**Subscriber Creation Fails**
- Ensure required fields are provided
- Check that subscriber names are valid
- Verify you have the necessary permissions to create subscribers

**Subscriber Not Found**
- Double-check the subscriber ID is correct
- Ensure the subscriber exists in your current environment
- Verify you're using the correct domain/environment

**Webhook Delivery Issues**
- Check webhook endpoint availability and configuration
- Verify endpoint authentication and authorization
- Monitor webhook delivery logs and error messages

**Pagination Issues**
- Use the `next_token` from previous responses for pagination
- Adjust `page_size` if you're getting too many or too few results
- Handle pagination properly in automated scripts

### Getting Help

For additional help with subscribers:

```bash
# Get help for the subscribers command group
kinde manage subscribers --help

# Get help for a specific command
kinde manage subscribers create --help
kinde manage subscribers get --help
```

## Related Commands

- **Webhooks Management** (`kinde manage webhooks`) - Configure webhook endpoints for subscribers
- **Events Management** (`kinde manage events`) - Monitor events that trigger notifications
- **Users Management** (`kinde manage users`) - Manage user-related events and notifications
- **Organizations Management** (`kinde manage organizations`) - Manage organization-related events

### Workflow Integration

Subscribers work in conjunction with:

1. **Webhooks** - Configure endpoints to receive notifications
2. **Events** - Monitor system events that trigger notifications
3. **Applications** - Integrate with application-specific events
4. **External Systems** - Connect to external services and APIs

Example workflow:
```bash
# 1. Create subscriber
kinde manage subscribers create --first_name "NotificationHandler"

# 2. Create webhook endpoint
kinde manage webhooks create --name "User Events" --endpoint "https://api.example.com/webhooks" --description "User event notifications"

# 3. Monitor events
kinde manage events get_types

# 4. Verify subscriber configuration
kinde manage subscribers get_all
```

### Integration with External Systems

Subscribers enable integration with:

1. **CRM Systems** - Sync user and organization data
2. **Analytics Platforms** - Track user behavior and events
3. **Marketing Tools** - Trigger campaigns based on user actions
4. **Monitoring Systems** - Alert on system events and issues

### Event-Driven Architecture

Subscribers support event-driven architecture patterns:

```bash
# Set up subscribers for different event types
kinde manage subscribers create --first_name "UserEventHandler"
kinde manage subscribers create --first_name "OrgEventHandler"
kinde manage subscribers create --first_name "SecurityEventHandler"

# Configure webhooks for different event categories
kinde manage webhooks create --name "User Events" --endpoint "https://api.example.com/users"
kinde manage webhooks create --name "Org Events" --endpoint "https://api.example.com/organizations"
kinde manage webhooks create --name "Security Events" --endpoint "https://api.example.com/security"
```

### Monitoring and Analytics

Use subscribers for monitoring and analytics:

```bash
# Create subscribers for monitoring
kinde manage subscribers create --first_name "MonitoringSystem"
kinde manage subscribers create --first_name "AnalyticsEngine"

# Set up webhooks for monitoring endpoints
kinde manage webhooks create --name "System Monitoring" --endpoint "https://monitoring.example.com/webhooks"
kinde manage webhooks create --name "Analytics" --endpoint "https://analytics.example.com/webhooks"

# Monitor subscriber activity
kinde manage subscribers get_all --page_size 100
```
