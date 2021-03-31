package camel

/*
  __author__ : stray_camel
  __description__ :群聊功能
  __REFERENCES__:
  __date__: 2021-03-10
*/
import (
	// "fmt"
	"bufio"
	"fmt"
	// "github.com/Logiase/gomirai"
	"encoding/json"
	"github.com/Mrs4s/MiraiGo/client"
	"gorm.io/gorm"
	// "github.com/Mrs4s/MiraiGo/client/pb/structmsg"
	"github.com/Mrs4s/MiraiGo/message"
	// con "github.com/StrayCamel247/BotCamel/apps/config"
	"github.com/StrayCamel247/BotCamel/apps/destiny"
	"github.com/StrayCamel247/BotCamel/apps/utils"
	"github.com/StrayCamel247/BotCamel/apps/lightGG"
	// "github.com/StrayCamel247/BotCamel/global"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	// url2 "net/url"
	"os"
	// "reflect"
	// "regexp"
	// "strconv"
	"strings"
	"time"
)

// var config *global.JSONConfig

// var config
func init() {
	command = CommandFilter()
	// config = con.GetConfig(false)
}

// 常量声明
const DayGenUrl string = "http://www.tianque.top/d2api/today/"

// AnalysisMsg 解析消息体的数据，对at类型、文本类型、链接、图片等不同格式的消息进行不同的处理

func AnalysisMsg(c *client.QQClient, ele []message.IMessageElement) (isAt bool, com, content string) {
	// 解析消息体
	for _, elem := range ele {
		switch e := elem.(type) {

		case *message.AtElement:
			if c.Uin == e.Target {
				isAt = true
			}
		case *message.TextElement:
			com = strings.TrimSpace(e.Content)
			slices := strings.Fields(com)
			if len(slices) < 1 {
				break
			} else if len(slices) >= 2 {
				content = slices[1]

			}
			com = slices[0]
			// log.Info(com)
		// case *message.ImageElement:
		// 	_msg += "[Image:" + e.Filename + "]"
		// 	log.Info(_msg)
		// 	continue
		// case *message.FaceElement:
		// 	_msg += "[" + e.Name + "]"
		// 	log.Info(_msg)
		// 	continue
		// case *message.GroupImageElement:
		// 	_msg += "[Image: " + e.ImageId + "]"
		// 	log.Info(_msg)
		// 	continue
		// case *message.GroupFlashImgElement:
		// 	// NOTE: ignore other components
		// 	_msg = "[Image (flash):" + e.Filename + "]"
		// 	log.Info(_msg)
		// 	continue
		// case *message.RedBagElement:
		// 	_msg += "[RedBag:" + e.Title + "]"
		// 	log.Info(_msg)
		// 	continue
		// case *message.ReplyElement:
		// 	_msg += "[Reply:" + strconv.FormatInt(int64(e.ReplySeq), 10) + "]"
		// 	log.Info(_msg)
		// 	continue
		default:
			break
		}
	}
	return isAt, com, content
}

func GetD2WeekDateOfWeek() string {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -4
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := weekStartDate.Format("2006-01-02")
	return weekMonday
}
func GetD2daykDateOfdayk() string {
	now := time.Now()
	currentDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02")
	return currentDay
}
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("File reading error", err)
	return false
}

// FileNameGenerator 文件名生成器
func FileNameGenerator(flag string) string {
	var _imgFileDate string
	if utils.EqualFolds(flag, command.D2xiu.Keys) || utils.EqualFolds(flag, command.D2day.Keys) {
		// 日更新
		_imgFileDate = GetD2daykDateOfdayk()
	} else if utils.EqualFolds(flag, command.D2week.Keys) || utils.EqualFolds(flag, command.D2trial.Keys) || utils.EqualFolds(flag, command.D2dust.Keys) {
		// 周更新 D2xiu D2week D2trial D2dust
		_imgFileDate = GetD2WeekDateOfWeek()
	}
	return fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
}
func D2DownloadHandler(flag string, url string) (fileName string, updated bool) {
	var _imgFileDate string
	if utils.EqualFolds(flag, command.D2xiu.Keys) || utils.EqualFolds(flag, command.D2day.Keys) {
		// 日更新
		_imgFileDate = GetD2daykDateOfdayk()
	} else if utils.EqualFolds(flag, command.D2week.Keys) || utils.EqualFolds(flag, command.D2trial.Keys) || utils.EqualFolds(flag, command.D2dust.Keys) {
		// 周更新 D2xiu D2week D2trial D2dust
		_imgFileDate = GetD2WeekDateOfWeek()
	}
	fileName = fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
	if !PathExists(fileName) {
		// 文件不存在-下载文件
		log.Info(fmt.Sprintf("正在下载文件 url: %s", url))
		err := downloadImg(fileName, url)
		if err != nil {
			log.WithError(err)
		}
		log.Info(fmt.Sprintf("文件下载完成 url: %s", url))
		updated = true
	}
	return fileName, updated
}

func d2uploadImgByUrl(flag string, url string, c *client.QQClient, msg *message.GroupMessage) (m *message.GroupImageElement, err error) {
	fileName, _ := D2DownloadHandler(flag, url)
	// 上传磁盘内指定的文件
	if PathExists(fileName) {
		_img, err := c.UploadGroupImageByFile(msg.GroupCode, fileName)
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

// 根据-老九-试炼-光尘-等关键词获取并上传最新数据
func d2uploadImgByFlag(flag string, c *client.QQClient, msg *message.GroupMessage) error {
	out := destiny.DataInfo(flag)
	m, err := d2uploadImgByUrl(flag, out, c, msg)
	if err != nil {
		log.WithError(err)
		return err
	}
	c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(m))
	return nil
}

// 文件下载
func downloadImg(filename, url string) error {
	// 记录下载时间
	_nowTime := time.Now()
	_timeCostLogger := func(start time.Time) {
		tc := time.Since(start)
		log.Info(fmt.Sprintf("time cost = %v\n", tc))
	}
	defer _timeCostLogger(_nowTime)
	// 构造请求头
	spaceClient := http.Client{
		// 请求时间
		Timeout: time.Minute * 10,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Warn(err)
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Warn(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		fmt.Println("图片下载失败；url")
		log.WithError(err)
		return err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(filename)
	if err != nil {
		log.WithError(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
	return nil
}

// getItemId 通过名字获取对应的itemid
func getItemId(content string, orm *gorm.DB) (itemids [][2]string) {
	itemids = IdQuery(orm, map[string]interface{}{"name": content})
	return itemids
}

// ItemGenerateImg
/*
	item 物品查询并上传
*/
func ItemGenerateImg(content, flag string, c *client.QQClient, msg *message.GroupMessage, orm *gorm.DB) {
	itemId := getItemId(content, orm)
	// 构造消息链-遍历返回的itemid在lightgg上进行批量截图-将图片传入消息链并返沪
	rMsg := message.NewSendingMessage()
	// 生成文件名
	_fileName := FileNameGenerator(flag + content)
	// 文件不存在则生成-若存在则直接上传
	if !utils.PathExists(_fileName) {
		// 检查item-id是否为正确的item
		log.Infof("item检查网页...")
		var checkedUrl string
		for _, info := range itemId {
			baseUrl := fmt.Sprintf("https://www.light.gg/db/zh-cht/items/%s/%s/", info[0], info[1])
			// url := url2.QueryEscape(baseUrl)
			// url = baseUrl
			if lightGG.LightGGChecker(baseUrl) {
				checkedUrl = baseUrl
			}

		}
		log.Infof("item网页检查完毕...")
		if checkedUrl != "" {
			log.Infof(fmt.Sprintf("[%s]网页截图ing", checkedUrl))
			lightGG.UrlShotCutHandler(checkedUrl, _fileName)
			log.Infof(fmt.Sprintf("%s网页截图完毕", checkedUrl))
		} else {
			log.Warnf(fmt.Sprintf("light 查无网页 %s", content+flag))
		}
	}
	// 文件存在则上传
	if PathExists(_fileName) {
		_ImgMsg, err := c.UploadGroupImageByFile(msg.GroupCode, _fileName)
		if err != nil {
			log.WithError(err)
		}
		c.SendGroupMessage(msg.GroupCode, rMsg.Append(_ImgMsg))
	} else {
		log.Warn(fmt.Sprintf("%s图片获取失败", _fileName))
		c.SendGroupMessage(msg.GroupCode, rMsg.Append(message.NewText("哎呀~出错了🤣，报告问题：https://github.com/StrayCamel247/BotCamel/issues")))
	}

}

// 介绍生成
func GenerateDes(content, flag string, c *client.QQClient, msg *message.GroupMessage, orm *gorm.DB) {

	des := DesQuery(orm, map[string]interface{}{"name": content})

	// 构造消息链-遍历返回的itemid在lightgg上进行批量截图-将图片传入消息链并返沪
	rMsg := message.NewSendingMessage()
	if des != "" {
		c.SendGroupMessage(msg.GroupCode, rMsg.Append(message.NewText(des)))
	} else {
		c.SendGroupMessage(msg.GroupCode, rMsg.Append(message.NewText("哎呀~出错了🤣，报告问题：https://github.com/StrayCamel247/BotCamel/issues")))
	}
	return
}

func dayGenerateImg(flag string, c *client.QQClient, msg *message.GroupMessage) {
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

	m, err := d2uploadImgByUrl(flag, ResJson.IMG_URL, c, msg)

	c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(m))
	if err == nil {
		return
	}
}
func randomHandler(c *client.QQClient, msg *message.GroupMessage) {
	out := fmt.Sprintf("%d", rand.Intn(10))
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(msg.GroupCode, m)
}
func menuHandler(c *client.QQClient, msg *message.GroupMessage) {
	out := `
		
	`
	m := message.NewSendingMessage().Append(message.NewText(string(out)))
	c.SendGroupMessage(msg.GroupCode, m)
}

// 玩家pvp数据信息的概览获取
func PvPInfoHandler(content string, c *client.QQClient, msg *message.GroupMessage) {
	res := "===== PVP =====\n"
	// 基本信息
	BaseInfo := destiny.PlayerBaseInfo(content)
	res += "Name: " + BaseInfo.Response.Profile.Data.UserInfo.DisplayName + "\n"
	// pvp记录信息
	AllData := destiny.AccountStatsFetchInfo(content)

	PVPData := AllData.Response.MergedAllCharacters.Results.AllPvP.AllTime
	// ==================kda信息解析==================
	// 总体pvp信息
	// 解析pvp数据
	_dataHandler := func(e destiny.AccountStatsInfo, time bool) (val string) {
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
	_dataPagHandler := func(e destiny.AccountStatsInfo, time bool) (val string) {
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
	c.SendGroupMessage(msg.GroupCode, m)

}

// 玩家PvE数据信息的概览获取
func PvEInfoHandler(content string, c *client.QQClient, msg *message.GroupMessage) {

}
