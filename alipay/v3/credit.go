package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 芝麻GO签约预创单 zhima.credit.pe.zmgo.preorder.create
// StatusCode = 200 is success
func (a *ClientV3) ZmGoPreorderCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoPreorderCreateRsp, err error) {
	err = bm.CheckEmptyError("partner_id", "template_id", "out_request_no", "biz_time")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPut, v3ZmGoPreorderCreate, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPut(ctx, bm, v3ZmGoPreorderCreate, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoPreorderCreateRsp{StatusCode: res.StatusCode}
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

// 商家芝麻GO累计数据回传接口 zhima.merchant.zmgo.cumulate.sync
// StatusCode = 200 is success
func (a *ClientV3) ZmGoCumulateSync(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoCumulateSyncRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "provider_pid", "out_biz_no", "biz_time", "biz_action", "sub_biz_action", "data_type")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3ZmGoCumulateSync, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZmGoCumulateSync, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoCumulateSyncRsp{StatusCode: res.StatusCode}
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

// 商家芝麻GO累计数据查询接口 zhima.merchant.zmgo.cumulate.query
// StatusCode = 200 is success
func (a *ClientV3) ZmGoCumulateQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoCumulateQueryRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "provider_pid")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3ZmGoCumulateQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoCumulateQueryRsp{StatusCode: res.StatusCode}
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

// 芝麻GO结算申请 zhima.credit.pe.zmgo.settle.apply
// StatusCode = 200 is success
func (a *ClientV3) ZmGoSettleApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoSettleApplyRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "partner_id", "out_request_no", "withhold_plan_no", "pay_amount")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3ZmGoSettleApply, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZmGoSettleApply, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoSettleApplyRsp{StatusCode: res.StatusCode}
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

// 芝麻GO结算退款 zhima.credit.pe.zmgo.settle.refund
// StatusCode = 200 is success
func (a *ClientV3) ZmGoSettleRefund(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoSettleRefundRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "partner_id", "refund_amount", "out_request_no")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3ZmGoSettleApplyRefund, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZmGoSettleApplyRefund, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoSettleRefundRsp{StatusCode: res.StatusCode}
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

// 芝麻Go协议查询接口 zhima.credit.pe.zmgo.agreement.query
// StatusCode = 200 is success
func (a *ClientV3) ZmGoAgreementQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoAgreementQueryRsp, err error) {
	err = bm.CheckEmptyError("agreement_id")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3ZmGoAgreementQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoAgreementQueryRsp{StatusCode: res.StatusCode}
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

// 芝麻GO协议解约 zhima.credit.pe.zmgo.agreement.unsign
// StatusCode = 200 is success
func (a *ClientV3) ZmGoAgreementQueryUnsign(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoAgreementQueryUnsignRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "partner_id")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3ZmGoAgreementQueryUnsign, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZmGoAgreementQueryUnsign, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoAgreementQueryUnsignRsp{StatusCode: res.StatusCode}
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

// 商户创建芝麻GO模板接口 zhima.merchant.zmgo.template.create
// StatusCode = 200 is success
func (a *ClientV3) ZmGoTemplateCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoTemplateCreateRsp, err error) {
	err = bm.CheckEmptyError("basic_config", "right_config", "open_config", "settlement_config")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3ZmGoTemplateCreate, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZmGoTemplateCreate, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoTemplateCreateRsp{StatusCode: res.StatusCode}
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

// 商家芝麻GO模板查询 zhima.merchant.zmgo.template.query
// StatusCode = 200 is success
func (a *ClientV3) ZmGoTemplateQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZmGoTemplateQueryRsp, err error) {
	err = bm.CheckEmptyError("template_no", "partner_id")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3ZmGoTemplateQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZmGoTemplateQueryRsp{StatusCode: res.StatusCode}
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
