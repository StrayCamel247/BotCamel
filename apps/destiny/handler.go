package destiny

import (
	"encoding/json"
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/lightGG"
	"github.com/StrayCamel247/BotCamel/apps/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	url2 "net/url"
	"time"
)

// 命运2插件结构体-消息片-非持久结构体
type Destiny struct {
	Orm *utils.CamelOrm
	Cli *client.QQClient
	Mes *message.GroupMessage
}

// 启动命运2插件
func Start() {
	d2 := Destiny{utils.Orm, nil, nil}
	log.Infof("命运2插件初始化开始；检查数据库；开启定时任务")
	d2.InfoMenifestBaseDBCheck()
	go d2.clickHandler()
	log.Infof("命运2插件初始化完成")
}

// 命运2-下载
func D2DownloadHandler(flag string, url string) (fileName string, updated bool) {
	var _imgFileDate string
	if utils.EqualFolds(flag, command.D2xiu.Keys) || utils.EqualFolds(flag, command.D2day.Keys) {
		// 日更新
		_imgFileDate = utils.GetDateViaWeekNum(0)
	} else if utils.EqualFolds(flag, command.D2week.Keys) || utils.EqualFolds(flag, command.D2trial.Keys) || utils.EqualFolds(flag, command.D2dust.Keys) {
		// 周更新 D2xiu D2week D2trial D2dust
		_imgFileDate = utils.GetDateViaWeekNum(3)
	}
	fileName = fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
	if !utils.PathExists(fileName) {
		// 文件不存在-下载文件
		log.Info(fmt.Sprintf("正在下载文件 url: %s", url))
		err := utils.DownloadImg(fileName, url)
		if err != nil {
			log.WithError(err)
		}
		log.Info(fmt.Sprintf("文件下载完成 url: %s", url))
		updated = true
	}
	return fileName, updated
}

// ==========业务处理==========

// fileNameGenerator 文件名生成器
func (r *Destiny) fileNameGenerator(flag string) string {
	var _imgFileDate string
	if utils.EqualFolds(flag, command.D2xiu.Keys) || utils.EqualFolds(flag, command.D2day.Keys) {
		// 日更新
		_imgFileDate = utils.GetDateViaWeekNum(0)
	} else if utils.EqualFolds(flag, command.D2week.Keys) || utils.EqualFolds(flag, command.D2trial.Keys) || utils.EqualFolds(flag, command.D2dust.Keys) {
		// 周更新 D2xiu D2week D2trial D2dust
		_imgFileDate = utils.GetDateViaWeekNum(3)
	}
	return fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
}

// 参数-url 链接
// 通过url发送图片消息
func (r *Destiny) d2uploadImgByUrl(flag string, url string) (m *message.GroupImageElement, err error) {
	fileName, _ := D2DownloadHandler(flag, url)
	// 上传磁盘内指定的文件
	if utils.PathExists(fileName) {
		_img, err := r.Cli.UploadGroupImageByFile(r.Mes.GroupCode, fileName)
		if err != nil {
			log.WithError(err)
			return m, err
		}
		return _img, nil
	} else {
		fmt.Println("图片获取失败")
	}
	return m, nil
}

// 通过物品名获取图片
func (r *Destiny) ItemGenerateImg(content, flag string) {
	itemId := r.IdQuery(map[string]interface{}{"name": content})
	// 构造消息链-遍历返回的itemid在lightgg上进行批量截图-将图片传入消息链并返沪
	rMsg := message.NewSendingMessage()
	// 生成文件名
	_fileName := r.fileNameGenerator(flag + content)
	// 文件不存在则生成-若存在则直接上传
	if !utils.PathExists(_fileName) {
		// 检查item-id是否为正确的item
		log.Infof("item检查网页...")
		var checkedUrl string
		for _, info := range itemId {
			baseUrl := fmt.Sprintf("https://www.light.gg/db/zh-cht/items/%s/%s", info[0], info[1])
			_ = url2.QueryEscape(info[1])
			// url = baseUrl
			if lightGG.LightGGChecker(baseUrl) {
				checkedUrl = baseUrl
			}

		}
		log.Infof("item网页检查完毕...")
		if checkedUrl != "" {
			log.Infof(fmt.Sprintf("[%s] 网页截图ing", checkedUrl))
			lightGG.UrlShotCutHandler(checkedUrl, _fileName)
			log.Infof(fmt.Sprintf("[%s] 网页截图完毕", checkedUrl))
		} else {
			log.Warnf(fmt.Sprintf("light 查无网页[%s]", flag+content))
		}
	}
	// 文件存在则上传
	if utils.PathExists(_fileName) {
		_ImgMsg, err := r.Cli.UploadGroupImageByFile(r.Mes.GroupCode, _fileName)
		if err != nil {
			log.WithError(err)
		}
		r.Cli.SendGroupMessage(r.Mes.GroupCode, rMsg.Append(_ImgMsg))
	} else {
		log.Warn(fmt.Sprintf("[%s]图片获取失败", flag+content))
		r.Cli.SendGroupMessage(r.Mes.GroupCode, rMsg.Append(message.NewText("哎呀~出错了🤣，报告问题：https://github.com/StrayCamel247/BotCamel/issues")))
	}

}

func (r *Destiny) dayGenerateImg(flag string) {
	url := DayGenUrl
	spaceClient := http.Client{
		Timeout: time.Second * 999, // Maximum of 10 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Warn(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

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
	var ResJson dayRes
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		log.Warn(jsonErr)
	}

	m, err := r.d2uploadImgByUrl(flag, ResJson.IMG_URL)

	r.Cli.SendGroupMessage(r.Mes.GroupCode, message.NewSendingMessage().Append(m))
	if err == nil {
		return
	}
}

// 试炼骰子功能
func (r *Destiny) randomHandler() {
	out := fmt.Sprintf("%d", rand.Intn(10))
	m := message.NewSendingMessage().Append(message.NewText(out))
	r.Cli.SendGroupMessage(r.Mes.GroupCode, m)
}

// 通过名称获取介绍信息
func (r *Destiny) GenerateDes(content, flag string) {

	des := r.DesQuery(map[string]interface{}{"name": content})

	// 构造消息链-遍历返回的itemid在lightgg上进行批量截图-将图片传入消息链并返沪
	rMsg := message.NewSendingMessage()
	if des != "" {
		r.Cli.SendGroupMessage(r.Mes.GroupCode, rMsg.Append(message.NewText(des)))
	} else {
		r.Cli.SendGroupMessage(r.Mes.GroupCode, rMsg.Append(message.NewText("哎呀~出错了🤣，报告问题：https://github.com/StrayCamel247/BotCamel/issues")))
	}
	return
}

// 根据-老九-试炼-光尘-等关键词获取并上传最新数据
func (r *Destiny) d2uploadImgByFlag(flag string) error {
	url := DataInfo(flag)
	m, err := r.d2uploadImgByUrl(flag, url)
	if err != nil {
		log.WithError(err)
		return err
	}
	r.Cli.SendGroupMessage(r.Mes.GroupCode, message.NewSendingMessage().Append(m))
	return nil
}

// 玩家pvp数据信息的概览获取
func (r *Destiny) pvpInfoHandler(content string) {
	res := "===== PVP =====\n"
	// 基本信息
	BaseInfo := PlayerBaseInfo(content)
	res += "Name: " + BaseInfo.Response.Profile.Data.UserInfo.DisplayName + "\n"
	// pvp记录信息
	AllData := AccountStatsFetchInfo(content)

	PVPData := AllData.Response.MergedAllCharacters.Results.AllPvP.AllTime
	// ==================kda信息解析==================
	// 总体pvp信息
	// 解析pvp数据
	_dataHandler := func(e AccountStatsInfo, time bool) (val string) {
		val += e.Basic.DisplayValue
		if !time {
			return val
		}
		return fmt.Sprintf("%.2f", e.Basic.Value/360)
	}
	res += "Total: "
	res += fmt.Sprintf("Kda %s/%s/%s-%s Suicides:%s Hours:%s ", _dataHandler(PVPData.Kills, false), _dataHandler(PVPData.Deaths, false), _dataHandler(PVPData.Assists, false), _dataHandler(PVPData.KillsDeathsAssists, false), _dataHandler(PVPData.Suicides, false), _dataHandler(PVPData.SecondsPlayed, true))
	// 场均pvp信息
	// 解析pvp数据
	_dataPagHandler := func(e AccountStatsInfo, time bool) (val string) {
		val += e.Pga.DisplayValue
		if !time {
			return val
		}
		return fmt.Sprintf("%.2f", e.Pga.Value/360)
	}
	res += "\nPer Ground: "
	res += fmt.Sprintf("Kda %s/%s/%s-%s Suicides:%s Hours:%s ", _dataPagHandler(PVPData.Kills, false), _dataPagHandler(PVPData.Deaths, false), _dataPagHandler(PVPData.Assists, false), _dataPagHandler(PVPData.KillsDeathsAssists, false), _dataPagHandler(PVPData.Suicides, false), _dataPagHandler(PVPData.SecondsPlayed, true))
	// 发送消息
	m := message.NewSendingMessage().Append(message.NewText(res))
	r.Cli.SendGroupMessage(r.Mes.GroupCode, m)

}

// 玩家PvE数据信息的概览获取
func (r *Destiny) pveInfoHandler(content string) {

}
