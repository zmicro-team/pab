package openpaa

import (
	"context"
)

// RequestData 请求数据里的公共字段
type RequestData struct {
	TxnCode     string // required(20), 交易码/接口id
	TxnClientNo string `json:",omitempty"` // optional(20), 交易客户码
	CnsmrSeqNo  string // required(22), 交易流水号
	TxnTime     string // required(22), 发送时间,格式: YYYYMMDDHHmmSSNNN, 后三位固定000(底层已实现)
	MrchCode    string // required(22), 商户号, 签约客户号, 见证宝产品必填(底层已实现)
}

// ResponseData 回复数据域的公共字段
type ResponseData struct {
	TxnReturnCode string `json:"TxnReturnCode"`
	TxnReturnMsg  string `json:"TxnReturnMsg"`
	CnsmrSeqNo    string `json:"CnsmrSeqNo"`

	RsaSign         string `json:"RsaSign"`
	TokenExpiryFlag string `json:"tokenExpiryFlag"`
}

// R: request type
// T: response type
// InvokeAny 处理请求, req 只需要非公共字段, 底层已加入公共字段
func Invoke[R any, T any](c *Client, ctx context.Context, interId string, req *R) (*T, error) {
	var result T

	err := c.Invoke(ctx, interId, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
