package camel

import (
	t "github.com/StrayCamel247/BotCamel/apps/utils"
)

// CommandStruct命令指令结构体
type CommandStruct struct {
	Menu t.Info
}

var Commands CommandStruct

func init() {
	Commands = CommandStruct{
		Menu: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"menu", "菜单", "功能"},
			Remark:  "功能列表"},
	}
}
