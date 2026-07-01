package douyin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
)

func (c *Client) requestId() string {
	return fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix())
}

func (c *Client) doProdPost(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	url := baseUrlProd + path
	if c.proxyHost != "" {
		url = c.proxyHost + path
	}
	req := c.hc.Req()
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestId())
	if sn := c.NewestPlatformSerialNo(); sn != "" {
		req.Header.Add(HeaderSerial, sn)
	}
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Douyin_Url: %s", url)
		c.logger.Debugf("Douyin_Req_Body: %s", bm.JsonBody())
		c.logger.Debugf("Douyin_Req_Headers: %#v", req.Header)
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
		c.logger.Debugf("Douyin_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Douyin_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Douyin_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *Client) doProdGet(ctx context.Context, uri, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	url := baseUrlProd + uri
	if c.proxyHost != "" {
		url = c.proxyHost + uri
	}
	req := c.hc.Req()
	req.Header.Add(HeaderAuthorization, authorization)
	req.Header.Add(HeaderRequestID, c.requestId())
	if sn := c.NewestPlatformSerialNo(); sn != "" {
		req.Header.Add(HeaderSerial, sn)
	}
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Douyin_Url: %s", url)
		c.logger.Debugf("Douyin_Req_Headers: %#v", req.Header)
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
		c.logger.Debugf("Douyin_Response: %d > %s", res.StatusCode, string(bs))
		c.logger.Debugf("Douyin_Rsp_Headers: %#v", res.Header)
		c.logger.Debugf("Douyin_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}
