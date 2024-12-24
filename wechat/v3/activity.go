package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建全场满额送活动
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityCreate(ctx context.Context, bm gopay.BodyMap) (*PayGiftActivityCreateRsp, error) {
	if err := bm.CheckEmptyError("activity_base_info", "award_send_rule"); err != nil {
		return nil, err
	}
	authorization, err := c.authorization(MethodPost, v3PayGiftActivityCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3PayGiftActivityCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityCreateRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityCreate{}}
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

// 获取支付有礼活动列表
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityList(ctx context.Context, bm gopay.BodyMap) (*PayGiftActivityListRsp, error) {
	uri := v3PayGiftActivityList + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityListRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityList{}}
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

// 获取活动详情
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityDetail(ctx context.Context, activityId string) (*PayGiftActivityDetailRsp, error) {
	uri := fmt.Sprintf(v3PayGiftActivityDetail, activityId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityDetailRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityDetail{}}
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

// 获取活动指定商品列表
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityGoods(ctx context.Context, activityId string, bm gopay.BodyMap) (*PayGiftActivityGoodsRsp, error) {
	uri := fmt.Sprintf(v3PayGiftActivityGoods, activityId) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityGoodsRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityGoods{}}
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

// 终止活动
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityTerminate(ctx context.Context, activityId string) (*PayGiftActivityTerminateRsp, error) {
	uri := fmt.Sprintf(v3PayGiftActivityTerminate, activityId)
	authorization, err := c.authorization(MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, nil, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityTerminateRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityTerminate{}}
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

// 获取活动发券商户号
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityMerchant(ctx context.Context, activityId string, bm gopay.BodyMap) (*PayGiftActivityMerchantRsp, error) {
	uri := fmt.Sprintf(v3PayGiftActivityMerchant, activityId) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityMerchantRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityMerchant{}}
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

// 新增活动发券商户号
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityMerchantAdd(ctx context.Context, activityId string, bm gopay.BodyMap) (*PayGiftActivityMerchantAddRsp, error) {
	if err := bm.CheckEmptyError("merchant_id_list", "add_request_no"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(v3PayGiftActivityMerchantAdd, activityId)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityMerchantAddRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityMerchantAdd{}}
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

// 删除活动发券商户号
// Code = 0 is success
func (c *ClientV3) V3PayGiftActivityMerchantDelete(ctx context.Context, activityId string, bm gopay.BodyMap) (*PayGiftActivityMerchantDeleteRsp, error) {
	if err := bm.CheckEmptyError("merchant_id_list", "delete_request_no"); err != nil {
		return nil, err
	}
	uri := fmt.Sprintf(v3PayGiftActivityMerchantDelete, activityId)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &PayGiftActivityMerchantDeleteRsp{Code: Success, SignInfo: si, Response: &PayGiftActivityMerchantDelete{}}
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
