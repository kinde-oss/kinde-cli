# Properties Management

## Overview

Properties allow you to store additional metadata and configuration data for users, organizations, and applications in your Kinde environment. They provide a flexible way to extend the default data model with custom fields, enabling you to capture business-specific information and configure application behavior dynamically.

**Global Flags:**
- `--domain value` - Kinde domain (environment) to work against
- `-h, --help` - Show help information

## Table of Contents

- [Overview](#overview)
- [Available Commands](#available-commands)
  - [create](#create)
  - [delete](#delete)
  - [get_all](#get_all)
  - [update](#update)
- [Common Usage Patterns](#common-usage-patterns)
  - [Creating Custom Properties](#creating-custom-properties)
  - [Organizing Properties with Categories](#organizing-properties-with-categories)
  - [Managing Property Lifecycle](#managing-property-lifecycle)
- [Property Types and Contexts](#property-types-and-contexts)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)
- [Related Commands](#related-commands)

## Available Commands

### `create`

Create a new custom property with specified type, context, and validation rules.

**Usage:**
```bash
kinde manage properties create [flags]
```

**Flags:**
- `--category_id string` - ID of the category to group this property
- `--context string` - Context where the property can be used (user, organization, application)
- `--description string` - Description of the property and its purpose
- `--key string` - Unique key identifier for the property
- `--name string` - Display name for the property
- `--type string` - Data type of the property (string, number, boolean, etc.)

**Example:**
```bash
kinde manage properties create --name "Department" --key "department" --type "string" --context "user" --description "User's department within the organization"
```

### `delete`

Delete a custom property permanently from your Kinde environment.

**Usage:**
```bash
kinde manage properties delete [flags]
```

**Flags:**
- `--property_id string` - ID of the property to delete

**Example:**
```bash
kinde manage properties delete --property_id "prop_123456789"
```

### `get_all`

List all custom properties in your Kinde environment with pagination support.

**Usage:**
```bash
kinde manage properties get_all [flags]
```

**Flags:**
- `--ending_before string` - Cursor for pagination to get results before this point
- `--page_size int` - Number of properties to return per page
- `--starting_after string` - Cursor for pagination to get results after this point

**Example:**
```bash
kinde manage properties get_all --page_size 20
```

### `update`

Update an existing custom property's name, description, or category.

**Usage:**
```bash
kinde manage properties update [flags]
```

**Flags:**
- `--category_id string` - New category ID for the property
- `--description string` - New description for the property
- `--name string` - New display name for the property
- `--property_id string` - ID of the property to update

**Example:**
```bash
kinde manage properties update --property_id "prop_123456789" --name "User Department" --description "Updated description for department property"
```

## Common Usage Patterns

### Creating Custom Properties

When setting up custom data fields for your application, you'll typically create properties for different contexts:

```bash
# Create user-specific properties
kinde manage properties create --name "Employee ID" --key "employee_id" --type "string" --context "user" --description "Unique employee identifier"
kinde manage properties create --name "Job Title" --key "job_title" --type "string" --context "user" --description "User's current job title"
kinde manage properties create --name "Hire Date" --key "hire_date" --type "string" --context "user" --description "Date when user was hired"

# Create organization-specific properties
kinde manage properties create --name "Industry" --key "industry" --type "string" --context "organization" --description "Organization's industry sector"
kinde manage properties create --name "Company Size" --key "company_size" --type "number" --context "organization" --description "Number of employees in the organization"

# Create application-specific properties
kinde manage properties create --name "API Version" --key "api_version" --type "string" --context "application" --description "API version used by the application"
kinde manage properties create --name "Environment" --key "environment" --type "string" --context "application" --description "Deployment environment (dev, staging, prod)"
```

### Organizing Properties with Categories

Use categories to group related properties for better organization:

```bash
# First, create a category (using categories management)
kinde manage categories create --name "User Profile" --description "Properties related to user profile information"

# Then create properties with the category
kinde manage properties create --name "First Name" --key "first_name" --type "string" --context "user" --category_id "cat_123" --description "User's first name"
kinde manage properties create --name "Last Name" --key "last_name" --type "string" --context "user" --category_id "cat_123" --description "User's last name"
kinde manage properties create --name "Phone Number" --key "phone" --type "string" --context "user" --category_id "cat_123" --description "User's contact phone number"
```

### Managing Property Lifecycle

Regular maintenance and updates of your property system:

```bash
# List all properties to review current setup
kinde manage properties get_all --page_size 50

# Update a property that needs clarification
kinde manage properties update --property_id "prop_123" --description "Updated description for better clarity"

# Move a property to a different category
kinde manage properties update --property_id "prop_456" --category_id "cat_new_category"

# Remove unused properties
kinde manage properties delete --property_id "prop_unused_789"
```

## Property Types and Contexts

### Supported Property Types

Properties support various data types to accommodate different use cases:

- **`string`** - Text values (names, descriptions, IDs)
- **`number`** - Numeric values (counts, scores, measurements)
- **`boolean`** - True/false values (flags, toggles)
- **`date`** - Date and time values
- **`json`** - Complex structured data

### Property Contexts

Properties can be associated with different entities in your Kinde environment:

- **`user`** - Properties that belong to individual users
- **`organization`** - Properties that belong to organizations
- **`application`** - Properties that belong to applications

### Example Property Definitions

```bash
# String properties for user data
kinde manage properties create --name "Bio" --key "bio" --type "string" --context "user" --description "User's biography or description"

# Number properties for metrics
kinde manage properties create --name "Login Count" --key "login_count" --type "number" --context "user" --description "Number of times user has logged in"

# Boolean properties for flags
kinde manage properties create --name "Email Verified" --key "email_verified" --type "boolean" --context "user" --description "Whether user's email is verified"

# JSON properties for complex data
kinde manage properties create --name "Preferences" --key "preferences" --type "json" --context "user" --description "User's application preferences as JSON"
```

## Best Practices

### Property Design and Naming

1. **Use Consistent Naming Conventions**
   - Use snake_case for property keys (e.g., `first_name`, `employee_id`)
   - Use descriptive but concise names
   - Avoid abbreviations unless they're widely understood

2. **Plan Property Structure**
   - Design properties that can be reused across different contexts
   - Consider data validation requirements
   - Plan for future extensibility

3. **Document Properties Thoroughly**
   - Always provide clear descriptions
   - Document expected data formats
   - Include examples of valid values

### Data Management

1. **Property Lifecycle Management**
   - Regularly review and audit property usage
   - Remove unused or obsolete properties
   - Update descriptions as requirements change

2. **Category Organization**
   - Group related properties into categories
   - Use consistent category naming
   - Plan category hierarchy for complex systems

3. **Data Validation**
   - Choose appropriate property types
   - Consider data format requirements
   - Plan for data migration if types change

### Security and Privacy

1. **Sensitive Data Handling**
   - Be cautious with sensitive information in properties
   - Consider encryption for sensitive data
   - Follow data privacy regulations

2. **Access Control**
   - Understand which properties are accessible to applications
   - Plan property visibility and access patterns
   - Consider user consent for data collection

## Troubleshooting

### Common Issues

**Property Creation Fails**
- Ensure the property key is unique and follows naming conventions
- Check that required fields (name, key, type, context) are provided
- Verify the context value is valid (user, organization, application)
- Ensure you have the necessary permissions to create properties

**Property Not Found**
- Double-check the property ID is correct
- Ensure the property exists in your current environment
- Verify you're using the correct domain/environment

**Pagination Issues**
- Use the `starting_after` and `ending_before` cursors properly
- Adjust `page_size` if you're getting too many or too few results
- Handle pagination properly in automated scripts

**Category Assignment Issues**
- Ensure the category ID exists and is valid
- Check that the category is accessible in your environment
- Verify category permissions

### Getting Help

For additional help with properties:

```bash
# Get help for the properties command group
kinde manage properties --help

# Get help for a specific command
kinde manage properties create --help
kinde manage properties update --help
```

## Related Commands

- **Categories Management** (`kinde manage categories`) - Organize properties into categories
- **Users Management** (`kinde manage users`) - Set and retrieve user property values
- **Organizations Management** (`kinde manage organizations`) - Set and retrieve organization property values
- **Applications Management** (`kinde manage applications`) - Set and retrieve application property values

### Workflow Integration

Properties are typically used in conjunction with:

1. **Categories** - Group and organize related properties
2. **Users** - Store custom user data and preferences
3. **Organizations** - Store organization-specific configuration
4. **Applications** - Store application settings and metadata

Example workflow:
```bash
# 1. Create a category for user profile data
kinde manage categories create --name "User Profile" --description "User profile information"

# 2. Create properties in the category
kinde manage properties create --name "Department" --key "department" --type "string" --context "user" --category_id "cat_123" --description "User's department"

# 3. Set property values for users
kinde manage users update_property --user_id "user_456" --property_key "department" --value "Engineering"

# 4. Retrieve property values
kinde manage users get_property_values --user_id "user_456"
```

### Integration with Applications

Properties enable applications to:
- Store custom user preferences and settings
- Configure application behavior dynamically
- Collect and store business-specific data
- Integrate with external systems through custom fields
- Support multi-tenant configurations
