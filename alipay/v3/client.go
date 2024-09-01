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

// NewClientV3 初始化支付宝客户端 V3
// mchid：商户ID 或者服务商模式的 sp_mchid
// serialNo：商户API证书的证书序列号
// apiV3Key：APIv3Key，商户平台获取
// privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
//func NewClientV3(mchid, serialNo, apiV3Key, privateKey string) (client *ClientV3, err error) {
//	if mchid == gopay.NULL || serialNo == gopay.NULL || apiV3Key == gopay.NULL || privateKey == gopay.NULL {
//		return nil, gopay.MissWechatInitParamErr
//	}
//	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))
//	if err != nil {
//		return nil, err
//	}
//	logger := xlog.NewLogger()
//	logger.SetLevel(xlog.DebugLevel)
//	client = &ClientV3{
//		Mchid:       mchid,
//		SerialNo:    serialNo,
//		ApiV3Key:    []byte(apiV3Key),
//		privateKey:  priKey,
//		ctx:         context.Background(),
//		DebugSwitch: gopay.DebugOff,
//		logger:      logger,
//		hc:          xhttp.NewClient(),
//	}
//	return client, nil
//}
