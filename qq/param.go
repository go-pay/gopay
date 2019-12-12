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

// 添加QQ证书 Byte 数组
//    certFile：apiclient_cert.pem byte数组
//    keyFile：apiclient_key.pem byte数组
//    pkcs12File：apiclient_cert.p12 byte数组
func (q *Client) AddCertFileByte(certFile, keyFile, pkcs12File []byte) {
	q.mu.Lock()
	q.CertFile = certFile
	q.KeyFile = keyFile
	q.Pkcs12File = pkcs12File
	q.mu.Unlock()
}

// 添加QQ证书 Path 路径
//    certFilePath：apiclient_cert.pem 路径
//    keyFilePath：apiclient_key.pem 路径
//    pkcs12FilePath：apiclient_cert.p12 路径
//    返回err
func (q *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath string) (err error) {
	cert, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return err
	}
	key, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return err
	}
	pkcs, err := ioutil.ReadFile(pkcs12FilePath)
	if err != nil {
		return err
	}
	q.mu.Lock()
	q.CertFile = cert
	q.KeyFile = key
	q.Pkcs12File = pkcs
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
	var (
		pkcs        []byte
		certificate tls.Certificate
		pkcsPool    = x509.NewCertPool()
	)

	if certFilePath == gopay.NULL && keyFilePath == gopay.NULL && pkcs12FilePath == gopay.NULL {
		q.mu.RLock()
		pkcsPool.AppendCertsFromPEM(q.Pkcs12File)
		certificate, err = tls.X509KeyPair(q.CertFile, q.KeyFile)
		q.mu.RUnlock()
		if err != nil {
			return nil, fmt.Errorf("tls.X509KeyPair：%s", err.Error())
		}
		tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{certificate},
			RootCAs:            pkcsPool,
			InsecureSkipVerify: true}
		return tlsConfig, nil
	}

	if certFilePath != gopay.NULL && keyFilePath != gopay.NULL && pkcs12FilePath != gopay.NULL {
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
