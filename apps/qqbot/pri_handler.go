package qqbot

/*
   __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-10
*/
import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

// PriMsgHandler 私人聊天处理
func PriMsgHandler(c *client.QQClient, msg *message.PrivateMessage) {
	out := BaseAutoreply(msg.ToString())
	// c.SolveGroupJoinRequest(c.UserJoinGroupRequest, true, false, "")
	// fmt.Printf("%+v", c)
	if out == "" {
		return
	}
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendPrivateMessage(msg.Sender.Uin, m)
}
