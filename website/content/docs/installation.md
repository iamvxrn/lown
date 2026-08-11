---
title: "Installation Track"
date: 2026-08-11
prev_title: "Overview & Philosophy"
prev_link: "/docs/overview/"
next_title: "5-Minute Quickstart"
next_link: "/docs/quickstart/"
---

## Installation Methods

Lown can be installed on macOS, Linux, WSL, and Git Bash using a single automated shell script.

### Shell Installer (Recommended)

Run the following command in your terminal:

```bash
curl -fsSL https://lown.pages.dev/install.sh | sh
```

### Manual Build from Source

If you have Go installed on your PATH:

```bash
git clone https://github.com/iamvxrn/lown.git
cd lown
go build -o ~/.lown/bin/lown main.go
```

## PATH Configuration

Ensure `~/.lown/bin` is exported in your shell configuration (`~/.bashrc` or `~/.zshrc`):

```bash
export PATH="$HOME/.lown/bin:$PATH"
```
