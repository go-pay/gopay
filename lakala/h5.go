package lakala

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
)

// 创建H5支付单
// 文档：https://payjp.lakala.com/docs/cn/#api-MobileH5-NewMobileH5Pay
func (c *Client) CreateH5PayOrder(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *PaymentRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price", "channel"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(newH5Order, c.PartnerCode, orderId)
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

// H5支付跳转页
// 文档：https://payjp.lakala.com/docs/cn/#api-MobileH5-MobileH5Pay
func (c *Client) H5Pay(ctx context.Context, orderId, redirect string) (rsp *ErrorCode, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if redirect == gopay.NULL {
		return nil, fmt.Errorf("redirect is empty")
	}
	url := fmt.Sprintf(h5Pay, c.PartnerCode, orderId)
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

// H5支付跳转页(Alipay+)
// 文档：https://payjp.lakala.com/docs/cn/#api-MobileH5-MobileH5PayAlipayPlus
func (c *Client) H5AlipayPlusPay(ctx context.Context, orderId, redirect string) (rsp *ErrorCode, err error) {
	if orderId == gopay.NULL {
		return nil, errors.New("order_id is empty")
	}
	if redirect == gopay.NULL {
		return nil, fmt.Errorf("redirect is empty")
	}
	url := fmt.Sprintf(alipayPlusH5Pay, c.PartnerCode, orderId)
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
