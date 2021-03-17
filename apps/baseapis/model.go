package baseapis

import (
	// "encoding/json"
	"fmt"
	// "github.com/bitly/go-simplejson"
	// "github.com/Mrs4s/MiraiGo/client"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	// "io/ioutil"
	// "net/http"
	// "regexp"
	// "strconv"
	"reflect"
	"time"
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

// InfoDisplayDBCheck 检查命运2 menifest表是否存在-若不存在则抽取
func InfoDisplayDBCheck(orm *gorm.DB) error {
	// 检查是否存在表
	if !orm.Migrator().HasTable(&InfoDisplayDB{}) {
		if err := orm.Migrator().CreateTable(&InfoDisplayDB{}); err != nil {
			log.Warn(err)
			// panic(err)
			return err
		}
		// 若数据库表不存在，并发查询数据并写入
		file, _ := ManifestFetchJson()

		typ := reflect.TypeOf(file)
		val := reflect.ValueOf(file) //获取reflect.Type类型

		kd := val.Kind() //获取到a对应的类别
		if kd != reflect.Struct {
			log.Info("expect struct")
			return nil
		}
		//获取到该结构体有几个字段
		num := val.NumField()

		//遍历结构体的所有字段
		start := time.Now()
		ch := make(chan bool)
		for i := 0; i < num; i++ {
			// goroutine的正确用法
			// 那怎么用goroutine呢？有没有像Python多进程/线程的那种等待子进/线程执行完的join方法呢？当然是有的，可以让Go 协程之间信道（channel）进行通信：从一端发送数据，另一端接收数据，信道需要发送和接收配对，否则会被阻塞：
			// log.Info("Field %d:值=%v\n", i, val.Field(i))
			tagVal := typ.Field(i).Tag.Get("json")
			//如果该字段有tag标签就显示，否则就不显示
			// if tagVal != "" {
			// 	log.Info("Field %d:tag=%v\n", i, tagVal)
			// }
			// 并发
			// go ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), orm, ch)
			// 串行
			print(tagVal)
			ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), orm, ch)
			// if tagVal == "DestinyInventoryItemLiteDefinition" {
			// 	ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), orm, ch)
			// }

		}
		elapsed := time.Since(start)
		log.Info(fmt.Sprintf("Took %s", elapsed))

	}

	return nil
}
