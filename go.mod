module github.com/StrayCamel247/BotCamel

go 1.15

require (
	github.com/Baozisoftware/qrcode-terminal-go v0.0.0-20170407111555-c0650d8dff0f
	github.com/Logiase/MiraiGo-Template v0.0.0-20210309153626-69cfb14d2cd1
	github.com/Mrs4s/MiraiGo v0.0.0-20210323143736-d233c90d5083
	github.com/Mrs4s/go-cqhttp v0.9.40
	github.com/bitly/go-simplejson v0.5.0
	github.com/chromedp/cdproto v0.0.0-20210323015217-0942afbea50e
	github.com/chromedp/chromedp v0.6.10
	github.com/guonaihong/gout v0.1.6
	github.com/json-iterator/go v1.1.10
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/t-tomalak/logrus-easy-formatter v0.0.0-20190827215021-c074f06c5816
	github.com/tidwall/gjson v1.7.3
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/term v0.0.0-20210317153231-de623e64d2a6
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.6
)

replace github.com/Mrs4s/go-cqhttp => ./go-cqhttp
