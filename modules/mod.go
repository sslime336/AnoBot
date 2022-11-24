package modules

import (
	. "github.com/sslime336/awbot/bot/export"
	"github.com/sslime336/awbot/res"

	"github.com/sslime336/awbot/modules/groups"
	"github.com/sslime336/awbot/modules/master"
	"github.com/sslime336/awbot/modules/ping"
)

func Load() {
	ping.Ping()
	master.ListenShutdownCmd()

	// 早安
	AddSchedule(res.M.Cron["daily_wake_up"], master.EveryDayWakeUpGreeting)

	// 群聊天气预报
	AddSchedule(res.M.Cron["daily_weather_report"], groups.DailyWeatherReport)
}
