package bot

import (
	"github.com/robfig/cron"
	"github.com/sslime336/awbot/utils"
)

var cronManager = cron.New()

// Start 开始定时任务
func StartScheduleMissions() {
	var err error
	for spec, mission := range Bot.CronList {
		err = cronManager.AddFunc(spec, mission)
		utils.Check(err)
	}

	cronManager.Start()
}
