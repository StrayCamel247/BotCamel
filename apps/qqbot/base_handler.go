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

func (a *qb) Serve(b *bot.Bot) {
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
