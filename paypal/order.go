package paypal

import (
	"context"

	"github.com/go-pay/gopay"
)

// 创建订单（Create order）
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_create
func (c *Client) CreateOrder(ctx context.Context, bm gopay.BodyMap) {

}

// 更新订单（Update order）
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_patch
func (c *Client) UpdateOrder(ctx context.Context, bm gopay.BodyMap) {

}

// 订单详情（Show order details）
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_get
func (c *Client) OrderDetail(ctx context.Context, bm gopay.BodyMap) {

}

// 订单支付授权（Authorize payment for order）
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_authorize
func (c *Client) OrderAuthorize(ctx context.Context, bm gopay.BodyMap) {

}

// 订单支付捕获（Capture payment for order）
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_capture
func (c *Client) OrderCapture(ctx context.Context, bm gopay.BodyMap) {

}
