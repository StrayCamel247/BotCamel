package destiny

/*
   __AUTHOR__ : stray_camel
  __DESCRIPTION__ : 接口存表-改变表结构或者新增表结构需要增加配置
  __REFERENCES__:
  __DATE__: 2021-03-18
*/
import (
	// "encoding/json"
	"fmt"
	"github.com/StrayCamel247/BotCamel/apps/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	// "io/ioutil"
	// "net/http"
	// "regexp"
	// "strconv"
	// "reflect"
	// "time"
	"sync"
)

// D2VersionHandler 返回true需要更新数据-false则不需要更新
func D2VersionHandler(orm *gorm.DB, params interface{}) bool {
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
	utils.Fetch_data_sql(orm, _sql, &resStruct, params)
	if resStruct.Version == "" {
		log.Errorf(fmt.Sprintf("Destiny2 数据不是最新-等待更新ing"))
		_insertBase := `
		INSERT INTO destiny2_version ("version") 
		VALUES
		(@version)
		`
		utils.Execute(orm, _insertBase, params)
		return true
	}
	log.Infof(fmt.Sprintf("Destiny2 数据已是最新 version: %s", resStruct.Version))
	return false
}

// InsertMenifestHandler 清空表后更新
func InsertMenifestHandler(orm *gorm.DB, dataArray [][]interface{}) {
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
	if _batch >= 1 {
		var wg = sync.WaitGroup{}
		wg.Add(_batch)
		_wgHandler := func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			utils.Execute_batch(orm, _insertBase, _insertSub, dataArray[i*800:(i+1)*800])
		}
		for i := 0; i < _batch; i++ {
			// utils.Execute_batch(orm, _insertBase, _insertSub, dataArray[i*800:(i+1)*800])
			_wgHandler(i, &wg)
		}
		wg.Wait()
	} else {
		utils.Execute_batch(orm, _insertBase, _insertSub, dataArray)
	}

	log.Infof(fmt.Sprintf("Destiny2 数据正在更新 Table: %s", "destiny2_menifest_base"))
}

// DBCheckHandler-检查表是否存在-若不存在-组装待初始化的sql
func DBCheckHandler(orm *gorm.DB, tableName, tableSql string, Sqls *[]string) {
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
	utils.Fetch_data_sql(orm, _sql, &resStruct, params)

	if resStruct.Count == 0 {
		*Sqls = append(*Sqls, tableSql)
		_msg := fmt.Sprintf("NO DB Named: %s", tableName)
		log.Errorf(_msg)
	}
	log.Infof(fmt.Sprintf("DB existed: %s", tableName))
}
