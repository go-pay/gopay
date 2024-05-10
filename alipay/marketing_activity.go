package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.marketing.activity.batchquery(条件查询活动列表)
// 文档地址：https://opendocs.alipay.com/open/04fgw9
func (a *Client) MarketingActivityBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("page_num", "page_size", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityBatchQueryRsp)
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

// alipay.marketing.activity.consult(活动领取咨询接口)
// 文档地址：https://opendocs.alipay.com/open/04fgwa
func (a *Client) MarketingActivityConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityConsultRsp, err error) {
	err = bm.CheckEmptyError("consult_activity_info_list", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityConsultRsp)
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

// alipay.marketing.activity.query(查询活动详情)
// 文档地址：https://opendocs.alipay.com/open/04fgwb
func (a *Client) MarketingActivityQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityQueryRsp)
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

// alipay.marketing.activity.merchant.batchquery(查询活动可用商户)
// 文档地址：https://opendocs.alipay.com/open/04fgwc
func (a *Client) MarketingActivityQueryMerchantBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryMerchantBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "page_num", "page_size")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.merchant.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityQueryMerchantBatchQueryRsp)
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

// alipay.marketing.activity.app.batchquery(查询活动可用小程序)
// 文档地址：https://opendocs.alipay.com/open/04fgwd
func (a *Client) MarketingActivityQueryAppBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryAppBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "page_num", "page_size", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.app.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityQueryAppBatchQueryRsp)
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

// alipay.marketing.activity.shop.batchquery(查询活动可用门店)
// 文档地址：https://opendocs.alipay.com/open/04fgwe
func (a *Client) MarketingActivityQueryShopBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryShopBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "page_num", "page_size")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.shop.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityQueryShopBatchQueryRsp)
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

// alipay.marketing.activity.goods.batchquery(查询活动适用商品)
// 文档地址：https://opendocs.alipay.com/open/04fgwf
func (a *Client) MarketingActivityQueryGoodsBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryGoodsBatchQueryRsp, err error) {
	err = bm.CheckEmptyError("activity_id", "page_num", "page_size")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.goods.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityQueryGoodsBatchQueryRsp)
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

// alipay.marketing.activity.user.batchqueryvoucher(条件查询用户券)
// 文档地址：https://opendocs.alipay.com/open/04fgwg
func (a *Client) MarketingActivityQueryUserBatchQueryVoucher(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryUserBatchQueryVoucherRsp, err error) {
	err = bm.CheckEmptyError("merchant_access_mode", "page_num", "page_size")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.user.batchqueryvoucher"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityQueryUserBatchQueryVoucherRsp)
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

// alipay.marketing.activity.user.queryvoucher(查询用户券详情)
// 文档地址：https://opendocs.alipay.com/open/04fgwh
func (a *Client) MarketingActivityQueryUserQueryVoucher(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityQueryUserQueryVoucherRsp, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.user.queryvoucher"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityQueryUserQueryVoucherRsp)
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
