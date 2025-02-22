package wallpaper

import (
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

type MacOSManager struct {
	logger *zap.Logger
}

func NewMacOSManager(logger *zap.Logger) *MacOSManager {
	return &MacOSManager{logger: logger}
}

func (m *MacOSManager) SetWallpaper(imagePath string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to set picture of every desktop to \"%s\"", imagePath))
	err := cmd.Run()
	if err != nil {
		m.logger.Error("Error changing wallpaper:", zap.Error(err))
	}
}
