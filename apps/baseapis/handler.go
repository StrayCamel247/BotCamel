package baseapis

/*
   __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-12
*/
/*
	sync.WaitGroup
	WaitGroup 用于等待一组 goroutine 结束,主 goroutine 调用 Add() 设置要等待的 goroutine 的数目。 每个 goroutine 结束时调用 Done()。同时主 goroutine 调用 Wait() ，阻塞主 goroutine 知道 所有的 goroutine 结束。 第一次使用 WatiGroup 实例后, 该 WaitGroup 一定不能被拷贝。更多的信息可以从不能被拷贝的结构 中获得。

	WaitGroup 是结构体，不是引用类型，所以传递给 goroutine 时不能直接传值，而要传递 WaitGroup 实例的指针.
*/
import (
	"encoding/json"
	"fmt"
	con "github.com/StrayCamel247/BotCamel/apps/config"
	"github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"sync"

	"github.com/StrayCamel247/BotCamel/apps/utils"
	"reflect"
	"strings"
	// "strings"
	"time"
)

func init() {
	config = con.GetConfig(false)

}

// NmslErrHandler 报错处理
func NmslErrHandler(err error) (_msg string) {
	concat := "xx"
	_msg = fmt.Sprintf("unable to fetch data from nmsl, pls concat %s", concat)
	if concat == "" {
		logger.WithError(err).Errorf("unable to read config file MASTERCONTACT")
	}
	return _msg

}

// AssKisserHandler 彩虹屁生成
func AssKisserHandler(from string) string {
	if from == "" {
		from = "xxx"
	}
	level := ""
	_levelParam := fmt.Sprintf("&level=%s", level)
	resp, err := http.Get(fmt.Sprintf(lickAPI, from) + _levelParam)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return fmt.Sprintf(string(body))
}

// MotherFuckerHandler 抠图芬芳
func MotherFuckerHandler(from string) string {
	if from == "" {
		from = "xxx"
	}
	level := ""
	_levelParam := fmt.Sprintf("&level=%s", level)
	resp, err := http.Get(fmt.Sprintf(nmslAPI, from) + _levelParam)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return fmt.Sprintf(string(body))
}

// ManifestFetchResponse 获取menifest接口返回
func ManifestFetchResponse() (Response ManifestWorldComponentContent, err error) {
	spaceClient := http.Client{
		Timeout: time.Second * 100, // Maximum of 100 secs
	}

	req, err := http.NewRequest(http.MethodGet, BunigieManifestUrl, nil)
	if err != nil {
		panic(err)

	}

	req.Header.Add("X-API-Key", config.BungieXApiKey)

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
func InfoMenifestBaseDBCheck(orm *gorm.DB) {
	// 检查是否存在表
	var wg sync.WaitGroup
	_Num := len(D2Table)
	log.Infof(fmt.Sprintf("正在检查%d张表...", _Num))
	wg.Add(_Num)
	var _InItSqls []string
	for tableName, tableSql := range D2Table {
		go DBCheckHandler(orm, tableName, tableSql, &_InItSqls, &wg)
	}
	wg.Wait()
	if len(_InItSqls) > 0 {
		// 直接调用执行-需等待率先你表后再进行后序的插入操作-阻塞
		utils.Execute(orm, strings.Join(_InItSqls, ";"), nil)
	}
	// 检查表里数据是否是最新-若不是或者数据为空-则重新抽数到数据库
	manifestRes, _ := ManifestFetchResponse()
	params := map[string]interface{}{"version": manifestRes.NewVersion}
	needUpdate := D2VersionHandler(orm, params)
	if needUpdate {
		// 强制更新数据-先清空后插入(以放原始数据被更改过)
		// 分批次写入-无需锁表
		// 写入中文数据
		ZhData := manifestRes.JsonWorldComponentContentPaths.ZhChs
		typ := reflect.TypeOf(ZhData)
		val := reflect.ValueOf(ZhData) //获取reflect.Type类型

		kd := val.Kind() //获取到a对应的类别
		if kd != reflect.Struct {
			log.Info("expect struct")
		}
		//获取到该结构体有几个字段
		num := val.NumField()

		//遍历结构体的所有字段
		for i := 0; i < num; i++ {
			tagVal := typ.Field(i).Tag.Get("json")
			ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), orm)

		}
	}
}

// ManifestFetchInfo 查询解析url数据并写入 InfoMenifestBaseDB 表
func ManifestFetchInfo(josnFile, tag string, orm *gorm.DB) {
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

	req.Header.Add("X-API-Key", config.BungieXApiKey)

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
			_params := []interface{}{itemid, _Description, _Name, _Icon, tag, _SeasonId}
			paramList = append(paramList, _params)
		}
	}

	// 写入数据库
	InsertMenifestHandler(orm, paramList)
	log.Infof(tag + " down!")
}

// PlayerBaseInfo 基础信息查询
func PlayerBaseInfo(steamId string) BaseprofileResult {
	// 构造url
	userUrl := baseprofileAPI(steamId)
	// 发送请求
	spaceClient := http.Client{
		Timeout: time.Second * 999, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	if err != nil {
		log.Warn(err)
	}

	req.Header.Add("X-API-Key", config.BungieXApiKey)

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
	var ResJson BaseprofileResult
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		log.Warn(jsonErr)
	}
	return ResJson
}

// AccountStats 数据实时解析返回
func AccountStatsFetchInfo(steamId string) AccountStatsResult {
	// 构造url
	userUrl := profileAPI(steamId)
	// 发送请求
	spaceClient := http.Client{
		Timeout: time.Second * 999, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	if err != nil {
		log.Warn(err)
	}

	req.Header.Add("X-API-Key", config.BungieXApiKey)

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
	var ResJson AccountStatsResult
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		log.Warn(jsonErr)
	}
	return ResJson
}

//
type ImgUrls struct {
	Zhoubao string `json:"zb"`
	Laojiu  string `json:"lj"`
	Shilian string `json:"sl"`
}

type Jsondata struct {
	Data interface{} `json:"data"`
}

type ReturnData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Jsondata
}

func DataInfo(flag string) string {
	var r ReturnData
	return r.gain(flag)
}

func (r *ReturnData) gain(flag string) string {

	//周报
	zhoubao, _, err := destinyWeekly("https://api.bilibili.com/x/article/list/web/articles?id=175327&jsonp=jsonp", 1, 0)
	//老九 试炼
	laojiu, shilian, err1 := destinyWeekly("https://api.bilibili.com/x/article/list/web/articles?id=175690&jsonp=jsonp", 1, 3)
	// 光尘商店
	dustDetailUrl := "https://cdn.jsdelivr.net/gh/azmiao/picture-bed/img/buy-13.jpg"
	if err != nil || err1 != nil {
		r.Code = 5001
		r.Msg = "获取失败"
	} else if flag == "week" {
		return string(zhoubao)
	} else if flag == "nine" {
		return string(laojiu)
	} else if flag == "trial" {
		return string(shilian)
	} else if flag == "dust" {
		return string(dustDetailUrl)
	} else {
		var imgurls = ImgUrls{Zhoubao: zhoubao, Laojiu: laojiu, Shilian: shilian}
		// 写入变量
		r.Data = imgurls
		r.Code = 200
		r.Msg = "获取成功"
		r.Time = time.Now().Format("2006-01-02 15:04:05")
	}
	g, _ := json.Marshal(r)
	return string(g)
}

/**
 *  哔哩哔哩 命运2周报
 */
func destinyWeekly(goUrl string, s int, m int) (imgUrl string, imgUrl2 string, err error) {
	bid := biList(goUrl)
	url := "https://www.bilibili.com/read/cv" + strconv.Itoa(bid) + "/?from=readlist"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		reg, err1 := regexp.Compile(`<img.*?data-src="(.*?)".*?>`)
		if err1 != nil {
			fmt.Println(err)
		}
		var respString string
		buf := make([]byte, 1024)
		for {
			n, _ := resp.Body.Read(buf)
			if n == 0 {
				break
			}
			respString += string(buf[:n])
		}
		data := reg.FindAllStringSubmatch(respString, -1)
		info := map[int]string{}
		for k, v := range data {
			info[k] = v[1]
		}
		imgUrl = "https:" + info[s]
		if m > 0 {
			imgUrl2 = "https:" + info[m]
		}
	}
	return imgUrl, imgUrl2, err
}

/**
 * 哔哩哔哩 json数据列 返回最新一条 ID
 */
func biList(url string) int {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	js, err := simplejson.NewJson([]byte(data))
	//放入map
	resultInfo, err := js.Get("data").Get("articles").Array()
	//获取最后一条
	info := js.Get("data").Get("articles").GetIndex((len(resultInfo) - 1))
	//取出最后一条ID
	s := info.Get("id").MustInt()
	return s
}
