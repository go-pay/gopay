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
)

// 添加QQ证书 Path 路径
//    certFilePath：apiclient_cert.pem 路径
//    keyFilePath：apiclient_key.pem 路径
//    pkcs12FilePath：apiclient_cert.p12 路径
//    返回err
func (q *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath string) (err error) {
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
	q.mu.Lock()
	q.certificate = certificate
	q.certPool = pkcsPool
	q.mu.Unlock()
	return nil
}

// 生成请求XML的Body体
func generateXml(bm gopay.BodyMap) (reqXml string) {
	bs, err := xml.Marshal(bm)
	if err != nil {
		return gopay.NULL
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

func (q *Client) addCertConfig(certFilePath, keyFilePath, pkcs12FilePath string) (tlsConfig *tls.Config, err error) {

	if certFilePath == gopay.NULL && keyFilePath == gopay.NULL && pkcs12FilePath == gopay.NULL {
		q.mu.RLock()
		defer q.mu.RUnlock()
		if &q.certificate != nil && q.certPool != nil {
			tlsConfig = &tls.Config{
				Certificates:       []tls.Certificate{q.certificate},
				RootCAs:            q.certPool,
				InsecureSkipVerify: true,
			}
			return tlsConfig, nil
		}
	}

	if certFilePath != gopay.NULL && keyFilePath != gopay.NULL && pkcs12FilePath != gopay.NULL {
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%w", err)
		}
		pkcsPool := x509.NewCertPool()
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
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
