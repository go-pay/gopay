package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.marketing.campaign.order.voucher.consult(订单优惠前置咨询)
// 文档地址：https://opendocs.alipay.com/open/04fgwi
func (a *Client) MarketingCampaignOrderVoucherConsult(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignOrderVoucherConsultRsp, err error) {
	err = bm.CheckEmptyError("scene_code", "order_amount")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.campaign.order.voucher.consult"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingCampaignOrderVoucherConsultRsp)
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

// alipay.marketing.activity.delivery.changed(推广计划状态变更消息)
// 文档地址：https://opendocs.alipay.com/open/85544608_alipay.marketing.activity.delivery.changed
func (a *Client) MarketingActivityDeliveryChanged(ctx context.Context, bm gopay.BodyMap) (success bool, err error) {
	err = bm.CheckEmptyError("event_time", "delivery_id", "delivery_status", "delivery_error_msg", "delivery_booth_code")
	if err != nil {
		return false, err
	}
	bs, err := a.doAliPay(ctx, bm, "alipay.marketing.activity.delivery.changed")
	if err != nil {
		return false, err
	}
	if strings.Contains(string(bs), "success") {
		return true, nil
	}
	return
}

// alipay.marketing.activity.delivery.stop(停止推广计划)
// 文档地址：https://opendocs.alipay.com/open/39c69f03_alipay.marketing.activity.delivery.stop
func (a *Client) MarketingActivityDeliveryStop(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityDeliveryStopRsp, err error) {
	err = bm.CheckEmptyError("delivery_id", "out_biz_no", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.delivery.stop"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityDeliveryStopRsp)
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

// alipay.marketing.activity.delivery.query(查询推广计划)
// 文档地址：https://opendocs.alipay.com/open/69c6d6c2_alipay.marketing.activity.delivery.query
func (a *Client) MarketingActivityDeliveryQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityDeliveryQueryRsp, err error) {
	err = bm.CheckEmptyError("delivery_id", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.delivery.query"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityDeliveryQueryRsp)
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

// alipay.marketing.activity.delivery.create(创建推广计划)
// 文档地址：https://opendocs.alipay.com/open/47498bf2_alipay.marketing.activity.delivery.create
func (a *Client) MarketingActivityDeliveryCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityDeliveryCreateRsp, err error) {
	err = bm.CheckEmptyError("delivery_booth_code", "out_biz_no", "delivery_base_info", "delivery_play_config", "delivery_target_rule", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.marketing.activity.delivery.create"); err != nil {
		return nil, err
	}
	aliRsp = new(MarketingActivityDeliveryCreateRsp)
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
