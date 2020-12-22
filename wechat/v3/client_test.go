package wecaht

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	client         *ClientV3
	appid          = ""
	mchid          = ""
	serialNo       = ""
	certKeyContent = ``
)

func TestMain(m *testing.M) {
	block, _ := pem.Decode([]byte(certKeyContent))
	pk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		xlog.Error(err)
		os.Exit(1)
	}
	key, ok := pk8.(*rsa.PrivateKey)
	if !ok {
		xlog.Error("parse PKCS8 key error")
		return
	}
	// NewClientV3 初始化微信客户端 V3
	//	appid：appid
	//	mchid：商户ID
	// 	serialNo 商户证书的证书序列号
	//	certContent：私钥 apiclient_key.pem 解析后的私钥key
	client = NewClientV3(appid, mchid, serialNo, key)

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	//err := client.AddCertFilePath(nil, nil, nil)
	//if err != nil {
	//	panic(err)
	//}

	os.Exit(m.Run())
}

func TestV3Jsapi(t *testing.T) {
	tradeNo := gotil.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("description", "测试支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", "https://www.gopay.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", "asdas")
		})

	wxRsp, err := client.V3TransactionJsapi(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp:", wxRsp)
}
