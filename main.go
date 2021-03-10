package main

import (
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/StrayCamel247/BotCamel/apps"
)

func init() {
	utils.WriteLogToFS()
	config.Init()
}

func main() {
	apps.QqBotInit()
}
