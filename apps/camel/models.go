package camel

/*
   __AUTHOR__ : stray_camel
  __DESCRIPTION__ : 接口存表-改变表结构或者新增表结构需要增加配置
  __REFERENCES__:
  __DATE__: 2021-03-18
*/
import (
	// "encoding/json"
	// "fmt"
	"github.com/StrayCamel247/BotCamel/apps/utils"
	// log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	// "io/ioutil"
	// "net/http"
	// "regexp"
	// "strconv"
	// "sync"
	// "reflect"
	// "time"
)

func init() {

}

// des 介绍信息查询
func DesQuery(orm *gorm.DB, params interface{}) string {
	type ReStruct struct {
		Des string `gorm:"column:description"`
	}
	var resStruct ReStruct
	_sql := `
		with tmpBase as (
			select
				description||'\n' description
			from
				destiny2_menifest_base
			where
				name = @name
				and description is not null
				and description != ' '
				and description != ''
			group by description
		)
		select GROUP_CONCAT(description) description from tmpBase
	`
	utils.Fetch_data_sql(orm, _sql, &resStruct, params)
	return resStruct.Des
}

// id itemid查询
func IdQuery(orm *gorm.DB, params interface{}) (res [][2]string) {
	type ReStruct struct {
		ItemId  string `gorm:"column:itemid"`
		ChtName string `gorm:"column:name"`
	}
	var resStruct []ReStruct
	_sql := `
		with zhChtData as (
			select itemid, name
			from
				destiny2_menifest_base
			where
				language = 'zh-cht'
		)
		, base as (
			select
				itemid
			from
				destiny2_menifest_base
			where
				name = @name
				and itemid is not null
				and itemid != ' '
				and itemid != ''
				and tag = 'DestinyInventoryItemLiteDefinition'
			group by itemid
		)
		select 
			base.itemid
			,t.name
		from base
		left join zhChtData t
			on base.itemid = t.itemid
		group by 
			base.itemid
			,t.name
		
	`
	utils.Fetch_data_sql(orm, _sql, &resStruct, params)
	for _, v := range resStruct {
		res = append(res, [2]string{v.ItemId, v.ChtName})
	}
	return res
}
