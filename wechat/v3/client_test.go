package wechat

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/cedarwu/gopay"
	"github.com/cedarwu/gopay/pkg/util"
	"github.com/cedarwu/gopay/pkg/xlog"
)

var (
	client            *ClientV3
	err               error
	MchId             = ""
	APIv3Key          = ""
	SerialNo          = ""
	PrivateKeyContent = `-----BEGIN PRIVATE KEY-----
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
	WxPublicKeySerialNo = "60A862B18FE9F86BF7075383F09C8092704A2B4D"
	WxPublicKeyContent  = `-----BEGIN CERTIFICATE-----
MIIDVzCCAj+gAwIBAgIJANfOWdH1ItcBMA0GCSqGSIb3DQEBCwUAMEIxCzAJBgNV
BAYTAlhYMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0RlZmF1bHQg
Q29tcGFueSBMdGQwHhcNMjEwNDI3MDg1NTIzWhcNMzEwNDI1MDg1NTIzWjBCMQsw
CQYDVQQGEwJYWDEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZh
dWx0IENvbXBhbnkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
2VCTd91fnUn73Xy9DLvt/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXag
HQ4vGeZN0yqm++rDsGK+U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOm
DeJf2vg6byms9RUipiq4SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B83
4uxl1zNLYQCrkdBzb8oUxwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGf
y8wus7hwKz6clt3Whmmda7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtk
aSef4FdEEgEXJiw0VAJT8wIDAQABo1AwTjAdBgNVHQ4EFgQUT1c7nd/SUO76HSoZ
umNUJv1R5PwwHwYDVR0jBBgwFoAUT1c7nd/SUO76HSoZumNUJv1R5PwwDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAfTjxKRQMzNB/U6ZoCUS+BSNfa2Oh
0plMN6ZuzwiVVZwg1jywvv5yv04koS7Pd4i9E4gt9ZBUQXlpq+A3oOCEEHNRR6b2
kyazGRM7s0OP5X21WrbpSmKmU6K7hkfx30yYs08LVs/Q8DIhvaj1FCFeJzUCzYn/
fHMq4tsbKO0dKAeydPM/nrUZBmaYQVKMVOORGLFjFKVO7JV6Kq/R86ouhjEPgJOe
2xulNBUcjicqtZlBdEh/PWCYP2SpGVDclKm8jeo175T3EVAkdKzzmfpxtMmnMlmq
cTJOU9TxuGvNASMtjj7pYIerTx+xgZDXEVBWFW9PjJ0TV06tCRsgSHItgg==
-----END CERTIFICATE-----`
)

func TestMain(m *testing.M) {
	// NewClientV3 初始化微信客户端 V3
	//	mchid：商户ID
	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：APIv3Key，商户平台获取
	//	privateKey：私钥 apiclient_key.pem 读取后的字符串内容
	client, err = NewClientV3(MchId, SerialNo, APIv3Key, PrivateKeyContent)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 设置微信平台证书和序列号，并启用自动同步返回验签
	//	注意：请预先通过 wechat.GetPlatformCerts() 获取并维护微信平台公钥证书和证书序列号
	client.SetPlatformCert([]byte(WxPublicKeyContent), WxPublicKeySerialNo).AutoVerifySign()

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	os.Exit(m.Run())
}

func TestGetPlatformCertsWithoutClient(t *testing.T) {
	certs, err := GetPlatformCerts(MchId, APIv3Key, SerialNo, PrivateKeyContent)
	if err != nil {
		xlog.Error(err)
		return
	}
	if certs.Code == Success {
		for _, v := range certs.Certs {
			xlog.Infof("生效时间: %s", v.EffectiveTime)
			xlog.Infof("到期时间: %s", v.ExpireTime)
			xlog.Infof("WxSerialNo: %s", v.SerialNo)
			xlog.Infof("WxContent: \n%s", v.PublicKey)
		}
		return
	}
	xlog.Errorf("certs:%s", certs.Error)
}

func TestGetPlatformCerts(t *testing.T) {
	certs, err := client.GetPlatformCerts()
	if err != nil {
		xlog.Error(err)
		return
	}
	if certs.Code == Success {
		for _, v := range certs.Certs {
			xlog.Infof("生效时间: %s", v.EffectiveTime)
			xlog.Infof("到期时间: %s", v.ExpireTime)
			xlog.Infof("WxSerialNo: %s", v.SerialNo)
			xlog.Infof("WxContent: \n%s", v.PublicKey)
		}
		return
	}
	xlog.Errorf("certs:%s", certs.Error)
}

func TestV3VerifySign(t *testing.T) {
	// type SignInfo struct {
	//	HeaderTimestamp string `json:"Wechatpay-Timestamp"`
	//	HeaderNonce     string `json:"Wechatpay-Nonce"`
	//	HeaderSignature string `json:"Wechatpay-Signature"`
	//	HeaderSerial    string `json:"Wechatpay-Serial"`
	//	SignBody        string `json:"sign_body"`
	//}
	timestamp := "1609149813"
	nonce := "c4682f0902e4c7fd5cfb7568a2a45e1b"
	signBody := `{"code_url":"weixin://wxpay/bizpayurl?pr=5zPMHa4zz"}`
	signature := "D/nRx+h1To/ybCJkJYTXptoSp6+UVPsKNlJ2AsHMf76rXq2qAYDSnoOTB4HRc8ZlPNck5JfeZ19lDXAJ/N9gyvWEwE3n01HNhaKqxOjW0C1KROCtxAj1Wd2qtMyiCzh/Azuk15eIHjht03teGQFDmowoOBSlMg9qOBaK8MNfwFcXvV3J12AMbFFR7s4cXbqzuk2qBeMAz6VrKDAwDHxZOWFqME59mg4bPWwBTNyYeCQVR2sqPflLvY1zttEGMN3s/CDvgLQ/SXZrAsHlS2lkDVHEc/sP9q0x9oU8lFL6DhD6eDU2mVP3pt7CPD/5QAnGnINaHIcZVj6Vb4l3PKzeog=="

	err = V3VerifySign(timestamp, nonce, signBody, signature, WxPublicKeyContent)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("sign ok")
}

func TestV3Jsapi(t *testing.T) {
	tradeNo := util.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("description", "测试Jsapi支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", "https://www.fmm.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", "asdas")
		})
	text, err := client.V3EncryptText("张三")
	if err != nil {
		xlog.Errorf("client.V3EncryptText(),err:%+v", err)
		err = nil
	}
	xlog.Debugf("加密text: %s", text)

	wxRsp, err := client.V3TransactionJsapi(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3Native(t *testing.T) {
	tradeNo := util.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("appid", "wx52xxxxxxxxxxx").
		Set("description", "测试Native支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", "https://www.fmm.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3TransactionNative(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3PartnerNative(t *testing.T) {
	tradeNo := util.GetRandomString(32)
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("description", "测试Native支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("sub_mchid", "1900000109").
		//Set("notify_url", "https://api2.fangyiyun.com/api/v1/wechat/callback").
		Set("notify_url", "https://www.fmm.ink").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3PartnerTransactionNative(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3QueryOrder(t *testing.T) {
	//wxRsp, err := client.V3TransactionQueryOrder(TransactionId, "42000008462020122402449153433")
	wxRsp, err := client.V3TransactionQueryOrder(OutTradeNo, "22LW55HDd8tuxgZgFM445kI52BZVk847")
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3CloseOrder(t *testing.T) {
	wxRsp, err := client.V3TransactionCloseOrder("FY160932049419637602")
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Error("success")
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3BillTradeBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2020-12-30").
		Set("tar_type", "GZIP")

	wxRsp, err := client.V3BillTradeBill(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3BillFundFlowBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2020-12-30").
		Set("tar_type", "GZIP")

	wxRsp, err := client.V3BillFundFlowBill(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestV3BillDownLoadBill(t *testing.T) {
	url := "https://api.mch.weixin.qq.com/v3/billdownload/file?token=4MWpG4bWfL3smAe2AeB8scfp1MN0LYORxW691-jI-wL9J9fA6F0qG0q66y44xrur&tartype=gzip"
	fileBytes, err := client.V3BillDownLoadBill(url)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("fileBytes:%v", fileBytes)

	// 申请账单时采用 GZIP 压缩，返回 bytes 为压缩文件
	//err = ioutil.WriteFile("bill.zip", fileBytes, 0666)
	//if err != nil {
	//	xlog.Error("ioutil.WriteFile:", err)
	//	return
	//}
}

func TestV3ProfitSharingOrder(t *testing.T) {
	client.autoSign = true
	var rs []*ProfitSharingReceiver
	item := &ProfitSharingReceiver{
		Type:        "PERSONAL_OPENID",
		Account:     "oOv-Z573Ktz7o2WRkzX98eAxePVE",
		Amount:      10,
		Description: "提现实时到账",
	}
	rs = append(rs, item)
	// bs, _ := json.Marshal(rs)

	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", "4200001149202106084654939138").
		Set("out_order_no", "202106071738581340").
		Set("unfreeze_unsplit", false).Set("receivers", rs)

	wxRsp, err := client.V3ProfitShareOrder(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp: %#v", wxRsp)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestV3ProfitSharingAddReceiver(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("type", "PERSONAL_OPENID").
		Set("account", "oOv-Z573Ktz7o2WRkzX98eAxePVE").
		Set("relation_type", "USER")

	wxRsp, err := client.V3ProfitShareAddReceiver(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp: %#v", wxRsp)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestV3ProfitSharingDeleteReceiver(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("type", "PERSONAL_OPENID").
		Set("account", "oOv-Z573Ktz7o2WRkzX98eAxePVE")

	wxRsp, err := client.V3ProfitShareDeleteReceiver(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp: %#v", wxRsp)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestV3ProfitSharingQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", "4200001149202106084654939138")
	wxRsp, err := client.V3ProfitShareOrderQuery("P20150806125346", bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp: %#v", wxRsp)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestV3ProfitSharingUnfreeze(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", "202106071738581338")
	bm.Set("out_order_no", "4200001037202106072686278117")
	bm.Set("description", "账单解冻")
	wxRsp, err := client.V3ProfitShareOrderUnfreeze(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp: %#v", wxRsp)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestV3ProfitSharingUnsplitQuery(t *testing.T) {
	wxRsp, err := client.V3ProfitShareUnsplitAmount("4200001149202106084654939138")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp: %#v", wxRsp)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestClientV3_V3MediaUploadImage(t *testing.T) {
	fileName := "logo.png"
	fileContent, err := ioutil.ReadFile("../../logo.png")
	if err != nil {
		xlog.Error(err)
		return
	}
	h := sha256.New()
	h.Write(fileContent)
	sha256Str := hex.EncodeToString(h.Sum(nil))
	xlog.Debug("sha256：", sha256Str)

	img := &util.File{
		Name:    fileName,
		Content: fileContent,
	}

	wxRsp, err := client.V3MediaUploadImage(fileName, sha256Str, img)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func TestClientV3_V3ComplaintUploadImage(t *testing.T) {
	fileName := "logo.png"
	fileContent, err := ioutil.ReadFile("../../logo.png")
	if err != nil {
		xlog.Error(err)
		return
	}
	h := sha256.New()
	h.Write(fileContent)
	sha256Str := hex.EncodeToString(h.Sum(nil))
	xlog.Debug("sha256：", sha256Str)

	img := &util.File{
		Name:    fileName,
		Content: fileContent,
	}

	wxRsp, err := client.V3ComplaintUploadImage(fileName, sha256Str, img)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp: %s", wxRsp.Error)
}
