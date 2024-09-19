package alipay

import (
	"crypto/rsa"
	"time"

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
	ivKey              []byte
	privateKey         *rsa.PrivateKey
	aliPayPublicKey    *rsa.PublicKey // 支付宝证书公钥内容 alipayPublicCert.crt
	autoSign           bool
	DebugSwitch        gopay.DebugSwitch
	logger             xlog.XLogger
	location           *time.Location
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
		AppId:       appid,
		IsProd:      isProd,
		privateKey:  priKey,
		DebugSwitch: gopay.DebugOff,
		logger:      logger,
		hc:          xhttp.NewClient(),
	}
	return client, nil
}
