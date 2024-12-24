package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 用户自主录掌&预授权
// Code = 0 is success
func (c *ClientV3) V3PalmServicePreAuthorize(ctx context.Context, bm gopay.BodyMap) (wxRsp *PalmServicePreAuthorizeRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3PalmServicePreAuthorize, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3PalmServicePreAuthorize, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PalmServicePreAuthorizeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PalmServicePreAuthorize)
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

// 预授权状态查询
// Code = 0 is success
func (c *ClientV3) V3PalmServiceOpenidQuery(ctx context.Context, openid, organizationId string) (wxRsp *PalmServiceOpenidQueryRsp, err error) {
	uri := fmt.Sprintf(v3PalmServiceOpenidQuery, openid) + "?organization_id=" + organizationId
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PalmServiceOpenidQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PalmServiceOpenidQuery)
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
