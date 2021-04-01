/*
	命运2数据库检查中心
*/
package destiny

import (
	"encoding/json"
	"fmt"
	// con "github.com/StrayCamel247/BotCamel/apps/config"
	// "github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	// "regexp"
	// "strconv"

	"bytes"
	"reflect"
	"strings"
	// "sync"
	"time"
)

// key表名-value建表语句
var D2Table map[string]string

// 常量声明
const DayGenUrl string = "http://www.tianque.top/d2api/today/"

// Bungie-api
const BungiePlatformRoot string = "https://www.bungie.net/Platform"
const BungieBase string = "https://www.bungie.net/"

// https://www.bungie.net/Platform/Destiny2/Manifest/ 所有点Definition对应表 zh-chs
const BunigieManifestUrl string = "https://www.bungie.net/Platform/Destiny2/Manifest/"

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

// 日报信息更新
func (r *Destiny) RefreshDayHandler(flag, url string) {
	ch := make(chan bool)
	// 每秒轮询日报信息-是否更新-若更新跳出
	go func() {
		for {
			_, updated := D2DownloadHandler(flag, url)
			if updated {
				log.Infof("定时器-日报数据已更新！")
				ch <- updated
			}
		}
	}()
	_ = <-ch
}

// ManifestFetchResponse 获取menifest接口返回
func (r *Destiny) ManifestFetchResponse() (Response ManifestWorldComponentContent, err error) {
	spaceClient := http.Client{
		Timeout: time.Second * 999, // Maximum of 100 secs
	}

	req, err := http.NewRequest(http.MethodGet, BunigieManifestUrl, nil)
	if err != nil {
		panic(err)

	}

	req.Header.Add("X-API-Key", "aff47ade61f643a19915148cfcfc6d7d")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		panic(getErr)

	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		panic(readErr)

	}
	ResJson := ManifestResult{}
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		panic(jsonErr)

	}
	return ResJson.Response, nil
}

// InfoMenifestBaseDBCheck 检查命运2 menifest表是否存在-若不存在则抽取
func (r *Destiny) InfoMenifestBaseDBCheck() {
	// 检查是否存在表
	_Num := len(D2Table)
	log.Infof(fmt.Sprintf("正在检查%d张表...", _Num))
	var _InItSqls []string
	for tableName, tableSql := range D2Table {
		r.DBCheckHandler(tableName, tableSql, &_InItSqls)
	}
	if len(_InItSqls) > 0 {
		// 直接调用执行-需等待率先你表后再进行后序的插入操作-阻塞
		r.Orm.Execute(strings.Join(_InItSqls, ";"), nil)
	}
	// 检查表里数据是否是最新-若不是或者数据为空-则重新抽数到数据库
	manifestRes, _ := r.ManifestFetchResponse()
	params := map[string]interface{}{"version": manifestRes.NewVersion}
	needUpdate := r.D2VersionHandler(params)

	if needUpdate || len(_InItSqls) == _Num {
		_handler := func(_Data interface{}, LangType string) {
			typ := reflect.TypeOf(_Data)
			val := reflect.ValueOf(_Data) //获取reflect.Type类型

			kd := val.Kind() //获取到a对应的类别
			if kd != reflect.Struct {
				log.Info("expect struct")
			}
			//获取到该结构体有几个字段
			num := val.NumField()

			//遍历结构体的所有字段
			for i := 0; i < num; i++ {
				tagVal := typ.Field(i).Tag.Get("json")
				r.ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), LangType)
			}
		}
		// 分批次写入-无需锁表
		// 写入中文数据
		AllData := manifestRes.JsonWorldComponentContentPaths
		typ := reflect.TypeOf(AllData)
		val := reflect.ValueOf(AllData) //获取reflect.Type类型
		kd := val.Kind()                //获取到a对应的类别
		if kd != reflect.Struct {
			log.Info("expect struct")
		}
		//获取到该结构体有几个字段
		num := val.NumField()
		//遍历结构体的所有字段
		for i := 0; i < num; i++ {
			LangType := typ.Field(i).Tag.Get("json")
			// fmt.Printf("%+v", val.Field(i))
			_handler(val.Field(i).Interface(), LangType)
		}
		r.D2VersionHandler(params)
	}
}

// ManifestFetchInfo 查询解析url数据并写入 InfoMenifestBaseDB 表
func (r *Destiny) ManifestFetchInfo(josnFile, tag string, LangType string) {
	u, err := url.Parse(BungieBase)
	u.Path = path.Join(u.Path, josnFile)
	url := u.String()
	spaceClient := http.Client{
		Timeout: time.Second * 999, // Maximum of 10 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Warn(err)
	}

	req.Header.Add("X-API-Key", "aff47ade61f643a19915148cfcfc6d7d")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Warn(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Warn(readErr)
	}
	var ResJson map[string]InfoDisplay
	// error "invalid character 'ï' looking for beginning of value” from json.Unmarsh https://blog.csdn.net/qq_30505673/article/details/97646315
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		log.Warn(jsonErr)
	}
	var paramList [][]interface{}
	// 对名字进行去重
	_nameTmp := make([]string, 0)
	_handler := func(name string) bool {
		for _, v := range _nameTmp {
			if v == name {
				return false
			}
		}
		_nameTmp = append(_nameTmp, name)
		return true
	}

	for itemid := range ResJson {

		// defer wg.Done()
		// 格式化数据-为字符串
		// "itemid", "description", "name", "icon", "tag", "seasonid" 遵循默认排序
		var _SeasonId string
		_Description := ResJson[itemid].Properties.Description
		_Icon := ResJson[itemid].Properties.Icon
		_Name := ResJson[itemid].Properties.Name
		if ResJson[itemid].SeasonHash > 0 {
			_SeasonId = fmt.Sprintf("%d", ResJson[itemid].SeasonHash)
		}

		if (_Description != "" || _Icon != "") && _handler(_Name) {
			// 单一字符串进行替换
			// var comDict map[string]string
			comDict := map[string]string{"'": "\"", "-": "="}
			for k, v := range comDict {
				itemid = strings.ReplaceAll(itemid, k, v)
				_Description = strings.ReplaceAll(_Description, k, v)
				_Name = strings.ReplaceAll(_Name, k, v)
				_Icon = strings.ReplaceAll(_Icon, k, v)
				tag = strings.ReplaceAll(tag, k, v)
				_SeasonId = strings.ReplaceAll(_SeasonId, k, v)
			}
			_params := []interface{}{itemid, _Description, _Name, _Icon, tag, _SeasonId, LangType}
			paramList = append(paramList, _params)

		}

	}
	// 写入数据库
	r.InsertMenifestHandler(paramList)
	log.Infof(tag + " down!")
}
