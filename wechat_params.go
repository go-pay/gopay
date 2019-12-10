package gopay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"strings"
)

type Country int

// 设置支付国家（默认：中国国内）
//    根据支付地区情况设置国家
//    country：<China：中国国内，China2：中国国内（冗灾方案），SoutheastAsia：东南亚，Other：其他国家>
func (w *WeChatClient) SetCountry(country Country) (client *WeChatClient) {
	w.mu.Lock()
	switch country {
	case China:
		w.BaseURL = wxBaseUrlCh
	case China2:
		w.BaseURL = wxBaseUrlCh2
	case SoutheastAsia:
		w.BaseURL = wxBaseUrlHk
	case Other:
		w.BaseURL = wxBaseUrlUs
	default:
		w.BaseURL = wxBaseUrlCh
	}
	w.mu.Unlock()
	return w
}

// 添加微信证书 Byte 数组
//    certFile：apiclient_cert.pem byte数组
//    keyFile：apiclient_key.pem byte数组
//    pkcs12File：apiclient_cert.p12 byte数组
func (w *WeChatClient) AddCertFileByte(certFile, keyFile, pkcs12File []byte) {
	w.mu.Lock()
	w.CertFile = certFile
	w.KeyFile = keyFile
	w.Pkcs12File = pkcs12File
	w.mu.Unlock()
}

// 添加微信证书 Path 路径
//    certFilePath：apiclient_cert.pem 路径
//    keyFilePath：apiclient_key.pem 路径
//    pkcs12FilePath：apiclient_cert.p12 路径
//    返回err
func (w *WeChatClient) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath string) (err error) {
	var (
		cert, key, pkcs []byte
	)
	if cert, err = ioutil.ReadFile(certFilePath); err != nil {
		return
	}
	if key, err = ioutil.ReadFile(keyFilePath); err != nil {
		return
	}
	if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
		return
	}
	w.mu.Lock()
	w.CertFile = cert
	w.KeyFile = key
	w.Pkcs12File = pkcs
	w.mu.Unlock()
	return
}

func (w *WeChatClient) addCertConfig(certFilePath, keyFilePath, pkcs12FilePath string) (tlsConfig *tls.Config, err error) {
	var (
		pkcs        []byte
		certificate tls.Certificate
		pkcsPool    = x509.NewCertPool()
	)

	if certFilePath == null && keyFilePath == null && pkcs12FilePath == null {
		w.mu.RLock()
		pkcsPool.AppendCertsFromPEM(w.Pkcs12File)
		certificate, err = tls.X509KeyPair(w.CertFile, w.KeyFile)
		w.mu.RUnlock()
		if err != nil {
			return nil, fmt.Errorf("tls.X509KeyPair：%s", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		return tlsConfig, nil
	}

	if certFilePath != null && keyFilePath != null && pkcs12FilePath != null {
		if pkcs, err = ioutil.ReadFile(pkcs12FilePath); err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%s", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.LoadX509KeyPair(certFilePath, keyFilePath); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%s", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		return tlsConfig, nil
	}

	return nil, errors.New("certificate file path must be all input or all input null")
}

// 获取微信支付正式环境Sign值
func getWeChatReleaseSign(apiKey string, signType string, bm BodyMap) (sign string) {
	var h hash.Hash
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(bm.EncodeWeChatSignParams(apiKey)))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// 获取微信支付沙箱环境Sign值
func getWeChatSignBoxSign(mchId, apiKey string, bm BodyMap) (sign string, err error) {
	var (
		sandBoxApiKey string
		h             hash.Hash
	)
	if sandBoxApiKey, err = getSanBoxKey(mchId, GetRandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	h = md5.New()
	h.Write([]byte(bm.EncodeWeChatSignParams(sandBoxApiKey)))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// 从微信提供的接口获取：SandboxSignKey
func getSanBoxKey(mchId, nonceStr, apiKey, signType string) (key string, err error) {
	bm := make(BodyMap)
	bm.Set("mch_id", mchId)
	bm.Set("nonce_str", nonceStr)
	//沙箱环境：获取沙箱环境ApiKey
	if key, err = getSanBoxSignKey(mchId, nonceStr, getWeChatReleaseSign(apiKey, signType, bm)); err != nil {
		return
	}
	return
}

// 从微信提供的接口获取：SandboxSignKey
func getSanBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(BodyMap)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)

	keyResponse := new(getSignKeyResponse)
	_, errs := NewHttpClient().Type(TypeXML).Post(wxSandboxGetsignkey).SendString(generateXml(reqs)).EndStruct(keyResponse)
	if len(errs) > 0 {
		return null, errs[0]
	}
	if keyResponse.ReturnCode == "FAIL" {
		return null, errors.New(keyResponse.ReturnMsg)
	}
	return keyResponse.SandboxSignkey, nil
}

// 生成请求XML的Body体
func generateXml(bm BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return null
	}
	return string(bs)
}
