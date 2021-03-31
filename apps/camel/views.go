package camel

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/utils"
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
	case utils.EqualFolds(com, command.D2pvp.Keys):
		PvPInfoHandler(content, r.Client, r.Message)

	case utils.EqualFolds(com, command.D2pve.Keys):
		PvEInfoHandler(content, r.Client, r.Message)

	case utils.EqualFolds(com, command.D2skill.Keys):
		GenerateDes(content, "skil", r.Client, r.Message, r.Orm)

	case utils.EqualFolds(com, command.D2npc.Keys):
		GenerateDes(content, "npc", r.Client, r.Message, r.Orm)

	case utils.EqualFolds(com, command.D2perk.Keys):
		GenerateDes(content, "perk", r.Client, r.Message, r.Orm)

	case utils.EqualFolds(com, command.D2item.Keys):
		ItemGenerateImg(content, "item", r.Client, r.Message, r.Orm)

	case utils.EqualFolds(com, command.D2day.Keys):
		dayGenerateImg("0x03", r.Client, r.Message)

	case utils.EqualFolds(com, command.D2week.Keys):
		d2uploadImgByFlag("0x02", r.Client, r.Message)

	case utils.EqualFolds(com, command.D2xiu.Keys):
		d2uploadImgByFlag("0x04", r.Client, r.Message)

	case utils.EqualFolds(com, command.D2trial.Keys):
		d2uploadImgByFlag("0x05", r.Client, r.Message)

	case utils.EqualFolds(com, command.D2dust.Keys):
		d2uploadImgByFlag("0x06", r.Client, r.Message)

	case utils.EqualFolds(com, command.D2random.Keys):
		randomHandler(r.Client, r.Message)
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
