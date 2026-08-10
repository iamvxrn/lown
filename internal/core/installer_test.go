package core

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"lown/internal/db"
	"lown/internal/manifest"
	"lown/internal/path"
)

func TestSmartInstallationGo(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("LOWN_ROOT", filepath.Join(tmpDir, ".lown"))
	t.Setenv("LOWN_CONFIG_DIR", filepath.Join(tmpDir, ".config", "lown"))

	if err := path.EnsureDirs(); err != nil {
		t.Fatalf("EnsureDirs failed: %v", err)
	}

	// Create a mock Go package repo
	pkgDir := filepath.Join(tmpDir, "mock-go-pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("failed to create mock pkg dir: %v", err)
	}

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

	goCode := `package main

import "fmt"

func main() {
	fmt.Println("Hello from mock Go tool")
}
`
	if err := os.WriteFile(filepath.Join(pkgDir, "main.go"), []byte(goCode), 0644); err != nil {
		t.Fatalf("failed to write main.go: %v", err)
	}
	if err := os.WriteFile(filepath.Join(pkgDir, "go.mod"), []byte("module mockgotool\n\ngo 1.22\n"), 0644); err != nil {
		t.Fatalf("failed to write go.mod: %v", err)
	}

	// Install package
	installer, err := NewInstaller()
	if err != nil {
		t.Fatalf("NewInstaller error: %v", err)
	}

	if err := installer.Install(pkgDir); err != nil {
		t.Fatalf("Install failed: %v", err)
	}

	// Verify binary exists in ~/.lown/bin/
	binPath := filepath.Join(path.BinDir(), "mock-go-tool")
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		t.Errorf("expected binary at %s was not created", binPath)
	}

	// Verify inventory entry
	store := db.NewStore()
	rec, exists, err := store.Get("mock-go-tool")
	if err != nil || !exists {
		t.Fatalf("inventory check failed: %v (exists: %v)", err, exists)
	}

	if rec.Version != "0.1.0" || rec.InstallType != "native-go" {
		t.Errorf("unexpected record data: %+v", rec)
	}

	// Remove package
	if err := installer.Remove("mock-go-tool"); err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	if _, err := os.Stat(binPath); !os.IsNotExist(err) {
		t.Errorf("binary at %s should have been removed", binPath)
	}

	_, exists, _ = store.Get("mock-go-tool")
	if exists {
		t.Errorf("package record should have been deleted from inventory")
	}
}

func TestSmartInstallationScriptFallback(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("LOWN_ROOT", filepath.Join(tmpDir, ".lown"))
	t.Setenv("LOWN_CONFIG_DIR", filepath.Join(tmpDir, ".config", "lown"))

	if err := path.EnsureDirs(); err != nil {
		t.Fatalf("EnsureDirs failed: %v", err)
	}

	pkgDir := filepath.Join(tmpDir, "mock-script-pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("failed to create mock pkg dir: %v", err)
	}

	manifestContent := `
[package]
name = "mock-script-tool"
version = "0.1.0"

[scripts]
install = "install.sh"
uninstall = "uninstall.sh"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "lown.toml"), []byte(manifestContent), 0644); err != nil {
		t.Fatalf("failed to write lown.toml: %v", err)
	}

	installScript := `#!/bin/sh
echo "Installing script tool to $LOWN_BIN"
mkdir -p "$LOWN_BIN"
echo "#!/bin/sh\necho script-tool-v0.1.0" > "$LOWN_BIN/mock-script-tool"
chmod +x "$LOWN_BIN/mock-script-tool"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "install.sh"), []byte(installScript), 0755); err != nil {
		t.Fatalf("failed to write install.sh: %v", err)
	}

	uninstallScript := `#!/bin/sh
echo "Uninstalling script tool from $LOWN_BIN"
rm -f "$LOWN_BIN/mock-script-tool"
`
	if err := os.WriteFile(filepath.Join(pkgDir, "uninstall.sh"), []byte(uninstallScript), 0755); err != nil {
		t.Fatalf("failed to write uninstall.sh: %v", err)
	}

	installer, err := NewInstaller()
	if err != nil {
		t.Fatalf("NewInstaller error: %v", err)
	}

	if err := installer.Install(pkgDir); err != nil {
		t.Fatalf("Install failed: %v", err)
	}

	binPath := filepath.Join(path.BinDir(), "mock-script-tool")
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		t.Errorf("expected binary at %s was not created by install script", binPath)
	}

	if err := installer.Remove("mock-script-tool"); err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	if _, err := os.Stat(binPath); !os.IsNotExist(err) {
		t.Errorf("binary at %s should have been removed by uninstall script/cleanup", binPath)
	}
}

func TestSmartInstallationValidationFailure(t *testing.T) {
	tmpDir := t.TempDir()
	pkgDir := filepath.Join(tmpDir, "invalid-pkg")
	_ = os.MkdirAll(pkgDir, 0755)

	manifestContent := `
[package]
name = "invalid-tool"
version = "0.1.0"
`
	_ = os.WriteFile(filepath.Join(pkgDir, "lown.toml"), []byte(manifestContent), 0644)

	m, err := manifest.LoadFromDir(pkgDir)
	if err == nil {
		t.Errorf("expected manifest validation to fail for missing language & script")
	}

	if m != nil {
		_, err = BuildAndInstall(pkgDir, m)
		if err == nil {
			t.Errorf("expected BuildAndInstall to fail")
		}
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

	resolvedGH := installer.cfg.ResolveURI("gh:user/repo")
	expectedGH := "https://github.com/user/repo.git"
	if resolvedGH != expectedGH {
		t.Errorf("gh: resolution failed: expected %s, got %s", expectedGH, resolvedGH)
	}

	fmt.Println("Config resolution tests passed successfully.")
}
