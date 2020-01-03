package wechat

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

	"github.com/iGoogle-ink/gopay"
)

type Country int

// 设置支付国家（默认：中国国内）
//    根据支付地区情况设置国家
//    country：<China：中国国内，China2：中国国内（冗灾方案），SoutheastAsia：东南亚，Other：其他国家>
func (w *Client) SetCountry(country Country) (client *Client) {
	w.mu.Lock()
	switch country {
	case China:
		w.BaseURL = baseUrlCh
	case China2:
		w.BaseURL = baseUrlCh2
	case SoutheastAsia:
		w.BaseURL = baseUrlHk
	case Other:
		w.BaseURL = baseUrlUs
	default:
		w.BaseURL = baseUrlCh
	}
	w.mu.Unlock()
	return w
}

// 添加微信证书 Path 路径
//    certFilePath：apiclient_cert.pem 路径
//    keyFilePath：apiclient_key.pem 路径
//    pkcs12FilePath：apiclient_cert.p12 路径
//    返回err
func (w *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath string) (err error) {
	cert, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile：%w", err)
	}
	key, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile：%w", err)
	}
	pkcs, err := ioutil.ReadFile(pkcs12FilePath)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile：%w", err)
	}
	certificate, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return fmt.Errorf("tls.LoadX509KeyPair：%w", err)
	}
	pkcsPool := x509.NewCertPool()
	pkcsPool.AppendCertsFromPEM(pkcs)
	w.mu.Lock()
	w.certificate = certificate
	w.certPool = pkcsPool
	w.mu.Unlock()
	return nil
}

func (w *Client) addCertConfig(certFilePath, keyFilePath, pkcs12FilePath string) (tlsConfig *tls.Config, err error) {

	if certFilePath == gopay.NULL && keyFilePath == gopay.NULL && pkcs12FilePath == gopay.NULL {
		w.mu.RLock()
		defer w.mu.RUnlock()
		if &w.certificate != nil && w.certPool != nil {
			tlsConfig = &tls.Config{
				Certificates:       []tls.Certificate{w.certificate},
				RootCAs:            w.certPool,
				InsecureSkipVerify: true,
			}
			return tlsConfig, nil
		}
	}

	if certFilePath != gopay.NULL && keyFilePath != gopay.NULL && pkcs12FilePath != gopay.NULL {
		cert, err := ioutil.ReadFile(certFilePath)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
		key, err := ioutil.ReadFile(keyFilePath)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
		pkcsPool := x509.NewCertPool()
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.X509KeyPair(cert, key)
		if err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%w", err)
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
func getReleaseSign(apiKey string, signType string, bm gopay.BodyMap) (sign string) {
	var h hash.Hash
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(bm.EncodeWeChatSignParams(apiKey)))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// 获取微信支付沙箱环境Sign值
func getSignBoxSign(mchId, apiKey string, bm gopay.BodyMap) (sign string, err error) {
	var (
		sandBoxApiKey string
		h             hash.Hash
	)
	if sandBoxApiKey, err = getSanBoxKey(mchId, gopay.GetRandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	h = md5.New()
	h.Write([]byte(bm.EncodeWeChatSignParams(sandBoxApiKey)))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// 从微信提供的接口获取：SandboxSignKey
func getSanBoxKey(mchId, nonceStr, apiKey, signType string) (key string, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("mch_id", mchId)
	bm.Set("nonce_str", nonceStr)
	//沙箱环境：获取沙箱环境ApiKey
	if key, err = getSanBoxSignKey(mchId, nonceStr, getReleaseSign(apiKey, signType, bm)); err != nil {
		return
	}
	return
}

// 从微信提供的接口获取：SandboxSignKey
func getSanBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(gopay.BodyMap)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)

	keyResponse := new(getSignKeyResponse)
	_, errs := gopay.NewHttpClient().Type(gopay.TypeXML).Post(sandboxGetSignKey).SendString(generateXml(reqs)).EndStruct(keyResponse)
	if len(errs) > 0 {
		return gopay.NULL, errs[0]
	}
	if keyResponse.ReturnCode == "FAIL" {
		return gopay.NULL, errors.New(keyResponse.ReturnMsg)
	}
	return keyResponse.SandboxSignkey, nil
}

// 生成请求XML的Body体
func generateXml(bm gopay.BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return gopay.NULL
	}
	return string(bs)
}
