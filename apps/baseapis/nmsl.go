package baseapis

import (
	"fmt"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"io/ioutil"
	"net/http"
)

var logger = utils.GetModuleLogger("NmslApi_handler")
var nmslAPI = "https://nmsl.shadiao.app/api.php?from=%s"
var lickAPI = "https://chp.shadiao.app/api.php?from=%s"

// NmslErrHandler 报错处理
func NmslErrHandler(err error) (_msg string) {
	concat := config.GlobalConfig.GetString("MASTERCONTACT")
	_msg = fmt.Sprintf("unable to fetch data from nmsl, pls concat %s", concat)
	if concat == "" {
		logger.WithError(err).Errorf("unable to read config file MASTERCONTACT")
	}
	return _msg

}

// AssKisserHandler 彩虹屁生成
func AssKisserHandler(from string) string {
	if from == "" {
		from = "xxx"
	}
	level := ""
	_levelParam := fmt.Sprintf("&level=%s", level)
	resp, err := http.Get(fmt.Sprintf(lickAPI, from) + _levelParam)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return fmt.Sprintf(string(body))
}

// MotherFuckerHandler 抠图芬芳
func MotherFuckerHandler(from string) string {
	if from == "" {
		from = "xxx"
	}
	level := ""
	_levelParam := fmt.Sprintf("&level=%s", level)
	resp, err := http.Get(fmt.Sprintf(nmslAPI, from) + _levelParam)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return fmt.Sprintf(string(body))
}
