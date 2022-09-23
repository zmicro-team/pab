package openpaa

import (
	"golang.org/x/exp/slices"
)

const (
	Success = "000000" // 成功
	Err936  = "ERR936" // 交易要素不全
	Err923  = "ERR923" // 数据库操作失败
	Err957  = "ERR957" // 交易异常,请稍候查询交易结果
	Err020  = "ERR020" // 无符合条件记录
	Err021  = "ERR021" // 错误的功能码
	Err022  = "ERR022" // 监管帐户必须是对公账号
	Err025  = "ERR025" // 主机返回错误
	Err026  = "ERR026" // 企业代码错误或未开通服务
	Err027  = "ERR027" // 所输监管账号未开户或已销户
	Err030  = "ERR030" // 请求页码错误
	Err031  = "ERR031" // 子账号不存在或已无效
	Err033  = "ERR033" // 绑定账号不存在
	Err042  = "ERR042" // 金额不能小于等于0
	Err044  = "ERR044" // 可用余额不足
	Err046  = "ERR046" // 银行类别错误
	Err047  = "ERR047" // 打开文件失败
	Err055  = "ERR055" // 输入账号错误
	Err059  = "ERR059" // 不允许发起此交易
	Err070  = "ERR070" // 交易网流水号不符
	Err074  = "ERR074" // 交易网返回失败描述
	Err075  = "ERR075" // 银行受理跨行转账失败
	Err089  = "ERR089" // 交易日期不符
	Err091  = "ERR091" // 交易网未签到
	Err095  = "ERR095" // 余额不为0不允许销户
	Err108  = "ERR108" // 系统处理失败
	Err111  = "ERR111" // 绑卡证件信息不符
	Err114  = "ERR114" // 该市场的交易网会员代码已存在
	Err115  = "ERR115" // 订单信息不存在
	Err116  = "ERR116" // 不在有效期内
	Err125  = "ERR125" // 账户名称不符
	Err126  = "ERR126" // 金额不符
	Err127  = "ERR127" // 预付卡号/订单号重复
	Err128  = "ERR128" // 身份证信息联网核查失败
	Err129  = "ERR129" // 监管落地审批信息不存在
	Err130  = "ERR130" // 无此交易权限
	Err131  = "ERR131" // 原流水号不存在
	Err132  = "ERR132" // 该功能暂未开通
	Err133  = "ERR133" // 单笔提现超过限额
	Err134  = "ERR134" // 账号已被绑定
	Err135  = "ERR135" // 短信动态码验证中，请稍后重新发起
	Err136  = "ERR136" // 系统繁忙中，请稍后重试
	Err137  = "ERR137" // 验证信息不存在
	Err138  = "ERR138" // 短信动态码验证不相符
	Err139  = "ERR139" // 短信发送失败
	Err140  = "ERR140" // 地区名称不正确，无法查询到对应县市
	Err141  = "ERR141" // 会员新绑定帐号与原账号户名不符
	Err142  = "ERR142" // 会员代码已存在
	Err143  = "ERR143" // 当前时间不允许发起此交易
	Err144  = "ERR144" // 该证件信息已签约过交易所
	Err145  = "ERR145" // 原有申请在待验证时间内，请稍后再发起
	Err146  = "ERR146" // 该帐号已被绑定，请更换银行帐号
	Err147  = "ERR147" // 时间已超过48小时，请重新发起
	Err148  = "ERR148" // 该流水号对应的交易已撤销
	Err149  = "ERR149" // 子账户与原订单不一致
	Err150  = "ERR150" // 证件信息或银行预留手机号不符
	Err151  = "ERR151" // 身份证格式错误
	Err152  = "ERR152" // 市场资金清算中，请稍后再试
	Err153  = "ERR153" // 企业代码不存在
	Err154  = "ERR154" // 参数未配置
	Err155  = "ERR155" // 子账户已实名不允许合并
	Err156  = "ERR156" // 未开通网上支付功能,无法验证
	Err157  = "ERR157" // 此卡号您已认证错误超过6次,请次日再试
	Err158  = "ERR158" // 其它错误
	Err159  = "ERR159" // 该证件已绑定其它子账户
	Err160  = "ERR160" // 支付指令号重复
	Err161  = "ERR161" // 支付指令号不存在或已处理
	Err162  = "ERR162" // 手机号码错误
	Err163  = "ERR163" // 短信指令号已失效或不存在
	Err164  = "ERR164" // 计算手续费失败
	Err165  = "ERR165" // 获取提现账户信息失败
	Err166  = "ERR166" // 获取联行号失败
	Err167  = "ERR167" // 获取联行号的城市码失败
	Err168  = "ERR168" // 不允许扣可提现的余额
	Err169  = "ERR169" // 批量次数超过限制
	Err170  = "ERR170" // 文件加解密失败
	Err171  = "ERR171" // 未开通理财服务
	Err172  = "ERR172" // 输入的理财客户号错误
	Err173  = "ERR173" // 赎回金额不能大于100万
	Err174  = "ERR174" // 鉴权失败：卡状态异常
	Err175  = "ERR175" // 户名不符
	Err176  = "ERR176" // 帐号不存在
	Err177  = "ERR177" // 冻结解冻相关交易金额必须大于0
	Err178  = "ERR178" // 交易初始化异常
	Err179  = "ERR179" // 该电商平台未配置参数信息
	Err180  = "ERR180" // 订单不存在
	Err181  = "ERR181" // 订单已存在,插入失败,事务回滚
	Err182  = "ERR182" // 获取平台流水号失败
	Err183  = "ERR183" // 提现关闭,请确认在途资金是否按规定进账
	Err184  = "ERR184" // 未开通橙e付,无法通过橙e付提现
	Err185  = "ERR185" // 橙e付开通失败,无法通过橙e付提现
	Err186  = "ERR186" // 赎回方式有误
	Err187  = "ERR187" // 门户预检查异常
	Err188  = "ERR188" // 门户预检查失败
	Err189  = "ERR189" // 银联鉴权异常
	Err190  = "ERR190" // 银联鉴权失败
	Err191  = "ERR191" // 平台帐号必须绑定对应本行一般结算户
	Err192  = "ERR192" // 未配置该会员属性
	Err193  = "ERR193" // 橙e门户未开通,无法理财开户
	Err194  = "ERR194" // 橙e付充值异常
	Err195  = "ERR195" // 橙e付充值失败
	Err196  = "ERR196" // 行内转账异常
	Err197  = "ERR197" // 行内转账失败
	Err198  = "ERR198" // 该子账户已绑定该账户
	Err199  = "ERR199" // 记录统计金额异常
	Err200  = "ERR200" // 记录日终余额异常
	Err201  = "ERR201" // __REQ__["Qydm"] + "-在途对账不平,请关注"]
	Err202  = "ERR202" // 对账异常
	Err203  = "ERR203" // 在途对账不平,请关注
	Err204  = "ERR204" // 行内转账异失败
	Err205  = "ERR205" // 橙e网(门户)注册异常
	Err206  = "ERR206" // 橙e网(门户)注册失败
	Err207  = "ERR207" // 橙e付提现异常
	Err208  = "ERR208" // 橙e付提现失败
	Err209  = "ERR209" // 橙e系统异常
	Err210  = "ERR210" // 未绑定橙e付
	Err211  = "ERR211" // 暂不支持查询全部
	Err212  = "ERR212" // 银联鉴权超级网银行号为空
	Err213  = "ERR213" // 超级网银行号对应的鉴权行号未配置
	Err214  = "ERR214" // 仍有未完成的订单不允许销户
	Err215  = "ERR215" // 橙E帐号开户未绑卡
	Err216  = "ERR216" // 橙E帐号开户未上传身份影像
	Err217  = "ERR217" // 橙E帐号待补充开户资料
	Err218  = "ERR218" // 帐号基本信息查询异常
	Err220  = "ERR220" // 第n对绑定关系不存在
	Err221  = "ERR221" // 该子帐号禁止提现
	Err222  = "ERR222" // 提现白名单验证失败
	Err223  = "ERR223" // 该子帐号禁止转账
	Err224  = "ERR224" // 转账白名单验证失败
	Err225  = "ERR225" // 文件下载失败
	Err226  = "ERR226" // 子交易调用失败
	Err227  = "ERR227" // 暂不支持该查询类型
	Err228  = "ERR228" // 功能子帐号异常
	Err229  = "ERR229" // 批次号已使用
	Err230  = "ERR230" // 文件笔数与实际读取不符
	Err231  = "ERR231" // 入库数据与文件首行不一致
	Err232  = "ERR232" // 代收付标志有误
	Err233  = "ERR233" // 代收代付金额不等
	Err234  = "ERR234" // 总分文件批次号不一致
	Err235  = "ERR235" // 绑卡信息不存在或不正常
	Err236  = "ERR236" // 超网和大小额行号必送其一
	Err237  = "ERR237" // 提现帐号不能是监管帐号
	Err238  = "ERR238" // 该会员已经绑定此帐号,请勿重复绑定
	Err239  = "ERR239" // 订单号不允许为空
	Err240  = "ERR240" // 金额不允许为空
	Err241  = "ERR241" // 出入金账户不能是监管账户
	Err242  = "ERR242" // 新监管账户余额不为零
	Err243  = "ERR243" // 账面余额不等于可用余额
	Err244  = "ERR244" // 原出入金账户不存在
	Err245  = "ERR245" // 设置监管标志失败
	Err246  = "ERR246" // 绑定关系变更失败
	Err247  = "ERR247" // 余额转移失败
	Err248  = "ERR248" // 余额转移异常,已登记冲正任务
	Err249  = "ERR249" // 分表参数有误
	Err250  = "ERR250" // 会员代码不符合要求
	Err251  = "ERR251" // 开橙e门户失败
	Err252  = "ERR252" // 仍有卡号未解绑不允许销户
	Err253  = "ERR253" // 平台未开通智能收款
	Err254  = "ERR254" // 会员代码不存在
	Err255  = "ERR255" // 收单渠道类型不正确
	Err256  = "ERR256" // 监管账户不正确
	Err257  = "ERR257" // 该鉴权流水还未同步结果
	Err258  = "ERR258" // 验证超过限定次数
	Err259  = "ERR259" // 通知网银失败
	Err260  = "ERR260" // 证件与姓名不一致
	Err261  = "ERR261" // 人行查询通道已关闭
	Err262  = "ERR262" // 新监管帐号不能为原结算帐号
	Err263  = "ERR263" // 金额提现中不能解绑
	Err264  = "ERR264" // 不存在来账鉴权流水
	Err265  = "ERR265" // 来账鉴权匹配中,请确认已转账验证
	Err266  = "ERR266" // 暂时只支持改为商户子账户
	Err267  = "ERR267" // 尚有收单渠道未解约,不能变更
	Err268  = "ERR268" // 非自营平台不能使用该分支
	Err269  = "ERR269" // 订单信息与上送信息不一致
	Err270  = "ERR270" // 原交易信息不存在
	Err271  = "ERR271" // 可退金额不足
	Err272  = "ERR272" // 原转出方信息不符
	Err273  = "ERR273" // 退款手续费须为0
	Err274  = "ERR274" // 出入金账户余额不足
	Err275  = "ERR275" // 第三方流水号与交易流水号重复,请核查
	Err276  = "ERR276" // 收单专属子账户不存在
	Err277  = "ERR277" // 是已经调过帐或者是已发起过退款
	Err278  = "ERR278" // 该笔充值当前入账帐号非该收单渠道子账户,不可调账
	Err279  = "ERR279" // 该笔清算流水不存在
	Err280  = "ERR280" // 只支持对账成功后发起调账
	Err281  = "ERR281" // 交易金额与原订单金额不一致
	Err282  = "ERR282" // 该订单号已经报送成功,无需进行平台补账
	Err283  = "ERR283" // 订单状态失败或异常
	Err284  = "ERR284" // 输入的交易金额与接口查询出的交易金额不一致
	Err285  = "ERR285" // 平台上送的渠道类型与接口返回不一致
	Err286  = "ERR286" // 内部调用交易504130在途充值失败
	Err287  = "ERR287" // 不支持的证件类型
	Err288  = "ERR288" // 备注字段为空,无法解析
	Err289  = "ERR289" // 客户注册任务仍在处理中,请稍后再尝试本请求
	Err290  = "ERR290" // 通过绑卡、或者存量客户通过补录接口注册数字口袋后再尝试申请,或者联系分行通过中台添加白名单
	Err291  = "ERR291" // 未配置白名单展示
	Err317  = "ERR317" // 订单可用余额不足

	E10002 = "E10002" // 无匹配的会员子账户或会员子账户已失效

	E80001 = "E80001" // 员信息鉴权失败
	E80011 = "E80011" // 待查询页码越界

	E90001 = "E90001" // 无此有效状态的大小额行号
)

const (
// OPEN-E-100074: 验签失败
// OPEN-E-100071: Token不存在
// OPEN-E-100073: Token有误，验证失败，请做开发者认证
)

type Error struct {
	ErrorCode    string
	ErrorMessage string
}

func (e *Error) Error() string {
	return "openpab: code(" + e.ErrorCode + "), msg(" + e.ErrorMessage + ")"
}

func NewError(code, msg string) *Error {
	return &Error{
		ErrorCode:    code,
		ErrorMessage: msg,
	}
}

// IsErrInvalidToken errorCode是否是token无效
func IsErrInvalidToken(errorCode string) bool {
	return slices.Contains([]string{"OPEN-E-100071", "OPEN-E-100073", "OPEN-E-100074"}, errorCode)
}

// ContainErr 是否包含指定的错误代码
func ContainErr(err error, errCode ...string) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	return slices.Contains(errCode, e.ErrorCode)
}

func ErrCodeText(err error) string {
	e, ok := err.(*Error)
	if !ok {
		return err.Error()
	}
	return e.ErrorMessage
}
