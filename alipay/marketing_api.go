package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.open.app.qrcode.create(小程序生成推广二维码接口)
// 文档地址：https://opendocs.alipay.com/apis/009zva
func (a *Client) OpenAppQrcodeCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppQrcodeCreateRsp, err error) {
	err = bm.CheckEmptyError("url_param", "query_param", "describe")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.app.qrcode.create"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppQrcodeCreateRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
