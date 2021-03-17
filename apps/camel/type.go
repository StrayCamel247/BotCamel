package camel

type dayRes struct {
	IMG_URL      string `json:"img_url"`
	IMG_HASH_MD5 string `json:"img_hash_md5"`
}
type Info struct {
	Keys   []string
	Remark string
}

// pvp查询返回
type PvPInfo struct {
	// 胜率
	KDA map[string]string
	// 游玩时长(小时)
	HoursPlayed map[string]string
}

// var PvPInfoResult = PvPInfo

// CommandStruct命令指令结构体
type CommandsStruct struct {
	Menu         Info
	Asskisser    Info
	Motherfucker Info
	D2week       Info
	D2day        Info
	D2xiu        Info
	D2trial      Info
	D2dust       Info
	D2random     Info
	D2perk       Info
	D2item       Info
	D2npc        Info
	D2skill      Info
	D2pve        Info
	D2pvp        Info
	Developers   Info
	DesChecker   Info
}

func init() {

}

// CommandFilter: 指令模糊判断
func CommandFilter() CommandsStruct {
	return CommandsStruct{
		Menu: Info{
			Keys:   []string{"menu", "菜单"},
			Remark: "查看所有指令"},
		Asskisser: Info{
			Keys:   []string{"0x00", "asskisser", "乖乖", "宝贝", "爱你"},
			Remark: "0x00 Asskisser 夸一下"},

		Motherfucker: Info{
			Keys:   []string{"0x01", "motherfucker", "傻逼", "cnm", "草泥马", "操你妈"},
			Remark: "0x01 Motherfucker 碧池一下"},

		D2week: Info{
			Keys:   []string{"0x02", "week", "周报"},
			Remark: "0x02 week 周报信息查询"},

		D2day: Info{
			Keys:   []string{"0x03", "day", "日报"},
			Remark: "0x03 日报信息查看"},

		D2xiu: Info{
			Keys:   []string{"0x04", "xiu", "nine", "老九"},
			Remark: "0x04 xiu 老九信息查询"},

		D2trial: Info{
			Keys:   []string{"0x05", "trail", "trial", "试炼", "train"},
			Remark: "0x05 trial 试炼最新动态"},

		D2dust: Info{
			Keys:   []string{"0x06", "dust", "dustdetail", "光尘", "光尘商店"},
			Remark: "0x06 赛季光尘商店"},

		D2random: Info{
			Keys:   []string{"0x07", "random", "random10", "骰子", "掷色子"},
			Remark: "0x07 骰子功能"},

		D2perk: Info{
			Keys:   []string{"0x08", "perk", "词条"},
			Remark: "0x08 perk 查询词条/模组信息"},

		D2item: Info{
			Keys:   []string{"0x09", "item", "物品"},
			Remark: "0x09 查询物品信息-提供light.gg信息"},

		D2npc: Info{
			Keys:   []string{"0x10", "npc"},
			Remark: "0x10 查询npc信息"},

		D2skill: Info{
			Keys:   []string{"0x1a", "skill"},
			Remark: "0x1a 查询技能等信息"},

		D2pve: Info{
			Keys:   []string{"0x1b", "pve"},
			Remark: "0x1b 查询pve信息"},

		D2pvp: Info{
			Keys:   []string{"0x1c", "pvp"},
			Remark: "0x1c 查询pvp信息"},

		Developers: Info{
			Keys:   []string{"0xFF", "developers", "developer", "开发人员"},
			Remark: "🙊 娃哈哈店长-StrayCamel247\n👋 期待你的加入"},

		DesChecker: Info{
			Keys:   []string{"", " ", "\n", "\n\n"},
			Remark: ""},
	}
}
