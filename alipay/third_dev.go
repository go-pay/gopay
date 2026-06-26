package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.open.app.api.query(查询应用可申请的接口出参敏感字段列表)
// 文档地址：https://opendocs.alipay.com/isv/c80094bf_alipay.open.app.api.query
func (a *Client) OpenAppApiQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppApiQueryResponse, err error) {
	err = bm.CheckEmptyError(AppAuthToken)
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

// alipay.open.app.api.field.apply(申请获取接口用户敏感信息字段)
// 文档地址：https://opendocs.alipay.com/isv/373ff9fc_alipay.open.app.api.field.apply
func (a *Client) OpenAppApiFieldApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppApiFieldApplyRsp, err error) {
	err = bm.CheckEmptyError("auth_field_apply")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.app.api.field.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppApiFieldApplyRsp)
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

// alipay.open.app.api.scene.query(查询接口字段使用场景)
// 文档地址：https://opendocs.alipay.com/isv/54e1e005_alipay.open.app.api.scene.query
func (a *Client) OpenAppApiSceneQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppApiSceneQueryRsp, err error) {
	err = bm.CheckEmptyError("field_name", "api_name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.app.api.scene.query"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppApiSceneQueryRsp)
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

// alipay.open.app.api.field.query(用户信息申请记录查询)
// 文档地址：https://opendocs.alipay.com/isv/a0d374a3_alipay.open.app.api.field.query
func (a *Client) OpenAppApiFieldQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppApiFieldQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.app.api.field.query"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppApiFieldQueryRsp)
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

// alipay.open.appinfo.modify(应用信息修改接口)
// 文档地址：https://opendocs.alipay.com/isv/b272268f_alipay.open.appinfo.modify
func (a *Client) OpenAppInfoModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppInfoModifyRsp, err error) {
	err = bm.CheckEmptyError("open_id_config")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.appinfo.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppInfoModifyRsp)
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

// alipay.open.appinfo.query(应用信息查询接口)
// 文档地址：https://opendocs.alipay.com/isv/e5a89be0_alipay.open.appinfo.query
func (a *Client) OpenAppInfoQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppInfoQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.appinfo.query"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppInfoQueryRsp)
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
