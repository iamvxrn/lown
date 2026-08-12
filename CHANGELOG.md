# Changelog

All notable changes to **Lown** will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.2.0] - 2026-08-12

### Added (Feature Release)
- **`lown list` / `lown ls`**: Formatted tabular view of all user-space installed packages.
- **`lown rollback <pkg>`**: One-command binary rollback to `.bak` executable backups.
- **Agent-Native Diagnostics**: Machine-readable `--json` output format for `lown list` and `lown doctor`.
- **Windows Support**: Native `.exe` extension handling for binary links and backups on Windows.
- **Fault-Tolerant `sync`**: Incremental repository sync across all installed tools.

## [v0.1.0] - 2026-08-11

### Added
- Initial v0.1.0 release with user-space installation, alias resolution, and doctor checks.
