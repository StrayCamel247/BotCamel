package main

import (
	"os"
	"os/signal"

	"fmt"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	_ "github.com/Logiase/MiraiGo-Template/modules/logging"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/StrayCamel247/BotCamel/apps"
)

func init() {
	utils.WriteLogToFS()
	config.Init()
}

// qqStart : qq机器人启动器
func qqStart() {
	// 快速初始化
	bot.Init()

	// 初始化 Modules
	bot.StartService()

	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)

	// 登录
	bot.Login()

	// 刷新好友列表，群列表
	bot.RefreshList()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	bot.Stop()
}
func main() {
	apps.QqBotInit()
	fmt.Print("This is main\n")
}
