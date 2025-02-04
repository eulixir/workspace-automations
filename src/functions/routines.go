package functions

import (
	"log"

	"github.com/eulixir/workspace-automations/config"
)

var currentTheme string

func MorningRoutine(cfg *config.Config) string {
	morningWallpaper := cfg.Wallpaper.MorningWallpaper
	log.Println("Morning routine...")

	SetMacOSTheme(false)
	// updateBackground(morningBackground)
	SetWallpaper(morningWallpaper)
	currentTheme = "light"

	return "light"
}

func NightRoutine(cfg *config.Config) string {
	nightWallpaper := cfg.Wallpaper.NightWallpaper
	log.Println("Night routine...")
	SetMacOSTheme(true)
	// updateBackground(nightBackground)
	SetWallpaper(nightWallpaper)
	return "dark"
}
