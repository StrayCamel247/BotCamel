// 周报-试炼等数据爬取b站数据
package destiny

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

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
	} else if flag == "0x02" {
		return string(zhoubao)
	} else if flag == "0x04" {
		return string(laojiu)
	} else if flag == "0x05" {
		return string(shilian)
	} else if flag == "0x06" {
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
