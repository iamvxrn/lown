package manifest

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

// ManifestFileName is the default manifest filename
const ManifestFileName = "lown.toml"

// PackageConfig holds package metadata in lown.toml
type PackageConfig struct {
	Name       string `toml:"name"`
	Version    string `toml:"version"`
	Language   string `toml:"language"`
	Executable string `toml:"executable"`
}

// ScriptConfig holds custom script hooks in lown.toml
type ScriptConfig struct {
	Install   string `toml:"install"`
	Uninstall string `toml:"uninstall"`
}

// Manifest represents the structure of lown.toml
type Manifest struct {
	Package PackageConfig `toml:"package"`
	Scripts ScriptConfig `toml:"scripts"`
}

// GetExecutable returns the configured executable name or defaults to package name.
func (m *Manifest) GetExecutable() string {
	exec := strings.TrimSpace(m.Package.Executable)
	if exec != "" {
		return exec
	}
	return strings.TrimSpace(m.Package.Name)
}

// GetLanguage returns normalized language (lowercase).
func (m *Manifest) GetLanguage() string {
	return strings.ToLower(strings.TrimSpace(m.Package.Language))
}

// IsNative returns true if language is a supported native language ("go", "rust", "revoq", "c", "cpp").
func (m *Manifest) IsNative() bool {
	lang := m.GetLanguage()
	return lang == "go" || lang == "rust" || lang == "revoq" || lang == "c" || lang == "cpp"
}

// Validate ensures manifest contains valid metadata and install instructions.
func (m *Manifest) Validate() error {
	if strings.TrimSpace(m.Package.Name) == "" {
		return errors.New("lown.toml: [package] name is required")
	}

	lang := m.GetLanguage()
	hasNative := m.IsNative()
	hasInstallScript := strings.TrimSpace(m.Scripts.Install) != ""

	if lang != "" && !hasNative {
		return fmt.Errorf("lown.toml: unsupported language '%s' (supported: go, rust, revoq, c, cpp)", m.Package.Language)
	}

	if !hasNative && !hasInstallScript {
		return fmt.Errorf("lown.toml: package '%s' provides neither a supported language ('go', 'rust', 'revoq', 'c', 'cpp') nor an install script", m.Package.Name)
	}

	if m.GetExecutable() == "" {
		return errors.New("lown.toml: package executable name could not be determined")
	}

	return nil
}

// LoadFromFile parses a lown.toml file at the given path.
func LoadFromFile(filePath string) (*Manifest, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read manifest file %s: %w", filePath, err)
	}

	var m Manifest
	if err := toml.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse manifest TOML %s: %w", filePath, err)
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return &m, nil
}

// LoadFromDir looks for lown.toml in directory and parses it.
func LoadFromDir(dir string) (*Manifest, error) {
	target := filepath.Join(dir, ManifestFileName)
	return LoadFromFile(target)
}
