package wechat

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/util"
)

var defaultRequestIdFunc = &requestIdFunc{}

type requestIdFunc struct{}

func (d *requestIdFunc) RequestId() string {
	return fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix())
}

func (c *ClientV3) doProdPostWithHeader(ctx context.Context, headerMap map[string]string, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	if c.proxyHost != "" {
		url = c.proxyHost + path
	}
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPostWithHost(ctx context.Context, bm gopay.BodyMap, host, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = host + path
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPost(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	if c.proxyHost != "" {
		url = c.proxyHost + path
	}
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdGet(ctx context.Context, uri, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + uri
	if c.proxyHost != "" {
		url = c.proxyHost + uri
	}
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Get(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPut(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	if c.proxyHost != "" {
		url = c.proxyHost + path
	}
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Put(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdDelete(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	if c.proxyHost != "" {
		url = c.proxyHost + path
	}
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Delete(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPostFile(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	if c.proxyHost != "" {
		url = c.proxyHost + path
	}
	req := c.hc.Req(xhttp.TypeMultipartFormData)
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Post(url).SendMultipartBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPatch(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	if c.proxyHost != "" {
		url = c.proxyHost + path
	}
	req := c.hc.Req() // default json
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestIdFunc.RequestId())
	req.Header.Add(HeaderSerial, c.WxSerialNo)
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Url: %s", url)
		c.logger.Debugf("Wechat_V3_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Wechat_V3_Req_Headers: %#v", req.Header)
	}
	res, bs, err = req.Patch(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	si = &SignInfo{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
		SignBody:        string(bs),
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Wechat_V3_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Wechat_V3_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Wechat_V3_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}
