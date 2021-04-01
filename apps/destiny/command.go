package destiny

import (
	t "github.com/StrayCamel247/BotCamel/apps/utils"
)

// CommandStruct命令指令结构体
type CommandStruct struct {
	D2week   t.Info
	D2day    t.Info
	D2xiu    t.Info
	D2trial  t.Info
	D2dust   t.Info
	D2random t.Info
	D2perk   t.Info
	D2item   t.Info
	D2npc    t.Info
	D2skill  t.Info
	D2pve    t.Info
	D2pvp    t.Info
}

var command CommandStruct

func init() {
	command = CommandStruct{
		D2week: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"week", "周报"},
			Remark:  "0x02 week 周报信息查询"},
		D2day: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"day", "日报"},
			Remark:  "0x03 日报信息查看"},
		D2xiu: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"xiu", "nine", "老九"},
			Remark:  "0x04 xiu 老九信息查询"},
		D2trial: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"trail", "trial", "试炼", "train"},
			Remark:  "0x05 trial 试炼最新动态"},
		D2dust: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"dust", "dustdetail", "光尘", "光尘商店"},
			Remark:  "0x06 赛季光尘商店"},
		D2random: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"random", "random10", "骰子", "掷色子"},
			Remark:  "0x07 骰子功能"},
		D2perk: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"perk", "词条"},
			Remark:  "0x08 perk 查询词条/模组信息"},
		D2item: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"item", "物品"},
			Remark:  "0x09 查询物品信息-提供light.gg信息"},
		D2npc: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"npc"},
			Remark:  "0x10 查询npc信息"},
		D2skill: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"skill"},
			Remark:  "0x0a 查询技能等信息"},
		D2pve: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"pve"},
			Remark:  "0x0b 查询pve信息"},
		D2pvp: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"pvp"},
			Remark:  "0x0c 查询pvp信息"},
	}
}
