package destiny

/*
   __author__ : stray_camel
  __description__ :
  __REFERENCES__:
  __date__: 2021-03-12
*/
/*
 */
import (
	"encoding/json"
	// "fmt"
	// con "github.com/StrayCamel247/BotCamel/apps/config"
	// "github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	// "net/url"
	// "path"
	// "regexp"
	// "strconv"

	// "bytes"
	// "github.com/StrayCamel247/BotCamel/apps/utils"
	// "reflect"
	// "strings"
	// "sync"
	"time"
)

func init() {
}

const bungieApiKey = "aff47ade61f643a19915148cfcfc6d7d"

// PlayerBaseInfo 基础信息查询
func PlayerBaseInfo(steamId string) BaseprofileResult {
	// 构造url
	userUrl := baseprofileAPI(steamId)
	// 发送请求
	spaceClient := http.Client{
		Timeout: time.Second * 999, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	if err != nil {
		log.Warn(err)
	}

	req.Header.Add("X-API-Key", bungieApiKey)

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Warn(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Warn(readErr)
	}
	var ResJson BaseprofileResult
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		log.Warn(jsonErr)
	}
	return ResJson
}

// AccountStats 数据实时解析返回
func AccountStatsFetchInfo(steamId string) AccountStatsResult {
	// 构造url
	userUrl := profileAPI(steamId)
	// 发送请求
	spaceClient := http.Client{
		Timeout: time.Second * 999, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	if err != nil {
		log.Warn(err)
	}

	req.Header.Add("X-API-Key", bungieApiKey)

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Warn(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Warn(readErr)
	}
	var ResJson AccountStatsResult
	jsonErr := json.Unmarshal(body, &ResJson)
	if jsonErr != nil {
		log.Warn(jsonErr)
	}
	return ResJson
}
