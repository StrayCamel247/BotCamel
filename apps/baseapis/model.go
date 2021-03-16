package baseapis

import (
	// "encoding/json"
	// "fmt"
	// "github.com/bitly/go-simplejson"
	// "github.com/Mrs4s/MiraiGo/client"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	// "io/ioutil"
	// "net/http"
	// "regexp"
	// "strconv"
	// "time"
)

func init() {

}

//
type InfoDisplayDB struct {
	gorm.Model
	ItemId      string `gorm:"primaryKey;SIZE:0;Name:item_id"`
	Description string `gorm:"SIZE:0;Name:des"`
	// Name        string `gorm:"SIZE:0;Name:name;index:,sort:desc,collate:utf8,type:btree"`
	Name string `gorm:"SIZE:0;Name:name;index:,sort:desc"`
	Icon string `gorm:"SIZE:0;Name:des"`
	Tag  string `gorm:"SIZE:0;Name:tag"`
}
type ItemIdDB struct {
	ItemId      string
	Description string
	Name        string
	Tag         string
}

func InfoDisplayDBCheck(orm *gorm.DB) (bool, error) {
	// 检查是否存在表
	if !orm.Migrator().HasTable(&InfoDisplayDB{}) {
		if err := orm.Migrator().CreateTable(&InfoDisplayDB{}); err != nil {
			log.Warn(err)
			// panic(err)
			return false, err
		}
		return false, nil
	}

	return true, nil
}
