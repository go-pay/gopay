package lakala

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
)

// 创建JSAPI订单
// 文档：https://payjp.lakala.com/docs/cn/#api-JSApi-NewJSAPI
func (c *Client) NewJSAPI(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price", "channel"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(newJSAPI, c.PartnerCode, orderId)
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

// 创建Native JSAPI订单(offline)
// 文档：https://payjp.lakala.com/docs/cn/#api-JSApi-NativeJSAPI
func (c *Client) NewNativeJSApi(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price", "channel"); err != nil {
		return nil, err
	}
	if bm.Get("channel") == "Wechat" && bm.Get("appid") == "" {
		return nil, fmt.Errorf("wechat appid is empty")

	}
	if bm.Get("channel") == "Alipay" && bm.Get("customer_id") == "" {
		return nil, fmt.Errorf("alipay customer_id is empty")
	}
	url := fmt.Sprintf(newNativeJSAPI, c.PartnerCode, orderId)
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

// 微信JSAPI支付跳转页
// 文档：https://payjp.lakala.com/docs/cn/#api-JSApi-WxJSAPIPay
func (c *Client) WechatJSAPIPay(ctx context.Context, orderId, redirect string, directPay bool) (rsp *ErrorCode, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if redirect == gopay.NULL {
		return nil, fmt.Errorf("redirect is empty")
	}
	url := fmt.Sprintf(wechatJSAPIPay, c.PartnerCode, orderId)
	bs, err := c.doGet(ctx, url, fmt.Sprintf("redirect=%s&directpay=%v", redirect, directPay))
	if err != nil {
		return nil, err
	}
	rsp = new(ErrorCode)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 支付宝JSAPI支付跳转页
// 文档：https://payjp.lakala.com/docs/cn/#api-JSApi-AliJSAPIPay
func (c *Client) AlipayJSAPIPay(ctx context.Context, orderId, redirect string, directPay bool) (rsp *ErrorCode, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if redirect == gopay.NULL {
		return nil, fmt.Errorf("redirect is empty")
	}
	url := fmt.Sprintf(alipayJSAPIPay, c.PartnerCode, orderId)
	bs, err := c.doGet(ctx, url, fmt.Sprintf("redirect=%s&directpay=%v", redirect, directPay))
	if err != nil {
		return nil, err
	}
	rsp = new(ErrorCode)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// Alipay+ JSAPI支付跳转页
// 文档：https://payjp.lakala.com/docs/cn/#api-JSApi-ApsJSAPIPAY
func (c *Client) WebJSAPIPay(ctx context.Context, orderId, redirect string) (rsp *ErrorCode, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if redirect == gopay.NULL {
		return nil, fmt.Errorf("redirect is empty")
	}
	url := fmt.Sprintf(alipayPlusJSAPIPay, c.PartnerCode, orderId)
	bs, err := c.doGet(ctx, url, "redirect="+redirect)
	if err != nil {
		return nil, err
	}
	rsp = new(ErrorCode)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}
