package wecaht

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xhttp"
	"github.com/iGoogle-ink/gotil/xlog"
)

// ClientV3 微信支付 V3
type ClientV3 struct {
	Appid       string
	Mchid       string
	SerialNo    string
	privateKey  *rsa.PrivateKey
	DebugSwitch gopay.DebugSwitch
	rwlock      sync.RWMutex
}

// NewClientV3 初始化微信客户端 V3
//	appid：appid
//	mchid：商户ID
// 	serialNo 商户证书的证书序列号
//	pkContent：私钥 apiclient_key.pem 读取后的内容
func NewClientV3(appid, mchid, serialNo string, pkContent []byte) (client *ClientV3, err error) {
	var (
		pk *rsa.PrivateKey
		ok bool
	)
	block, _ := pem.Decode(pkContent)
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
		privateKey:  pk,
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}

// 微信 v3 鉴权请求Header
func (c *ClientV3) Authorization(method, path, nonceStr string, timestamp int64, bm gopay.BodyMap) (string, error) {
	func() {
		c.rwlock.RLock()
		defer c.rwlock.RUnlock()
		if bm.Get("appid") == gotil.NULL {
			bm.Set("appid", c.Appid)
		}
		if bm.Get("mchid") == gotil.NULL {
			bm.Set("mchid", c.Mchid)
		}
	}()
	jb := bm.JsonBody()
	ts := gotil.Int642String(timestamp)
	_str := method + "\n" + path + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return "", err
	}
	return Authorization + ` mchid="` + c.Mchid + `",nonce_str="` + nonceStr + `",timestamp="` + ts + `",serial_no="` + c.SerialNo + `",signature="` + sign + `"`, nil
}

func (c *ClientV3) rsaSign(str string) (string, error) {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()
	if c.privateKey == nil {
		return "", errors.New("privateKey can't be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

func (c *ClientV3) doProdPost(bm gopay.BodyMap, path, authorization string) (bs []byte, err error) {
	var url = v3BaseUrlCh + path

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		jb, _ := json.Marshal(bm)
		xlog.Debugf("Wechat_V3_RequestBody: %s", jb)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d, ResponseBody = %s", res.StatusCode, string(bs))
	}
	if strings.Contains(string(bs), "HTML") || strings.Contains(string(bs), "html") {
		return nil, errors.New(string(bs))
	}
	return bs, nil
}
