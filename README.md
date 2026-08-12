# Lown

> **Lown** is a lightweight, user-space binary manager and task orchestrator designed to compile, link, and sync developer tools without `sudo` privileges or centralized repositories.

> **Note:** Lown was created as an engineering experiment in automated toolchain design — a study in what AI agents can build under human direction. The engine, manifest specification, and documentation site were generated under maintainer guidance.

[![Go Version](https://img.shields.io/badge/go-1.26-4a7bc0.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-4a7bc0.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-4a7bc0.svg)](#)

## Quick Start

Install Lown with a single command:
```bash
curl -fsSL https://lown.pages.dev/install.sh | sh
```

Or install any package via Lown:
```bash
lown install gh:iamvxrn/revoq
```

## Features

- **User-Space Isolation**: Installs binaries to `~/.lown/bin/` and sources to `~/.lown/apps/`. Zero `sudo` required.
- **Smart Native Compilation**: Auto-detects Go, Rust, and Revoq (C/C++) manifests, running `go build`, `cargo build --release`, or `revoq build` automatically.
- **Git Synchronization**: `lown sync` pulls updates across all managed applications and triggers automatic recompilation upon commit SHA or version changes.

## License

Released under the [MIT License](LICENSE).
