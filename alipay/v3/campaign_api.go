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
