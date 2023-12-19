package openpaa

import (
	"errors"
)

// 接口id/服务id
const (
	KFEJZB6248 = "KFEJZB6248" // 6248 实名开户
	KFEJZB6293 = "KFEJZB6293" // 6293 实名开同名账户关联关系维护
	KFEJZB6092 = "KFEJZB6092" // 6092 根据会员代码查询会员子账号
	KFEJZB6238 = "KFEJZB6238" // 6238 会员绑定提现账户银联鉴权-校验法人
	KFEJZB6239 = "KFEJZB6239" // 6239 银联鉴权回填短信码-校验法人 与6238搭配使用
	KFEJZB6240 = "KFEJZB6240" // 6240 会员绑定提现账户小额鉴权-校验法人
	KFEJZB6241 = "KFEJZB6241" // 6241 小额鉴权回填金额-校验法人 与6240搭配使用
	KFEJZB6065 = "KFEJZB6065" // 6065 会员解绑提现账户
	KFEJZB6083 = "KFEJZB6083" // 6083 申请修改手机号码
	KFEJZB6084 = "KFEJZB6084" // 6084 回填动态码-修改手机
	KFEJZB6138 = "KFEJZB6138" // 6138 维护会员绑定提现账户联行号
	KFEJZB6244 = "KFEJZB6244" // 6244 登记行为记录信息
	KFEJZB6139 = "KFEJZB6139" // 6139 登记挂账
	KFEJZB6140 = "KFEJZB6140" // 6140 登记挂账撤消
	KFEJZB6145 = "KFEJZB6145" // 6145 调账-见证收单
	KFEJZB6147 = "KFEJZB6147" // 6147 平台补账-见证收单
	KFEJZB6171 = "KFEJZB6171" // 6171 受理维护会员信息申请
	KFEJZB6296 = "KFEJZB6296" // 6296 会员信息修改
	// 交易
	KFEJZB6034 = "KFEJZB6034" // 6034 会员间交易 - 不验证
	KFEJZB6033 = "KFEJZB6033" // 6033 会员提现 - 不验证
	KFEJZB6007 = "KFEJZB6007" // 6007 会员资金冻结-不验证
	KFEJZB6216 = "KFEJZB6216" // 6216 第三方支付渠道在途充值(分账)
	KFEJZB6217 = "KFEJZB6217" // 6217 第三方支付渠道在途充值撤销（分账）
	KFEJZB6082 = "KFEJZB6082" // 6082 申请提现或支付短信动态码
	KFEJZB6101 = "KFEJZB6101" // 6101 会员间交易-验证短信动态码
	KFEJZB6164 = "KFEJZB6164" // 6164 会员间交易退款-不验证

	// 查询
	KFEJZB6098 = "KFEJZB6098" // 6098 会员绑定信息查询
	KFEJZB6061 = "KFEJZB6061" // 6061 查询小额鉴权转账结果
	KFEJZB6037 = "KFEJZB6037" // 6037 查询会员子账号
	KFEJZB6093 = "KFEJZB6093" // 6093 查询会员子账号余额
	KFEJZB6010 = "KFEJZB6010" // 6010 查询银行子账户余额
	KFEJZB6110 = "KFEJZB6110" // 6110 查询银行单笔交易状态
	KFEJZB6108 = "KFEJZB6108" // 6108 查询银行在途清算结果
	KFEJZB6109 = "KFEJZB6109" // 6109 查询银行费用扣收结果
	KFEJZB6114 = "KFEJZB6114" // 6114 查询子帐号历史余额及待转可提现状态信息
	KFEJZB6011 = "KFEJZB6011" // 6011 查询资金汇总账户余额
	KFEJZB6048 = "KFEJZB6048" // 6048 查询银行提现退单信息
	KFEJZB6050 = "KFEJZB6050" // 6050 查询普通转账充值明细
	KFEJZB6072 = "KFEJZB6072" // 6072 查询银行时间段内交易明细
	KFEJZB6073 = "KFEJZB6073" // 6073 查询银行时间段内清分提现明细
	KFEJZB6142 = "KFEJZB6142" // 6142 查询明细单验证码
	KFEJZB6103 = "KFEJZB6103" // 6103 查询对账文件信息
	KFEJZB6146 = "KFEJZB6146" // 6146 查询充值明细的交易状态-见证收单
	KFEJZB6324 = "KFEJZB6324" // 6324 见证子台帐信息查询
	// 测试使用
	KFEJZB6211 = "KFEJZB6211" // 6211 指定转账划款
)

// 接口id --> 服务id 映射
var mappingServiceId = map[string]string{
	KFEJZB6248: "AutonymOpenCustAcctId",
	KFEJZB6293: "MntMbrBindSameRealCustNameAcct",
	KFEJZB6092: "QueryCustAcctIdByThirdCustId",
	KFEJZB6238: "BindUnionPayWithCheckCorp",
	KFEJZB6239: "CheckMsgCodeWithCorp",
	KFEJZB6240: "BindSmallAmountWithCheckCorp",
	KFEJZB6241: "CheckAmountWithCorp",
	KFEJZB6065: "UnbindRelateAcct",
	KFEJZB6083: "ApplyForChangeOfCellPhoneNum",
	KFEJZB6084: "BackfillDynamicPassword",
	KFEJZB6138: "MntMbrBindRelateAcctBankCode",
	KFEJZB6244: "RegisterBehaviorRecordInfo",
	KFEJZB6139: "RegisterBillSupportWithdraw",
	KFEJZB6140: "RevRegisterBillSupportWithdraw",
	KFEJZB6145: "AccountRegulation",
	KFEJZB6147: "PlatformAccountSupply",
	KFEJZB6171: "RcvMntMbrInfoAply",
	KFEJZB6296: "MemberInformationChange",
	// 交易
	KFEJZB6034: "MemberTransaction",
	KFEJZB6033: "MembershipWithdrawCash",
	KFEJZB6007: "MembershipTrancheFreeze",
	KFEJZB6216: "OnWayTopThirdPaySplit",
	KFEJZB6217: "RevokeOnWayTopThirdPaySplit",
	KFEJZB6082: "ApplicationTextMsgDynamicCode",
	KFEJZB6101: "MemberTranVerifyTextMsgs",
	KFEJZB6164: "MemberTransactionRefund",

	// 查询
	KFEJZB6098: "MemberBindQuery",
	KFEJZB6061: "SmallAmountTransferQuery",
	KFEJZB6037: "QueryCustAcctId",
	KFEJZB6093: "QueryCustAcctIdBalance",
	KFEJZB6010: "CustAcctIdBalanceQuery",
	KFEJZB6110: "SingleTransactionStatusQuery",
	KFEJZB6108: "BankClearQuery",
	KFEJZB6109: "BankCostDsDealResultQuery",
	KFEJZB6114: "CustAcctIdHistoryBalanceQuery",
	KFEJZB6011: "SupAcctIdBalanceQuery",
	KFEJZB6048: "BankWithdrawCashBackQuery",
	KFEJZB6050: "CommonTransferRechargeQuery",
	KFEJZB6072: "BankTransactionDetailsQuery",
	KFEJZB6073: "BankWithdrawCashDetailsQuery",
	KFEJZB6142: "DetailVerifiedCodeQuery",
	KFEJZB6103: "ReconciliationDocumentQuery",
	KFEJZB6146: "ChargeDetailQuery",
	KFEJZB6211: "ApntTransfer",
	KFEJZB6324: "EJZBCustInformationQuery",
}

// RegisterServiceId registered before it's registered'
func RegisterServiceId(interId, serviceId string) {
	mappingServiceId[interId] = serviceId
}

// GetServiceId 通过接口id获取服务id
func GetServiceId(interId string) (string, error) {
	v, ok := mappingServiceId[interId]
	if ok {
		return v, nil
	}
	return "", errors.New("openpab: interface id mapping service id not existed")
}

// MustGetServiceId 通过接口id获取服务id, 未找到则panic
func MustGetServiceId(interId string) string {
	v, err := GetServiceId(interId)
	if err != nil {
		panic(err)
	}
	return v
}
