package routines

import (
	"go.uber.org/zap"
)

func (r *RoutineManager) MorningRoutine() {
	r.Logger.Info("Starting morning routine...")

	r.ThemeManager.SetMacOSTheme(false)

	err := r.CodeEditor.UpdateEditorSettings(
		r.Settings.CodeMorningTheme,
		r.Settings.CodeMorningBg,
	)
	if err != nil {
		r.Logger.Error("Error updating theme:", zap.Error(err))
	}

	r.WallpaperManager.SetWallpaper(r.Settings.MorningWallpaper)
}
