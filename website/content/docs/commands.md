---
title: "CLI Command Reference"
date: 2026-08-11
prev_title: "Manifest Specification"
prev_link: "/docs/manifest/"
next_title: "Engine Architecture"
next_link: "/docs/architecture/"
---

## Command Flags & Subcommands

### lown install `<uri>`

Installs a package from a Git URL, GitHub shorthand (`gh:user/repo`), short alias (`revoq`, `muth`, `runa`), or local path.

```bash
lown install revoq
lown install gh:iamvxrn/muth
lown install ./my-local-tool
```

### lown rollback `<name>`

Restores a package binary to its previous backup executable (`.bak`) in `~/.lown/bin/`.

```bash
lown rollback revoq
```

### lown sync `[name]`

Pulls latest upstream commits and re-compiles binaries if changes are detected. Continues safely if individual repositories fail.

```bash
lown sync
lown sync revoq
```

### lown remove `<name>`

Removes a managed binary from `~/.lown/bin/` and deletes source directory.

```bash
lown remove muth
```

### lown list

Lists all installed tools, versions, languages, and source paths.

### lown doctor

Runs environment diagnostics (PATH validation, Go/Rust compiler detection).
