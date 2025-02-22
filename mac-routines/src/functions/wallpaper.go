package functions

import (
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

func SetWallpaper(imagePath string, logger *zap.Logger) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to set picture of every desktop to \"%s\"", imagePath))
	err := cmd.Run()
	if err != nil {
		logger.Error("Error changing wallpaper:", zap.Error(err))
	}
}
