package douyin

import (
	"context"
	"crypto/rsa"
	"errors"
	"strings"
	"sync"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/xlog"
)

// Client 抖音支付客户端
// mchid：直连商户号，由抖音支付生成
// serialNo：商户 API 证书序列号
// apiKey：接口加密密钥（32 字节 AES-256-GCM 对称密钥），用于回调解密
// privateKey：商户 API 证书私钥内容
type Client struct {
	Mchid       string
	SerialNo    string
	ApiKey      []byte
	privateKey  *rsa.PrivateKey
	proxyHost   string
	autoSign    bool
	hc          *xhttp.Client
	ctx         context.Context
	DebugSwitch gopay.DebugSwitch
	logger      xlog.XLogger

	// RespTimestampWindow 响应时间戳允许的最大偏差（秒），默认 300（与官方 SDK 一致）
	// 设为 <= 0 可关闭校验（例如本地时钟不准时）
	RespTimestampWindow int64

	// 抖音支付平台公钥（多序列号并存，key: serial_no）
	certMu    sync.RWMutex
	certMap   map[string]*rsa.PublicKey
	newestSno string // 最新有效的平台证书序列号，用于加密敏感字段时的 Douyinpay-Serial 头
}

// NewClient 初始化抖音支付客户端
// mchid：直连商户号
// serialNo：商户 API 证书序列号
// apiKey：接口加密密钥（32 字节），用于回调 AES-256-GCM 解密
// privateKey：商户 API 证书私钥 apiclient_key.pem 读取后的字符串内容
func NewClient(mchid, serialNo, apiKey, privateKey string) (client *Client, err error) {
	if mchid == gopay.NULL || serialNo == gopay.NULL || apiKey == gopay.NULL || privateKey == gopay.NULL {
		return nil, gopay.MissDouyinInitParamErr
	}
	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	client = &Client{
		Mchid:               mchid,
		SerialNo:            serialNo,
		ApiKey:              []byte(apiKey),
		privateKey:          priKey,
		ctx:                 context.Background(),
		DebugSwitch:         gopay.DebugOff,
		logger:              logger,
		hc:                  xhttp.NewClient(),
		certMap:             make(map[string]*rsa.PublicKey),
		RespTimestampWindow: DefaultRespTimestampWindow,
	}
	return client, nil
}

// SetPlatformCert 设置抖音支付平台证书公钥
// pubKeyContent：平台证书 PEM 内容（在抖音支付商户平台【产品中心】->【密钥管理】->【接口加签证书】下载）
// serialNo：平台证书序列号
// 支持多次调用注册多张平台证书，用于验签时按 Douyinpay-Serial 匹配
func (c *Client) SetPlatformCert(pubKeyContent []byte, serialNo string) (err error) {
	if len(pubKeyContent) == 0 || serialNo == gopay.NULL {
		return errors.New("pubKeyContent or serialNo is empty")
	}
	pubKey, err := xpem.DecodePublicKey(pubKeyContent)
	if err != nil {
		return err
	}
	if pubKey == nil {
		return errors.New("xpem.DecodePublicKey() failed, pubKey is nil")
	}
	c.certMu.Lock()
	c.certMap[serialNo] = pubKey
	c.newestSno = serialNo
	c.autoSign = true
	c.certMu.Unlock()
	return nil
}

// PlatformCertMap 获取平台证书副本（key: serial_no）
func (c *Client) PlatformCertMap() map[string]*rsa.PublicKey {
	c.certMu.RLock()
	defer c.certMu.RUnlock()
	dst := make(map[string]*rsa.PublicKey, len(c.certMap))
	for k, v := range c.certMap {
		dst[k] = v
	}
	return dst
}

// NewestPlatformSerialNo 返回最新一次通过 SetPlatformCert 注册的证书序列号
// 用于上送敏感信息加密时携带的 Douyinpay-Serial 请求头
func (c *Client) NewestPlatformSerialNo() string {
	c.certMu.RLock()
	defer c.certMu.RUnlock()
	return c.newestSno
}

// getPlatformKey 按序列号取出平台公钥
func (c *Client) getPlatformKey(serialNo string) (*rsa.PublicKey, bool) {
	c.certMu.RLock()
	defer c.certMu.RUnlock()
	pk, ok := c.certMap[serialNo]
	return pk, ok
}

// SetBodySize 设置 http response body size (MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}

// SetHttpClient 设置自定义的 xhttp.Client
func (c *Client) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		c.hc = client
	}
}

// SetLogger 设置自定义 logger
func (c *Client) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		c.logger = logger
	}
}

// SetProxyHost 设置代理 Host（用于部署环境无法直接访问 https://api.douyinpay.com 的场景）
func (c *Client) SetProxyHost(proxyHost string) {
	before, found := strings.CutSuffix(proxyHost, "/")
	if found {
		c.proxyHost = before
		return
	}
	c.proxyHost = proxyHost
}

// GetProxyHost 返回当前代理 Host
func (c *Client) GetProxyHost() string {
	return c.proxyHost
}

// GetHttpClient 获取 xhttp.Client，用于自定义调整 http 请求参数
func (c *Client) GetHttpClient() *xhttp.Client {
	return c.hc
}
