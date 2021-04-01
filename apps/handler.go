package apps

import (
	"github.com/Mrs4s/MiraiGo/client"
	// "github.com/Mrs4s/MiraiGo/message"
	"github.com/Mrs4s/go-cqhttp/coolq"
	// "github.com/StrayCamel247/BotCamel/apps/camel"
	"github.com/StrayCamel247/BotCamel/apps/destiny"
	"github.com/StrayCamel247/BotCamel/apps/utils"
	"gorm.io/gorm"
)

var dbOrm *gorm.DB

func init() {

}

// CamelBot=结构体实例
type CamelBot struct {
	db  *utils.CamelOrm
	bot *coolq.CQBot
}

// mod模块实例
type mod interface {
}

// 触发
func Start(bot *coolq.CQBot) {

	// 初始化机器基本功能
	Camel := CamelBot{db: utils.Orm, bot: bot}
	// 初始化命运2基础功能
	// 初始化时检查命运2数据库是否存在
	destiny.Start()
	// 命运2群聊天功能启动
	Camel.bot.Client.OnGroupMessage(destiny.GroupMessageEvent)
}

// 收到加群邀请
func GroReciveInviteEvent(c *client.QQClient, e *client.GroupInvitedRequest) {
	c.SolveGroupJoinRequest(e, true, false, "")
}
