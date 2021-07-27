package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// koubei.trade.order.aggregate.consult(聚合支付订单咨询服务)
//	文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.order.aggregate.consult
func (a *Client) KoubeiTradeOrderAggregateConsult(bm gopay.BodyMap) (aliRsp *KoubeiTradeOrderAggregateConsultRsp, err error) {
	err = bm.CheckEmptyError("shop_id", "total_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "koubei.trade.order.aggregate.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeOrderAggregateConsultRsp)
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

// koubei.trade.order.precreate(口碑订单预下单)
//	文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.order.precreate
func (a *Client) KoubeiTradeOrderPrecreate(bm gopay.BodyMap) (aliRsp *KoubeiTradeOrderPrecreateRsp, err error) {
	err = bm.CheckEmptyError("request_id", "biz_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "koubei.trade.order.precreate"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeOrderPrecreateRsp)
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

// koubei.trade.itemorder.buy(口碑商品交易购买接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.itemorder.buy
func (a *Client) KoubeiTradeItemorderBuy(bm gopay.BodyMap) (aliRsp *KoubeiTradeItemorderBuyRsp, err error) {
	err = bm.CheckEmptyError("out_order_no", "subject", "biz_product", "biz_scene", "shop_id", "buyer_id", "total_amount", "item_order_details")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "koubei.trade.itemorder.buy"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeItemorderBuyRsp)
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
