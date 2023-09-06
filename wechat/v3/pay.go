package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// APP下单API
// Code = 0 is success
func (c *ClientV3) V3TransactionApp(ctx context.Context, bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("mchid") == util.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiApp, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// JSAPI/小程序下单API
// Code = 0 is success
func (c *ClientV3) V3TransactionJsapi(ctx context.Context, bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("mchid") == util.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiJsapi, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// Native下单API
// Code = 0 is success
func (c *ClientV3) V3TransactionNative(ctx context.Context, bm gopay.BodyMap) (wxRsp *NativeRsp, err error) {
	if bm.GetString("mchid") == util.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiNative, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Native)
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

// H5下单API
// Code = 0 is success
func (c *ClientV3) V3TransactionH5(ctx context.Context, bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	if bm.GetString("mchid") == util.NULL {
		bm.Set("mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiH5, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(H5Url)
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

// 查询订单API
// Code = 0 is success
func (c *ClientV3) V3TransactionQueryOrder(ctx context.Context, orderNoType OrderNoType, orderNo string) (wxRsp *QueryOrderRsp, err error) {
	var uri string
	switch orderNoType {
	case TransactionId:
		uri = fmt.Sprintf(v3ApiQueryOrderTransactionId, orderNo) + "?mchid=" + c.Mchid
	case OutTradeNo:
		uri = fmt.Sprintf(v3ApiQueryOrderOutTradeNo, orderNo) + "?mchid=" + c.Mchid
	default:
		return nil, errors.New("unsupported order number type")
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &QueryOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(QueryOrder)
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

// 关闭订单API
// Code = 0 is success
func (c *ClientV3) V3TransactionCloseOrder(ctx context.Context, tradeNo string) (wxRsp *CloseOrderRsp, err error) {
	url := fmt.Sprintf(v3ApiCloseOrder, tradeNo)
	bm := make(gopay.BodyMap)
	bm.Set("mchid", c.Mchid)
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &CloseOrderRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
