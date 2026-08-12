package cmbpay

import (
	"context"
)

// 本文件为常用接口提供类型化的便捷方法。若需调用未在此封装的接口，
// 可直接使用 Client.Execute / Client.ExecuteEncrypted 传入自定义结构体或 map。

// QrCodeApply 收款码申请（4.1）。
func (c *Client) QrCodeApply(ctx context.Context, req *QrCodeApplyReq) (*QrCodeApplyResp, error) {
	var resp QrCodeApplyResp
	if err := c.Execute(ctx, PathQrCodeApply, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QrCodeApplyEncrypted 收款码申请，并携带敏感字段加密所需的数字信封（4.1 + 2.4.4）。
// req 中的敏感字段（如 EncryptTerminalInfo、EncryptIdentity）密文需先由 enc.Encrypt 生成。
func (c *Client) QrCodeApplyEncrypted(ctx context.Context, req *QrCodeApplyReq, enc *Encryptor) (*QrCodeApplyResp, error) {
	var resp QrCodeApplyResp
	if err := c.ExecuteEncrypted(ctx, PathQrCodeApply, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderQuery 支付结果查询（4.2）。
func (c *Client) OrderQuery(ctx context.Context, req *OrderQueryReq) (*OrderQueryResp, error) {
	var resp OrderQueryResp
	if err := c.Execute(ctx, PathOrderQuery, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Refund 退款申请（4.4）。
func (c *Client) Refund(ctx context.Context, req *RefundReq) (*RefundResp, error) {
	var resp RefundResp
	if err := c.Execute(ctx, PathRefund, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefundQuery 退款结果查询（4.5）。
func (c *Client) RefundQuery(ctx context.Context, req *RefundQueryReq) (*RefundQueryResp, error) {
	var resp RefundQueryResp
	if err := c.Execute(ctx, PathRefundQuery, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Close 关闭订单（4.7）。
func (c *Client) Close(ctx context.Context, req *CloseReq) (*CloseResp, error) {
	var resp CloseResp
	if err := c.Execute(ctx, PathClose, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Pay 付款码收款（4.8）。若 req 含加密后的终端信息，请改用 PayEncrypted。
func (c *Client) Pay(ctx context.Context, req *PayReq) (*PayResp, error) {
	var resp PayResp
	if err := c.Execute(ctx, PathPay, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayEncrypted 付款码收款，携带敏感字段加密数字信封（4.8 + 2.4.4）。
func (c *Client) PayEncrypted(ctx context.Context, req *PayReq, enc *Encryptor) (*PayResp, error) {
	var resp PayResp
	if err := c.ExecuteEncrypted(ctx, PathPay, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Cancel 付款码支付撤销（4.9）。
func (c *Client) Cancel(ctx context.Context, req *CancelReq) (*CancelResp, error) {
	var resp CancelResp
	if err := c.Execute(ctx, PathCancel, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OnlinePay 微信统一下单（4.10）。
func (c *Client) OnlinePay(ctx context.Context, req *OnlinePayReq) (*OnlinePayResp, error) {
	var resp OnlinePayResp
	if err := c.Execute(ctx, PathOnlinePay, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
