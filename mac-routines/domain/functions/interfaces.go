package functions

import "go.uber.org/zap"

type ThemeManager interface {
	SetMacOSTheme(isDarkMode bool)
	ChangeThemeOnStart()
}

type WallpaperManager interface {
	SetWallpaper(imagePath string)
}

type CodeEditorManager interface {
	UpdateEditorSettings(theme, wallpaper string) error
}

type RoutineManager struct {
	logger           *zap.Logger
	themeManager     ThemeManager
	wallpaperManager WallpaperManager
	codeEditor       CodeEditorManager
	settings         RoutineSettings
}

type RoutineSettings struct {
	MorningWallpaper string
	NightWallpaper   string
	CodeMorningTheme string
	CodeNightTheme   string
	CodeMorningBg    string
	CodeNightBg      string
}

func NewRoutineManager(
	logger *zap.Logger,
	themeManager ThemeManager,
	wallpaperManager WallpaperManager,
	codeEditor CodeEditorManager,
	settings RoutineSettings,
) *RoutineManager {
	return &RoutineManager{
		logger:           logger,
		themeManager:     themeManager,
		wallpaperManager: wallpaperManager,
		codeEditor:       codeEditor,
		settings:         settings,
	}
}
