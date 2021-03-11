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
	"github.com/Mrs4s/MiraiGo/message"
	"gopkg.in/yaml.v2"
	"strings"
)

func init() {
	bot.RegisterModule(instance)
}

var instance = &ar{}
var logger = utils.GetModuleLogger("QQBot_Handler")
var tem map[string]string

type ar struct {
}

func (a *ar) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "QQBot_Handler",
		Instance: instance,
	}
}

func (a *ar) Init() {
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

func (a *ar) PostInit() {
}

func (a *ar) Serve(b *bot.Bot) {
	b.OnGroupMessage(GroMsgHandler)

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {
		out := BaseAutoreply(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendPrivateMessage(msg.Sender.Uin, m)
	})
}

func (a *ar) Start(bot *bot.Bot) {
}

func (a *ar) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

// BaseAutoreply 根据配置的文本进行基础信息回复
func BaseAutoreply(in string) string {
	out, ok := tem[in]
	if !ok {
		for k, v := range tem {
			if strings.EqualFold(out, k) {
				return v
			}

		}
		return ""
	}
	return out
}
