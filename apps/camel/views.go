package camel

import (
	"time"
	// "github.com/Logiase/gomirai"
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"gorm.io/gorm"
	// "github.com/Mrs4s/MiraiGo/client/pb/structmsg"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/destiny"
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
				go RefreshDayHandler("0x02", destiny.DataInfo("0x02"))
			}
		// 每天触发
		default:
			// 更新日报信息-每天一点开始轮询更新
			if t.Hour() == 1 {
				go RefreshDayHandler("0x03", DayGenUrl)
				go RefreshDayHandler("0x04", destiny.DataInfo("0x04"))
				go RefreshDayHandler("0x05", destiny.DataInfo("0x05"))
				go RefreshDayHandler("0x06", destiny.DataInfo("0x06"))
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
	case handler.EqualFolds(com, command.Menu.Keys):
		menuHandler(c, msg)

	case handler.EqualFolds(com, command.D2pvp.Keys):
		PvPInfoHandler(content, c, msg)

	case handler.EqualFolds(com, command.D2pve.Keys):
		PvEInfoHandler(content, c, msg)

	case handler.EqualFolds(com, command.D2skill.Keys):
		GenerateDes(content, "skil", c, msg, orm)

	case handler.EqualFolds(com, command.D2npc.Keys):
		GenerateDes(content, "npc", c, msg, orm)

	case handler.EqualFolds(com, command.D2perk.Keys):
		GenerateDes(content, "perk", c, msg, orm)

	case handler.EqualFolds(com, command.D2item.Keys):
		ItemGenerateImg(content, "item", c, msg, orm)

	case handler.EqualFolds(com, command.D2day.Keys):
		dayGenerateImg("0x03", c, msg)

	case handler.EqualFolds(com, command.D2week.Keys):
		d2uploadImgByFlag("0x02", c, msg)

	case handler.EqualFolds(com, command.D2xiu.Keys):
		d2uploadImgByFlag("0x04", c, msg)

	case handler.EqualFolds(com, command.D2trial.Keys):
		d2uploadImgByFlag("0x05", c, msg)

	case handler.EqualFolds(com, command.D2dust.Keys):
		d2uploadImgByFlag("0x06", c, msg)

	case handler.EqualFolds(com, command.D2random.Keys):
		randomHandler(c, msg)

	case out == "":
		c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText(fmt.Sprintf("%s", BaseAutoreply("0x00")))))
	default:
		c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText(fmt.Sprintf("%s", out))))
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
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(group.Code, m)
}
