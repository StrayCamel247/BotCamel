package apps

import (
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/StrayCamel247/BotCamel/apps/camel"
	"github.com/StrayCamel247/BotCamel/apps/destiny"
	"github.com/robfig/cron"
)

func init() {

}

// CamelBot
type CamelBot struct {
	clickElements map[string][]func()
}

// 触发
func Start(bot *coolq.CQBot) {
	// 载入自定义模块
	bot.Client.OnGroupMessage(GroupMessageEvent)
	// 定时任务=每周三-1点-20分
	c := cron.New()
	c.AddFunc("*0 20 1 * * 3", func() {
		go camel.RefreshDayHandler("0x02", destiny.DataInfo("0x02"))
		// 检查数据库更新
		destiny.InfoMenifestBaseDBCheck(dbGorm)
	})
	// 定时任务=每天-1点-20分触发
	c.AddFunc("*0 20 1 * * *", func() {
		go camel.RefreshDayHandler("0x03", camel.DayGenUrl)
		go camel.RefreshDayHandler("0x04", destiny.DataInfo("0x04"))
		go camel.RefreshDayHandler("0x05", destiny.DataInfo("0x05"))
		go camel.RefreshDayHandler("0x06", destiny.DataInfo("0x06"))
	})
	c.Start()
}
