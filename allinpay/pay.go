package allinpay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// Pay 统一支付接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=88
func (c *Client) Pay(ctx context.Context, bm gopay.BodyMap) (rsp *PayRsp, err error) {
	err = bm.CheckEmptyError("reqsn", "paytype")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, payPath, bm); err != nil {
		return nil, err
	}
	rsp = new(PayRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	rsp = new(PayRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, c.verifySign(bs)
}

// ScanPay 统一扫码接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=89
func (c *Client) ScanPay(ctx context.Context, bm gopay.BodyMap) (rsp *ScanPayRsp, err error) {
	err = bm.CheckEmptyError("reqsn", "authcode", "terminfo")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = c.doPost(ctx, scanQrPath, bm); err != nil {
		return nil, err
	}
	rsp = new(ScanPayRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, c.verifySign(bs)
}

// Query 统一查询接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=836
func (c *Client) Query(ctx context.Context, orderType string, no string) (rsp *ScanPayRsp, err error) {
	bm := gopay.BodyMap{}
	switch orderType {
	case OrderTypeReqSN:
		bm.Set("reqsn", no)
	case OrderTypeTrxId:
		bm.Set("trxid", no)
	}
	var bs []byte
	if bs, err = c.doPost(ctx, queryPath, bm); err != nil {
		return nil, err
	}
	rsp = new(ScanPayRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, nil
}

// Refund 统一退款接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=838
func (c *Client) Refund(ctx context.Context, bm gopay.BodyMap) (rsp *RefundRsp, err error) {
	err = bm.CheckEmptyError("reqsn", "trxamt")
	if err != nil {
		return nil, err
	}
	if bm.GetString("oldreqsn") == gopay.NULL && bm.GetString("oldtrxid") == gopay.NULL {
		return nil, fmt.Errorf("[%w], %v", gopay.MissParamErr, "oldreqsn和oldtrxid必填其一")
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
	return rsp, nil
}

// Cancel 统一撤销接口 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=837
func (c *Client) Cancel(ctx context.Context, bm gopay.BodyMap) (rsp *RefundRsp, err error) {
	err = bm.CheckEmptyError("reqsn", "trxamt")
	if err != nil {
		return nil, err
	}
	if bm.GetString("oldreqsn") == gopay.NULL && bm.GetString("oldtrxid") == gopay.NULL {
		return nil, fmt.Errorf("[%w], %v", gopay.MissParamErr, "oldreqsn和oldtrxid必填其一")
	}
	var bs []byte
	if bs, err = c.doPost(ctx, cancelPath, bm); err != nil {
		return nil, err
	}
	rsp = new(RefundRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, nil
}

// Close 订单关闭 https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=424
func (c *Client) Close(ctx context.Context, bm gopay.BodyMap) (rsp *CloseRsp, err error) {
	if bm.GetString("oldreqsn") == gopay.NULL && bm.GetString("oldtrxid") == gopay.NULL {
		return nil, fmt.Errorf("[%w], %v", gopay.MissParamErr, "oldreqsn和oldtrxid必填其一")
	}
	var bs []byte
	if bs, err = c.doPost(ctx, closePath, bm); err != nil {
		return nil, err
	}
	rsp = new(CloseRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err := bizErrCheck(rsp.RspBase); err != nil {
		return nil, err
	}
	return rsp, nil
}
