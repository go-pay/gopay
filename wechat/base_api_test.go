package wechat

import (
	"strconv"
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

func TestClient_UnifiedOrder(t *testing.T) {
	number := util.RandomString(32)
	xlog.Info("out_trade_no:", number)
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("body", "H5支付").
		Set("out_trade_no", number).
		Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("notify_url", "https://www.fmm.ink").
		Set("trade_type", TradeType_H5).
		Set("device_info", "WEB").
		Set("sign_type", SignType_MD5).
		SetBodyMap("scene_info", func(bm gopay.BodyMap) {
			bm.SetBodyMap("h5_info", func(bm gopay.BodyMap) {
				bm.Set("type", "Wap")
				bm.Set("wap_url", "https://www.fmm.ink")
				bm.Set("wap_name", "H5测试支付")
			})
		}) /*.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")*/

	// 请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(ctx, bm)
	if err != nil {
		xlog.Errorf("client.UnifiedOrder(%+v),error:%+v", bm, err)
		return
	}
	xlog.Info("wxRsp:", *wxRsp)
	//xlog.Info("wxRsp.MwebUrl:", wxRsp.MwebUrl)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	// 获取小程序支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := GetMiniPaySign(appId, wxRsp.NonceStr, pac, SignType_MD5, timeStamp, apiKey)
	//xlog.Info("paySign:", paySign)

	// 获取H5支付需要的paySign
	pac := "prepay_id=" + wxRsp.PrepayId
	paySign := GetJsapiPaySign(appId, wxRsp.NonceStr, pac, SignType_MD5, timeStamp, apiKey)
	xlog.Debug("paySign:", paySign)

	// 获取小程序需要的paySign
	//paySign := GetAppPaySign(appId,"partnerid", wxRsp.NonceStr, wxRsp.PrepayId, SignType_MD5, timeStamp, apiKey)
	//xlog.Info("paySign:", paySign)
}

func TestClient_Micropay(t *testing.T) {
	number := util.RandomString(32)
	xlog.Info("out_trade_no:", number)
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("body", "扫用户付款码支付").
		Set("out_trade_no", number).
		Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("auth_code", "134622817080551492").
		Set("sign_type", SignType_MD5)

	// 请求支付，成功后得到结果
	wxRsp, err := client.Micropay(ctx, bm)
	if err != nil {
		xlog.Errorf("client.Micropay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("Response:", *wxRsp)
	ok, err := VerifySign(apiKey, SignType_MD5, wxRsp)
	if err != nil {
		xlog.Error(err)
	}
	xlog.Debug("同步验签结果：", ok) // 沙箱环境验签失败请用正式环境测
}

func TestClient_QueryOrder(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ").
		Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_MD5)

	// 请求订单查询，成功后得到结果
	wxRsp, resBm, err := client.QueryOrder(ctx, bm)
	if err != nil {
		xlog.Errorf("client.QueryOrder(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
	xlog.Debug("resBm：", resBm)
}

func TestClient_CloseOrder(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ").
		Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_MD5)

	// 请求关闭订单，成功后得到结果
	wxRsp, err := client.CloseOrder(ctx, bm)
	if err != nil {
		xlog.Errorf("client.CloseOrder(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
}

func TestClient_Refund(t *testing.T) {
	// 初始化参数结构体
	s := util.RandomString(64)
	xlog.Info("out_refund_no:", s)
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "QRcTBTbJLoDrWSW9FtpSFlgWhft2QbaY").
		Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_MD5).
		Set("out_refund_no", s).
		Set("total_fee", 101).
		Set("refund_fee", 101).
		Set("notify_url", "https://www.fmm.ink")

	// 请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	wxRsp, resBm, err := client.Refund(ctx, bm)
	if err != nil {
		xlog.Errorf("client.Refund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
	xlog.Debug("resBm：", resBm)
}

func TestClient_QueryRefund(t *testing.T) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk").
		Set("nonce_str", util.RandomString(32)).
		Set("sign_type", SignType_MD5) /*.
		Set("transaction_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk").
		Set("out_refund_no", "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7").
		Set("refund_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")*/

	// 请求申请退款
	wxRsp, resBm, err := client.QueryRefund(ctx, bm)
	if err != nil {
		xlog.Errorf("client.QueryRefund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("wxRsp：", *wxRsp)
	xlog.Debug("resBm：", resBm)
}

func TestClient_Reverse(t *testing.T) {
	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("out_trade_no", "6aDCor1nUcAihrV5JBlI09tLvXbUp02B").
		Set("sign_type", SignType_MD5)

	// 请求撤销订单，成功后得到结果，沙箱环境下，证书路径参数可传nil
	wxRsp, err := client.Reverse(ctx, bm)
	if err != nil {
		xlog.Errorf("client.Reverse(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("Response:", wxRsp)
}
