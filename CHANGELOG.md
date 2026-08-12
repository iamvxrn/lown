# Changelog

All notable changes to **Lown** will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.1.2] - 2026-08-12

### Added
- **Windows Executable Support**: Automatic `.exe` extension handling for binary links and backups on Windows.
- **Agent-Native Diagnostics**: Machine-readable `--json` output format for `lown list` and `lown doctor`.

## [v0.1.1] - 2026-08-11

### Added
- **`lown list` / `lown ls`**: Formatted tabular view of all user-space installed packages and linked binaries.
- **`lown rollback <pkg>`**: One-command binary rollback to `.bak` executable backups.
- **Fault-Tolerant `sync`**: Incremental repository sync across all installed tools without halting on single remote failures.

## [v0.1.0] - 2026-08-11

### Added
- Initial v0.1.0 release with user-space installation, alias resolution, and doctor checks.
