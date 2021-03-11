package camel

import (
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/StrayCamel247/BotCamel/apps/baseapis"
	"github.com/StrayCamel247/BotCamel/global"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"strings"
)

// var logger = utils.GetModuleLogger("QQBot_Handler")
var tem map[string]string

// JSONConfig
var JSONConfig *global.JSONConfig

// GetConf 获取当前配置文件信息
func GetConf() *global.JSONConfig {
	if JSONConfig != nil {
		return JSONConfig
	}
	conf := global.LoadConfig("./config.hjson")
	return conf
}
func init() {

	config := GetConf()
	path := config.DialogueFilePath

	if path == "" {
		path = "./apps/base_default.yaml"
	}
	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &tem)
	if err != nil {
		log.WithError(err).Errorf("unable to read config file in %s", path)
	}
}

// SendGroupMessage
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
