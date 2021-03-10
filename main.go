package main

import (
    "github.com/StrayCamel247/BotCamel/apps"
    "github.com/StrayCamel247/BotCamel/config"
    "github.com/StrayCamel247/BotCamel/apps/utils"
)

func init() {
    utils.WriteLogToFS()
    config.Init()
}

func main() {
    // tools_device.TestGenDevice()
    apps.QqBotInit()
}
