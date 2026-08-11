---
title: "Engine Architecture"
date: 2026-08-11
prev_title: "CLI Command Reference"
prev_link: "/docs/commands/"
next_title: ""
next_link: ""
---

## Internal Directory Structure

Lown isolates all user-space state under `~/.lown/`:

```
~/.lown/
├── bin/          # Extracted & compiled executables (on PATH)
├── apps/         # Cloned Git repositories & source code
└── lown.db       # SQLite / JSON metadata registry
```

## Build Engine Dispatch

When `lown install` is called:
1. `ResolveURI` parses shorthand input or alias lookup.
2. `Git.CloneOrFetch` clones source into `~/.lown/apps/<pkg>`.
3. `Manifest.LoadFromDir` validates `lown.toml`.
4. `Builder` invokes language compiler (`go`, `cargo`, `revoq`) or script fallback.
5. Executable is symlinked/copied to `~/.lown/bin/`.
