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
	"github.com/StrayCamel247/BotCamel/apps/baseapis"
	con "github.com/StrayCamel247/BotCamel/apps/config"
	"github.com/StrayCamel247/BotCamel/apps/handler"
	"github.com/StrayCamel247/BotCamel/global"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	url2 "net/url"
	"os"
	// "reflect"
	"strings"
	"time"
)

// var bot *gomirai.Bot
var GroupMenu = "├─	Destiny 2\n│  ├─ 0x02 week 周报信息查询\n│  ├─ 0x03 day 日报信息查询\n│  ├─ 0x04 xiu 老九\n│  ├─ 0x05 trial 试炼信息查询\n│  ├─ 0x06 dust 光尘信息查询\n│  ├─ 0x07 random 掷骰子功能\n│  ├─ 0x08 perk 词条信息查询\n│  ├─ 0x09 item 物品信息查询\n│  ├─ 0x10 npc 查询npc信息\n│  ├─ 0x0a skill 查询技能等信息\n│  ├─ 0x0c pvp 查询pvp信息\n└─ more-devploping"
var config *global.JSONConfig

// var config
func init() {
	command = CommandFilter()
	config = con.GetConfig(false)
}

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

func d2uploadImgByUrl(flag string, url string, c *client.QQClient, msg *message.GroupMessage) (m *message.GroupImageElement, err error) {
	var _imgFileDate string
	if handler.EqualFolds(flag, command.D2xiu.Keys) || handler.EqualFolds(flag, command.D2day.Keys) {
		// 日更新
		_imgFileDate = GetD2daykDateOfdayk()
	} else if handler.EqualFolds(flag, command.D2week.Keys) || handler.EqualFolds(flag, command.D2trial.Keys) || handler.EqualFolds(flag, command.D2dust.Keys) {
		// 周更新 D2xiu D2week D2trial D2dust
		_imgFileDate = GetD2WeekDateOfWeek()
	}
	fileName := fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
	if !PathExists(fileName) {
		err := downloadImg(fileName, url)
		if err != nil {
			log.Println(err)
			// panic(err)
			return m, err
		}
	}
	// 上传磁盘内指定的文件
	if PathExists(fileName) {
		_img, err := c.UploadGroupImageByFile(msg.GroupCode, fileName)
		if err != nil {
			log.Println(err)
			// log.Warnf(err)
			return m, err
		}
		// m := message.NewSendingMessage().Append(_img)
		return _img, nil
		// c.SendGroupMessage(msg.GroupCode, m)
	} else {
		fmt.Println("File downloading error")
	}
	return m, nil
}

func d2uploadImgByFlag(flag string, c *client.QQClient, msg *message.GroupMessage) error {
	out := baseapis.DataInfo(flag)
	m, err := d2uploadImgByUrl(flag, out, c, msg)
	if err != nil {
		return err
	}
	c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(m))
	return nil
}

func downloadImg(filename, url string) error {
	res, err := http.Get(url)
	log.Info(fmt.Sprintf("正在下载%s", url))
	if err != nil {
		fmt.Println("图片下载失败；url")
		return err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
	return nil
}

func getItemId(content string, orm *gorm.DB) (itemids []string) {
	// 若表不存在-则创建表-并查询menifest接口解析json并写入数据
	// db.Create(&models.User{Profile: profile, Name: "silence"})
	baseapis.InfoMenifestBaseDBCheck(orm)

	// 获取item id
	// var results = []baseapis.ItemIdDB{}
	// _ = orm.Model(&baseapis.InfoMenifestBaseDB{}).Find(&results, baseapis.InfoMenifestBaseDB{Name: content})
	// for _, v := range results {
	// 	// 只返回固定tag的标签
	// 	if v.Tag == "DestinyInventoryItemLiteDefinition" {
	// 		itemids = append(itemids, v.ItemId)
	// 	}
	// 	// 将标签数据进行返回
	// 	if !handler.EqualFolds(v.Description, command.DesChecker.Keys) {
	// 		_des := strings.ReplaceAll(v.Description, "\n\n", "\n")
	// 		if !strings.Contains(des, _des) {
	// 			if des != "" {
	// 				des += "\n" + _des
	// 			} else {
	// 				des += _des
	// 			}

	// 		}
	// 	}

	// }
	itemids = IdQuery(orm, map[string]interface{}{"name": content})
	return itemids
}

// item 图片生成
func itemGenerateImg(content, flag string, c *client.QQClient, msg *message.GroupMessage, orm *gorm.DB) {

	itemId := getItemId(content, orm)
	// if err != nil {
	// 	panic(err)
	// }

	// 构造消息链-遍历返回的itemid在lightgg上进行批量截图-将图片传入消息链并返沪
	rMsg := message.NewSendingMessage()
	for _, v := range config.MasterShotTokens {

		// 上传文件是否报错
		_errFlag := false
		for _, _id := range itemId {
			baseUrl := fmt.Sprintf("https://www.light.gg/db/zh-cht/items/%s", _id)
			url := url2.QueryEscape(baseUrl)
			width := 1280
			height := 800
			full_page := 1
			query := "https://www.screenshotmaster.com/api/v1/screenshot"
			query += fmt.Sprintf("?token=%s&url=%s&width=%d&height=%d&full_page=%d",
				v, url, width, height, full_page)
			m, err := d2uploadImgByUrl(flag+_id, query, c, msg)
			rMsg = rMsg.Append(m)
			_errFlag = _errFlag || err != nil
		}
		if _errFlag {
			// 图片获取失败-重新构造消息链
			rMsg = message.NewSendingMessage()
		} else {
			// 图片调用成功
			if len(rMsg.Elements) > 0 {
				c.SendGroupMessage(msg.GroupCode, rMsg)
				// } else if des != "" {
				// 	c.SendGroupMessage(msg.GroupCode, rMsg.Append(message.NewText(des)))
			} else {
				c.SendGroupMessage(msg.GroupCode, rMsg.Append(message.NewText("哎呀~出错了🤣，报告问题：https://github.com/StrayCamel247/BotCamel/issues")))
			}
			return

		}
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
	url := "http://www.tianque.top/d2api/today/"
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
	out := BaseAutoreply("menu")
	out += GroupMenu
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(msg.GroupCode, m)
}

// 玩家pvp数据信息的概览获取
func PvPInfoHandler(content string, c *client.QQClient, msg *message.GroupMessage) {
	res := "===== PVP =====\n"
	// 基本信息
	BaseInfo := baseapis.PlayerBaseInfo(content)
	res += "Name: " + BaseInfo.Response.Profile.Data.UserInfo.DisplayName + "\n"
	// pvp记录信息
	AllData := baseapis.AccountStatsFetchInfo(content)

	PVPData := AllData.Response.MergedAllCharacters.Results.AllPvP.AllTime
	// ==================kda信息解析==================
	// 总体pvp信息
	// 解析pvp数据
	_dataHandler := func(e baseapis.AccountStatsInfo, time bool) (val string) {
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
	_dataPagHandler := func(e baseapis.AccountStatsInfo, time bool) (val string) {
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
