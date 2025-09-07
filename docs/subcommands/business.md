# `kinde manage business`

The `kinde manage business` CLI provides various subcommands to manage business operations in Kinde. Below is a detailed list of available commands and their usage.

## Overview

Business commands allow you to manage business settings and configuration in your Kinde environment. These operations are typically used for configuring business-level settings and preferences.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [get](#get)
  - [update](#update)
- [Common Usage Patterns](#common-usage-patterns)
  - [Managing Business Settings](#managing-business-settings)
  - [Complete Business Setup Workflow](#complete-business-setup-workflow)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `get`

Retrieve current business settings and configuration including business name, contact information, branding settings, and other business-level preferences.

**Usage:**
```bash
kinde manage business get [flags]
```

**Flags:**
- No additional flags required for this command

**Example:**
```bash
kinde manage business get
```

### `update`

Update business settings and configuration including business information, contact details, branding preferences, and compliance settings.

**Usage:**
```bash
kinde manage business update [flags]
```

**Flags:**
- `--business_name string` - Business name
- `--email string` - Business contact email address
- `--industry_key string` - Industry classification key (use `kinde manage industries get_all` to see available options)
- `--is_click_wrap` - Enable click-wrap agreement for terms and privacy
- `--is_show_kinde_branding` - Show Kinde branding in authentication flows
- `--kinde_perk_code string` - Kinde perk code for special features
- `--phone string` - Business contact phone number
- `--privacy_url string` - URL to privacy policy
- `--terms_url string` - URL to terms of service
- `--timezone_key string` - Business timezone (use `kinde manage timezones get_all` to see available options)

**Example:**
```bash
# Update basic business information
kinde manage business update --business_name "My Company" --email "admin@mycompany.com"

# Update business with industry and timezone
kinde manage business update --business_name "Tech Corp" --industry_key "technology" --timezone_key "UTC"

# Update branding and compliance settings
kinde manage business update --is_show_kinde_branding --privacy_url "https://mycompany.com/privacy" --terms_url "https://mycompany.com/terms"
```

## Common Usage Patterns

### Managing Business Settings

```bash
# Get current business settings
kinde manage business get

# Update basic business information
kinde manage business update --business_name "My Company" --email "admin@mycompany.com"

# Configure business with industry and timezone
kinde manage business update --industry_key "technology" --timezone_key "America/New_York"

# Set up compliance and branding
kinde manage business update --privacy_url "https://mycompany.com/privacy" --terms_url "https://mycompany.com/terms" --is_click_wrap

# Update contact information
kinde manage business update --phone "+1-555-123-4567" --email "contact@mycompany.com"
```

### Complete Business Setup Workflow

```bash
# 1. Get available industries and timezones
kinde manage industries get_all
kinde manage timezones get_all

# 2. Set up complete business profile
kinde manage business update \
  --business_name "My Tech Company" \
  --email "admin@mytechcompany.com" \
  --phone "+1-555-123-4567" \
  --industry_key "technology" \
  --timezone_key "America/New_York" \
  --privacy_url "https://mytechcompany.com/privacy" \
  --terms_url "https://mytechcompany.com/terms" \
  --is_click_wrap

# 3. Verify the configuration
kinde manage business get
```

## Best Practices

1. **Complete Setup**: Use the complete business setup workflow to configure all necessary business information at once
2. **Industry Classification**: Always set an appropriate industry key to enable industry-specific features and reporting
3. **Timezone Configuration**: Set the correct timezone for your business to ensure proper time handling across your applications
4. **Compliance URLs**: Always provide privacy and terms URLs when enabling click-wrap agreements
5. **Regular Review**: Periodically review business settings using `kinde manage business get` to ensure they're current
6. **Backup Settings**: Keep track of important business configuration changes for audit purposes
7. **Branding Consistency**: Decide whether to show Kinde branding based on your brand requirements

## Troubleshooting

- **Permission Errors**: Ensure your M2M application has business management scopes
- **Invalid Industry Key**: Use `kinde manage industries get_all` to see valid industry options
- **Invalid Timezone Key**: Use `kinde manage timezones get_all` to see valid timezone options
- **Configuration Issues**: Verify that business settings are properly configured by running `kinde manage business get`
- **Missing Required Fields**: Some business settings may be required depending on your Kinde plan
- **URL Validation**: Ensure privacy and terms URLs are accessible and properly formatted

## Related Commands

- `kinde manage organizations` - Manage organizations within the business
- `kinde manage applications` - Manage applications for the business
