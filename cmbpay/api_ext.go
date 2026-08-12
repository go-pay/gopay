package cmbpay

import "context"

// 本文件为其余接口提供类型化方法。凡请求含需加密的敏感字段（如 encryptTerminalInfo、
// encryptIdentity），先用 Encryptor 加密字段值，再调用对应的 *Encrypted 方法。

// ServPay 服务窗支付（4.11）。
func (c *Client) ServPay(ctx context.Context, req *ServPayReq) (*ServPayResp, error) {
	var resp ServPayResp
	if err := c.Execute(ctx, PathServPay, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ServPayEncrypted 服务窗支付，携带敏感字段加密数字信封（4.11 + 2.4.4）。
func (c *Client) ServPayEncrypted(ctx context.Context, req *ServPayReq, enc *Encryptor) (*ServPayResp, error) {
	var resp ServPayResp
	if err := c.ExecuteEncrypted(ctx, PathServPay, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZfbNative 支付宝 native 码支付（4.12）。
func (c *Client) ZfbNative(ctx context.Context, req *ZfbNativeReq) (*ZfbNativeResp, error) {
	var resp ZfbNativeResp
	if err := c.Execute(ctx, PathZfbQrCode, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZfbNativeEncrypted 支付宝 native 码支付，携带敏感字段加密数字信封（4.12 + 2.4.4）。
func (c *Client) ZfbNativeEncrypted(ctx context.Context, req *ZfbNativeReq, enc *Encryptor) (*ZfbNativeResp, error) {
	var resp ZfbNativeResp
	if err := c.ExecuteEncrypted(ctx, PathZfbQrCode, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderQrCodeApply 订单二维码申请（4.14）。
func (c *Client) OrderQrCodeApply(ctx context.Context, req *OrderQrCodeApplyReq) (*OrderQrCodeApplyResp, error) {
	var resp OrderQrCodeApplyResp
	if err := c.Execute(ctx, PathOrderQrCodeApply, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderQrCodeApplyEncrypted 订单二维码申请，携带敏感字段加密数字信封（4.14 + 2.4.4）。
func (c *Client) OrderQrCodeApplyEncrypted(ctx context.Context, req *OrderQrCodeApplyReq, enc *Encryptor) (*OrderQrCodeApplyResp, error) {
	var resp OrderQrCodeApplyResp
	if err := c.ExecuteEncrypted(ctx, PathOrderQrCodeApply, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MiniAppOrder 微信小程序下单（4.15）。
func (c *Client) MiniAppOrder(ctx context.Context, req *MiniAppOrderReq) (*MiniAppOrderResp, error) {
	var resp MiniAppOrderResp
	if err := c.Execute(ctx, PathMiniAppOrderApply, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MiniAppOrderEncrypted 微信小程序下单，携带敏感字段加密数字信封（4.15 + 2.4.4）。
func (c *Client) MiniAppOrderEncrypted(ctx context.Context, req *MiniAppOrderReq, enc *Encryptor) (*MiniAppOrderResp, error) {
	var resp MiniAppOrderResp
	if err := c.ExecuteEncrypted(ctx, PathMiniAppOrderApply, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloudPay 银联云闪付（4.16）。
func (c *Client) CloudPay(ctx context.Context, req *CloudPayReq) (*CloudPayResp, error) {
	var resp CloudPayResp
	if err := c.Execute(ctx, PathCloudPay, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EcnyUnifiedOrder 数字人民币统一下单（4.17）。
func (c *Client) EcnyUnifiedOrder(ctx context.Context, req *EcnyUnifiedOrderReq) (*EcnyUnifiedOrderResp, error) {
	var resp EcnyUnifiedOrderResp
	if err := c.Execute(ctx, PathEcnyUnifiedOrder, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EcnyContractUnifiedOrder 数字人民币统一下单-带合约（4.21）。请求需上送 ContractReq。
func (c *Client) EcnyContractUnifiedOrder(ctx context.Context, req *EcnyUnifiedOrderReq) (*EcnyUnifiedOrderResp, error) {
	var resp EcnyUnifiedOrderResp
	if err := c.Execute(ctx, PathEcnyContractUnifiedOrder, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EcnyUnifiedPayment 数字人民币统一支付（4.18）。
func (c *Client) EcnyUnifiedPayment(ctx context.Context, req *EcnyUnifiedPaymentReq) (*EcnyUnifiedPaymentResp, error) {
	var resp EcnyUnifiedPaymentResp
	if err := c.Execute(ctx, PathEcnyUnifiedPayment, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EcnySubwalletPay 数字人民币子钱包支付（4.19）。
func (c *Client) EcnySubwalletPay(ctx context.Context, req *EcnySubwalletPayReq) (*EcnySubwalletPayResp, error) {
	var resp EcnySubwalletPayResp
	if err := c.Execute(ctx, PathEcnySubwalletPay, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EcnyContractSubwalletPay 数字人民币子钱包支付-带合约（4.20）。请求需上送 ContractReq。
func (c *Client) EcnyContractSubwalletPay(ctx context.Context, req *EcnySubwalletPayReq) (*EcnySubwalletPayResp, error) {
	var resp EcnySubwalletPayResp
	if err := c.Execute(ctx, PathEcnyContractSubwalletPay, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZfbApp 支付宝 APP 支付（4.24）。
func (c *Client) ZfbApp(ctx context.Context, req *ZfbAppReq) (*ZfbAppResp, error) {
	var resp ZfbAppResp
	if err := c.Execute(ctx, PathZfbApp, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZfbAppEncrypted 支付宝 APP 支付，携带敏感字段加密数字信封（4.24 + 2.4.4）。
func (c *Client) ZfbAppEncrypted(ctx context.Context, req *ZfbAppReq, enc *Encryptor) (*ZfbAppResp, error) {
	var resp ZfbAppResp
	if err := c.ExecuteEncrypted(ctx, PathZfbApp, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZfbWap 支付宝手机网站支付（4.25）。
func (c *Client) ZfbWap(ctx context.Context, req *ZfbWapReq) (*ZfbWapResp, error) {
	var resp ZfbWapResp
	if err := c.Execute(ctx, PathZfbWap, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZfbWapEncrypted 支付宝手机网站支付，携带敏感字段加密数字信封（4.25 + 2.4.4）。
func (c *Client) ZfbWapEncrypted(ctx context.Context, req *ZfbWapReq, enc *Encryptor) (*ZfbWapResp, error) {
	var resp ZfbWapResp
	if err := c.ExecuteEncrypted(ctx, PathZfbWap, req, enc, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PapQuery 微信委托代扣查询（4.23）。
func (c *Client) PapQuery(ctx context.Context, req *PapQueryReq) (*PapQueryResp, error) {
	var resp PapQueryResp
	if err := c.Execute(ctx, PathPapOrderQuery, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OpenIDQuery 微信授权码查询 openid（4.37）。
func (c *Client) OpenIDQuery(ctx context.Context, req *OpenIDQryReq) (*OpenIDQryResp, error) {
	var resp OpenIDQryResp
	if err := c.Execute(ctx, PathOpenIDQryByAC, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
