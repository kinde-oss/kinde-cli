# Kinde command line utility

##### This is work in progress. API is subject to change without warning or be missing.

## Features

- Official CLI for interacting with Kinde APIs.
- Secure handling of credentials via OS-specific secrets store.
- 100% verifyiable authenticity
  - 100% built in go without external native dependencies
  - Compiles into a single executable for ease of maintenance and deployment
  - Ships for multiple platforms - Linux, MacOS, Windows
  - Ships for multiple architectures - x64, ARM, Apple ARM
  - Each variation comes with the checksum to verify authenticity
  - Where applicable the binary is signed/notarized with Kinde private key
- Secure by default. Kinde business cannot be managed via the CLI until a new M2M application is created with the management API authorized and the select scopes granted.

TODO:
- [x] Client-credentials authentication
- [x] Managaement API top-level sub-commands and flags
  - [ ] Support for array arguments
- [x] Online token verification
  - [ ] Offline token verification
- [ ] Support for user tokens and account API

## Getting Started

### Installation

Download a pre-built binary from [Releases](https://github.com/kinde-oss/kinde-cli/releases).

Or install from source using `go`:

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

Please note, although this authentication method is supported, account API is in TODO.

Please configure a default application for defice flow or specify `client_id`

```bash
kinde login
    --domain <your Kinde business domain>
    --client_id [optional]
```

## Check authenticaton status

```bash
kinde whoami
```

## Multiple businesses and environments

Kinde CLI supports authentication to multiple businesses via the `--domain` argument.

Enviornments are independent from each-other and cannot be cross-authenticated, each token is linked to an environment, and cannot be used to cross-authenticate into multiple environments.

When `--domain` is not specified, the last logged-in business becomes `current`.

## Terminal autocomplete

Kinde CLI comes with terminal autompletion support for `bash`, `fish`, `powershell` and `zsh`

You can see instruction for your shell/OS combination by executing the following

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
- `apis`
- `applications`
- `billing`
- `business`
- `categories`
- `connected_apps`
- `connections`
- `environment_variables`
- `environments`
- `events`
- `feature_flags`
- `industries`
- `organizations`
- `permissions`
- `portal`
- `properties`
- `roles`
- `subscribers`
- `timezones`
- [`users`](docs/subcommands/users.md)
- `webhooks`

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
- Create an enviornment-level feature-flag of type `string`
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