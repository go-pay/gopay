package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建现金活动 alipay.marketing.campaign.cash.create
// StatusCode = 200 is success
func (a *ClientV3) MarketingCampaignCashCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashCreateRsp, err error) {
	err = bm.CheckEmptyError("coupon_name", "prize_type", "total_money", "total_num", "prize_msg", "start_time", "end_time")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MarketingCampaignCashCreate, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingCampaignCashCreate, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingCampaignCashCreateRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 触发现金红包 alipay.marketing.campaign.cash.trigger
// StatusCode = 200 is success
func (a *ClientV3) MarketingCampaignCashTrigger(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashTriggerRsp, err error) {
	err = bm.CheckEmptyError("crowd_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MarketingCampaignCashTrigger, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingCampaignCashTrigger, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingCampaignCashTriggerRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 更改现金活动状态 alipay.marketing.campaign.cash.status.modify
// StatusCode = 200 is success
func (a *ClientV3) MarketingCampaignCashStatusModify(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashStatusModifyRsp, err error) {
	err = bm.CheckEmptyError("crowd_no", "camp_status")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MarketingCampaignCashStatusModify, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingCampaignCashStatusModify, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingCampaignCashStatusModifyRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 现金活动列表查询 alipay.marketing.campaign.cash.list.query
// StatusCode = 200 is success
func (a *ClientV3) MarketingCampaignCashListQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashListQueryRsp, err error) {
	err = bm.CheckEmptyError("page_size", "page_index")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3MarketingCampaignCashListQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingCampaignCashListQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 现金活动详情查询 alipay.marketing.campaign.cash.detail.query
// StatusCode = 200 is success
func (a *ClientV3) MarketingCampaignCashDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashDetailQueryRsp, err error) {
	err = bm.CheckEmptyError("crowd_no")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3MarketingCampaignCashDetailQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingCampaignCashDetailQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}
