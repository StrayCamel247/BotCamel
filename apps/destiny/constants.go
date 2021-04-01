package destiny

import (
	"fmt"
)

// 玩家pvp-pve生涯记录数据查询-接口
var profileAPI = func(fireId string) string {
	return fmt.Sprintf("https://www.bungie.net/Platform/Destiny2/3/Account/%s/Stats/", fireId)
}

//
var baseprofileAPI = func(fireId string) string {
	return fmt.Sprintf("https://www.bungie.net/Platform/Destiny2/3/Profile/%s/?components=100", fireId)
}

// var config *global.JSONConfig
