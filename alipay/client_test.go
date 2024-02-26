package alipay

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error
	// 普通公钥模式时
	//aliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAk3/P3KX7rqwVEMkbZlY80PqBKxNHsMB4q5wKbpKU2KZffDiilV3bPpHPm1wHaJHKHudmirVWjn65lIspqIMgA4RoU3Cf0sEs8Mf9aAzv0F47KAWb2oBhDyQcYGtHftLG7q2mGBalXK7TJWWA2+DB2UAtfZELBRaVhsaF7bCkGd3GjTSXBZcyXAShgnH94C7yTr2IRbv+SnwedTdCKHxXvoPRvABylO9Krx8MiyJECBBfQScA67SsL+E1MKiMVdMomzlTEQ7W0UEAtzmG4aRMzSp30Lggit1xH7HKyOWgSE3Xy0LT78VKbqGRuCAT6IT5AA9jqSAbOBft7igJFkz9swIDAQAB"
	//appPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw0x6QX0eeW+XRNh4u/BG1RsyuYYsHmqenk4GrPV8ElrUEN6nLRwSXGSIIuIuyCo0t4swCNp9Q54g+AUmpAddeBCYhHKrAYG4n11MnXUYosEe43wzUhd7PaXpxctFlSKhFgYBiX7cQg/O8ThHJJ0H0Qy+k9NNJ2gMnfPypk9/43DstHmKEVQvgQpcgnhlbK8X4thIK4zW0xwHgDhSAeZu9QLvf/cc2PdKmd5xiUUM7J7PtwT7VvbKI27fYOBe1hxxrvpY6vcGGGal8xhNOKD+eiqXAgPgIRYuXqLpPgaYOImjnnH9lIpzjrCO0nsba1T4tonDjp5jv0FaSKwOTAjO6QIDAQAB"
	//appPrivateKey = "MIIEpQIBAAKCAQEAw0x6QX0eeW+XRNh4u/BG1RsyuYYsHmqenk4GrPV8ElrUEN6nLRwSXGSIIuIuyCo0t4swCNp9Q54g+AUmpAddeBCYhHKrAYG4n11MnXUYosEe43wzUhd7PaXpxctFlSKhFgYBiX7cQg/O8ThHJJ0H0Qy+k9NNJ2gMnfPypk9/43DstHmKEVQvgQpcgnhlbK8X4thIK4zW0xwHgDhSAeZu9QLvf/cc2PdKmd5xiUUM7J7PtwT7VvbKI27fYOBe1hxxrvpY6vcGGGal8xhNOKD+eiqXAgPgIRYuXqLpPgaYOImjnnH9lIpzjrCO0nsba1T4tonDjp5jv0FaSKwOTAjO6QIDAQABAoIBAQCjxUUcF7zvXml+XPzZtQLg/97IfsAOfaAn5gxpC66QgcQlpWCRTmIDQnZ6sitCxUnRxJFySy4R/s9szHz7vgVegqQzJSlLqSlV0lpGDAStrr6lSuiKZZB+QNxJdk0SY6irnDu7vjsb1r/VvjjCdkAwyLwjoGSpr/Isnn4TgsUexoBJOVBRvfsVmzNq0oaD12jEHPLnOgyBOUxN54A64mz7H7VBnYhG2TOL2ECqiQ/bAD3hy4KGoU1y3uT5gcC1pOXTE1XP4d1LDt434G4nPUIMxkMLhotecviWsbJ4FocRcPXs8qVptgarj5h9IMrjyLXDm17hfUqTtImhMEDqE/JFAoGBAP5RWI9qhIzEAWxualhAc2IssQ8MPFGkXFWB4NWYmXVMxodVonyQp9P4AFdL3wsiVhvSmccxqCy5R6Mbv81rFnb29meL08eqAy3HAW/br3jcnbN/W799OGaMXBi+DDES1xBmQndFVGFfm09xcndyTrmGiMgHW1kGgz6WFEHuRXI/AoGBAMSXMMqHEZxD9T4tahFV4xlUP67mLDrETQNOT0vX4NrLS4CXZMkt+IVqy3Z0TRapyIrkBpEUoH3ScGzzbIOdLz57SS7D+ZbsU5kCTDXKfPVxww4RjZ92xNJiEHBMzta9Ku5+D3mBznBFm9dMO8+E+0PSKMxm8n4AovQ9SU/hGyTXAoGBALrU4+yoYixPqoQQMcwXvSx4jLLzWDTaPIMM4THJ46MK/iZaQP6l/sV4Qjffo0I4vW2/L/3oexYwH3KyZhvw+hX3pFm5naHnQmKU+ndEuwpdePVvMOXihla/8sCyjZ5Xqut/VIDuy+ilJiIcw+0Aatlc/ouE7BTg9fY6pzMwapBdAoGBAIkMvb7zGpvN5JJMJr2fGor16M+NNxhg8S900GMXRHJDd4dWA7UcjzyzjtQtj/BUvLHW9Zz+vEP7CNVrfiLi2aS9Xe90P/OvHTh2GZsGZsbVYB3WrtyUd/IS21LuuOOLTPqmdzNGAxzR6irVwnyRQHmvcTHOMw8UcoXCk/FUBRBRAoGAYx52azbgbw1dDoxgGWHWS31zpYdGdg2pxPkr1TInrKFbrQEwnNTqcnPIHqytkvrc9gpqW4XU0Ux7pU/twPO9JXHfXDNWXmJUHyYIeLwbeHrhJ+avtPDs8VwsI8zbBu84D13NJq7dWxPLEUYQNrCrAw4+ywvXmBzXkvEdLaQq1S4="
)

func TestMain(m *testing.M) {
	xlog.Level = xlog.DebugLevel
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

	// 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用，设置后会报错）
	//client.SetAESKey("KvKUTqSVZX2fUgmxnFyMaQ==")

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
