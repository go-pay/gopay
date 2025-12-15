package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 商品图片上传接口 alipay.merchant.item.file.upload
// StatusCode = 200 is success
func (a *ClientV3) MerchantItemFileUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *MerchantItemFileUploadRsp, err error) {
	err = bm.CheckEmptyError("scene", "file_content")
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

	authorization, err := a.authorization(MethodPost, v3MerchantItemFileUpload, signMap, aat)
	if err != nil {
		return nil, err
	}
	// 重新把file设置到原map中
	tempFile.Range(func(k string, v any) bool {
		bm.SetFormFile(k, v.(*gopay.File))
		return true
	})

	res, bs, err := a.doProdPostFile(ctx, bm, v3MerchantItemFileUpload, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &MerchantItemFileUploadRsp{StatusCode: res.StatusCode}
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
