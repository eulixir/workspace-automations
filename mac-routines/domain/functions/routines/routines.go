package routines

import (
	"time"

	"go.uber.org/zap"
)

type RoutineManager struct {
	Logger           *zap.Logger
	ThemeManager     ThemeManager
	WallpaperManager WallpaperManager
	CodeEditor       CodeEditor
	Settings         Settings
}

type Settings struct {
	MorningWallpaper string
	NightWallpaper   string
	CodeMorningTheme string
	CodeNightTheme   string
	CodeMorningBg    string
	CodeNightBg      string
}

type ThemeManager interface {
	SetMacOSTheme(isDarkMode bool)
}

type WallpaperManager interface {
	SetWallpaper(imagePath string)
}

type CodeEditor interface {
	UpdateEditorSettings(theme, wallpaper string) error
}

func (r *RoutineManager) ChangeThemeOnStart() {
	currentHour := time.Now().Hour()

	if currentHour >= 6 && currentHour < 18 {
		r.MorningRoutine()
	} else {
		r.NightRoutine()
	}
}
