package qq

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"hash"
	"io/ioutil"
	"strings"

	"github.com/iGoogle-ink/gopay"
)

// 添加QQ证书 Byte 数组
//    certFile：apiclient_cert.pem byte数组
//    keyFile：apiclient_key.pem byte数组
//    pkcs12File：apiclient_cert.p12 byte数组
func (w *Client) AddCertFileByte(certFile, keyFile, pkcs12File []byte) {
	w.mu.Lock()
	w.CertFile = certFile
	w.KeyFile = keyFile
	w.Pkcs12File = pkcs12File
	w.mu.Unlock()
}

// 添加QQ证书 Path 路径
//    certFilePath：apiclient_cert.pem 路径
//    keyFilePath：apiclient_key.pem 路径
//    pkcs12FilePath：apiclient_cert.p12 路径
//    返回err
func (w *Client) AddCertFilePath(certFilePath, keyFilePath, pkcs12FilePath string) (err error) {
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
	w.mu.Lock()
	w.CertFile = cert
	w.KeyFile = key
	w.Pkcs12File = pkcs
	w.mu.Unlock()
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
