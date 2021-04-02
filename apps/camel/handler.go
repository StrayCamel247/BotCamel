package camel

/*
  __author__ : stray_camel
  __description__ :基础群聊功能
  __REFERENCES__:
  __date__: 2021-03-10
*/
import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	log "github.com/sirupsen/logrus"
	"strings"
)

// var config
func init() {
	command = CommandFilter()
	// config = con.GetConfig(false)
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

// 生成菜单消息
func menuHandler(c *client.QQClient, msg *message.GroupMessage) {
	_ImgMsg, err := c.UploadGroupImageByFile(msg.GroupCode, fmt.Sprintf("./tmp/%s.jpg", "menu"))
	if err != nil {
		log.WithError(err)
	}
	c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(_ImgMsg), true)
}
