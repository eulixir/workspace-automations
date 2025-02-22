package functions

import (
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

type MacOSWallpaperManager struct {
	logger *zap.Logger
}

func NewMacOSWallpaperManager(logger *zap.Logger) *MacOSWallpaperManager {
	return &MacOSWallpaperManager{logger: logger}
}

func (m *MacOSWallpaperManager) SetWallpaper(imagePath string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to set picture of every desktop to \"%s\"", imagePath))
	err := cmd.Run()
	if err != nil {
		m.logger.Error("Error changing wallpaper:", zap.Error(err))
	}
}
