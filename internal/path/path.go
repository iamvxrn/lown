package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LownRoot returns the root directory for lown state (~/.lown by default).
// Can be overridden via LOWN_ROOT environment variable.
func LownRoot() string {
	if env := os.Getenv("LOWN_ROOT"); env != "" {
		return ExpandPath(env)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ".lown"
	}
	return filepath.Join(home, ".lown")
}

// BinDir returns ~/.lown/bin
func BinDir() string {
	return filepath.Join(LownRoot(), "bin")
}

// AppsDir returns ~/.lown/apps
func AppsDir() string {
	return filepath.Join(LownRoot(), "apps")
}

// AppPath returns ~/.lown/apps/<name>
func AppPath(name string) string {
	return filepath.Join(AppsDir(), name)
}

// CacheDir returns ~/.lown/cache
func CacheDir() string {
	return filepath.Join(LownRoot(), "cache")
}

// InventoryPath returns ~/.lown/inventory.json
func InventoryPath() string {
	return filepath.Join(LownRoot(), "inventory.json")
}

// ConfigDir returns ~/.config/lown (respects XDG_CONFIG_HOME if set)
func ConfigDir() string {
	if env := os.Getenv("LOWN_CONFIG_DIR"); env != "" {
		return ExpandPath(env)
	}
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(ExpandPath(xdg), "lown")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(LownRoot(), "config")
	}
	return filepath.Join(home, ".config", "lown")
}

// ConfigPath returns ~/.config/lown/config.toml
func ConfigPath() string {
	return filepath.Join(ConfigDir(), "config.toml")
}

// EnsureDirs creates necessary directory structures (~/.lown/bin, apps, cache, config).
func EnsureDirs() error {
	dirs := []string{
		BinDir(),
		AppsDir(),
		CacheDir(),
		ConfigDir(),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

// ExpandPath expands leading ~ to user's home directory.
func ExpandPath(p string) string {
	if !strings.HasPrefix(p, "~") {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return p
	}
	if p == "~" {
		return home
	}
	if strings.HasPrefix(p, "~/") {
		return filepath.Join(home, p[2:])
	}
	return p
}
