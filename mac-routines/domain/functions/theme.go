package functions

import (
	"fmt"
	"os/exec"
	"time"

	"go.uber.org/zap"
)

type MacOSThemeManager struct {
	logger *zap.Logger
}

func NewMacOSThemeManager(logger *zap.Logger) *MacOSThemeManager {
	return &MacOSThemeManager{logger: logger}
}

func (m *MacOSThemeManager) SetMacOSTheme(isDarkMode bool) {
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

func (m *MacOSThemeManager) ChangeThemeOnStart() {
	currentHour := time.Now().Hour()
	if currentHour >= 6 && currentHour < 18 {
		m.SetMacOSTheme(false)
	} else {
		m.SetMacOSTheme(true)
	}
}
