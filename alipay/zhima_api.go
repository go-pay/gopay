package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// Deprecated
// zhima.credit.score.get(查询芝麻用户的芝麻分)
// 文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.score.get
func (a *Client) ZhimaCreditScoreGet(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditScoreGetResponse, err error) {
	if bm.GetString("product_code") == gopay.NULL {
		bm.Set("product_code", "w1010100100000000001")
	}
	err = bm.CheckEmptyError("transaction_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.score.get"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditScoreGetResponse)
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

// zhima.credit.ep.scene.rating.initialize(芝麻企业信用信用评估初始化)
// 文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.rating.initialize
func (a *Client) ZhimaCreditEpSceneRatingInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneRatingInitializeRsp, err error) {
	if bm.GetString("product_code") == gopay.NULL {
		bm.Set("product_code", "w1010100100000000001")
	}
	err = bm.CheckEmptyError("credit_category", "out_order_no", "user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.ep.scene.rating.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneRatingInitializeRsp)
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

// zhima.credit.ep.scene.fulfillment.sync(信用服务履约同步)
// 文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.fulfillment.sync
func (a *Client) ZhimaCreditEpSceneFulfillmentSync(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneFulfillmentSyncRsp, err error) {
	err = bm.CheckEmptyError("credit_order_no", "out_order_no", "biz_time", "biz_ext_param")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.ep.scene.fulfillment.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneFulfillmentSyncRsp)
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

// zhima.credit.ep.scene.agreement.use(加入信用服务)
// 文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.agreement.use
func (a *Client) ZhimaCreditEpSceneAgreementUse(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneAgreementUseRsp, err error) {
	err = bm.CheckEmptyError("rating_order_no", "out_order_no", "biz_time", "provision_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.ep.scene.agreement.use"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneAgreementUseRsp)
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

// zhima.credit.ep.scene.agreement.cancel(取消信用服务)
// 文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.agreement.cancel
func (a *Client) ZhimaCreditEpSceneAgreementCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneAgreementCancelRsp, err error) {
	err = bm.CheckEmptyError("credit_order_no", "out_order_no", "biz_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.ep.scene.agreement.cancel"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneAgreementCancelRsp)
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

// zhima.credit.ep.scene.fulfillmentlist.sync(信用服务履约同步(批量))
// 文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.fulfillmentlist.sync
func (a *Client) ZhimaCreditEpSceneFulfillmentlistSync(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneFulfillmentlistSyncRsp, err error) {
	err = bm.CheckEmptyError("credit_order_no", "fulfillment_info_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.ep.scene.fulfillmentlist.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneFulfillmentlistSyncRsp)
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

// zhima.credit.pe.zmgo.cumulation.sync(芝麻go用户数据回传)
// 文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.pe.zmgo.cumulation.sync
func (a *Client) ZhimaCreditPeZmgoCumulationSync(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoCumulationSyncRsp, err error) {
	err = bm.CheckEmptyError("agreement_no", "user_id", "partner_id", "out_biz_no", "biz_time", "request_from", "biz_action", "cumulate_data_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.cumulation.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoCumulationSyncRsp)
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

// zhima.merchant.zmgo.cumulate.sync(商家芝麻GO累计数据回传接口)
// 文档地址：https://opendocs.alipay.com/apis/01ol9h
func (a *Client) ZhimaMerchantZmgoCumulateSync(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaMerchantZmgoCumulateSyncRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "user_id", "provider_pid", "out_biz_no", "biz_time", "biz_action", "sub_biz_action", "data_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.merchant.zmgo.cumulate.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaMerchantZmgoCumulateSyncRsp)
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

// zhima.merchant.zmgo.cumulate.query(商家芝麻GO累计数据查询接口)
// 文档地址：https://opendocs.alipay.com/open/03ui2q
func (a *Client) ZhimaMerchantZmgoCumulateQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaMerchantZmgoCumulateQueryRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "user_id", "provider_pid")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.merchant.zmgo.cumulate.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaMerchantZmgoCumulateQueryRsp)
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

// zhima.merchant.zmgo.template.create(商户创建芝麻GO模板接口)
// 文档地址：https://opendocs.alipay.com/open/03uq08
func (a *Client) ZhimaMerchantZmgoTemplateCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaMerchantZmgoTemplateCreateRsp, err error) {
	err = bm.CheckEmptyError("basic_config", "right_config", "open_config", "settlement_config")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.merchant.zmgo.template.create"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaMerchantZmgoTemplateCreateRsp)
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

// zhima.merchant.zmgo.template.query(芝麻GO模板查询)
// 文档地址：https://opendocs.alipay.com/open/04m8ci
func (a *Client) ZhimaMerchantZmgoTemplateQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaMerchantZmgoTemplateQueryRsp, err error) {
	err = bm.CheckEmptyError("template_no", "partner_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.merchant.zmgo.template.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaMerchantZmgoTemplateQueryRsp)
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

// zhima.credit.pe.zmgo.settle.apply(芝麻GO结算申请)
// 文档地址：https://opendocs.alipay.com/open/03usxk
func (a *Client) ZhimaCreditPeZmgoSettleApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoSettleApplyRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "partner_id", "out_request_no", "withhold_plan_no", "pay_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.settle.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoSettleApplyRsp)
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

// zhima.credit.pe.zmgo.bizopt.close(芝麻GO签约关单)
// 文档地址：https://opendocs.alipay.com/apis/01qii3
func (a *Client) ZhimaCreditPeZmgoBizoptClose(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoBizoptCloseRsp, err error) {
	err = bm.CheckEmptyError("alipay_user_id", "partner_id", "out_request_no", "template_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.bizopt.close"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoBizoptCloseRsp)
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

// zhima.credit.pe.zmgo.settle.refund(芝麻GO结算退款接口)
// 文档地址：https://opendocs.alipay.com/open/03ub1e
func (a *Client) ZhimaCreditPeZmgoSettleRefund(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoSettleRefundRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "partner_id", "alipay_user_id", "refund_amount", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.settle.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoSettleRefundRsp)
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

// zhima.credit.pe.zmgo.preorder.create(芝麻GO签约预创单)
// 文档地址：https://opendocs.alipay.com/open/03ujao
func (a *Client) ZhimaCreditPeZmgoPreorderCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoPreorderCreateRsp, err error) {
	err = bm.CheckEmptyError("partner_id", "template_id", "out_request_no", "biz_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.preorder.create"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoPreorderCreateRsp)
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

// zhima.credit.pe.zmgo.sign.apply(芝麻GO页面签约接口)
// 文档地址：https://opendocs.alipay.com/open/03u934
func (a *Client) ZhimaCreditPeZmgoSignApply(ctx context.Context, bm gopay.BodyMap) (orderStr string, err error) {
	err = bm.CheckEmptyError("partner_id", "template_id", "out_request_no")
	if err != nil {
		return gopay.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.sign.apply"); err != nil {
		return gopay.NULL, err
	}
	orderStr = string(bs)
	return orderStr, nil
}

// zhima.credit.pe.zmgo.agreement.unsign(芝麻GO协议解约)
// 文档地址：https://opendocs.alipay.com/open/03ub1g
func (a *Client) ZhimaCreditPeZmgoAgreementUnsign(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoAgreementUnsignRsp, err error) {
	err = bm.CheckEmptyError("partner_id", "agreement_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.agreement.unsign"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoAgreementUnsignRsp)
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

// zhima.credit.pe.zmgo.agreement.query(芝麻Go协议查询接口)
// 文档地址：https://opendocs.alipay.com/open/03utv3
func (a *Client) ZhimaCreditPeZmgoAgreementQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoAgreementQueryRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "alipay_user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.agreement.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoAgreementQueryRsp)
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

// zhima.credit.pe.zmgo.settle.unfreeze(芝麻Go解冻接口)
// 文档地址：https://opendocs.alipay.com/apis/01vx41
func (a *Client) ZhimaCreditPeZmgoSettleUnfreeze(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoSettleUnfreezeRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "out_request_no", "unfreeze_amount", "biz_time", "alipay_user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.settle.unfreeze"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoSettleUnfreezeRsp)
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

// zhima.credit.pe.zmgo.paysign.apply(芝麻GO支付下单链路签约申请)
// 文档地址：https://opendocs.alipay.com/apis/01xdtu
func (a *Client) ZhimaCreditPeZmgoPaysignApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoPaysignApplyRsp, err error) {
	err = bm.CheckEmptyError("alipay_user_id", "partner_id", "template_id", "merchant_app_id", "out_request_no", "biz_time", "timeout_express")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.paysign.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoPaysignApplyRsp)
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

// zhima.credit.pe.zmgo.paysign.confirm(芝麻GO支付下单链路签约确认)
// 文档地址：https://opendocs.alipay.com/apis/01xcif
func (a *Client) ZhimaCreditPeZmgoPaysignConfirm(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoPaysignConfirmRsp, err error) {
	err = bm.CheckEmptyError("alipay_user_id", "partner_id", "merchant_app_id", "zmgo_opt_no", "biz_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.pe.zmgo.paysign.confirm"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoPaysignConfirmRsp)
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

// zhima.credit.payafteruse.creditagreement.sign(信用服务开通/授权)
// 文档地址：https://opendocs.alipay.com/open/03uloz
func (a *Client) ZhimaCreditPayAfterUseAgreementSign(ctx context.Context, bm gopay.BodyMap) (orderStr string, err error) {
	err = bm.CheckEmptyError("out_agreement_no", "zm_service_id")
	if err != nil {
		return gopay.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.payafteruse.creditagreement.sign"); err != nil {
		return gopay.NULL, err
	}
	orderStr = string(bs)
	return orderStr, nil
}

// zhima.credit.payafteruse.creditagreement.query(查询服务开通/授权信息)
// 文档地址：https://opendocs.alipay.com/open/03ulp0
func (a *Client) ZhimaCreditPayAfterUseAgreementQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPayAfterUseAgreementQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.payafteruse.creditagreement.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPayAfterUseAgreementQueryRsp)
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

// zhima.credit.payafteruse.creditbizorder.order(芝麻信用服务下单（免用户确认场景）)
// 文档地址：https://opendocs.alipay.com/open/03ulpo
func (a *Client) ZhimaCreditPayAfterUseCreditBizOrder(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPayAfterUseCreditBizOrderRsp, err error) {
	err = bm.CheckEmptyError("out_order_no", "credit_agreement_id", "subject")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.payafteruse.creditbizorder.order"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPayAfterUseCreditBizOrderRsp)
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

// zhima.credit.payafteruse.creditbizorder.query(信用服务订单查询)
// 文档地址：https://opendocs.alipay.com/open/03vtet
func (a *Client) ZhimaCreditPayAfterUseCreditBizOrderQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPayAfterUseCreditBizOrderQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.payafteruse.creditbizorder.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPayAfterUseCreditBizOrderQueryRsp)
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

// zhima.credit.payafteruse.creditbizorder.finish(结束信用服务订单)
// 文档地址：https://opendocs.alipay.com/open/03vteu
func (a *Client) ZhimaCreditPayAfterUseCreditBizOrderFinish(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCreditPayAfterUseCreditBizOrderFinishRsp, err error) {
	err = bm.CheckEmptyError("out_request_no", "credit_biz_order_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.credit.payafteruse.creditbizorder.finish"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPayAfterUseCreditBizOrderFinishRsp)
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

// zhima.customer.jobworth.adapter.query(职得工作证信息匹配度查询)
// 文档地址：https://opendocs.alipay.com/apis/022mvz
func (a *Client) ZhimaCustomerJobworthAdapterQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCustomerJobworthAdapterQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.customer.jobworth.adapter.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCustomerJobworthAdapterQueryRsp)
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

// zhima.customer.jobworth.scene.use(职得工作证外部渠道应用数据回流)
// 文档地址：https://opendocs.alipay.com/apis/022waz
func (a *Client) ZhimaCustomerJobworthSceneUse(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCustomerJobworthSceneUseRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.customer.jobworth.scene.use"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCustomerJobworthSceneUseRsp)
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

// zhima.customer.jobworth.authentication.query(职得身份认证查询接口)
// 文档地址：https://opendocs.alipay.com/open/351177b5_zhima.customer.jobworth.authentication.query
func (a *Client) ZhimaCustomerJobworthAuthQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCustomerJobworthAuthQueryRsp, err error) {
	err = bm.CheckEmptyError("conn_key", "once_token", "query_type", "identity_type", "service_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.customer.jobworth.authentication.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCustomerJobworthAuthQueryRsp)
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

// zhima.customer.jobworth.authentication.preconsult(身份验真预咨询服务)
// 文档地址：https://opendocs.alipay.com/open/16f72e43_zhima.customer.jobworth.authentication.preconsult
func (a *Client) ZhimaCustomerJobworthAuthPreConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZhimaCustomerJobworthAuthPreConsultRsp, err error) {
	err = bm.CheckEmptyError("out_agreement_no", "once_token", "identity_type", "query_type", "zm_service_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "zhima.customer.jobworth.authentication.preconsult"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCustomerJobworthAuthPreConsultRsp)
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
