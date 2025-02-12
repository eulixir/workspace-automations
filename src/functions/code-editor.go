package functions

import (
	"encoding/json"
	"fmt"
	"os"
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

func readSettingsRaw(cfg *config.Config) (*map[string]interface{}, error) {
	path, err := expandPath(cfg.CodeEditor.SettingsPath)
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

func RunCodeEditorChanges(cfg *config.Config, theme string, wallpaper string) error {
	settings, err := readSettingsRaw(cfg)
	if err != nil {
		return fmt.Errorf("error reading settings: %w", err)
	}

	updateCodeTheme(theme, settings)
	updateBackground(wallpaper, settings)

	err = writeSettings(settings, cfg)
	if err != nil {
		return fmt.Errorf("error writing settings: %w", err)
	}

	return nil
}

func updateCodeTheme(theme string, settings *map[string]interface{}) {
	(*settings)["workbench.colorTheme"] = theme

}

func updateBackground(url string, settings *map[string]interface{}) {
	bgSettings, ok := (*settings)["background.fullscreen"].(map[string]interface{})
	if !ok {
		bgSettings = map[string]interface{}{
			"opacity":  0.85,
			"position": "center",
			"size":     "cover",
		}
	}

	bgSettings["images"] = []string{url}

	(*settings)["background.fullscreen"] = bgSettings
}

func writeSettings(settings *map[string]interface{}, cfg *config.Config) error {
	jsonData, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling settings: %w", err)
	}

	path, err := expandPath(cfg.CodeEditor.SettingsPath)
	if err != nil {
		return fmt.Errorf("error expanding path: %w", err)
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing settings: %w", err)
	}

	return nil
}
