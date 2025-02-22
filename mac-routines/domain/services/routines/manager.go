package routines

import (
	"github.com/eulixir/workspace-automations/domain/functions/editor"
	"github.com/eulixir/workspace-automations/domain/functions/macos"
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
	Logger     *zap.Logger
	OsManager  macos.OsManager
	CodeEditor editor.EditorManager
	Settings   Settings
}

func NewManager(
	logger *zap.Logger,
	osManager macos.OsManager,
	codeEditor editor.EditorManager,
	settings Settings,
) *Manager {
	return &Manager{
		Logger:     logger,
		OsManager:  osManager,
		CodeEditor: codeEditor,
		Settings:   settings,
	}
}
