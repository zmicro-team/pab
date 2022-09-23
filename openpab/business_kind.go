package openpab

import (
	"golang.org/x/exp/slices"
)

const (
	TRUE  = "1" // 是
	FALSE = "2" // 否
)

// MemberProperty 成员属性,即子账号类型
const (
	MemberPropertyNormal = "00"
	MemberPropertySH     = "SH"

	// MemberGlobalType 证件类型
	MemberGlobalTypeIdentityCard            = "1"  // 身份证
	MemberGlobalTypeResidentPass            = "3"  // 港澳台居民通行证(即回乡证)
	MemberGlobalTypeChinesePassport         = "4"  // 中国护照
	MemberGlobalTypeMTPs                    = "5"  // 台湾居民来往大陆通行证（即台胞证）
	MemberGlobalTypeForeignPassport         = "19" // 外国护照
	MemberGlobalTypeOrgCodeCertificate      = "52" // 组织机构代码证
	MemberGlobalTypeBusinessLicense         = "68" // 营业执照
	MemberGlobalTypeUnifiedSocialCreditCode = "73" // 统一社会信用代码

	// 银行类型
	BankTypeOur   = "1" // 本行
	BankTypeOther = "2" // 他行

	// 修改方式 - 申请修改手机号
	ModifyTypeSms      = "1"
	ModifyTypeUnionPay = "2"

	SignChannelNone   = "0" // 无
	SignChannelApp    = "1" // app
	SignChannelH5     = "2" // 平台H5页面
	SignChannelPublic = "3" // 公众号
	SignChannelApplet = "4" // 小程序

	// 补录是否成功标志
	ReinFlagSuccess = "S"
	ReinFlagFailed  = "F"

	// 币种
	CcyRMB = "RMB"

	// 交易类型
	TranTypeNormal = "01" // 普通交易

	// 第三方 支付渠道类型
	PayChannelTypeWechat              = "0001" // 微信
	PayChannelTypeAlipay              = "0002" // 支付宝
	PayChannelTypeJD                  = "0003" // 京东支付
	PayChannelTypeBaidu               = "0004" // 百度支付
	PayChannelTypeKJT                 = "0005" // 快捷通支付
	PayChannelTypeYF                  = "0006" // 裕福支付
	PayChannelTypeChinaUmsInstallment = "0007" // 银联商务分期
	PayChannelTypeLaKaLa              = "0008" // 拉卡拉支付
	PayChannelTypePinAn               = "0009" // 平安付
	PayChannelTypeQQ                  = "0010" // QQ钱包
	PayChannelTypeAllinPay            = "0011" // 通联
	PayChannelTypeChinaUms            = "0012" // 银联商务
	PayChannelTypeUnionPay            = "0013" // 银联
	PayChannelTypeSuNing              = "0014" // 苏宁支付
	PayChannelTypeYeahKa              = "0015" // 乐刷支付
	PayChannelTypeCPCN                = "0016" // 中金支付
	PayChannelTypeHeli                = "0017" // 合利宝
	PayChannelTypeYee                 = "0018" // 易宝支付
	PayChannelTypeCMB                 = "0019" // 招行一网通
	PayChannelTypeXiaoMi              = "0020" // 小米支付
	PayChannelTypeUl                  = "0021" // 合众支付
	PayChannelTypeCm                  = "0022" // 和包支付
	PayChannelTypeBest                = "0023" // 翼支付
	PayChannelTypeBCM                 = "0024" // 交通银行
	PayChannelTypeCCB                 = "0025" // 建设银行
	PayChannelTypeBaoFu               = "0026" // 宝付
	PayChannelTypeICBC                = "0027" // 工商银行
	PayChannelTypeLianLian            = "0028" // 连连支付
	PayChannelTypeQianDai             = "0029" // 钱袋宝
	PayChannelTypeETC                 = "0030" // ETC
	PayChannelTypeNetEase             = "0031" // 网易宝

	// 6082 交易类型
	ApplicationTextMsgDynamicCodeTranTypeWithdraw    = "1" // 提现
	ApplicationTextMsgDynamicCodeTranTypePay         = "2" // 支付
	ApplicationTextMsgDynamicCodeTranTypeBatchPay    = "3" // 批次号支付
	ApplicationTextMsgDynamicCodeTranTypeBatchCancel = "4" // 批次号作废
)

// 6103 对账文件类型
const (
	ReconciliationDocFileTypeCZ  = "CZ"
	ReconciliationDocFileTypeTX  = "TX"
	ReconciliationDocFileTypeJY  = "JY"
	ReconciliationDocFileTypeYE  = "YE"
	ReconciliationDocFileTypeJQ  = "JQ"
	ReconciliationDocFileTypePOS = "POS"
	ReconciliationDocFileTypeJG  = "JG"
	ReconciliationDocFileTypeGJ  = "GJ"
)

var reconciliationDocFileTypeSlices = []string{
	ReconciliationDocFileTypeCZ,
	ReconciliationDocFileTypeTX,
	ReconciliationDocFileTypeJY,
	ReconciliationDocFileTypeYE,
	ReconciliationDocFileTypeJQ,
	ReconciliationDocFileTypePOS,
	ReconciliationDocFileTypeJG,
	ReconciliationDocFileTypeGJ,
}

// IsReconciliationDocFileType 是否是对账文件类型
func IsReconciliationDocFileType(ft string) bool {
	return slices.Contains(reconciliationDocFileTypeSlices, ft)
}
