package functions

import (
	"log"

	"github.com/eulixir/workspace-automations/config"
)

func MorningRoutine(cfg *config.Config) {
	morningWallpaper := cfg.Wallpaper.MorningWallpaper
	theme := cfg.CodeEditor.MorningTheme
	codeWallpaper := cfg.CodeEditor.MorningBackground

	log.Println("Morning routine...")

	SetMacOSTheme(false)

	err := RunCodeEditorChanges(cfg, theme, codeWallpaper)
	if err != nil {
		log.Println("Error updating theme:", err)
	}
	SetWallpaper(morningWallpaper)

}

func NightRoutine(cfg *config.Config) {
	nightWallpaper := cfg.Wallpaper.NightWallpaper
	theme := cfg.CodeEditor.NightTheme
	codeWallpaper := cfg.CodeEditor.NightBackground

	log.Println("Night routine...")

	SetMacOSTheme(true)

	err := RunCodeEditorChanges(cfg, theme, codeWallpaper)
	if err != nil {
		log.Println("Error updating theme:", err)
	}

	SetWallpaper(nightWallpaper)
}
