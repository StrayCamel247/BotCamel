package apps

import (
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/StrayCamel247/BotCamel/apps/camel"
	"github.com/StrayCamel247/BotCamel/apps/destiny"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func init() {

}

// CamelBot=结构体实例
type CamelBot struct {
	db  *gorm.DB
	bot *coolq.CQBot
}

// mod模块实例
type mod interface {
}

// 触发
func Start(bot *coolq.CQBot) {
	log.Infof("载入sqlite数据库...")
	// 启用sqlite数据库
	dbGorm, _ = gorm.Open(sqlite.Open("./data/sqlite3.db"), &gorm.Config{
		// PrepareStmt: true,
		// 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。但这会降低性能，你可以在初始化时禁用这种方式
		SkipDefaultTransaction: true,
	})
	Camel := CamelBot{db: dbGorm}
	// 初始化时检查命运2数据库是否存在
	destiny.InfoMenifestBaseDBCheck(Camel.db)
	Camel.handlerMessage()
	Camel.clickHandler()
}

// handlerMessage 处理消息
func (self *CamelBot) handlerMessage() {
	// 载入自定义模块
	self.bot.Client.OnGroupMessage(GroupMessageEvent)
}

// clickHandler 定时任务
func (self *CamelBot) clickHandler() {
	c := cron.New()
	// 定每周三-1点-20分
	c.AddFunc("*0 20 1 * * 3", func() {
		go camel.RefreshDayHandler("0x02", destiny.DataInfo("0x02"))
		// 检查数据库更新
		destiny.InfoMenifestBaseDBCheck(dbGorm)
	})
	// 每天-1点-20分触发
	c.AddFunc("*0 20 1 * * *", func() {
		go camel.RefreshDayHandler("0x03", camel.DayGenUrl)
		go camel.RefreshDayHandler("0x04", destiny.DataInfo("0x04"))
		go camel.RefreshDayHandler("0x05", destiny.DataInfo("0x05"))
		go camel.RefreshDayHandler("0x06", destiny.DataInfo("0x06"))
	})
	c.Start()
}
