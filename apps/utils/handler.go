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
	"gorm.io/gorm"
	"strings"
	// "gorm.io/gorm/logger"
)

func init() {

}

// Execute_batch 批量处理数据
func Execute_batch(orm *gorm.DB, sql string, orderParamsList [][]interface{}) (lines int) {
	_paramsLen := len(orderParamsList)
	if _paramsLen == 0 {
		return _paramsLen
	}
	// format 参数列表
	var _sqlArrys = make([]string, _paramsLen)
	for i := 0; i < _paramsLen; i++ {
		_sqlArrys[i] += fmt.Sprintf(string(sql), orderParamsList[i]...)
	}

	res := orm.Debug().Exec(strings.Join(_sqlArrys, ";"))
	if res.Error != nil {
		panic(res.Error)
	}
	log.Infof(fmt.Sprintf("execute successed lines: %d", res.RowsAffected))
	return int(res.RowsAffected)
}

// Execute 处理数据
func Execute(orm *gorm.DB, sql string, params interface{}) int64 {
	res := orm.Debug().Exec(string(sql), params)
	if res.Error != nil {
		panic(res.Error)
	}

	log.Infof(fmt.Sprintf("execute successed lines: %d", res.RowsAffected))
	return res.RowsAffected
}
func Fetch_data_sql(orm *gorm.DB, sql string, resStruct interface{}, params interface{}) (res *gorm.DB) {
	res = orm.Debug().Raw(string(sql), params).Scan(resStruct)
	if res.Error != nil {
		panic(res.Error)
	}
	log.Infof(fmt.Sprintf("fetch successed lines: %d", res.RowsAffected))
	return res
}
