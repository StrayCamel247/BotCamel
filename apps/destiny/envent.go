package destiny

// 命运2插件
import (
	// "fmt"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/camel"
	"github.com/StrayCamel247/BotCamel/apps/utils"
)

func init() {
}

// ========对外事件

// 群消息处理
func GroupMessageEvent(c *client.QQClient, m *message.GroupMessage) {
	// isAt, com, content := camel.AnalysisMsg(c, m.Elements)
	// if isAt {
	// 	handler := Destiny{utils.Orm, c, m}
	// 	go handler.groMsgHandler(com, content)
	// }
	_, com, content := camel.AnalysisMsg(c, m.Elements)
	handler := Destiny{utils.Orm, c, m}
	go handler.groMsgHandler(com, content)
}

// ===========时间处理
// 命运2群聊处理
func (r *Destiny) groMsgHandler(com, content string) {
	// 基础消息返回
	// out := camel.BaseAutoreply(com)
	// 判断指令处理消息
	switch {
	case utils.EqualFolds(com, Commands.D2pvp.Key()):
		r.pvpInfoHandler(content)

	case utils.EqualFolds(com, Commands.D2pve.Key()):
		r.pveInfoHandler(content)

	case utils.EqualFolds(com, Commands.D2skill.Key()):
		r.GenerateDes(content, "skil")

	case utils.EqualFolds(com, Commands.D2npc.Key()):
		r.GenerateDes(content, "npc")

	case utils.EqualFolds(com, Commands.D2perk.Key()):
		r.GenerateDes(content, "perk")

	case utils.EqualFolds(com, Commands.D2item.Key()):
		r.ItemGenerateImg(content, "item")

	case utils.EqualFolds(com, Commands.D2day.Key()):
		r.dayGenerateImg("0x03")

	case utils.EqualFolds(com, Commands.D2week.Key()):
		r.d2uploadImgByFlag("0x02")

	case utils.EqualFolds(com, Commands.D2xiu.Key()):
		r.d2uploadImgByFlag("0x04")

	case utils.EqualFolds(com, Commands.D2trial.Key()):
		r.d2uploadImgByFlag("0x05")

	case utils.EqualFolds(com, Commands.D2dust.Key()):
		r.d2uploadImgByFlag("0x06")

	case utils.EqualFolds(com, Commands.D2random.Key()):
		r.randomHandler()
		// case out == "":
		// 	r.Cli.SendGroupMessage(r.Mes.GroupCode, message.NewSendingMessage().Append(message.NewText(fmt.Sprintf("%s", camel.BaseAutoreply("0x00")))))
		// default:
		// 	r.Cli.SendGroupMessage(r.Mes.GroupCode, message.NewSendingMessage().Append(message.NewText(fmt.Sprintf("%s", out))))
	}
}
