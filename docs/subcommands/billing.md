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
  - [Usage Tracking and Metering](#usage-tracking-and-metering)
  - [Entitlements Management](#entitlements-management)
  - [Complete Billing Workflow](#complete-billing-workflow)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create_agreement`

Create a new billing agreement with a customer for a specific plan, including options for prorating and immediate invoicing.

**Usage:**
```bash
kinde manage billing create_agreement [flags]
```

**Flags:**
- `--customer_id string` - Customer ID for the agreement
- `--is_invoice_now` - Invoice immediately upon agreement creation
- `--is_prorate` - Enable prorating for mid-cycle plan changes
- `--plan_code string` - Plan code for the agreement

**Example:**
```bash
# Create a basic billing agreement
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan"

# Create agreement with prorating enabled
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan" --is_prorate

# Create agreement with immediate invoicing
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan" --is_invoice_now
```

### `create_meter_usage_record`

Create a meter usage record for billing tracking and usage-based billing calculations.

**Usage:**
```bash
kinde manage billing create_meter_usage_record [flags]
```

**Flags:**
- `--billing_feature_code string` - Code identifying the billing feature being tracked
- `--customer_agreement_id string` - ID of the customer agreement this usage applies to
- `--meter_value string` - The usage value to record (e.g., number of API calls, storage used)

**Example:**
```bash
# Record API usage for a customer
kinde manage billing create_meter_usage_record --customer_agreement_id "agreement_123" --billing_feature_code "api_calls" --meter_value "1000"

# Record storage usage
kinde manage billing create_meter_usage_record --customer_agreement_id "agreement_123" --billing_feature_code "storage_gb" --meter_value "5.5"
```

### `get_agreements`

Retrieve billing agreements with optional filtering by customer, feature, and pagination support.

**Usage:**
```bash
kinde manage billing get_agreements [flags]
```

**Flags:**
- `--customer_id string` - Filter agreements by specific customer ID
- `--ending_before string` - Pagination: get agreements ending before this cursor
- `--feature_code string` - Filter agreements by specific feature code
- `--page_size int` - Number of agreements to return per page
- `--starting_after string` - Pagination: get agreements starting after this cursor

**Example:**
```bash
# Get all agreements
kinde manage billing get_agreements

# Get agreements for a specific customer
kinde manage billing get_agreements --customer_id "customer_123"

# Get agreements with pagination
kinde manage billing get_agreements --page_size 10 --starting_after "cursor_123"

# Get agreements for a specific feature
kinde manage billing get_agreements --feature_code "api_calls"
```

### `get_entitlements`

Retrieve billing entitlements with optional filtering by customer and pagination support.

**Usage:**
```bash
kinde manage billing get_entitlements [flags]
```

**Flags:**
- `--customer_id string` - Filter entitlements by specific customer ID
- `--ending_before string` - Pagination: get entitlements ending before this cursor
- `--max_value string` - Filter entitlements by maximum value
- `--page_size int` - Number of entitlements to return per page
- `--starting_after string` - Pagination: get entitlements starting after this cursor

**Example:**
```bash
# Get all entitlements
kinde manage billing get_entitlements

# Get entitlements for a specific customer
kinde manage billing get_entitlements --customer_id "customer_123"

# Get entitlements with pagination
kinde manage billing get_entitlements --page_size 20 --starting_after "cursor_456"

# Get entitlements with value filtering
kinde manage billing get_entitlements --max_value "1000"
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

# Get agreements for a specific customer
kinde manage billing get_agreements --customer_id "customer_123"

# Get agreements with pagination
kinde manage billing get_agreements --page_size 10 --starting_after "cursor_123"
```

### Usage Tracking and Metering

```bash
# Record API usage for a customer
kinde manage billing create_meter_usage_record --customer_agreement_id "agreement_123" --billing_feature_code "api_calls" --meter_value "1000"

# Record storage usage
kinde manage billing create_meter_usage_record --customer_agreement_id "agreement_123" --billing_feature_code "storage_gb" --meter_value "5.5"

# Record monthly active users
kinde manage billing create_meter_usage_record --customer_agreement_id "agreement_123" --billing_feature_code "mau" --meter_value "250"
```

### Entitlements Management

```bash
# Check all entitlements
kinde manage billing get_entitlements

# Get entitlements for a specific customer
kinde manage billing get_entitlements --customer_id "customer_123"

# Get entitlements with value filtering
kinde manage billing get_entitlements --max_value "1000"

# Get entitlements with pagination
kinde manage billing get_entitlements --page_size 20 --starting_after "cursor_456"
```

### Complete Billing Workflow

```bash
# 1. Create a billing agreement for a new customer
kinde manage billing create_agreement --customer_id "customer_123" --plan_code "premium_plan" --is_prorate

# 2. Record initial usage
kinde manage billing create_meter_usage_record --customer_agreement_id "agreement_123" --billing_feature_code "api_calls" --meter_value "100"

# 3. Check customer entitlements
kinde manage billing get_entitlements --customer_id "customer_123"

# 4. Monitor agreements
kinde manage billing get_agreements --customer_id "customer_123"

# 5. Record ongoing usage
kinde manage billing create_meter_usage_record --customer_agreement_id "agreement_123" --billing_feature_code "api_calls" --meter_value "500"
```

## Best Practices

1. **Agreement Management**: Use clear customer IDs and plan codes for easy tracking and management
2. **Prorating**: Enable prorating for mid-cycle plan changes to ensure fair billing
3. **Usage Records**: Create usage records regularly and consistently for accurate billing calculations
4. **Entitlements Monitoring**: Regularly check entitlements to ensure proper access and feature availability
5. **Pagination**: Use pagination parameters when dealing with large datasets to improve performance
6. **Customer Filtering**: Use customer ID filtering to focus on specific customer billing data
7. **Feature Codes**: Use consistent billing feature codes across your application for accurate metering
8. **Immediate Invoicing**: Use immediate invoicing for urgent billing situations or new customer onboarding

## Troubleshooting

- **Customer Not Found**: Verify the customer ID exists in your system and is correctly formatted
- **Plan Issues**: Ensure the plan code is valid, active, and available for the customer
- **Permission Errors**: Check that your M2M application has billing management scopes
- **Invalid Agreement ID**: Verify the customer agreement ID exists and is associated with the correct customer
- **Feature Code Issues**: Ensure billing feature codes are properly configured and match your billing setup
- **Pagination Errors**: Check that cursor values are valid and not expired when using pagination
- **Usage Recording Failures**: Verify that meter values are in the correct format and within acceptable ranges
- **Entitlement Access**: Ensure the customer has valid agreements before checking entitlements

## Related Commands

- `kinde manage organizations` - Manage organizations that may have billing agreements
- `kinde manage applications` - Manage applications that may be subject to billing
