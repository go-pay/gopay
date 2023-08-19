package lakala

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// 创建线下支付订单
// 文档：https://payjp.lakala.com/docs/cn/#api-RetailPay-RetailMicroPay
func (c *Client) CreateRetailOrder(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *RetailPayRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(newRetailOrder, c.PartnerCode, orderId)
	bs, err := c.doPut(ctx, url, bm)
	if err != nil {
		return nil, err
	}
	rsp = new(RetailPayRsp)
	err = json.Unmarshal(bs, rsp)
	if err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, nil
}

// 创建线下QRCode支付单
// 文档：https://payjp.lakala.com/docs/cn/#api-RetailPay-RetailQRCode
func (c *Client) CreateRetailQRCodeOrder(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *PaymentRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(newRetailQrcodeOrder, c.PartnerCode, orderId)
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
