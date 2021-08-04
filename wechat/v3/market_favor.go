package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建代金券批次API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_1.shtml
func (c *ClientV3) V3FavorBatchCreate(bm gopay.BodyMap) (wxRsp *FavorBatchCreateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3FavorBatchCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3FavorBatchCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchCreateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchCreate)
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

// 发放代金券批次API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_2.shtml
func (c *ClientV3) V3FavorBatchGrant(openid string, bm gopay.BodyMap) (wxRsp *FavorBatchGrantRsp, err error) {
	url := fmt.Sprintf(v3FavorBatchGrant, openid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchGrantRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchGrant)
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

// 激活代金券批次API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_3.shtml
func (c *ClientV3) V3FavorBatchStart(stockId, stockCreatorMchid string) (wxRsp *FavorBatchStartRsp, err error) {
	url := fmt.Sprintf(v3FavorBatchStart, stockId)
	bm := make(gopay.BodyMap)
	bm.Set("stock_creator_mchid", stockCreatorMchid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchStartRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchStart)
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

// 条件查询批次列表API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_4.shtml
func (c *ClientV3) V3FavorBatchList(bm gopay.BodyMap) (wxRsp *FavorBatchListRsp, err error) {
	uri := v3FavorBatchList + "?" + bm.EncodeGetParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatchList)
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

// 查询批次详情API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_5.shtml
func (c *ClientV3) V3FavorBatchDetail(stockId, stockCreatorMchid string) (wxRsp *FavorBatchDetailRsp, err error) {
	uri := fmt.Sprintf(v3FavorBatchDetail, stockId) + "?stock_creator_mchid=" + stockCreatorMchid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorBatchDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorBatch)
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

// 查询代金券详情API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_6.shtml
func (c *ClientV3) V3FavorDetailQuery(appid, couponId, openid string) (wxRsp *FavorDetailRsp, err error) {
	uri := fmt.Sprintf(v3FavorDetail, openid, couponId) + "?appid=" + appid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &FavorDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(FavorDetail)
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
