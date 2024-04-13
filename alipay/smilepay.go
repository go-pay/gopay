package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// zoloz.authentication.smilepay.initialize(刷脸支付初始化)
// 文档地址：https://opendocs.alipay.com/open/2f7c1d5f_zoloz.authentication.smilepay.initialize
func (a *Client) ZolozAuthenticationSmilepayInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZolozAuthenticationSmilepayInitializeRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zoloz.authentication.smilepay.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(ZolozAuthenticationSmilepayInitializeRsp)
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

// zoloz.authentication.customer.ftoken.query(查询刷脸结果信息接口)
// 文档地址：https://opendocs.alipay.com/open/c8e4d285_zoloz.authentication.customer.ftoken.query
func (a *Client) ZolozAuthenticationCustomerFtokenQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZolozAuthenticationCustomerFtokenQueryRsp, err error) {
	err = bm.CheckEmptyError("ftoken", "biz_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zoloz.authentication.customer.ftoken.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZolozAuthenticationCustomerFtokenQueryRsp)
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
