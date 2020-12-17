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
	"github.com/iGoogle-ink/gopay/wechat"
)

const (
	AUTH_MSG = "WECHATPAY2-SHA256-RSA2048"
)

// ClientV3 微信支付 V3
type ClientV3 struct {
	privateKey *rsa.PrivateKey
	mchID      string
	serialNo   string
	rwlock     sync.RWMutex
}

// 初始化微信客户端 V3
//	appId：应用ID
//	mchId：商户ID
//	ApiKey：API秘钥值
//	IsProd：是否是正式环境
func NewClientV3(appId, mchId, apiKey string) (client *wechat.Client) {
	return &wechat.Client{
		AppId:       appId,
		MchId:       mchId,
		ApiKey:      apiKey,
		DebugSwitch: gopay.DebugOff,
	}
}

// NewClient 微信支付 V3
// privateKey 私钥 apiclient_key.pem
// mechID 商户号
// serialNo 商户证书的证书序列号
func NewClient(key *rsa.PrivateKey, mechID, serialNo string) *ClientV3 {
	return &ClientV3{
		privateKey: key,
		mchID:      mechID,
		serialNo:   serialNo,
	}
}

// 微信 v3 鉴权请求头 Authorization: xxx 获取
func (c *ClientV3) Authorization(method, url, timestamp, randomStr, body string) (string, error) {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()
	_str := method + "\n" + url + "\n" + timestamp + "\n" + randomStr + "\n" + body + "\n"
	sign, err := c.rsa2(_str)
	if err != nil {
		return "", err
	}
	return AUTH_MSG + ` mchid="` + c.mchID + `",nonce_str="` + randomStr + `",timestamp="` + timestamp + `",serial_no="` + c.serialNo + `",signature="` + sign + `"`, nil
}

func (c *ClientV3) rsa2(str string) (string, error) {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()
	if c.privateKey == nil {
		return "", errors.New("privateKey cant be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, h.Sum(nil))
	return string(base64.StdEncoding.EncodeToString(result)), err
}
