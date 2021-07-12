package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 特约商户进件
// 提交申请单API
// Code = 0 is success
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_1.shtml
// 本接口会提交一些敏感信息，需调用V3EncryptText进行加密
func (c *ClientV3) V3Applyment4subSubmit(bm gopay.BodyMap) (*Applyment4SubSubmitResp, error) {
	if err := bm.CheckEmptyError(
		"business_code", "contact_info", "subject_info",
		"business_info", "settlement_info", "bank_account_info"); err != nil {
		return nil, err
	}

	authorization, err := c.authorization(MethodPost, v3Applyment4SubSubmit, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3Applyment4SubSubmit, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &Applyment4SubSubmitResp{Code: Success, SignInfo: si}
	wxResp.Response = new(Applyment4SubSubmit)
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
// Code = 0 is success
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (c *ClientV3) V3Applyment4subQueryByBusinessCode(businessCode string) (*Applyment4SubQueryResp, error) {
	uri := fmt.Sprintf(v3Applyment4SubQueryByBusinessCode, businessCode)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &Applyment4SubQueryResp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Applyment4SubQuery)
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
// Code = 0 is success
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (c *ClientV3) V3Applyment4subQueryByApplymentId(applymentId uint64) (*Applyment4SubQueryResp, error) {
	uri := fmt.Sprintf(v3Applyment4SubQueryByApplymentId, applymentId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &Applyment4SubQueryResp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Applyment4SubQuery)
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
// sub_mchid长度最小8个字节
// Code = 0 is success
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_3.shtml
func (c *ClientV3) V3Applyment4subModifySettlement(bm gopay.BodyMap) (*Applyment4SubModifySettlementResp, error) {
	if err := bm.CheckEmptyError(
		"sub_mchid", "account_type", "account_bank", "bank_address_code", "account_number",
	); err != nil {
		return nil, err
	}

	authorization, err := c.authorization(MethodPost, v3Applyment4SubModifySettlement, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3Applyment4SubModifySettlement, authorization)
	if err != nil {
		return nil, err
	}
	wxResp := &Applyment4SubModifySettlementResp{Code: Success, SignInfo: si}

	if res.StatusCode != http.StatusNoContent {
		wxResp.Code = res.StatusCode
		wxResp.Error = string(bs)
		return wxResp, nil
	}
	return wxResp, c.verifySyncSign(si)
}

// 查询结算账户 API
// Code = 0 is success
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_4.shtml
func (c *ClientV3) V3Applyment4subQuerySettlement(subMchId string) (*Applyment4SubQuerySettlementResp, error) {
	uri := fmt.Sprintf(v3Applyment4SubQuerySettlement, subMchId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &Applyment4SubQuerySettlementResp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Applyment4SubQuerySettlement)
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
