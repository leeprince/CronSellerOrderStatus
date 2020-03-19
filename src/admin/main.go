package main

import (
	"CronSellerOrderStatus/src/admin/db"
	"CronSellerOrderStatus/src/admin/helper"
	"CronSellerOrderStatus/src/admin/models"
	"github.com/robfig/cron"
	"time"
)

func main() {
	/*cron 定时任务*/
	spec := "*/5, *, *, *, *, ?"
	helper.InitConfig() // 初始化配置文件

	c := cron.New()
	c.AddFunc(spec, func() {
		helper.DebugLog("============================定时开始=========================", "")
		db.InitReadMysql()

		var cacelPayTime string = "-5m" // 待支付 => 取消支付
		var preReceiptTime string = "-10m" // 支付成功 => 取消接单
		var comfirmServerTime string = "-24h" // 待确认 => 完成服务
		cacelPayTimeC := calculationDurationTime(cacelPayTime)
		preReceiptTimeC := calculationDurationTime(preReceiptTime)
		comfirmServerTimeC := calculationDurationTime(comfirmServerTime)

		models.CancelPayStatusModel(cacelPayTimeC)
		models.CancelReceiptStatusModel(preReceiptTimeC)
		models.ComfirmServerStatusModel(comfirmServerTimeC)

		db.CloseDb()
		helper.DebugLog("============================定时结束=========================", "")

	})
	c.Start()
	select{} // 堵塞主进程不退出
	/*cron 定时任务_end*/
}

// 计算距离当前时间一段间隔后的时间戳
func calculationDurationTime(duration string) int64 {
	pTime, _ := time.ParseDuration(duration)
	cTime := time.Now().Add(pTime)
	return cTime.Unix()
}
