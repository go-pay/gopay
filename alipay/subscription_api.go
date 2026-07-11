package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.trade.product.create(商品创建接口)
// 文档地址：https://opendocs.alipay.com/pre-open/0a257866_alipay.trade.product.create
func (a *Client) TradeProductCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeProductCreateResponse, err error) {
	err = bm.CheckEmptyError("name", "description")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.product.create"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeProductCreateResponse)
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

// alipay.trade.product.modify(商品修改接口)
// 文档地址：https://opendocs.alipay.com/pre-open/ed1d3c17_alipay.trade.product.modify
func (a *Client) TradeProductModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeProductModifyResponse, err error) {
	err = bm.CheckEmptyError("product_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.product.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeProductModifyResponse)
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

// alipay.trade.product.query(商品查询接口)
// 文档地址：https://opendocs.alipay.com/pre-open/81b9f0f0_alipay.trade.product.query
func (a *Client) TradeProductQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeProductQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.product.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeProductQueryResponse)
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

// alipay.trade.price.create(价格创建接口)
// 文档地址：https://opendocs.alipay.com/pre-open/df4e04f6_alipay.trade.price.create
func (a *Client) TradePriceCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradePriceCreateResponse, err error) {
	err = bm.CheckEmptyError("unit_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.price.create"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePriceCreateResponse)
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

// alipay.trade.price.query(价格查询接口)
// 文档地址：https://opendocs.alipay.com/pre-open/071f0840_alipay.trade.price.query
func (a *Client) TradePriceQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradePriceQueryResponse, err error) {
	err = bm.CheckEmptyError("price_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.price.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradePriceQueryResponse)
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

// alipay.trade.customer.create(客户创建接口)
// 文档地址：https://opendocs.alipay.com/pre-open/acce49e2_alipay.trade.customer.create
func (a *Client) TradeCustomerCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCustomerCreateResponse, err error) {
	err = bm.CheckEmptyError("name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.customer.create"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeCustomerCreateResponse)
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

// alipay.trade.subscription.create(订阅创建接口)
// 文档地址：https://opendocs.alipay.com/pre-open/201c1381_alipay.trade.subscription.create
func (a *Client) TradeSubscriptionCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeSubscriptionCreateResponse, err error) {
	err = bm.CheckEmptyError("items", "customer_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.subscription.create"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeSubscriptionCreateResponse)
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

// alipay.trade.subscription.modify(订阅修改接口)
// 文档地址：https://opendocs.alipay.com/pre-open/3b44eb16_alipay.trade.subscription.modify
func (a *Client) TradeSubscriptionModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeSubscriptionModifyResponse, err error) {
	err = bm.CheckEmptyError("subscription_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.subscription.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeSubscriptionModifyResponse)
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

// alipay.trade.subscription.query(订阅查询接口)
// 文档地址：https://opendocs.alipay.com/pre-open/9acd5c9e_alipay.trade.subscription.query
func (a *Client) TradeSubscriptionQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeSubscriptionQueryResponse, err error) {
	err = bm.CheckEmptyError("customer_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.subscription.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeSubscriptionQueryResponse)
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
