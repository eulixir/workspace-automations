package macos

import (
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

func (o *OsManager) SetMacOSTheme(isDarkMode bool) {
	mode := "false"
	if isDarkMode {
		mode = "true"
	}

	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to tell appearance preferences to set dark mode to %s", mode))
	err := cmd.Run()
	if err != nil {
		o.logger.Error("Error changing theme:", zap.Error(err))
	}
}
