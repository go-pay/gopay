package alipay

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error
	// 普通公钥模式时，验签使用
	//aliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
)

func TestMain(m *testing.M) {

	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err = NewClient(cert.Appid, cert.PrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// Debug开关，输出/关闭日志
	client.DebugSwitch = gopay.DebugOff

	// 配置公共参数
	client.SetCharset("utf-8").
		SetSignType(RSA2).
		// SetAppAuthToken("")
		SetReturnUrl("https://www.fmm.ink").
		SetNotifyUrl("https://www.fmm.ink")

	// 自动同步验签（只支持证书模式）
	// 传入 支付宝公钥证书 alipayPublicCert.crt 内容
	client.AutoVerifySign(cert.AlipayPublicContentRSA2)

	// 传入证书内容
	err := client.SetCertSnByContent(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)
	// 传入证书文件路径
	//err := client.SetCertSnByPath("cert/appPublicCert.crt", "cert/alipayRootCert.crt", "cert/alipayPublicCert.crt")
	if err != nil {
		xlog.Debug("SetCertSn:", err)
		return
	}
	os.Exit(m.Run())
}

func TestClient_PostAliPayAPISelfV2(t *testing.T) {
	bm := make(gopay.BodyMap)

	// 自定义公共参数（根据自己需求，需要独立设置的自行设置，不需要单独设置的，共享client的配置）
	bm.Set("app_id", "appid")
	bm.Set("app_auth_token", "app_auth_token")
	bm.Set("auth_token", "auth_token")

	// biz_content
	bm.SetBodyMap("biz_content", func(bz gopay.BodyMap) {
		bz.Set("subject", "预创建创建订单")
		bz.Set("out_trade_no", util.RandomString(32))
		bz.Set("total_amount", "100")
	})

	aliPsp := new(TradePrecreateResponse)
	err := client.PostAliPayAPISelfV2(ctx, bm, "alipay.trade.precreate", aliPsp)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug(aliPsp.Response)
}

// =================================================

func TestDecryptOpenDataToBodyMap(t *testing.T) {
	data := "MkvuiIZsGOC8S038cu/JIpoRKnF+ZFjoIRGf5d/K4+ctYjCtb/eEkwgrdB5TeH/93bxff1Ylb+SE+UGStlpvcg=="
	key := "TDftre9FpItr46e9BVNJcw=="
	bm, err := DecryptOpenDataToBodyMap(data, key)
	if err != nil {
		xlog.Errorf("DecryptOpenDataToBodyMap(%s,%s),error:%+v", data, key, err)
		return
	}
	xlog.Info("bm:", bm)
}
