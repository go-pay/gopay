package paypal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建订单（Create order）
//	Code = 0 is success
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_create
func (c *Client) CreateOrder(ctx context.Context, bm gopay.BodyMap) (ppRsp *CreateOrderRsp, err error) {
	if err = bm.CheckEmptyError("intent", "purchase_units"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, orderCreate)
	if err != nil {
		return nil, err
	}
	ppRsp = &CreateOrderRsp{Code: Success}
	ppRsp.Response = new(OrderDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
	}
	return ppRsp, nil
}

// 更新订单（Update order）
//	Code = 0 is success
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_patch
func (c *Client) UpdateOrder(ctx context.Context, orderId string, patchs []*Patch) (ppRsp *EmptyRsp, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_is is empty")
	}
	url := fmt.Sprintf(orderUpdate, orderId)
	res, bs, err := c.doPayPalPatch(ctx, patchs, url)
	if err != nil {
		return nil, err
	}
	ppRsp = &EmptyRsp{Code: Success}
	if res.StatusCode != http.StatusNoContent {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
	}
	return ppRsp, nil
}

// 订单详情（Show order details）
//	Code = 0 is success
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_get
func (c *Client) OrderDetail(ctx context.Context, orderId string, bm gopay.BodyMap) (ppRsp *OrderDetailRsp, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_is is empty")
	}
	uri := fmt.Sprintf(orderDetail, orderId) + "?" + bm.EncodeURLParams()
	res, bs, err := c.doPayPalGet(ctx, uri)
	if err != nil {
		return nil, err
	}
	ppRsp = &OrderDetailRsp{Code: Success}
	ppRsp.Response = new(OrderDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
	}
	return ppRsp, nil
}

// 订单支付授权（Authorize payment for order）
//	Code = 0 is success
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_authorize
func (c *Client) OrderAuthorize(ctx context.Context, bm gopay.BodyMap) {

}

// 订单支付捕获（Capture payment for order）
//	Code = 0 is success
//	文档：https://developer.paypal.com/docs/api/orders/v2/#orders_capture
func (c *Client) OrderCapture(ctx context.Context, bm gopay.BodyMap) {

}
