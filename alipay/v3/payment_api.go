package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 统一收单线下交易预创建
// Code = 200 is success
func (a *ClientV3) TradePrecreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradePrecreateRsp, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TradePrecreate, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradePrecreate, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradePrecreateRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}
