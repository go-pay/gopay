package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建电子发票卡券模板
// Code = 0 is success
func (c *ClientV3) V3InvoiceCardTemplateCreate(ctx context.Context, bm gopay.BodyMap) (*InvoiceCardTemplateCreateRsp, error) {
	if err := bm.CheckEmptyError("card_appid", "card_template_information"); err != nil {
		return nil, err
	}
	authorization, err := c.authorization(MethodPost, v3InvoiceCardTemplateCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3InvoiceCardTemplateCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceCardTemplateCreateRsp{Code: Success, SignInfo: si, Response: &InvoiceCardTemplateCreate{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 配置开发选项
// Code = 0 is success
func (c *ClientV3) V3InvoiceMerchantDevConfig(ctx context.Context, bm gopay.BodyMap) (*InvoiceMerchantDevConfigRsp, error) {
	authorization, err := c.authorization(MethodPATCH, v3InvoiceMerchantDevConfig, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPatch(ctx, bm, v3InvoiceMerchantDevConfig, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceMerchantDevConfigRsp{Code: Success, SignInfo: si, Response: &InvoiceMerchantDevConfig{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询商户配置的开发选项
// Code = 0 is success
func (c *ClientV3) V3InvoiceMerchantDevConfigQuery(ctx context.Context) (*InvoiceMerchantDevConfigQueryRsp, error) {
	authorization, err := c.authorization(MethodGet, v3InvoiceMerchantDevConfigQuery, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, v3InvoiceMerchantDevConfigQuery, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceMerchantDevConfigQueryRsp{Code: Success, SignInfo: si, Response: &InvoiceMerchantDevConfigQuery{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询电子发票
// Code = 0 is success
func (c *ClientV3) V3InvoiceQuery(ctx context.Context, fapiaoApplyId string, bm gopay.BodyMap) (*InvoiceQueryRsp, error) {
	uri := fmt.Sprintf(v3InvoiceQuery, fapiaoApplyId) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceQueryRsp{Code: Success, SignInfo: si, Response: &InvoiceQuery{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 获取抬头填写链接
// Code = 0 is success
func (c *ClientV3) V3InvoiceUserTitleUrl(ctx context.Context, bm gopay.BodyMap) (*InvoiceUserTitleUrlRsp, error) {
	if err := bm.CheckEmptyError("fapiao_apply_id", "appid", "openid", "total_amount", "source"); err != nil {
		return nil, err
	}
	uri := v3InvoiceUserTitleUrl + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceUserTitleUrlRsp{Code: Success, SignInfo: si, Response: &InvoiceUserTitleUrl{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 获取用户填写的抬头
// Code = 0 is success
func (c *ClientV3) V3InvoiceUserTitle(ctx context.Context, bm gopay.BodyMap) (*InvoiceUserTitleRsp, error) {
	if err := bm.CheckEmptyError("fapiao_apply_id", "scene"); err != nil {
		return nil, err
	}
	uri := v3InvoiceUserTitle + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceUserTitleRsp{Code: Success, SignInfo: si, Response: &InvoiceUserTitle{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 获取商户开票基础信息
// Code = 0 is success
func (c *ClientV3) V3InvoiceMerchantBaseInfo(ctx context.Context) (*InvoiceMerchantBaseInfoRsp, error) {
	authorization, err := c.authorization(MethodGet, v3InvoiceMerchantBaseInfo, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, v3InvoiceMerchantBaseInfo, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceMerchantBaseInfoRsp{Code: Success, SignInfo: si, Response: &InvoiceMerchantBaseInfo{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 获取商户可开具的商品和服务税收分类编码对照表
// Code = 0 is success
func (c *ClientV3) V3InvoiceMerchantTaxCodes(ctx context.Context, bm gopay.BodyMap) (*InvoiceMerchantTaxCodesRsp, error) {
	if err := bm.CheckEmptyError("offset", "limit"); err != nil {
		return nil, err
	}
	uri := v3InvoiceMerchantTaxCodes + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceMerchantTaxCodesRsp{Code: Success, SignInfo: si, Response: &InvoiceMerchantTaxCodes{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 开具电子发票
// Code = 0 is success
func (c *ClientV3) V3InvoiceCreate(ctx context.Context, bm gopay.BodyMap) (*EmptyRsp, error) {
	if err := bm.CheckEmptyError("scene", "fapiao_apply_id", "buyer_information", "fapiao_information"); err != nil {
		return nil, err
	}
	authorization, err := c.authorization(MethodPost, v3InvoiceCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3InvoiceCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusAccepted {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 冲红电子发票
// Code = 0 is success
func (c *ClientV3) V3InvoiceReverse(ctx context.Context, fapiaoApplyId string, bm gopay.BodyMap) (*EmptyRsp, error) {
	if err := bm.CheckEmptyError("reverse_reason"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(v3InvoiceReverse, fapiaoApplyId)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusAccepted {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 获取发票下载信息
// Code = 0 is success
func (c *ClientV3) V3InvoiceFileUrl(ctx context.Context, fapiaoApplyId string, bm gopay.BodyMap) (*InvoiceFileUrlRsp, error) {
	uri := fmt.Sprintf(v3InvoiceFileUrl, fapiaoApplyId) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &InvoiceFileUrlRsp{Code: Success, SignInfo: si, Response: &InvoiceFileUrl{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 上传电子发票文件
// 注意：非服务商时 subMchid 字段传空
// Code = 0 is success
func (c *ClientV3) V3InvoiceUploadFile(ctx context.Context, subMchid, fileType, digestAlogrithm, digest string, invoiceFile *gopay.File) (wxRsp *InvoiceUploadFileRsp, err error) {
	bmFile := make(gopay.BodyMap)
	bmFile.Set("file_type", fileType).
		Set("digest_alogrithm", digestAlogrithm).
		Set("digest", digest)

	if subMchid != "" {
		bmFile.Set("sub_mchid", subMchid)
	}
	authorization, err := c.authorization(MethodPost, v3InvoiceUploadFile, bmFile)
	if err != nil {
		return nil, err
	}
	bm := make(gopay.BodyMap)
	bm.Set("meta", bmFile).
		SetFormFile("file", invoiceFile)
	res, si, bs, err := c.doProdPostFile(ctx, bm, v3InvoiceUploadFile, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &InvoiceUploadFileRsp{Code: Success, SignInfo: si, Response: &InvoiceUploadFile{}}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 将电子发票插入微信用户卡包
// Code = 0 is success
func (c *ClientV3) V3InvoiceInsertCard(ctx context.Context, fapiaoApplyId string, bm gopay.BodyMap) (*EmptyRsp, error) {
	if err := bm.CheckEmptyError("scene", "buyer_information", "fapiao_card_information"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(v3InvoiceInsertCard, fapiaoApplyId)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusAccepted {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
