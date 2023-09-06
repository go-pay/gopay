package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建商家券
// Code = 0 is success
func (c *ClientV3) V3BusiFavorBatchCreate(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorCreateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorBatchCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorBatchCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorCreateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchCreate)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询商家券详情
// Code = 0 is success
func (c *ClientV3) V3BusiFavorBatchDetail(ctx context.Context, stockId string) (wxRsp *BusiFavorBatchDetailRsp, err error) {
	uri := fmt.Sprintf(v3BusiFavorBatchDetail, stockId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorBatchDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorBatchDetail)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 核销用户券
// Code = 0 is success
func (c *ClientV3) V3BusiFavorUse(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorUseRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorUse, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorUse, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorUseRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorUse)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 根据过滤条件查询用户券
// Code = 0 is success
func (c *ClientV3) V3BusiFavorUserCoupons(ctx context.Context, openid string, bm gopay.BodyMap) (wxRsp *BusiFavorUserCouponsRsp, err error) {
	uri := fmt.Sprintf(v3BusiFavorUserCoupons, openid) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorUserCouponsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorUserCoupons)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询用户单张券详情
// Code = 0 is success
func (c *ClientV3) V3BusiFavorUserCouponDetail(ctx context.Context, openid, couponCode, appid string) (wxRsp *BusiFavorUserCouponDetailRsp, err error) {
	uri := fmt.Sprintf(v3BusiFavorUserCouponDetail, openid, couponCode, appid)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorUserCouponDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiUserCoupon)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 上传预存code
// Code = 0 is success
func (c *ClientV3) V3BusiFavorCodeUpload(ctx context.Context, stockId string, bm gopay.BodyMap) (wxRsp *BusiFavorCodeUploadRsp, err error) {
	url := fmt.Sprintf(v3BusiFavorCodeUpload, stockId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorCodeUploadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorCodeUpload)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 设置商家券事件通知地址
// Code = 0 is success
func (c *ClientV3) V3BusiFavorCallbackUrlSet(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorCallbackUrlSetRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorCallbackUrlSet, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorCallbackUrlSet, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorCallbackUrlSetRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorCallbackUrlSet)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询商家券事件通知地址
// Code = 0 is success
func (c *ClientV3) V3BusiFavorCallbackUrl(ctx context.Context, mchid string) (wxRsp *BusiFavorCallbackUrlRsp, err error) {
	uri := v3BusiFavorCallbackUrl + "?mchid=" + mchid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorCallbackUrlRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorCallbackUrl)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 关联订单信息
// Code = 0 is success
func (c *ClientV3) V3BusiFavorAssociate(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorAssociateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorAssociate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorAssociate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorAssociateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorAssociate)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 取消关联订单信息
// Code = 0 is success
func (c *ClientV3) V3BusiFavorDisassociate(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorDisassociateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorDisassociate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorDisassociate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorDisassociateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorDisassociate)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 修改批次预算
// Code = 0 is success
func (c *ClientV3) V3BusiFavorBatchUpdate(ctx context.Context, stockId string, bm gopay.BodyMap) (wxRsp *BusiFavorBatchUpdateRsp, err error) {
	url := fmt.Sprintf(v3BusiFavorBatchUpdate, stockId)
	authorization, err := c.authorization(MethodPATCH, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPatch(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorBatchUpdateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorBatchUpdate)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 修改商家券基本信息
// Code = 0 is success
func (c *ClientV3) V3BusiFavorInfoUpdate(ctx context.Context, stockId string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3BusiFavorInfoUpdate, stockId)
	authorization, err := c.authorization(MethodPATCH, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPatch(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 发放消费卡
// Code = 0 is success
func (c *ClientV3) V3BusiFavorSend(ctx context.Context, cardId string, bm gopay.BodyMap) (wxRsp *BusiFavorSendRsp, err error) {
	url := fmt.Sprintf(v3BusiFavorSend, cardId)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorSendRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorSend)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 申请退券
// Code = 0 is success
func (c *ClientV3) V3BusiFavorReturn(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorReturnRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorReturn, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorReturn, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorReturnRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorReturn)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 使券失效
// Code = 0 is success
func (c *ClientV3) V3BusiFavorDeactivate(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorDeactivateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorDeactivate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorDeactivate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorDeactivateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorDeactivate)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 营销补差付款
// Code = 0 is success
func (c *ClientV3) V3BusiFavorSubsidyPay(ctx context.Context, bm gopay.BodyMap) (wxRsp *BusiFavorSubsidyPayRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusiFavorSubsidyPay, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusiFavorSubsidyPay, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorSubsidyPayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorSubsidyPay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询营销补差付款单详情
// Code = 0 is success
func (c *ClientV3) V3BusiFavorSubsidyPayDetail(ctx context.Context, subsidyReceiptId string) (wxRsp *BusiFavorSubsidyPayDetailRsp, err error) {
	url := fmt.Sprintf(v3BusiFavorSubsidyPayDetail, subsidyReceiptId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BusiFavorSubsidyPayDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusiFavorSubsidyPay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
