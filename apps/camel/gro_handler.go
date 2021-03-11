package camel

/*
  __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-10
*/
import (
	// "fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	log "github.com/sirupsen/logrus"
	"strings"
	// "strconv"
)

// AnalysisMsg 解析消息体的数据，对at类型、文本类型、链接、图片等不同格式的消息进行不同的处理
func AnalysisMsg(botUin int64, ele []message.IMessageElement) (isAt bool, content string) {
	// 解析消息体
	for _, elem := range ele {
		switch e := elem.(type) {

		case *message.AtElement:
			if botUin == e.Target {
				// qq聊天机器人当at机器人时触发
				println(e.Display)
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

// GroMsgHandler 群聊信息获取并返回
func GroMsgHandler(c *client.QQClient, msg *message.GroupMessage) {
	var out string
	IsAt, content := AnalysisMsg(c.Uin, msg.Elements)
	if IsAt {
		out = BaseAutoreply(content)
		switch content {
		default:
			if strings.EqualFold(content, "menu") {
				out += "--more--\ndeving..."
			}
			if out == "" {
				out = "作甚😜\nMenu即可查看功能菜单👻"
			}
		}
		/*

			type ReplyElement struct {
				ReplySeq int32
				Sender   int64
				Time     int32
				Elements []IMessageElement

				//original []*msg.Elem

				NewReply

			func NewReply(m *GroupMessage) *ReplyElement {
				return &ReplyElement{
					ReplySeq: m.Id,
					Sender:   m.Sender.Uin,
					Time:     m.Time,
					//original: m.OriginalElements,
					Elements: m.Elements,
				}
			}
			}
		*/
		// _AtEle = message.AtElement{Target: msg.Sender.Uin, Display: ""}
		m := message.NewSendingMessage().Append(message.NewText(out)).Append(message.NewReply(msg))
		c.SendGroupMessage(msg.GroupCode, m)
	}
}
