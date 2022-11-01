package apple

import (
	"context"
	"crypto/rsa"
	"sync"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xpem"
)

// Client AppleClient
type Client struct {
	Mchid       string
	ApiV3Key    []byte
	SerialNo    string
	WxSerialNo  string
	autoSign    bool
	bodySize    int // http response body size(MB), default is 10MB
	rwMu        sync.RWMutex
	privateKey  *rsa.PrivateKey
	wxPublicKey *rsa.PublicKey
	ctx         context.Context
	DebugSwitch gopay.DebugSwitch
	SnCertMap   map[string]*rsa.PublicKey // key: serial_no
}

// NewClient 初始化Apple客户端
// mchid：商户ID 或者服务商模式的 sp_mchid
// serialNo：商户API证书的证书序列号
// apiV3Key：APIv3Key，商户平台获取
// privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
func NewClient(mchid, serialNo, apiV3Key, privateKey string) (client *Client, err error) {
	if mchid == util.NULL || serialNo == util.NULL || apiV3Key == util.NULL || privateKey == util.NULL {
		return nil, gopay.MissWechatInitParamErr
	}
	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	client = &Client{
		Mchid:       mchid,
		SerialNo:    serialNo,
		ApiV3Key:    []byte(apiV3Key),
		privateKey:  priKey,
		ctx:         context.Background(),
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}
