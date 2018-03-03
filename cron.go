package main


import (
	"github.com/robfig/cron"
	"github.com/olling/slog"
)


var (
	cBreadDate = cron.New()
	cRemind = cron.New()
	cNextDate = cron.New()
)

func InitializeCron() {
	slog.PrintTrace("Func called: initializeCron")
	slog.PrintDebug("Initializing Cron")
	ReloadCron()
}

func NextDate () {
	slog.PrintInfo(cNextDate.Entries()[0].Schedule)
}

func ReloadCron() {
	slog.PrintTrace("Func called: ReloadCron")

	cBreadDate.Stop()
	cRemind.Stop()
	cNextDate.Stop()

	cBreadDate = cron.New()
	cBreadDate.AddFunc(CurrentConfiguration.CronBreadDate,BreadDate)

	cRemind = cron.New()
	cRemind.AddFunc(CurrentConfiguration.CronRemind,Remind)

	cNextDate = cron.New()
	cNextDate.AddFunc(CurrentConfiguration.CronNext,GetNext)
	slog.PrintTrace("cronneext",CurrentConfiguration.CronNext)

	cBreadDate.Start()
	cRemind.Start()
	cNextDate.Start()
}
