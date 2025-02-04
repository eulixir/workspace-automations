package functions

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/eulixir/workspace-automations/config"
)

func SetMacOSTheme(isDarkMode bool) {
	mode := "false"
	if isDarkMode {
		mode = "true"
	}

	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to tell appearance preferences to set dark mode to %s", mode))
	err := cmd.Run()
	if err != nil {
		log.Println("Error changing theme:", err)
	}
}

func ChangeThemeOnStart(cfg *config.Config) {
	currentHour := time.Now().Hour()

	if currentHour >= 6 && currentHour < 18 {
		MorningRoutine(cfg)
	} else {
		NightRoutine(cfg)
	}
}
