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

var Commands CommandStruct

func init() {
	Commands = CommandStruct{
		D2week: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"week", "周报"},
			Remark:  "周报信息查询"},
		D2day: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"day", "日报"},
			Remark:  "日报信息查看"},
		D2xiu: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"xiu", "nine", "老九"},
			Remark:  "老九信息查询"},
		D2trial: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"trial", "试炼", "train"},
			Remark:  "试炼最新动态"},
		D2dust: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"dust", "光尘"},
			Remark:  "赛季光尘商店"},
		D2random: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"random", "骰子"},
			Remark:  "骰子功能"},
		D2perk: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"perk", "词条"},
			Remark:  "查询词条/模组信息"},
		D2item: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"item", "物品"},
			Remark:  "查询物品信息-提供light-gg信息"},
		D2npc: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"npc"},
			Remark:  "查询npc信息"},
		D2skill: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"skill"},
			Remark:  "查询技能等信息"},
		D2pve: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"pve"},
			Remark:  "查询pve信息"},
		D2pvp: t.Info{
			Command: t.FetchCommandNum(),
			Keys:    []string{"pvp"},
			Remark:  "查询pvp信息"},
	}
}
