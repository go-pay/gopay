package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// ant.merchant.expand.shop.modify(修改蚂蚁店铺)
// 文档地址：https://opendocs.alipay.com/apis/014tmb
func (a *Client) AntMerchantShopModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopModifyRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopModifyRsp)
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

// ant.merchant.expand.shop.create(蚂蚁店铺创建)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.create
func (a *Client) AntMerchantShopCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopCreateRsp, err error) {
	err = bm.CheckEmptyError("business_address", "shop_category", "store_id", "shop_type", "ip_role_id", "shop_name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.create"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopCreateRsp)
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

// ant.merchant.expand.shop.consult(蚂蚁店铺创建咨询)
// 文档地址：https://opendocs.alipay.com/apis/014yig
func (a *Client) AntMerchantShopConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopConsultRsp, err error) {
	err = bm.CheckEmptyError("business_address", "shop_category", "store_id", "shop_type", "ip_role_id", "shop_name")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopConsultRsp)
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

// ant.merchant.expand.order.query(商户申请单查询)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.order.query
func (a *Client) AntMerchantOrderQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantOrderQueryRsp, err error) {
	err = bm.CheckEmptyError("order_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.order.query"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantOrderQueryRsp)
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

// ant.merchant.expand.shop.query(店铺查询接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.query
func (a *Client) AntMerchantShopQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopQueryRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.query"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopQueryRsp)
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

// ant.merchant.expand.shop.close(蚂蚁店铺关闭)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.close
func (a *Client) AntMerchantShopClose(ctx context.Context, bm gopay.BodyMap) (aliRsp *AntMerchantShopCloseRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "ant.merchant.expand.shop.close"); err != nil {
		return nil, err
	}
	aliRsp = new(AntMerchantShopCloseRsp)
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
