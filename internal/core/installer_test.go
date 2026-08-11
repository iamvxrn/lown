package core

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"lown/internal/path"
)

func TestSmartInstallationGo(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("LOWN_BIN", filepath.Join(tmpDir, ".lown", "bin"))
	t.Setenv("LOWN_APPS", filepath.Join(tmpDir, ".lown", "apps"))

	installer, err := NewInstaller()
	if err != nil {
		t.Fatalf("NewInstaller error: %v", err)
	}

	// Create mock Go package directory
	pkgDir := filepath.Join(tmpDir, "mock-go-pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("failed to create pkg dir: %v", err)
	}

	// Write lown.toml
	manifestContent := `
[package]
name = "mock-go-tool"
version = "0.1.0"
language = "go"
executable = "mock-go-tool"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "lown.toml"), []byte(manifestContent), 0644); err != nil {
		t.Fatalf("failed to write lown.toml: %v", err)
	}

	// Write mock main.go & go.mod
	if err := os.WriteFile(filepath.Join(pkgDir, "go.mod"), []byte("module mock-go-tool\n\ngo 1.20\n"), 0644); err != nil {
		t.Fatalf("failed to write go.mod: %v", err)
	}
	mainGoContent := `package main
import "fmt"
func main() { fmt.Println("mock go tool") }
`
	if err := os.WriteFile(filepath.Join(pkgDir, "main.go"), []byte(mainGoContent), 0644); err != nil {
		t.Fatalf("failed to write main.go: %v", err)
	}

	// Install package
	if err := installer.Install(pkgDir); err != nil {
		t.Fatalf("Install failed: %v", err)
	}

	// Verify binary exists
	binPath := filepath.Join(installer.cfg.BinDir, "mock-go-tool")
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		t.Errorf("binary was not created at expected path: %s", binPath)
	}

	// Remove package
	if err := installer.Remove("mock-go-tool"); err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	// Verify binary was removed
	if _, err := os.Stat(binPath); !os.IsNotExist(err) {
		t.Errorf("binary was not removed from path: %s", binPath)
	}
}

func TestSmartInstallationScriptFallback(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("LOWN_BIN", filepath.Join(tmpDir, ".lown", "bin"))
	t.Setenv("LOWN_APPS", filepath.Join(tmpDir, ".lown", "apps"))

	installer, err := NewInstaller()
	if err != nil {
		t.Fatalf("NewInstaller error: %v", err)
	}

	// Create mock script package directory
	pkgDir := filepath.Join(tmpDir, "mock-script-pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("failed to create pkg dir: %v", err)
	}

	// Write lown.toml
	manifestContent := `
[package]
name = "mock-script-tool"
version = "0.1.0"
executable = "mock-script-tool"

[scripts]
install = "install.sh"
uninstall = "uninstall.sh"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "lown.toml"), []byte(manifestContent), 0644); err != nil {
		t.Fatalf("failed to write lown.toml: %v", err)
	}

	// Write mock install.sh
	installShContent := `#!/bin/sh
echo "Installing script tool to $LOWN_BIN"
touch "$LOWN_BIN/mock-script-tool"
chmod +x "$LOWN_BIN/mock-script-tool"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "install.sh"), []byte(installShContent), 0755); err != nil {
		t.Fatalf("failed to write install.sh: %v", err)
	}

	// Write mock uninstall.sh
	uninstallShContent := `#!/bin/sh
echo "Uninstalling script tool from $LOWN_BIN"
rm -f "$LOWN_BIN/mock-script-tool"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "uninstall.sh"), []byte(uninstallShContent), 0755); err != nil {
		t.Fatalf("failed to write uninstall.sh: %v", err)
	}

	// Install package
	if err := installer.Install(pkgDir); err != nil {
		t.Fatalf("Install failed: %v", err)
	}

	// Verify binary exists
	binPath := filepath.Join(installer.cfg.BinDir, "mock-script-tool")
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		t.Errorf("binary was not created at expected path: %s", binPath)
	}

	// Remove package
	if err := installer.Remove("mock-script-tool"); err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	// Verify binary was removed
	if _, err := os.Stat(binPath); !os.IsNotExist(err) {
		t.Errorf("binary was not removed from path: %s", binPath)
	}
}

func TestSmartInstallationValidationFailure(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("LOWN_BIN", filepath.Join(tmpDir, ".lown", "bin"))
	t.Setenv("LOWN_APPS", filepath.Join(tmpDir, ".lown", "apps"))

	installer, err := NewInstaller()
	if err != nil {
		t.Fatalf("NewInstaller error: %v", err)
	}

	// Create mock invalid package directory (no native source, no install script)
	pkgDir := filepath.Join(tmpDir, "mock-invalid-pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("failed to create pkg dir: %v", err)
	}

	manifestContent := `
[package]
name = "mock-invalid-tool"
version = "0.1.0"
language = "python"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "lown.toml"), []byte(manifestContent), 0644); err != nil {
		t.Fatalf("failed to write lown.toml: %v", err)
	}

	// Attempt install, expect error
	if err := installer.Install(pkgDir); err == nil {
		t.Errorf("expected Install to fail for invalid package manifest")
	}
}

func TestConfigResolveURI(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("LOWN_CONFIG_DIR", filepath.Join(tmpDir, ".config", "lown"))

	cfgPath := path.ConfigPath()
	_ = os.MkdirAll(filepath.Dir(cfgPath), 0755)

	tomlContent := `
[aliases]
mytool = "gh:foo/bar"
`
	_ = os.WriteFile(cfgPath, []byte(tomlContent), 0644)

	installer, err := NewInstaller()
	if err != nil {
		t.Fatalf("NewInstaller error: %v", err)
	}

	resolved := installer.cfg.ResolveURI("mytool")
	expected := "https://github.com/foo/bar.git"
	if resolved != expected {
		t.Errorf("alias resolution failed: expected %s, got %s", expected, resolved)
	}

	resolvedRevoq := installer.cfg.ResolveURI("revoq")
	expectedRevoq := "https://github.com/iamvxrn/revoq.git"
	if resolvedRevoq != expectedRevoq {
		t.Errorf("revoq built-in alias failed: expected %s, got %s", expectedRevoq, resolvedRevoq)
	}

	resolvedMuth := installer.cfg.ResolveURI("muth")
	expectedMuth := "https://github.com/iamvxrn/muth.git"
	if resolvedMuth != expectedMuth {
		t.Errorf("muth built-in alias failed: expected %s, got %s", expectedMuth, resolvedMuth)
	}

	resolvedRuna := installer.cfg.ResolveURI("runa")
	expectedRuna := "https://github.com/iamvxrn/runa.git"
	if resolvedRuna != expectedRuna {
		t.Errorf("runa built-in alias failed: expected %s, got %s", expectedRuna, resolvedRuna)
	}

	resolvedGH := installer.cfg.ResolveURI("gh:user/repo")
	expectedGH := "https://github.com/user/repo.git"
	if resolvedGH != expectedGH {
		t.Errorf("gh: resolution failed: expected %s, got %s", expectedGH, resolvedGH)
	}

	fmt.Println("Config resolution tests passed successfully.")
}
