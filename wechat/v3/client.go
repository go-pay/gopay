package wecaht

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sync"

	"github.com/iGoogle-ink/gopay"
)

// ClientV3 微信支付 V3
type ClientV3 struct {
	MchId       string
	SerialNo    string
	privateKey  *rsa.PrivateKey
	DebugSwitch gopay.DebugSwitch
	rwlock      sync.RWMutex
}

// NewClientV3 初始化微信客户端 V3
//	mchId：商户ID
// 	serialNo 商户证书的证书序列号
//	certContent：私钥 apiclient_key.pem 内容
func NewClientV3(mchId, serialNo string, privateKey *rsa.PrivateKey) (client *ClientV3) {
	return &ClientV3{
		MchId:       mchId,
		SerialNo:    serialNo,
		privateKey:  privateKey,
		DebugSwitch: gopay.DebugOff,
	}
}

// 微信 v3 鉴权请求头 Authorization: xxx 获取
func (c *ClientV3) Authorization(method, path, timestamp, nonceStr, body string) (string, error) {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()
	_str := method + "\n" + path + "\n" + timestamp + "\n" + nonceStr + "\n" + body + "\n"
	sign, err := c.rsa2(_str)
	if err != nil {
		return "", err
	}
	return Authorization + ` mchid="` + c.MchId + `",nonce_str="` + nonceStr + `",timestamp="` + timestamp + `",serial_no="` + c.SerialNo + `",signature="` + sign + `"`, nil
}

func (c *ClientV3) rsa2(str string) (string, error) {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()
	if c.privateKey == nil {
		return "", errors.New("privateKey can't be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, h.Sum(nil))
	return string(base64.StdEncoding.EncodeToString(result)), err
}
