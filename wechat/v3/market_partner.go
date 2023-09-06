package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 建立合作关系
// Code = 0 is success
func (c *ClientV3) V3PartnershipsBuild(ctx context.Context, idempotencyKey string, bm gopay.BodyMap) (wxRsp *PartnershipsBuildRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3PartnershipsBuild, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPostWithHeader(ctx, map[string]string{"Idempotency-Key": idempotencyKey}, bm, v3PartnershipsBuild, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PartnershipsBuildRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnershipsBuild)
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

// 终止合作关系
// Code = 0 is success
func (c *ClientV3) V3PartnershipsTerminate(ctx context.Context, idempotencyKey string, bm gopay.BodyMap) (wxRsp *PartnershipsTerminateRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3PartnershipsTerminate, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPostWithHeader(ctx, map[string]string{"Idempotency-Key": idempotencyKey}, bm, v3PartnershipsTerminate, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PartnershipsTerminateRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnershipsTerminate)
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

// 查询合作关系列表
// Code = 0 is success
func (c *ClientV3) V3PartnershipsList(ctx context.Context, bm gopay.BodyMap) (wxRsp *PartnershipsListRsp, err error) {
	uri := v3PartnershipsList + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PartnershipsListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnershipsList)
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
