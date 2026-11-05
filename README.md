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
  - [x] Support for account API

## Getting Started

### Installation

#### macOS with Homebrew
```bash
brew tap kinde-oss/kinde-cli
brew install kinde-cli
```

**Update:**
```bash
brew upgrade kinde-cli
```

#### Windows with Scoop
```bash
scoop bucket add kinde https://github.com/kinde-oss/scoop-kinde-cli
scoop install kinde-cli
```

**Update:**
```bash
scoop update kinde-cli
```

#### Linux

##### APT Repository (Recommended for Debian/Ubuntu)
```bash
curl -s https://pkg.kinde.com/apt/gpg.key | gpg --dearmor | sudo tee /usr/share/keyrings/kinde.gpg
sudo sh -c 'echo "deb [signed-by=/usr/share/keyrings/kinde.gpg] https://pkg.kinde.com/apt/ * *" > /etc/apt/sources.list.d/kinde.list'
sudo apt update
sudo apt install kinde-cli
```

**Update:**
```bash
apt update
apt install --only-upgrade kinde-cli
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

## Quick start

1. In Kinde, create a machine-to-machine (M2M) application, authorize it for the Kinde Management API, and grant it the scopes you need.
2. Log in with its client ID. You'll be asked for the client secret:

    ```bash
    kinde login --domain mybusiness.kinde.com --client-id <client ID>
    ```

3. Run a command:

    ```bash
    kinde users list --page-size 5
    ```

Run `kinde --help` to see every command, and `kinde <command> --help` for its flags.

Commands print the API's response as JSON on stdout, so you can pipe it to tools like `jq`. Messages and errors go to stderr, and a failed command exits with status 1.

## Authentication

The Kinde Management API only accepts tokens issued to an M2M application, so every resource command (`users`, `organizations`, `roles`, ...) needs an M2M login. Its secret is stored in your operating system's keychain, never in the config file.

```bash
kinde login --domain <your Kinde domain> --client-id <client ID>
Client secret:
```

The secret isn't shown as you type. There's no flag for it, because flag values are kept in your shell history. In a script, pipe it in instead:

```bash
echo "$CLIENT_SECRET" | kinde login --client-id <client ID> --client-secret-stdin
```

You can also log in as a user through the browser (device authorization flow), and use the [Account API](#account-api) to see what that user can access. A user login can't call the Management API:

```bash
kinde account login --domain <your Kinde domain> --client-id <application client ID>
```

`kinde whoami` shows both logins for the domain as JSON, and `kinde logout` removes both.

### Environment variables

In CI, or anywhere you'd rather not store credentials, set these instead of running `kinde login`:

| Variable | Purpose |
|----------|---------|
| `KINDE_DOMAIN` | The Kinde domain to use |
| `KINDE_MANAGEMENT_CLIENT_ID` | M2M application client ID |
| `KINDE_MANAGEMENT_CLIENT_SECRET` | M2M application client secret |

### Multiple businesses

`kinde login` remembers the domain you logged in to as the default. To run a single command against another domain, use `--domain` (or `KINDE_DOMAIN`); it doesn't change the default.

```bash
kinde users list --domain otherbusiness.kinde.com
```

## Commands

Management API commands are grouped by resource, and each one maps to a single API operation. For example, `kinde users create-identity` calls `POST /api/v1/users/{user_id}/identities`. Each command's help shows the scope your M2M application needs.

See the **[command reference](docs/commands/kinde.md)** for every command and flag.

Commands that can't be undone (every delete, plus `api-keys rotate` and `connected-apps revoke-token`) ask you to confirm first. In scripts, pass `--yes`; without it they fail rather than wait for an answer.

### Passing values

- **Simple values** are flags: `--name "Sally"`, `--page-size 10`, `--has-organization`.
- **Lists** are comma-separated: `--scope-ids s1,s2`.
- **Fields that take one of a set of values** list them in `--help` and reject anything else.
- **Objects and lists of objects** are JSON: `--users '[{"id":"kp_123"}]'`.
- **`--body`** sends the whole request body as JSON, inline, from a file (`--body @user.json`), or from stdin (`--body -`). Flags override fields in the body.

### Examples

```bash
# List users with their identities
kinde users list --page-size 2 --expand identities

# Create a user from a file
kinde users create --body @user.json

# Add an email identity to a user
kinde users create-identity --user-id kp_123 --type email --value sally@example.com

# Create a string feature flag
kinde feature-flags create --name "First flag" --key first_flag --type str --default-value hello
```

## Account API

The `kinde account` commands call the Account API as the user logged in with `kinde account login`. They show what your application would see for that user:

```bash
kinde account permissions
kinde account roles
kinde account feature-flags
kinde account entitlements
kinde account entitlement --key <feature key>
kinde account properties
kinde account profile
kinde account portal-link --subnav profile
```

`kinde account token` prints the user's access token, for calling the Account API or your own APIs with other tools:

```bash
curl -H "Authorization: Bearer $(kinde account token)" https://mybusiness.kinde.com/account_api/v1/permissions
```

## Updates

`kinde version` shows the installed version and checks GitHub for a newer release. When you run commands in a terminal, the CLI also checks at most once a day and prints a notice on stderr. It never checks in scripts or CI. Set `KINDE_NO_UPDATE_NOTIFIER=1` to turn the check off.

## Shell completion

Kinde CLI can generate completion scripts for `bash`, `zsh`, `fish` and `powershell`. Run `kinde completion <shell> --help` for setup instructions.

## Configuration

The config file is `~/.config/kinde/config.json` (or `$XDG_CONFIG_HOME/kinde/config.json`). It records the default domain and which logins exist for each domain; secrets and tokens are kept in the keychain.

```json
{
  "current": "mybusiness.kinde.com",
  "environments": {
    "mybusiness.kinde.com": {
      "domain_name": "mybusiness.kinde.com",
      "m2m": { "client_id": "a1b2c3" }
    }
  }
}
```

## Troubleshooting

- **"no M2M login"**: run `kinde login` with an M2M application's credentials, or set the environment variables above.
- **403 Forbidden**: the error names the scope the command needs (also shown in its `--help`). Grant that scope to your M2M application in Kinde.
- **Keychain**: on Linux, credentials are kept in an encrypted file instead of a system keychain, so the CLI also works on servers and in CI. Set `KINDE_KEYCHAIN_PASS` to supply its password when there's no terminal to prompt in.
- **Logging**: `KINDE_LOG_LEVEL` sets the log level (e.g. `debug`), and `KINDE_STRUCTURED_LOG=true` writes logs as JSON.

## Development

The command help and the [command reference](docs/commands/kinde.md) are generated from the Kinde [Management API](api/kinde-management-api-spec.yaml) and [Account API](api/kinde-account-api-spec.yaml) specs pinned in `api/`. After updating the spec, or changing commands, regenerate them:

```bash
go generate ./...
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release notes and version history.