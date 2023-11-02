package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
)

// alipay.trade.customs.declare(统一收单报关接口)
// 文档地址：https://opendocs.alipay.com/apis/api_29/alipay.trade.customs.declare
func (a *Client) TradeCustomsDeclare(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCustomsDeclareRsp, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "merchant_customs_code", "merchant_customs_name", "amount", "customs_place")
	if err != nil {
		return nil, err
	}
	var bs []byte

	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.customs.declare"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCustomsDeclareRsp)
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

// alipay.acquire.customs(报关接口)
// 文档地址：https://opendocs.alipay.com/pre-open/01x3kh
func (a *Client) AcquireCustoms(ctx context.Context, bm gopay.BodyMap) (aliRspBs []byte, err error) {
	err = bm.CheckEmptyError("partner", "out_request_no", "trade_no", "merchant_customs_code", "amount", "customs_place", "merchant_customs_name")
	if err != nil {
		return nil, err
	}
	bs, err := a.doAliPayCustoms(ctx, bm, "alipay.acquire.customs")
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// alipay.overseas.acquire.customs.query(报关查询接口)
// 文档地址：https://opendocs.alipay.com/pre-open/01x3ki
func (a *Client) AcquireCustomsQuery(ctx context.Context, bm gopay.BodyMap) (aliRspBs []byte, err error) {
	err = bm.CheckEmptyError("partner", "out_request_nos")
	if err != nil {
		return nil, err
	}
	bs, err := a.doAliPayCustoms(ctx, bm, "alipay.overseas.acquire.customs.query")
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// 向支付宝发送请求
func (a *Client) doAliPayCustoms(ctx context.Context, bm gopay.BodyMap, service string) (bs []byte, err error) {
	bm.Set("service", service).
		Set("_input_charset", "utf-8")
	bm.Remove("sign_type")
	bm.Remove("sign")

	sign, err := a.getRsaSign(bm, RSA)
	if err != nil {
		return nil, fmt.Errorf("GetRsaSign Error: %v", err)
	}

	bm.Set("sign_type", RSA).Set("sign", sign)
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Request: %s", bm.JsonBody())
	}
	// request
	res, bs, err := a.hc.Req(xhttp.TypeFormData).Post("https://mapi.alipay.com/gateway.do").SendString(bm.EncodeURLParams()).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}
