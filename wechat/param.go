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
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xhttp"
)

type Country int

// 设置支付国家（默认：中国国内）
//	根据支付地区情况设置国家
//	country：<China：中国国内，China2：中国国内（冗灾方案），SoutheastAsia：东南亚，Other：其他国家>
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
//	certFilePath：apiclient_cert.pem 路径
//	keyFilePath：apiclient_key.pem 路径
//	pkcs12FilePath：apiclient_cert.p12 路径
//	返回err
func (w *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath interface{}) (err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return
	}
	var config *tls.Config
	if config, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return
	}
	w.mu.Lock()
	w.certificate = config.Certificates[0]
	w.certPool = config.RootCAs
	w.mu.Unlock()
	return
}

// 添加微信证书内容
//	certFileContent：apiclient_cert.pem 内容
//	keyFileContent：apiclient_key.pem 内容
//	pkcs12FileContent：apiclient_cert.p12 内容
//	返回err
func (w *Client) AddCertFileContent(certFileContent, keyFileContent, pkcs12FileContent []byte) (err error) {
	if err = checkCertFilePath(certFileContent, keyFileContent, pkcs12FileContent); err != nil {
		return
	}
	var config *tls.Config
	if config, err = w.addCertConfig(certFileContent, keyFileContent, pkcs12FileContent); err != nil {
		return
	}
	w.mu.Lock()
	w.certificate = config.Certificates[0]
	w.certPool = config.RootCAs
	w.mu.Unlock()
	return
}

func (w *Client) addCertConfig(certFile, keyFile, pkcs12File interface{}) (tlsConfig *tls.Config, err error) {
	if certFile == nil && keyFile == nil && pkcs12File == nil {
		w.mu.RLock()
		defer w.mu.RUnlock()
		if w.certPool != nil {
			tlsConfig = &tls.Config{
				Certificates:       []tls.Certificate{w.certificate},
				RootCAs:            w.certPool,
				InsecureSkipVerify: true,
			}
			return tlsConfig, nil
		}
	}

	if certFile != nil && keyFile != nil && pkcs12File != nil {
		var (
			cert, key, pkcs []byte
			certificate     tls.Certificate
		)
		if _, ok := certFile.([]byte); ok {
			cert = certFile.([]byte)
		} else {
			cert, err = ioutil.ReadFile(certFile.(string))
		}
		if _, ok := keyFile.([]byte); ok {
			key = keyFile.([]byte)
		} else {
			key, err = ioutil.ReadFile(keyFile.(string))
		}
		if _, ok := pkcs12File.([]byte); ok {
			pkcs = pkcs12File.([]byte)
		} else {
			pkcs, err = ioutil.ReadFile(pkcs12File.(string))
		}
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
		pkcsPool := x509.NewCertPool()
		pkcsPool.AppendCertsFromPEM(pkcs)
		if certificate, err = tls.X509KeyPair(cert, key); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%w", err)
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true,
		}
		return tlsConfig, nil
	}
	return nil, errors.New("cert files must all nil or all not nil")
}

func checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath interface{}) error {
	if certFilePath != nil && keyFilePath != nil && pkcs12FilePath != nil {
		files := map[string]interface{}{
			"certFilePath":   certFilePath,
			"keyFilePath":    keyFilePath,
			"pkcs12FilePath": pkcs12FilePath,
		}
		for varName, v := range files {
			switch v.(type) {
			case string:
				if v.(string) == gotil.NULL {
					return fmt.Errorf("%s is empty", varName)
				}
			case []byte:
				if len(v.([]byte)) == 0 {
					return fmt.Errorf("%s is empty", varName)
				}
			default:
				return fmt.Errorf("%s type error", varName)
			}
		}
		return nil
	}
	if !(certFilePath == nil && keyFilePath == nil && pkcs12FilePath == nil) {
		return errors.New("cert paths must all nil or all not nil")
	}
	return nil
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
	if sandBoxApiKey, err = getSanBoxKey(mchId, gotil.GetRandomString(32), apiKey, SignType_MD5); err != nil {
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
	// 沙箱环境：获取沙箱环境ApiKey
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
	_, errs := xhttp.NewClient().Type(xhttp.TypeXML).Post(sandboxGetSignKey).SendString(generateXml(reqs)).EndStruct(keyResponse)
	if len(errs) > 0 {
		return gotil.NULL, errs[0]
	}
	if keyResponse.ReturnCode == "FAIL" {
		return gotil.NULL, errors.New(keyResponse.ReturnMsg)
	}
	return keyResponse.SandboxSignkey, nil
}

// 生成请求XML的Body体
func generateXml(bm gopay.BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return gotil.NULL
	}
	return string(bs)
}
