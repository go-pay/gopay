package alipay

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/crypto/xrsa"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/xlog"
)

// ClientV3 支付宝 V3
type ClientV3 struct {
	AppId              string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	AppAuthToken       string
	IsProd             bool
	aesKey             string // biz_content 加密的 AES KEY
	proxyHost          string // 代理host地址
	ivKey              []byte
	privateKey         *rsa.PrivateKey
	aliPayPublicKey    *rsa.PublicKey // 支付宝证书公钥内容 alipayPublicCert.crt
	DebugSwitch        gopay.DebugSwitch
	logger             xlog.XLogger
	requestIdFunc      xhttp.RequestIdHandler
	hc                 *xhttp.Client
}

// NewClientV3 初始化支付宝客户端 V3
// appid：应用ID
// privateKey：应用私钥，支持PKCS1和PKCS8
// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
func NewClientV3(appid, privateKey string, isProd bool) (client *ClientV3, err error) {
	if appid == gopay.NULL || privateKey == gopay.NULL {
		return nil, gopay.MissAlipayInitParamErr
	}
	key := xrsa.FormatAlipayPrivateKey(privateKey)
	priKey, err := xpem.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	client = &ClientV3{
		AppId:         appid,
		IsProd:        isProd,
		privateKey:    priKey,
		DebugSwitch:   gopay.DebugOff,
		logger:        logger,
		requestIdFunc: defaultRequestIdFunc,
		hc:            xhttp.NewClient(),
	}
	return client, nil
}

// 设置自定义RequestId生成函数
func (a *ClientV3) SetRequestIdFunc(requestIdFunc xhttp.RequestIdHandler) {
	if requestIdFunc != nil {
		a.requestIdFunc = requestIdFunc
	}
}

// 应用公钥证书内容设置 app_cert_sn、alipay_root_cert_sn、alipay_cert_sn
// appCertContent：应用公钥证书文件内容
// alipayRootCertContent：支付宝根证书文件内容
// alipayPublicCertContent：支付宝公钥证书文件内容
func (a *ClientV3) SetCert(appCertContent, alipayRootCertContent, alipayPublicCertContent []byte) (err error) {
	appCertSn, err := getCertSN(appCertContent)
	if err != nil {
		return fmt.Errorf("get app_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	rootCertSn, err := getRootCertSN(alipayRootCertContent)
	if err != nil {
		return fmt.Errorf("get alipay_root_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	publicCertSn, err := getCertSN(alipayPublicCertContent)
	if err != nil {
		return fmt.Errorf("get alipay_cert_sn return err, but alse return alipay client. err: %w", err)
	}

	// alipay public key
	pubKey, err := xpem.DecodePublicKey(alipayPublicCertContent)
	if err != nil {
		return fmt.Errorf("decode alipayPublicCertContent err: %w", err)
	}

	a.AppCertSN = appCertSn
	a.AliPayRootCertSN = rootCertSn
	a.AliPayPublicCertSN = publicCertSn
	a.aliPayPublicKey = pubKey
	return nil
}

// SetBodySize 设置http response body size(MB)
func (a *ClientV3) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		a.hc.SetBodySize(sizeMB)
	}
}

// SetHttpClient 设置自定义的xhttp.Client
func (a *ClientV3) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		a.hc = client
	}
}

// SetLogger 设置自定义的logger
func (a *ClientV3) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		a.logger = logger
	}
}

// SetAESKey 设置 biz_content 的AES加密key，设置此参数默认开启 biz_content 参数加密
// 注意：目前不可用，设置后会报错
func (a *ClientV3) SetAESKey(aesKey string) {
	a.aesKey = aesKey
	a.ivKey = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

// SetProxyHost 设置的 ProxyHost
// 使用场景：
// 1. 部署环境无法访问互联网，可以通过代理服务器访问
// 2. 不设置则默认 https://api.mch.weixin.qq.com
func (a *ClientV3) SetProxyHost(proxyHost string) {
	before, found := strings.CutSuffix(proxyHost, "/")
	if found {
		a.proxyHost = before
		return
	}
	a.proxyHost = proxyHost
}

// GetProxyHost 返回当前的 ProxyHost
func (a *ClientV3) GetProxyHost() string {
	return a.proxyHost
}
