package baseapis

/*
   __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-12
*/
import (
	"encoding/json"
	"fmt"
	con "github.com/StrayCamel247/BotCamel/apps/config"
	"github.com/StrayCamel247/BotCamel/global"
	"github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	// "strings"
	"time"
)

var config *global.JSONConfig

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

// ManifestFetchJson 获取menifest接口返回的json文件的路径
func ManifestFetchJson() (jsonpath interface{}, err error) {
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, BunigieManifestUrl, nil)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")
	// req.Header.Add("X-API-Key", "aff47ade61f643a19915148cfcfc6d7d")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Warn(getErr)
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Warn(readErr)
		return nil, readErr
	}
	ResJson := ManifestResult{}
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		log.Warn(jsonErr)
		return nil, jsonErr
	}
	return ResJson.Response.JsonWorldComponentContentPaths.ZhChs, nil
}

// type  struct {
// 	info
// }
// func HttpRequest(method, url string) (*Request, error){

// }
// var JsonRes map[string]InfoDisplay
// ManifestFetchInfo 查询解析url数据并写入 InfoDisplayDB 表
func ManifestFetchInfo(josnFile, tag string, orm *gorm.DB, ch chan bool) {
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

	req.Header.Set("User-Agent", "spacecount-tutorial")
	// req.Header.Add("X-API-Key", "aff47ade61f643a19915148cfcfc6d7d")

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
	// 二维数组存在插入数据库的数据，没
	data := make([]InfoDisplayDB, 0)
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
		// 只记录有Description或者存在icon的数据
		// strings.Replace(ResJson[itemid].Properties.Description, "<", "",-1)
		// pat := "[0-9]+.[0-9]+"
		_Description := ResJson[itemid].Properties.Description
		_Icon := ResJson[itemid].Properties.Icon
		_Name := ResJson[itemid].Properties.Name
		if (_Description != "" || _Icon != "") && _handler(_Name) {
			// u.Path = path.Join(u.Path, _Icon)
			// iconurl := u.String()
			data = append(data, InfoDisplayDB{ItemId: itemid, Tag: tag, Icon: _Icon, Description: _Description, Name: _Name})
		}
	}

	// 写入数据库
	if len(data) > 0 {
		// 大于 800 分批处理
		// https://gitsea.com/2013/04/23/sqlite-too-many-sql-variables/
		// sqlite报错:Sqlite too many SQL variables
		_batch := len(data) / 800
		if _batch >= 1 {
			for i := 0; i < _batch; i++ {
				_data := data[i*800 : (i+1)*800]
				r := orm.Create(&_data)
				if r.Error != nil {
					log.Info(r.Error)
				}
			}
		} else {
			r := orm.Create(&data)
			if r.Error != nil {
				log.Info(r.Error)
			}
		}

	}
	log.Infof(tag + " down!")
	// jsonBytes, _ := json.Marshal(p)
	// 写入InfoDisplayDB表
	// return ResJson.Response.JsonWorldComponentContentPaths.ZhChs, nil
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
