package editor

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (v *EditorManager) WriteSettings(settings *map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling settings: %w", err)
	}

	path, err := expandPath(v.settingsPath)
	if err != nil {
		return fmt.Errorf("error expanding path: %w", err)
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing settings: %w", err)
	}

	return nil
}

func (v *EditorManager) ReadSettingsRaw() (*map[string]interface{}, error) {
	path, err := expandPath(v.settingsPath)
	if err != nil {
		return nil, fmt.Errorf("error expanding path: %w", err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading settings: %w", err)
	}

	var settings map[string]interface{}
	if err := json.Unmarshal(content, &settings); err != nil {
		return nil, fmt.Errorf("error parsing settings: %w", err)
	}

	return &settings, nil
}

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
