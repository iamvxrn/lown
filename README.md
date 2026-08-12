# Lown

> **Lown** is a decentralized, user-space package manager and tool orchestrator for Linux and macOS. It installs developer tools directly into `~/.local/bin` with zero root privileges, automatic rollback backups (`.bak`), multi-forge resolution (GitHub, GitLab, Codeberg, Sourcehut, raw HTTPS), and agent-native JSON diagnostics.

[![Latest Release](https://img.shields.io/github/v/release/iamvxrn/lown?color=4a7bc0&label=release)](https://github.com/iamvxrn/lown/releases)
[![Documentation](https://img.shields.io/badge/docs-lown.pages.dev-4a7bc0.svg)](https://lown.pages.dev)
[![Deployment](https://img.shields.io/badge/deployment-active-4a7bc0.svg)](https://lown.pages.dev)
[![OS Matrix](https://img.shields.io/badge/OS-Linux%20%7C%20macOS%20%7C%20Windows-4a7bc0.svg)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-4a7bc0.svg)](LICENSE)
[![AI Pairing](https://img.shields.io/badge/AI%20Pairing-Antigravity%20Agent-4a7bc0.svg)](#)

```
       Git Remotes (GitHub, GitLab, Codeberg, Sourcehut)
                               │
                       [ Lown Engine ]
                               │
                ~/.local/bin/ & ~/.lown/pkg/
```

## 🚀 Key Capabilities

- **User-Space Isolation**: No `sudo` or root permissions required. Installs directly to `~/.local/bin`.
- **Multi-Forge Resolvers**: Supports GitHub (`gh:`), GitLab (`gl:`), Codeberg (`cb:`), Sourcehut (`sh:`), and direct HTTPS tarballs/zips.
- **Legacy Build Auto-Detection**: Native build engine detection for `make` (Makefile), `cmake` (CMakeLists.txt), and `./configure && make` (Autotools) alongside Go, Cargo, and Revoq.
- **One-Command Rollback**: `lown rollback <pkg>` instantly restores previous executable backups (`.bak`).
- **Agent-Native `--json` Flags**: Machine-readable status and diagnostics for LLMs and automated scripts.
- **Cross-OS Compatibility**: Full support for Linux (`x86_64`, `aarch64`, `armv7`, `i686` 32-bit), macOS (Intel & Apple Silicon), and Windows.

## 📦 Quick Start

### Installation

Install via shell script:
```bash
curl -fsSL https://lown.pages.dev/install.sh | sh
```

Or download precompiled binary release archives from [GitHub Releases](https://github.com/iamvxrn/lown/releases).

### Usage

```bash
# Install tool from GitHub or alias
lown install revoq
lown install gh:user/repository

# List installed user-space tools
lown list

# Rollback to previous version
lown rollback revoq

# Run system health check
lown doctor
```

## 📖 Documentation & Releases

- **Documentation & Guides**: [https://lown.pages.dev](https://lown.pages.dev)
- **Changelog & Version History**: [https://lown.pages.dev/changelog/](https://lown.pages.dev/changelog/)
- **GitHub Release Assets**: [https://github.com/iamvxrn/lown/releases](https://github.com/iamvxrn/lown/releases)

## 📄 License

Released under the [MIT License](LICENSE). Built autonomously by AI agentic pairing under human direction.
