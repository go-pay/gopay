package lakala

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
)

// 文档：https://payjp.lakala.com/docs/cn/#api-_
// JSApi - 创建JSAPI订单 1.2.0
// JSAPI适用于在微信/支付宝内打开的网页进行支付，如果用户从微信公众号进入支付页要求公众号已完成认证。用户下单后跳转至Lakala订单页，并拉起微信或支付宝内置收银台完成支付
// https://pay.lakala-japan.com/api/v1.0/jsapi_gateway/partners/{partner_code}/orders/{order_id}
func (c *Client) CreateJSApi(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(jsApiUrl, c.PartnerCode, orderId)
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
// JSApi - 创建Native JSAPI订单(offline) 1.2.0
// 接入公众号需完成主体认证，并且需要Lakala工作人员完成Appid绑定
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/native_jsapi/{order_id}
func (c *Client) CreateNativeJSApi(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price"); err != nil {
		return nil, err
	}
	if bm.Get("channel") == "Wechat" {
		if c.WxAppid == "" {
			return nil, fmt.Errorf("AppId null")
		}
		bm.Set("appid", c.WxAppid)
	}
	url := fmt.Sprintf(jsApiNativeUrl, c.PartnerCode, orderId)
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
// JSApi - 微信JSAPI支付跳转页 1.2.0
// 必须先调用创建JSAPI订单接口再进行跳转，最终URL以下单api返回的pay_url为准
// https://pay.lakala-japan.com/api/v1.0/wechat_jsapi_gateway/partners/{partner_code}_order_{order_id}
func (c *Client) GetJSApiWechat(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if err = bm.CheckEmptyError("redirect"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(jsApiWechatUrl, c.PartnerCode, orderId)
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

// 文档：https://payjp.lakala.com/docs/cn/#api-JSApi-NewJSAPI
// JSApi - 支付宝JSAPI支付跳转页 1.2.0
// 必须先调用创建JSAPI订单接口再进行跳转，最终URL以下单api返回的pay_url为准
// https://pay.lakala-japan.com/api/v1.0/gateway/alipay/partners/{partner_code}/orders/{order_id}/app_pay
func (c *Client) GetJSApiAlipay(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if err = bm.CheckEmptyError("redirect"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(jsApiAppPayUrl, c.PartnerCode, orderId)
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

// 文档：https://payjp.lakala.com/docs/cn/#api-_
// JSApi - Alipay+ JSAPI支付跳转页 1.3.0
// 必须先调用创建JSAPI支付订单接口再进行跳转，最终URL以下单api返回的pay_url为准
// https://pay.lakala-japan.com/api/v1.0/alipay_connect/partners/{partner_code}/orders/{order_id}/web_pay
func (c *Client) GetJSApiWebPay(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if err = bm.CheckEmptyError("redirect"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(jsApiWebPayUrl, c.PartnerCode, orderId)
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
