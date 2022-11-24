package master

import (
	. "github.com/sslime336/awbot/bot/export"
	"github.com/sslime336/awbot/config"
	"github.com/sslime336/awbot/utils"
)

func EveryDayWakeUpGreeting() {
	Bot.SendPrivateMessage(
		config.Master.Uin,
		utils.NewText(config.Master.Salutation+" "+config.Master.GreetingText))
}
