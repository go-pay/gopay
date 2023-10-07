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
	rsp = new(MiniPayRsp)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return rsp, c.verifySign(bs)
}
