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
func AnalysisMsg(botUin int64, ele []message.IMessageElement) (isAt bool, content string) {
	// 解析消息体
	for _, elem := range ele {
		switch e := elem.(type) {

		case *message.AtElement:
			if botUin == e.Target {
				// qq聊天机器人当at机器人时触发
				isAt = true
			}
		case *message.TextElement:
			content = strings.TrimSpace(e.Content)
			log.Info(content)
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
	return isAt, content
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

func d2uploadImgByUrl(flag string, c *client.QQClient, msg *message.GroupMessage) {
	_imgFileDate := GetD2WeekDateOfWeek()
	out := baseapis.DataInfo(flag)
	fileName := fmt.Sprintf("./tmp/%s%s.jpg", flag, _imgFileDate)
	if !PathExists(fileName) {
		downloadImg(fileName, out)
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

// GroMsgHandler 群聊信息获取并返回
func GroMsgHandler(c *client.QQClient, msg *message.GroupMessage) {
	var out string
	IsAt, content := AnalysisMsg(c.Uin, msg.Elements)
	if IsAt {
		out = BaseAutoreply(content)
		switch {
		case handler.EqualFolds(content, command.Menu.Keys):
			out += GroupMenu
			m := message.NewSendingMessage().Append(message.NewText(out))
			c.SendGroupMessage(msg.GroupCode, m)

		case handler.EqualFolds(content, command.D2week.Keys):
			d2uploadImgByUrl("week", c, msg)

		case handler.EqualFolds(content, command.D2xiu.Keys):
			d2uploadImgByUrl("nine", c, msg)

		case handler.EqualFolds(content, command.D2trial.Keys):
			d2uploadImgByUrl("trial", c, msg)

		case handler.EqualFolds(content, command.D2dust.Keys):
			d2uploadImgByUrl("dust", c, msg)

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
