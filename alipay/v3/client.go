package alipay

import (
	"crypto/rsa"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/xlog"
	"time"
)

// ClientV3 支付宝 V3
type ClientV3 struct {
	//Mchid       string
	//ApiV3Key    []byte
	//SerialNo    string
	//WxSerialNo  string
	//autoSign    bool
	//rwMu        sync.RWMutex
	//hc          *xhttp.Client
	//privateKey  *rsa.PrivateKey
	//wxPublicKey *rsa.PublicKey
	//ctx         context.Context
	//DebugSwitch gopay.DebugSwitch
	//logger      xlog.XLogger
	//SnCertMap   map[string]*rsa.PublicKey // key: serial_no

	AppId              string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	ReturnUrl          string
	NotifyUrl          string
	Charset            string
	SignType           string
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
