package libraries

import (
	"CronSellerOrderStatus/src/admin/helper"
	"github.com/objcoding/wxpay"
)

// 初始化微信配置
func initClientConfig() *wxpay.Client {
	// 创建支付账户
	account := wxpay.NewAccount(
		"wx7b2702cb7237f697",
		"1494219842",
		"e9483bcfa2cfa9414c27a332316f964e",
		false) // sandbox环境请传true

	// 设置证书
	account.SetCertData("./weixin_cert/apiclient_cert.p12")

	// 新建微信支付客户端
	client := wxpay.NewClient(account)

	// 设置http请求超时时间
	client.SetHttpConnectTimeoutMs(2000)

	// 设置http读取信息流超时时间
	client.SetHttpReadTimeoutMs(1000)

	// 更改签名类型
	//client.SetSignType(HMACSHA256)

	// 设置支付账户
	//client.setAccount(account)

	return client
}

// 订单查询
// map[sign:C8486C0F99DD2181F5A9D5126CD720B6 result_code:SUCCESS bank_type:CFT trade_state:SUCCESS trade_state_desc:支付成功 return_code:SUCCESS nonce_str:7Zwq2QT2HuKVhC3l openid:ogLYF0YeJrnLvtLoENC7DLULw_9Q trade_type:JSAPI fee_type:CNY time_end:20181122135108 cash_fee:1 out_trade_no:S20181122135053629525730145 attach:云享云技术提供 return_msg:OK appid:wx7b2702cb7237f697 mch_id:1494219842 is_subscribe:N total_fee:1 transaction_id:4200000214201811220888219487]
func OrderQuery(out_trade_no string) wxpay.Params {
	client := initClientConfig()

	params := make(wxpay.Params)
	params.SetString("out_trade_no", out_trade_no) // S20181121140004859429926672
	queryResult, err := client.OrderQuery(params)
	helper.CheckErr(err, "OrderQuery 发生错误！")
	return queryResult
}

// 退款
func OrderRefund(out_trade_no string, out_refund_no string, total_fee int64, refund_fee int64) wxpay.Params {
	client := initClientConfig()

	params := make(wxpay.Params)
	params.SetString("out_trade_no", out_trade_no).SetString("out_refund_no", out_refund_no).SetInt64("total_fee", total_fee).SetInt64("refund_fee", refund_fee)
	refundResult, err := client.Refund(params)
	helper.CheckErr(err, "OrderRefund 发生错误！")
	return refundResult
}

// 退款查询
func OrderRefundQuery(out_refund_no string) wxpay.Params {
	client := initClientConfig()

	params := make(wxpay.Params)
	params.SetString("out_refund_no", out_refund_no)
	queryRefundResult, err := client.RefundQuery(params)
	helper.CheckErr(err, "OrderRefundQuery 发生错误！")
	return queryRefundResult
}


