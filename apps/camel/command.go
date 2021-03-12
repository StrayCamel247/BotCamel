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
	D2day        CommandEleStruct
	D2xiu        CommandEleStruct
	D2trial      CommandEleStruct
	D2dust       CommandEleStruct
	D2random     CommandEleStruct
	D2perk       CommandEleStruct
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
			Keys:   []string{"0x00", "asskisser", "乖乖", "宝贝", "爱你"},
			Remark: "0x00 Asskisser 夸一下"},
		Motherfucker: CommandEleStruct{
			Keys:   []string{"0x01", "motherfucker", "傻逼", "cnm", "草泥马", "操你妈"},
			Remark: "0x01 Motherfucker 碧池模式"},
		D2week: CommandEleStruct{
			Keys:   []string{"0x02", "week", "周报"},
			Remark: "0x02 week 周报信息查询"},
		D2day: CommandEleStruct{
			Keys:   []string{"0x03", "day", "日报"},
			Remark: "0x03 日报信息查看"},
		D2xiu: CommandEleStruct{
			Keys:   []string{"0x04", "xiu", "nine", "老九"},
			Remark: "0x04 xiu 老九信息查询"},
		D2trial: CommandEleStruct{
			Keys:   []string{"0x05", "trail", "trial", "试炼", "train"},
			Remark: "0x05 trial 试炼最新动态"},
		D2dust: CommandEleStruct{
			Keys:   []string{"0x06", "dust", "dustdetail", "光尘", "光尘商店"},
			Remark: "0x06 赛季光尘商店"},
		D2random: CommandEleStruct{
			Keys:   []string{"0x07", "random", "random10", "骰子", "掷色子"},
			Remark: "0x07 筛子功能"},
		D2perk: CommandEleStruct{
			Keys:   []string{"0x08", "perk", "词条"},
			Remark: "0x08 perk查看"},
		Developers: CommandEleStruct{
			Keys:   []string{"0xFF", "developers", "developer", "开发人员"},
			Remark: "🙊 娃哈哈店长-StrayCamel247\n👋 期待你的加入"},
	}
}
