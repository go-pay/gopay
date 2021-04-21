package wechat

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"strings"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
	"github.com/iGoogle-ink/gopay/pkg/xhttp"
	"golang.org/x/crypto/pkcs12"
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

// 添加微信证书文件 Path 路径或证书内容
//	注意：只传pem证书或只传pkcs12证书均可，无需3个证书全传
//	certFilePath：apiclient_cert.pem 路径或内容
//	keyFilePath：apiclient_key.pem 路径或内容
//	pkcs12FilePath：apiclient_cert.p12 路径或内容
//	返回err
func (w *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath interface{}) (err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return
	}
	var config *tls.Config
	if config, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return
	}
	w.mu.Lock()
	w.certificate = &config.Certificates[0]
	w.mu.Unlock()
	return
}

// Deprecated
// 推荐使用 AddCertPemFileContent() 或 AddCertPkcs12FileContent()
// 添加微信证书内容
//	certFileContent：apiclient_cert.pem 内容
//	keyFileContent：apiclient_key.pem 内容
//	pkcs12FileContent：apiclient_cert.p12 内容
//	返回err
func (w *Client) AddCertFileContent(certFileContent, keyFileContent, pkcs12FileContent []byte) (err error) {
	return w.AddCertFilePath(certFileContent, keyFileContent, pkcs12FileContent)
}

// 添加微信pem证书内容或证书文件Path
func (w *Client) AddCertPemFileContent(certFileContent, keyFileContent []byte) (err error) {
	return w.AddCertFilePath(certFileContent, keyFileContent, nil)
}

// 添加微信pkcs12证书内容证书文件Path
func (w *Client) AddCertPkcs12FileContent(pkcs12FileContent []byte) (err error) {
	return w.AddCertFilePath(nil, nil, pkcs12FileContent)
}

func (w *Client) addCertConfig(certFile, keyFile, pkcs12File interface{}) (tlsConfig *tls.Config, err error) {
	if certFile == nil && keyFile == nil && pkcs12File == nil {
		w.mu.RLock()
		defer w.mu.RUnlock()
		if w.certificate != nil {
			tlsConfig = &tls.Config{
				Certificates:       []tls.Certificate{*w.certificate},
				InsecureSkipVerify: true,
			}
			return tlsConfig, nil
		}
		return nil, errors.New("cert parse failed")
	}

	var (
		certPem, keyPem []byte
		certificate     tls.Certificate
	)
	if certFile != nil && keyFile != nil {
		if _, ok := certFile.([]byte); ok {
			certPem = certFile.([]byte)
		} else {
			certPem, err = ioutil.ReadFile(certFile.(string))
		}
		if _, ok := keyFile.([]byte); ok {
			keyPem = keyFile.([]byte)
		} else {
			keyPem, err = ioutil.ReadFile(keyFile.(string))
		}
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
	} else if pkcs12File != nil {
		var pfxData []byte
		if _, ok := pkcs12File.([]byte); ok {
			pfxData = pkcs12File.([]byte)
		} else {
			if pfxData, err = ioutil.ReadFile(pkcs12File.(string)); err != nil {
				return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
			}
		}
		blocks, err := pkcs12.ToPEM(pfxData, w.MchId)
		if err != nil {
			return nil, fmt.Errorf("pkcs12.ToPEM：%w", err)
		}
		for _, b := range blocks {
			keyPem = append(keyPem, pem.EncodeToMemory(b)...)
		}
		certPem = keyPem
	}
	if certPem != nil && keyPem != nil {
		if certificate, err = tls.X509KeyPair(certPem, keyPem); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%w", err)
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			InsecureSkipVerify: true,
		}
		return tlsConfig, nil
	}
	return nil, errors.New("cert files must all nil or all not nil")
}

func checkCertFilePathOrContent(certFile, keyFile, pkcs12File interface{}) error {
	if certFile == nil && keyFile == nil && pkcs12File == nil {
		return nil
	}
	if certFile != nil && keyFile != nil {
		files := map[string]interface{}{"certFile": certFile, "keyFile": keyFile}
		for varName, v := range files {
			switch v.(type) {
			case string:
				if v.(string) == util.NULL {
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
	} else if pkcs12File != nil {
		switch pkcs12File.(type) {
		case string:
			if pkcs12File.(string) == util.NULL {
				return errors.New("pkcs12File is empty")
			}
		case []byte:
			if len(pkcs12File.([]byte)) == 0 {
				return errors.New("pkcs12File is empty")
			}
		default:
			return errors.New("pkcs12File type error")
		}
		return nil
	} else {
		return errors.New("certFile keyFile must all nil or all not nil")
	}
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
	if sandBoxApiKey, err = getSanBoxKey(mchId, util.GetRandomString(32), apiKey, SignType_MD5); err != nil {
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
	_, errs := xhttp.NewClient().Type(xhttp.TypeXML).Post(sandboxGetSignKey).SendString(GenerateXml(reqs)).EndStruct(keyResponse)
	if len(errs) > 0 {
		return util.NULL, errs[0]
	}
	if keyResponse.ReturnCode == "FAIL" {
		return util.NULL, errors.New(keyResponse.ReturnMsg)
	}
	return keyResponse.SandboxSignkey, nil
}

// 生成请求XML的Body体
func GenerateXml(bm gopay.BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return util.NULL
	}
	return string(bs)
}
