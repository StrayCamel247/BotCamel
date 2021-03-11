package qqbot

/*
   __author__ : stray_camel
  __description__ : 消除处理基本处理逻辑
  __REFERENCES__: https://github.com/Logiase/MiraiGo-module-autoreply
  __date__: 2021-03-10
*/
import (
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	// "github.com/Mrs4s/MiraiGo/message"
	"github.com/StrayCamel247/BotCamel/apps/qqbot/baseapis"
	"gopkg.in/yaml.v2"
	"strings"
)

func init() {
	bot.RegisterModule(instance)
}

var instance = &qb{}
var logger = utils.GetModuleLogger("QQBot_Handler")
var tem map[string]string

type qb struct {
}

func (a *qb) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "QQBot_Handler",
		Instance: instance,
	}
}

func (a *qb) Init() {
	path := config.GlobalConfig.GetString("logiase.autoreply.path")

	if path == "" {
		path = "./autoreply.yaml"
	}

	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &tem)
	if err != nil {
		logger.WithError(err).Errorf("unable to read config file in %s", path)
	}
}

func (a *qb) PostInit() {
}

/*
func (c *QQClient) SolveGroupJoinRequest(i interface{}, accept, block bool, reason string) {
	if accept {
		block = false
		reason = ""
	}

	switch req := i.(type) {
	case *UserJoinGroupRequest:
		_, pkt := c.buildSystemMsgGroupActionPacket(req.RequestId, req.RequesterUin, req.GroupCode, func() int32 {
			if req.Suspicious {
				return 2
			} else {
				return 1
			}
		}(), false, accept, block, reason)
		_ = c.send(pkt)
	case *GroupInvitedRequest:
		_, pkt := c.buildSystemMsgGroupActionPacket(req.RequestId, req.InvitorUin, req.GroupCode, 1, true, accept, block, reason)
		_ = c.send(pkt)
	}
}
*/
func OnGroupInvitedHandler(c *client.QQClient, r *client.GroupInvitedRequest) {

}
func (a *qb) Serve(b *bot.Bot) {
	// 群组消息处理
	b.OnGroupMessage(GroMsgHandler)
	// 私人发消息处理
	b.OnPrivateMessage(PriMsgHandler)
	// 群组邀请加入消息处理
	/*
		func (c *QQClient) OnPrivateMessage(f func(*QQClient, *message.PrivateMessage)) {
			c.eventHandlers.privateMessageHandlers = append(c.eventHandlers.privateMessageHandlers, f)
		}
			func (c *QQClient) OnGroinviteupInvited(f func(*QQClient, *GroupInvitedRequest)) {
				c.eventHandlers.groupInvitedHandlers = append(c.eventHandlers.groupInvitedHandlers, f)
		}
	*/
	// b.OnGroupInvited(OnGroupInvitedHandler)
	// 新成员加入
	// b.OnGroupMemberJoined(nil)
}

func (a *qb) Start(bot *bot.Bot) {
}

func (a *qb) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func commandHandler(com string) string {
	if strings.EqualFold(com, "motherfucker") {
		_From := strings.TrimLeft(com, "Motherfucker")
		return baseapis.MotherFuckerHandler(_From)
	}
	if strings.EqualFold(com, "asskisser") {
		_From := strings.TrimLeft(com, "Asskisser")
		return baseapis.AssKisserHandler(_From)
	}
	return ""
}

// BaseAutoreply 根据配置的文本进行基础信息回复
func BaseAutoreply(in string) string {
	out, ok := tem[in]
	if !ok {
		for k, v := range tem {
			if strings.EqualFold(in, string(k)) {
				return v
			}
		}
		_arrayIn := strings.Split(in, " ")
		for _, _ele := range _arrayIn {
			return commandHandler(_ele)
		}

		out = ""
	}
	return out
}
