package lakala

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// 文档：https://payjp.lakala.com/docs/cn/#api-Channel_Web_Gateway-WebGateway
/*
SDKPayment - 创建SDK订单(Online) 1.3.0
用于移动端APP调用支付渠道SDK支付，调用API创建订单，得到SDK调用参数，将参数传递给SDK拉起对应渠道App(微信、支付宝)支付，并由渠道App直接返回支付结果。 强烈建议获得支付结果后再调用Lakala订单查询API确认完成支付后再进行后续流程，避免因超时自动撤单导致资金损失

AlipayPlus的Bundle模式和Tile模式：

Bundle模式：即先跳转Alipay+页面再由消费者选择付款钱包
Tile模式：先调用查询可用钱包接口获取可用钱包供消费者选择，下单后直接跳转对应钱包支付。Tile模式需要传递extra.pay_type参数指定钱包

*/
// https://pay.lakala-japan.com/api/v1.0/gateway/partners/{partner_code}/app_orders/{order_id}
func (c *Client) CreateAppOrders(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *SDKPaymentRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price", "channel"); err != nil {
		return nil, err
	}
	if bm.Get("channel") == "Wechat" {
		if c.WxAppid == "" {
			return nil, fmt.Errorf("AppId null")
		}
		bm.Set("appid", c.WxAppid)
	}
	url := fmt.Sprintf(appOrdersUrl, c.PartnerCode, orderId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(SDKPaymentRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}
