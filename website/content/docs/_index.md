---
title: "Overview"
---

Lown is a CLI tool for compiling and installing user-space binaries directly from Git repositories or local source paths.

## Execution Flow

1. **Fetch**: Clones the source repository into `~/.lown/apps/<package-name>/`.
2. **Inspect**: Parses `lown.toml` at the repository root.
3. **Build**:
   - If `language = "go"`: Runs `go build -o ~/.lown/bin/<executable> .`
   - If `language = "rust"`: Runs `cargo build --release` and copies target binary to `~/.lown/bin/<executable>`
   - Fallback: Runs `scripts.install` with `LOWN_BIN` and `LOWN_APP_ROOT` environment variables set.
4. **Register**: Writes package metadata to `~/.lown/inventory.json`.

## Installation

Via Go:

```bash
go install github.com/iamvxrn/lown@latest
```

Via Git source:

```bash
git clone https://github.com/iamvxrn/lown.git
cd lown
go build -o lown main.go
mv lown ~/.lown/bin/
```

Environment check:

```bash
lown doctor
```
