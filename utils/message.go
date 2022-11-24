package utils

import "github.com/Mrs4s/MiraiGo/message"

// NewText 构建一个只包含文本的 SendingMessage
func NewText(msg string) *message.SendingMessage {
	return message.NewSendingMessage().Append(message.NewText(msg))
}

func NewAtAll() *message.SendingMessage {
	return message.NewSendingMessage().Append(message.AtAll())
}

// NewAtAllWithText 联合 @all 和 信息，option 为 1 时，@all 在 信息后面，0 则先 @all
func NewAtAllWithText(msg string, option ...int) (retmsg *message.SendingMessage) {
	retmsg = new(message.SendingMessage)
	atall, txt := message.AtAll(), message.NewText(msg)
	space := message.NewText(" ")

	if option == nil || len(option) < 1 || option[0] == 0 {
		retmsg = message.NewSendingMessage().Append(atall).Append(space).Append(txt)
	} else if option[1] == 1 {
		retmsg = message.NewSendingMessage().Append(txt).Append(space).Append(atall)
	}
	return
}
