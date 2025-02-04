package functions

import (
	"github.com/eulixir/workspace-automations/config"
	"github.com/robfig/cron/v3"
)

func SetupCronJob(cfg *config.Config) {
	c := cron.New()

	c.AddFunc("0 6 * * *", func() { MorningRoutine(cfg) })
	c.AddFunc("0 18 * * *", func() { NightRoutine(cfg) })

	c.Start()
}
