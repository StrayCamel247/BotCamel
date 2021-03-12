package camel

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

/*
   __author__ : stray_camel
  __description__ :私人聊天功能
  __REFERENCES__:
  __date__: 2021-03-12
*/

func SoloMsgHandler(c *client.QQClient, msg *message.PrivateMessage) {
	out := BaseAutoreply(msg.ToString())
	if out == "" {
		return
	}
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendPrivateMessage(msg.Sender.Uin, m)
}
