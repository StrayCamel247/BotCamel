package camel

/*
   __author__ : stray_camel
  __description__ :基础聊天功能
  __REFERENCES__:
  __date__: 2021-03-12
*/
import (
	// "fmt"
	"github.com/StrayCamel247/BotCamel/apps/utils"
	// "github.com/StrayCamel247/BotCamel/global"
	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"

	"github.com/StrayCamel247/BotCamel/apps/xxapi"
	"strings"
	// "time"
)

var tem map[string]string

// JSONConfig
// var JSONConfig *global.JSONConfig
var command CommandsStruct

func init() {
	path := "./apps/base_default.yaml"
	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &tem)
	if err != nil {
		log.WithError(err).Errorf("unable to read config file in %s", path)
	}
	command = CommandFilter()
}

// SendGroupMessage
func xxApiHandler(com string) string {
	switch {
	case utils.EqualFolds(com, command.Asskisser.Keys):
		_From := strings.TrimLeft(com, "Asskisser")
		return xxapi.AssKisserHandler(_From)
	case utils.EqualFolds(com, command.Motherfucker.Keys):
		_From := strings.TrimLeft(com, "Motherfucker")
		return xxapi.MotherFuckerHandler(_From)
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
			return xxApiHandler(_ele)
		}
		out = ""
	}
	return out
}
