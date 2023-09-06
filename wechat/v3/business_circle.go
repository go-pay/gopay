package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 商圈积分同步
// Code = 0 is success
func (c *ClientV3) V3BusinessPointsSync(ctx context.Context, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusinessPointsSync, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusinessPointsSync, authorization)
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

// 商圈积分授权查询
// Code = 0 is success
// openid：path 参数
// bm：query 参数
func (c *ClientV3) V3BusinessAuthPointsQuery(ctx context.Context, openid string, bm gopay.BodyMap) (*BusinessAuthPointsQueryRsp, error) {
	uri := fmt.Sprintf(v3BusinessAuthPointsQuery, openid) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &BusinessAuthPointsQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusinessAuthPointsQuery)
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

// 商圈会员待积分状态查询
// Code = 0 is success
// openid：path 参数
// bm：query 参数
func (c *ClientV3) V3BusinessPointsStatusQuery(ctx context.Context, openid string, bm gopay.BodyMap) (*BusinessPointsStatusQueryRsp, error) {
	uri := fmt.Sprintf(v3BusinessPointsStatusQuery, openid) + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &BusinessPointsStatusQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusinessPointsStatusQuery)
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

// 商圈会员停车状态同步
// Code = 0 is success
func (c *ClientV3) V3BusinessParkingSync(ctx context.Context, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusinessParkingSync, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusinessParkingSync, authorization)
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
