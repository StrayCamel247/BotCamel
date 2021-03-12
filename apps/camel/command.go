package camel

type CommandEleStruct struct {
	Keys   []string
	Remark string
}

// CommandStruct命令指令结构体
type CommandsStruct struct {
	Menu         CommandEleStruct
	Asskisser    CommandEleStruct
	Motherfucker CommandEleStruct
	D2week       CommandEleStruct
	D2xiu        CommandEleStruct
	D2trial      CommandEleStruct
	Developers   CommandEleStruct
}

func init() {

}

// CommandFilter: 指令模糊判断
func CommandFilter() CommandsStruct {
	return CommandsStruct{
		Menu: CommandEleStruct{
			Keys:   []string{"menu", "菜单"},
			Remark: "查看所有指令"},
		Asskisser: CommandEleStruct{
			Keys:   []string{"0x00", "asskisser"},
			Remark: "0x00 Asskisser 夸一下"},
		Motherfucker: CommandEleStruct{
			Keys:   []string{"0x01", "motherfucker"},
			Remark: "0x01 Motherfucker 碧池模式"},
		D2week: CommandEleStruct{
			Keys:   []string{"0x02", "week", "周报"},
			Remark: "0x02 week 周报信息查询"},
		D2xiu: CommandEleStruct{
			Keys:   []string{"0x03", "xiu", "nine", "老九"},
			Remark: "0x03 xiu 老九信息查询"},
		D2trial: CommandEleStruct{
			Keys:   []string{"0x04", "trail", "trial", "试炼", "train"},
			Remark: "0x04 trial 试炼最新动态"},
		Developers: CommandEleStruct{
			Keys:   []string{"0xFF", "developers", "developer", "开发人员"},
			Remark: "🙊 娃哈哈店长-StrayCamel247\n👋 期待你的加入"},
	}
}
