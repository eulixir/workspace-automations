package main

import (
	"log"
	"time"

	"github.com/eulixir/workspace-automations/config"
	"github.com/eulixir/workspace-automations/src/functions"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Starting daemon...")

	functions.ChangeThemeOnStart(cfg, logger)
	functions.SetupCronJob(cfg, logger)

	for {
		time.Sleep(1 * time.Hour)
	}
}
