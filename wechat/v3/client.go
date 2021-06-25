package wechat

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
)

// ClientV3 微信支付 V3
type ClientV3 struct {
	Appid       string
	Mchid       string
	SerialNo    string
	apiV3Key    string
	wxPkContent []byte
	autoSign    bool
	privateKey  *rsa.PrivateKey
	DebugSwitch gopay.DebugSwitch
	rwlock      sync.RWMutex
}

// NewClientV3 初始化微信客户端 V3
//	appid：appid 或者服务商模式的 sp_appid
//	mchid：商户ID 或者服务商模式的 sp_mchid
// 	serialNo：商户证书的证书序列号
//	apiV3Key：apiV3Key，商户平台获取
//	pkContent：私钥 apiclient_key.pem 读取后的字符串内容
func NewClientV3(appid, mchid, serialNo, apiV3Key, pkContent string) (client *ClientV3, err error) {
	var (
		pk *rsa.PrivateKey
		ok bool
	)
	block, _ := pem.Decode([]byte(pkContent))
	if block == nil {
		return nil, errors.New(fmt.Sprintf("pem.Decode(%s),error", pkContent))
	}
	pk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		pk, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	if pk == nil {
		pk, ok = pk8.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
	}
	client = &ClientV3{
		Appid:       appid,
		Mchid:       mchid,
		SerialNo:    serialNo,
		apiV3Key:    apiV3Key,
		privateKey:  pk,
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}

// AutoVerifySign 开启请求完自动验签功能（默认不开启，推荐开启）
//	注意：未获取到微信平台公钥时，不要开启，请调用 client.GetPlatformCerts() 获取微信平台证书公钥
func (c *ClientV3) AutoVerifySign(wxPkContent []byte) {
	if wxPkContent != nil {
		c.wxPkContent = wxPkContent
		c.autoSign = true
	}
}

func (c *ClientV3) doProdPost(bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		jb, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_V3_RequestBody: %s", jb)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderSerial, c.SerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes()
	if len(errs) > 0 {
		return nil, nil, nil, errs[0]
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

func (c *ClientV3) doProdGet(uri, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + uri

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_Url: %s", url)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderSerial, c.SerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes()
	if len(errs) > 0 {
		return nil, nil, nil, errs[0]
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

func (c *ClientV3) doProdPut(bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		jb, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_V3_RequestBody: %s", jb)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderSerial, c.SerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Put(url).SendBodyMap(bm).EndBytes()
	if len(errs) > 0 {
		return nil, nil, nil, errs[0]
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

func (c *ClientV3) doProdDelete(bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		jb, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_V3_RequestBody: %s", jb)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderSerial, c.SerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Delete(url).SendBodyMap(bm).EndBytes()
	if len(errs) > 0 {
		return nil, nil, nil, errs[0]
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

func (c *ClientV3) doProdPostFile(bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_RequestBody: %s", bm.GetString("meta"))
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderSerial, c.SerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeMultipartFormData).Post(url).SendMultipartBodyMap(bm).EndBytes()
	if len(errs) > 0 {
		return nil, nil, nil, errs[0]
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

func (c *ClientV3) doProdPatch(bm gopay.BodyMap, path, authorization string) (res *http.Response, si *SignInfo, bs []byte, err error) {
	var url = v3BaseUrlCh + path

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		jb, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_V3_RequestBody: %s", jb)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add(HeaderSerial, c.SerialNo)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Patch(url).SendBodyMap(bm).EndBytes()
	if len(errs) > 0 {
		return nil, nil, nil, errs[0]
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
