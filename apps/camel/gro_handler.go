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
	// "github.com/Mrs4s/MiraiGo/client/pb/structmsg"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/baseapis"
	con "github.com/StrayCamel247/BotCamel/apps/config"
	"github.com/StrayCamel247/BotCamel/apps/handler"
	"github.com/StrayCamel247/BotCamel/global"
	// log "github.com/sirupsen/logrus"
	"io"

	"math/rand"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
	"time"
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

func d2uploadImgByUrl(flag string, url string, c *client.QQClient, msg *message.GroupMessage) error {
	var _imgFileDate string
	if handler.EqualFolds(flag, command.D2perk.Keys) {
		// 日更新
		_imgFileDate = GetD2daykDateOfdayk()
	} else {
		// 周更新
		_imgFileDate = GetD2WeekDateOfWeek()
	}
	fileName := fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
	if !PathExists(fileName) {
		err := downloadImg(fileName, url)
		if err != nil {
			return err
		}
	}
	if PathExists(fileName) {
		_img, err := c.UploadGroupImageByFile(msg.GroupCode, fileName)
		if err != nil {
			return err
		}
		m := message.NewSendingMessage().Append(_img)
		c.SendGroupMessage(msg.GroupCode, m)
	} else {
		fmt.Println("File downloading error")
	}
	return nil
}

func d2uploadImgByFlag(flag string, c *client.QQClient, msg *message.GroupMessage) error {
	out := baseapis.DataInfo(flag)
	err := d2uploadImgByUrl(flag, out, c, msg)
	if err != nil {
		return err
	}
	return nil
}

func downloadImg(filename, url string) error {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(filename)
	if err != nil {
		return err
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
	return nil
}
func getItemId(content string) {
	client := &http.Client{}
	//生成要访问的url
	url := "https://www.bungie.net/Platform/Destiny2/Manifest/"
	//提交请求
	reqest, _ := http.NewRequest("GET", url, nil)
	//增加header选项
	reqest.Header.Add("X-API-Key", "aff47ade61f643a19915148cfcfc6d7d")
	res, _ := client.Do(reqest)
	defer res.Body.Close()

}
func perkGenerateImg(content, flag string, c *client.QQClient, msg *message.GroupMessage) {
	baseUrl := fmt.Sprintf("https://www.light.gg/db/zh-cht/items/%s", content)
	// 参数
	// token := "604b1394d1a2d"
	url := url2.QueryEscape(baseUrl)
	width := 1280
	height := 800
	full_page := 1
	// 构造URL
	for _, v := range config.MasterShotTokens {
		query := "https://www.screenshotmaster.com/api/v1/screenshot"
		query += fmt.Sprintf("?token=%s&url=%s&width=%d&height=%d&full_page=%d",
			v, url, width, height, full_page)
		err := d2uploadImgByUrl(flag+content, query, c, msg)
		println(url, v, query)
		if err == nil {
			println(url)
			return
		}
	}

}

func dayGenerateImg(flag string, c *client.QQClient, msg *message.GroupMessage) {
	// 参数
	url := url2.QueryEscape("http://download.kamuxiy.top:88/destiny2/")
	width := 1280
	height := 800
	full_page := 1
	// 构造URL
	for _, v := range config.MasterShotTokens {
		query := "https://www.screenshotmaster.com/api/v1/screenshot"
		query += fmt.Sprintf("?token=%s&url=%s&width=%d&height=%d&full_page=%d",
			v, url, width, height, full_page)
		err := d2uploadImgByUrl(flag, query, c, msg)
		if err == nil {
			return
		}
	}
}

// GroMsgHandler 群聊信息获取并返回

func GroMsgHandler(c *client.QQClient, msg *message.GroupMessage) {
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
			perkGenerateImg(content, "perk", c, msg)

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
