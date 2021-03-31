package apps

/*
	Bot Camel插件
*/
import (
	// "encoding/hex"
	// "io/ioutil"
	// "path"
	// "strconv"
	"strings"
	// "time"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/camel"
)

var format = "string"

func init() {
}

// SetMessageFormat 设置消息上报格式，默认为string
func SetMessageFormat(f string) {
	format = f
}

// AnalysisMsg
// 参数-客户端-消息列表
// 返回-是否@触发-指令-内容
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

// 收到加群邀请
func GroReciveInviteEvent(c *client.QQClient, e *client.GroupInvitedRequest) {
	c.SolveGroupJoinRequest(e, true, false, "")
}

// 群消息处理
func GroupMessageEvent(c *client.QQClient, m *message.GroupMessage) {
	isAt, com, content := camel.AnalysisMsg(c, m.Elements)
	if isAt {
		_camlMsg := camel.BaseMsg{Orm: dbGorm, Client: c, Message: m}
		_camlMsg.GroMsgHandler(com, content)
	}
}
