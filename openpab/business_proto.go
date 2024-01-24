package openpab

// ResponseReserved 返回仅带保留域
type ResponseReserved struct {
	ResponseData
	ReservedMsg string
}

// ResponseFrontSeqNo 返回带前置流水和保留域
type ResponseFrontSeqNo struct {
	ResponseData
	FrontSeqNo  string // 见证系统流水号
	ReservedMsg string
}

// AutonymOpenCustAcctIdReq 子账户实名开立
// NOTE: 销户时仍有卡在绑, 不能解绑
type AutonymOpenCustAcctIdReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	FunctionFlag      string // required(1), 功能标志, 1: 开户, 2: 销户, 3: 存量实名
	TranNetMemberCode string // required(32), 交易网会员代码(即在平台端系统的会员编号)
	MemberName        string // required(120), 客户真实姓名(销户时可选)
	MemberGlobalType  string // required(2), 会员证件类型
	MemberGlobalId    string // required(20), 证件号码(销户时可选)
	UserNickname      string `json:",omitempty"` // optional(120), 用户昵称
	Mobile            string // required(12), 手机号(测试环境送11个1)
	MemberProperty    string // required(2), 会员属性, SH: 商户子账户, 00: 普通子账户
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

// AutonymOpenCustAcctIdRsp 开户回复
type AutonymOpenCustAcctIdRsp struct {
	ResponseData
	SubAcctNo   string // 会员子账号
	ReservedMsg string // 保留信息
}

type MntMbrBindSameRealCustNameAcctReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	// 1. 申请开立商户子账户和普通会员子账户
	// 2. 申请为已有的商户子账户开立关联的同名的普通会员子账户
	// 3. 申请为已有的普通会员子账户开立关联的同名的商户子账户
	// 4. 查询两个会员子账户是否存在同名账户关系
	// 5. 维护两个子台帐信息一致的SH和00户为同名户关系
	FunctionFlag            string // required(1), 功能标志
	MerSubAcctMemberCode    string // required(32), 商户子账户的交易网会员代码
	MerSubAcctNickname      string `json:",omitempty"` // optional(64), 商户子账户的交易网会员昵称
	CommonSubAcctMemberCode string // required(32), 普通子账户的交易网会员代码
	CommonSubAcctNickname   string `json:",omitempty"` // optional(64), 普通子账户的交易网会员昵称
	Mobile                  string `json:",omitempty"` // optional(11), 手机号
	Email                   string `json:",omitempty"` // optional(64), 邮箱
	// 同6248实名开户接口描述
	// 1分支: 必传
	// 2分支,3分支：
	//      当会员名称+会员证件类型+会员证件号码不传则实名原子账户信息;
	//      传一个值验证一个，传2个值验证2个;传三个验证三个值是否与原实名子账户相同
	MemberName       string `json:",omitempty"` // optional(120), 客户真实姓名
	MemberGlobalType string `json:",omitempty"` // optional(2), 会员证件类型(1,52,68,73)
	MemberGlobalId   string `json:",omitempty"` // optional(20), 证件号码
	// 保留域
	ReservedMsgOne   string `json:",omitempty"` // optional(120), 保留域1
	ReservedMsgTwo   string `json:",omitempty"` // optional(120), 保留域2
	ReservedMsgThree string `json:",omitempty"` // optional(120), 保留域3
	ReservedMsgFour  string `json:",omitempty"` // optional(120), 保留域3
}

type MntMbrBindSameRealCustNameAcctRsp struct {
	ResponseData
	SameNameAcctRelRelation string // required(1), 1: 是
	MerSubAcctNo            string // required(32), 商户子账户账号
	CommonSubAcctNo         string // required(32), 普通子账户账号
	// 保留域
	ReservedMsgOne   string `json:",omitempty"` // optional(120), 保留域1
	ReservedMsgTwo   string `json:",omitempty"` // optional(120), 保留域2
	ReservedMsgThree string `json:",omitempty"` // optional(120), 保留域3
	ReservedMsgFour  string `json:",omitempty"` // optional(120), 保留域3
}

// QueryCustAcctIdByThirdCustIdReq 根据会员代码查询会员子账号请求
type QueryCustAcctIdByThirdCustIdReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	TranNetMemberCode string // required(32), 交易网会员代码(即在平台端系统的会员编号)
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

// QueryCustAcctIdByThirdCustIdRsp 根据会员代码查询会员子账号回复
type QueryCustAcctIdByThirdCustIdRsp struct {
	ResponseData
	SubAcctNo   string // 会员子账号
	ReservedMsg string // 返回会员属性 00: 普通子账户, SH: 商户子账户
}

type MemberBindQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo  string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	QueryFlag   string // required(1), 查询标志, 1: 全部会员, 2: 单个会员
	SubAcctNo   string `json:",omitempty"` // optional(32), 见证子账号, QueryFlag = 2 必填
	PageNum     string // required(6), 页码,每次返回20条
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
}

type TranItem struct {
	FundSummaryAcctNo  string `json:"FundSummaryAcctNo"`
	SubAcctNo          string `json:"SubAcctNo"`
	TranNetMemberCode  string `json:"TranNetMemberCode"`
	MemberName         string `json:"MemberName"`
	MemberGlobalType   string `json:"MemberGlobalType"`
	MemberGlobalID     string `json:"MemberGlobalId"`
	MemberAcctNo       string `json:"MemberAcctNo"`
	BankType           string `json:"BankType"`
	AcctOpenBranchName string `json:"AcctOpenBranchName"`
	CnapsBranchID      string `json:"CnapsBranchId"`
	EiconBankBranchID  string `json:"EiconBankBranchId"`
	Mobile             string `json:"Mobile"`
}

type MemberBindQueryRsp struct {
	ResponseData
	ResultNum     string     `json:"ResultNum"`
	StartRecordNo string     `json:"StartRecordNo"`
	EndFlag       string     `json:"EndFlag"`
	TotalNum      string     `json:"TotalNum"`
	ReservedMsg   string     `json:"ReservedMsg"`
	TranItemArray []TranItem `json:"TranItemArray"`
}

type BindUnionPayWithCheckCorpReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个会员子账号, 并且 '|::|' 分隔, 注意顺序
	TranNetMemberCode string // required(32), 交易网会员代码, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个交易网会员代码, 并且 '|::|' 分隔, 注意顺序
	// 个人工商户此信息默认为法人信息, 当 RepFlag = 2 时,该信息是代办人信息
	// 会员为纯个人时,证件仅支持(1,3,4,5,19), 为个体工商户时仅支持1
	MemberName       string // required(120), 客户真实姓名
	MemberGlobalType string // required(2), 会员证件类型(1,3,4,5,19)
	MemberGlobalId   string // required(20), 证件号码
	// 银行信息
	MemberAcctNo       string // required(32), 提现的银行卡
	BankType           string // required(1), 银行类型 1: 本行 2: 他行
	AcctOpenBranchName string // required(120), 开户行名称
	CnapsBranchId      string `json:",omitempty"` // optional(14), 大小额行号, NOTE: CnapsBranchId 和 EiconBankBranchId 必填其一
	EiconBankBranchId  string `json:",omitempty"` // optional(14), 超级网银行号
	Mobile             string // required(12), 手机号
	// IndivBusinessFlag = 1 个体工商户必填信息
	IndivBusinessFlag string `json:",omitempty"` // required(1), 是否个体工商户, 1: 是 2: 否
	CompanyName       string `json:",omitempty"` // optional(120), 公司名称, 个体工商户必填
	CompanyGlobalType string `json:",omitempty"` // optional(4), 公司证件类型(52,68,73), 个体工商户必填
	CompanyGlobalId   string `json:",omitempty"` // optional(32), 公司证件号码, 个体工商户必填
	ShopId            string `json:",omitempty"` // optional(32), 店铺id, 个体工商户必填,可与交易网会员代码相同
	ShopName          string `json:",omitempty"` // optional(120), 店铺名称, 个体工商户必填
	// RepFlag = 2 会员名称不是法人信息时必填
	RepFlag        string // required(1), 会员名称是否为法人, 1: 是 2: 否, 个体工商户必填
	ReprName       string `json:",omitempty"` // optional(32), 法人名称, RepFlag = 2 必填
	ReprGlobalType string `json:",omitempty"` // optional(32), 法人证件类型, RepFlag = 2 必填
	ReprGlobalId   string `json:",omitempty"` // optional(32), 法人证件号码, RepFlag = 2 必填
	// 其它
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
	Flag        string `json:",omitempty"` // optional(120), 标志(未启用)
}

type CheckMsgCodeWithCorpReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个会员子账号, 并且 '|::|' 分隔, 注意顺序
	TranNetMemberCode string // required(32), 交易网会员代码, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个会员子账号, 并且 '|::|' 分隔, 注意顺序
	MemberAcctNo      string // required(32), 提现的银行卡
	MessageCheckCode  string // required(7), 短信验证码
	// 其它
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
}

type BindSmallAmountWithCheckCorpReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个会员子账号, 并且 '|::|' 分隔, 注意顺序
	TranNetMemberCode string // required(32), 交易网会员代码, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个会员的交易网代码, 并且 '|::|' 分隔, 注意顺序
	// 个人工商户此信息默认为法人信息, RepFlag = 2 时,该信息是代办人信息
	// 会员为企业时,证件仅支持(52,68,73), 为纯个人或个体工商户时仅支持1
	MemberName       string // required(120), 客户真实姓名
	MemberGlobalType string // required(2), 会员证件类型(1,52,68,73)
	MemberGlobalId   string // required(20), 证件号码
	// 银行信息
	// 银行账号为平安银行账户时, BankType = 1, 大小额行号和超级网银行号都可不填
	MemberAcctNo       string // required(32), 提现的银行卡
	BankType           string // required(1), 银行类型 1: 本行 2: 他行
	AcctOpenBranchName string // required(120), 开户行名称
	CnapsBranchId      string `json:",omitempty"` // optional(14), 大小额行号, NOTE: CnapsBranchId 和 EiconBankBranchId 必填其一
	EiconBankBranchId  string `json:",omitempty"` // optional(14), 超级网银行号
	Mobile             string // required(12), 手机号
	// IndivBusinessFlag = 1 个体工商户必填信息
	IndivBusinessFlag string `json:",omitempty"` // optional(1), 是否个体工商户, 1: 是 2: 否
	CompanyName       string `json:",omitempty"` // optional(120), 公司名称, 个体工商户必填
	CompanyGlobalType string `json:",omitempty"` // optional(4), 公司证件类型(52,68,73), 个体工商户必填
	CompanyGlobalId   string `json:",omitempty"` // optional(32), 公司证件号码, 个体工商户必填
	// 个体工商户或企业必输
	ShopId   string `json:",omitempty"` // optional(32), 店铺id, 个体工商户,企业必填,可以和交易网会员代码相同
	ShopName string `json:",omitempty"` // optional(120), 店铺名称, 个体工商户,企业必填
	// AgencyClientFlag = 1 经办人信息必填
	AgencyClientFlag       string `json:",omitempty"` // required(1), 是否存在经办人, 1: 是, 2: 否
	AgencyClientName       string `json:",omitempty"` // optional(120), 经办人姓名, AgencyClientFlag = 1 必填
	AgencyClientGlobalType string `json:",omitempty"` // optional(4), 经办人证件类型,仅支持身份证, AgencyClientFlag = 1 必填
	AgencyClientGlobalId   string `json:",omitempty"` // optional(32), 经办人证件号, AgencyClientFlag = 1 必填
	AgencyClientMobile     string `json:",omitempty"` // optional(32), 经办人手机号, AgencyClientFlag = 1 必填
	// RepFlag = 2 会员名称不是法人信息时必填
	RepFlag        string `json:",omitempty"` // required(1), 会员名称是否为法人, 1: 是 2: 否, 个体工商户、企业必填，纯个人无需填写
	ReprName       string `json:",omitempty"` // optional(32), 法人名称, RepFlag = 2 必填
	ReprGlobalType string `json:",omitempty"` // optional(32), 法人证件类型(1,3,5,19), RepFlag = 2 必填
	ReprGlobalId   string `json:",omitempty"` // optional(32), 法人证件号码, RepFlag = 2 必填
	// 其它
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
	Flag        string `json:",omitempty"` // optional(120), 标志(未启用)
}

type CheckAmountWithCorpReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个会员子账号, 并且 '|::|' 分隔, 注意顺序
	TranNetMemberCode string // required(32), 交易网会员代码 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个交易网会员代码, 并且 '|::|' 分隔, 注意顺序
	TakeCashAcctNo    string // required(32), 提现的银行卡
	AuthAmt           string // required(15), 鉴权金额
	OrderNo           string // required(5), 指令号
	Ccy               string // required(3), 币种, 默认RMB
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

type UnbindRelateAcctReq struct {
	// required(32), 资金监管账号(底层已实现)
	// FundSummaryAcctNo string
	// optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	CnsmrSeqNo string
	// required(1), 功能标志, 1: 解绑
	FunctionFlag string
	// required(32), 子账户账号, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个会员子账号, 并且 '|::|' 分隔, 注意顺序
	SubAcctNo string
	// required(32), 交易网会员代码, 若需要把待绑定账户关联到两个会员名下, 此字段可上送两个交易网代码, 并且 '|::|' 分隔, 注意顺序
	TranNetMemberCode string
	// required(32), 待解绑的提现的银行卡
	MemberAcctNo string
	// optional(120), 保留域
	ReservedMsg string `json:",omitempty"`
}

type ApplyForChangeOfCellPhoneNumReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码
	ModifyType        string // required(2), 修改方法, 1: 短信验证码, 2: 银联鉴权
	NewMobile         string // required(12) 新手机号, ModifyType=2 应为银行卡预留手机号
	BankCardNo        string `json:",omitempty"` // optional(32), 银行卡号 ModifyType=2 必填
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

type ApplyForChangeOfCellPhoneNumRsp struct {
	ResponseData
	ReceiveMobile  string // required(12), 接收的手机号码,银行只返回后四位
	MessageOrderNo string // required(32), 短信指令号
	ReservedMsg    string // optional(120), 保留域
}

type BackfillDynamicPasswordReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码
	ModifyType        string // required(2), 修改方法, 1: 短信验证码, 2: 银联鉴权
	MessageOrderNo    string // required(32)  短信指令号
	MessageCheckCode  string // required(7) 短信验证码
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

type MntMbrBindRelateAcctBankCodeReq struct {
	// required(32), 资金监管账号(底层已实现)
	// FundSummaryAcctNo string
	// optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	CnsmrSeqNo string
	// required(32), 子账户账号
	SubAcctNo string
	// required(32), 会员绑定账号
	MemberBindAcctNo string
	// 若大小额行号不填则送超级网银号对应的银行名称
	// 若填大小额行号则送大小额行号对应的银行名称
	// required(120), 开户行名称
	AcctOpenBranchName string
	// NOTE: CnapsBranchId 和 EiconBankBranchId 必填其一
	// optional(14), 大小额行号,
	CnapsBranchId string `json:",omitempty"`
	// optional(14), 超级网银行号
	EiconBankBranchId string `json:",omitempty"`
	// optional(120), 保留域
	ReservedMsg string `json:",omitempty"`
}

type RegisterBehaviorRecordInfoReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码
	FunctionFlag      string // required(2), 功能标志,1: 登记行为记录信息, 2: 查询补录信息
	OpClickTime       string `json:",omitempty"` // optional(14), 操作点击的时间, 格式: 20201222171623,  FunctionFlag = 1 必填
	IpAddress         string `json:",omitempty"` // optional(20), ip地址, 格式: 22.25.12.33, FunctionFlag = 1 必填
	MacAddress        string `json:",omitempty"` // optional(20), MAC地址, 格式: 5E:E3:90:4R:E5:33, FunctionFlag = 1 必填
	SignChannel       int    `json:",omitempty"` // optional(2), 签约渠道, 1-app 2-平台H5网页 3-公众号 4-小程序, FunctionFlag = 1 必填
	ReservedMsgOne    string `json:",omitempty"` // optional(120), 保留域
	ReservedMsgTwo    string `json:",omitempty"` // optional(120), 保留域
}

type RegisterBehaviorRecordInfoRsp struct {
	ReinSuccessFlag   string `json:",omitempty"` // 补录是否成功标志, S: 成功, F: 失败, 功能标志=2时必填
	CompanyName       string `json:",omitempty"` // 单位名称, 功能标志=2时且商户为非自然人
	CompanyGlobalType string `json:",omitempty"` // 公司证件类型, 功能标志=2时且商户为非自然人
	CompanyGlobalId   string `json:",omitempty"` // 公司证件号码,  功能标志=2时且商户为非自然人
	ReprName          string `json:",omitempty"` // 法人名称,  功能标志=2时且商户为非自然人
	ReprGlobalType    string `json:",omitempty"` // 法人证件类型,  功能标志=2时且商户为非自然人
	ReprGlobalId      string `json:",omitempty"` // 法人证件号码,  功能标志=2时且商户为非自然人
	ReservedMsg       string `json:",omitempty"` // 保留域
}

type RegisterBillSupportWithdrawReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码
	OrderNo           string // required(32), 订单号, 全局唯一
	SuspendAmt        string // required(15), 挂账金额,包含交易费用，即挂账金额=会员实际到账金额+交易费用
	TranFee           string // required(15), 交易费用,平台收取用户的费用
	Remark            string `json:",omitempty"` // optional(120), 备注
	ReservedMsgOne    string `json:",omitempty"` // optional(120), 保留域1
	ReservedMsgTwo    string `json:",omitempty"` // optional(120), 保留域2
	ReservedMsgThree  string `json:",omitempty"` // optional(120), 保留域3
}

type RevRegisterBillSupportWithdrawReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码
	OldOrderNo        string // required(32), 原订单号, 原 6139 登记挂账的订单号
	CancelAmt         string // required(15), 挂账金额, 支持部分撤销，不能大于原订单可用金额，包含交易费用，即撤销金额=实际会员到账的撤销金额+交易费用
	TranFee           string // required(15), 交易费用, 平台退给用户的手续费，不能超过原订单的手续费
	Remark            string `json:",omitempty"` // optional(120), 备注
	ReservedMsgOne    string `json:",omitempty"` // optional(120), 保留域1
	ReservedMsgTwo    string `json:",omitempty"` // optional(120), 保留域2
	ReservedMsgThree  string `json:",omitempty"` // optional(120), 保留域3
}

type AccountRegulationReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo           string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo            string // required(32), 子账户账号
	TranNetMemberCode    string // required(32), 交易网会员代码
	SubAcctName          string // required(120), 见证子账户名称
	AcquiringChannelType string // required(2), 收单渠道类型, 01-橙E收款, YST1-云收款
	OrderNo              string // required(30), 订单号
	Amt                  string // required(15), 金额
	Ccy                  string // required(3), 币种
	Remark               string `json:",omitempty"` // optional(120), 备注
	ReservedMsg          string `json:",omitempty"` // optional(120), 保留域
}

type PlatformAccountSupplyReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo           string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	AcquiringChannelType string // required(2), 收单渠道类型, 01-橙E收款, YST1-云收款
	OrderNo              string // required(30), 订单号, 根据所填渠道所返回的订单号，这里是总订单号
	Amt                  string // required(15), 金额
	ReservedMsg          string `json:",omitempty"` // optional(120), 保留域
}

type PlatformAccountSupplyRsp struct {
	ResponseData
	FrontSeqNo  string // required(16)
	SubAcctNo   string // required(32)
	Amt         string // required(15)
	Remark      string // required(120)
	ReservedMsg string // required(20)
}

type RcvMntMbrInfoApplyReq struct {
	CnsmrSeqNo               string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	FunctionFlag             string // M-修改 Q-查询
	WitnessSubAcctNo         string // 见证子账户的账号
	TranNetMemberCode        string // 交易网会员代码
	ModifyReason             string // 修改原因
	ModifyWitnessSubAcctName string // 修改后的见证子账户的户名
	ModifyMemberGlobalType   string // 修改后的会员的证件类型
	ModifyMemberGlobalId     string // 修改后的会员的证件号码
	HoldOne                  string // 预留字段1
	HoldTwo                  string // 预留字段2
	HoldThree                string // 预留字段3
	HoldFour                 string // 预留字段4
	HoldFive                 string // 预留字段5
}
type MemberInformationChangeReq struct {
	CnsmrSeqNo        string // optional(22), 交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码
	MemberName        string // required(120), 会员名称
	CompanyName       string `json:",omitempty"` // optional(120), 公司名称
	ReprName          string `json:",omitempty"` // optional(120), 法人名称
	ReprGlobalType    string `json:",omitempty"` // optional(32), 法人证件类型
	ReprGlobalId      string `json:",omitempty"` // optional(32), 法人证件号码
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

/************************************* 交易 ************************************/

type MemberTransactionReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo     string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	FunctionFlag   string // required(1), 功能标志, 1: 下单预支付, 2: 确认并支付, 3: 退款(仅退预支付) 4: 同名子账户支付, 6: 直接支付T+1, 7: 支付到平台, 9: 直接支付T+0
	OutSubAcctNo   string // required(32), 转出方的见证子账户账号
	OutMemberCode  string // required(32), 转出方的交易网会员代码
	OutSubAcctName string `json:",omitempty"` // optional(120), 转出方的见证子账户户名
	InSubAcctNo    string // required(32), 转入方的见证子账户账号
	InMemberCode   string // required(32), 转入方的交易网会员代码
	InSubAcctName  string `json:",omitempty"` // optional(120), 转入方的见证子账户户名
	TranAmt        string // required(15), 交易金额(单位:元), 含交易费用, 会员的实际到账金额=交易金额-交易费用
	TranFee        string // required(15), 交易费用(单位:元), 平台收取交易费用
	TranType       string // required(2), 交易类型, 01: 普通交易
	Ccy            string // required(3), 币种, 默认: RMB
	OrderNo        string `json:",omitempty"` // optional(30), 订单号, 功能为1,2,3时必填,全局唯一,不能与6139/6007/6135/6134订单号相同
	OrderContent   string `json:",omitempty"` // optional(500), 订单内容
	Remark         string `json:",omitempty"` // optional(120), 备注, 建议送订单号,可以对账备注字段获取到
	ReservedMsg    string `json:",omitempty"` // optional(120), 保留域, 若需短信验证码则此项必输短信指令号
	WebSign        string `json:",omitempty"` // optional(256), 网银签名, 若需短信验证码则此项必输
}

type MembershipWithdrawCashReq struct {
	// required(32), 资金监管账号(底层已实现)
	// FundSummaryAcctNo string
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	// required(32), 子账户账号
	SubAcctNo string
	// required(32), 交易网会员代码
	TranNetMemberCode string
	// required(120), 交易网名称, 平台自定义名称或填写4位的市场代码（MrchCode）
	TranWebName string
	// 客户信息
	// required(120), 客户真实姓名
	MemberName string
	// optional(2), 会员证件类型
	MemberGlobalType string `json:",omitempty"`
	// optional(20), 会员证件号码
	MemberGlobalId string `json:",omitempty"`
	// required(32), 提现账号(银行卡)
	TakeCashAcctNo string
	// required(120), 出账账户名称(银行卡名称)
	OutAmtAcctName string
	// required(3), 币种, 默认RMB
	Ccy string
	// required(15), 可提现金额(单位: 元),不包含手续费，实际到账金额=申请提现的金额
	CashAmt string
	// optional(120), 备注, 建议可送订单号, 可在对账文件的备注字段获取到
	Remark string `json:",omitempty"`
	// optional(120), 提现手续费(单位: 元), 格式: 0.00
	ReservedMsg string `json:",omitempty"`
	// optional(256), 网银签名
	WebSign string `json:",omitempty"`
}

type MembershipWithdrawCashRsp struct {
	ResponseData
	// required(16), 见证系统流水号
	FrontSeqNo string
	// required(15), 转账手续费, 未启用, 固定0.00
	TransferFee string
	// required(20), 保留域
	ReservedMsg string
}

type MembershipTrancheFreezeReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	// 1: 冻结 (会员→担保)
	// 2: 解冻 (担保→会员)
	// 4: 见证+收单的冻结资金解冻
	// 5: 可提现冻结 (会员→担保)
	// 6: 可提现解冻 (担保→会员)
	// 7: 在途充值解冻 担保→会员
	FunctionFlag      string // required(1), 功能标志
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码
	ConsumeAmt        string // required(15), 消费的金额(单位: 元)
	TranAmt           string // required(15), 交易金额(单位: 元), 需解冻/冻结的金额，包含手续费，即交易金额=实际解冻到账金额+手续费
	TranCommission    string // required(15), 交易手续费(单位: 元),解冻时，将根据该金额收取手续费，若无手续费则送0. 冻结时，该字段不启用
	Ccy               string // required(3), 币种 RMB
	OrderNo           string // required(30), 订单号, 全局唯一，不能与6034/6101/6006/6139订单号相同. 如果是解冻，需要送冻结时送的子订单号或订单号
	OrderContent      string `json:",omitempty"` // optional(500), 订单内容
	Remark            string `json:",omitempty"` // optional(120), 备注
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

type OnWayTopThirdPaySplitReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo            string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	PayChannelType        string // required(4), 支付渠道类型
	PayChannelAssignMerNo string // required(64), 支付渠道所分配的商户号, 即市场在第三方支付渠道的商户号
	TotalOrderNo          string // required(64) 总订单号
	TranTotalAmt          string // required(32) 交易总金额(单位: 元)
	OrdersCount           string // required(2) 订单数量(1-25条)
	TranItemArray         []OnWayTopThirdPaySplitTranItem
	ReservedMsgOne        string `json:",omitempty"` // optional(120),保贸域1
	ReservedMsgTwo        string `json:",omitempty"` // optional(120),保留域2
}

type OnWayTopThirdPaySplitTranItem struct {
	RechargeSubAcctNo    string // required(32), 充值子账户, 同一笔总订单号下子账户不能重复
	SubOrderFillMemberCd string // required(32), 子订单充值会员代码
	SubOrderTranAmt      string // required(32), 子订单交易金额(单位: 元),会员的可用余额的调增金额，不包含手续费
	SubOrderTranFee      string // required(32), 子订单交易费用(单位: 元),本子订单交易充值收取的手续费
	SubOrderNo           string // required(64), 子订单号,全局唯一
	SubOrderContent      string // optional(500), 子订单内容
	SubOrderBussSeqNo    string // optional(64), 子订单业务流水号,全局唯一
	SubOrderRemark       string // optional(120), 子订单备注
	ReservedMsg          string // optional(32), 保留域
}

type OnWayTopThirdPaySplitRsp struct {
	ResponseData
	OrdersCount    string                      // 订单数量(1-25条)
	TranItemArray  []OnWayTopThirdPaySplitItem // 信息数组
	ReservedMsgOne string                      // 保留域1
	ReservedMsgTwo string                      // 保留域2
}

type OnWayTopThirdPaySplitItem struct {
	FrontSeqNo string // required(16), 前置流水号, 电商见证宝系统生成的流水号，可关联具体一笔请求
}

type RevokeOnWayTopThirdPaySplitReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	OldPayChannelType string // required(4), 原支[6216]付渠道类型
	OldTotalOrderNo   string // required(64) 原总订单号
	TotalRefundAmt    string // required(32) 退款总金额, 支付总金额（子订单会员退款金额和子订单手续费退款金额的总和）
	RefundOrderNum    string // required(4) 退款订单数量(1-25条)
	TranItemArray     []RevokeOnWayTopThirdPaySplitTranItem
	ReservedMsgOne    string `json:",omitempty"` // optional(120),保贸域1
	ReservedMsgTwo    string `json:",omitempty"` // optional(120),保留域2
}

type RevokeOnWayTopThirdPaySplitTranItem struct {
	SubOrderRefundSubAcctNo string // required(32),子订单退款子账户
	SubOrderRefundMemberCd  string // required(32),子订单退款会员代码
	SubOrderMemberRefundAmt string // required(32),子订单会员退款金额, 从会员子账号扣除的金额，不能大于原订单的交易金额
	SubOrderFeeRefundAmt    string // required(32),子订单手续费退款金额, 从手续费子账号（或者出入金账户）扣除的金额，不能大于原订单的交易费用
	SubOrderRefundOrderNo   string // required(64),子订单退款订单号
	SubOrderRefundSeqNo     string // optional(32),子订单退款流水号, 全局唯一
	ReservedMsg             string // optional(32),保留域
}

type RevokeOnWayTopThirdPaySplitRsp struct {
	ResponseData
	OrdersCount    string
	TranItemArray  []OnWayTopThirdPaySplitItem
	ReservedMsgOne string
	ReservedMsgTwo string
}

type ApplicationTextMsgDynamicCodeReq struct {
	// FundSummaryAcctNo string// required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo         string // required(32), 子账户账号
	TranNetMemberCode string // required(32), 交易网会员代码, 按批次号提现时非必输
	// 1：提现
	// 2：支付
	// 3：批次号支付
	// 4：批次号作废
	TranType    string // required(2), 交易类型
	TranAmt     string // required(15), 交易金额(单位: 元), 按批次号提现时非必输
	BankCardNo  string `json:",omitempty"` // optional(32), 银行卡号
	OrderNo     string `json:",omitempty"` // optional(30), 订单号,按批次号提现时，此字段必须, 上送批次号字段
	Remark      string `json:",omitempty"` // optional(120), 备注
	ReservedMsg string `json:",omitempty"` // optional(120), 当所申请的短信验证码是用于进行6101接口的功能分支2或4或6或9的场景，须上送转入方的见证子账户的账号
}

type ApplicationTextMsgDynamicCodeRsp struct {
	ResponseData
	ReceiveMobile  string // 接收手机号码, 只有后4位
	MessageOrderNo string // 短信指令号
	ReservedMsg    string // 保留域
}

type MemberTranVerifyTextMsgsReq struct {
	// FundSummaryAcctNo string// required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	// 6: 直接支付T+1
	// 9: 直接支付T+0
	FunctionFlag     string // required(1), 功能标志
	OutSubAcctNo     string // required(32), 转出方的见证子账户的账号, 营销子账号
	OutMemberCode    string // required(32), 转出方的交易网会员代码
	OutSubAcctName   string `json:",omitempty"` // optional(120), 户名就送开户接口 AutonymOpenCustAcctId 接口上送的用户MemberName会员名称
	InSubAcctNo      string // required(32), 转入方的见证子账户的账号
	InMemberCode     string // required(32), 转入方的交易网会员代码
	InSubAcctName    string `json:",omitempty"` // optional(120), 转入方的, 户名就送开户接口 AutonymOpenCustAcctId 接口上送的用户MemberName会员名称
	TranAmt          string // required(15), 交易金额(单位: 元), 包含交易费用，会员的实际到账金额=交易金额-交易费用
	TranFee          string // required(15), 交易费用(单位: 元), 平台收取交易费用
	TranType         string // required(2), 交易类型, 01：普通交易
	Ccy              string // required(3), 币种, 默认: RMB
	OrderNo          string `json:",omitempty"` // optional(20), 订单号
	OrderContent     string `json:",omitempty"` // optional(500), 订单内容
	Remark           string `json:",omitempty"` // optional(120), 备注, 建议可送订单号，可在对账文件的备注字段获取到.
	MessageOrderNo   string `json:",omitempty"` // optional(120), 短信指令号, 当使用短信验证时，必输
	MessageCheckCode string // required(7), 短信验证码
	ReservedMsg      string `json:",omitempty"` // optional(120), 当所申请的短信验证码是用于进行6101接口的功能分支2或4或6或9的场景，须上送转入方的见证子账户的账号
}

type MemberTransactionRefundReq struct {
	// FundSummaryAcctNo string// required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	// 1：会员交易接口 [确认付款] 退款, 针对 [6006/6034/6101] funcflag=2-确认并付款的退款
	// 2：会员交易接口 [直接支付] 退款, 针对 [6006/6034/6101] funcflag=6、9-直接支付的退款
	// 3：平台订单管理接口 [平台代理确认收货] 退款, 针对 [6031] 2-平台代理确认收货的退款
	// 4：会员批量交易接口 [批量确认] 退款, 针对 [6052/6120/6133] 2-批量确认的退款
	// 5：会员批量交易接口 [直接支付] 退款, 针对 [6052/6120/6133] 3-直接支付的退款. 未提供使用.
	// 6：会员资金支付接口的退款, 针对 [6163/6165/6166] 会员资金支付的退款
	// 7：会员交易（多借多贷）接口的退款, 针对 [6119] 6、9-直接支付的退款. 6119是专用接口，不提供开放使用.
	// X：会员交易接口 [资方交易] 退款, 针对 [6034] X-资方交易
	// H: [6034] 同名子账户支付（商户子账户A→普通会员A→商户子账户B）或 （商户子账户A→商户子账户B）
	FunctionFlag     string // required(1), 功能标志
	OldTranSeqNo     string // required(20), 原交易流水号
	OldOutSubAcctNo  string // required(32), 原转出见证子账户的帐号
	OldOutMemberCode string // required(32), 原转出会员代码
	OldInSubAcctNo   string // required(32), 原转入见证子账户的账号
	OldInMemberCode  string // required(32), 原转入会员代码
	OldOrderNo       string `json:",omitempty"` // optional(30), 原订单号, 1、3、4、5、6分支必输
	ReturnAmt        string // required(15), 退款金额, 包含退款手续费
	ReturnCommission string // required(15), 退款手续费(单位: 元), 6分支填0.00
	Remark           string `json:",omitempty"` // optional(120), 备注
	ReservedMsg      string `json:",omitempty"` // optional(120),
}

/************************************* 查询 ************************************/

type SmallAmountTransferQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	OldTranSeqNo string // required(32), 原交易流水号, 小额鉴权交易请求时的CnsmrSeqNo值（第一次申请成功时的流水号）
	TranDate     string // required(32), 格式：20060102
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type SmallAmountTransferQueryRsp struct {
	ResponseData
	ReturnStatu string // required(1), 0: 成功, 1: 失败, 2: 待确认(说明已转账成功，但是收款账户还未收到资金)
	ReturnMsg   string // optional(80), 失败返回的具体信息
	ReservedMsg string // optional(120),
}

type QueryCustAcctIdReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo        string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	TranNetMemberCode string // required(32), 交易网会员代码
	ReservedMsg       string `json:",omitempty"` // optional(120), 保留域
}

type QueryCustAcctIdRsp struct {
	ResponseData
	SubAcctNo        string // required(32), 子账户账号
	SubAcctCashBal   string // required(15), 见证子账户可提现余额
	SubAcctAvailBal  string // required(15), 见证子账户可用余额
	SubAcctFreezeAmt string // required(15), 见证子账户冻结金额(指在担保子账户里冻结的金额)
	ReservedMsg      string // required(120), 见证子账户可用余额(若平台开通智能收款,保留域返回智能收款子帐号)
}

type QueryCustAcctIdBalanceReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo  string // required(32), 子账户账号
	// 可送三类值
	// 1， 送空值
	// 2， 送对手方子账户（已废除）
	// 3， 送值“SameAccount”
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
}

type QueryCustAcctIdBalanceRsp struct {
	ResponseData
	TranNetMemberCode string // required(32), 交易网会员代码
	SubAcctAvailBal   string // required(15), 见证子账户可用余额(当入参[保留域]值为“SameAccount”时，含其同名子账户的可用余额，如有)
	SubAcctAssureAmt  string // required(15), 见证子账户担保余额(指在担保子账户里待支付或冻结的金额)
	ReservedMsg       string // optional(120), 保留域
}

type CustAcctIdBalanceQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo  string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	QueryFlag   string // required(1), 查询标志, 2：普通会员子账号, 3：功能子账号
	SubAcctNo   string `json:",omitempty"` // optional(32), 子账户账号, QueryFlag = 2 必填
	PageNum     string // required(6), 页码, 起始值为1, 每页最多20条
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
}

type CustAcctIdBalanceItem struct {
	SubAcctNo string // 见证子账户的账号
	// 1：普通会员子账号
	// 2：挂账子账号
	// 3：手续费子账号
	// 4：利息子账号
	// 5 ：平台担保子账号
	// 7：平台在途子账户
	// 11：汇总影子子账户
	// 12：提现在途子账户
	// 13：营销子账户
	// 22：特殊挂账子账户（22类型如没有返回，则表示没有）
	SubAcctProperty   string // 见证子账户的属性
	TranNetMemberCode string // 交易网会员代码
	SubAcctName       string // 见证子账户的名称
	AcctAvailBal      string // 见证子账户可用余额
	CashAmt           string // 见证子账户可提现金额
	MaintenanceDate   string // 维护日期(开户日期或修改日期)
}

type CustAcctIdBalanceQueryRsp struct {
	ResponseData
	ResultNum     string // 本次交易返回查询结果记录数
	StartRecordNo string // 起始记录号
	EndFlag       string // 结束标志
	TotalNum      string // 符合业务查询条件的记录总数
	AcctArray     []CustAcctIdBalanceItem
	ReservedMsg   string // 保留域
}

type SingleTransactionStatusQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	// 2: 会员间交易
	// 3: 提现
	// 4: 充值
	FunctionFlag string // required(1), 功能标志
	TranNetSeqNo string // required(32), 交易网流水号, 提现，充值或会员交易请求时的CnsmrSeqNo值,6216分账时请求时的SubOrderBussSeqNo（如果6216没有送，就用返回的前置流水号FrontSeqNo）
	SubAcctNo    string `json:",omitempty"` // optional(32), 子账户 (暂时未启用)
	TranDate     string `json:",omitempty"` // optional(8), 交易日期,(格式: 20060102) 当功能标志为2时，查询三天前的记录需上送交易日期
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type SingleTransactionStatusQueryRsp struct {
	ResponseData
	// 1：登记挂账
	// 2：支付
	// 3：提现
	// 4：清分
	// 5:下单预支付
	// 6：确认并付款
	// 7：退款
	// 8：支付到平台
	// N:其他
	BookingFlag string // required(1), 记账标志
	// 0：成功
	// 1：失败
	// 2：待确认
	// 5：待处理
	// 6：处理中
	// 若系统返回状态不为0或者1， 返回其他任何状态均为交易状态不明，5分钟后重新查询
	TranStatus        string // required(1), 交易状态
	TranAmt           string // required(16), 交易金额, 包含手续费，即交易金额=实际到账金额+手续费
	TranDate          string // required(8), 交易日期
	TranTime          string // required(6), 交易时间
	InSubAcctNo       string // required(32), 转入子账户号
	OutSubAcctNo      string // required(32), 转出子账户号
	FailMsg           string // required(120), 失败信息
	OldTranFrontSeqNo string // required(32), 原交易前置流水号
}

type BankClearQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	FunctionFlag string // required(1), 功能标志, 1: 全部, 2: 指定时间段
	PageNum      string // required(6), 页码, 起始值为1, 最多返回20条
	StartDate    string `json:",omitempty"` // optional(8), 开始日期, 格式: 20181201
	EndDate      string `json:",omitempty"` // optional(8), 终止日期, 格式: 20181201
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type BankClearQueryItem struct {
	Date               string // 日期, 格式: 20060102
	SubAcctType        string // 子账号类型, 7：在途子账户
	ReconcileStatus    string // 对账状态, 0；成功 1：失败
	ReconcileReturnMsg string // 对账返回的信息
	TotalAmt           string // 待清算总金额
	// 0：成功
	// 1：失败
	// 2：异常
	// 3:待处理
	ClearingStatus    string // 清算状态
	ClearingReturnMsg string // 清算返回的信息
}

type BankClearQueryRsp struct {
	ResponseData
	ResultNum     string // 本次交易返回查询结果记录数
	StartRecordNo string // 起始记录号
	EndFlag       string // 结束标志, 0: 否, 1: 是
	TotalNum      string // 符合业务查询条件的记录总数, 一次最多返回20条
	ReservedMsg   string // 保留域
	TranItemArray []BankClearQueryItem
}

type BankCostDsDealResultQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	FunctionFlag string // required(1), 1: 全部, 2: 指定时间段
	PageNum      string // required(6), 页码
	StartDate    string `json:",omitempty"` // optional(8), 开始日期, 格式: 20060102, FunctionFlag = 2 必填
	EndDate      string `json:",omitempty"` // optional(8), 终止日期, 格式: 20060102, FunctionFlag = 2 必填
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type BankCostDsDealResultQueryItem struct {
	FeeType       string // 费用类型, 1：提现手续费 2：会员验证费 3：服务费
	FeeStartDate  string // 费用起始日期
	FeeEndDate    string // 费用结束日期
	FeeDeductDate string // 费用扣收日期
	ChargeAmt     string // 收费金额
	TranStatus    string // 交易状态, 0 成功 1失败 2超时 3未处理
}

type BankCostDsDealResultQueryRsp struct {
	ResponseData
	ResultNum     string // 本次交易返回查询结果记录数
	StartRecordNo string // 起始记录号
	EndFlag       string // 结束标志, 0: 否, 1: 是
	TotalNum      string // 符合业务查询条件的记录总数, 一次最多返回20条
	ReservedMsg   string // 保留域
	FeeArray      []BankCostDsDealResultQueryItem
}

type CustAcctIdHistoryBalanceQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo    string // required(32), 子账户账号
	FunctionFlag string // required(1), 1: 全部, 2: 指定时间段
	PageNum      string // required(6), 页码, 每次最多返回20条
	StartDate    string `json:",omitempty"` // optional(8), 开始日期, 格式: 20060102
	EndDate      string `json:",omitempty"` // optional(8), 终止日期, 格式: 20060102
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type CustAcctIdHistoryBalanceQueryItem struct {
	Date            string // 日期, 格式: yymmdd
	DayAcctAvailBal string // 日终可用余额
	DayCashBal      string // 日终冻结余额
	DayFreezeBal    string // 当日待转可提现发生额
	DayCashOccurAmt string // 日终待转可提现余额
	DayWaitCashBal  string // 待转可提现金额
	CashStatus      string // 待转可提现状态 0: 待转 1: 已转 2: 无需转 3: 异常
}

type CustAcctIdHistoryBalanceQueryRsp struct {
	ResponseData
	ResultNum     string // 本次交易返回查询结果记录数
	StartRecordNo string // 起始记录号
	EndFlag       string // 结束标志, 0：否  1：是
	TotalNum      string // 符合业务查义监票件的记录总数
	AcctArray     []CustAcctIdHistoryBalanceQueryItem
	ReservedMsg   string // 保留域
}

type SupAcctIdBalanceQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo  string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
}

type SupAcctIdBalanceQueryRsp struct {
	ResponseData
	LastBalance  string // required(15), 上日余额
	CurBalabce   string // required(15), 当前余额
	Balance      string // required(15), 账户余额
	AddedBalance string // required(15), 今日增加的余额
	ReservedMsg  string // required(20), 保留域
}

type BankWithdrawCashBackQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	FunctionFlag string // required(1), 1: 提现退票, 2: 小额鉴权退票
	StartDate    string // required(8), 开始日期
	EndDate      string // required(8), 终止日期
	ReservedMsg  string // optional(120), 保留域
}

type BankWithdrawCashBackQueryItem struct {
	OldTranSeqNo            string // 原提现的交易流水号
	OldFrontSeqNo           string // 原提现的见证系统流水号
	OldMarketSeqNo          string // 原提现的市场流水号
	OldAddMsg               string // 原提现的附言信息
	RejectBillReason        string // 退票原因
	RejectBillDate          string // 退票日期
	RejectInAcctTranSeqNo   string // 退票入账的交易流水号
	RejectInAcctTranAmt     string // 退票入账的交易金额
	RejectInPayerAcctNo     string // 退票入账的付款账号
	RejectInPayerAcctName   string // 退票入账的付款户名
	RejectInPayerBranchId   string // 退票入账的付款方行号
	RejectInPayerBranchName string // 退票入账的付款方行名
	PayeeWitnessSubAcctNo   string // 退票入账的收款方见证子账户
	PayeeFrontSeqNo         string // 退票入账的收款方见证系统流水号
	ReservedMsgOne          string // 保留域1
	ReservedMsgTwo          string // 保留域2
	ReservedMsgThree        string // 保留域3
}

type BankWithdrawCashBackQueryRsp struct {
	ResponseData
	ResultNum     string // 本次交易返回查询结果记录数
	ReservedMsg   string // 保留域
	TranItemArray []BankWithdrawCashBackQueryItem
}

type CommonTransferRechargeQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	FunctionFlag string // required(1), 0: 查询历史数据 1: 为查询当日数据
	StartDate    string // required(8), 开始日期, 格式: 20060102
	PageNum      string // required(6), 页码
	EndDate      string // required(8), 终止日期, 格式: 20060102
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type CommonTransferRechargeQueryItem struct {
	FrontSeqNo        string // 见证系统流水号
	SubAcctNo         string // 见证子帐户的帐号
	TranNetMemberCode string // 交易网会员代码
	InAcctType        string // 入账类型 02：会员充值/退票入账, 03：资金挂账
	InAcctNo          string // 入金账号
	InAcctName        string // 入金账户名称
	TranAmt           string // 入金金额
	Ccy               string // 币种
	BankName          string // 银行名称
	AccountingDate    string // 会计日期
	Remark            string // 转账备注
}

type CommonTransferRechargeQueryRsp struct {
	ResponseData
	EndFlag       string // 结束标志
	TotalNum      string // 符合业务查询条件的记录总数
	ResultNum     string // 本次交易返回查询结果记录数
	StartRecordNo string // 起始记录号
	ReservedMsg   string // 保留域
	TranItemArray []CommonTransferRechargeQueryItem
}

type BankTransactionDetailsQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo    string // required(32), 子账户账号
	FunctionFlag string // required(1), 1: 当日 2: 历史
	QueryFlag    string // required(1), 1: 全部 2: 转出 3: 转入
	PageNum      string // required(6), 页码
	StartDate    string `json:",omitempty"` // optional(8), 开始日期, 格式: 20060102, FunctionFlag = 2 必填
	EndDate      string `json:",omitempty"` // optional(8), 终止日期, 格式: 20060102,  FunctionFlag = 2 必填
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type BankTransactionDetailsQueryItem struct {
	BookingFlag string // 记账标志, 1：转出 2：转入
	TranStatus  string // 交易状态, 0：成功
	TranAmt     string // 交易金额, 包含手续费，即交易金额=实际到账金额+手续费
	TranDate    string // 交易日期
	TranTime    string // 交易时间
	FrontSeqNo  string // 见证系统流水号
	// 1：会员支付（6034-6/9分支;6101-6/9分支；6006-6/9分支）
	// 2：会员冻结（6007-1/5分支；6135-1/5分支；6134-1/5分支）
	// 3：会员解冻 （6007-2/6分支；6135-2/6分支；6134-2/6分支）
	// 4：登记挂账（6139）
	// 6：下单预支付 （6034-1分支；6101-1分支；6006-1分支）
	// 7：确认并付款 （6034-2分支；6101-2分支；6006-2分支；6163；6166；6165）
	// 8：会员退款 （6034-3分支；6101-3分支；6006-3分支）
	// 22：见证+收单平台调账（6145）
	// 23：见证+收单资金冻结  (Note字段放:见证+收单资金冻结,订单号)
	// 24：见证+收单资金解冻（6007-4分支；6135-4分支；6134-4分支）
	// 25：会员间交易退款（6164）
	// 33：在途充值解冻(6007-7分支)
	BookingType  string // 记账类型
	InSubAcctNo  string // 转入见证子账户的帐号
	OutSubAcctNo string // 转出见证子账户的帐号
	Remark       string // 备注
}

type BankTransactionDetailsQueryRsp struct {
	ResponseData
	EndFlag       string // 结束标志
	TotalNum      string // 符合业务查询条件的记录总数
	ResultNum     string // 本次交易返回查询结果记录数
	StartRecordNo string // 起始记录号
	ReservedMsg   string // 保留域
	TranItemArray []BankTransactionDetailsQueryItem
}

type BankWithdrawCashDetailsQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo   string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	SubAcctNo    string // required(32), 子账户账号
	FunctionFlag string // required(1), 1: 当日 2: 历史
	QueryFlag    string // required(1), 2: 提现 3: 清分
	BeginDate    string `json:",omitempty"` // optional(8), 开始日期, 格式: 20060102, FunctionFlag=2 必填
	EndDate      string `json:",omitempty"` // optional(8), 终止日期, 格式: 20060102, FunctionFlag=2 必填
	PageNum      string // required(6), 页码
	ReservedMsg  string `json:",omitempty"` // optional(120), 保留域
}

type BankWithdrawCashDetailsQueryItem struct {
	// 01:提现（6033；6085；6111；见证+收单退款）
	// 02:清分（6129；6216；6130；6217;见证+收单充值；
	BookingFlag       string // 记账标志
	TranStatus        string // 交易状态, 0：成功
	BookingMsg        string // 记账说明
	TranNetMemberCode string // 交易网会员代码
	SubAcctNo         string // 见证子帐户的帐号
	SubAcctName       string // 见证子账户的名称
	TranAmt           string // 交易金额, 包含手续费，即交易金额=实际到账金额+手续费
	Commission        string // 手续费
	TranDate          string // 交易日期
	TranTime          string // 交易时间
	FrontSeqNo        string // 见证系统流水号
	Remark            string // 备注, 如果是见证+收单的交易，返回交易订单号
}

type BankWithdrawCashDetailsQueryRsp struct {
	ResponseData
	EndFlag       string // 结束标志
	TotalNum      string // 符合业务查询条件的记录总数
	ResultNum     string // 本次交易返回查询结果记录数
	StartRecordNo string // 起始记录号
	ReservedMsg   string // 保留域
	TranItemArray []BankWithdrawCashDetailsQueryItem
}

type DetailVerifiedCodeQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo    string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	OldFrontSeqNo string // required(32), 原前置流水号
	OldTranType   string // required(2), 原交易类型, 默认上送1: 担保交易
	ReservedMsg   string `json:",omitempty"` // optional(120), 保留域
}

type DetailVerifiedCodeQueryRsp struct {
	ResponseData
	OldFrontSeqNo   string // 原前置流水号
	DetailCheckCode string // 明细单验证码
	ReservedMsg     string // 保留域
}

type ReconciliationDocumentQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	// 充值文件-CZ
	// 提现文件-TX
	// 交易文件-JY
	// 余额文件-YE
	// 鉴权文件-JQ
	// POS文件-POS
	// 资金汇总账户明细文件-JG
	// 平台归集账户明细文件-GJ
	FileType    string // required(2), 文件类型
	FileDate    string // required(10), 文件日期, 格式: 20060102
	ReservedMsg string `json:",omitempty"` // optional(120), 保留域
}

type ReconciliationDocumentQueryItem struct {
	FileName       string // 文件名称
	RandomPassword string // 随机密码
	FilePath       string // 文件路径
	DrawCode       string // 提取码
}

type ReconciliationDocumentQueryRsp struct {
	ResultNum     string // 本次交易返回查询结果记录数
	ReservedMsg   string // 保留域
	TranItemArray []ReconciliationDocumentQueryItem
}

type ChargeDetailQueryReq struct {
	// FundSummaryAcctNo string // required(32), 资金监管账号(底层已实现)
	CnsmrSeqNo           string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	AcquiringChannelType string // required(2), 收单渠道类型, 01-橙E收款, YST1-云收款
	OrderNo              string // required(30), 订单号(下单时的子订单号, 不是总订单号)
	ReservedMsg          string `json:",omitempty"` // optional(120), 保留域
}

type ChargeDetailQueryRsp struct {
	ResponseData
	TranStatus            string
	TranAmt               string
	CommissionAmt         string
	PayMode               string
	TranDate              string
	TranTime              string
	OrderInSubAcctNo      string
	OrderInSubAcctName    string
	OrderActInSubAcctNo   string
	OrderActInSubAcctName string
	FrontSeqNo            string
	TranDesc              string
}

type EJZBCustInformationQueryReq struct {
	CnsmrSeqNo  string // optional(22),交易网业务流水号(如果未设置,底层将自动生成一个)
	CustAcctId  string // required(32), 子台账账号
	ThirdCustId string // required(32), 交易网会员代码
	Reserve     string `json:",omitempty"` // optional(100), 保留域
}

type EJZBCustInformationQueryRsp struct {
	CustAcctId   string // 子台账账号
	ThirdCustId  string // 交易网会员代码
	CustName     string // 会员名称
	IdType       string // 会员证件类型
	IdCode       string // 会员证件号码
	ClientLvl    string // 会员属性
	NickName     string // 用户昵称
	MobilePhone  string // 手机号码
	BusinessFlag string // 个体工商户标志
	ComepanyName string // 公司名称
	CreditidCode string // 公司证件号码
	StoreId      string // 店铺编
	StoreName    string // 店铺名称
	LegalFlag    string // 法人标志
	LegalName    string // 法人名称
	CPFlag       string // 实名认证标识
	OrangePay    string // 是否有登记行为记录信息, S:是,F:否
	LecertiType  string // 法人证件类型
	LecertiCode  string // 法人证件号码
	Reserve      string // 保留域
}
