package macos

import (
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

func (o *OsManager) SetWallpaper(imagePath string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to set picture of every desktop to \"%s\"", imagePath))
	err := cmd.Run()
	if err != nil {
		o.logger.Error("Error changing wallpaper:", zap.Error(err))
	}
}
