package lakala

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// 创建QRCode支付单
// 文档：https://payjp.lakala.com/docs/cn/#api-QRCode-NewQRCode
func (c *Client) NewQRCode(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price", "channel"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(newQrcode, c.PartnerCode, orderId)
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

// 创建Native QRCode支付单
// 文档：https://payjp.lakala.com/docs/cn/#api-QRCode-NativeQRCode
func (c *Client) NewNativeQRCode(ctx context.Context, orderId string, bm gopay.BodyMap) (rsp *QRCodeRsp, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if err = bm.CheckEmptyError("description", "price", "channel"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf(nativeQrcode, c.PartnerCode, orderId)
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

// QRCode支付跳转页
// 文档：https://payjp.lakala.com/docs/cn/#api-QRCode-QRCodePay
func (c *Client) QRCodePay(ctx context.Context, orderId, redirect string) (rsp *ErrorCode, err error) {
	if orderId == gopay.NULL {
		return nil, fmt.Errorf("orderId is empty")
	}
	if redirect == gopay.NULL {
		return nil, fmt.Errorf("redirect is empty")
	}
	url := fmt.Sprintf(qrcodePay, c.PartnerCode, orderId)
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
