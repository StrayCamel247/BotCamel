package utils

import (
	"time"
)

// 通过数字获取当周指定周数的日期
// 0-1-2-3-4-5-6-7：分别对用当周和周一到周日
func GetDateViaWeekNum(weekNum int) (res string) {
	var format string = "2006-01-02"
	now := time.Now()
	if weekNum == 0 {
		res = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format(format)
		return
	}
	offset := weekNum - int(now.Weekday())
	res = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Format(format)
	return
}
