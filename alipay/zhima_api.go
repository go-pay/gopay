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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
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
		return nil, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
