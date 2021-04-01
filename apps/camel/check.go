// 数据库检查-更新
package camel

import (
	log "github.com/sirupsen/logrus"
)

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
