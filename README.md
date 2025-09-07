# Kinde command line utility

## Status

This is a work in progress. APIs may change without notice or be missing.

## Features

- Official CLI for interacting with Kinde APIs.
- Secure handling of credentials via OS-specific secrets store.
- Verifiable authenticity
  - single native binary for ease of maintenance and deployment
  - Ships for multiple platforms - Linux, MacOS, Windows
  - Ships for multiple architectures - x64, i386, ARM, Apple ARM
  - Each variation comes with the checksum to verify authenticity
- Secure by default. Kinde business cannot be managed via the CLI until a new M2M application is created with the management API authorized and the select scopes granted.

TODO:
- [x] Online token verification
  - [ ] Offline token verification
- [x] Support for user tokens and account API
  - [ ] Spport for account API

## Getting Started

### Installation

#### macOS with Homebrew
```bash
brew tap kinde-oss/kinde-cli
brew install kinde-cli
```

#### Windows with Scoop
```bash
scoop bucket add kinde https://github.com/kinde-oss/scoop-kinde-cli
scoop install kinde-cli
```

#### Linux

##### APT Repository (Recommended for Debian/Ubuntu)
```bash
curl -s https://pkg.kinde.com/apt/gpg.key | gpg --dearmor | sudo tee /usr/share/keyrings/kinde.gpg
sudo sh -c 'echo "deb [signed-by=/usr/share/keyrings/kinde.gpg] https://pkg.kinde.com/apt/ * *" > /etc/apt/sources.list.d/kinde.list'
sudo apt update
sudo apt install kinde-cli
```

##### Manual Package Installation
Download and install the appropriate package for your architecture:

Example uses `bash`, the following shells are also supported:
- bash
- zsh
- fish
- powershell

**ARM64 (aarch64):**
```bash
curl -L https://github.com//kinde-oss/kinde-cli/releases/latest/download/kinde-cli_linux_arm64.deb > kinde-cli_linux_arm64.deb ; sudo dpkg -i kinde-cli_linux_arm64.deb ; source <(kinde completion bash)
```

**x64 (amd64):**
```bash
curl -L https://github.com//kinde-oss/kinde-cli/releases/latest/download/kinde-cli_linux_amd64.deb > kinde-cli_linux_amd64.deb ; sudo dpkg -i kinde-cli_linux_amd64.deb ; source <(kinde completion bash)
```

#### Download pre-built binary for your architecture/OS
Download a pre-built binary from [Releases](https://github.com/kinde-oss/kinde-cli/releases).

**📋 Supported Architectures and Formats:** See [Installation Architectures](docs/installation-architectures.md) for comprehensive information about all supported operating systems, architectures, package formats, and installation methods.

#### Install from source using Go
```bash
go install github.com/kinde-oss/kinde-cli/cmd/kinde@latest
```

### Quick Start and top-level commands

`kinde help` for top-level help

`kinde help <command>` to get options for each command

## Authentication

### To login using client credentials

```bash
kinde login --domain <your Kinde business domain> \
--client_id <M2M application> \
--client_secret <M2M application secret> 
```

### To login using device authorization flow

Please configure a default application for device flow or specify `client_id`

```bash
kinde login
    --domain <your Kinde business domain>
    --client_id [optional]
```

## Check authentication status

## Terminal autocomplete

Kinde CLI comes with terminal autocompletion support for `bash`, `fish`, `powershell`, and `zsh`

You can see instructions for your shell/OS combination by executing the following

```
kinde help completion <bash|fish|powershell|ash>
```

1. Use `kinde <command>` to interact with your Kinde account.

## Usage

### Commands

- `login` – Authenticate with Kinde.
  - Multi-tenancy is supported via `--domain <your Kinde business domain>`
- `logout` – Remove authentication credentials.
- `whoami` – Show current authenticated token details.
- `manage` - provides access to management API
- `help` – Show help for commands.

### Flags

- `--help` – Show help information.

### Management API sub-commands

| Subcommand | Description |
|------------|-------------|
| [api_keys](docs/manage-commands-reference.md#1-api-keys-management-api_keys) | Manage API keys |
| [apis](docs/manage-commands-reference.md#2-apis-management-apis) | Manage APIs |
| [applications](docs/manage-commands-reference.md#3-applications-management-applications) | Manage applications |
| [billing](docs/manage-commands-reference.md#4-billing-management-billing) | Manage billing |
| [business](docs/manage-commands-reference.md#5-business-management-business) | Manage business settings |
| [categories](docs/manage-commands-reference.md#6-categories-management-categories) | Manage categories |
| [connected_apps](docs/manage-commands-reference.md#7-connected-apps-management-connected_apps) | Manage connected applications |
| [connections](docs/manage-commands-reference.md#8-connections-management-connections) | Manage connections |
| [environment_variables](docs/manage-commands-reference.md#9-environment-variables-management-environment_variables) | Manage environment variables |
| [environments](docs/manage-commands-reference.md#10-environments-management-environments) | Manage environments |
| [events](docs/manage-commands-reference.md#11-events-management-events) | Manage events |
| [feature_flags](docs/manage-commands-reference.md#12-feature-flags-management-feature_flags) | Manage feature flags |
| [industries](docs/manage-commands-reference.md#13-industries-management-industries) | Manage industries |
| [organizations](docs/manage-commands-reference.md#14-organizations-management-organizations) | Manage organizations |
| [permissions](docs/manage-commands-reference.md#15-permissions-management-permissions) | Manage permissions |
| [properties](docs/manage-commands-reference.md#16-properties-management-properties) | Manage properties |
| [roles](docs/manage-commands-reference.md#17-roles-management-roles) | Manage roles |
| [subscribers](docs/manage-commands-reference.md#18-subscribers-management-subscribers) | Manage subscribers |
| [timezones](docs/manage-commands-reference.md#19-timezones-management-timezones) | Manage timezones |
| [users](docs/subcommands/users.md) | Manage users |
| [webhooks](docs/manage-commands-reference.md#21-webhooks-management-webhooks) | Manage webhooks |

**📖 Complete Reference:** See [Manage Commands Reference](docs/manage-commands-reference.md) for detailed documentation of all subcommands and their available operations.

## Configuration

Configuration is stored in `$HOME/.kinde/config.json` by default for *nix-based systems, `<user profile>\.config\kinde\config.json` for Windows. You can set API keys, business, and other preferences.

Example:

```json
{
  "current": "mybusiness.kinde.com",
  "environments": {
    "mybusiness.kinde.com": {
      "domain_name": "app.kinde.com"
    },
    "mybusiness2.localkinde.com": {
      "domain_name": "app.localkinde.com"
    },
    "mybusiness2.kinde.com": {
      "domain_name": "mybusiness2.kinde.com",
      "client_id": "xxx",
      "client_secret": "yyy"
    }
  }
}
```

## Examples

- List users:
    ```bash
    kinde manage users get_all --page_size 2
    ```
- Show token details of the logged-in user:
    ```bash
    kinde whoami
    ```
- Create a user
    ```bash
    kinde manage users create
    ```
- Create an environment-level feature-flag of type `string`
  ```
    kinde manage feature_flags create --name "first flag" --key "first_key" --allow_override_level "env" --type "str"
   ```

## Troubleshooting

- CLI validates the token against public JWKS and requires internet connectivity.
- Check your configuration file for errors.
- Ensure your API credentials are valid. To use management API, please create an M2M application and grant it access to the management API, enable appropriate scopes to grant the specific level of access.

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release notes and version history.