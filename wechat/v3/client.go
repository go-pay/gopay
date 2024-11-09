package wechat

import (
	"context"
	"crypto/rsa"
	"sync"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/xlog"
)

// ClientV3 微信支付 V3
type ClientV3 struct {
	Mchid         string
	ApiV3Key      []byte
	SerialNo      string
	WxSerialNo    string
	autoSign      bool
	rwMu          sync.RWMutex
	hc            *xhttp.Client
	privateKey    *rsa.PrivateKey
	wxPublicKey   *rsa.PublicKey
	ctx           context.Context
	DebugSwitch   gopay.DebugSwitch
	requestIdFunc xhttp.RequestIdHandler
	logger        xlog.XLogger
	SnCertMap     map[string]*rsa.PublicKey // key: serial_no
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

func (c *ClientV3) SetRequestIdFunc(requestIdFunc xhttp.RequestIdHandler) {
	if requestIdFunc != nil {
		c.requestIdFunc = requestIdFunc
	}
}

// AutoVerifySign 开启请求完自动验签功能（默认不开启，推荐开启）
// 开启自动验签，自动开启每12小时一次轮询，请求最新证书操作
func (c *ClientV3) AutoVerifySign(autoRefresh ...bool) (err error) {
	wxSerialNo, certMap, err := c.GetAndSelectNewestCert()
	if err != nil {
		return err
	}
	if len(c.SnCertMap) <= 0 {
		c.SnCertMap = make(map[string]*rsa.PublicKey)
	}
	for sn, cert := range certMap {
		// decode cert
		pubKey, err := xpem.DecodePublicKey([]byte(cert))
		if err != nil {
			return err
		}
		c.SnCertMap[sn] = pubKey
	}
	c.WxSerialNo = wxSerialNo
	c.wxPublicKey = c.SnCertMap[wxSerialNo]
	if len(autoRefresh) == 1 && !autoRefresh[0] {
		return
	}
	c.autoSign = true
	go c.autoCheckCertProc()
	return
}

// wxPublicKeyContent：微信公钥证书文件内容[]byte
// wxPublicKeyID：微信公钥证书ID
func (c *ClientV3) AutoVerifySignByCert(wxPublicKeyContent []byte, wxPublicKeyID string) {
	pubKey, err := xpem.DecodePublicKey(wxPublicKeyContent)
	if err != nil {
		c.logger.Errorf("AutoVerifySignByCert(%s),err:%+v", wxPublicKeyContent, err)
	}
	if pubKey != nil {
		if len(c.SnCertMap) <= 0 {
			c.SnCertMap = make(map[string]*rsa.PublicKey)
		}
		c.SnCertMap[wxPublicKeyID] = pubKey
		c.wxPublicKey = pubKey
		c.WxSerialNo = wxPublicKeyID
		c.autoSign = true
	}
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

func (c *ClientV3) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		c.logger = logger
	}
}
