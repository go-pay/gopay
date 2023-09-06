package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 点金计划管理API
// Code = 0 is success
func (c *ClientV3) V3GoldPlanManage(ctx context.Context, bm gopay.BodyMap) (wxRsp *GoldPlanManageRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3GoldPlanManage, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanManage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &GoldPlanManageRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(GoldPlanManage)
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

// 商家小票管理API
// Code = 0 is success
func (c *ClientV3) V3GoldPlanBillManage(ctx context.Context, bm gopay.BodyMap) (wxRsp *GoldPlanManageRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3GoldPlanBillManage, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanBillManage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &GoldPlanManageRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(GoldPlanManage)
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

// 同业过滤标签管理API
// Code = 0 is success
func (c *ClientV3) V3GoldPlanFilterManage(ctx context.Context, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3GoldPlanFilterManage, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanFilterManage, authorization)
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

// 开通广告展示API
// Code = 0 is success
func (c *ClientV3) V3GoldPlanOpenAdShow(ctx context.Context, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPATCH, v3GoldPlanOpenAdShow, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPatch(ctx, bm, v3GoldPlanOpenAdShow, authorization)
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

// 关闭广告展示API
// Code = 0 is success
func (c *ClientV3) V3GoldPlanCloseAdShow(ctx context.Context, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPATCH, v3GoldPlanCloseAdShow, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanCloseAdShow, authorization)
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
