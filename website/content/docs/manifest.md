---
title: "lown.toml Manifest Schema"
---

The `lown.toml` manifest file defines target build parameters and installation scripts for a package.

## Field Reference

### `[package]`

- `name` (string, required): Package name and directory identifier under `~/.lown/apps/`.
- `version` (string, required): Package semantic version string.
- `language` (string, optional): Native language identifier (`"go"` or `"rust"`).
- `executable` (string, optional): Binary output filename linked to `~/.lown/bin/`. Defaults to `name` if omitted.

### `[scripts]`

- `install` (string, optional): Relative path to installation shell script.
- `uninstall` (string, optional): Relative path to uninstallation shell script.

## Native Build Example

```toml
[package]
name = "my-tool"
version = "0.1.0"
language = "go"
executable = "my-tool"
```

## Script Fallback Example

```toml
[package]
name = "script-tool"
version = "0.1.0"
executable = "script-tool"

[scripts]
install = "install.sh"
uninstall = "uninstall.sh"
```

## Process Environment

During script execution, Lown sets the following environment variables:

| Variable | Description |
| :--- | :--- |
| `LOWN_BIN` | Path to `~/.lown/bin` |
| `LOWN_APP_ROOT` | Path to `~/.lown/apps/<package-name>` |
| `LOWN_PACKAGE_NAME` | Value of `package.name` |
| `LOWN_PACKAGE_VERSION` | Value of `package.version` |
| `LOWN_EXECUTABLE` | Output binary filename |
