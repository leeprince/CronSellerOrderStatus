package models

import (
	"CronSellerOrderStatus/src/admin/db"
	"CronSellerOrderStatus/src/admin/helper"
	"time"
)

//待支付状态下，ctime 后用户未支付则转换为：支付取消状态
func CancelPayStatusModel(ctime int64) {
	helper.DebugLog("CancelPayStatusModel.待支付=>支付取消 time：", ctime)

	var tbJgxSellerOrder JgxSellerOrder
	tbJgxSellerOrder.EndTime = time.Now().Unix()
	tbJgxSellerOrder.Status = 2
	tbJgxSellerOrder.Timing = 1
	affected, err := db.Engine.Where("status =? and create_time <=?", 0, ctime).Update(&tbJgxSellerOrder)
	helper.CheckErr(err, "CancelPayStatusModel Update fail. ")
	helper.DebugLog("CancelPayStatusModel 更新支付取消状态成功. 影响行数: ", affected)
}
