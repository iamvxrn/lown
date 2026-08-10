package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"

	"lown/internal/path"
)

// Config represents global configuration in ~/.config/lown/config.toml
type Config struct {
	Aliases map[string]string `toml:"aliases"`
}

// NewConfig returns default configuration.
func NewConfig() *Config {
	return &Config{
		Aliases: make(map[string]string),
	}
}

// LoadConfig reads configuration from ~/.config/lown/config.toml
func LoadConfig() (*Config, error) {
	cfgPath := path.ConfigPath()
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		return NewConfig(), nil
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", cfgPath, err)
	}

	cfg := NewConfig()
	if err := toml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", cfgPath, err)
	}

	if cfg.Aliases == nil {
		cfg.Aliases = make(map[string]string)
	}

	return cfg, nil
}

// ResolveURI converts aliases or shorthand syntax (e.g. gh:owner/repo) into full Git URIs.
func (c *Config) ResolveURI(input string) string {
	input = strings.TrimSpace(input)
	if resolved, ok := c.Aliases[input]; ok {
		input = strings.TrimSpace(resolved)
	}

	// Handle gh:owner/repo format
	if strings.HasPrefix(input, "gh:") {
		repoPath := strings.TrimPrefix(input, "gh:")
		return fmt.Sprintf("https://github.com/%s.git", strings.Trim(repoPath, "/"))
	}

	// Handle owner/repo format if not local path or URL
	if !strings.Contains(input, "://") && !strings.HasPrefix(input, "git@") && !strings.HasPrefix(input, "/") && !strings.HasPrefix(input, ".") && strings.Count(input, "/") == 1 {
		return fmt.Sprintf("https://github.com/%s.git", input)
	}

	return input
}
