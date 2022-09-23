package openpac

import (
	"encoding/json"
	"unsafe"
)

// CashierOderItem CashierOrderRemark 的 OderList
type CashierOderItem struct {
	SubAccNo   string `json:"SubAccNo"`        // required(32), 入账会员子账户
	PayModel   int    `json:"PayModel,string"` // required(1), 支付模式, 0: 冻结支付, 1: 普通支付
	TranFee    string `json:"TranFee"`         // required(13), 手续费, 金额为两位小数格式的字符串
	SubAmount  string `json:"subamount"`       // required(13), 子订单金额(包含手续费), 金额为两位小数格式的字符串
	SuborderId string `json:"suborderId"`      // required(22), 子订单号, 不可超过22位, 且全局唯一,不可跟本订单或其他订单明细重复
	Object     string `json:"Object"`          // optional(500), 子订单信息
}

// CashierOrderRemark 云收款 OrderRemark 字段结构
type CashierOrderRemark struct {
	SFJOrderType int               `json:"SFJOrdertype,string"` // required(1), 订单类型, 取值: 1: 子订单
	RemarkType   string            `json:"remarktype"`          // required(10), 备注类型, 取值: JHS0100000
	PlantCode    string            `json:"plantCode"`           // required(4), 平台代码, 即见证宝的平台代码
	OderList     []CashierOderItem `json:"oderlist"`            // 子订单信息
}

func CashierOrderRemarkString(cor CashierOrderRemark) string {
	b, _ := json.Marshal([]CashierOrderRemark{cor})
	return *(*string)(unsafe.Pointer(&b))
}

// Cashier 云收款支付
type Cashier struct {
	// TraderNo      string     // required(20), 商户编号(云收款商户号), 商户在云收款系统中的编号
	ClientNo      string // required(50), 客户号, 商户自己客户体系唯一标识(交易网中付款方唯一识别码)
	TraderOrderNo string // required(20), 商户系统生成的订单号(总订单号)
	TranAmt       string // required(18), 交易金额(子订单总额), 整数, 单位: 分.
	OrderSendTime string // required(14), 商户系统订单生成时间, 格式: yyyyMMddHHmmss
	OrderName     string // required(50), 订单名称
	OrderRemark   string // required(2500), 订单备注,见证空分商户必须上送, OrderRemark 的数组字符串
	ClientType    string // required(50), 登录类型, pc: pc端, h5: 移动端

	FrontSkipUrl      string // required(100), 前端跳转url
	CallBackNoticeUrl string // required(100), 回调通知url
	// Signature         string // required(2000), 使用商户私钥对报文数据签名(底层已实现)
}

// PayNotify 支付前后台通知
type PayNotify struct {
	TraderNo        string // required(18), 商户编号(云收款商户号), 商户在云收款系统的编号
	TraderOrderNo   string // required(32), 商户系统生成的订单号(总订单号)
	BankOrderNo     string // required(32), 银行订单号
	PayModeNo       string // required(42), 支付方式编号, 可选值为 B2B,B2C,WAP,OnlineGateway
	Fee             string // TODO: 未知
	StoreId         string // optional(18), 门店编号
	CashierNo       string // optional(18), 收银员编号
	TranAmt         string // required(17), 交易金额, 整数, 单位: 分
	OrderType       string // required(2), 订单类型, 1: 支付, 2: 退款, 3: 撤销
	OrderStatus     string // required(2), 订单状态, 0: 已受理, 1: 交易成功, 2: 交易中, 3: 用户支付中, 4: 交易关闭, 9: 已撤销
	OrderSendTime   string // required(14), 订单发送时间, 格式: yyyyMMddHHmmss
	PaySuccessTime  string // optional(14), 支付成功时间, 格式: yyyyMMddHHmmss
	ChannelOrderNo  string // optional(18), 通道流水号
	PayerInfo       string // optional(50), 支付账户信息
	PayCardType     string // optional(2), 通道类型, 1: 借记通道, 2: 贷记通道, 支付方式为B2C时必填
	DimensionalCode string // optional(50), 二维码
	Signature       string // required(2000), 签名, 使用平安银行公钥对报文数据签名验证
}
