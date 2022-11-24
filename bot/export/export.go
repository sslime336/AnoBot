package export

import "github.com/sslime336/awbot/bot"

var Bot = bot.Bot

func AddSchedule(spec string, cmd func()) {
	Bot.CronList[spec] = cmd
}
