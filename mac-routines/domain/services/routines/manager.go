package routines

import (
	"github.com/eulixir/workspace-automations/domain/interfaces"
	"go.uber.org/zap"
)

type Settings struct {
	MorningWallpaper string
	NightWallpaper   string
	CodeMorningTheme string
	CodeNightTheme   string
	CodeMorningBg    string
	CodeNightBg      string
}

type Manager struct {
	Logger           *zap.Logger
	ThemeManager     interfaces.ThemeManager
	WallpaperManager interfaces.WallpaperManager
	CodeEditor       interfaces.Editor
	Settings         Settings
}

func NewManager(
	logger *zap.Logger,
	theme interfaces.ThemeManager,
	wallpaper interfaces.WallpaperManager,
	editor interfaces.Editor,
	settings Settings,
) *Manager {
	return &Manager{
		Logger:           logger,
		ThemeManager:     theme,
		WallpaperManager: wallpaper,
		CodeEditor:       editor,
		Settings:         settings,
	}
}
