package wechat

import (
	"context"
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
	"os"
	"strings"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/util"
	"golang.org/x/crypto/pkcs12"
)

type Country int

// 设置支付国家（默认：中国国内）
// 根据支付地区情况设置国家
// country：<China：中国国内，China2：中国国内（冗灾方案），SoutheastAsia：东南亚，Other：其他国家>
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

// SetProxyUrl 设置代理 Url
// 使用场景：
// 1. 部署环境无法访问互联网，可以通过代理服务器访问
func (w *Client) SetProxyUrl(proxyUrl string) (client *Client) {
	w.mu.Lock()
	w.BaseURL = proxyUrl
	w.mu.Unlock()
	return w
}

// 添加微信pem证书文件路径
// certFilePath：apiclient_cert.pem 文件路径
// keyFilePath：apiclient_key.pem 文件路径
func (w *Client) AddCertPemFilePath(certFilePath, keyFilePath string) (err error) {
	return w.addCertFileContentOrPath(certFilePath, keyFilePath, nil)
}

// 添加微信pkcs12证书文件路径
// pkcs12FilePath：apiclient_cert.p12 文件路径
func (w *Client) AddCertPkcs12FilePath(pkcs12FilePath string) (err error) {
	return w.addCertFileContentOrPath(nil, nil, pkcs12FilePath)
}

// 添加微信pem证书内容[]byte
// certFileContent：apiclient_cert.pem 证书内容[]byte
// keyFileContent：apiclient_key.pem 证书内容[]byte
func (w *Client) AddCertPemFileContent(certFileContent, keyFileContent []byte) (err error) {
	return w.addCertFileContentOrPath(certFileContent, keyFileContent, nil)
}

// 添加微信pkcs12证书内容[]byte
// p12FileContent：apiclient_cert.p12 证书内容[]byte
func (w *Client) AddCertPkcs12FileContent(p12FileContent []byte) (err error) {
	return w.addCertFileContentOrPath(nil, nil, p12FileContent)
}

// 添加微信证书文件 Path 路径或证书内容
// 注意：只传pem证书或只传pkcs12证书均可，无需3个证书全传
func (w *Client) addCertFileContentOrPath(certFile, keyFile, pkcs12File any) (err error) {
	if err = checkCertFilePathOrContent(certFile, keyFile, pkcs12File); err != nil {
		return
	}
	config, err := w.addCertConfig(certFile, keyFile, pkcs12File)
	if err != nil {
		return
	}
	w.tlsHc.SetHttpTLSConfig(config)
	return
}

func (w *Client) addCertConfig(certFile, keyFile, pkcs12File any) (tlsConfig *tls.Config, err error) {
	if certFile == nil && keyFile == nil && pkcs12File == nil {
		return nil, errors.New("cert parse failed or nil")
	}

	var (
		certPem, keyPem []byte
		certificate     tls.Certificate
	)
	if certFile != nil && keyFile != nil {
		if _, ok := certFile.([]byte); ok {
			certPem = certFile.([]byte)
		} else {
			certPem, err = os.ReadFile(certFile.(string))
		}
		if _, ok := keyFile.([]byte); ok {
			keyPem = keyFile.([]byte)
		} else {
			keyPem, err = os.ReadFile(keyFile.(string))
		}
		if err != nil {
			return nil, fmt.Errorf("os.ReadFile: %w", err)
		}
	} else if pkcs12File != nil {
		var pfxData []byte
		if _, ok := pkcs12File.([]byte); ok {
			pfxData = pkcs12File.([]byte)
		} else {
			if pfxData, err = os.ReadFile(pkcs12File.(string)); err != nil {
				return nil, fmt.Errorf("os.ReadFile: %w", err)
			}
		}
		blocks, err := pkcs12.ToPEM(pfxData, w.MchId)
		if err != nil {
			return nil, fmt.Errorf("pkcs12.ToPEM: %w", err)
		}
		for _, b := range blocks {
			keyPem = append(keyPem, pem.EncodeToMemory(b)...)
		}
		certPem = keyPem
	}
	if certPem != nil && keyPem != nil {
		if certificate, err = tls.X509KeyPair(certPem, keyPem); err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair: %w", err)
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			InsecureSkipVerify: true,
		}
		return tlsConfig, nil
	}
	return nil, errors.New("cert files must all nil or all not nil")
}

func checkCertFilePathOrContent(certFile, keyFile, pkcs12File any) error {
	if certFile == nil && keyFile == nil && pkcs12File == nil {
		return nil
	}
	if certFile != nil && keyFile != nil {
		files := map[string]any{"certFile": certFile, "keyFile": keyFile}
		for varName, v := range files {
			switch v := v.(type) {
			case string:
				if v == gopay.NULL {
					return fmt.Errorf("%s is empty", varName)
				}
			case []byte:
				if len(v) == 0 {
					return fmt.Errorf("%s is empty", varName)
				}
			default:
				return fmt.Errorf("%s type error", varName)
			}
		}
		return nil
	} else if pkcs12File != nil {
		switch pkcs12File := pkcs12File.(type) {
		case string:
			if pkcs12File == gopay.NULL {
				return errors.New("pkcs12File is empty")
			}
		case []byte:
			if len(pkcs12File) == 0 {
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
func GetReleaseSign(apiKey string, signType string, bm gopay.BodyMap) (sign string) {
	var h hash.Hash
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, []byte(apiKey))
	} else {
		h = md5.New()
	}
	h.Write([]byte(bm.EncodeWeChatSignParams(apiKey)))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// 获取微信支付正式环境Sign值
func (w *Client) getReleaseSign(apiKey string, signType string, bm gopay.BodyMap) (sign string) {
	signParams := bm.EncodeWeChatSignParams(apiKey)
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Request_SignStr: %s", signParams)
	}
	var h hash.Hash
	if signType == SignType_HMAC_SHA256 {
		h = w.sha256Hash
	} else {
		h = w.md5Hash
	}
	w.mu.Lock()
	defer func() {
		h.Reset()
		w.mu.Unlock()
	}()
	h.Write([]byte(signParams))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// 获取微信支付沙箱环境Sign值
func GetSandBoxSign(ctx context.Context, mchId, apiKey string, bm gopay.BodyMap) (sign string, err error) {
	var (
		sandBoxApiKey string
		h             hash.Hash
	)
	if sandBoxApiKey, err = getSanBoxKey(ctx, mchId, util.RandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	h = md5.New()
	h.Write([]byte(bm.EncodeWeChatSignParams(sandBoxApiKey)))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// 获取微信支付沙箱环境Sign值
func (w *Client) getSandBoxSign(ctx context.Context, mchId, apiKey string, bm gopay.BodyMap) (sign string, err error) {
	var (
		sandBoxApiKey string
		h             hash.Hash
	)
	if sandBoxApiKey, err = getSanBoxKey(ctx, mchId, util.RandomString(32), apiKey, SignType_MD5); err != nil {
		return
	}
	h = md5.New()
	signParams := bm.EncodeWeChatSignParams(sandBoxApiKey)
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("Wechat_Request_SignStr: %s", signParams)
	}
	h.Write([]byte(signParams))
	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

// 从微信提供的接口获取：SandboxSignKey
func getSanBoxKey(ctx context.Context, mchId, nonceStr, apiKey, signType string) (key string, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("mch_id", mchId)
	bm.Set("nonce_str", nonceStr)
	// 沙箱环境：获取沙箱环境ApiKey
	if key, err = getSanBoxSignKey(ctx, mchId, nonceStr, GetReleaseSign(apiKey, signType, bm)); err != nil {
		return
	}
	return
}

// 从微信提供的接口获取：SandboxSignKey
func getSanBoxSignKey(ctx context.Context, mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(gopay.BodyMap)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)

	keyResponse := new(getSignKeyResponse)
	_, err = xhttp.NewClient().Req(xhttp.TypeXML).Post(sandboxGetSignKey).SendString(GenerateXml(reqs)).EndStruct(ctx, keyResponse)
	if err != nil {
		return gopay.NULL, err
	}
	if keyResponse.ReturnCode == "FAIL" {
		return gopay.NULL, errors.New(keyResponse.ReturnMsg)
	}
	return keyResponse.SandboxSignkey, nil
}

// 生成请求XML的Body体
func GenerateXml(bm gopay.BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return gopay.NULL
	}
	return string(bs)
}
