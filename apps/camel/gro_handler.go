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
	"github.com/StrayCamel247/BotCamel/apps/handler"
	log "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// var bot *gomirai.Bot
var GroupMenu = "├─	Destiny 2\n│  ├─ 0x02 week 周报信息查询\n│  └─ 0x03 xiu 老九信息查询\n│  └ 0x04 trial 试炼最新动态\n└─ more-devploping"

func init() {
	command = CommandFilter()
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
			slices, _ := c.GetWordSegmentation(com)
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
func d2uploadImgByUrl(flag string, url string, c *client.QQClient, msg *message.GroupMessage) {
	_imgFileDate := GetD2WeekDateOfWeek()
	fileName := fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
	if !PathExists(fileName) {
		downloadImg(fileName, url)
	}
	if PathExists(fileName) {
		_img, err := c.UploadGroupImageByFile(msg.GroupCode, fileName)
		if err != nil {
			panic(err)
		}
		m := message.NewSendingMessage().Append(_img)
		c.SendGroupMessage(msg.GroupCode, m)

	} else {
		fmt.Println("File downloading error")
	}
}
func d2uploadImgByFlag(flag string, c *client.QQClient, msg *message.GroupMessage) {
	out := baseapis.DataInfo(flag)
	d2uploadImgByUrl(flag, out, c, msg)
}
func downloadImg(filename, url string) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return
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
}
func perkGenerateImg(content, flag string, c *client.QQClient, msg *message.GroupMessage) {
	url := fmt.Sprintf("https://www.light.gg/db/zh-cht/items/%s", content)
	d2uploadImgByUrl(flag, url, c, msg)
}

// GroMsgHandler 群聊信息获取并返回
func GroMsgHandler(c *client.QQClient, msg *message.GroupMessage) {
	var out string
	IsAt, com, _ := AnalysisMsg(c, msg.Elements)
	if IsAt {
		out = BaseAutoreply(com)
		switch {
		// case
		case handler.EqualFolds(com, command.D2perk.Keys):
			// out += GroupMenu
			// m := message.NewSendingMessage().Append(message.NewText(out))
			// c.SendGroupMessage(msg.GroupCode, m)

		case handler.EqualFolds(com, command.D2week.Keys):
			d2uploadImgByFlag("week", c, msg)

		case handler.EqualFolds(com, command.D2xiu.Keys):
			d2uploadImgByFlag("nine", c, msg)

		case handler.EqualFolds(com, command.D2trial.Keys):
			d2uploadImgByFlag("trial", c, msg)

		case handler.EqualFolds(com, command.D2dust.Keys):
			d2uploadImgByFlag("dust", c, msg)

		case handler.EqualFolds(com, command.D2random.Keys):
			out := string(rand.Intn(10))
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
// 		println(out)
// 		message.NewSendingMessage().Append(message.NewAt(e.Operator.Uin, e.Operator.Nickname)).Append(message.NewText(out))
// 	} else {
// 		out := BaseAutoreply("0x00") + "\n"
// 		println(out)
// 		message.NewSendingMessage().Append(message.NewText(out))
// 		// log.Infof("Bot退出了群 %v.", formatGroupName(e.Group))
// 	}
// }
