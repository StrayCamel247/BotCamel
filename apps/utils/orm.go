/*
   __author__ : stray_camel
  __description__ : 继承使用gorm构造orm对象
  __REFERENCES__:
  __date__: 2021-03-16
*/
package utils

import (
	// "database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	// "gorm.io/gorm/logger"
)

var Orm *CamelOrm

func init() {
	Orm = GetOrm()
}

// 触发
func GetOrm() *CamelOrm {
	log.Infof("载入sqlite数据库...")
	// 启用sqlite数据库
	orm, _ := gorm.Open(sqlite.Open("./data/sqlite3.db"), &gorm.Config{
		// PrepareStmt: true,
		// 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。但这会降低性能，你可以在初始化时禁用这种方式
		SkipDefaultTransaction: true,
	})
	return &CamelOrm{orm}
}

type CamelOrm struct {
	Orm *gorm.DB
}

// Execute_batch 批量处理数据
func (r *CamelOrm) Execute_batch(baseSql, sql string, orderParamsList [][]interface{}) (lines int) {
	_paramsLen := len(orderParamsList)
	if _paramsLen == 0 {
		return _paramsLen
	}
	// format 参数列表
	var _sqlArrys = make([]string, _paramsLen)
	for i := 0; i < _paramsLen; i++ {
		_sqlArrys[i] += fmt.Sprintf(string(sql), orderParamsList[i]...)
	}
	if baseSql == "" {
		res := r.Orm.Debug().Exec(strings.Join(_sqlArrys, ";"))
		if res.Error != nil {
			log.WithError(res.Error)
		}
		log.Infof(fmt.Sprintf("execute successed lines: %d", res.RowsAffected))
		return int(res.RowsAffected)
	} else {
		// 如果数据量大于5 不打印日志
		var res *gorm.DB
		if len(_sqlArrys) > 5 {
			res = r.Orm.Exec(baseSql + strings.Join(_sqlArrys, ","))
		} else {
			res = r.Orm.Debug().Exec(baseSql + strings.Join(_sqlArrys, ","))
		}
		if res.Error != nil {
			log.WithError(res.Error)
		}
		log.Infof(fmt.Sprintf("execute successed lines: %d", res.RowsAffected))
		return int(res.RowsAffected)
	}

}

// Execute 处理数据
func (r *CamelOrm) Execute(sql string, params interface{}) int64 {
	var res *gorm.DB
	if params == nil {
		res = r.Orm.Debug().Exec(string(sql))
	} else {
		res = r.Orm.Debug().Exec(string(sql), params)
	}

	if res.Error != nil {
		log.WithError(res.Error)
	}

	log.Infof(fmt.Sprintf("execute successed lines: %d", res.RowsAffected))
	return res.RowsAffected
}

// 获取数据

func (r *CamelOrm) Fetch_data_sql(sql string, resStruct interface{}, params interface{}) (res *gorm.DB) {
	res = r.Orm.Debug().Raw(string(sql), params).Scan(resStruct)
	if res.Error != nil {
		log.WithError(res.Error)
	}
	log.Infof(fmt.Sprintf("fetch successed lines: %d", res.RowsAffected))
	return res
}
