package camel

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"gorm.io/gorm"
)

type BaseMsg struct {
	// 基础消息片段
	Orm     *gorm.DB
	Client  *client.QQClient
	Message *message.GroupMessage
}

func (r *BaseMsg) GroMsgHandler(com, content string) {
	// 基础消息返回
	out := BaseAutoreply(com)
	// 判断指令处理消息
	switch {

	case out == "":
		r.Client.SendGroupMessage(r.Message.GroupCode, message.NewSendingMessage().Append(message.NewText(fmt.Sprintf("%s", BaseAutoreply("0x00")))))
	default:
		r.Client.SendGroupMessage(r.Message.GroupCode, message.NewSendingMessage().Append(message.NewText(fmt.Sprintf("%s", out))))
	}
}

// 加入群聊
func GroJoinHandler(c *client.QQClient, group *client.GroupInfo) {
	out := BaseAutoreply("f48dcc50457d") + "\n"
	out += BaseAutoreply("0x00") + "\n"
	out += BaseAutoreply("menu")
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(group.Code, m)
}
