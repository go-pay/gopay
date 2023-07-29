package lakala

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// 文档：https://payjp.lakala.com/docs/cn/#api-_
// QRCode - 创建QRCode支付单 1.2.0
// QRCode支付单适用于PC端网页/应用进行支付，用户使用微信/支付宝客户端扫描下单后生成的二维码完成支付。
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}
func (c *Client) CreateQRCode(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(ordersUrl, c.PartnerCode, orderId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(QRCodeRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 文档：https://payjp.lakala.com/docs/cn/#api-_
// QRCode - 创建Native QRCode支付单  1.2.0
// QRCode支付单适用于PC端网页/应用进行支付，用户使用微信/支付宝客户端扫描下单后生成的二维码完成支付。
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/native_orders/{order_id}
func (c *Client) CreateNativeQRCode(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(nativeOrdersUrl, c.PartnerCode, orderId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(QRCodeRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 文档：https://payjp.lakala.com/docs/cn/#api-_
// QRCode - QRCode支付跳转页 1.2.0
// 必须先调用创建QRCode订单接口再进行跳转。 建议在用户回调到对应页时通过后台查询订单状态接口确认订单的支付状态。 (Alipay/Wechat支持支付页面跳转,UnionPay不支持)
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}/pay
func (c *Client) GetQRCode(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("redirect"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(orderspayUrl, c.PartnerCode, orderId)
	_, bs, err := c.doGetParams(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(QRCodeRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}
