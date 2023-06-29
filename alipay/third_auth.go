package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// alipay.open.auth.token.app(换取应用授权令牌)
// 文档地址：https://opendocs.alipay.com/isv/04h3uf
func (a *Client) OpenAuthTokenApp(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAuthTokenAppResponse, err error) {
	if bm.GetString("code") == util.NULL && bm.GetString("refresh_token") == util.NULL {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("grant_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.auth.token.app"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAuthTokenAppResponse)
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

// alipay.open.auth.token.app.query(查询某个应用授权AppAuthToken的授权信息)
// 文档地址：https://opendocs.alipay.com/isv/04hgcp
func (a *Client) OpenAuthTokenAppQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAuthTokenAppQueryResponse, err error) {
	err = bm.CheckEmptyError("app_auth_token")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.auth.token.app.query"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAuthTokenAppQueryResponse)
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

// alipay.open.auth.appauth.invite.create(ISV向商户发起应用授权邀约)
// 文档地址：https://opendocs.alipay.com/isv/06evao
func (a *Client) OpenAuthTokenAppInviteCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAuthTokenAppInviteCreateResponse, err error) {
	if err = bm.CheckEmptyError("auth_app_id"); err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.auth.appauth.invite.create"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAuthTokenAppInviteCreateResponse)
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
