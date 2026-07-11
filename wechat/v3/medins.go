package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// 医保自费混合收款下单
// Code = 0 is success
func (c *ClientV3) V3MedInsOrder(ctx context.Context, bm gopay.BodyMap) (wxRsp *MedInsOrderRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3MedInsOrder, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3MedInsOrder, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MedInsOrderRsp{Code: Success, SignInfo: si, Response: new(MedInsOrder)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 使用医保自费混合订单号查看下单结果
// Code = 0 is success
func (c *ClientV3) V3MedInsOrderQueryByMixNo(ctx context.Context, mixTradeNo string) (wxRsp *MedInsQueryOrderRsp, err error) {
	uri := fmt.Sprintf(v3MedInsOrderQueryByMixNo, mixTradeNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MedInsQueryOrderRsp{Code: Success, SignInfo: si, Response: new(MedInsQueryOrder)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}

// 使用商户订单号查看下单结果
// Code = 0 is success
func (c *ClientV3) V3MedInsOrderQueryByOutNo(ctx context.Context, outTradeNo string) (wxRsp *MedInsQueryOrderRsp, err error) {
	uri := fmt.Sprintf(v3MedInsOrderQueryByOutNo, outTradeNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MedInsQueryOrderRsp{Code: Success, SignInfo: si, Response: new(MedInsQueryOrder)}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &wxRsp.ErrResponse)
		return wxRsp, nil
	}
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, c.verifySyncSign(si)
}
