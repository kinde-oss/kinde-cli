# `kinde manage billing`

The `kinde manage billing` CLI provides various subcommands to manage billing operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Billing commands allow you to manage billing agreements, usage records, and entitlements in your Kinde environment. These operations are typically used for subscription management and usage tracking.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create_agreement](#create_agreement)
  - [create_meter_usage_record](#create_meter_usage_record)
  - [get_agreements](#get_agreements)
  - [get_entitlements](#get_entitlements)
- [Common Usage Patterns](#common-usage-patterns)
  - [Managing Billing Agreements](#managing-billing-agreements)
  - [Usage Tracking](#usage-tracking)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create_agreement`

Create a new billing agreement.

**Usage:**
```bash
kinde manage billing create_agreement [flags]
```

**Flags:**
- `--customer_id string` - Customer ID for the agreement
- `--is_invoice_now` - Invoice immediately
- `--is_prorate` - Enable prorating
- `--plan_code string` - Plan code for the agreement

**Example:**
```bash
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan" --is_prorate
```

### `create_meter_usage_record`

Create a meter usage record for billing tracking.

**Usage:**
```bash
kinde manage billing create_meter_usage_record [flags]
```

**Flags:**

**Example:**
```bash
kinde manage billing create_meter_usage_record
```

### `get_agreements`

Retrieve billing agreements.

**Usage:**
```bash
kinde manage billing get_agreements [flags]
```

**Flags:**

**Example:**
```bash
kinde manage billing get_agreements
```

### `get_entitlements`

Retrieve billing entitlements.

**Usage:**
```bash
kinde manage billing get_entitlements [flags]
```

**Flags:**

**Example:**
```bash
kinde manage billing get_entitlements
```

## Common Usage Patterns

### Managing Billing Agreements

```bash
# Create a new billing agreement
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan"

# Create agreement with prorating
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan" --is_prorate

# Create agreement with immediate invoicing
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan" --is_invoice_now

# List all agreements
kinde manage billing get_agreements
```

### Usage Tracking

```bash
# Create usage records for metering
kinde manage billing create_meter_usage_record

# Check entitlements
kinde manage billing get_entitlements
```

## Best Practices

1. **Agreement Management**: Use clear customer IDs and plan codes for easy tracking
2. **Prorating**: Enable prorating for mid-cycle plan changes
3. **Usage Records**: Create usage records regularly for accurate billing
4. **Entitlements**: Regularly check entitlements to ensure proper access

## Troubleshooting

- **Customer Not Found**: Verify the customer ID exists in your system
- **Plan Issues**: Ensure the plan code is valid and active
- **Permission Errors**: Check that your M2M application has billing scopes

## Related Commands

- `kinde manage organizations` - Manage organizations that may have billing agreements
- `kinde manage applications` - Manage applications that may be subject to billing
