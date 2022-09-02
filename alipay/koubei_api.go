package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// koubei.trade.order.aggregate.consult(聚合支付订单咨询服务)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.order.aggregate.consult
func (a *Client) KoubeiTradeOrderAggregateConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeOrderAggregateConsultRsp, err error) {
	err = bm.CheckEmptyError("shop_id", "total_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.order.aggregate.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeOrderAggregateConsultRsp)
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

// koubei.trade.order.precreate(口碑订单预下单)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.order.precreate
func (a *Client) KoubeiTradeOrderPrecreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeOrderPrecreateRsp, err error) {
	err = bm.CheckEmptyError("request_id", "biz_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.order.precreate"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeOrderPrecreateRsp)
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

// koubei.trade.itemorder.buy(口碑商品交易购买接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.itemorder.buy
func (a *Client) KoubeiTradeItemorderBuy(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeItemorderBuyRsp, err error) {
	err = bm.CheckEmptyError("out_order_no", "subject", "biz_product", "biz_scene", "shop_id", "buyer_id", "total_amount", "item_order_details")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.itemorder.buy"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeItemorderBuyRsp)
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

// koubei.trade.order.consult(口碑订单预咨询)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.order.consult
func (a *Client) KoubeiTradeOrderConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeOrderConsultRsp, err error) {
	err = bm.CheckEmptyError("request_id", "user_id", "total_amount", "shop_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.order.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeOrderConsultRsp)
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

// koubei.trade.itemorder.refund(口碑商品交易退货接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.itemorder.refund
func (a *Client) KoubeiTradeItemorderRefund(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeItemorderRefundRsp, err error) {
	err = bm.CheckEmptyError("order_no", "out_request_no", "refund_infos")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.itemorder.refund"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeItemorderRefundRsp)
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

// koubei.trade.itemorder.query(口碑商品交易查询接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.itemorder.query
func (a *Client) KoubeiTradeItemorderQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeItemorderQueryRsp, err error) {
	err = bm.CheckEmptyError("order_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.itemorder.query"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeItemorderQueryRsp)
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

// koubei.trade.ticket.ticketcode.send(码商发码成功回调接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.ticket.ticketcode.send
func (a *Client) KoubeiTradeTicketTicketcodeSend(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeTicketTicketcodeSendRsp, err error) {
	err = bm.CheckEmptyError("request_id", "isv_ma_list", "send_order_no", "send_token", "order_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.ticket.ticketcode.send"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeTicketTicketcodeSendRsp)
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

// koubei.trade.ticket.ticketcode.delay(口碑凭证延期接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.ticket.ticketcode.delay
func (a *Client) KoubeiTradeTicketTicketcodeDelay(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeTicketTicketcodeDelayRsp, err error) {
	err = bm.CheckEmptyError("request_id", "end_date", "ticket_code", "code_type", "order_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.ticket.ticketcode.delay"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeTicketTicketcodeDelayRsp)
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

// koubei.trade.ticket.ticketcode.query(口碑凭证码查询)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.ticket.ticketcode.query
func (a *Client) KoubeiTradeTicketTicketcodeQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeTicketTicketcodeQueryRsp, err error) {
	err = bm.CheckEmptyError("ticket_code", "shop_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.ticket.ticketcode.query"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeTicketTicketcodeQueryRsp)
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

// koubei.trade.ticket.ticketcode.cancel(口碑凭证码撤销核销)
// 文档地址：https://opendocs.alipay.com/apis/api_1/koubei.trade.ticket.ticketcode.cancel
func (a *Client) KoubeiTradeTicketTicketcodeCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *KoubeiTradeTicketTicketcodeCancelRsp, err error) {
	err = bm.CheckEmptyError("request_id", "request_biz_no", "ticket_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "koubei.trade.ticket.ticketcode.cancel"); err != nil {
		return nil, err
	}
	aliRsp = new(KoubeiTradeTicketTicketcodeCancelRsp)
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
