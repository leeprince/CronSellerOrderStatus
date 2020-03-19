package models
/*models 包中定义的结构体*/

type JgxSellerOrder struct {
	Id int `xorm:" int(11) notnull pk autoincr comment('自增ID')"`
	CreateTime int `xorm:"int(11) not null"`
	Status int `xorm:"int(11) not null comment('订单状态。0=>待支付,1=>支付成功(待接单),2=>支付取消，3=>服务中,4=>待确认,5=>已完成')"`
	EndTime int64 `xorm:"int(11) comment('订单完成时间')"`
	Timing int `xorm:"tinyint(1) comment('定时任务0=>未处理过，1=>已处理过')"`
}

type JgxSellerOrderRefund struct {
	Id            int    `xorm:"int(11) notnull pk autoincr comment('自增ID')"`
	Uuid          string `xorm:"varchar(50) notnull comment('uuid')"`
	TransactionId string `xorm:"varchar(64) notnull comment('支付平台订单号')"`
	OutTradeNo    string `xorm:"varchar(32) notnull comment('商户交易单号: p+14位时间+3位毫秒+9位随机数')"`
	RefundId      string `xorm:"varchar(64) comment('退款平台订单号')"`
	OutRefundNo  string `xorm:" varchar(24) notnull comment('商户退款单号: p+14位时间+3位毫秒+9位随机数')"`
	CashFee       string `xorm:"decimal(10,2) notnull comment('支付总金额')"`
    CreateTime    int64 `xorm:"created comment('创建退单时间')"`
	RefundReason  string `xorm:"varchar(255) notnull comment('退款原因')"`
	Mobile        string `xorm:"varchar(20) notnull comment('用户手机号码')"`
	Status        int `xorm:"tinyint(1) notnull comment('退款状态0=>待退款，1=>退款成功，2=>退款失败')"`
	Accept        int `xorm:"tinyint(1) notnull comment('受理状态:0=>未处理;1=>受理; 2=>不受理')"`
	RefundFee       string `xorm:" decimal(10,2) comment('退款金额')"`
	PayType       string `xorm:"tinyint(1) comment('1=>微信小程序')"`
	PayTime       int64 `xorm:"int(11) notnull comment('退款时间')"`
	AcceptTime int64 `xorm:"int(50) notnull comment('处理时间')"`
	ApplyRole int `xorm:"tinyint(1) notnull comment('角色：1=>平台2=>商家,3=>用户,4=>定时器')"`
	Source int `xorm:"tinyint(1) comment('退款申请的前置状态,对应订单表状态1=>支付成功(待接单),3=>服务中,4=>待确认')"`
}