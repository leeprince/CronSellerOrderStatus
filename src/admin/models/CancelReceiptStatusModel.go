package models

import (
	"github.com/objcoding/wxpay"
	"time"
	"CronSellerOrderStatus/src/admin/helper"
	"CronSellerOrderStatus/src/admin/db"
	"CronSellerOrderStatus/src/admin/libraries"
)

//用户支付成功状态下, ctime 商家还未接单则更新状态为：取消接单， 并自动退款
func CancelReceiptStatusModel(ctime int64) {
	helper.DebugLog("CancelReceiptStatusModel.支付成功=>取消接单，并退款。 time：", ctime)

	paytime := helper.StrconvInt64ToString(ctime)
	sqlSellerOrderRefund := "SELECT t1.uuid, t1.transaction_id, t1.out_trade_no, t1.cash_fee, t1.mobile, t1.pay_type " +
		"FROM jgx_seller_order t1 LEFT JOIN jgx_seller_order_refund t2 on t2.out_trade_no = t1.out_trade_no " +
		"WHERE t2.id is null and t1.status = 1 and t1.pay_time <= " + paytime
	result, err := db.Engine.QueryString(sqlSellerOrderRefund)
	helper.CheckErr(err, "条件查询支付成功状态下， 用户未退款商家 & 未接单的订单")

	helper.DebugLog("CancelReceiptStatusModel.支付成功=>取消接单-数据：", result)
	currentTime := time.Now().Unix()
	for _, v := range result {
		var dataJgxSellerOrderRefund JgxSellerOrderRefund
		dataJgxSellerOrderRefund.Uuid = v["uuid"]
		dataJgxSellerOrderRefund.TransactionId = v["transaction_id"]
		dataJgxSellerOrderRefund.OutTradeNo = v["out_trade_no"]
		dataJgxSellerOrderRefund.CashFee = v["cash_fee"]
		dataJgxSellerOrderRefund.RefundFee = v["cash_fee"]
		dataJgxSellerOrderRefund.Mobile = v["mobile"]
		dataJgxSellerOrderRefund.PayType = v["pay_type"]

		OutRefundNo := helper.GetOutTradeNo("s")
		dataJgxSellerOrderRefund.OutRefundNo = OutRefundNo
		dataJgxSellerOrderRefund.RefundReason = "系统自动处理取消接单"
		dataJgxSellerOrderRefund.Status = 0
		dataJgxSellerOrderRefund.Accept = 1
		dataJgxSellerOrderRefund.AcceptTime = currentTime
		dataJgxSellerOrderRefund.ApplyRole = 4
		dataJgxSellerOrderRefund.Source = 1

		time.Sleep(1000)
		// 事务
		engineSession := db.Engine.NewSession()
		defer engineSession.Close()
		err := engineSession.Begin()

		_, err = engineSession.Insert(&dataJgxSellerOrderRefund)
		if err != nil {
			helper.ErrorLog("CancelReceiptStatusModel 插入退款订单失败:", dataJgxSellerOrderRefund)
			engineSession.Rollback()
		}
		helper.DebugLog("CancelReceiptStatusModel 插入退款订单成功。插入ID: ", dataJgxSellerOrderRefund.Id)

		isNotFree := false // 免费订单不操作退款
		if helper.StrconvStringToFloat64(v["cash_fee"]) != 0 {
			isNotFree = true
		}

		var refundStatus bool = true
		var isHaveRefund bool = false
		refundResult := wxpay.Params{}
		if isNotFree == true {
			// 开始退款
			totalFee :=  helper.StrconvStringToFloat64(v["cash_fee"]) * 100
			cashFee :=  helper.StrconvStringToFloat64(v["cash_fee"]) * 100

			refundResult = libraries.OrderRefund(v["out_trade_no"], OutRefundNo, int64(totalFee), int64(cashFee))

			if refundResult["return_code"] != "SUCCESS" {
				refundStatus = false
				helper.ErrorLog("CancelReceiptStatusModel 微信退款失败。OutTradeNo:"+v["out_trade_no"]+"-refundResult：", refundResult)
				engineSession.Rollback()
			} else if refundResult["result_code"] == "FAIL" {
				if refundResult["err_code_des"] == "订单已全额退款" {
					refundStatus = true
					isHaveRefund = true
				} else {
					refundStatus = false
					isHaveRefund = false
					engineSession.Rollback()
				}

			}
		}

		// 更新数据
		if refundStatus == true {
			var updateJgxSellerOrderRefund JgxSellerOrderRefund
			updateJgxSellerOrderRefund.Status = 1
			updateJgxSellerOrderRefund.PayTime = currentTime
			if isNotFree == true {
				if isHaveRefund == false {
					updateJgxSellerOrderRefund.RefundId = refundResult["refund_id"]
				} else {
					updateJgxSellerOrderRefund.RefundReason = "系统自动处理取消接单(已退款订单)"
				}
			} else {
				updateJgxSellerOrderRefund.RefundReason = "系统自动处理取消接单(免费订单)"
			}


			_, err = engineSession.Where("out_trade_no = ?", v["out_trade_no"]).Update(&updateJgxSellerOrderRefund)
			helper.CheckErr(err, "CancelReceiptStatusModel 更新退款失败")
			helper.DebugLog("CancelReceiptStatusModel 更新退款成功。OutTradeNo:", v["out_trade_no"])
		}

		err = engineSession.Commit()
		helper.CheckErr(err, "CancelReceiptStatusModel engineSession 事务提交发生错误")
	}
}