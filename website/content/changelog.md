---
title: "Changelog"
description: "Release history and version changelog for Lown"
---

# Changelog & Release History

All notable changes to **Lown** are documented on this page.

---

## [v0.3.0] — 2026-08-13

### Added (Feature Release)
- **Multi-Forge Resolvers**: Direct package resolution for GitHub (`gh:`), GitLab (`gl:`), Codeberg (`cb:`), Sourcehut (`sh:`), and raw HTTPS tarballs/zip archives.
- **Legacy Build Auto-Detection**: Native build engine detection for `make` (Makefile), `cmake` (CMakeLists.txt), and `./configure && make` (Autotools) alongside Go, Cargo, and Revoq.
- **Cross-OS Compatibility**: Enhanced 32-bit Linux (`i686`), ARMv7, Apple Silicon (`arm64`), and Windows binary link management.
- **Expanded Feature Matrix**: 6-feature comparison matrix vs Homebrew, Cargo, and Nix/Devbox.
- **AI Development Transparency**: Explicit disclosure notice in footer and docs.

---

## [v0.2.0] — 2026-08-12

### Added
- **`lown list` / `lown ls`**: Formatted tabular view of all user-space installed packages.
- **`lown rollback <pkg>`**: One-command binary rollback to `.bak` executable backups.
- **Agent-Native Diagnostics**: Machine-readable `--json` output format for `lown list` and `lown doctor`.
- **Windows Support**: Native `.exe` extension handling for binary links and backups on Windows.
- **Fault-Tolerant `sync`**: Incremental repository sync across all installed tools.

---

## [v0.1.0] — 2026-08-11

### Added
- Initial v0.1.0 release with user-space installation, alias resolution, and doctor checks.
