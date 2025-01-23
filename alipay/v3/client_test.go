package alipay

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *ClientV3
	err    error
)

func TestMain(m *testing.M) {
	xlog.SetLevel(xlog.DebugLevel)
	// 初始化支付宝客V3户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err = NewClientV3(cert.Appid, cert.PrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// Debug开关，输出/关闭日志
	client.DebugSwitch = gopay.DebugOn

	// 设置自定义RequestId生成方法
	//client.SetRequestIdFunc()

	// 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用）
	//client.SetAESKey("KvKUTqSVZX2fUgmxnFyMaQ==")

	// 传入证书内容
	err = client.SetCert(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)
	if err != nil {
		xlog.Debug("SetCert:", err)
		return
	}
	os.Exit(m.Run())
}

func TestDoAliPayAPISelfV3(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "预创建创建订单").
		Set("out_trade_no", util.RandomString(32)).
		Set("total_amount", "0.01")

	rsp := new(struct {
		OutTradeNo string `json:"out_trade_no"`
		QrCode     string `json:"qr_code"`
	})
	// 创建订单
	res, err := client.DoAliPayAPISelfV3(ctx, MethodPost, v3TradePrecreate, bm, rsp)
	if err != nil {
		xlog.Errorf("client.TradePrecreate(), err:%v", err)
		return
	}
	xlog.Debugf("aliRsp:%s", js.MarshalString(rsp))
	if res.StatusCode != Success {
		xlog.Errorf("aliRsp.StatusCode:%d", res.StatusCode)
		return
	}
}

func TestClientV3_Transfer(t *testing.T) {
	bm := make(gopay.BodyMap)
	// 收款方信息
	type PayeeInfo struct {
		Identity     string `json:"identity"`      // 必选
		IdentityType string `json:"identity_type"` // 必选
		CertNo       string `json:"cert_no"`       // 可选
		CertType     string `json:"cert_type"`     // 可选
		Name         string `json:"name"`          // 可选
	}
	payeeInfo := &PayeeInfo{
		Identity:     "sjrngj9819@sandbox.com",
		IdentityType: "ALIPAY_LOGON_ID",
		Name:         "sjrngj9819",
	}

	bm.Set("out_biz_no", util.RandomString(32)).
		Set("trans_amount", "0.01").
		Set("biz_scene", "DIRECT_TRANSFER").
		Set("product_code", "TRANS_ACCOUNT_NO_PWD").
		Set("order_title", "转账测试").
		Set("payee_info", payeeInfo).
		Set("remark", "转账测试").
		Set("business_params", struct{}{})

	res, err := client.FundTransUniTransfer(ctx, bm)
	if err != nil {
		xlog.Errorf("client.FundTransUniTransfer(), err:%v", err)
		return
	}

	xlog.Debugf("aliRsp:%s", js.MarshalString(res))
	if res.StatusCode != Success {
		xlog.Errorf("aliRsp.StatusCode:%d", res.StatusCode)
		return
	}
}
