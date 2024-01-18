package openpab

import (
	"context"
)

// AutonymOpenCustAcctId 子账户实名开立, 返回见证保子账户的账号
// NOTE: 重复开户(交易网会员代码已存), 返回(Err114)错误
func (c *Client) AutonymOpenCustAcctId(ctx context.Context, req *AutonymOpenCustAcctIdReq) (*AutonymOpenCustAcctIdRsp, error) {
	return Invoke2[AutonymOpenCustAcctIdReq, AutonymOpenCustAcctIdRsp](c, ctx, KFEJZB6248, req)
}

// MntMbrBindSameRealCustNameAcct 实名开同名账户关联关系维护
// 1：可用于6248接口实名开同名账户关联关系维护:
// 2：非自营市场可通过本接口
//
//	分支1. 申请开立商户子账户和普通会员子账户
//	分支2. 申请为已有的商户子账户开立关联的同名的普通会员子账户
//	分支3. 申请为已有的普通会员子账户开立关联的同名的商户子账户
//	分支4. 查询两个会员子账户是否存在同名账户关系
//	分支5. 维护两个子台帐信息一致的SH和00户为同名户关系
func (c *Client) MntMbrBindSameRealCustNameAcct(ctx context.Context, req *MntMbrBindSameRealCustNameAcctReq) (*MntMbrBindSameRealCustNameAcctRsp, error) {
	return Invoke2[MntMbrBindSameRealCustNameAcctReq, MntMbrBindSameRealCustNameAcctRsp](c, ctx, KFEJZB6293, req)
}

// QueryCustAcctIdByThirdCustId 根据会员代码查询会员子账号
// 返回见证宝子账户的账号
func (c *Client) QueryCustAcctIdByThirdCustId(ctx context.Context, req *QueryCustAcctIdByThirdCustIdReq) (*QueryCustAcctIdByThirdCustIdRsp, error) {
	return Invoke2[QueryCustAcctIdByThirdCustIdReq, QueryCustAcctIdByThirdCustIdRsp](c, ctx, KFEJZB6092, req)
}

// BindUnionPayWithCheckCorp 会员绑定提现账户银联鉴权-校验法人
// 此鉴权用户: 1、纯个人 2、个体工商户
func (c *Client) BindUnionPayWithCheckCorp(ctx context.Context, req *BindUnionPayWithCheckCorpReq) (*ResponseReserved, error) {
	return Invoke2[BindUnionPayWithCheckCorpReq, ResponseReserved](c, ctx, KFEJZB6238, req)
}

// CheckMsgCodeWithCorp 银联鉴权回填短信码-校验法人
// 与 BindUnionPayWithCheckCorp 配合使用
func (c *Client) CheckMsgCodeWithCorp(ctx context.Context, req *CheckMsgCodeWithCorpReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[CheckMsgCodeWithCorpReq, ResponseFrontSeqNo](c, ctx, KFEJZB6239, req)
}

// BindSmallAmountWithCheckCorp 会员绑定提现账户小额鉴权-校验法人
// 该接口发起成功后，银行会向提现账户转入小于等于0.5元的随机金额，并短信通知客户查看，客户查看后，需将收到的金额大小，
// 在电商平台页面上回填，并通知银行。银行验证通过后，完成提现账户绑定。
//   - 企业对公账户只能使用小额鉴权绑卡
//   - 企业与个体工商户增加校验工商五要素
//   - 五要素校验: 姓名,证件,卡号,银行预留手机以及校验个体工商户
//
// NOTE:
//  1. 大小额联行号和超级网银号的上送说明
//     - 本行账户（平安银行）：大小额行号和超级网银号都不用送
//     - 他行账户, 个人可以只送超级网银号, 企业则需填大小额行号
//  2. 随机金额回填金额有效期48小时
//  3. 鉴权发起次数限制每24小时只能发起一次
//  4. 鉴权发起成功后,如收不到短信,可以重新调用此接口触发短信(2分钟后).每天最多5次
func (c *Client) BindSmallAmountWithCheckCorp(ctx context.Context, req *BindSmallAmountWithCheckCorpReq) (*ResponseReserved, error) {
	return Invoke2[BindSmallAmountWithCheckCorpReq, ResponseReserved](c, ctx, KFEJZB6240, req)
}

// CheckAmountWithCorp 小额鉴权回填金额-校验法人
func (c *Client) CheckAmountWithCorp(ctx context.Context, req *CheckAmountWithCorpReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[CheckAmountWithCorpReq, ResponseFrontSeqNo](c, ctx, KFEJZB6241, req)
}

// UnbindRelateAcct 会员解绑提现账户
func (c *Client) UnbindRelateAcct(ctx context.Context, req *UnbindRelateAcctReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[UnbindRelateAcctReq, ResponseFrontSeqNo](c, ctx, KFEJZB6065, req)
}

// ApplyForChangeOfCellPhoneNum 申请修改手机号码
// 1: 短信验证码方式,适合原手机仍可以接收到短信
// 2: 银联鉴权,适合原手机不能收到短信, 提供卡号,并且手机号为银联预留手机号(这种方式只支持个人卡的修改)
func (c *Client) ApplyForChangeOfCellPhoneNum(ctx context.Context, req *ApplyForChangeOfCellPhoneNumReq) (*ApplyForChangeOfCellPhoneNumRsp, error) {
	return Invoke2[ApplyForChangeOfCellPhoneNumReq, ApplyForChangeOfCellPhoneNumRsp](c, ctx, KFEJZB6083, req)
}

// BackfillDynamicPassword 回填动态码-修改手机
func (c *Client) BackfillDynamicPassword(ctx context.Context, req *BackfillDynamicPasswordReq) (*ResponseReserved, error) {
	return Invoke2[BackfillDynamicPasswordReq, ResponseReserved](c, ctx, KFEJZB6084, req)
}

// MntMbrBindRelateAcctBankCode 维护会员绑定提现账户联行号
// 支持市场修改会员的提现账户的开户行信息.
// 具体包括开户行行名、开户行的银行联行号（大小额联行号）和超级网银行号。
func (c *Client) MntMbrBindRelateAcctBankCode(ctx context.Context, req *MntMbrBindRelateAcctBankCodeReq) (*ResponseReserved, error) {
	return Invoke2[MntMbrBindRelateAcctBankCodeReq, ResponseReserved](c, ctx, KFEJZB6138, req)
}

// RegisterBehaviorRecordInfo 登记行为记录信息
// 功能分支1：用于签订协议上送用户行为信息。无输出反馈，仅会实时反馈记录成功、失败情况。
// 功能分支2：用于SH商户子账户查询是否完成补录，其中用新鉴权接口鉴权的SH商户子账户也属于补录成功的子账户。
func (c *Client) RegisterBehaviorRecordInfo(ctx context.Context, req *RegisterBehaviorRecordInfoReq) (*RegisterBehaviorRecordInfoRsp, error) {
	return Invoke2[RegisterBehaviorRecordInfoReq, RegisterBehaviorRecordInfoRsp](c, ctx, KFEJZB6244, req)
}

// RegisterBillSupportWithdraw 登记挂账(支持撤销)
// 可实现把不明来账或自有资金等已登记在挂账子账户下的资金调整到普通会员子账户。
// 即通过申请调用此接口，将会减少挂账子账户的资金，调增指定的普通会员子账户的可提现余额及可用余额。
// 此接口不支持把挂账子账户资金清分到功能子账户（营销子账户除外）
func (c *Client) RegisterBillSupportWithdraw(ctx context.Context, req *RegisterBillSupportWithdrawReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[RegisterBillSupportWithdrawReq, ResponseFrontSeqNo](c, ctx, KFEJZB6139, req)
}

// RevRegisterBillSupportWithdraw 登记挂账撤销
// 可以实现把原6139指令完成的登记挂账进行撤销，
// 即调减普通会员子账户的可提现和可用余额，调增挂账子账户的可用余额。
func (c *Client) RevRegisterBillSupportWithdraw(ctx context.Context, req *RevRegisterBillSupportWithdrawReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[RevRegisterBillSupportWithdrawReq, ResponseFrontSeqNo](c, ctx, KFEJZB6140, req)
}

// RcvMntMbrInfoApply 受理维护会员信息申请
func (c *Client) RcvMntMbrInfoApply(ctx context.Context, req *RcvMntMbrInfoApplyReq) (*ResponseReserved, error) {
	return Invoke2[RcvMntMbrInfoApplyReq, ResponseReserved](c, ctx, KFEJZB6171, req)
}

// MemberInformationChange 会员信息修改
// 企业用户支持修改会员名称、公司名称、法人名称、法人证件类型、法人证件号码,个人用户只支持修改会员名称,
// 接口暂不支持个体工商户,若企业同名户为个体工商户,不解绑提现账户,不更新同名户信息,若个人同名户为个体工商户,解绑提现账户,不更新同名户信息
func (c *Client) MemberInformationChange(ctx context.Context, req *MemberInformationChangeReq) (*ResponseReserved, error) {
	return Invoke2[MemberInformationChangeReq, ResponseReserved](c, ctx, KFEJZB6296, req)
}

/************************************* 交易 ************************************/

// AccountRegulation 调账-见证收单
// 见证+收单模式，若前一日存在子账户记账失败的情况(如平台上送子钱包信息错误)，
// 对应清算资金挂在快收专属子账户，
// 平台核实资金归属后，调用该接口，将资金清分到具体的会员子账户。(需要关联原资金流水)
func (c *Client) AccountRegulation(ctx context.Context, req *AccountRegulationReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[AccountRegulationReq, ResponseFrontSeqNo](c, ctx, KFEJZB6145, req)
}

// PlatformAccountSupply 平台补账-见证收单
func (c *Client) PlatformAccountSupply(ctx context.Context, req *PlatformAccountSupplyReq) (*PlatformAccountSupplyRsp, error) {
	return Invoke2[PlatformAccountSupplyReq, PlatformAccountSupplyRsp](c, ctx, KFEJZB6147, req)
}

// MemberTransaction 会员间交易 - 不验证
// 实现会员间的余额的交易, 实现资金在会员之间流动
func (c *Client) MemberTransaction(ctx context.Context, req *MemberTransactionReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[MemberTransactionReq, ResponseFrontSeqNo](c, ctx, KFEJZB6034, req)
}

// MembershipWithdrawCash 会员提现 - 不验证
// 此接口受理会员发起的提现申请。会员子账户的可提现余额、可用余额会减少，
// 市场的资金汇总账户(监管账户)会减少相应的发生金额，提现到会员申请的收款账户。
func (c *Client) MembershipWithdrawCash(ctx context.Context, req *MembershipWithdrawCashReq) (*MembershipWithdrawCashRsp, error) {
	return Invoke2[MembershipWithdrawCashReq, MembershipWithdrawCashRsp](c, ctx, KFEJZB6033, req)
}

// MembershipTrancheFreeze 会员资金冻结-不验证
// 可以实现减少会员的余额，增加市场的担保子账户的余额。
// 资金从会员子账户划转到担保子账户后，可用于支付到其他的会员子账户或被解冻退回。
// 跟 6034 会员间交易”相比，进行会员资金冻结的操作，不须要明确最终的收款方信息。
func (c *Client) MembershipTrancheFreeze(ctx context.Context, req *MembershipTrancheFreezeReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[MembershipTrancheFreezeReq, ResponseFrontSeqNo](c, ctx, KFEJZB6007, req)
}

// OnWayTopThirdPaySplit 第三方支付渠道在途充值(分账)
// 市场通过第三方支付渠道(如支付宝、微信等)进行成功扣款后，
// 可通过此接口通知电商见证宝系统为指定的会员增加可用余额, 并收取指定金额的手续费(如有),可支持多笔充值.
func (c *Client) OnWayTopThirdPaySplit(ctx context.Context, req *OnWayTopThirdPaySplitReq) (*OnWayTopThirdPaySplitRsp, error) {
	return Invoke2[OnWayTopThirdPaySplitReq, OnWayTopThirdPaySplitRsp](c, ctx, KFEJZB6216, req)
}

// RevokeOnWayTopThirdPaySplit 第三方支付渠道在途充值撤消(分账)
// 市场若需要对经 6216【会员在途充值(经第三方支付渠道)】的指令进行撤销，可通过本接口实现。支持多笔撤销。
func (c *Client) RevokeOnWayTopThirdPaySplit(ctx context.Context, req *RevokeOnWayTopThirdPaySplitReq) (*RevokeOnWayTopThirdPaySplitRsp, error) {
	return Invoke2[RevokeOnWayTopThirdPaySplitReq, RevokeOnWayTopThirdPaySplitRsp](c, ctx, KFEJZB6217, req)
}

// ApplicationTextMsgDynamicCode 申请提现或支付短信动态码
// 用于有需要进行短信动态码验证的平台，申请短信动态验证码，以便于 6085,6101 等接口的交易进行短信动态码验证。
func (c *Client) ApplicationTextMsgDynamicCode(ctx context.Context, req *ApplicationTextMsgDynamicCodeReq) (*ApplicationTextMsgDynamicCodeRsp, error) {
	return Invoke2[ApplicationTextMsgDynamicCodeReq, ApplicationTextMsgDynamicCodeRsp](c, ctx, KFEJZB6082, req)
}

// MemberTranVerifyTextMsgs 会员间交易-验证短信动态码
// 可以实现会员间的余额的交易，实现资金在会员之间流动。此接口须通过短信验证码校验
// NOTE: 目前仅使用于营销子账号和商户子账号之间的转账
func (c *Client) MemberTranVerifyTextMsgs(ctx context.Context, req *MemberTranVerifyTextMsgsReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[MemberTranVerifyTextMsgsReq, ResponseFrontSeqNo](c, ctx, KFEJZB6101, req)
}

// MemberTransactionRefund 会员间交易退款-不验证
// 用于会员间交易做售后退款处理。从卖方子账户扣退款金额（不含退款手续费），如退手续费金额大于0，
// 则从退款垫款户（默认为平台自有一般账户）扣该退手续费金额，支付给买方子账户可用金额（T+1日变为可提现）。
func (c *Client) MemberTransactionRefund(ctx context.Context, req *MemberTransactionRefundReq) (*ResponseFrontSeqNo, error) {
	return Invoke2[MemberTransactionRefundReq, ResponseFrontSeqNo](c, ctx, KFEJZB6164, req)
}

/************************************* 查询 ************************************/

// MemberBindQuery 会员绑定信息查询
// 查询标志为“单个会员”的情况下，返回该会员的有效的绑定账户信息。
// 查询标志为“全部会员”的情况下，返回市场下的全部的有效的绑定账户信息。
// 查询标志为“单个会员的证件信息”的情况下，返回市场下的指定的会员的留存在电商见证宝系统的证件信息。
func (c *Client) MemberBindQuery(ctx context.Context, req *MemberBindQueryReq) (*MemberBindQueryRsp, error) {
	return Invoke2[MemberBindQueryReq, MemberBindQueryRsp](c, ctx, KFEJZB6098, req)
}

// SmallAmountTransferQuery 查询小额鉴权转账结果
func (c *Client) SmallAmountTransferQuery(ctx context.Context, req *SmallAmountTransferQueryReq) (*SmallAmountTransferQueryRsp, error) {
	return Invoke2[SmallAmountTransferQueryReq, SmallAmountTransferQueryRsp](c, ctx, KFEJZB6061, req)
}

// QueryCustAcctId 查询会员子账号
// 显示可提现余额,可用余额,冻结金额
func (c *Client) QueryCustAcctId(ctx context.Context, req *QueryCustAcctIdReq) (*QueryCustAcctIdRsp, error) {
	return Invoke2[QueryCustAcctIdReq, QueryCustAcctIdRsp](c, ctx, KFEJZB6037, req)
}

// QueryCustAcctIdBalance 查询会员子账号余额
// 1，当入参【保留域】为空时，查询子账户的可用余额和待支付或冻结的担保金额。
// 2，当入参【保留域】为见证有效子账户时，查询【子账户账号】的可用余额和待支付或冻结的担保金额，且出参【保留域】返回入参【保留域】待支付给【子账户账号】的担保金额。（已废除）
// 3，当入参【保留域】值为“SameAccount”时，查询子账户的可用余额（含其同名子账户，如有）和待支付或冻结的担保金额，且出参【保留域】返回其代收担保金额。
func (c *Client) QueryCustAcctIdBalance(ctx context.Context, req *QueryCustAcctIdBalanceReq) (*QueryCustAcctIdBalanceRsp, error) {
	return Invoke2[QueryCustAcctIdBalanceReq, QueryCustAcctIdBalanceRsp](c, ctx, KFEJZB6093, req)
}

// CustAcctIdBalanceQuery 查询银行子账户余额
func (c *Client) CustAcctIdBalanceQuery(ctx context.Context, req *CustAcctIdBalanceQueryReq) (*CustAcctIdBalanceQueryRsp, error) {
	return Invoke2[CustAcctIdBalanceQueryReq, CustAcctIdBalanceQueryRsp](c, ctx, KFEJZB6010, req)
}

// SingleTransactionStatusQuery 查询银行单笔交易状态
func (c *Client) SingleTransactionStatusQuery(ctx context.Context, req *SingleTransactionStatusQueryReq) (*SingleTransactionStatusQueryRsp, error) {
	return Invoke2[SingleTransactionStatusQueryReq, SingleTransactionStatusQueryRsp](c, ctx, KFEJZB6110, req)
}

// BankClearQuery 查询银行在途清算结果
func (c *Client) BankClearQuery(ctx context.Context, req *BankClearQueryReq) (*BankClearQueryRsp, error) {
	return Invoke2[BankClearQueryReq, BankClearQueryRsp](c, ctx, KFEJZB6108, req)
}

// BankCostDsDealResultQuery 查询银行费用扣收结果
func (c *Client) BankCostDsDealResultQuery(ctx context.Context, req *BankCostDsDealResultQueryReq) (*BankCostDsDealResultQueryRsp, error) {
	return Invoke2[BankCostDsDealResultQueryReq, BankCostDsDealResultQueryRsp](c, ctx, KFEJZB6109, req)
}

// CustAcctIdHistoryBalanceQuery 查询子帐号历史余额及待转可提现状态信息
func (c *Client) CustAcctIdHistoryBalanceQuery(ctx context.Context, req *CustAcctIdHistoryBalanceQueryReq) (*CustAcctIdHistoryBalanceQueryRsp, error) {
	return Invoke2[CustAcctIdHistoryBalanceQueryReq, CustAcctIdHistoryBalanceQueryRsp](c, ctx, KFEJZB6114, req)
}

// SupAcctIdBalanceQuery 查询资金汇总账户的当前余额及上日余额
func (c *Client) SupAcctIdBalanceQuery(ctx context.Context, req *SupAcctIdBalanceQueryReq) (*SupAcctIdBalanceQueryRsp, error) {
	return Invoke2[SupAcctIdBalanceQueryReq, SupAcctIdBalanceQueryRsp](c, ctx, KFEJZB6011, req)
}

// BankWithdrawCashBackQuery 查询银行提现退单信息
// 通过大小额支付系统进行的提现交易，会存在被收款行退票的情况。
// 本接口用于查询此类退票交易的记录，可以退票日期为区间查询该段时间区间内的退票记录。
func (c *Client) BankWithdrawCashBackQuery(ctx context.Context, req *BankWithdrawCashBackQueryReq) (*BankWithdrawCashBackQueryRsp, error) {
	return Invoke2[BankWithdrawCashBackQueryReq, BankWithdrawCashBackQueryRsp](c, ctx, KFEJZB6048, req)
}

// CommonTransferRechargeQuery 查询普通转账充值明细
func (c *Client) CommonTransferRechargeQuery(ctx context.Context, req *CommonTransferRechargeQueryReq) (*CommonTransferRechargeQueryRsp, error) {
	return Invoke2[CommonTransferRechargeQueryReq, CommonTransferRechargeQueryRsp](c, ctx, KFEJZB6050, req)
}

// BankTransactionDetailsQuery 查询银行时间段内交易明细
func (c *Client) BankTransactionDetailsQuery(ctx context.Context, req *BankTransactionDetailsQueryReq) (*BankTransactionDetailsQueryRsp, error) {
	return Invoke2[BankTransactionDetailsQueryReq, BankTransactionDetailsQueryRsp](c, ctx, KFEJZB6072, req)
}

// BankWithdrawCashDetailsQuery 查询银行时间段内清分提现明细
func (c *Client) BankWithdrawCashDetailsQuery(ctx context.Context, req *BankWithdrawCashDetailsQueryReq) (*BankWithdrawCashDetailsQueryRsp, error) {
	return Invoke2[BankWithdrawCashDetailsQueryReq, BankWithdrawCashDetailsQueryRsp](c, ctx, KFEJZB6073, req)
}

// DetailVerifiedCodeQuery 查询明细单验证码
// 根据见证宝系统的交易流水号(前置流水号)查询其对应的明细单验证码
func (c *Client) DetailVerifiedCodeQuery(ctx context.Context, req *DetailVerifiedCodeQueryReq) (*DetailVerifiedCodeQueryRsp, error) {
	return Invoke2[DetailVerifiedCodeQueryReq, DetailVerifiedCodeQueryRsp](c, ctx, KFEJZB6142, req)
}

// ReconciliationDocumentQuery 查询对账文件信息
// 平台调用该接口获取需下载对账文件的文件名称以及密钥.
// 平台获取到信息后, 可以再调用OPENAPI的文件下载功能
func (c *Client) ReconciliationDocumentQuery(ctx context.Context, req *ReconciliationDocumentQueryReq) (*ReconciliationDocumentQueryRsp, error) {
	return Invoke2[ReconciliationDocumentQueryReq, ReconciliationDocumentQueryRsp](c, ctx, KFEJZB6103, req)
}

// ChargeDetailQuery 查询充值明细交易的状态-见证收单
func (c *Client) ChargeDetailQuery(ctx context.Context, req *ChargeDetailQueryReq) (*ChargeDetailQueryRsp, error) {
	return Invoke2[ChargeDetailQueryReq, ChargeDetailQueryRsp](c, ctx, KFEJZB6146, req)
}

// EJZBCustInformationQuery 见证子台账信息查询接口，支持产业结算通平台调用，查询子台账开立和补录的信息。
func (c *Client) EJZBCustInformationQuery(ctx context.Context, req *EJZBCustInformationQueryReq) (*EJZBCustInformationQueryRsp, error) {
	return Invoke2[EJZBCustInformationQueryReq, EJZBCustInformationQueryRsp](c, ctx, KFEJZB6324, req)
}
