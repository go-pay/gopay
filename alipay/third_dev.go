package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.open.app.api.query(查询应用可申请的接口出参敏感字段列表)
// 文档地址：https://opendocs.alipay.com/isv/04f74n
func (a *Client) OpenAppApiQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppApiQueryResponse, err error) {
	err = bm.CheckEmptyError("app_auth_token")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.app.api.query"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppApiQueryResponse)
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
