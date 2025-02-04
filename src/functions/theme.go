package functions

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	// currentTheme := GetCurrentTheme()

	// if currentHour >= 6 && currentHour < 18 {
	// 	if currentTheme == "dark" {
	// 		MorningRoutine(cfg)
	// 	}
	// } else {
	// 	if currentTheme == "light" {
	// 		NightRoutine(cfg)
	// 	}
	// }
	if currentHour >= 6 && currentHour < 18 {
		MorningRoutine(cfg)
	} else {
		NightRoutine(cfg)
	}
}

func GetCurrentTheme() string {
	cmd := exec.Command("osascript", "-e", "tell application \"System Events\" to tell appearance preferences to get dark mode")
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error getting current theme:", err)
		return "false"
	}

	theme := strings.TrimSpace(string(output))
	if theme == "true" {
		return "dark"
	}

	return "light"
}
