package theme

import (
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

type MacOSManager struct {
	logger *zap.Logger
}

func NewMacOSManager(logger *zap.Logger) *MacOSManager {
	return &MacOSManager{logger: logger}
}

func (m *MacOSManager) SetMacOSTheme(isDarkMode bool) {
	mode := "false"
	if isDarkMode {
		mode = "true"
	}

	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to tell appearance preferences to set dark mode to %s", mode))
	err := cmd.Run()
	if err != nil {
		m.logger.Error("Error changing theme:", zap.Error(err))
	}
}
