---
title: "Overview & Configuration"
date: 2026-08-13
prev_title: ""
prev_link: ""
next_title: "Installation Track"
next_link: "/docs/installation/"
---

# Lown Configuration & Architecture

Lown is configured flexibly via `~/.config/lown/config.toml` (XDG standard) or local `lown.toml` files.

## Configuration Schema (`~/.config/lown/config.toml`)

```toml
bin_dir = "~/.local/bin"          # Directory where binary links are created
apps_dir = "~/.lown/apps"         # Directory where package repositories are cloned
default_forge = "github"          # "github", "gitlab", "codeberg", "sourcehut"
auto_backup = true                # Create executable backup (.bak) on update

[aliases]
revoq = "gh:iamvxrn/revoq"
muth  = "gh:iamvxrn/muth"
trigg = "gh:iamvxrn/trigg"
mytool = "gl:user/repo"           # Custom GitLab alias
```

## Core Principles

- **Zero Root Privileges**: All compiled binaries reside strictly inside `~/.local/bin/` or `~/.lown/bin/`.
- **Native Builder Engine**: Automatically detects project types (`Go`, `Rust`, `Revoq`, `Make`, `CMake`, `Autotools`) and builds native executables.
- **One-Command Rollback**: Restore previous executable backups instantly via `lown rollback <pkg>`.
