package openpac

import (
	"encoding/json"
	"unsafe"
)

type PayReq struct {
	//	TraderNo          string     `json:"TraderNo"`                    // required(32), 商户在云收款系统的编号(底层已实现)
	TraderOrderNo     string `json:"TraderOrderNo"`               // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	StoreId           string `json:"StoreId"`                     // optional(32), 商户在云收款系统的门店编号
	OrderSendTime     string `json:"OrderSendTime"`               // required(20), 商户系统订单生成时间 格式：yyyyMMddHHmmss
	OrderName         string `json:"OrderName,omitempty"`         // optional(60), 订单名称
	CashierNo         string `json:"CashierNo,omitempty"`         // optional(20), 收银员在云收款系统的编号
	PayModeNo         string `json:"PayModeNo"`                   // required(50), 云收款支付方式编号
	Scene             string `json:"Scene,omitempty"`             // Optional(20), 场景
	AuthCode          string `json:"AuthCode,omitempty"`          // optional(6), 授权码
	TranAmt           string `json:"TranAmt"`                     // required(20), 交易金额（单位：分）
	CallBackNoticeURL string `json:"CallBackNoticeUrl,omitempty"` // optional(200), 回调通知url
	FrontSkipURL      string `json:"FrontSkipUrl"`                // optional(200), 前端跳转url
	OrderRemark       string `json:"OrderRemark,omitempty"`       // optional(300), 订单备注,见证宝商户必填
	PayLimit          string `json:"PayLimit,omitempty"`          // optional(20), limit_pay=no_credit，微信支付限制用户不能使用信用卡支付
	MessageCheckCode  string `json:"MessageCheckCode,omitempty"`  // optional(20), 固定6位数字,支付方式为FastPay、DirectPay、AgreePay时必填
	BindCardNo        string `json:"BindCardNo,omitempty"`        // optional(50), 商户在云收款系统的绑卡编号, 支付方式为FastPay时必填，需与发送短信接口绑卡编号一致
	ClientNo          string `json:"ClientNo,omitempty"`          // optional(32), 客户号
	// TODO:
}

type PayRsp struct {
	// TODO
}

type QuerySingleOrderReq struct {
	// TraderNo      string     // required(32), 商户在云收款系统的编号(底层已实现)
	TraderOrderNo string // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	OrderSendTime string // required(20), 订单发送时间
}

type QuerySingleOrderRsp struct {
	ResponseData   `json:",squash"`
	TraderNo       string `json:"TraderNo"`           // required(32), 商户在云收款系统的编号
	TraderOrderNo  string `json:"TraderOrderNo"`      // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	BankOrderNo    string `json:"BankOrderNo"`        // required(50), 云收款系统订单号
	PayModeNo      string `json:"PayModeNo"`          // required(50), 云收款支付方式编号
	StoreId        string `json:"StoreId"`            // optional(32), 商户在云收款系统的门店编号
	CashierNo      string `json:"CashierNo"`          // optional(20), 收银员在云收款系统的编号
	TranAmt        string `json:"TranAmt"`            // required(20), 订单金额（单位：分）
	OrderType      int    `json:"OrderType,string"`   // required(3), 订单类型, 1支付 2退款 3撤销
	OrderStatus    int    `json:"OrderStatus,string"` // required(20), 订单状态, 0 已受理;1 交易成功 ;2 交易中; 3 用户支付中;  4 交易关闭; 9 已撤销
	OrderSendTime  string `json:"OrderSendTime"`      // required(20), 商户系统订单生成时间 格式：yyyyMMddHHmmss
	PaySuccessTime string `json:"PaySuccessTime"`     // optional(20), 云收款系统订单支付成功时间，格式：yyyyMMddHHmmss
	ChannelOrderNo string `json:"ChannelOrderNo"`     // optional(50), 上游通道返回的订单号
	PayerInfo      string `json:"PayerInfo"`          // optional(100), 付款方信息, 支付账户信息
	PayCardType    int    `json:"PayCardType,string"` // optional(20), 1：借记卡/储蓄卡2：贷记卡/信用卡

	Fee            string `json:"Fee"`            // required(50), 未知使用
	SettleProperty string `json:"SettleProperty"` // required(50), 未知使用
}

type RevokeOrderReq struct {
	// TraderNo         string     `json:"TraderNo"`               // required(32), 商户在云收款系统的编号(底层已实现)
	OldMerOrderNo    string `json:"OldMerOrderNo"`          // required(32), 支付接口商户上送的订单号
	OldOrderSendTime string `json:"OldOrderSendTime"`       // required(20), 支付接口商户上送的发送时间，格式：yyyyMMddHHmmss
	CashierNo        string `json:"CashierNo,omitempty"`    // optional(20), 收银员在云收款系统的编号
	CancelRemark     string `json:"CancelRemark,omitempty"` // optional(300), 撤销备注
}

type RevokeOrderRsp struct {
	ResponseData      `json:",squash"`
	CancelOrderStatus int `json:"CancelOrderStatus"` // required(20), 撤销订单状态, 1:成功 2:处理中 4:失败
}

type ApplyRefundOderItem struct {
	SubAccNo      string `json:"SubAccNo"`           // required(32), 入账会员子账户
	RefundModel   int    `json:"refundModel,string"` // required(1), 退款模式, 0: 冻结支付, 1: 普通支付
	RefundTranFee string `json:"RefundTranFee"`      // required(13), 退款手续费, 金额为两位小数格式的字符串
	SubRefundAmt  string `json:"subrefundamt"`       // required(13), 子订单退款金额, 金额为两位小数格式
	SubOrderID    string `json:"suborderId"`         // required(22), 子订单号, 原支付时的子订单好，不可超过22位，且全局唯一，不可跟本订单或其他订单明细重复
	SubRefundId   string `json:"subrefundId"`        // required(22), 子订单退款号, 不可超过22位，且全局唯一，不可跟本订单或其他订单明细重复
	Object        string `json:"object,omitempty"`   // optional(500)子订单信息
}

type ApplyRefundRemark struct {
	SFJOrderType int                   `json:"SFJOrdertype,string"` // required(1), 订单类型, 取值: 1: 子订单
	RemarkType   string                `json:"remarktype"`          // required(10), 备注类型, 取值: JHT0100000
	PlantCode    string                `json:"plantCode"`           // required(4), 平台代码, 即见证宝的平台代码
	OderList     []ApplyRefundOderItem `json:"oderlist"`            // 子订单信息
}

func ApplyRefundRemarkString(arp ApplyRefundRemark) string {
	b, _ := json.Marshal([]ApplyRefundRemark{arp})
	return *(*string)(unsafe.Pointer(&b))
}

type ApplyRefundReq struct {
	// TraderNo            string     `json:"TraderNo"`            // required(32), 商户在云收款系统的编号(底层已实现)
	ReturnOrderNo       string `json:"ReturnOrderNo"`       // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	OldMerOrderNo       string `json:"OldMerOrderNo"`       // required(32), 支付接口商户上送的订单号
	OldOrderSendTime    string `json:"OldOrderSendTime"`    // required(20), 支付接口商户上送的发送时间，格式：yyyyMMddHHmmss
	ReturnOrderSendTime string `json:"ReturnOrderSendTime"` // required(20), 商户系统订单生成时间 格式：yyyyMMddHHmmss
	CashierNo           string `json:"CashierNo,omitempty"` // optional(20), 收银员在云收款系统的编号
	ReturnAmt           string `json:"ReturnAmt"`           // required(20), 退款金额,单位: 分
	RefundRemark        string `json:"RefundRemark"`        // optional(300), 退款备注, 见证空分商户必须上送
}

type ApplyRefundRsp struct {
	ResponseData      `json:",squash"`
	ReturnOrderNo     string `json:"ReturnOrderNo"`     // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	BankOrderNo       string `json:"BankOrderNo"`       // required(50), 云收款系统订单号
	ChannelOrderNo    string `json:"ChannelOrderNo"`    // optional(50), 上游通道返回的订单号
	ReturnOrderStatus int    `json:"ReturnOrderStatus"` // required(20), 退款订单状态, 1:成功 2:处理中 4:失败
	CreateDate        string `json:"CreateDate"`        // required(20), 创建时间
}

type QuerySingleRefundReq struct {
	// TraderNo            string     `json:"TraderNo"`            // required(32), 商户在云收款系统的编号(底层已实现)
	ReturnOrderNo       string `json:"ReturnOrderNo"`       // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	ReturnOrderSendTime string `json:"ReturnOrderSendTime"` // required(20), 商户系统订单生成时间 格式：yyyyMMddHHmmss
}

type QuerySingleRefundRsp struct {
	ResponseData      `json:",squash"`
	ReturnOrderNo     string `json:"ReturnOrderNo"`     // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	BankOrderNo       string `json:"BankOrderNo"`       // required(50), 云收款系统订单号
	ReturnAmt         string `json:"ReturnAmt"`         // required(20), 退款金额,单位为分
	ReturnOrderStatus int    `json:"ReturnOrderStatus"` // required(20), 退款订单状态, 0 已受理;1 交易成功 ;2 交易中; 3 用户支付中;  4 交易关闭; 9 已撤销
}

type SendShortMessageReq struct {
	// TraderNo      string     `json:"TraderNo"`           // required(32), 商户在云收款系统的编号(底层已实现)
	TraderOrderNo string `json:"TraderOrderNo"`      // required(100), 商户系统生成的订单号，需要保证在商户系统唯一
	OrderSendTime string `json:"OrderSendTime"`      // required(20), 商户系统订单生成时间 格式：yyyyMMddHHmmss
	TranAmt       string `json:"TranAmt"`            // required(20), 交易金额（单位：分）
	BindCardNo    string `json:"BindCardNo"`         // required(50), 云收款系统的绑卡编号
	ClientNo      string `json:"ClientNo,omitempty"` // required(32), 客户号, 商户自己客户体系唯一标识
}

type SendShortMessageRsp struct {
	ResponseData `json:",squash"`
	Status       int `json:"Status"` // 状态, 0: 失败, 1: 成功
}

type DownloadReconciliationFileReq struct {
	// TraderNo      string     `json:"TraderNo"`      // required(32), 商户在云收款系统的编号(底层已实现)
	ReconcileDate string `json:"ReconcileDate"` // required(8), 对账日期, 格式：yyyyMMdd
}

type DownloadReconciliationFileRsp struct {
	ResponseData     `json:",squash"`
	CreateBillStatus int // required(20), 生成对账单的状态, 0初始状态,1对账成功,2,对账中,4对账失败
	// 以下当 CreateBillStatus=1时有值
	IncomeCount int64  // optional(20), 收入笔数
	IncomeAmt   string // optional(20), 收入金额
	ReturnNum   int64  // optional(20), 退款笔数
	ReturnAmt   string // optional(20), 退款金额
	CancelNum   int64  // optional(20), 撤销笔数
	CancelAmt   string // optional(20), 撤销金额
	DrawCode    string // optional(64), 提取码
	EncryptFlag int    // optional(1), 加密标志
	UnzipCode   string // optional(64), 解压码
	FilePath    string // optional(512), 文件路径
}
