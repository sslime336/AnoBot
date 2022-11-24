// Package ping is used to test the bot
package ping

import (
	"time"

	. "github.com/sslime336/awbot/bot/export"
	"github.com/sslime336/awbot/config"
	"github.com/sslime336/awbot/utils"
)

func Ping() {
	Bot.SendPrivateMessage(
		config.Master.Uin,
		utils.NewText("Bot report: running smoothly + time:"+time.Now().Local().String()),
	)
}
