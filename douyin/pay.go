package douyin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util/js"
)

// AppOrder App 支付下单
// Code = 0 表示 HTTP 200 成功
func (c *Client) AppOrder(ctx context.Context, bm gopay.BodyMap) (dyRsp *PrepayRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, appOrder, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, appOrder, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &PrepayRsp{Code: Success, SignInfo: si, Response: new(Prepay)}
	if res.StatusCode != http.StatusOK {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	if err = json.Unmarshal(bs, dyRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return dyRsp, c.verifySyncSign(si)
}

// JsapiOrder JSAPI 支付下单
func (c *Client) JsapiOrder(ctx context.Context, bm gopay.BodyMap) (dyRsp *PrepayRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, jsapiOrder, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, jsapiOrder, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &PrepayRsp{Code: Success, SignInfo: si, Response: new(Prepay)}
	if res.StatusCode != http.StatusOK {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	if err = json.Unmarshal(bs, dyRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return dyRsp, c.verifySyncSign(si)
}

// H5Order H5 支付下单
// 响应字段仅 h5_url（直接跳转支付页），不返回 prepay_id
func (c *Client) H5Order(ctx context.Context, bm gopay.BodyMap) (dyRsp *H5OrderRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, h5Order, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, h5Order, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &H5OrderRsp{Code: Success, SignInfo: si, Response: new(H5Order)}
	if res.StatusCode != http.StatusOK {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	if err = json.Unmarshal(bs, dyRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return dyRsp, c.verifySyncSign(si)
}

// NativeOrder Native 支付下单
// 响应字段仅 code_url（收银台二维码链接）
func (c *Client) NativeOrder(ctx context.Context, bm gopay.BodyMap) (dyRsp *NativeOrderRsp, err error) {
	if bm.GetString("mchid") == gopay.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, nativeOrder, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, nativeOrder, authorization)
	if err != nil {
		return nil, err
	}
	dyRsp = &NativeOrderRsp{Code: Success, SignInfo: si, Response: new(NativeOrder)}
	if res.StatusCode != http.StatusOK {
		dyRsp.Code = res.StatusCode
		dyRsp.Error = string(bs)
		_ = js.UnmarshalBytes(bs, &dyRsp.ErrResponse)
		return dyRsp, nil
	}
	if err = json.Unmarshal(bs, dyRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return dyRsp, c.verifySyncSign(si)
}
