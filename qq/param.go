package qq

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
	"os"
	"strings"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
	"golang.org/x/crypto/pkcs12"
)

// 添加QQ证书 Path 路径
// certFilePath：apiclient_cert.pem 路径
// keyFilePath：apiclient_key.pem 路径
// pkcs12FilePath：apiclient_cert.p12 路径
// 返回err
func (q *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath any) (err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return err
	}
	config, err := q.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return
	}
	q.tlsHc.SetHttpTLSConfig(config)
	return nil
}

// 添加QQ证书内容
// certFileContent：apiclient_cert.pem 内容
// keyFileContent：apiclient_key.pem 内容
// pkcs12FileContent：apiclient_cert.p12 内容
// 返回err
func (q *Client) AddCertFileContent(certFileContent, keyFileContent, pkcs12FileContent []byte) (err error) {
	return q.AddCertFilePath(certFileContent, keyFileContent, pkcs12FileContent)
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

// 生成请求XML的Body体
func GenerateXml(bm gopay.BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return gopay.NULL
	}
	return string(bs)
}

// 获取QQ支付正式环境Sign值
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

func (q *Client) getReleaseSign(apiKey string, signType string, bm gopay.BodyMap) (sign string) {
	signParams := bm.EncodeWeChatSignParams(apiKey)
	if q.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("QQ_Request_SignStr: %s", signParams)
	}
	var h hash.Hash
	if signType == SignType_HMAC_SHA256 {
		h = q.sha256Hash
	} else {
		h = q.md5Hash
	}
	q.mu.Lock()
	defer func() {
		h.Reset()
		q.mu.Unlock()
	}()
	h.Write([]byte(signParams))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (q *Client) addCertConfig(certFile, keyFile, pkcs12File any) (tlsConfig *tls.Config, err error) {
	if certFile == nil && keyFile == nil && pkcs12File == nil {
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
		blocks, err := pkcs12.ToPEM(pfxData, q.MchId)
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
