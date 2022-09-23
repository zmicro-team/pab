package openpac

type ResponseData struct {
	TxnReturnCode string `json:"TxnReturnCode"` // required(20), 映射返回码
	TxnReturnMsg  string `json:"TxnReturnMsg"`  // required(100), 映射返回信息
	CnsmrSeqNo    string `json:"CnsmrSeqNo"`    // required(22), 系统流水号，同输入
	// TokenExpiryFlag bool   `json:"tokenExpiryFlag"` // required(50), token过期标识(底层处理), true: 过期
}
