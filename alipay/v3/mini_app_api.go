package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 小程序退回开发 alipay.open.mini.version.audited.cancel
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionAuditedCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionAuditedCancelRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionAuditedCancel, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionAuditedCancel, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionAuditedCancelRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序灰度上架 alipay.open.mini.version.gray.online
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionGrayOnline(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionGrayOnlineRsp, err error) {
	err = bm.CheckEmptyError("app_version", "gray_strategy")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionGrayOnline, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionGrayOnline, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionGrayOnlineRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序结束灰度 alipay.open.mini.version.gray.cancel
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionGrayCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionGrayCancelRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionGrayCancel, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionGrayCancel, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionGrayCancelRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序上架 alipay.open.mini.version.online
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionOnline(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionOnlineRsp, err error) {
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionOnline, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionOnline, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionOnlineRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序下架 alipay.open.mini.version.offline
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionOffline(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionOfflineRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionOffline, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionOffline, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionOfflineRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序回滚 alipay.open.mini.version.rollback
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionRollback(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionRollbackRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionRollback, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionRollback, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionRollbackRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序删除版本 alipay.open.mini.version.delete
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionDelete(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionDeleteRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionDelete, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionDelete, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionDeleteRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序提交审核 alipay.open.mini.version.audit.apply
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionAuditApply(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionAuditApplyRsp, err error) {
	err = bm.CheckEmptyError("app_version", "version_desc")
	if err != nil {
		return nil, err
	}
	if bm.GetString("service_email") == gopay.NULL && bm.GetString("service_phone") == gopay.NULL {
		return nil, errors.New("service_email and service_phone are not allowed to be null at the same time")
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

	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionAuditApply, signMap, aat)
	if err != nil {
		return nil, err
	}
	// 重新把file设置到原map中
	tempFile.Range(func(k string, v any) bool {
		bm.SetFormFile(k, v.(*gopay.File))
		return true
	})

	res, bs, err := a.doProdPostFile(ctx, bm, v3OpenMiniVersionAuditApply, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionAuditApplyRsp{StatusCode: res.StatusCode}
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

// 小程序基于模板上传版本 alipay.open.mini.version.upload
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionUploadRsp, err error) {
	err = bm.CheckEmptyError("template_id", "app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniVersionUpload, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniVersionUpload, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionUploadRsp{StatusCode: res.StatusCode}
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

// 查询使用模板的小程序列表 alipay.open.mini.template.usage.query
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniTemplateUsageQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniTemplateUsageQueryRsp, err error) {
	err = bm.CheckEmptyError("template_id")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3OpenMiniTemplateUsageQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniTemplateUsageQueryRsp{StatusCode: res.StatusCode}
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

// 小程序查询版本构建状态 alipay.open.mini.version.build.query
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionBuildQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionBuildQueryRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3OpenMiniVersionBuildQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionBuildQueryRsp{StatusCode: res.StatusCode}
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

// 小程序版本详情查询 alipay.open.mini.version.detail.query
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionDetailQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionDetailQueryRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3OpenMiniVersionDetailQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionDetailQueryRsp{StatusCode: res.StatusCode}
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

// 小程序版本列表查询 alipay.open.mini.version.list.query
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniVersionListQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniVersionListQueryRsp, err error) {
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3OpenMiniVersionListQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniVersionListQueryRsp{StatusCode: res.StatusCode}
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

// 小程序生成体验版 alipay.open.mini.experience.create
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniExperienceCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniExperienceCreateRsp, err error) {
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniExperienceCreate, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniExperienceCreate, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniExperienceCreateRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 小程序体验版状态查询接口 alipay.open.mini.experience.query
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniExperienceQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniExperienceQueryRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	bm.Remove(HeaderAppAuthToken)
	uri := v3OpenMiniExperienceQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniExperienceQueryRsp{StatusCode: res.StatusCode}
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

// 小程序取消体验版 alipay.open.mini.experience.cancel
// StatusCode = 200 is success
func (a *ClientV3) OpenMiniExperienceCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *OpenMiniExperienceCancelRsp, err error) {
	err = bm.CheckEmptyError("app_version")
	if err != nil {
		return nil, err
	}
	aat := bm.GetString(HeaderAppAuthToken)
	authorization, err := a.authorization(MethodPost, v3OpenMiniExperienceCancel, bm, aat)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3OpenMiniExperienceCancel, authorization, aat)
	if err != nil {
		return nil, err
	}
	aliRsp = &OpenMiniExperienceCancelRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}
