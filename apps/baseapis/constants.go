package baseapis

import (
	"fmt"
	// "github.com/Logiase/MiraiGo-Template/utils"
	"github.com/StrayCamel247/BotCamel/global"
)

// key表名-value建表语句
var D2Table map[string]string

func init() {
	D2Table = make(map[string]string)
	// 构造D2需要的表和建表语句
	// D2 version 版本控制表-只保留最新版本的数据
	D2Table["destiny2_version"] = `
		-- ----------------------------
		-- Table structure for destiny2_version
		-- 命运2  
		-- ----------------------------
		DROP TABLE IF EXISTS "destiny2_version";
		CREATE TABLE "destiny2_version" (
		"version" text
		);
	`
	// 基础表
	D2Table["destiny2_menifest_base"] = `
		-- ----------------------------
		-- Table structure for destiny2_menifest_base
		-- 命运2  数据主表-itemid-name-des主要信息
		-- ----------------------------
		DROP TABLE IF EXISTS "destiny2_menifest_base";
		CREATE TABLE "destiny2_menifest_base" (
		"created_at" datetime,
		"updated_at" datetime,
		"deleted_at" datetime,
		"itemid" text,    -- 主键itemid
		"description" text,   -- 描述
		"name" text,  -- 中文名称
		"language" text,  -- 语言类型
		"icon" text,  -- destiny2官方图标
		"tag" text,   -- tag标签
		"seasonid" text  -- 赛季itemid
		);

		-- ----------------------------
		-- Indexes structure for table destiny2_menifest_base
		-- ----------------------------
		DROP INDEX IF EXISTS "destiny2_menifest_base_name";
		CREATE INDEX "destiny2_menifest_base"
		ON "destiny2_menifest_base_name" (
		"name" DESC
		);
	`
	// item perk 关系表
	D2Table["destiny2_item_perk"] = `
		-- ----------------------------
		-- Table structure for destiny2_item_perk
		-- 命运2
		-- ----------------------------
		DROP TABLE IF EXISTS "destiny2_item_perk";
		CREATE TABLE "destiny2_item_perk" (
		"created_at" datetime,
		"updated_at" datetime,
		"deleted_at" datetime,
		"itemid" text,    -- 主键itemid
		"perkitemid" text,    -- 主键itemid

		CONSTRAINT "destiny2_item_perk_uni" UNIQUE ("itemid", "perkitemid")
		);
		DROP INDEX IF EXISTS "destiny2_item_itemid_perk";
		CREATE INDEX "destiny2_item_itemid_perk"
		ON "destiny2_item_perk" (
		"itemid" DESC
		);
	`
}

// var logger = utils.GetModuleLogger("NmslApi_handler")

// Bungie-api
var BungiePlatformRoot = "https://www.bungie.net/Platform"
var BungieBase = "https://www.bungie.net/"

// https://www.bungie.net/Platform/Destiny2/Manifest/ 所有点Definition对应表 zh-chs
var BunigieManifestUrl = "https://www.bungie.net/Platform/Destiny2/Manifest/"

// 沙雕app-骂人
var nmslAPI = "https://nmsl.shadiao.app/api.php?from=%s"

// 沙雕app-夸人
var lickAPI = "https://chp.shadiao.app/api.php?from=%s"

// 玩家pvp-pve生涯记录数据查询-接口
var profileAPI = func(fireId string) string {
	return fmt.Sprintf("https://www.bungie.net/Platform/Destiny2/3/Account/%s/Stats/", fireId)
}

//
var baseprofileAPI = func(fireId string) string {
	return fmt.Sprintf("https://www.bungie.net/Platform/Destiny2/3/Profile/%s/?components=100", fireId)
}
var config *global.JSONConfig
