package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
)

// APP下单API
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_1.shtml
func (c *ClientV3) V3TransactionApp(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := gotil.GetRandomString(32)
	authorization, err := c.Authorization(MethodPost, v3ApiPayApp, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, hs, bs, err := c.doProdPost(bm, v3ApiPayApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{StatusCode: res.StatusCode, Headers: hs}
	if res.StatusCode != http.StatusOK {
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	wxRsp.Response = new(Prepay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// JSAPI/小程序下单API
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_2.shtml
func (c *ClientV3) V3TransactionJsapi(bm gopay.BodyMap) (wxRsp *PrepayRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := gotil.GetRandomString(32)
	authorization, err := c.Authorization(MethodPost, v3ApiJsapi, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, hs, bs, err := c.doProdPost(bm, v3ApiJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{StatusCode: res.StatusCode, Headers: hs}
	if res.StatusCode != http.StatusOK {
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	wxRsp.Response = new(Prepay)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// Native下单API
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_3.shtml
func (c *ClientV3) V3TransactionNative(bm gopay.BodyMap) (wxRsp *NativeRsp, err error) {
	ts := time.Now().Unix()
	nonceStr := gotil.GetRandomString(32)
	authorization, err := c.Authorization(MethodPost, v3ApiNative, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, hs, bs, err := c.doProdPost(bm, v3ApiNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{StatusCode: res.StatusCode, Headers: hs}
	if res.StatusCode != http.StatusOK {
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	wxRsp.Response = new(Native)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// H5下单API
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_4.shtml
func (c *ClientV3) V3TransactionH5(bm gopay.BodyMap) (wxRsp *H5Rsp, err error) {
	ts := time.Now().Unix()
	nonceStr := gotil.GetRandomString(32)
	authorization, err := c.Authorization(MethodPost, v3ApiH5, nonceStr, ts, bm)
	if err != nil {
		return nil, err
	}
	res, hs, bs, err := c.doProdPost(bm, v3ApiH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{StatusCode: res.StatusCode, Headers: hs}
	if res.StatusCode != http.StatusOK {
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	wxRsp.Response = new(H5Url)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 查询订单API
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_5.shtml
func (c *ClientV3) V3TransactionQueryOrder(orderNoType OrderNoType, orderNo string) (wxRsp *QueryOrderRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = gotil.GetRandomString(32)
		uri      string
	)

	switch orderNoType {
	case TransactionId:
		uri = fmt.Sprintf(v3ApiQueryOrderTransactionId, orderNo) + "?mchid=" + c.Mchid
	case OutTradeNo:
		uri = fmt.Sprintf(v3ApiQueryOrderOutTradeNo, orderNo) + "?mchid=" + c.Mchid
	default:
		return nil, errors.New("unsupported order number type")
	}
	authorization, err := c.Authorization(MethodGet, uri, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}
	res, hs, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &QueryOrderRsp{StatusCode: res.StatusCode, Headers: hs}
	if res.StatusCode != http.StatusOK {
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	wxRsp.Response = new(QueryOrder)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}
