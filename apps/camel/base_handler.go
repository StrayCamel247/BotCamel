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
	// "reflect"
	"strings"
)

// var logger = utils.GetModuleLogger("QQBot_Handler")
var tem map[string]string

// JSONConfig
var JSONConfig *global.JSONConfig
var command CommandsStruct

// GetConfig 获取当前配置文件信息
// func GetConfig() *global.JSONConfig {
// 	if JSONConfig != nil {
// 		return JSONConfig
// 	}
// 	conf := global.LoadConfig("./config.hjson")
// 	return conf
// }
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

// TODO: 通过指令结构体动态生成菜单栏
// GenerateMenu 通过指令结构体动态生成菜单栏
// func GenerateMenu(command CommandsStruct) (res string) {
// 	// 动态生成菜单
// 	res += "github.com/StrayCamel247/BotCamel\n"
// 	res += "快来领bug修修我吧~\n"
// 	res += "===== command =====\n"
// 	typ := reflect.TypeOf(command)
// 	val := reflect.ValueOf(command)
// 	comsNum := val.NumField()
// 	//遍历结构体的所有字段
// 	for i := 0; i < comsNum; i++ {
// 		// 获取结构体实例的反射类型对象
// 		if value, ok := val.Field(i).(CommandEleStruct{}); ok == true {
// 			fmt.Printf("x[%d] 类型为int,内容为%d\n", index, value)
// 		}
// 		//
// 		res += fmt.Sprintf("Field %d:值=%v\n", i, val.Field(i))
// 		//获取到struct标签，需要通过reflect.Type来获取tag标签的值
// 		tagVal := typ.Field(i).Tag.Get("json")
// 		//如果该字段有tag标签就显示，否则就不显示
// 		if tagVal != "" {
// 			res += fmt.Sprintf("Field %d:tag=%v\n", i, tagVal)
// 		}
// 	}
// 	return res
// }

// SendGroupMessage
func commandHandler(com string) string {
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
	switch {

	}
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
