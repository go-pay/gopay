package wecaht

import (
	"os"
	"testing"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	client         *ClientV3
	err            error
	appid          = "wx52a25f196830f677"
	mchid          = "1604896569"
	serialNo       = "298A5BA7E00AF6E71579E81D9CB1AC7037A51471"
	certKeyContent = `-----BEGIN PRIVATE KEY-----
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
)

func TestMain(m *testing.M) {
	// NewClientV3 初始化微信客户端 V3
	//	appid：appid
	//	mchid：商户ID
	// 	serialNo 商户证书的证书序列号
	//	pkContent：私钥 apiclient_key.pem 读取后的内容
	client, err = NewClientV3(appid, mchid, serialNo, []byte(certKeyContent))
	if err != nil {
		xlog.Error(err)
		os.Exit(1)
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	os.Exit(m.Run())
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
	xlog.Debug("wxRsp:", wxRsp)
}

func TestV3QueryOrder(t *testing.T) {
	//wxRsp, err := client.V3TransactionQueryOrder(TransactionId, "42000008462020122402449153433")
	wxRsp, err := client.V3TransactionQueryOrder(OutTradeNo, "22LW55HDd8tuxgZgFM445kI52BZVk848")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp:%#v", wxRsp)
}
