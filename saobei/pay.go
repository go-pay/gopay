package saobei

// 支付2.0接口

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// MiniPay 小程序支付接口 https://help.lcsw.cn/xrmpic/tisnldchblgxohfl/rinsc3#title-node17
func (c *Client) MiniPay(ctx context.Context, bm gopay.BodyMap) (rsp *MiniPayRsp, err error) {
	err = bm.CheckEmptyError("pay_type", "terminal_ip", "terminal_trace", "terminal_time", "total_fee", "sub_appid", "open_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, miniPayPath, bm); err != nil {
		return nil, err
	}
	rsp = new(MiniPayRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, c.verifySign(bs)
}

// BarcodePay 付款码支付(扫码支付) https://help.lcsw.cn/xrmpic/tisnldchblgxohfl/rinsc3#title-node14
func (c *Client) BarcodePay(ctx context.Context, bm gopay.BodyMap) (rsp *BarcodePayRsp, err error) {
	err = bm.CheckEmptyError("pay_type", "terminal_ip", "terminal_trace", "terminal_time", "total_fee", "auth_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, barcodePayPath, bm); err != nil {
		return nil, err
	}
	rsp = new(BarcodePayRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, nil
}

// Query 支付查询 https://help.lcsw.cn/xrmpic/tisnldchblgxohfl/rinsc3#title-node18
func (c *Client) Query(ctx context.Context, bm gopay.BodyMap) (rsp *QueryRsp, err error) {
	err = bm.CheckEmptyError("pay_type", "terminal_trace", "terminal_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, queryPath, bm); err != nil {
		return nil, err
	}
	rsp = new(QueryRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, c.verifySign(bs)
}

// Refund 退款申请 https://help.lcsw.cn/xrmpic/tisnldchblgxohfl/rinsc3#title-node19
func (c *Client) Refund(ctx context.Context, bm gopay.BodyMap) (rsp *RefundRsp, err error) {
	err = bm.CheckEmptyError("pay_type", "terminal_trace", "terminal_time", "refund_fee")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, refundPath, bm); err != nil {
		return nil, err
	}
	rsp = new(RefundRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, c.verifySign(bs)
}

// QueryRefund 退款订单查询 https://help.lcsw.cn/xrmpic/tisnldchblgxohfl/rinsc3#title-node22
func (c *Client) QueryRefund(ctx context.Context, bm gopay.BodyMap) (rsp *QueryRefundRsp, err error) {
	err = bm.CheckEmptyError("pay_type", "terminal_trace", "terminal_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, queryRefundPath, bm); err != nil {
		return nil, err
	}
	rsp = new(QueryRefundRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, c.verifySign(bs)
}
