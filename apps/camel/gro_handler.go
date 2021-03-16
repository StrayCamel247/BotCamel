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
	"math/rand"
	"net/http"
	url2 "net/url"
	"os"
	"reflect"
	"strings"
	"time"
	// "io/ioutil"
)

// var bot *gomirai.Bot
var GroupMenu = "├─	Destiny 2\n│  ├─ 0x02 week 周报信息查询\n│  ├─ 0x03 day 日报信息查询\n│  ├─ 0x04 xiu 老九\n│  ├─ 0x05 trial 试炼信息查询\n│  ├─ 0x06 dust 光尘信息查询\n│  ├─ 0x07 random 掷骰子功能\n└─ more-devploping"
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
				// qq聊天机器人当at机器人时触发
				isAt = true
			}
		case *message.TextElement:
			com = strings.TrimSpace(e.Content)
			// slices, _ := c.GetWordSegmentation(com)
			slices := strings.Fields(com)
			print(len(slices))
			for _, v := range slices {
				print(v)
			}
			if len(slices) < 1 {
				break
			} else if len(slices) >= 2 {
				content = slices[1]

			}
			com = slices[0]
			print(com, content)
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
			return m, nil
		}
	}
	if PathExists(fileName) {
		_img, err := c.UploadGroupImageByFile(msg.GroupCode, fileName)
		if err != nil {
			return m, nil
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

func getItemId(content string, orm *gorm.DB) (itemids []string, des string, err error) {
	// 若表不存在-则创建表-并查询menifest接口解析json并写入数据
	// db.Create(&models.User{Profile: profile, Name: "silence"})
	isexisted, err := baseapis.InfoDisplayDBCheck(orm)
	if err != nil {
		// 数据库校验报错-直接返回
		return itemids, des, nil
	}
	if !isexisted {
		// 若数据库表不存在，并发查询数据并写入
		file, _ := baseapis.ManifestFetchJson(content)

		typ := reflect.TypeOf(file)
		val := reflect.ValueOf(file) //获取reflect.Type类型

		kd := val.Kind() //获取到a对应的类别
		if kd != reflect.Struct {
			fmt.Println("expect struct")
			return
		}
		//获取到该结构体有几个字段
		num := val.NumField()

		//遍历结构体的所有字段
		start := time.Now()
		ch := make(chan bool)
		for i := 0; i < num; i++ {
			// goroutine的正确用法
			// 那怎么用goroutine呢？有没有像Python多进程/线程的那种等待子进/线程执行完的join方法呢？当然是有的，可以让Go 协程之间信道（channel）进行通信：从一端发送数据，另一端接收数据，信道需要发送和接收配对，否则会被阻塞：
			// fmt.Printf("Field %d:值=%v\n", i, val.Field(i))
			tagVal := typ.Field(i).Tag.Get("json")
			//如果该字段有tag标签就显示，否则就不显示
			// if tagVal != "" {
			// 	fmt.Printf("Field %d:tag=%v\n", i, tagVal)
			// }
			// 并发
			// go baseapis.ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), orm, ch)
			// 串行
			print(tagVal)
			baseapis.ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), orm, ch)
			// if tagVal == "DestinyInventoryItemLiteDefinition" {
			// 	baseapis.ManifestFetchInfo(fmt.Sprintf("%v", val.Field(i)), fmt.Sprintf("%v", tagVal), orm, ch)
			// }

		}
		elapsed := time.Since(start)
		fmt.Printf("Took %s", elapsed)

		// println(file)
	}
	// 获取item id
	var results = []baseapis.ItemIdDB{}
	_ = orm.Model(&baseapis.InfoDisplayDB{}).Find(&results, baseapis.InfoDisplayDB{Name: content})
	for _, v := range results {
		// 只返回固定tag的标签
		if v.Tag == "DestinyInventoryItemLiteDefinition" {
			itemids = append(itemids, v.ItemId)
		}
		if v.Description != "" {
			des += strings.ReplaceAll(v.Description, "\n\n", "\n")
		}

		// 对item id进行判断是否可获取perk
	}
	return itemids, des, nil
}

func perkGenerateImg(content, flag string, c *client.QQClient, msg *message.GroupMessage, orm *gorm.DB) {

	itemId, des, err := getItemId(content, orm)
	if err != nil {
		panic(err)
	}

	// 构造消息链-遍历返回的itemid在lightgg上进行批量截图-将图片传入消息链并返沪
	rMsg := message.NewSendingMessage()
	// c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(m))
	// 构造URL
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
			c.SendGroupMessage(msg.GroupCode, rMsg.Append(message.NewText(des)))
			return
		}
	}
}

func dayGenerateImg(flag string, c *client.QQClient, msg *message.GroupMessage) {
	// 参数
	// url := url2.QueryEscape("http://www.tianque.top/d2api/today/")
	// width := 1280
	// height := 800
	// full_page := 1
	// 构造URL
	m, err := d2uploadImgByUrl(flag, "http://www.tianque.top/d2api/today/", c, msg)

	c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(m))
	if err == nil {
		return
	}
}

// GroMsgHandler 群聊信息获取并返回

func GroMsgHandler(orm *gorm.DB, c *client.QQClient, msg *message.GroupMessage) {
	var out string
	IsAt, com, content := AnalysisMsg(c, msg.Elements)
	if IsAt {
		out = BaseAutoreply(com)
		switch {
		// case
		case handler.EqualFolds(com, command.Menu.Keys):
			// content := com
			out = BaseAutoreply("menu")
			out += GroupMenu
			m := message.NewSendingMessage().Append(message.NewText(out))
			c.SendGroupMessage(msg.GroupCode, m)
		// case
		case handler.EqualFolds(com, command.D2perk.Keys):
			// content := com
			perkGenerateImg(content, "perk", c, msg, orm)

		case handler.EqualFolds(com, command.D2day.Keys):
			dayGenerateImg("day", c, msg)

		case handler.EqualFolds(com, command.D2week.Keys):
			_ = d2uploadImgByFlag("week", c, msg)

		case handler.EqualFolds(com, command.D2xiu.Keys):
			_ = d2uploadImgByFlag("nine", c, msg)

		case handler.EqualFolds(com, command.D2trial.Keys):
			_ = d2uploadImgByFlag("trial", c, msg)

		case handler.EqualFolds(com, command.D2dust.Keys):
			_ = d2uploadImgByFlag("dust", c, msg)

		case handler.EqualFolds(com, command.D2random.Keys):
			out := fmt.Sprintf("%d", rand.Intn(10))
			m := message.NewSendingMessage().Append(message.NewText(out))
			c.SendGroupMessage(msg.GroupCode, m)

		case out == "":
			out = "作甚😜\nmenu-菜单"
			m := message.NewSendingMessage().Append(message.NewText(out))
			c.SendGroupMessage(msg.GroupCode, m)

		default:
			m := message.NewSendingMessage().Append(message.NewText(out))
			c.SendGroupMessage(msg.GroupCode, m)
		}

	}
}

// 收到加群邀请

func GroReciveInviteHandler(c *client.QQClient, e *client.GroupInvitedRequest) {
	print("testtest")
	c.SolveGroupJoinRequest(e, true, false, "")
}

// 加入群聊

func GroJoinHandler(c *client.QQClient, group *client.GroupInfo) {
	out := BaseAutoreply("f48dcc50457d") + "\n"
	out += BaseAutoreply("0x00") + "\n"
	out += BaseAutoreply("menu")
	out += GroupMenu
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(group.Code, m)
}

// 离开群聊-都被t了怎么发文字，，，，开发了个寂寞

// func GroLeaveHandler(c *client.QQClient, e *client.GroupLeaveEvent) {
// 	if e.Operator != nil {
// 		out := BaseAutoreply("0x01") + "\n"
// 		message.NewSendingMessage().Append(message.NewAt(e.Operator.Uin, e.Operator.Nickname)).Append(message.NewText(out))
// 	} else {
// 		out := BaseAutoreply("0x00") + "\n"
// 		message.NewSendingMessage().Append(message.NewText(out))
// 		// log.Infof("Bot退出了群 %v.", formatGroupName(e.Group))
// 	}
// }
