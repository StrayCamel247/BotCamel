package qqbot

/*
  __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-10
*/
import (
	// "fmt"
	// "string"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
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
				isAt = true
			}
		case *message.TextElement:
			content = e.Content
			logger.Info(content)
		// case *message.ImageElement:
		// 	_msg += "[Image:" + e.Filename + "]"
		// 	logger.Info(_msg)
		// 	continue
		// case *message.FaceElement:
		// 	_msg += "[" + e.Name + "]"
		// 	logger.Info(_msg)
		// 	continue
		// case *message.GroupImageElement:
		// 	_msg += "[Image: " + e.ImageId + "]"
		// 	logger.Info(_msg)
		// 	continue
		// case *message.GroupFlashImgElement:
		// 	// NOTE: ignore other components
		// 	_msg = "[Image (flash):" + e.Filename + "]"
		// 	logger.Info(_msg)
		// 	continue
		// case *message.RedBagElement:
		// 	_msg += "[RedBag:" + e.Title + "]"
		// 	logger.Info(_msg)
		// 	continue
		// case *message.ReplyElement:
		// 	_msg += "[Reply:" + strconv.FormatInt(int64(e.ReplySeq), 10) + "]"
		// 	logger.Info(_msg)
		// 	continue
		default:
			break
		}
	}
	return isAt, content
}

// GroMsgHandler 群聊信息获取并返回
func GroMsgHandler(c *client.QQClient, msg *message.GroupMessage) {
	// fmt.Printf("用户信息: \n", "%+v", c)
	// fmt.Printf("消息信息: \n", "%+v", msg, "\n")
	// println(msg.Target)
	IsAt, content := AnalysisMsg(c.Uin, msg.Elements)
	out := BaseAutoreply(content)
	if out == "" && IsAt {
		out = "作甚😜"
	}
	m := message.NewSendingMessage().Append(message.NewText(out))
	c.SendGroupMessage(msg.GroupCode, m)
}
