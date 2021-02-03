package qq

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
	"github.com/iGoogle-ink/gopay/pkg/util"
)

// 添加QQ证书 Path 路径
//	certFilePath：apiclient_cert.pem 路径
//	keyFilePath：apiclient_key.pem 路径
//	pkcs12FilePath：apiclient_cert.p12 路径
//	返回err
func (w *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath interface{}) (err error) {
	if err = checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return err
	}
	cert, err := ioutil.ReadFile(certFilePath.(string))
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile：%w", err)
	}
	key, err := ioutil.ReadFile(keyFilePath.(string))
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile：%w", err)
	}
	pkcs, err := ioutil.ReadFile(pkcs12FilePath.(string))
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

func checkCertFilePath(certFilePath, keyFilePath, pkcs12FilePath interface{}) error {
	if certFilePath != nil && keyFilePath != nil && pkcs12FilePath != nil {
		if v, ok := certFilePath.(string); !ok || v == util.NULL {
			return errors.New("certFilePath not string type or is null string")
		}
		if v, ok := keyFilePath.(string); !ok || v == util.NULL {
			return errors.New("keyFilePath not string type or is null string")
		}
		if v, ok := pkcs12FilePath.(string); !ok || v == util.NULL {
			return errors.New("pkcs12FilePath not string type or is null string")
		}
		return nil
	}
	if !(certFilePath == nil && keyFilePath == nil && pkcs12FilePath == nil) {
		return errors.New("cert paths must all nil or all not nil")
	}
	return nil
}

// 生成请求XML的Body体
func generateXml(bm gopay.BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return util.NULL
	}
	return string(bs)
}

// 获取QQ支付正式环境Sign值
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

func (q *Client) addCertConfig(certFilePath, keyFilePath, pkcs12FilePath interface{}) (tlsConfig *tls.Config, err error) {
	if certFilePath == nil && keyFilePath == nil && pkcs12FilePath == nil {
		q.mu.RLock()
		defer q.mu.RUnlock()
		if q.certPool != nil {
			tlsConfig = &tls.Config{
				Certificates:       []tls.Certificate{q.certificate},
				RootCAs:            q.certPool,
				InsecureSkipVerify: true,
			}
			return tlsConfig, nil
		}
	}

	if certFilePath != nil && keyFilePath != nil && pkcs12FilePath != nil {
		cert, err := ioutil.ReadFile(certFilePath.(string))
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
		key, err := ioutil.ReadFile(keyFilePath.(string))
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
		pkcs, err := ioutil.ReadFile(pkcs12FilePath.(string))
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
	return nil, errors.New("cert paths must all nil or all not nil")
}
