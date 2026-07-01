package douyin

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error

	// 商户凭证：请填写自己的抖音支付商户信息
	Mchid      = ""
	SerialNo   = ""
	ApiKey     = "" // 32 字节接口加密密钥，用于回调 AES-256-GCM 解密
	PrivateKey = ``

	// 抖音支付平台证书（可注册多张，请填写自己的平台证书 PEM 与序列号）
	PlatformCert     = ``
	PlatformCertSNo  = ""

	// 业务测试用参数：请按实际填写
	Appid     = ""
	NotifyUrl = "https://www.example.com/douyin/notify"
	Openid    = ""

	// 已存在的订单号（用于查询 / 关单 / 退款测试，请填写自己环境下的真实单号）
	TransactionId = ""
	OutTradeNo    = ""
	OutRefundNo   = ""
	OutOrderNo    = ""
	OutReturnNo   = ""
	OutBillNo     = ""
	TransferBillNo = ""
)

func TestMain(m *testing.M) {
	// NewClient 初始化抖音支付客户端
	//	mchid：商户号
	//	serialNo：商户 API 证书序列号
	//	apiKey：接口加密密钥（32 字节）
	//	privateKey：商户 API 私钥 apiclient_key.pem 内容
	client, err = NewClient(Mchid, SerialNo, ApiKey, PrivateKey)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 注册抖音支付平台证书（用于响应/回调验签）
	if PlatformCert != "" && PlatformCertSNo != "" {
		if err = client.SetPlatformCert([]byte(PlatformCert), PlatformCertSNo); err != nil {
			xlog.Error(err)
			return
		}
	}

	// 打开 Debug 日志开关，输出请求/响应报文
	client.DebugSwitch = gopay.DebugOff

	// 给 HTTP 客户端设置整体超时，避免测试偶发卡住
	client.GetHttpClient().SetTimeout(15 * time.Second)

	os.Exit(m.Run())
}

// ===================== 基础支付 =====================

func TestAppOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音App支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		Set("time_expire", time.Now().Add(15*time.Minute).Format(time.RFC3339)).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.AppOrder(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestJsapiOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音JSAPI支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(b gopay.BodyMap) {
			b.Set("openid", Openid)
		})

	rsp, err := client.JsapiOrder(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestH5Order(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音H5支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.H5Order(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("h5_url: %s", rsp.Response.H5Url)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestNativeOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音Native支付测试").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.NativeOrder(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("code_url: %s", rsp.Response.CodeUrl)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

// ===================== 订单查询与关单 =====================

func TestOrderQueryByTransactionId(t *testing.T) {
	rsp, err := client.OrderQueryByTransactionId(ctx, TransactionId)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestOrderQueryByOutTradeNo(t *testing.T) {
	rsp, err := client.OrderQueryByOutTradeNo(ctx, OutTradeNo)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestCloseOrder(t *testing.T) {
	rsp, err := client.CloseOrder(ctx, OutTradeNo, nil)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debug("close order success")
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

// ===================== 退款 =====================

func TestRefund(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("out_trade_no", OutTradeNo).
		Set("out_refund_no", "REF_"+util.RandomString(16)).
		Set("reason", "用户申请退款").
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("refund", 1).Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.Refund(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestRefundQuery(t *testing.T) {
	rsp, err := client.RefundQuery(ctx, OutRefundNo, "", Appid)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

// ===================== 账单 =====================

func TestApplyTradeBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2026-06-30")

	rsp, err := client.ApplyTradeBill(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestApplyFundBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2026-06-30").Set("account_type", "BaseAccount")

	rsp, err := client.ApplyFundBill(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestApplyProfitBill(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_date", "2026-06-30")

	rsp, err := client.ApplyProfitBill(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestDownloadBillFile(t *testing.T) {
	// 从 ApplyXxxBill 的返回中取 download_url，填入下面
	url := ""
	fileBytes, err := client.DownloadBillFile(ctx, url)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("fileBytes len:%d", len(fileBytes))

	// 若为 GZIP 压缩包，可解压
	//raw, err := UngzipBill(fileBytes)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Debugf("raw len:%d", len(raw))

	// 完整性校验（hashType 与 hashValue 来自 ApplyXxxBill 响应）
	//if err := VerifyBillHash(raw, "SHA1", "hash_value_from_apply_response"); err != nil {
	//	xlog.Error(err)
	//}
}

// ===================== 分账 =====================

func TestProfitRequest(t *testing.T) {
	// 若接收方需要 name（type=MERCHANT_ID 时必传），请先加密
	//encName, err := client.EncryptText("接收方商户名称")
	//if err != nil { xlog.Error(err); return }

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("transaction_id", TransactionId).
		Set("out_order_no", "SPLIT_"+util.RandomString(16)).
		Set("unfreeze_unsplit", false).
		Set("notify_url", NotifyUrl).
		Set("receivers", []ProfitReceiverReq{
			{Type: "MERCHANT_ID", Account: "6020230307605084", /* Name: encName, */ Amount: 100, Description: "分给合作方"},
		})

	rsp, err := client.ProfitRequest(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", TransactionId)

	rsp, err := client.ProfitQuery(ctx, OutOrderNo, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitRollback(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("order_id", "抖音支付分账单号").
		Set("out_order_no", OutOrderNo).
		Set("out_return_no", "OUT_"+util.RandomString(16)).
		Set("return_mchid", Mchid).
		Set("amount", 100).
		Set("description", "退分账")

	rsp, err := client.ProfitRollback(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitRollbackQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_order_no", OutOrderNo)

	rsp, err := client.ProfitRollbackQuery(ctx, OutReturnNo, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitComplete(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", TransactionId).
		Set("out_order_no", "FIN_"+util.RandomString(16)).
		Set("description", "完结分账").
		Set("notify_url", NotifyUrl)

	rsp, err := client.ProfitComplete(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitBalanceQuery(t *testing.T) {
	rsp, err := client.ProfitBalanceQuery(ctx, TransactionId, "")
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitReceiverAdd(t *testing.T) {
	// name 需先加密
	//encName, err := client.EncryptText("接收方名称")
	//if err != nil { xlog.Error(err); return }

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("type", "MERCHANT_ID").
		Set("account", "6020230307605084").
		//Set("name", encName).
		Set("relation_type", "STORE")

	rsp, err := client.ProfitReceiverAdd(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestProfitReceiverDelete(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("type", "MERCHANT_ID").
		Set("account", "6020230307605084")

	rsp, err := client.ProfitReceiverDelete(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

// ===================== 转账 =====================

func TestTransfer(t *testing.T) {
	// 大额转账（≥2000元）时 user_name 必填并先加密
	//encName, err := client.EncryptText("张三")
	//if err != nil { xlog.Error(err); return }

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("out_bill_no", "OUT_"+util.RandomString(16)).
		Set("transfer_scene_id", "SCENE_001").
		Set("openid", Openid).
		//Set("user_name", encName).
		Set("transfer_amount", 100).
		Set("transfer_remark", "商户转账").
		Set("notify_url", NotifyUrl)

	rsp, err := client.Transfer(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestTransferQueryByOutBillNo(t *testing.T) {
	rsp, err := client.TransferQueryByOutBillNo(ctx, OutBillNo)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}

func TestTransferQueryByTransferBillNo(t *testing.T) {
	rsp, err := client.TransferQueryByTransferBillNo(ctx, TransferBillNo)
	if err != nil {
		xlog.Error(err)
		return
	}
	if rsp.Code == Success {
		xlog.Debugf("rsp: %+v", rsp.Response)
		return
	}
	xlog.Errorf("rsp:%s", rsp.Error)
}
