package wechat

import (
	"context"
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
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 预授权状态查询
// Code = 0 is success
func (c *ClientV3) V3PalmServiceOpenidQuery(ctx context.Context, bm gopay.BodyMap) (wxRsp *PalmServiceOpenidQueryRsp, err error) {
	uri := v3PalmServiceOpenidQuery + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PalmServiceOpenidQueryRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
