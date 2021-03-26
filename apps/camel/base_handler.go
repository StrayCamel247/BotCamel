package camel

/*
   __author__ : stray_camel
  __description__ :基础聊天功能
  __REFERENCES__:
  __date__: 2021-03-12
*/
import (
	// "fmt"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/StrayCamel247/BotCamel/apps/baseapis"
	"github.com/StrayCamel247/BotCamel/apps/handler"
	"github.com/StrayCamel247/BotCamel/global"
	log "github.com/sirupsen/logrus"

	con "github.com/StrayCamel247/BotCamel/apps/config"
	"gopkg.in/yaml.v2"
	"strings"
	// "time"
)

var tem map[string]string

// JSONConfig
var JSONConfig *global.JSONConfig
var command CommandsStruct

func init() {
	config := con.GetConfig(false)
	path := config.DialogueFilePath

	if path == "" {
		path = "./apps/base_default.yaml"
	}
	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &tem)
	if err != nil {
		log.WithError(err).Errorf("unable to read config file in %s", path)
	}
	command = CommandFilter()
}

// SendGroupMessage
func shadiaoApiHandler(com string) string {
	switch {
	case handler.EqualFolds(com, command.Asskisser.Keys):
		_From := strings.TrimLeft(com, "Asskisser")
		return baseapis.AssKisserHandler(_From)
	case handler.EqualFolds(com, command.Motherfucker.Keys):
		_From := strings.TrimLeft(com, "Motherfucker")
		return baseapis.MotherFuckerHandler(_From)
		// case handler.EqualFolds(com, command.Menu.Keys):

		// 	return GenerateMenu(command)
	}
	return ""
}

// BaseAutoreply 根据配置的文本进行基础信息回复
func BaseAutoreply(in string) string {
	out, ok := tem[in]
	if !ok {
		// 查询对话配置表 apps\base_default.yaml
		for k, v := range tem {
			if strings.EqualFold(in, string(k)) {
				return v
			}
		}
		// 查询沙雕api
		_arrayIn := strings.Split(in, " ")
		for _, _ele := range _arrayIn {
			return shadiaoApiHandler(_ele)
		}
		out = ""
	}
	return out
}

// 日报信息更新
func RefreshDayHandler(flag, url string) {
	ch := make(chan bool)
	// 每秒轮询日报信息-是否更新-若更新跳出
	go func() {
		for {
			_, updated := D2DownloadHandler(flag, url)
			if updated {
				log.Infof("定时器-日报数据已更新！")
				ch <- updated
			}
		}
	}()
	_ = <-ch
}
