package openpab

import (
	"context"
)

// Response 通用回复
type Response[T any] struct {
	Code       string // 错误码,表式成功返回
	Message    string // 错误信息
	Errors     Errors // 业务错误
	ExtendData *T     // 失败时返回
	Data       *T     // 成功时返回
}

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
	TxnReturnCode string // 返回码
	TxnReturnMsg  string // 返回信息
	CnsmrSeqNo    string // 交易流水号, 同输入
}

// R: request type
// T: response type
// Invoke 处理请求, req 只需要非公共字段, 底层已加入公共字段
func Invoke[R any, T any](c *Client, ctx context.Context, interId string, req *R) (*Response[T], error) {
	var result Response[T]

	err := c.Invoke(ctx, interId, req, &result)
	if err != nil {
		return nil, err
	}
	if result.Code == "" {
		return &result, nil
	}
	return &result, result.Errors
}

// R: request type
// T: response type
// Invoke2 处理请求, req 只需要非公共字段, 底层已加入公共字段
func Invoke2[R any, T any](c *Client, ctx context.Context, interId string, req *R) (*T, error) {
	result, err := Invoke[R, T](c, ctx, interId, req)
	if err != nil {
		return nil, err
	}
	if result.Code == "" {
		return result.Data, nil
	}
	return result.ExtendData, result.Errors
}
