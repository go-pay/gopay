package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 提交申请单API
//	注意：本接口会提交一些敏感信息，需调用 client.V3EncryptText() 进行加密
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_1.shtml
func (c *ClientV3) V3Apply4SubSubmit(ctx context.Context, bm gopay.BodyMap) (*Apply4SubSubmitRsp, error) {
	if err := bm.CheckEmptyError(
		"business_code", "contact_info", "subject_info",
		"business_info", "settlement_info", "bank_account_info"); err != nil {
		return nil, err
	}

	authorization, err := c.authorization(MethodPost, v3Apply4SubSubmit, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3Apply4SubSubmit, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &Apply4SubSubmitRsp{Code: Success, SignInfo: si}
	wxResp.Response = new(Apply4SubSubmit)
	if err = json.Unmarshal(bs, wxResp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}

	if res.StatusCode != http.StatusOK {
		wxResp.Code = res.StatusCode
		wxResp.Error = string(bs)
		return wxResp, nil
	}
	return wxResp, c.verifySyncSign(si)
}

// 通过业务申请编号查询申请状态API
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (c *ClientV3) V3Apply4SubQueryByBusinessCode(ctx context.Context, businessCode string) (*Apply4SubQueryRsp, error) {
	uri := fmt.Sprintf(v3Apply4SubQueryByBusinessCode, businessCode)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &Apply4SubQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Apply4SubQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 通过申请单号查询申请状态API
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (c *ClientV3) V3Apply4SubQueryByApplyId(ctx context.Context, applyId string) (*Apply4SubQueryRsp, error) {
	uri := fmt.Sprintf(v3Apply4SubQueryByApplyId, applyId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &Apply4SubQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Apply4SubQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 修改结算账号 API
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_3.shtml
// 	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_1_4.shtml
func (c *ClientV3) V3Apply4SubModifySettlement(ctx context.Context, bm gopay.BodyMap) (*EmptyRsp, error) {
	if err := bm.CheckEmptyError(
		"sub_mchid", "account_type", "account_bank", "bank_address_code", "account_number",
	); err != nil {
		return nil, err
	}

	authorization, err := c.authorization(MethodPost, v3Apply4SubModifySettlement, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3Apply4SubModifySettlement, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &EmptyRsp{Code: Success, SignInfo: si}

	if res.StatusCode != http.StatusNoContent {
		wxResp.Code = res.StatusCode
		wxResp.Error = string(bs)
		return wxResp, nil
	}
	return wxResp, c.verifySyncSign(si)
}

// 查询结算账户 API
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_4.shtml
// 	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_1_5.shtml
func (c *ClientV3) V3Apply4SubQuerySettlement(ctx context.Context, subMchId string) (*Apply4SubQuerySettlementRsp, error) {
	uri := fmt.Sprintf(v3Apply4SubQuerySettlement, subMchId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &Apply4SubQuerySettlementRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Apply4SubQuerySettlement)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 二级商户进件API
//	注意：本接口会提交一些敏感信息，需调用 client.V3EncryptText() 进行加密。部分图片参数，请先调用 client.V3MediaUploadImage() 上传，获取MediaId
//	Code = 0 is success
// 	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_1_1.shtml
func (c *ClientV3) V3EcommerceApply(ctx context.Context, bm gopay.BodyMap) (*EcommerceApplyRsp, error) {
	authorization, err := c.authorization(MethodPost, v3EcommerceApply, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3EcommerceApply, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &EcommerceApplyRsp{Code: Success, SignInfo: si}
	wxResp.Response = new(EcommerceApply)
	if err = json.Unmarshal(bs, wxResp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxResp.Code = res.StatusCode
		wxResp.Error = string(bs)
		return wxResp, nil
	}
	return wxResp, c.verifySyncSign(si)
}

// 查询申请状态API
//	注意：applyId 和 outRequestNo 二选一
//	Code = 0 is success
// 	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_1_2.shtml
func (c *ClientV3) V3EcommerceApplyStatus(ctx context.Context, applyId int64, outRequestNo string) (*EcommerceApplyStatusRsp, error) {
	if applyId == 0 && outRequestNo == gopay.NULL {
		return nil, fmt.Errorf("applyId[%d] and outRequestNo[%s] empty at the same time", applyId, outRequestNo)
	}
	var url string
	if applyId != 0 {
		url = fmt.Sprintf(v3EcommerceApplyQueryById, applyId)
	} else {
		url = fmt.Sprintf(v3EcommerceApplyQueryByNo, outRequestNo)
	}
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EcommerceApplyStatusRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceApplyStatus)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
