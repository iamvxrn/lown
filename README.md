# Lown

> **Lown** is a lightweight, user-space package manager and toolchain orchestrator designed to compile, link, and sync developer tools without `sudo` privileges or centralized registries.

[![Go Version](https://img.shields.io/badge/go-1.22-4a7bc0.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-4a7bc0.svg)](LICENSE)
[![Nightly CI](https://github.com/iamvxrn/lown/actions/workflows/nightly.yml/badge.svg)](https://github.com/iamvxrn/lown/actions/workflows/nightly.yml)

```
                       ~/.lown/
                           ├── bin/       <- User-Space Binaries ($PATH)
                           ├── apps/      <- Git Sources
                           └── config.toml <- Aliases & Curated Universe
```

## Quick Start

Install Lown with a single command:
```bash
curl -fsSL https://lown.pages.dev/install.sh | sh
```

Install the complete **FAF Developer Suite** (`lown`, `muth`, `revoq`, `trigg`) in one go:
```bash
lown install faf
```

Or install any package via Git shorthand or curated alias:
```bash
lown install gh:iamvxrn/revoq
lown install ripgrep
lown install fzf
```

## Features

- **User-Space Isolation**: Installs binaries to `~/.lown/bin/` and sources to `~/.lown/apps/`. Zero `sudo` required.
- **Unified FAF Meta-Package**: `lown install faf` deploys all 4 ecosystem tools automatically.
- **Smart Native Compilation**: Auto-detects Go, Rust, and Revoq (C/C++) manifests, running `go build`, `cargo build --release`, or `revoq build` automatically.
- **Curated CLI Universe**: Instant alias resolution for `ripgrep`, `fd`, `bat`, `fzf`, `zoxide`, `eza`, `jq`, and `vhs`.
- **Instant Rollback**: Restore previous executable versions instantly with `lown rollback <package>`.
- **Agent-Native `--json` Flags**: Machine-readable outputs for AI coding agents (`lown list --json`, `lown doctor --json`).

## Documentation

Full guides and architecture documentation: https://lown.pages.dev

## License

Released under the [MIT License](LICENSE).
