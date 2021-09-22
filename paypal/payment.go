package paypal

import (
	"context"

	"github.com/go-pay/gopay"
)

// 支付授权详情（Show details for authorized payment）
//	文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_get
func (c *Client) PaymentAuthorizeDetail(ctx context.Context, bm gopay.BodyMap) {

}

// 支付授权捕获（Capture authorized payment）
//	文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
func (c *Client) PaymentAuthorizeCapture(ctx context.Context, bm gopay.BodyMap) {

}

// 重新授权支付授权（Reauthorize authorized payment）
//	Note：This request is currently not supported for Partner use cases.
//	文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_reauthorize
func (c *Client) PaymentReauthorize(ctx context.Context, bm gopay.BodyMap) {

}

// 作废支付授权（Void authorized payment）
//	文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_void
func (c *Client) PaymentAuthorizeVoid(ctx context.Context, bm gopay.BodyMap) {

}

// 支付捕获详情（Show captured payment details）
//	文档：https://developer.paypal.com/docs/api/payments/v2/#captures_get
func (c *Client) PaymentCaptureDetail(ctx context.Context, bm gopay.BodyMap) {

}

// 支付捕获退款（Refund captured payment）
//	文档：https://developer.paypal.com/docs/api/payments/v2/#captures_refund
func (c *Client) PaymentCaptureRefund(ctx context.Context, bm gopay.BodyMap) {

}

// 支付退款详情（Show refund details）
//	文档：https://developer.paypal.com/docs/api/payments/v2/#refunds_get
func (c *Client) PaymentRefundDetail(ctx context.Context, bm gopay.BodyMap) {

}
