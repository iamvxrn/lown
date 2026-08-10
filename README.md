# Lown

> **Lown** is a lightweight, user-space binary manager and task orchestrator designed to compile, link, and sync developer tools without `sudo` privileges or centralized repositories.

> **Note:** Lown was created as an engineering experiment in automated toolchain design — a study in what AI agents can build under human direction. The engine, manifest specification, and documentation site were generated under maintainer guidance.

[![Go Version](https://img.shields.io/badge/go-1.26-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

---

## 🌟 Philosophy

1. **No `sudo` (User-space)**: Everything lives inside `~/.lown/` (`bin/`, `apps/`, `cache/`) and `~/.config/lown/config.toml`. Your system root stays untouched.
2. **No Central Registry (Decentralized)**: Install software directly from Git repositories (`gh:user/repo`, `https://...`, `git@...`, or local directories).
3. **Transparency**: Explicit manifests (`lown.toml`) define native language build targets or clear shell scripts.

---

## 🚀 Quick Start

### Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/faf/lown.git
cd lown
go build -o lown main.go
mv lown ~/.lown/bin/ # or move to any directory in your PATH
```

Ensure `~/.lown/bin` is in your shell `PATH`:

```bash
export PATH="$HOME/.lown/bin:$PATH"
```

Run diagnostics to verify your setup:

```bash
lown doctor
```

---

## 📄 Manifest Specification (`lown.toml`)

Projects declare how Lown should build and install them by including a `lown.toml` file in their repository root.

### 1. Native Build (Go or Rust)

Lown automatically runs `go build -o <executable>` or `cargo build --release` and links the resulting binary to `~/.lown/bin/`.

```toml
[package]
name = "my-tool"
version = "0.1.0"
language = "go" # or "rust"
executable = "my-tool" # Optional: defaults to package name if omitted
```

### 2. Custom Script Fallback

For non-compiled software, multi-binary suites, or custom installation procedures, omit `language` and specify `scripts.install`.

```toml
[package]
name = "custom-app"
version = "2.1.0"
executable = "custom-app"

[scripts]
install = "install.sh"
uninstall = "uninstall.sh"
```

When executing fallback scripts, Lown exports the following environment variables:
- `LOWN_BIN`: Path to `~/.lown/bin`
- `LOWN_APP_ROOT`: Path to `~/.lown/apps/<package-name>`
- `LOWN_PACKAGE_NAME`: Package name
- `LOWN_PACKAGE_VERSION`: Package version
- `LOWN_EXECUTABLE`: Declared binary name

---

## ⚙️ Configuration (`~/.config/lown/config.toml`)

Create custom shorthand aliases for frequently installed packages:

```toml
[aliases]
mytool = "gh:myuser/mytool"
coolcli = "https://github.com/example/coolcli.git"
```

---

## 🛠️ Commands

| Command | Shorthand | Description |
| :--- | :--- | :--- |
| `lown install <uri>` | `lown i <uri>` | Install package from Git URL, `gh:user/repo`, alias, or local directory |
| `lown remove <name>` | `lown rm <name>` | Execute `uninstall.sh` (if present) and remove binary + app directory |
| `lown sync [name]` | `lown update` | Pull latest Git changes and re-build updated packages |
| `lown list` | `lown ls` | Display table of installed packages and metadata |
| `lown doctor` | | Check environment paths and Lown state |
| `lown version` | `lown -v` | Print Lown CLI version |

---

## 🧪 Running Tests

```bash
go test -v ./...
```

---

## 📜 License

Distributed under the [MIT License](LICENSE).
