package openpaa

import (
	"context"
)

type ApntTransferReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo                    string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	RecvAcctNo                    string // required(32), 收款账户的账号(见证子账户绑定的银行卡账户)
	RecvAcctName                  string `json:",omitempty"` // optional(120), 收款账户的户名
	RecvAcctOpenBranchName        string `json:",omitempty"` // optional(100), 收款账户的开户行的行名
	RecvAcctOpenBranchInterbankId string `json:",omitempty"` // optional(20), 收款账户的联行号
	ApplyTakeCashAmt              string // required(20), 申请提现的金额(单位: 元)
	MarketChargeCommission        string // required(15), 市场收取的手续费(单位: 元), 未启用, 须送 0.00
	Remark                        string `json:",omitempty"` // optional(120), 备注, 此栏位内容会作为提现交易的附言信息，若未上送则后台自动补全为该笔转账申请的第三方交易流水号
	HoldOne                       string `json:",omitempty"` // optional(120), 预留字段1
	HoldTwo                       string `json:",omitempty"` // optional(120), 预留字段2
	HoldThree                     string `json:",omitempty"` // optional(120), 预留字段3
}

type ApntTransferRsp struct {
	ResponseData
	WitnessSysSeqNo string // 见证系统流水号,电商见证宝系统生成的流水号, 可关联具体一笔请求
	HoldOne         string
	HoldTwo         string
}

// ApntTransfer 指定转账划款
// 从固定的一个账户支取,向指定的账户(输入的收款账户支持绑定卡、监管户或智能收款子账户)中转入指定的金额.
// NOTE: 该接口只能在测试环境使用, 不可投产
func (c *Client) ApntTransfer(ctx context.Context, req *ApntTransferReq) (*ApntTransferRsp, error) {
	return Invoke[ApntTransferReq, ApntTransferRsp](c, ctx, KFEJZB6211, req)
}
