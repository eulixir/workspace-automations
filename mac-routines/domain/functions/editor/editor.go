package editor

import "go.uber.org/zap"

type EditorManager struct {
	settingsPath string
	logger       *zap.Logger
}

func NewEditorManager(settingsPath string, logger *zap.Logger) *EditorManager {
	return &EditorManager{settingsPath: settingsPath, logger: logger}
}
