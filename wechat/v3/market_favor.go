package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// 创建代金券批次
// Code = 0 is success
func (c *ClientV3) V3FavorBatchCreate(ctx context.Context, bm gopay.BodyMap) (wxRsp *FavorBatchCreateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3FavorBatchCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3FavorBatchCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchCreateRsp{Code: Success, SignInfo: si}
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

// 发放代金券批次
// Code = 0 is success
func (c *ClientV3) V3FavorBatchGrant(ctx context.Context, openid string, bm gopay.BodyMap) (wxRsp *FavorBatchGrantRsp, err error) {
	url := fmt.Sprintf(v3FavorBatchGrant, openid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchGrantRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchGrant)
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

// 激活代金券批次
// Code = 0 is success
func (c *ClientV3) V3FavorBatchStart(ctx context.Context, stockId, stockCreatorMchid string) (wxRsp *FavorBatchStartRsp, err error) {
	url := fmt.Sprintf(v3FavorBatchStart, stockId)
	bm := make(gopay.BodyMap)
	bm.Set("stock_creator_mchid", stockCreatorMchid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchStartRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchStart)
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

// 条件查询批次列表
// Code = 0 is success
func (c *ClientV3) V3FavorBatchList(ctx context.Context, bm gopay.BodyMap) (wxRsp *FavorBatchListRsp, err error) {
	uri := v3FavorBatchList + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchList)
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

// 查询批次详情
// Code = 0 is success
func (c *ClientV3) V3FavorBatchDetail(ctx context.Context, stockId, stockCreatorMchid string) (wxRsp *FavorBatchDetailRsp, err error) {
	uri := fmt.Sprintf(v3FavorBatchDetail, stockId) + "?stock_creator_mchid=" + stockCreatorMchid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatch)
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

// 查询代金券详情
// Code = 0 is success
func (c *ClientV3) V3FavorDetail(ctx context.Context, appid, couponId, openid string) (wxRsp *FavorDetailRsp, err error) {
	uri := fmt.Sprintf(v3FavorDetail, openid, couponId) + "?appid=" + appid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorDetail)
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

// 查询代金券可用商户
// Code = 0 is success
func (c *ClientV3) V3FavorMerchant(ctx context.Context, stockId, stockCreatorMchid string, limit, offset int) (wxRsp *FavorMerchantRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := fmt.Sprintf(v3FavorMerchant, stockId) + "?stock_creator_mchid=" + stockCreatorMchid + "&limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorMerchantRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorMerchant)
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

// 查询代金券可用单品
// Code = 0 is success
func (c *ClientV3) V3FavorItems(ctx context.Context, stockId, stockCreatorMchid string, limit, offset int) (wxRsp *FavorItemsRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := fmt.Sprintf(v3FavorItems, stockId) + "?stock_creator_mchid=" + stockCreatorMchid + "&limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorItemsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorItems)
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

// 根据商户号查用户的券
// Code = 0 is success
func (c *ClientV3) V3FavorUserCoupons(ctx context.Context, openid string, bm gopay.BodyMap) (wxRsp *FavorUserCouponsRsp, err error) {
	uri := fmt.Sprintf(v3FavorUserCoupons, openid) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorUserCouponsRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorUserCoupons)
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

// 下载批次核销明细
// Code = 0 is success
func (c *ClientV3) V3FavorUseFlowDownload(ctx context.Context, stockId string) (wxRsp *FavorUseFlowDownloadRsp, err error) {
	url := fmt.Sprintf(v3FavorUseFlowDownload, stockId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorUseFlowDownloadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorFlowDownload)
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

// 下载批次退款明细
// Code = 0 is success
func (c *ClientV3) V3FavorRefundFlowDownload(ctx context.Context, stockId string) (wxRsp *FavorRefundFlowDownloadRsp, err error) {
	url := fmt.Sprintf(v3FavorRefundFlowDownload, stockId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorRefundFlowDownloadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorFlowDownload)
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

// 设置消息通知地址
// Code = 0 is success
func (c *ClientV3) V3FavorCallbackUrlSet(ctx context.Context, bm gopay.BodyMap) (wxRsp *FavorCallbackUrlSetRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3FavorCallbackUrlSet, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3FavorCallbackUrlSet, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorCallbackUrlSetRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorCallbackUrl)
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

// 暂停代金券批次
// Code = 0 is success
func (c *ClientV3) V3FavorBatchPause(ctx context.Context, stockId, stockCreatorMchid string) (wxRsp *FavorBatchPauseRsp, err error) {
	url := fmt.Sprintf(v3FavorBatchPause, stockId)
	bm := make(gopay.BodyMap)
	bm.Set("stock_creator_mchid", stockCreatorMchid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchPauseRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchPause)
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

// 重启代金券批次
// Code = 0 is success
func (c *ClientV3) V3FavorBatchRestart(ctx context.Context, stockId, stockCreatorMchid string) (wxRsp *FavorBatchRestartRsp, err error) {
	url := fmt.Sprintf(v3FavorBatchRestart, stockId)
	bm := make(gopay.BodyMap)
	bm.Set("stock_creator_mchid", stockCreatorMchid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchRestartRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchRestart)
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
