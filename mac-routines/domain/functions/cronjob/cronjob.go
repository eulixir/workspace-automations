package cronjob

import (
	"github.com/eulixir/workspace-automations/domain/functions/routines"
	"github.com/robfig/cron/v3"
)

func SetupCronJob(routineManager *routines.RoutineManager) {
	c := cron.New()

	routineManager.Logger.Info("Setting up cron job...")

	c.AddFunc("0 6 * * *", routineManager.MorningRoutine)
	c.AddFunc("0 18 * * *", routineManager.NightRoutine)

	routineManager.Logger.Info("Cron job setup complete")

	c.Start()
}
