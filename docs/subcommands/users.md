# `kinde manage users`

The `kinde manage users` CLI provides various subcommands to manage user operations. Below is a list of available subcommands and their descriptions:

| Subcommand                      | Description                                                      |
|----------------------------------|------------------------------------------------------------------|
| `create`                        | Create a new user                                                |
| `create_identity`               | Add a new identity to a user                                     |
| `delete`                        | Delete a user                                                    |
| `delete_identity`               | Remove an identity from a user                                   |
| `delete_sessions`               | Delete all sessions for a user                                   |
| [`get_all`](#get_all)           | Retrieve all users                                               |
| `get_data`                      | Get detailed data for a user                                     |
| `get_identities`                | List all identities associated with a user                       |
| `get_mfa`                       | Get multi-factor authentication settings for a user              |
| `get_permissions`               | List permissions assigned to a user                              |
| `get_profile_v2`                | Retrieve the user's profile (v2 format)                          |
| `get_properties`                | List custom properties for a user                                |
| `get_property_values`           | Get values of specific properties for a user                     |
| `get_roles`                     | List roles assigned to a user                                    |
| `get_sessions`                  | List active sessions for a user                                  |
| `refresh_claims`                | Refresh authentication claims for a user                         |
| `replace_mfa`                   | Replace multi-factor authentication settings for a user          |
| `reset_mfa`                     | Reset multi-factor authentication for a user                     |
| `reset_mfa_all`                 | Reset multi-factor authentication for all users                  |
| `search`                        | Search for users by criteria                                     |
| `set_password`                  | Set or update a user's password                                  |
| `update`                        | Update user details                                              |
| `update_feature_flag_override`  | Update feature flag overrides for a user                         |
| `update_identity`               | Update an identity for a user                                    |
| `update_properties`             | Update multiple properties for a user                            |
| `update_property`               | Update a specific property for a user                            |

## Usage

Each subcommand can be used as follows:

```sh
kinde manage users <subcommand> [options]
```

For detailed help on a specific subcommand, use:

```sh
kinde manage users <subcommand> --help
```

### `get_all`

Flags for the `get_all` subcommand:

- `--email string` — Filter results by user email address.
- `--expand string` — Specify related resources to expand in the response.
- `--has_organization` — Filter results to include only users with an organization.
- `--next_token string` — Token for fetching the next page of results.
- `--page_size int` — Number of results to return per page.
- `--phone string` — Filter results by user phone number.
- `--user_id string` — Filter results by user ID.
- `--username string` — Filter results by username.
- `-h, --help` — Display help information for the get_all command.
