package routines

import (
	"time"

	"github.com/eulixir/workspace-automations/domain/functions/editor"
	"github.com/eulixir/workspace-automations/domain/functions/macos"
	"go.uber.org/zap"
)

type RoutineManager struct {
	Logger     *zap.Logger
	OsManager  macos.OsManager
	CodeEditor editor.EditorManager
	Settings   Settings
}

type Settings struct {
	MorningWallpaper string
	NightWallpaper   string
	CodeMorningTheme string
	CodeNightTheme   string
	CodeMorningBg    string
	CodeNightBg      string
}

func (r *RoutineManager) ChangeThemeOnStart() {
	currentHour := time.Now().Hour()

	if currentHour >= 6 && currentHour < 18 {
		r.MorningRoutine()
	} else {
		r.NightRoutine()
	}
}
