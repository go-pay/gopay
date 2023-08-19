package lakala

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// 创建SDK订单(Online)
// 文档：https://payjp.lakala.com/docs/cn/#api-SDKPayment-SDK
func (c *Client) CreateSDKPaymentOrder(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *PaymentRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price", "channel"); err != nil {
		return nil, err
	}
	if bm.Get("channel") == "Wechat" && bm.Get("appid") == "" {
		return nil, fmt.Errorf("wechat appid is empty")
	}
	url := fmt.Sprintf(newSDKPaymentOrder, c.PartnerCode, orderId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(PaymentRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}
