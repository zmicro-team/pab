package openpaa

// 功能标志
const (
	// 6248 开户 - 功能标志
	FFAutonymOpenAccountOpen  = "1" // 开户
	FFAutonymOpenAccountClose = "2" // 销户
	FFAutonymOpenAccountStock = "3" // 存量实名

	// 6293 实名开同名账户关联关系维护的 - 功能标志
	FFMntMbrBindAccountSame     = "1" // 申请开立商户子账户和普通会员子账户
	FFMntMbrBindAccountCommon   = "2" // 申请为已有的商户子账户开立关联的同名的普通会员子账户
	FFMntMbrBindAccountMer      = "3" // 申请为已有的普通会员子账户开立关联的同名的商户子账户
	FFMntMbrBindAccountQuery    = "4" // 查询两个会员子账户是否存在同名账户关系
	FFMntMbrBindAccountMaintain = "5" // 维护两个子台帐信息一致的SH和00户为同名户关系

	// 6065 解绑 - 功能标志
	FFUnbindAcctFunction = "1"

	// 6034 会员间交易的 - 功能标志
	// 6101 会员间交易-验证短信动态码
	FFMemberTranPrepay      = "1" // 下单预支付
	FFMemberTranConfirmPay  = "2" // 确认并付款
	FFMemberTranRefund      = "3" // 退款(只能退预支付)
	FFMemberTranSameAccount = "4" // 同名子账户支付 (商户A -> 商户B)
	FFMemberTranDirectPayT1 = "6" // 直接支付T+1
	FFMemberTranToPlatform  = "7" // 支付到平台
	FFMemberTranDirectPayT0 = "9" // 直接支付T+0

	// 6164 会员间交易退款-不验证
	FFMemberTranRefundConfirmPay      = "1" // 1：会员交易接口 [确认付款] 退款, 针对 [6006/6034/6101] funcflag=2-确认并付款的退款
	FFMemberTranRefundDirectPay       = "2" // 2：会员交易接口 [直接支付] 退款, 针对 [6006/6034/6101] funcflag=6、9-直接支付的退款
	FFMemberTranRefundAgentConfirm    = "3" // 3：平台订单管理接口 [平台代理确认收货] 退款, 针对 [6031] 2-平台代理确认收货的退款
	FFMemberTranRefundBatchConfirmPay = "4" // 4：会员批量交易接口 [批量确认] 退款, 针对 [6052/6120/6133] 2-批量确认的退款
	FFMemberTranRefundBatchDirectPay  = "5" // 5：会员批量交易接口 [直接支付] 退款, 针对 [6052/6120/6133] 3-直接支付的退款. 未提供使用.
	FFMemberTranRefundFund            = "6" // 6：会员资金支付接口的退款, 针对 [6163/6165/6166] 会员资金支付的退款
	FFMemberTranRefundTrade           = "7" // 7：会员交易（多借多贷）接口的退款, 针对 [6119] 6、9-直接支付的退款. 6119是专用接口，不提供开放使用.
	FFMemberTranRefundX               = "X" // X：会员交易接口 [资方交易] 退款, 针对 [6034] X-资方交易
	FFMemberTranRefundH               = "H" // H: [6034] 同名子账户支付（商户子账户A→普通会员A→商户子账户B）或 （商户子账户A→商户子账户B）

	// 6007 会员资金冻结 - 功能标志
	FFMembershipTrancheFreeze                = "1" // 冻结 (会员→担保)
	FFMembershipTrancheUnfreeze              = "2" // 解冻 (担保→会员)
	FFMembershipTrancheJZBUnfreeze           = "4" // 见证+收单的冻结资金解冻
	FFMembershipTrancheWithdrawFreeze        = "5" // 可提现冻结 (会员→担保)
	FFMembershipTrancheWithdrawUnfreeze      = "6" // 可提现解冻 (担保→会员)
	FFMembershipTrancheOnWayRechargeUnfreeze = "7" // 在途充值解冻 担保→会员

	// 6164 会员间交易退款-不验证 - 功能标志
	FFMemberTransactionRefundDirectPay = "2" // 会员交易接口[直接支付]退款

	//
	FFRegisterBehaviorRecord = "1"
	FFRegisterBehaviorQuery  = "2"

	// 6110 查询银行单笔交易状态 - 功能标志
	FFSingleTransactionStatusMemberShip = "2" // 会员间交易
	FFSingleTransactionStatusWithdraw   = "3" // 提现
	FFSingleTransactionStatusRecharge   = "4" // 充值

	// 6108 查询银行在途清算结果 - 功能标志
	FFBankClearQueryAll       = "1" // 全部
	FFBankClearQuerySpecified = "2" // 指定时间段

	// 6114 查询子帐号历史余额及待转可提现状态信息 - 功能标志
	FFCustAcctIdHistoryBalanceQueryAll       = "1" // 全部
	FFCustAcctIdHistoryBalanceQuerySpecified = "2" // 指定时间段
)

const (
	// 结束标志
	EndFlagNo  = "0" // 否
	EndFlagYes = "1" // 是

	// 查询标志
	QueryFlagAll    = "1"
	QueryFlagSingle = "2"
)
