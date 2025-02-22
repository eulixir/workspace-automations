package main

import (
	"os"
	"time"

	"github.com/eulixir/workspace-automations/config"
	"github.com/eulixir/workspace-automations/domain/functions/cronjob"
	"github.com/eulixir/workspace-automations/domain/functions/editor"
	"github.com/eulixir/workspace-automations/domain/functions/macos"
	"github.com/eulixir/workspace-automations/domain/functions/routines"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Error loading config:", zap.Error(err))
		os.Exit(1)
	}

	routineManager := initInterfaces(cfg, logger)

	routineManager.ChangeThemeOnStart()

	cronjob.SetupCronJob(routineManager)

	for {
		time.Sleep(1 * time.Hour)
	}
}

func initInterfaces(cfg *config.Config, logger *zap.Logger) *routines.RoutineManager {

	osManager := macos.NewOsManager(logger)

	codeEditor := editor.NewEditorManager(cfg.CodeEditor.SettingsPath, logger)

	settings := routines.Settings{
		MorningWallpaper: cfg.Wallpaper.MorningWallpaper,
		NightWallpaper:   cfg.Wallpaper.NightWallpaper,
		CodeMorningTheme: cfg.CodeEditor.MorningTheme,
		CodeNightTheme:   cfg.CodeEditor.NightTheme,
		CodeMorningBg:    cfg.CodeEditor.MorningBackground,
		CodeNightBg:      cfg.CodeEditor.NightBackground,
	}

	return &routines.RoutineManager{
		Logger:     logger,
		OsManager:  *osManager,
		CodeEditor: *codeEditor,
		Settings:   settings,
	}
}
