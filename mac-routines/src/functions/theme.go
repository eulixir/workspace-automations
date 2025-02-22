package functions

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/eulixir/workspace-automations/config"
	"go.uber.org/zap"
)

func SetMacOSTheme(isDarkMode bool, logger *zap.Logger) {
	mode := "false"
	if isDarkMode {
		mode = "true"
	}

	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to tell appearance preferences to set dark mode to %s", mode))
	err := cmd.Run()
	if err != nil {
		logger.Error("Error changing theme:", zap.Error(err))
	}
}

func ChangeThemeOnStart(cfg *config.Config, logger *zap.Logger) {
	currentHour := time.Now().Hour()

	if currentHour >= 6 && currentHour < 18 {
		MorningRoutine(cfg, logger)
	} else {
		NightRoutine(cfg, logger)
	}
}
