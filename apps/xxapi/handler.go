package xxapi

/*
   __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-12
*/
/*
	沙雕api调用-夸人/骂人功能
*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const xxAPI string = "https://nmsl.shadiao.app/api.php?from=%s"
const lickAPI string = "https://chp.shadiao.app/api.php?from=%s"

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
	resp, err := http.Get(fmt.Sprintf(xxAPI, from) + _levelParam)

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
