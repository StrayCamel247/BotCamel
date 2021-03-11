package camel

import "github.com/StrayCamel247/BotCamel/server"

// var logger = utils.GetModuleLogger("QQBot_Handler")
var tem map[string]string

func init() {

	config := server.GetConf()
	path := config.DialogueFilePath

	if path == "" {
		path = "./apps/base_default.yaml"
	}

	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &tem)
	if err != nil {
		logger.WithError(err).Errorf("unable to read config file in %s", path)
	}
}
