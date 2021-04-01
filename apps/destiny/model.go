package destiny

/*
   __AUTHOR__ : stray_camel
  __DESCRIPTION__ : 接口存表-改变表结构或者新增表结构需要增加配置
  __REFERENCES__:
  __DATE__: 2021-03-18
*/
import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// des 介绍信息查询
func (r *Destiny) DesQuery(params interface{}) string {
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
	r.Orm.Fetch_data_sql(_sql, &resStruct, params)
	return resStruct.Des
}

// id itemid查询
func (r *Destiny) IdQuery(params interface{}) (res [][2]string) {
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
	r.Orm.Fetch_data_sql(_sql, &resStruct, params)
	for _, v := range resStruct {
		res = append(res, [2]string{v.ItemId, v.ChtName})
	}
	return res
}

// D2VersionHandler 返回true需要更新数据-false则不需要更新
func (r *Destiny) D2VersionHandler(params interface{}) bool {
	type versionRestruct struct {
		Version string `json:"version"`
	}
	var resStruct versionRestruct
	_sql := `
		select 
			version
		from destiny2_version
		where version = @version
	`
	r.Orm.Fetch_data_sql(_sql, &resStruct, params)
	if resStruct.Version == "" {
		log.Errorf(fmt.Sprintf("Destiny2 数据不是最新-等待更新ing"))
		_insertBase := `
		INSERT INTO destiny2_version ("version") 
		VALUES
		(@version)
		`
		r.Orm.Execute(_insertBase, params)
		return true
	}
	log.Infof(fmt.Sprintf("Destiny2 数据已是最新 version: %s", resStruct.Version))
	return false
}

// InsertMenifestHandler 清空表后更新
func (r *Destiny) InsertMenifestHandler(dataArray [][]interface{}) {
	if len(dataArray) == 0 {
		return
	}
	_insertBase := `
	
	INSERT INTO destiny2_menifest_base 
	("created_at", "updated_at", "deleted_at", 
	"itemid", "description", "name", "icon", "tag", "seasonid", "language") 
	VALUES
	`
	_insertSub := `
		(CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL, 
			'%s', '%s', '%s', '%s', '%s', '%s', '%s')
	`
	_batch := len(dataArray) / 800
	if _batch > 1 {
		for i := 0; i < _batch; i++ {
			r.Orm.Execute_batch(_insertBase, _insertSub, dataArray[i*800:(i+1)*800])
		}
	} else if _batch == 1 {
		r.Orm.Execute_batch(_insertBase, _insertSub, dataArray)
	}

	log.Infof(fmt.Sprintf("Destiny2 数据正在更新 Table: %s", "destiny2_menifest_base"))
}

// DBCheckHandler-检查表是否存在-若不存在-组装待初始化的sql
func (r *Destiny) DBCheckHandler(tableName, tableSql string, Sqls *[]string) {
	// 定义查询传参和返回结果对象

	type countResult struct {
		Count int64 `json:"count"`
	}
	_sql := `
			SELECT 
				count(1) count
			FROM sqlite_master 
			WHERE 
				type ='table' 
				and name =@table_name
			ORDER BY name
		`
	var resStruct countResult
	params := map[string]interface{}{"table_name": tableName}
	r.Orm.Fetch_data_sql(_sql, &resStruct, params)

	if resStruct.Count == 0 {
		*Sqls = append(*Sqls, tableSql)
		_msg := fmt.Sprintf("NO DB Named: %s", tableName)
		log.Errorf(_msg)
	}
	log.Infof(fmt.Sprintf("DB existed: %s", tableName))
}
