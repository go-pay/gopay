package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// Deprecated
// zhima.credit.score.get(查询芝麻用户的芝麻分)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.score.get
func (a *Client) ZhimaCreditScoreGet(bm gopay.BodyMap) (aliRsp *ZhimaCreditScoreGetResponse, err error) {
	if bm.GetString("product_code") == util.NULL {
		bm.Set("product_code", "w1010100100000000001")
	}
	err = bm.CheckEmptyError("transaction_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.score.get"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditScoreGetResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// zhima.credit.ep.scene.rating.initialize(芝麻企业信用信用评估初始化)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.rating.initialize
func (a *Client) ZhimaCreditEpSceneRatingInitialize(bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneRatingInitializeRsp, err error) {
	if bm.GetString("product_code") == util.NULL {
		bm.Set("product_code", "w1010100100000000001")
	}
	err = bm.CheckEmptyError("credit_category", "out_order_no", "user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.ep.scene.rating.initialize"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneRatingInitializeRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// zhima.credit.ep.scene.fulfillment.sync(信用服务履约同步)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.fulfillment.sync
func (a *Client) ZhimaCreditEpSceneFulfillmentSync(bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneFulfillmentSyncRsp, err error) {
	err = bm.CheckEmptyError("credit_order_no", "out_order_no", "biz_time", "biz_ext_param")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.ep.scene.fulfillment.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneFulfillmentSyncRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.credit.ep.scene.agreement.use(加入信用服务)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.agreement.use
func (a *Client) ZhimaCreditEpSceneAgreementUse(bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneAgreementUseRsp, err error) {
	err = bm.CheckEmptyError("rating_order_no", "out_order_no", "biz_time", "provision_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.ep.scene.agreement.use"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneAgreementUseRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.credit.ep.scene.agreement.cancel(取消信用服务)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.agreement.cancel
func (a *Client) ZhimaCreditEpSceneAgreementCancel(bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneAgreementCancelRsp, err error) {
	err = bm.CheckEmptyError("credit_order_no", "out_order_no", "biz_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.ep.scene.agreement.cancel"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneAgreementCancelRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.credit.ep.scene.fulfillmentlist.sync(信用服务履约同步(批量))
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.ep.scene.fulfillmentlist.sync
func (a *Client) ZhimaCreditEpSceneFulfillmentlistSync(bm gopay.BodyMap) (aliRsp *ZhimaCreditEpSceneFulfillmentlistSyncRsp, err error) {
	err = bm.CheckEmptyError("credit_order_no", "fulfillment_info_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.ep.scene.fulfillmentlist.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditEpSceneFulfillmentlistSyncRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.credit.pe.zmgo.cumulation.sync(芝麻go用户数据回传)
//	文档地址：https://opendocs.alipay.com/apis/api_8/zhima.credit.pe.zmgo.cumulation.sync
func (a *Client) ZhimaCreditPeZmgoCumulationSync(bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoCumulationSyncRsp, err error) {
	err = bm.CheckEmptyError("agreement_no", "user_id", "partner_id", "out_biz_no", "biz_time", "request_from", "biz_action", "cumulate_data_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.pe.zmgo.cumulation.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoCumulationSyncRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.merchant.zmgo.cumulate.sync(商家芝麻GO累计数据回传接口)
//	文档地址：https://opendocs.alipay.com/apis/01ol9h
func (a *Client) ZhimaMerchantZmgoCumulateSync(bm gopay.BodyMap) (aliRsp *ZhimaMerchantZmgoCumulateSyncRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "user_id", "provider_pid", "out_biz_no", "biz_time", "biz_action", "sub_biz_action", "data_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.merchant.zmgo.cumulate.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaMerchantZmgoCumulateSyncRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.merchant.zmgo.cumulate.query(商家芝麻GO累计数据查询接口)
//	文档地址：https://opendocs.alipay.com/apis/01ooeo
func (a *Client) ZhimaMerchantZmgoCumulateQuery(bm gopay.BodyMap) (aliRsp *ZhimaMerchantZmgoCumulateQueryRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "user_id", "provider_pid")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.merchant.zmgo.cumulate.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaMerchantZmgoCumulateQueryRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.credit.pe.zmgo.bizopt.close(芝麻GO签约关单)
//	文档地址：https://opendocs.alipay.com/apis/01qii3
func (a *Client) ZhimaCreditPeZmgoBizoptClose(bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoBizoptCloseRsp, err error) {
	err = bm.CheckEmptyError("alipay_user_id", "partner_id", "out_request_no", "template_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.pe.zmgo.bizopt.close"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoBizoptCloseRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.credit.pe.zmgo.settle.refund(芝麻GO结算退款接口)
//	文档地址：https://opendocs.alipay.com/apis/01rhsf
func (a *Client) ZhimaCreditPeZmgoSettleRefund(bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoSettleRefundRsp, err error) {
	err = bm.CheckEmptyError("agreement_id", "partner_id", "alipay_user_id", "refund_amount", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.pe.zmgo.settle.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoSettleRefundRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//  zhima.credit.pe.zmgo.preorder.create(芝麻GO签约预创单)
//	文档地址：https://opendocs.alipay.com/apis/01rhsk
func (a *Client) ZhimaCreditPeZmgoPreorderCreate(bm gopay.BodyMap) (aliRsp *ZhimaCreditPeZmgoPreorderCreateRsp, err error) {
	err = bm.CheckEmptyError("partner_id", "template_id", "out_request_no", "biz_time")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "zhima.credit.pe.zmgo.preorder.create"); err != nil {
		return nil, err
	}
	aliRsp = new(ZhimaCreditPeZmgoPreorderCreateRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
