package wechat

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"sync"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/aes"
	"github.com/iGoogle-ink/gotil/errgroup"
	"github.com/iGoogle-ink/gotil/xhttp"
	"github.com/iGoogle-ink/gotil/xlog"
	"github.com/pkg/errors"
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
	if block == nil {
		return nil, errors.Errorf("pem.Decode(%s),error", string(pkContent))
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
		privateKey:  pk,
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}

// 获取微信平台证书
func (c *ClientV3) GetPlatformCerts(apiV3Key string) (certs *PlatformCertRsp, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = gotil.GetRandomString(32)
		eg       = new(errgroup.Group)
		mu       sync.Mutex
	)

	authorization, err := c.Authorization(MethodGet, v3GetCerts, nonceStr, ts, nil)
	if err != nil {
		return nil, err
	}

	res, hs, bs, err := c.doProdGet(v3GetCerts, authorization)
	if err != nil {
		return nil, err
	}
	certs = &PlatformCertRsp{StatusCode: res.StatusCode, Headers: hs}
	if res.StatusCode != http.StatusOK {
		certs.Error = string(bs)
		return certs, nil
	}
	certRsp := new(PlatformCert)
	if err = json.Unmarshal(bs, certRsp); err != nil {
		return nil, errors.Errorf("json.Unmarshal(%s)：%+v", string(bs), err)
	}
	for _, v := range certRsp.Data {
		cert := v
		if cert.EncryptCertificate != nil {
			ec := cert.EncryptCertificate
			eg.Go(func(ctx context.Context) error {
				pubKey, err := c.DecryptCerts(ec.Ciphertext, ec.Nonce, ec.AssociatedData, apiV3Key)
				if err != nil {
					return err
				}
				pci := &PlatformCertItem{
					EffectiveTime: cert.EffectiveTime,
					ExpireTime:    cert.ExpireTime,
					PublicKey:     pubKey,
					SerialNo:      cert.SerialNo,
				}
				mu.Lock()
				certs.Certs = append(certs.Certs, pci)
				mu.Unlock()
				return nil
			})
		}
	}
	if err = eg.Wait(); err != nil {
		return nil, err
	}
	return certs, nil
}

func (c *ClientV3) DecryptCerts(ciphertext, nonce, additional, apiV3Key string) (wxCerts string, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return "", errors.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	return string(decrypt), nil
}

// v3 鉴权请求Header
func (c *ClientV3) Authorization(method, path, nonceStr string, timestamp int64, bm gopay.BodyMap) (string, error) {
	var (
		jb = ""
	)
	if bm != nil {
		if bm.Get("appid") == gotil.NULL {
			bm.Set("appid", c.Appid)
		}
		if bm.Get("mchid") == gotil.NULL {
			bm.Set("mchid", c.Mchid)
		}
		jb = bm.JsonBody()
	}
	ts := gotil.Int642String(timestamp)
	_str := method + "\n" + path + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return "", err
	}
	return Authorization + ` mchid="` + c.Mchid + `",nonce_str="` + nonceStr + `",timestamp="` + ts + `",serial_no="` + c.SerialNo + `",signature="` + sign + `"`, nil
}

func (c *ClientV3) rsaSign(str string) (string, error) {
	if c.privateKey == nil {
		return "", errors.New("privateKey can't be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return "", errors.Errorf("rsa.SignPKCS1v15(),err:%+v", err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

func (c *ClientV3) doProdPost(bm gopay.BodyMap, path, authorization string) (res *http.Response, hs *Headers, bs []byte, err error) {
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
		return nil, nil, nil, errs[0]
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %s", res.Header)
	}
	hs = &Headers{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
	}
	return res, hs, bs, nil
}

func (c *ClientV3) doProdGet(uri, authorization string) (res *http.Response, hs *Headers, bs []byte, err error) {
	var url = v3BaseUrlCh + uri

	httpClient := xhttp.NewClient()
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Wechat_V3_Url: %s", url)
		xlog.Debugf("Wechat_V3_Authorization: %s", authorization)
	}
	httpClient.Header.Add(HeaderAuthorization, authorization)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes()
	if len(errs) > 0 {
		return nil, nil, nil, errs[0]
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Errorf("StatusCode = %d, ResponseBody = %s", res.StatusCode, string(bs))
		xlog.Debugf("Wechat_Headers: %s", res.Header)
	}
	hs = &Headers{
		HeaderTimestamp: res.Header.Get(HeaderTimestamp),
		HeaderNonce:     res.Header.Get(HeaderNonce),
		HeaderSignature: res.Header.Get(HeaderSignature),
		HeaderSerial:    res.Header.Get(HeaderSerial),
	}
	return res, hs, bs, nil
}
