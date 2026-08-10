---
title: "CLI Command Reference"
---

Lown provides subcommands for package installation, updates, removal, and diagnostics.

## `lown install <uri>` (Alias: `lown i`)

Installs a package from a Git URL, GitHub shorthand (`gh:user/repo`), custom alias, or local directory path.

```bash
lown install gh:user/repo
lown install https://github.com/user/repo.git
lown install ./my-local-package
```

## `lown sync [name]` (Alias: `lown update`)

Pulls latest changes from Git for a single package or **all** installed packages. If the package version or commit hash has changed, Lown automatically re-builds or re-executes the installer.

```bash
lown sync          # Sync all installed packages
lown sync my-tool  # Sync specific package
```

## `lown remove <name>` (Alias: `lown rm`)

Executes `scripts.uninstall` (if present), removes linked binaries from `~/.lown/bin/`, and purges application sources from `~/.lown/apps/`.

```bash
lown remove my-tool
```

## `lown list` (Alias: `lown ls`)

Displays a clean table of installed packages, versions, build types, and installation dates.

## `lown doctor`

Checks shell environment paths, verifying `~/.lown/bin` is in `$PATH` and reporting directory permissions.
