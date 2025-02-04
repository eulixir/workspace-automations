package main

import (
	"log"
	"time"

	"github.com/eulixir/workspace-automations/config"
	"github.com/eulixir/workspace-automations/src/functions"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting daemon...")

	functions.ChangeThemeOnStart(cfg)
	functions.SetupCronJob(cfg)

	for {
		time.Sleep(1 * time.Hour)
	}
}
