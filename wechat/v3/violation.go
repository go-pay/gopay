package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 创建商户违规通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ViolationNotifyUrlCreate(ctx context.Context, url string) (wxRsp *ViolationNotifyUrlRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("notify_url", url)
	authorization, err := c.authorization(MethodPost, v3ViolationNotifyUrlCreate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ViolationNotifyUrlCreate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ViolationNotifyUrlRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ViolationNotifyUrl)
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

// 查询商户违规通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ViolationNotifyUrlQuery(ctx context.Context) (wxRsp *ViolationNotifyUrlRsp, err error) {
	authorization, err := c.authorization(MethodGet, v3ViolationNotifyUrlQuery, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, v3ViolationNotifyUrlQuery, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ViolationNotifyUrlRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ViolationNotifyUrl)
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

// 更新商户违规通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ViolationNotifyUrlUpdate(ctx context.Context, url string) (wxRsp *ViolationNotifyUrlRsp, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("notify_url", url)
	authorization, err := c.authorization(MethodPut, v3ViolationNotifyUrlUpdate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPut(ctx, bm, v3ViolationNotifyUrlUpdate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &ViolationNotifyUrlRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(ViolationNotifyUrl)
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

// 删除商户违规通知回调地址API
// Code = 0 is success
func (c *ClientV3) V3ViolationNotifyUrlDelete(ctx context.Context) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodDelete, v3ViolationNotifyUrlDelete, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdDelete(ctx, nil, v3ViolationNotifyUrlDelete, authorization)
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
