package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// alipay.open.app.qrcode.create(小程序生成推广二维码接口)
// 文档地址：https://opendocs.alipay.com/apis/009zva
func (a *Client) OpenAppQrcodeCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenAppQrcodeCreateRsp, err error) {
	err = bm.CheckEmptyError("url_param", "query_param", "describe")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.open.app.qrcode.create"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppQrcodeCreateRsp)
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

// alipay.marketing.campaign.cash.create(创建现金活动接口)
// 文档地址：https://opendocs.alipay.com/open/029yy9
func (a *Client) MarketingCampaignCashCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashCreateRsp, err error) {
	err = bm.CheckEmptyError("coupon_name", "prize_type", "total_money", "total_num", "prize_msg", "start_time", "end_time", "merchant_link")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.campaign.cash.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCampaignCashCreateRsp)
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

// alipay.marketing.campaign.cash.trigger(触发现金红包活动)
// 文档地址：https://opendocs.alipay.com/open/029yya
func (a *Client) MarketingCampaignCashTrigger(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashTriggerRsp, err error) {
	err = bm.CheckEmptyError("user_id", "crowd_no", "login_id", "order_price", "out_biz_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.campaign.cash.trigger"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCampaignCashTriggerRsp)
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

// alipay.marketing.campaign.cash.status.modify(更改现金活动状态接口)
// 文档地址：https://opendocs.alipay.com/open/029yyb
func (a *Client) MarketingCampaignCashStatusModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashStatusModifyRsp, err error) {
	err = bm.CheckEmptyError("crowd_no", "camp_status")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.campaign.cash.status.modify"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCampaignCashStatusModifyRsp)
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

// alipay.marketing.campaign.cash.list.query(现金活动列表查询接口)
// 文档地址：https://opendocs.alipay.com/open/02a1f9
func (a *Client) MarketingCampaignCashListQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashListQueryRsp, err error) {
	err = bm.CheckEmptyError("camp_status", "page_size", "page_index")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.campaign.cash.list.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCampaignCashListQueryRsp)
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

// alipay.marketing.campaign.cash.detail.query(现金活动详情查询)
// 文档地址：https://opendocs.alipay.com/open/02a1fa
func (a *Client) MarketingCampaignCashDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashDetailQueryRsp, err error) {
	err = bm.CheckEmptyError("crowd_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.campaign.cash.detail.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCampaignCashDetailQueryRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	fmt.Println(err)
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
