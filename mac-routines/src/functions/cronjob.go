package functions

import (
	"github.com/eulixir/workspace-automations/config"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func SetupCronJob(cfg *config.Config, logger *zap.Logger) {
	c := cron.New()

	c.AddFunc("0 6 * * *", func() { MorningRoutine(cfg, logger) })
	c.AddFunc("0 18 * * *", func() { NightRoutine(cfg, logger) })

	c.Start()
}
