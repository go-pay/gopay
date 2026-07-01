package douyin

import (
	"context"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/douyin"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

const (
	Mchid          = "600000000000001"
	SerialNo       = "商户 API 证书序列号"
	ApiKey       = "抖音支付接口加密密钥32字节"
	PrivateKey     = "商户 API 证书私钥 apiclient_key.pem 内容"
	PlatformKey    = "抖音支付平台证书 PEM 内容"
	PlatformSerial = "抖音支付平台证书序列号"

	Appid     = "awz9w2wncdof4ba6"
	NotifyUrl = "https://www.example.com/douyin/notify"
)

// 初始化抖音支付客户端
func newClient() (*douyin.Client, error) {
	client, err := douyin.NewClient(Mchid, SerialNo, ApiKey, PrivateKey)
	if err != nil {
		return nil, err
	}
	client.DebugSwitch = gopay.DebugOn
	if err = client.SetPlatformCert([]byte(PlatformKey), PlatformSerial); err != nil {
		return nil, err
	}
	return client, nil
}

// AppOrder App 支付下单 + 生成 App 端调起签名
func AppOrder() {
	client, err := newClient()
	if err != nil {
		xlog.Error("newClient err:", err)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音支付测试商品").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		Set("time_expire", time.Now().Add(15*time.Minute).Format(time.RFC3339)).
		SetBodyMap("amount", func(sub gopay.BodyMap) {
			sub.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.AppOrder(context.Background(), bm)
	if err != nil {
		xlog.Error("AppOrder err:", err)
		return
	}
	if rsp.Code != douyin.Success {
		xlog.Errorf("AppOrder fail, code=%d, err_response=%+v", rsp.Code, rsp.ErrResponse)
		return
	}
	xlog.Infof("prepay_id=%s", rsp.Response.PrepayId)

	// 生成 App 端调起签名
	params, err := client.PaySignOfApp(Appid, rsp.Response.PrepayId)
	if err != nil {
		xlog.Error("PaySignOfApp err:", err)
		return
	}
	xlog.Infof("app pay params: %+v", params)
}

// JsapiOrder JSAPI 支付下单 + 前端调起签名
func JsapiOrder() {
	client, err := newClient()
	if err != nil {
		xlog.Error("newClient err:", err)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音JSAPI支付").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(sub gopay.BodyMap) {
			sub.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.JsapiOrder(context.Background(), bm)
	if err != nil {
		xlog.Error("JsapiOrder err:", err)
		return
	}
	if rsp.Code != douyin.Success {
		xlog.Errorf("JsapiOrder fail: %+v", rsp.ErrResponse)
		return
	}
	xlog.Infof("prepay_id=%s", rsp.Response.PrepayId)

	params, err := client.PaySignOfJSAPI(Appid, rsp.Response.PrepayId)
	if err != nil {
		xlog.Error("PaySignOfJSAPI err:", err)
		return
	}
	xlog.Infof("jsapi params: %+v", params)
}

// H5Order H5 支付下单（响应直接返回可跳转的 h5_url）
func H5Order() {
	client, err := newClient()
	if err != nil {
		xlog.Error("newClient err:", err)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音H5支付").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(sub gopay.BodyMap) {
			sub.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.H5Order(context.Background(), bm)
	if err != nil {
		xlog.Error("H5Order err:", err)
		return
	}
	if rsp.Code != douyin.Success {
		xlog.Errorf("H5Order fail: %+v", rsp.ErrResponse)
		return
	}
	xlog.Infof("h5_url=%s", rsp.Response.H5Url)
}

// NativeOrder Native 支付下单（响应直接返回二维码 code_url）
func NativeOrder() {
	client, err := newClient()
	if err != nil {
		xlog.Error("newClient err:", err)
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("description", "抖音Native支付").
		Set("out_trade_no", "OUT_"+util.RandomString(16)).
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(sub gopay.BodyMap) {
			sub.Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.NativeOrder(context.Background(), bm)
	if err != nil {
		xlog.Error("NativeOrder err:", err)
		return
	}
	if rsp.Code != douyin.Success {
		xlog.Errorf("NativeOrder fail: %+v", rsp.ErrResponse)
		return
	}
	xlog.Infof("code_url=%s", rsp.Response.CodeUrl)
}

// QueryOrder 查询订单（两种方式）
func QueryOrder() {
	client, _ := newClient()

	rsp, err := client.OrderQueryByOutTradeNo(context.Background(), "OUT_1666688488")
	if err != nil {
		xlog.Error("OrderQueryByOutTradeNo err:", err)
		return
	}
	xlog.Infof("query by out_trade_no => trade_state=%s", rsp.Response.TradeState)

	rsp2, err := client.OrderQueryByTransactionId(context.Background(), "TP2022101317144741443210681000")
	if err != nil {
		xlog.Error("OrderQueryByTransactionId err:", err)
		return
	}
	xlog.Infof("query by transaction_id => trade_state=%s", rsp2.Response.TradeState)
}

// CloseOrder 关闭订单
func CloseOrder() {
	client, _ := newClient()

	rsp, err := client.CloseOrder(context.Background(), "OUT_1666688488", nil)
	if err != nil {
		xlog.Error("CloseOrder err:", err)
		return
	}
	if rsp.Code != douyin.Success {
		xlog.Errorf("CloseOrder fail: %+v", rsp.ErrResponse)
		return
	}
	xlog.Info("CloseOrder success")
}

// Refund 申请退款
func Refund() {
	client, _ := newClient()

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("out_trade_no", "OUT_1666688488").
		Set("out_refund_no", "REF_"+util.RandomString(16)).
		Set("reason", "用户申请退款").
		Set("notify_url", NotifyUrl).
		SetBodyMap("amount", func(sub gopay.BodyMap) {
			sub.Set("refund", 1).Set("total", 1).Set("currency", "CNY")
		})

	rsp, err := client.Refund(context.Background(), bm)
	if err != nil {
		xlog.Error("Refund err:", err)
		return
	}
	if rsp.Code != douyin.Success {
		xlog.Errorf("Refund fail: %+v", rsp.ErrResponse)
		return
	}
	xlog.Infof("refund_id=%s status=%s", rsp.Response.RefundId, rsp.Response.Status)
}

// RefundQuery 查询退款
func RefundQuery() {
	client, _ := newClient()

	rsp, err := client.RefundQuery(context.Background(), "REF_XXX", "", "")
	if err != nil {
		xlog.Error("RefundQuery err:", err)
		return
	}
	xlog.Infof("refund_status=%s", rsp.Response.Status)
}

// Profit 分账主流程（请求 → 查询 → 完结）
func Profit() {
	client, _ := newClient()

	// 加密接收方名称（RSA-PKCS1v15，使用平台证书公钥）
	encryptedName, err := client.EncryptText("接收方商户名称")
	if err != nil {
		xlog.Error("EncryptText err:", err)
		return
	}

	// 1) 请求分账（异步受理）
	reqBM := make(gopay.BodyMap)
	reqBM.Set("appid", Appid).
		Set("transaction_id", "TP2022101317144741443210681000").
		Set("out_order_no", "SPLIT_"+util.RandomString(16)).
		Set("unfreeze_unsplit", false).
		Set("notify_url", NotifyUrl).
		Set("receivers", []douyin.ProfitReceiverReq{
			{Type: "MERCHANT_ID", Account: "6020230307605084", Name: encryptedName, Amount: 100, Description: "分给合作方"},
		})
	reqRsp, err := client.ProfitRequest(context.Background(), reqBM)
	if err != nil {
		xlog.Error("ProfitRequest err:", err)
		return
	}
	if reqRsp.Code != douyin.Success {
		xlog.Errorf("ProfitRequest fail: %+v", reqRsp.ErrResponse)
		return
	}
	xlog.Infof("profit request => order_id=%s state=%s", reqRsp.Response.OrderId, reqRsp.Response.State)

	// 2) 查询分账
	queryBM := make(gopay.BodyMap)
	queryBM.Set("transaction_id", reqRsp.Response.TransactionId)
	queryRsp, err := client.ProfitQuery(context.Background(), reqRsp.Response.OutOrderNo, queryBM)
	if err != nil {
		xlog.Error("ProfitQuery err:", err)
		return
	}
	xlog.Infof("profit query => state=%s receivers=%d", queryRsp.Response.State, len(queryRsp.Response.Receivers))

	// 3) 完结分账（unfreeze_unsplit=false 时才需要）
	finBM := make(gopay.BodyMap)
	finBM.Set("transaction_id", reqRsp.Response.TransactionId).
		Set("out_order_no", "FIN_"+util.RandomString(16)).
		Set("description", "完结")
	finRsp, err := client.ProfitComplete(context.Background(), finBM)
	if err != nil {
		xlog.Error("ProfitComplete err:", err)
		return
	}
	xlog.Infof("profit complete => state=%s", finRsp.Response.State)
}

// Transfer 商户转账到用户零钱 + 查询
func Transfer() {
	client, _ := newClient()

	// 大额转账（≥2000元）时 user_name 需先加密
	encryptedUserName, err := client.EncryptText("张三")
	if err != nil {
		xlog.Error("EncryptText err:", err)
		return
	}

	bm := make(gopay.BodyMap)
	bm.Set("appid", Appid).
		Set("out_bill_no", "OUT_"+util.RandomString(16)).
		Set("transfer_scene_id", "SCENE_001").
		Set("openid", "oUpF8uMEB4jR").
		Set("user_name", encryptedUserName).
		Set("transfer_amount", 100).
		Set("transfer_remark", "商户转账").
		Set("notify_url", NotifyUrl)

	rsp, err := client.Transfer(context.Background(), bm)
	if err != nil {
		xlog.Error("Transfer err:", err)
		return
	}
	if rsp.Code != douyin.Success {
		xlog.Errorf("Transfer fail: %+v", rsp.ErrResponse)
		return
	}
	xlog.Infof("transfer => transfer_bill_no=%s state=%s", rsp.Response.TransferBillNo, rsp.Response.State)

	// 按抖音转账单号查询
	q1, err := client.TransferQueryByTransferBillNo(context.Background(), rsp.Response.TransferBillNo)
	if err != nil {
		xlog.Error("TransferQueryByTransferBillNo err:", err)
		return
	}
	xlog.Infof("query by transfer_bill_no => state=%s", q1.Response.State)

	// 按商户单号查询
	q2, err := client.TransferQueryByOutBillNo(context.Background(), rsp.Response.OutBillNo)
	if err != nil {
		xlog.Error("TransferQueryByOutBillNo err:", err)
		return
	}
	xlog.Infof("query by out_bill_no => state=%s", q2.Response.State)
}

// ApplyBill 申请三类账单并下载 + 校验 + 解压
func ApplyBill() {
	client, _ := newClient()

	// 1) 申请交易账单
	tradeBM := make(gopay.BodyMap)
	tradeBM.Set("bill_date", "2026-06-30")
	tradeRsp, err := client.ApplyTradeBill(context.Background(), tradeBM)
	if err != nil {
		xlog.Error("ApplyTradeBill err:", err)
		return
	}
	if tradeRsp.Code != douyin.Success {
		xlog.Errorf("ApplyTradeBill fail: %+v", tradeRsp.ErrResponse)
		return
	}
	xlog.Infof("trade bill: %+v", tradeRsp.Response)

	// 2) 申请资金账单（account_type 默认 BaseAccount，可显式指定 OperationAccount）
	fundBM := make(gopay.BodyMap)
	fundBM.Set("bill_date", "2026-06-30").Set("account_type", "BaseAccount")
	fundRsp, err := client.ApplyFundBill(context.Background(), fundBM)
	if err != nil {
		xlog.Error("ApplyFundBill err:", err)
		return
	}
	xlog.Infof("fund bill: %+v", fundRsp.Response)

	// 3) 申请分账账单
	profitBM := make(gopay.BodyMap)
	profitBM.Set("bill_date", "2026-06-30")
	profitRsp, err := client.ApplyProfitBill(context.Background(), profitBM)
	if err != nil {
		xlog.Error("ApplyProfitBill err:", err)
		return
	}
	xlog.Infof("profit bill: %+v", profitRsp.Response)

	// 4) 下载 + 解压 + 校验（示例以交易账单为例）
	gzipData, err := client.DownloadBillFile(context.Background(), tradeRsp.Response.DownloadUrl)
	if err != nil {
		xlog.Error("DownloadBillFile err:", err)
		return
	}
	raw, err := douyin.UngzipBill(gzipData)
	if err != nil {
		xlog.Error("UngzipBill err:", err)
		return
	}
	if err = douyin.VerifyBillHash(raw, tradeRsp.Response.HashType, tradeRsp.Response.HashValue); err != nil {
		xlog.Error("VerifyBillHash err:", err)
		return
	}
	xlog.Infof("bill csv bytes=%d", len(raw))
}
