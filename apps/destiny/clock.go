package destiny

import (
	"github.com/robfig/cron"
)

// clickHandler 定时任务
func (r *Destiny) clickHandler() {
	c := cron.New()
	// 定每周三-1点-20分
	c.AddFunc("*0 20 1 * * 3", func() {
		go r.RefreshDayHandler("0x02", DataInfo("0x02"))
		// 检查数据库更新
		r.InfoMenifestBaseDBCheck()
	})
	// 每天-1点-20分触发
	c.AddFunc("*0 20 1 * * *", func() {
		go r.RefreshDayHandler("0x03", DayGenUrl)
		go r.RefreshDayHandler("0x04", DataInfo("0x04"))
		go r.RefreshDayHandler("0x05", DataInfo("0x05"))
		go r.RefreshDayHandler("0x06", DataInfo("0x06"))
	})
	c.Start()
}
