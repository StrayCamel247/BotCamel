package apps

/*
	Bot Camel插件
*/
import (
	// "encoding/hex"
	// "io/ioutil"
	// "path"
	// "strconv"
	// "strings"
	// "time"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/camel"
	"github.com/StrayCamel247/BotCamel/apps/destiny"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var format = "string"
var dbGorm *gorm.DB

func init() {
	log.Infof("载入sqlite数据库...")
	// 启用sqlite数据库
	dbGorm, _ = gorm.Open(sqlite.Open("./data/sqlite3.db"), &gorm.Config{
		// PrepareStmt: true,
		// 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。但这会降低性能，你可以在初始化时禁用这种方式
		SkipDefaultTransaction: true,
	})
	// 初始化时检查命运2数据库是否存在
	destiny.InfoMenifestBaseDBCheck(dbGorm)
}

// SetMessageFormat 设置消息上报格式，默认为string
func SetMessageFormat(f string) {
	format = f
}

// 群消息处理
func GroupMessageEvent(c *client.QQClient, m *message.GroupMessage) {
	isAt, com, content := camel.AnalysisMsg(c, m.Elements)
	if isAt {
		camel.GroMsgHandler(dbGorm, c, m, com, content)
	}
}
