package camel

import (
	"time"
	// "github.com/Logiase/gomirai"
	"github.com/Mrs4s/MiraiGo/client"
	"gorm.io/gorm"
	// "github.com/Mrs4s/MiraiGo/client/pb/structmsg"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/baseapis"
	"github.com/StrayCamel247/BotCamel/apps/handler"
)

func BaseRefreshHandler() {
	for {
		// 每分钟触发一次
		time.Sleep(time.Minute * 1)
		t := time.Now()
		_weekNum := int(t.Weekday())
		switch {
		case _weekNum == 3:
			// 周三触发
			if t.Hour() == 1 {
				go RefreshDayHandler("0x02", baseapis.DataInfo("0x02"))
			}
		// 每天触发
		default:
			// 更新日报信息-每天一点开始轮询更新
			if t.Hour() == 1 {
				go RefreshDayHandler("0x03", DayGenUrl)
				go RefreshDayHandler("0x04", baseapis.DataInfo("0x04"))
				go RefreshDayHandler("0x05", baseapis.DataInfo("0x05"))
				go RefreshDayHandler("0x06", baseapis.DataInfo("0x06"))
			}

		}
	}
}

// GroMsgHandler 群聊信息获取并返回

func GroMsgHandler(orm *gorm.DB, c *client.QQClient, msg *message.GroupMessage, com, content string) {
	var out string
	// 若@机器人则触发
	out = BaseAutoreply(com)
	switch {
	// case
	case handler.EqualFolds(com, command.Menu.Keys):
		go menuHandler(c, msg)

	case handler.EqualFolds(com, command.D2pvp.Keys):
		go PvPInfoHandler(content, c, msg)

	case handler.EqualFolds(com, command.D2pve.Keys):
		go PvEInfoHandler(content, c, msg)

	case handler.EqualFolds(com, command.D2skill.Keys):
		go GenerateDes(content, "skil", c, msg, orm)

	case handler.EqualFolds(com, command.D2npc.Keys):
		go GenerateDes(content, "npc", c, msg, orm)

	case handler.EqualFolds(com, command.D2perk.Keys):
		go GenerateDes(content, "perk", c, msg, orm)

	case handler.EqualFolds(com, command.D2item.Keys):
		go ItemGenerateImg(content, "item", c, msg, orm)

	case handler.EqualFolds(com, command.D2day.Keys):
		go dayGenerateImg("0x03", c, msg)

	case handler.EqualFolds(com, command.D2week.Keys):
		go d2uploadImgByFlag("0x02", c, msg)

	case handler.EqualFolds(com, command.D2xiu.Keys):
		go d2uploadImgByFlag("0x04", c, msg)

	case handler.EqualFolds(com, command.D2trial.Keys):
		go d2uploadImgByFlag("0x05", c, msg)

	case handler.EqualFolds(com, command.D2dust.Keys):
		go d2uploadImgByFlag("0x06", c, msg)

	case handler.EqualFolds(com, command.D2random.Keys):
		go randomHandler(c, msg)
	case out == "":
		out = BaseAutoreply("0x00")
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendGroupMessage(msg.GroupCode, m)

	default:
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendGroupMessage(msg.GroupCode, m)
	}
}

// 收到加群邀请

func GroReciveInviteHandler(c *client.QQClient, e *client.GroupInvitedRequest) {
	c.SolveGroupJoinRequest(e, true, false, "")
}

// 加入群聊

func GroJoinHandler(c *client.QQClient, group *client.GroupInfo) {
	out := BaseAutoreply("f48dcc50457d") + "\n"
	out += BaseAutoreply("0x00") + "\n"
	out += BaseAutoreply("menu")
	out += GroupMenu
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(group.Code, m)
}

// 离开群聊-都被t了怎么发文字，，，，开发了个寂寞

// func GroLeaveHandler(c *client.QQClient, e *client.GroupLeaveEvent) {
// 	if e.Operator != nil {
// 		out := BaseAutoreply("0x01") + "\n"
// 		message.NewSendingMessage().Append(message.NewAt(e.Operator.Uin, e.Operator.Nickname)).Append(message.NewText(out))
// 	} else {
// 		out := BaseAutoreply("0x00") + "\n"
// 		message.NewSendingMessage().Append(message.NewText(out))
// 		// log.Infof("Bot退出了群 %v.", formatGroupName(e.Group))
// 	}
// }
