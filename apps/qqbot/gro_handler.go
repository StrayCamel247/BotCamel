package qqbot

/*
  __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-10
*/
import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

// GroMsgHandler 群聊信息获取并返回
func GroMsgHandler(c *client.QQClient, msg *message.GroupMessage) {
	out := BaseAutoreply(msg.ToString())
	if out == "" {
		return
	}
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(msg.GroupCode, m)
}
