# Changelog

All notable changes to Lown will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - Planned / Roadmap

### Planned Features
- **Enhanced Windows Support**: Native PowerShell (`.ps1`) and CMD batch installer hooks alongside Linux/macOS `/bin/sh`.
- **Custom Build Directives**: Flexible `[build]` block in `lown.toml` allowing custom command triggers (e.g. `make`, `ninja`, `zig build`).
- **Environment Variable Overrides**: Support for custom user-defined env maps in `lown.toml`.
- **Global Package Locking**: Optional `lown.lock` for pinning specific Git commit SHAs across installations.

---

## [0.1.0] - 2026-08-10

### Added
- **User-space Package Manager Core**: Initial release operating strictly in `~/.lown/` (`bin/`, `apps/`, `cache/`).
- **Smart Installation Engine**:
  - Native compilation support for `go` (`go build -o`) and `rust` (`cargo build --release`).
  - Fallback script execution (`install.sh` / `uninstall.sh`) with `$LOWN_BIN` and `$LOWN_APP_ROOT` exported.
- **URI Resolution**: Direct Git clone (`https://`, `git@`), GitHub shorthand (`gh:org/repo`), local directory paths, and custom aliases in `~/.config/lown/config.toml`.
- **Git Synchronization**: `lown sync [name]` to pull updates and trigger automatic recompilation upon commit SHA or version changes.
- **Diagnostics**: `lown doctor` for environment and `$PATH` verification.
- **Documentation Sites**: Hugo landing and documentation site for Lown.
