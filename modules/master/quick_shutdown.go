package master

import (
	"os"
	"strings"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	. "github.com/sslime336/awbot/bot/export"
	"github.com/sslime336/awbot/config"
	"github.com/sslime336/awbot/utils"
)

func ListenShutdownCmd() {
	Bot.PrivateMessageEvent.Subscribe(func(client *client.QQClient, event *message.PrivateMessage) {
		senderUin := event.Sender.Uin
		if senderUin != config.Master.Uin {
			return
		}
		if strings.Contains(event.ToString(), "stop") {
			client.SendPrivateMessage(senderUin, utils.NewText("Bot offline"))
			os.Exit(-1)
		}
	})
}
