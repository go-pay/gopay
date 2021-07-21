package alipay

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// alipay.trade.pay(统一收单交易支付接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.pay
func (a *Client) TradePay(bm gopay.BodyMap) (aliRsp *TradePayResponse, err error) {
	err = bm.CheckEmptyError("out_trade_no", "subject")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.pay"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePayResponse)
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

// alipay.trade.precreate(统一收单线下交易预创建)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate
func (a *Client) TradePrecreate(bm gopay.BodyMap) (aliRsp *TradePrecreateResponse, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.precreate"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePrecreateResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	if aliRsp.NullResponse != nil {
		info := aliRsp.NullResponse
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.app.pay(app支付接口2.0)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
func (a *Client) TradeAppPay(bm gopay.BodyMap) (payParam string, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return util.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.app.pay"); err != nil {
		return util.NULL, err
	}
	payParam = string(bs)
	return payParam, nil
}

// alipay.trade.wap.pay(手机网站支付接口2.0)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay
func (a *Client) TradeWapPay(bm gopay.BodyMap) (payUrl string, err error) {
	bm.Set("product_code", "QUICK_WAP_WAY")
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return util.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.wap.pay"); err != nil {
		return util.NULL, err
	}
	payUrl = string(bs)
	return payUrl, nil
}

// alipay.trade.page.pay(统一收单下单并支付页面接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
func (a *Client) TradePagePay(bm gopay.BodyMap) (payUrl string, err error) {
	bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return util.NULL, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.page.pay"); err != nil {
		return util.NULL, err
	}
	payUrl = string(bs)
	return payUrl, nil
}

// alipay.trade.create(统一收单交易创建接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.create
func (a *Client) TradeCreate(bm gopay.BodyMap) (aliRsp *TradeCreateResponse, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.create"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCreateResponse)
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

// alipay.trade.query(统一收单线下交易查询)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.query
func (a *Client) TradeQuery(bm gopay.BodyMap) (aliRsp *TradeQueryResponse, err error) {
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("trade_no") == util.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeQueryResponse)
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

// alipay.trade.cancel(统一收单交易撤销接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.cancel
func (a *Client) TradeCancel(bm gopay.BodyMap) (aliRsp *TradeCancelResponse, err error) {
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("trade_no") == util.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.cancel"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCancelResponse)
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

// alipay.trade.close(统一收单交易关闭接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.close
func (a *Client) TradeClose(bm gopay.BodyMap) (aliRsp *TradeCloseResponse, err error) {
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("trade_no") == util.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.close"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCloseResponse)
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

// alipay.trade.refund(统一收单交易退款接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.refund
func (a *Client) TradeRefund(bm gopay.BodyMap) (aliRsp *TradeRefundResponse, err error) {
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("trade_no") == util.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("refund_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRefundResponse)
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

// alipay.trade.page.refund(统一收单退款页面接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.refund
func (a *Client) TradePageRefund(bm gopay.BodyMap) (aliRsp *TradePageRefundResponse, err error) {
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("trade_no") == util.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("out_request_no", "refund_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.page.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePageRefundResponse)
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

// alipay.trade.fastpay.refund.query(统一收单交易退款查询)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.fastpay.refund.query
func (a *Client) TradeFastPayRefundQuery(bm gopay.BodyMap) (aliRsp *TradeFastpayRefundQueryResponse, err error) {
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("trade_no") == util.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.fastpay.refund.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeFastpayRefundQueryResponse)
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

// alipay.trade.order.settle(统一收单交易结算接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.order.settle
func (a *Client) TradeOrderSettle(bm gopay.BodyMap) (aliRsp *TradeOrderSettleResponse, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "royalty_parameters")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.order.settle"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleResponse)
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

// alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.orderinfo.sync
func (a *Client) TradeOrderInfoSync(bm gopay.BodyMap) (aliRsp *TradeOrderInfoSyncRsp, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "biz_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.orderinfo.sync"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderInfoSyncRsp)
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

// alipay.trade.advance.consult(订单咨询服务)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.advance.consult
func (a *Client) TradeAdvanceConsult(bm gopay.BodyMap) (aliRsp *TradeAdvanceConsultRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.trade.advance.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeAdvanceConsultRsp)
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

// Deprecated
// koubei.trade.order.aggregate.consult(聚合支付订单咨询服务)
//	文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.order.aggregate.consult
func (a *Client) TradeOrderAggregateConsult(bm gopay.BodyMap) (aliRsp *TradeOrderAggregateConsultRsp, err error) {
	err = bm.CheckEmptyError("shop_id", "total_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "koubei.trade.order.aggregate.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderAggregateConsultRsp)
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

// alipay.pcredit.huabei.auth.settle.apply(花芝轻会员结算申请)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.pcredit.huabei.auth.settle.apply
func (a *Client) PcreditHuabeiAuthSettleApply(bm gopay.BodyMap) (aliRsp *PcreditHuabeiAuthSettleApplyRsp, err error) {
	err = bm.CheckEmptyError("agreement_no", "pay_amount", "out_request_no", "alipay_user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.pcredit.huabei.auth.settle.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(PcreditHuabeiAuthSettleApplyRsp)
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

// alipay.commerce.transport.nfccard.send(NFC用户卡信息同步)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.commerce.transport.nfccard.send
func (a *Client) CommerceTransportNfccardSend(bm gopay.BodyMap) (aliRsp *CommerceTransportNfccardSendRsp, err error) {
	err = bm.CheckEmptyError("issue_org_no", "card_no", "card_status")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.commerce.transport.nfccard.send"); err != nil {
		return nil, err
	}
	aliRsp = new(CommerceTransportNfccardSendRsp)
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

// alipay.data.dataservice.ad.data.query(广告投放数据查询)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.data.dataservice.ad.data.query
func (a *Client) DataDataserviceAdDataQuery(bm gopay.BodyMap) (aliRsp *DataDataserviceAdDataQueryRsp, err error) {
	err = bm.CheckEmptyError("query_type", "biz_token", "ad_level", "start_date", "end_date", "outer_id_list")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.data.dataservice.ad.data.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataDataserviceAdDataQueryRsp)
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

// alipay.commerce.air.callcenter.trade.apply(航司电话订票待申请接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.commerce.air.callcenter.trade.apply
func (a *Client) CommerceAirCallcenterTradeApply(bm gopay.BodyMap) (aliRsp *CommerceAirCallcenterTradeApplyRsp, err error) {
	err = bm.CheckEmptyError("scene_code", "op_code", "channel", "target_id", "target_id_type", "trade_apply_params")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.commerce.air.callcenter.trade.apply"); err != nil {
		return nil, err
	}
	aliRsp = new(CommerceAirCallcenterTradeApplyRsp)
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

// mybank.payment.trade.order.create(网商银行全渠道收单业务订单创建)
//	文档地址：https://opendocs.alipay.com/apis/api_1/mybank.payment.trade.order.create
func (a *Client) PaymentTradeOrderCreate(bm gopay.BodyMap) (aliRsp *PaymentTradeOrderCreateRsp, err error) {
	err = bm.CheckEmptyError("partner_id", "out_trade_no", "recon_related_no", "pd_code", "ev_code", "total_amount", "currency_code", "goods_info", "seller_id", "pay_type", "pay_date")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "mybank.payment.trade.order.create"); err != nil {
		return nil, err
	}
	aliRsp = new(PaymentTradeOrderCreateRsp)
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
