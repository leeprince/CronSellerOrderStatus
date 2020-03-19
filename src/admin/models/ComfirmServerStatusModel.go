package models

import (
	"time"
	"CronSellerOrderStatus/src/admin/helper"
	"CronSellerOrderStatus/src/admin/db"
)

//待确认状态下，ctime 后用户未确认则更新为：完成服务状态
func ComfirmServerStatusModel(ctime int64) {
	helper.DebugLog("ComfirmServerStatusModel.待确认=>完成服务 time：", ctime)

	endTime := time.Now().Unix()
	sqlComfirmServerStatusModel := "UPDATE jgx_seller_order t1 LEFT JOIN jgx_seller_order_refund t2 ON t1.out_trade_no = t2.out_trade_no " +
		"SET t1.status = 5, t1.timing = 1, t1.end_time = ? " +
		"WHERE t2.id is null and t1.status = 4 and update_time <= ?"
	res, err := db.Engine.Exec(sqlComfirmServerStatusModel, helper.StrconvInt64ToString(endTime), helper.StrconvInt64ToString(ctime))
	helper.CheckErr(err, "ComfirmServerStatusModel 更新完成服务状态失败。")
	num, _ := res.RowsAffected()
	helper.DebugLog("ComfirmServerStatusModel 更新完成服务状态成功。影响行数: ", num)
}
