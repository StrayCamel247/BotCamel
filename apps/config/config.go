package config

import (
	"github.com/StrayCamel247/BotCamel/global"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func init() {

}
func GetConfig(isFastStart bool) *global.JSONConfig {
	var conf *global.JSONConfig
	if global.PathExists("config.json") {
		conf = global.LoadConfig("config.json")
		_ = conf.Save("config.hjson")
		_ = os.Remove("config.json")
	} else if os.Getenv("UIN") != "" {
		log.Infof("将从环境变量加载配置.")
		uin, _ := strconv.ParseInt(os.Getenv("UIN"), 10, 64)
		pwd := os.Getenv("PASS")
		post := os.Getenv("HTTP_POST")
		conf = &global.JSONConfig{
			Uin:      uin,
			Password: pwd,
			HTTPConfig: &global.GoCQHTTPConfig{
				Enabled:  true,
				Host:     "0.0.0.0",
				Port:     5700,
				PostUrls: map[string]string{},
			},
			WSConfig: &global.GoCQWebSocketConfig{
				Enabled: true,
				Host:    "0.0.0.0",
				Port:    6700,
			},
			PostMessageFormat: "string",
			Debug:             os.Getenv("DEBUG") == "true",
		}
		if post != "" {
			conf.HTTPConfig.PostUrls[post] = os.Getenv("HTTP_SECRET")
		}
	} else {
		conf = global.LoadConfig("config.hjson")
	}
	if conf == nil {
		err := global.WriteAllText("config.hjson", global.DefaultConfigWithComments)
		if err != nil {
			log.Fatalf("创建默认配置文件时出现错误: %v", err)
			return nil
		}
		log.Infof("默认配置文件已生成, 请编辑 config.hjson 后重启程序.")
		if !isFastStart {
			time.Sleep(time.Second * 5)
		}
		return nil
	}
	return conf
}
