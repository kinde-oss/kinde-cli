# Kinde CLI Installation - Supported Architectures and Formats

## Overview

The Kinde CLI is available for multiple operating systems and architectures, providing native support for Linux, macOS, and Windows platforms. This document provides comprehensive information about all supported architectures, package formats, and installation methods.

## Supported Operating Systems

### Linux

#### Debian/Ubuntu Systems
- **x86 (32-bit)**: `.deb`, `.tar.gz`
- **x86_64 (64-bit)**: `.deb`, `.tar.gz`  
- **ARM64 (aarch64)**: `.deb`, `.tar.gz`

#### Red Hat/Fedora/CentOS Systems
- **x86 (32-bit)**: `.rpm`, `.tar.gz`
- **x86_64 (64-bit)**: `.rpm`, `.tar.gz`
- **ARM64 (aarch64)**: `.rpm`, `.tar.gz`

#### Package Manager Installation
- **Debian/Ubuntu**: Use `.deb` files with `dpkg -i`
- **Red Hat/Fedora/CentOS**: Use `.rpm` files with `rpm -i`
- **Generic Linux**: Use `.tar.gz` files and extract to your preferred location

### macOS

- **Universal Binary**: Single binary supports both Intel (x86_64) and Apple Silicon (ARM64)
- **Available via**: Homebrew or direct download
- **Formats**: `.tar.gz` files

### Windows

- **x86_64 (64-bit)**: Available via Scoop or direct download
- **Formats**: `.exe` files and `.tar.gz` archives

## Package Formats Explained

### .deb Packages
- **Use case**: Debian-based Linux distributions (Ubuntu, Debian, etc.)
- **Installation**: `sudo dpkg -i package.deb`
- **Benefits**: Automatic dependency management, system integration

### .rpm Packages  
- **Use case**: Red Hat-based Linux distributions (Fedora, CentOS, RHEL, etc.)
- **Installation**: `sudo rpm -i package.rpm`
- **Benefits**: Automatic dependency management, system integration

### .tar.gz Archives
- **Use case**: All platforms, portable installation
- **Installation**: Extract and place binary in PATH
- **Benefits**: No root access required, portable across systems

### .exe Files
- **Use case**: Windows systems
- **Installation**: Direct execution or via package managers
- **Benefits**: Native Windows integration

## Architecture Support Details

### x86 (32-bit)
- **Linux**: Full support with .deb, .rpm, and .tar.gz formats
- **Use cases**: Legacy systems, embedded devices, resource-constrained environments

### x86_64 (64-bit)
- **Linux**: Full support with .deb, .rpm, and .tar.gz formats
- **macOS**: Supported via universal binary
- **Windows**: Full support with .exe and .tar.gz formats
- **Use cases**: Modern desktop and server systems

### ARM64 (aarch64)
- **Linux**: Full support with .deb, .rpm, and .tar.gz formats
- **macOS**: Supported via universal binary (Apple Silicon)
- **Use cases**: ARM-based servers, Apple Silicon Macs, ARM development boards

## Download Sources

### Official Releases
- **GitHub Releases**: [https://github.com/kinde-oss/kinde-cli/releases](https://github.com/kinde-oss/kinde-cli/releases)
- **Latest Version**: Always available for all supported architectures
- **Checksums**: Provided for verification of download integrity

### Package Managers
- **Homebrew (macOS)**: `brew install kinde-cli`
- **Scoop (Windows)**: `scoop install kinde-cli`
- **APT (Linux)**: Official APT repository available

## Verification and Security

### Checksum Verification
Each release includes checksums for all architecture variants to ensure download integrity and authenticity.

### Native Binaries
- Single native binary per architecture
- No external dependencies required
- Optimized for each target platform

## Installation Recommendations

### For Development
- Use package managers (Homebrew, Scoop, APT) for easy updates
- Prefer official repositories for automatic dependency management

### For Production
- Download specific versions from GitHub releases
- Verify checksums before installation
- Use .tar.gz for portable deployments

### For CI/CD
- Use .tar.gz archives for cross-platform compatibility
- Pin to specific versions for reproducible builds
- Verify checksums in automated pipelines

## Troubleshooting

### Architecture Mismatch
- Verify your system architecture: `uname -m` (Linux/macOS) or `systeminfo` (Windows)
- Download the correct package for your architecture

### Package Installation Issues
- Ensure you have appropriate permissions (sudo for system-wide installation)
- Check for conflicting versions
- Verify package integrity with checksums

### Binary Execution Issues
- Ensure the binary has execute permissions: `chmod +x kinde`
- Add the binary location to your PATH environment variable
- Check for missing system dependencies (rare, as binaries are mostly static)
