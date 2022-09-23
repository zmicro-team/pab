package openpac

import (
	"context"
)

// Pay 支付
// TODO: 未实现
func (c *Client) Pay(ctx context.Context, req PayReq) (*PayRsp, error) {
	var result PayRsp

	err := c.InvokeCloudPay(ctx, "Pay", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// QuerySingleOrder 查询单笔订单
func (c *Client) QuerySingleOrder(ctx context.Context, req QuerySingleOrderReq) (*QuerySingleOrderRsp, error) {
	var result QuerySingleOrderRsp

	err := c.InvokeCloudPay(ctx, "QuerySingleOrder", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// RevokeOrder 撤销订单
func (c *Client) RevokeOrder(ctx context.Context, req RevokeOrderReq) (*RevokeOrderRsp, error) {
	var result RevokeOrderRsp

	err := c.InvokeCloudPay(ctx, "RevokeOrder", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// ApplyRefund 申请退款
func (c *Client) ApplyRefund(ctx context.Context, req ApplyRefundReq) (*ApplyRefundRsp, error) {
	var result ApplyRefundRsp

	err := c.InvokeCloudPay(ctx, "ApplyRefund", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// QuerySingleRefund 查询单笔退款
func (c *Client) QuerySingleRefund(ctx context.Context, req QuerySingleRefundReq) (*QuerySingleRefundRsp, error) {
	var result QuerySingleRefundRsp

	err := c.InvokeCloudPay(ctx, "QuerySingleRefund", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// SendShortMessage 发送短信接口
// 快捷支付时先发送短信接口，再进行支付。短信验证码验证提供短信验证码发送、验证流程，
// 短信验证码为6位随机数字，3分钟内有效，一分钟内同一手机号只能下发一次验证码。
// 同一个手机号限制一小时发6次，两次间隔必须大于一分钟，成功支付后对这个手机号会重置限制次数。
// 一个短信码有效期5分钟，验证成功后会失效，失败3次会失效，失效后需要重新获取验证码。短信码重发的话3分钟内验证码不变。
func (c *Client) SendShortMessage(ctx context.Context, req SendShortMessageReq) (*SendShortMessageRsp, error) {
	var result SendShortMessageRsp

	err := c.InvokeCloudPay(ctx, "SendShortMessage", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// DownloadReconciliationFile 下载对账单的文件路径、提取码、解压码
func (c *Client) DownloadReconciliationFile(ctx context.Context, req DownloadReconciliationFileReq) (*DownloadReconciliationFileRsp, error) {
	var result DownloadReconciliationFileRsp

	err := c.InvokeCloudPay(ctx, "DownloadReconciliationFile", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
