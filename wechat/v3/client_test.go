package wechat

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
)

var (
	client      *ClientV3
	err         error
	Appid       = ""
	MchId       = ""
	ApiV3Key    = ""
	SerialNo    = ""
	PKContent   = ``
	WxPkContent = ``
)

func TestMain(m *testing.M) {
	// NewClientV3 初始化微信客户端 V3
	//	appid：appid
	//	mchid：商户ID
	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：apiV3Key，商户平台获取
	//	pkContent：私钥 apiclient_key.pem 读取后的字符串内容
	client, err = NewClientV3(Appid, MchId, SerialNo, ApiV3Key, PKContent)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 自动验签
	// 注意：未获取到微信平台公钥时，不要开启，请调用 client.GetPlatformCerts() 获取微信平台证书公钥
	//client.AutoVerifySign(WxPkContent)

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
	if certs.Code == Success {
		for _, v := range certs.Certs {
			xlog.Debug("cert:", v)
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

	err = V3VerifySign(timestamp, nonce, signBody, signature, WxPkContent)
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
		Set("notify_url", "https://www.fumm.cc").
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
	bm.Set("description", "测试Native支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		//Set("notify_url", "https://api2.fangyiyun.com/api/v1/wechat/callback").
		Set("notify_url", "https://www.fumm.cc").
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
		Set("notify_url", "https://www.fumm.cc").
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
	wxRsp, err := client.V3ProfitShareOrderQuery("202106071738581340", "4200001149202106084654939138")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("wxRsp: %#v", wxRsp)
	xlog.Debugf("wxRsp.Response:%#v", wxRsp.Response)
}

func TestV3ProfitSharingUnfreeze(t *testing.T) {
	wxRsp, err := client.V3ProfitShareOrderUnfreeze("202106071738581338", "4200001037202106072686278117", "账单解冻")
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
