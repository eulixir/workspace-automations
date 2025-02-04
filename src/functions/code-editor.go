package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/eulixir/workspace-automations/config"
)

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting home directory: %w", err)
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
}

func UpdateTheme(theme string, cfg *config.Config) error {
	path, err := expandPath(cfg.CodeEditor.SettingsPath)
	if err != nil {
		return fmt.Errorf("error expanding path: %w", err)
	}

	cmd := exec.Command("jq", fmt.Sprintf(".\"workbench.colorTheme\" = \"%s\"", theme), path)

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading settings file: %w", err)
	}

	cmd.Stdin = strings.NewReader(string(content))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error updating theme: %w", err)
	}

	err = os.WriteFile(path, output, 0644)
	if err != nil {
		return fmt.Errorf("error writing updated settings: %w", err)
	}

	return nil
}

func UpdateBackground(url string, cfg *config.Config) error {
	path, err := expandPath(cfg.CodeEditor.SettingsPath)
	if err != nil {
		return fmt.Errorf("error expanding path: %w", err)
	}

	cmd := exec.Command("jq", fmt.Sprintf(".\"workbench.background\" = \"%s\"", url), path)

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading settings file: %w", err)
	}

	cmd.Stdin = strings.NewReader(string(content))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error updating background: %w", err)
	}

	err = os.WriteFile(path, output, 0644)
	if err != nil {
		return fmt.Errorf("error writing updated settings: %w", err)
	}

	return nil
}

func ReadSettings(cfg *config.Config) (*SettingsJson, error) {
	path, err := expandPath(cfg.CodeEditor.SettingsPath)
	if err != nil {
		return nil, fmt.Errorf("error expanding path: %w", err)
	}

	path = filepath.Clean(path)

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("settings file does not exist at %s", path)
		}
		if os.IsPermission(err) {
			return nil, fmt.Errorf("permission denied accessing settings at %s", path)
		}
		return nil, fmt.Errorf("error accessing settings file: %w", err)
	}

	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("%s is not a regular file", path)
	}

	settings, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading settings at %s: %w", path, err)
	}

	var settingsJson SettingsJson
	err = json.Unmarshal(settings, &settingsJson)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling settings: %w\nContent: %s", err, settings)
	}

	return &settingsJson, nil
}

type SettingsJson struct {
	WorkbenchColorTheme string `json:"workbench.colorTheme"`
}
