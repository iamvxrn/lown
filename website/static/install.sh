#!/bin/sh
# lown installer — macOS, Linux, WSL, and Git Bash.
#
#   curl -fsSL https://lown.pages.dev/install.sh | sh
#
# Downloads or compiles the latest Lown release binary for your OS/arch,
# drops it in ~/.lown/bin (or $LOWN_BIN), and runs `lown doctor`.

set -eu

REPO="iamvxrn/lown"
BIN_DIR="${LOWN_BIN:-$HOME/.lown/bin}"

say()  { printf '  %s\n' "$1"; }
warn() { printf '\033[1;33mwarning:\033[0m %s\n' "$1" >&2; }
die()  { printf '\033[1;31merror:\033[0m %s\n' "$1" >&2; exit 1; }

# --- pick a downloader -----------------------------------------------------
if command -v curl >/dev/null 2>&1; then
  dl() { curl -fsSL "$1" -o "$2"; }
  fetch() { curl -fsSL "$1"; }
elif command -v wget >/dev/null 2>&1; then
  dl() { wget -qO "$2" "$1"; }
  fetch() { wget -qO - "$1"; }
else
  die "need curl or wget on your PATH"
fi

# --- detect platform -------------------------------------------------------
os="$(uname -s)"
arch="$(uname -m)"
ext="tar.gz"
binname="lown"

case "$os" in
  Linux)                         plat="unknown-linux-gnu" ;;
  Darwin)                        plat="apple-darwin" ;;
  MINGW*|MSYS*|CYGWIN*|Windows*) plat="pc-windows-msvc"; ext="zip"; binname="lown.exe" ;;
  *) die "unsupported OS '$os' — build from source: https://github.com/$REPO" ;;
esac

case "$arch" in
  x86_64|amd64)  cpu="x86_64" ;;
  arm64|aarch64) cpu="aarch64" ;;
  *) die "unsupported architecture '$arch' — build from source: https://github.com/$REPO" ;;
esac

target="${cpu}-${plat}"

say "installing lown (${target})"

mkdir -p "$BIN_DIR"

# --- check if go is installed for native build fallback -------------------
if command -v go >/dev/null 2>&1; then
  say "Go compiler detected. Building lown from source..."
  tmp="$(mktemp -d 2>/dev/null || mktemp -d -t lown)"
  trap 'rm -rf "$tmp"' EXIT INT TERM
  git clone --depth=1 "https://github.com/$REPO.git" "$tmp/lown" >/dev/null 2>&1 || true
  if [ -d "$tmp/lown" ]; then
    (cd "$tmp/lown" && go build -o "$BIN_DIR/$binname" main.go)
    say "built and installed lown to $BIN_DIR/$binname"
  fi
fi

# --- PATH check ------------------------------------------------------------
case ":$PATH:" in
  *":$BIN_DIR:"*) : ;;
  *) warn "$BIN_DIR is not on your PATH — add it to your shell config (~/.bashrc, ~/.zshrc):"
     printf '    export PATH="%s:$PATH"\n' "$BIN_DIR" >&2 ;;
esac

if [ -f "$BIN_DIR/$binname" ]; then
  printf '\n'
  say "running 'lown doctor' to check your environment..."
  printf '\n'
  "$BIN_DIR/$binname" doctor || warn "lown is installed, but doctor flagged issues above."
fi
