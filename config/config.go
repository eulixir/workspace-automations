package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	CodeEditor CodeEditor
	Wallpaper  Wallpaper
}

type CodeEditor struct {
	MorningBackground string `envconfig:"MORNING_BACKGROUND"`
	NightBackground   string `envconfig:"NIGHT_BACKGROUND"`

	MorningTheme string `envconfig:"MORNING_THEME"`
	NightTheme   string `envconfig:"NIGHT_THEME"`

	SettingsPath string `envconfig:"CODE_EDITOR_SETTINGS_PATH"`
	Editor       string `envconfig:"EDITOR"`
}

type Wallpaper struct {
	MorningWallpaper string `envconfig:"MORNING_WALLPAPER"`
	NightWallpaper   string `envconfig:"NIGHT_WALLPAPER"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed loading config: %w", err)
	}

	return &cfg, nil
}
