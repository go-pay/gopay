package wechat

import (
	"context"
	"crypto/rsa"
	"errors"
	"strings"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/smap"
	"github.com/go-pay/xlog"
)

// ClientV3 微信支付 V3
type ClientV3 struct {
	Mchid         string
	ApiV3Key      []byte
	SerialNo      string
	WxSerialNo    string // 微信支付公钥ID（微信平台证书序列号）
	proxyHost     string // 代理host地址
	autoSign      bool
	hc            *xhttp.Client
	privateKey    *rsa.PrivateKey
	wxPublicKey   *rsa.PublicKey
	ctx           context.Context
	DebugSwitch   gopay.DebugSwitch
	requestIdFunc xhttp.RequestIdHandler
	logger        xlog.XLogger
	SnCertMap     smap.Map[string, *rsa.PublicKey] // key: serial_no
}

// NewClientV3 初始化微信客户端 V3
// mchid：商户ID 或者服务商模式的 sp_mchid
// serialNo：商户API证书的证书序列号
// apiV3Key：APIv3Key，商户平台获取
// privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
func NewClientV3(mchid, serialNo, apiV3Key, privateKey string) (client *ClientV3, err error) {
	if mchid == gopay.NULL || serialNo == gopay.NULL || apiV3Key == gopay.NULL || privateKey == gopay.NULL {
		return nil, gopay.MissWechatInitParamErr
	}
	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	client = &ClientV3{
		Mchid:         mchid,
		SerialNo:      serialNo,
		ApiV3Key:      []byte(apiV3Key),
		privateKey:    priKey,
		ctx:           context.Background(),
		DebugSwitch:   gopay.DebugOff,
		logger:        logger,
		requestIdFunc: defaultRequestIdFunc,
		hc:            xhttp.NewClient(),
	}
	return client, nil
}

// SetRequestIdFunc 设置自定义请求头 Request-Id 处理方法
func (c *ClientV3) SetRequestIdFunc(requestIdFunc xhttp.RequestIdHandler) {
	if requestIdFunc != nil {
		c.requestIdFunc = requestIdFunc
	}
}

// AutoVerifySign 开启请求完自动验签功能（默认不开启，推荐开启）
// 开启自动验签，自动开启每12小时一次轮询，请求最新证书操作
// autoRefresh：是否自动刷新证书，默认 true：自动刷新
// 说明：开启自动验签功能，会自动获取并刷新微信平台证书，并同步验签。
// 注意：此方法仅支持微信平台证书验签，不支持微信支付公钥验签，如使用微信支付公钥验签请使用 AutoVerifySignByPublicKey() 方法
func (c *ClientV3) AutoVerifySign(autoRefresh ...bool) (err error) {
	wxSerialNo, certMap, err := c.GetAndSelectNewestCert()
	if err != nil {
		return err
	}
	for sn, cert := range certMap {
		// decode cert
		pubKey, err := xpem.DecodePublicKey([]byte(cert))
		if err != nil {
			return err
		}
		c.SnCertMap.Store(sn, pubKey)
		if sn == wxSerialNo {
			c.wxPublicKey = pubKey
		}
	}
	c.WxSerialNo = wxSerialNo
	if len(autoRefresh) == 1 && !autoRefresh[0] {
		return nil
	}
	c.autoSign = true
	go c.autoCheckCertProc()
	return nil
}

// AutoVerifySignByPublicKey 微信支付公钥自动验签
// wxPublicKeyContent：微信支付公钥内容[]byte
// wxPublicKeyID：微信支付公钥ID
func (c *ClientV3) AutoVerifySignByPublicKey(wxPublicKeyContent []byte, wxPublicKeyID string) (err error) {
	return c.AutoVerifySignByCert(wxPublicKeyContent, wxPublicKeyID)
}

// Deprecated
// AutoVerifySignByCert 微信平台证书自动验签（微信支付公钥验签同样适用）
// wxPublicKeyContent：微信公钥证书文件内容[]byte
// wxPublicKeyID：微信公钥证书ID，即 证书序列号
func (c *ClientV3) AutoVerifySignByCert(wxPublicKeyContent []byte, wxPublicKeyID string) (err error) {
	pubKey, err := xpem.DecodePublicKey(wxPublicKeyContent)
	if err != nil {
		return err
	}
	if pubKey == nil {
		return errors.New("xpem.DecodePublicKey() failed, pubKey is nil")
	}
	c.SnCertMap.Store(wxPublicKeyID, pubKey)
	c.wxPublicKey = pubKey
	c.WxSerialNo = wxPublicKeyID
	c.autoSign = true
	return nil
}

// SetBodySize 设置http response body size(MB)
func (c *ClientV3) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}

// SetHttpClient 设置自定义的xhttp.Client
func (c *ClientV3) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		c.hc = client
	}
}

// SetLogger 设置自定义 logger
func (c *ClientV3) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		c.logger = logger
	}
}

// SetProxyHost 设置的 ProxyHost
// 使用场景：
// 1. 部署环境无法访问互联网，可以通过代理服务器访问
// 2. 不设置则默认 https://api.mch.weixin.qq.com
func (c *ClientV3) SetProxyHost(proxyHost string) {
	before, found := strings.CutSuffix(proxyHost, "/")
	if found {
		c.proxyHost = before
		return
	}
	c.proxyHost = proxyHost
}

// GetProxyHost 返回当前的 ProxyHost
func (c *ClientV3) GetProxyHost() string {
	return c.proxyHost
}
