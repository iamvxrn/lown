---
title: "Overview & Configuration Precedence"
date: 2026-08-13
prev_title: ""
prev_link: ""
next_title: "Installation Track"
next_link: "/docs/installation/"
---

# Lown Architecture & Configuration Precedence

Lown resolves configuration settings across 5 distinct layers to gracefully resolve conflicts between global user settings and local repository manifests.

## Configuration Precedence Hierarchy

Settings are resolved in order of priority (highest to lowest):

1. **CLI Arguments** (e.g. `lown install --bin-dir ~/.bin gh:user/repo`)
2. **Environment Variables** (e.g. `LOWN_BIN_DIR=~/.local/bin`)
3. **Repository Config** (Local `lown.toml` in project directory)
4. **Global User Config** (`~/.config/lown/config.toml`)
5. **Built-in Defaults** (`~/.local/bin`, GitHub registry)

> **Conflict Resolution**: Local project aliases and build configurations override global user settings, ensuring deterministic builds across machines while allowing users to define global fallback aliases.

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
