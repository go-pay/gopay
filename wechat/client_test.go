package wechat

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

var (
	client *Client
	appId  = "wxdaa2ab9ef87b5497"
	mchId  = "1368139502"
	apiKey = "GFDS8j98rewnmgl45wHTt980jg543wmg"
	ctx    = context.Background()
)

func TestMain(m *testing.M) {

	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client = NewClient(appId, mchId, apiKey, false)

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOn

	// 设置国家，不设置默认就是 China
	client.SetCountry(China)

	// 不使用证书接口，可以不添加证书
	// 添加pkcs12内容
	//err = client.AddCertPkcs12FilePath("apiclient_cert.p12")
	//if err != nil {
	//	panic(err)
	//}

	// 或

	// 添加pem证书路径
	//err := client.AddCertPemFilePath("apiclient_cert.pem", "apiclient_key.pem")
	//if err != nil {
	//	panic(err)
	//}

	os.Exit(m.Run())
}

func TestClient_AuthCodeToOpenId(t *testing.T) {
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("auth_code", "134753997737645794")

	wxRsp, err := client.AuthCodeToOpenId(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AuthCodeToOpenId(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("Response:", *wxRsp)
}

func TestClient_GetTransferInfo(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("partner_trade_no", util.RandomString(32))

	// 查询企业付款
	//    body：参数Body
	wxRsp, err := client.GetTransferInfo(ctx, bm)
	if err != nil {
		xlog.Errorf("client.GetTransferInfo(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
}

func TestClient_DownloadBill(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_MD5).
		Set("bill_date", "20190722").
		Set("bill_type", "ALL")

	// 请求下载对账单，成功后得到结果（string类型字符串）
	wxRsp, err := client.DownloadBill(ctx, bm)
	if err != nil {
		xlog.Errorf("client.DownloadBill(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_DownloadFundFlow(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_HMAC_SHA256).
		Set("bill_date", "20190122").
		Set("account_type", "Basic")

	// 请求下载资金账单，成功后得到结果，沙箱环境下，证书路径参数可传nil
	wxRsp, err := client.DownloadFundFlow(ctx, bm)
	if err != nil {
		xlog.Errorf("client.DownloadFundFlow(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}

func TestClient_BatchQueryComment(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_HMAC_SHA256).
		Set("begin_time", "20190120000000").
		Set("end_time", "20190122174000").
		Set("offset", "0")

	// 请求拉取订单评价数据，成功后得到结果，沙箱环境下，证书路径参数可传nil
	wxRsp, err := client.BatchQueryComment(ctx, bm)
	if err != nil {
		xlog.Errorf("client.BatchQueryComment(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}
