package apps

import (
	"github.com/Mrs4s/go-cqhttp/coolq"
)

func init() {

}

// CamelBot
type CamelBot struct {
}

// 触发
func Start(bot *coolq.CQBot) {
	// 载入自定义模块
	bot.Client.OnGroupMessage(GroupMessageEvent)
}
