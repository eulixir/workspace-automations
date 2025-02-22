package functions

import (
	"fmt"
	"log"
	"os/exec"
)

func SetWallpaper(imagePath string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to set picture of every desktop to \"%s\"", imagePath))
	err := cmd.Run()
	if err != nil {
		log.Println("Error changing wallpaper:", err)
	}
}
