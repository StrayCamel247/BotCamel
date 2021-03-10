package qqbot

import (
    "os"
    "os/signal"

    "github.com/StrayCamel247/BotCamel/apps/bot"

    "github.com/StrayCamel247/BotCamel/config"
    _ "github.com/StrayCamel247/BotCamel/apps/modules/logging"
    "github.com/StrayCamel247/BotCamel/apps/utils"
)

func init() {
    utils.WriteLogToFS()
    config.Init()
    print("BotCamel inited.")
}

// QqStart : qq机器人启动器
func QqStart() {
    // 新键设备文件
    bot.GenRandomDevice()
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
    // Go 中如何使用通道来处理信号https://gobyexample-cn.github.io/signals
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, os.Interrupt, os.Kill)

    <-ch
    bot.Stop()
    print("BotCamel shutdown...")
}
