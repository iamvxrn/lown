package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestManifestValidation(t *testing.T) {
	tests := []struct {
		name    string
		toml    string
		wantErr bool
	}{
		{
			name: "Valid Go Native",
			toml: `
[package]
name = "my-go-tool"
version = "0.1.0"
language = "go"
executable = "my-go-tool"
`,
			wantErr: false,
		},
		{
			name: "Valid Rust Native",
			toml: `
[package]
name = "my-rust-tool"
version = "0.1.0"
language = "rust"
`,
			wantErr: false,
		},
		{
			name: "Valid Script Fallback",
			toml: `
[package]
name = "script-tool"
version = "0.1.0"

[scripts]
install = "install.sh"
uninstall = "uninstall.sh"
`,
			wantErr: false,
		},
		{
			name: "Invalid No Native No Script",
			toml: `
[package]
name = "empty-tool"
version = "0.1.0"
`,
			wantErr: true,
		},
		{
			name: "Unsupported Language No Script",
			toml: `
[package]
name = "py-tool"
version = "0.1.0"
language = "python"
`,
			wantErr: true,
		},
		{
			name: "Missing Package Name",
			toml: `
[package]
version = "0.1.0"
language = "go"
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			manifestPath := filepath.Join(tmpDir, "lown.toml")
			if err := os.WriteFile(manifestPath, []byte(tt.toml), 0644); err != nil {
				t.Fatalf("failed to write test manifest: %v", err)
			}

			m, err := LoadFromFile(manifestPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if m.Package.Name == "" {
					t.Errorf("Expected non-empty package name")
				}
			}
		})
	}
}

func TestDefaultExecutableName(t *testing.T) {
	m := &Manifest{
		Package: PackageConfig{
			Name: "auto-name",
		},
	}

	if m.GetExecutable() != "auto-name" {
		t.Errorf("Expected default executable 'auto-name', got '%s'", m.GetExecutable())
	}

	m.Package.Executable = "custom-exec"
	if m.GetExecutable() != "custom-exec" {
		t.Errorf("Expected custom executable 'custom-exec', got '%s'", m.GetExecutable())
	}
}
