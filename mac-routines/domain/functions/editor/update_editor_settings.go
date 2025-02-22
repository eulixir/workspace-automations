package editor

import (
	"fmt"
)

func (v *EditorManager) UpdateEditorSettings(theme, wallpaper string) error {
	settings, err := v.ReadSettingsRaw()
	if err != nil {
		return fmt.Errorf("error reading settings: %w", err)
	}

	v.UpdateEditorTheme(theme, settings)
	v.UpdateEditorBackground(wallpaper, settings)

	err = v.WriteSettings(settings)
	if err != nil {
		return fmt.Errorf("error writing settings: %w", err)
	}

	return nil
}
