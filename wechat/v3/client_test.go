package wechat

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"testing"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	client    *ClientV3
	err       error
	Appid     = "wx52a25f196830f677"
	MchId     = "1604896569"
	ApiV3Key  = "j7XthyAeqmeKPNjECOEd60YVG1Knwr3Y"
	SerialNo  = "298A5BA7E00AF6E71579E81D9CB1AC7037A51471"
	PKContent = `-----BEGIN PRIVATE KEY-----
MIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDV523KVXZaaZI3
WxQiaz0J8o8nxAYsxBjrfcaKPnLo+r5uFME7GPV+4UHEZWG6ZogJ87yBt8L4IV8q
/2n0MPKV5qNtS0htG7G0Mtyw7lPmdXUXsA0ionbL2mzz0kgJ1S6FJwyZRRZNJ08Q
/GQE3TWqErbxL/2ITuzTeHrdTNL0i9oNxtB92EWFZ0gSL677zEiz41EVog24SyOd
TFqxjGFd9DR0CeRNU/oQPplFnM9YFseRuhEZ/jLATEvubH/U1qGqTlW0UHvIn14j
NqRxyAjDI/HfXl3Bo7Fx0QCJkVkqb+5ou8KFRchbcixRU0khbrxTy7dDJj60vSmr
PySqqZLFAgMBAAECggEBAKHPN9ZfX/B0/A6z7z86MCpeOryyJJmondFGi/H326Uy
SOus959k+hDJBZ8zsgH3neEpZ+gYwnxBgmRcYiI/BMMwfWAoGtmuoXbXIusU3pLv
N2x72PPiQktjKBgpciU+BrrjFzy6bmxe2AjZZC/pxrapAYrh6sA6NBykfwz5GHu0
DQmjHYqSlghDDljCzVR3Gcs/KicCMw6eQ0JlWDqtDEDoENlBkfn4spHwocV7HtSq
0bnUrQqqMtpZjbMJzZxJc39qkyNNDosuNy5GXYLQE7lr9RuRqLxEfg6KfGUS5bAZ
eJ5pizql7+c0viUtiXG17PYp8QR4c5G+54RlQd1pPuECgYEA9UBi5rFJzK0/n4aO
lsrp6BvUOSherp57SNYvpsRuBPU0odyH2/McLNphisKTxfSm0/hADaTmnzAnOUVg
cduc/5/5tVaaqyLL3SemxJhwqVsL3tE/KAN7HUBhhQrqD+H8r39TAoIkyfjCOHzS
74rygZ35x0kXNMavXQFB0RE2fEcCgYEA30dWaLddGmTvUXwhyTWcsiDfrsKbw8+n
MhAlSCXE42v9Uo3ULqD3/rpUQlMhoqaZb3cSyOyQwJvv0tp/g3hM7Q4usLxkdysc
KA9HmmZ4Q2P2838cqvNr/Dz1UAnfdDryMEnbiv1MUKYqFFTVX6oR0iH544JgDFCG
YLQA2M+3GpMCgYEAh+ax51v+pSirxN5vTSgMDc69/x5buS+g6W+m4CahQKYQEFGA
B2XkCwbIXngMIvm7KGK8O9NQ6I1qbtX+55jmmtAvM0lWU9boWRiL1Q0UAQSuwz34
XVfwdPkkEPFHWp3DxAwuF4m+kR0DowGocYzxbNn5e3EJJvmiW0tDCXMcWikCgYEA
tyNxWcUFBdBCh+i0YbCqzWSvdE3Fq8/YSPT7T3lDTHLYPu18W57Gq1Y0JI7BaQMT
mVzmuI1pkcKV7LIxoyl6l3ppi6eLFD/1AVq/FYL1I/mLpl/dqM6vBR8O686dTV3I
Jxl9jTyEayZQH4sR1TzPDze1GwpmM9Oc1RbwFuYRPycCgYEAzYaRKh6EQ+s37HDv
e/ZGMs3PI+CoA/x6lx4Owa7amRsWRKys45NV6gcC8pkbN4IeFaYXVHmJ1Yaef3xn
0VxHAzWI4BF+1pUwXzS2rAMBZR/VKS0XA856NauAC3mKHipoOWVVs+uFP3VMUQ79
hSImAa7UBzss6b6ie7AYxXtZBjY=
-----END PRIVATE KEY-----`

	WxPkContent = `-----BEGIN CERTIFICATE-----
MIID3DCCAsSgAwIBAgIUYKhisY/p+Gv3B1OD8JyAknBKK00wDQYJKoZIhvcNAQEL
BQAwXjELMAkGA1UEBhMCQ04xEzARBgNVBAoTClRlbnBheS5jb20xHTAbBgNVBAsT
FFRlbnBheS5jb20gQ0EgQ2VudGVyMRswGQYDVQQDExJUZW5wYXkuY29tIFJvb3Qg
Q0EwHhcNMjAxMjE3MDkzNjM3WhcNMjUxMjE2MDkzNjM3WjBuMRgwFgYDVQQDDA9U
ZW5wYXkuY29tIHNpZ24xEzARBgNVBAoMClRlbnBheS5jb20xHTAbBgNVBAsMFFRl
bnBheS5jb20gQ0EgQ2VudGVyMQswCQYDVQQGDAJDTjERMA8GA1UEBwwIU2hlblpo
ZW4wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC3kgThR8l6z60QgzJq
AnlES5wiQAdjQNig2beUJz5MMzia2+TsGsyK5FqhcLTPAZxluCJ6o8jhvYgg76NN
NQB+5HHNP+SYrCmYVpuq/0UNMM0tcWRV6OhBWkbQ13BY4aryC2seuky27aVvSP6z
jtioYOSiaMDlpx7KHWFNC2RTegoKV+Q+oxa8e0hq4t7tYeNss5/WVijCIkPg7Rgb
hzrqw/g7W6As1HZs37WOpPZ25DD55ztHqsNFMgo/jz79ob57yLhJAhGqzNaGC1mr
tvh0H3Aq8AIIuuEoNXYWAXzQg6MtfeBmVOklOC3jaZEZCVevGK1kNZCV3JURS1pd
8PklAgMBAAGjgYEwfzAJBgNVHRMEAjAAMAsGA1UdDwQEAwIE8DBlBgNVHR8EXjBc
MFqgWKBWhlRodHRwOi8vZXZjYS5pdHJ1cy5jb20uY24vcHVibGljL2l0cnVzY3Js
P0NBPTFCRDQyMjBFNTBEQkMwNEIwNkFEMzk3NTQ5ODQ2QzAxQzNFOEVCRDIwDQYJ
KoZIhvcNAQELBQADggEBAJum+aqHwh8usDfLp3tX/W2O+9WAXNfZucdeYTgAnhDh
0qjNN4pQCCHkiP/zUQGp0gbSsI/c+CDjHZ4zRnHV3leDystQZiIxeJ005pQz/SY4
mUOgeMFQC8DeGq0WUCtMYJCdKLz43XennMOSJzFYisp6c9vUZ/7CXl2qEbVfJ0Um
v4/yw6Y6o08eMk8jHAlTJCUsKefjS3OsIXWTTlQv4N6tvui7rWOjux2oQ37pJIT5
HrSIbzvplW2BjfPptspK+eQNSK+WAatSmfxU2vi8fS2BK1SeK/S1bvXqzcpohHtw
sUg2x/kdyA1Vas1TDLJHueVfNIZQF8sLFiAP/q33Jvo=
-----END CERTIFICATE-----`
)

func TestMain(m *testing.M) {
	// NewClientV3 初始化微信客户端 V3
	//	appid：appid
	//	mchid：商户ID
	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：apiV3Key，商户平台获取
	//	pkContent：私钥 apiclient_key.pem 读取后的内容
	client, err = NewClientV3(Appid, MchId, SerialNo, ApiV3Key, []byte(PKContent))
	if err != nil {
		xlog.Error(err)
		return
	}
	// 自动验签
	// 注意：未获取到微信平台公钥时，不要开启，请调用 client.GetPlatformCerts() 获取微信平台公钥
	//client.AutoVerifySign("微信平台公钥")

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	os.Exit(m.Run())
}

func TestGetPlatformCerts(t *testing.T) {
	certs, err := client.GetPlatformCerts()
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("certs.StatusCode", certs.StatusCode)
	xlog.Debug("certs.SignInfo", certs.SignInfo)

	for _, v := range certs.Certs {
		xlog.Debug("cert:", v)
	}
}

func TestV3VerifySign(t *testing.T) {
	//应答时间戳\n
	//应答随机串\n
	//应答报文主体\n

	ts := "1609149813"
	nonce := "c4682f0902e4c7fd5cfb7568a2a45e1b"
	signBody := `{"code_url":"weixin://wxpay/bizpayurl?pr=5zPMHa4zz"}`
	sign := "D/nRx+h1To/ybCJkJYTXptoSp6+UVPsKNlJ2AsHMf76rXq2qAYDSnoOTB4HRc8ZlPNck5JfeZ19lDXAJ/N9gyvWEwE3n01HNhaKqxOjW0C1KROCtxAj1Wd2qtMyiCzh/Azuk15eIHjht03teGQFDmowoOBSlMg9qOBaK8MNfwFcXvV3J12AMbFFR7s4cXbqzuk2qBeMAz6VrKDAwDHxZOWFqME59mg4bPWwBTNyYeCQVR2sqPflLvY1zttEGMN3s/CDvgLQ/SXZrAsHlS2lkDVHEc/sP9q0x9oU8lFL6DhD6eDU2mVP3pt7CPD/5QAnGnINaHIcZVj6Vb4l3PKzeog=="
	//serialNo := "60A862B18FE9F86BF7075383F09C8092704A2B4D"

	_str := ts + "\n" + nonce + "\n" + signBody + "\n"
	var (
		block     *pem.Block
		pubKey    *x509.Certificate
		publicKey *rsa.PublicKey
		ok        bool
	)

	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	if block, _ = pem.Decode([]byte(WxPkContent)); block == nil {
		xlog.Error("解析微信平台公钥失败")
		return
	}
	if pubKey, err = x509.ParseCertificate(block.Bytes); err != nil {
		xlog.Errorf("x509.ParseCertificate：%+v", err)
		return
	}
	if publicKey, ok = pubKey.PublicKey.(*rsa.PublicKey); !ok {
		xlog.Error("微信平台公钥转换错误")
		return
	}
	hashs := crypto.SHA256
	h := hashs.New()
	h.Write([]byte(_str))

	err = rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
	if err != nil {
		xlog.Debug(" sign error:", err)
		return
	}
	xlog.Debug("sign ok")
}

func TestV3Jsapi(t *testing.T) {
	tradeNo := gotil.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("description", "测试Jsapi支付商品").
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

func TestV3Native(t *testing.T) {
	tradeNo := gotil.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("description", "测试Native支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		//Set("notify_url", "https://api2.fangyiyun.com/api/v1/wechat/callback").
		Set("notify_url", "https://www.gopay.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3TransactionNative(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp.StatusCode:", wxRsp.StatusCode)
	xlog.Debugf("wxRsp.SignInfo:%#v", wxRsp.SignInfo)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestV3QueryOrder(t *testing.T) {
	//wxRsp, err := client.V3TransactionQueryOrder(TransactionId, "42000008462020122402449153433")
	wxRsp, err := client.V3TransactionQueryOrder(OutTradeNo, "22LW55HDd8tuxgZgFM445kI52BZVk847")
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.StatusCode == 200 {
		xlog.Debugf("wxRsp:%#v", wxRsp.Response)
	}
	xlog.Debugf("wxRsp:%s", wxRsp.Error)
}
