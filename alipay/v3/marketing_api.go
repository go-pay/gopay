package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建推广计划 alipay.marketing.activity.delivery.create
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityDeliveryCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingActivityDeliveryCreateRsp, err error) {
	err = bm.CheckEmptyError("delivery_booth_code", "out_biz_no", "delivery_base_info", "delivery_play_config", "merchant_access_mode")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MarketingActivityDeliveryCreate, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingActivityDeliveryCreate, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityDeliveryCreateRsp{StatusCode: res.StatusCode}
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

// 查询推广计划 alipay.marketing.activity.delivery.query
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityDeliveryQuery(ctx context.Context, deliveryId string, bm gopay.BodyMap) (aliRsp *MarketingActivityDeliveryQueryRsp, err error) {
	err = bm.CheckEmptyError("merchant_access_mode")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(v3MarketingActivityDeliveryQuery, deliveryId)
	authorization, err := a.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityDeliveryQueryRsp{StatusCode: res.StatusCode}
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

// 停止推广计划 alipay.marketing.activity.delivery.stop
// StatusCode = 200 is success
func (a *ClientV3) MarketingActivityDeliveryStop(ctx context.Context, deliveryId string, bm gopay.BodyMap) (aliRsp *MarketingActivityDeliveryStopRsp, err error) {
	err = bm.CheckEmptyError("out_biz_no")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(v3MarketingActivityDeliveryStop, deliveryId)
	authorization, err := a.authorization(MethodPatch, url, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPatch(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingActivityDeliveryStopRsp{StatusCode: res.StatusCode}
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

// 营销图片资源上传接口 alipay.marketing.material.image.upload
// StatusCode = 200 is success
func (a *ClientV3) MarketingMaterialImageUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingMaterialImageUploadRsp, err error) {
	err = bm.CheckEmptyError("file_key", "file_content")
	if err != nil {
		return nil, err
	}
	upfile := bm.GetAny("file_content")
	// 签名时需要移除文件字段
	bm.Remove("file_content")
	// 遍历 map，把除了 file_content 字段之外的参数重新 set 到 bm 的 data 字段里，然后移除自身
	bm.SetBodyMap("data", func(b gopay.BodyMap) {
		bm.Range(func(k string, v any) bool {
			if k != "file_content" {
				b.Set(k, v)
				bm.Remove(k)
			}
			return true
		})
	})
	authorization, err := a.authorization(MethodPost, v3MarketingMaterialImageUpload, bm)
	if err != nil {
		return nil, err
	}
	bm.Set("file_content", upfile)
	// 至此，bodymap 内容 key 如下：file_content, data
	res, bs, err := a.doProdPostFile(ctx, bm, v3MarketingMaterialImageUpload, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &MarketingMaterialImageUploadRsp{StatusCode: res.StatusCode}
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

// MarketingCampaignCashCreate 创建现金活动
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

// MarketingCampaignCashTrigger 触发现金红包
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

// MarketingCampaignCashStatusModify 更改现金活动状态
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

// MarketingCampaignCashListQuery 现金活动列表查询
func (a *ClientV3) MarketingCampaignCashListQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashListQueryRsp, err error) {
	err = bm.CheckEmptyError("page_size", "page_index")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MarketingCampaignCashListQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingCampaignCashListQuery, authorization)
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

// MarketingCampaignCashDetailQuery 现金活动详情查询
func (a *ClientV3) MarketingCampaignCashDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *MarketingCampaignCashDetailQueryRsp, err error) {
	err = bm.CheckEmptyError("crowd_no")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3MarketingCampaignCashDetailQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingCampaignCashDetailQuery, authorization)
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
