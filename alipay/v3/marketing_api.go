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
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3MarketingActivityDeliveryCreate, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3MarketingActivityDeliveryCreate, authorization, aat)
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
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityDeliveryQuery, deliveryId)
	authorization, err := a.authorization(MethodPost, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, url, authorization, aat)
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
	aat := bm.GetString(HeaderAppAuthToken)
	url := fmt.Sprintf(v3MarketingActivityDeliveryStop, deliveryId)
	authorization, err := a.authorization(MethodPatch, url, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPatch(ctx, bm, url, authorization, aat)
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
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	// 临时存放 body file
	tempFile := make(gopay.BodyMap)
	signMap := make(gopay.BodyMap)
	// 遍历 map，把除了 file文件 字段之外的参数重新 set 到 bm 的 data 字段里签名用，然后移除自身
	bm.SetBodyMap("data", func(b gopay.BodyMap) {
		bm.Range(func(k string, v any) bool {
			// 取出 file 类型文件，签名时需要移除文件字段
			if file, ok := v.(*gopay.File); ok {
				// 保存到临时存放的 map 中
				tempFile.SetFormFile(k, file)
				// 原map删除此文件
				bm.Remove(k)
				return true
			}
			// 非 file 类型的参数 set 到签名用的 map 中
			signMap.Set(k, v)
			// 非 file 类型的参数 set 到 data 字段中，然后从原map中删除
			b.Set(k, v)
			bm.Remove(k)
			return true
		})
	})

	authorization, err := a.authorization(MethodPost, v3MarketingMaterialImageUpload, bm, aat)
	if err != nil {
		return nil, err
	}
	// 重新把file设置到原map中
	tempFile.Range(func(k string, v any) bool {
		bm.SetFormFile(k, v.(*gopay.File))
		return true
	})

	res, bs, err := a.doProdPostFile(ctx, bm, v3MarketingMaterialImageUpload, authorization, aat)
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
