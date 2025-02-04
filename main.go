package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

const configFile = "/Users/seuusuario/Library/Application Support/SeuApp/settings.json"

const morningBackground = "https://seu-link-matinal.jpg"
const nightBackground = "https://seu-link-noturno.jpg"

const morningWallpaper = "https://github.com/eulixir/workspace-automations/blob/main/assets/light.jpg"
const nightWallpaper = "https://github.com/eulixir/workspace-automations/blob/main/assets/dark.jpg"

func setMacOSTheme(isDarkMode bool) {
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

func updateBackground(url string) {
	cmd := exec.Command("jq", "--arg newURL", url, ".background = $newURL", configFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error updating JSON:", err, string(output))
	}
}

func setWallpaper(imagePath string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to set picture of every desktop to \"%s\"", imagePath))
	err := cmd.Run()
	if err != nil {
		log.Println("Error changing wallpaper:", err)
	}
}

func morningRoutine() {
	log.Println("Morning routine...")
	setMacOSTheme(false)
	// updateBackground(morningBackground)
	setWallpaper(morningWallpaper)
}

func nightRoutine() {
	log.Println("Night routine...")
	setMacOSTheme(true)
	// updateBackground(nightBackground)
	setWallpaper(nightWallpaper)
}

func main() {
	// logFile, err := os.OpenFile("/Users/seuusuario/theme_switcher.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer logFile.Close()
	// log.SetOutput(logFile)

	log.Println("Iniciando o daemon...")

	changeThemeOnStart()

	c := cron.New()

	c.AddFunc("0 6 * * *", morningRoutine)
	c.AddFunc("0 18 * * *", nightRoutine)

	c.Start()

	for {
		time.Sleep(1 * time.Hour)
	}
}

func changeThemeOnStart() {
	now := time.Now()
	hour := now.Hour()

	if hour >= 6 && hour < 18 {
		morningRoutine()
	} else {
		nightRoutine()
	}
}
