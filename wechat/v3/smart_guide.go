package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 服务人员注册API
// 注意：入参加密字段数据加密：client.V3EncryptText()
// Code = 0 is success
func (c *ClientV3) V3SmartGuideReg(ctx context.Context, bm gopay.BodyMap) (wxRsp *SmartGuideRegRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3GuideReg, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GuideReg, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &SmartGuideRegRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(SmartGuideReg)
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

// 服务人员分配API
// Code = 0 is success
func (c *ClientV3) V3SmartGuideAssign(ctx context.Context, guideId, tradeNo string) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3GuideAssign, guideId)
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", tradeNo)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
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

// 服务人员查询API
// 注意：入参加密字段数据加密：client.V3EncryptText()，返回参数加密字段解密：client.V3DecryptText()
// Code = 0 is success
func (c *ClientV3) V3SmartGuideQuery(ctx context.Context, bm gopay.BodyMap) (wxRsp *SmartGuideQueryRsp, err error) {
	if err = bm.CheckEmptyError("store_id"); err != nil {
		return nil, err
	}
	uri := v3GuideQuery + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &SmartGuideQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(SmartGuideQuery)
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

// 服务人员信息更新API
// 注意：入参加密字段数据加密：client.V3EncryptText()
// Code = 0 is success
func (c *ClientV3) V3SmartGuideUpdate(ctx context.Context, guideId string, bm gopay.BodyMap) (wxRsp *EmptyRsp, err error) {
	url := fmt.Sprintf(v3GuideUpdate, guideId)
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
