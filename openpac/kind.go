package openpac

// 登录类型
const (
	ClientTypePc = "pc" // pc端
	ClientTypeH5 = "h5" // h5端
)

// 订单类型
const (
	SFJOrderTypeSubOrder = 1 // 子订单
)

// RemarkType 备注类型
const (
	RemarkTypeCashier = "JHS0100000"
	RemarkTypeRefund  = "JHT0100000"
)

// PayModel 支付模式
const (
	PayModelFreeze = 0 // 冻结支付
	PayModelNormal = 1 // 普通支付
)

// RefundMode 退款模式
const (
	RefundModelFreeze = 0 // 冻结支付
	RefundModelNormal = 1 // 普通支付
)

// 订单类型, 支付前后台通知的订单类型
const (
	NotifyOrderTypePay    = "1" // 支付
	NotifyOrderTypeRefund = "2" // 退款
	NotifyOrderTypeUndo   = "3" // 撤消
)

// 订单状态, 支付前后台通知的订单状态
const (
	NotifyOrderStatusAccepted = "0" // 已受理
	NotifyOrderStatusSuccess  = "1" // 交易成功
	NotifyOrderStatusTrading  = "2" // 交易中
	NotifyOrderStatusPaying   = "3" // 支付中
	NotifyOrderStatusClose    = "4" // 交易关闭
	NotifyOrderStatusUndo     = "5" // 已撤消
)

// 通道类型
const (
	NotifyPayCardTypeDebit = "1" // 借记通道
	NotifyPayCardTypeCargo = "2" // 货记通道
)
