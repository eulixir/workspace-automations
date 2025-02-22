package routines

import (
	"go.uber.org/zap"
)

func (r *RoutineManager) NightRoutine() {
	r.Logger.Info("Starting night routine...")

	r.OsManager.SetMacOSTheme(true)

	err := r.CodeEditor.UpdateEditorSettings(
		r.Settings.CodeNightTheme,
		r.Settings.CodeNightBg,
	)
	if err != nil {
		r.Logger.Error("Error updating theme:", zap.Error(err))
	}

	r.OsManager.SetWallpaper(r.Settings.NightWallpaper)
}
