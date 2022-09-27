package wechat

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/pkg/xpem"
)

// ClientV3 微信支付 V3
type ClientV3 struct {
	Mchid       string
	ApiV3Key    []byte
	SerialNo    string
	WxSerialNo  string
	autoSign    bool
	bodySize    int // http response body size(MB), default is 10MB
	privateKey  *rsa.PrivateKey
	ctx         context.Context
	DebugSwitch gopay.DebugSwitch
	CertMap     map[string]*rsa.PublicKey
}

// NewClientV3 初始化微信客户端 V3
//	mchid：商户ID 或者服务商模式的 sp_mchid
// 	serialNo：商户API证书的证书序列号
//	apiV3Key：APIv3Key，商户平台获取
//	privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
func NewClientV3(mchid, serialNo, apiV3Key, privateKey string) (client *ClientV3, err error) {
	if mchid == util.NULL || serialNo == util.NULL || apiV3Key == util.NULL || privateKey == util.NULL {
		return nil, gopay.MissWechatInitParamErr
	}
	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	client = &ClientV3{
		Mchid:       mchid,
		SerialNo:    serialNo,
		ApiV3Key:    []byte(apiV3Key),
		privateKey:  priKey,
		ctx:         context.Background(),
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}

// AutoVerifySign 开启请求完自动验签功能（默认不开启，推荐开启）
//	开启自动验签，在校验时，如果证书不存在，会获取所有已生效的证书
func (c *ClientV3) AutoVerifySign() error {
	certMap, err := c.GetAllValidCert()
	if err != nil {
		return err
	}
	newCertMap := map[string]*rsa.PublicKey{}
	for key, value := range certMap {
		pubKey, err := xpem.DecodePublicKey([]byte(value))
		if err != nil {
			return err
		}
		newCertMap[key] = pubKey
	}
	c.CertMap = newCertMap
	c.autoSign = true
	return nil
}

// SetBodySize 设置http response body size(MB)
func (c *ClientV3) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.bodySize = sizeMB
	}
}

func (c *ClientV3) doProdPostWithHeader(ctx context.Context, headerMap map[string]string, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	for k, v := range headerMap {
		httpClient.Header.Add(k, v)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, c.WxSerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes(ctx)
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
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %#v", res.Header)
		xlog.Debugf("Wechat_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPost(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, c.WxSerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes(ctx)
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
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %#v", res.Header)
		xlog.Debugf("Wechat_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdGet(ctx context.Context, uri, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + uri
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_Url: %s", url)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, c.WxSerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes(ctx)
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
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %#v", res.Header)
		xlog.Debugf("Wechat_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPut(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, c.WxSerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Put(url).SendBodyMap(bm).EndBytes(ctx)
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
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %#v", res.Header)
		xlog.Debugf("Wechat_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdDelete(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, c.WxSerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Delete(url).SendBodyMap(bm).EndBytes(ctx)
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
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %#v", res.Header)
		xlog.Debugf("Wechat_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPostFile(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_RequestBody: %s", bm.GetString("meta"))
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, c.WxSerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeMultipartFormData).Post(url).SendMultipartBodyMap(bm).EndBytes(ctx)
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
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %#v", res.Header)
		xlog.Debugf("Wechat_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}

func (c *ClientV3) doProdPatch(ctx context.Context, bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	httpClient.Header.Add(HeaderSerial, c.WxSerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Patch(url).SendBodyMap(bm).EndBytes(ctx)
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
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %#v", res.Header)
		xlog.Debugf("Wechat_SignInfo: %#v", si)
	}
	return res, si, bs, nil
}
