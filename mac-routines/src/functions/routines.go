package functions

import (
	"github.com/eulixir/workspace-automations/config"
	"go.uber.org/zap"
)

func MorningRoutine(cfg *config.Config, logger *zap.Logger) {
	morningWallpaper := cfg.Wallpaper.MorningWallpaper
	theme := cfg.CodeEditor.MorningTheme
	codeWallpaper := cfg.CodeEditor.MorningBackground

	logger.Info("Morning routine...")

	SetMacOSTheme(false, logger)

	err := RunCodeEditorChanges(cfg, theme, codeWallpaper)
	if err != nil {
		logger.Error("Error updating theme:", zap.Error(err))
	}
	SetWallpaper(morningWallpaper, logger)

}

func NightRoutine(cfg *config.Config, logger *zap.Logger) {
	nightWallpaper := cfg.Wallpaper.NightWallpaper
	theme := cfg.CodeEditor.NightTheme
	codeWallpaper := cfg.CodeEditor.NightBackground

	logger.Info("Night routine...")

	SetMacOSTheme(true, logger)

	err := RunCodeEditorChanges(cfg, theme, codeWallpaper)
	if err != nil {
		logger.Error("Error updating theme:", zap.Error(err))
	}

	SetWallpaper(nightWallpaper, logger)
}
