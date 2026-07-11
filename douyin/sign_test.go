package douyin

import (
	"testing"

	"github.com/go-pay/xlog"
)

func TestPaySignOfApp(t *testing.T) {
	app, err := client.PaySignOfApp("appid", "prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("app:%#v", app)
}

func TestPaySignOfJSAPI(t *testing.T) {
	jsapi, err := client.PaySignOfJSAPI("appid", "prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("jsapi:%#v", jsapi)
}

func TestVerifySignByPK(t *testing.T) {
	// 抖音支付响应/回调验签使用抖音支付平台证书公钥
	// 验签串 3 行：{timestamp}\n{nonce}\n{body}\n
	//	type SignInfo struct {
	//		HeaderTimestamp string `json:"Douyinpay-Timestamp"`
	//		HeaderNonce     string `json:"Douyinpay-Nonce"`
	//		HeaderSignature string `json:"Douyinpay-Signature"`
	//		HeaderSerial    string `json:"Douyinpay-Serial"`
	//		SignBody        string `json:"sign_body"`
	//	}

	timestamp := ""
	nonce := ""
	signBody := ""
	signature := ""

	// 从已注册的平台证书中获取公钥
	pubKeyMap := client.PlatformCertMap()
	pubKey := pubKeyMap[""] // 传入应答/回调 Header 里的 Douyinpay-Serial

	if err = VerifySignByPK(timestamp, nonce, signBody, signature, pubKey); err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("sign ok")
}
