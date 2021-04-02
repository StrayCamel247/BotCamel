// Bot Camel
//
package apps

import (
	"bytes"
	"fmt"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/StrayCamel247/BotCamel/apps/camel"
	"github.com/StrayCamel247/BotCamel/apps/destiny"
	t "github.com/StrayCamel247/BotCamel/apps/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

var dbOrm *gorm.DB

func init() {

}

// CamelBot=结构体实例
type CamelBot struct {
	db  *t.CamelOrm
	bot *coolq.CQBot
}

// mod模块实例
type mod interface {
}

// 触发
func Start(bot *coolq.CQBot) {

	// 初始化机器基本功能
	Camel := CamelBot{db: t.Orm, bot: bot}
	// 初始化命运2基础功能
	// 初始化时检查命运2数据库是否存在
	destiny.Start()
	// 命运2群聊天功能启动
	Camel.bot.Client.OnGroupMessage(destiny.GroupMessageEvent)
	Camel.bot.Client.OnGroupMessage(MenuSendEvent)
}

// 综合菜单发送
func MenuSendEvent(c *client.QQClient, m *message.GroupMessage) {
	// c.SolveGroupJoinRequest(e, true, false, "")
	_, com, _ := camel.AnalysisMsg(c, m.Elements)

	switch {
	case t.EqualFolds(com, camel.Commands.Menu.Key()):
		var menu bytes.Buffer

		menu.WriteString("https://github.com/StrayCamel247/BotCamel 领bug修修我吧 (^^ゞ !!!\n")
		menu.WriteString("──────BotCamel基础功能─────\n")
		menu.WriteString(strings.Join(comStrings(camel.Commands), "\n") + "\n")
		menu.WriteString("──────命运2基础功能────────\n")
		menu.WriteString(strings.Join(comStrings(destiny.Commands), "\n") + "\n")
		res := strings.ReplaceAll(menu.String(), "{", "")
		res = strings.ReplaceAll(res, "}", "")
		imgPath, err := t.Text2ImgHandler(res)
		if err != nil {
			log.WithError(err)
		}
		_imgM, err := c.UploadGroupImageByFile(m.GroupCode, imgPath)
		if err != nil {
			log.WithError(err)
		}
		c.SendGroupMessage(m.GroupCode, message.NewSendingMessage().Append(_imgM))
	}

}

// 字符串化指令结构体
func comStrings(a interface{}) (res []string) {

	// var buf bytes.Buffer

	val := reflect.ValueOf(a) //获取reflect.Type类型

	kd := val.Kind() //获取到a对应的类别
	if kd != reflect.Struct {
		fmt.Println("expect struct")
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	// fmt.Printf("该结构体有%d个字段\n", num) //4个

	//遍历结构体的所有字段
	for i := num - 1; i > -1; i-- {
		// buf.WriteString(fmt.Sprintf("├─%v\n", val.Field(i)))
		res = append([]string{fmt.Sprintf("├─ %v", val.Field(i))}, res...)
	}
	return
	// return buf.String()
}

// 收到加群邀请
func GroReciveInviteEvent(c *client.QQClient, e *client.GroupInvitedRequest) {
	c.SolveGroupJoinRequest(e, true, false, "")
}
