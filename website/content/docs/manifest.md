---
title: "Manifest Specification"
date: 2026-08-11
prev_title: "5-Minute Quickstart"
prev_link: "/docs/quickstart/"
next_title: "CLI Command Reference"
next_link: "/docs/commands/"
---

## lown.toml Schema

The `lown.toml` manifest file defines how Lown compiles, links, and manages a package.

### Package Metadata

```toml
[package]
name = "my-tool"
version = "0.1.0"
language = "go"
executable = "my-tool"
```

### Supported Languages

- `go`: Invokes `go build`
- `rust`: Invokes `cargo build --release`
- `revoq`: Invokes `revoq build`
- `c` / `cpp`: Native Revoq C/C++ compilation

### Script Fallbacks

```toml
[scripts]
install = "install.sh"
uninstall = "uninstall.sh"
```
