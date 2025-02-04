package main

import (
	"log"
	"time"

	"github.com/eulixir/workspace-automations/config"
	"github.com/eulixir/workspace-automations/src/functions"
	"github.com/robfig/cron/v3"
)

var currentTheme string

func main() {
	// logFile, err := os.OpenFile("/Users/seuusuario/theme_switcher.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer logFile.Close()
	// log.SetOutput(logFile)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting daemon...")

	functions.ChangeThemeOnStart(cfg)

	c := cron.New()

	c.AddFunc("0 06 * * *", func() { functions.MorningRoutine(cfg) })
	c.AddFunc("0 18 * * *", func() { functions.NightRoutine(cfg) })

	c.Start()

	for {
		time.Sleep(1 * time.Hour)
	}
}
