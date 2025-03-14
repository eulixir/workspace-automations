package routines

import (
	"go.uber.org/zap"
)

func (r *RoutineManager) MorningRoutine() {
	r.Logger.Info("Starting morning routine...")

	r.OsManager.SetMacOSTheme(false)

	err := r.CodeEditor.UpdateEditorSettings(
		r.Settings.CodeMorningTheme,
		r.Settings.CodeMorningBg,
	)
	if err != nil {
		r.Logger.Error("Error updating theme:", zap.Error(err))
	}

	r.OsManager.SetWallpaper(r.Settings.MorningWallpaper)
}
