package alipay

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

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
	// alipay/v3 包内测试均依赖真实支付宝接口（集成测试），
	// 在 `go test -short ./...` 模式下整体跳过，避免拖慢整个仓库的测试。
	flag.Parse()
	if testing.Short() {
		os.Exit(0)
	}
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

	// 设置自定义配置（如需要）
	//client.
	//	SetAppAuthToken("xxx").    // 设置授权token
	//	SetBodySize().                        // 自定义配置http请求接收返回结果body大小，默认 10MB，没有特殊需求，可忽略此配置
	//	SetRequestIdFunc().                   // 设置自定义RequestId生成方法
	//	SetAESKey("KvKUTqSVZX2fUgmxnFyMaQ==") // 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用）

	// 如果需要单个请求独立设置 Alipay-App-Auth-Token，每个请求的 body map 中 bm.Set(alipay.HeaderAppAuthToken, "xxx")

	// Debug开关，输出/关闭日志
	client.DebugSwitch = gopay.DebugOn
	// 给 HTTP 客户端设置整体超时，避免支付宝某些接口偶发卡住导致 go test 整体超时
	client.GetHttpClient().SetTimeout(15 * time.Second)

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
