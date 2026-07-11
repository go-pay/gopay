// Command generic 演示如何调用 SDK 未提供类型化方法的接口。
//
// SDK 已为常用接口封装了强类型方法，其余 30+ 接口（数字人民币、微信支付分、
// 支付宝 APP/WAP、银联云闪付等）可用通用的 Execute / ExecuteEncrypted 调用：
// 传入接口路径常量、任意可 JSON 序列化的请求体，以及接收响应的目标。
//
//	go run ./example/generic
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-pay/gopay/cmbpay"
	"github.com/go-pay/gopay/examples/cmbpay/config"
)

func main() {
	client := config.MustClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 方式一：请求体用 map，响应解析到 map。
	// 以银联云闪付（4.16）为例，业务字段按接口文档填写。
	req := map[string]string{
		"merId":     client.MerID(),
		"orderId":   fmt.Sprintf("G%d", time.Now().UnixNano()),
		"userId":    "N003109945",
		"notifyUrl": config.NotifyURL(),
		"txnAmt":    "1",
		"authCode":  "6220000000000000000",
	}
	var resp map[string]any
	if err := client.Execute(ctx, cmbpay.PathCloudPay, req, &resp); err != nil {
		log.Printf("银联云闪付调用返回: %v", err)
	} else {
		pretty, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Printf("响应业务体:\n%s\n", pretty)
	}

	// 方式二：请求体用自定义结构体，响应解析到自定义结构体。
	type ecnyOrderReq struct {
		MerID     string `json:"merId"`
		OrderID   string `json:"orderId"`
		UserID    string `json:"userId"`
		NotifyURL string `json:"notifyUrl"`
		TxnAmt    string `json:"txnAmt"`
	}
	type ecnyOrderResp struct {
		MerID      string `json:"merId"`
		OrderID    string `json:"orderId"`
		CmbOrderID string `json:"cmbOrderId"`
		QrCode     string `json:"qrCode"`
	}
	var eResp ecnyOrderResp
	err := client.Execute(ctx, cmbpay.PathEcnyUnifiedOrder, &ecnyOrderReq{
		MerID:     client.MerID(),
		OrderID:   fmt.Sprintf("E%d", time.Now().UnixNano()),
		UserID:    "N003109945",
		NotifyURL: config.NotifyURL(),
		TxnAmt:    "1",
	}, &eResp)
	if err != nil {
		log.Printf("数字人民币统一下单调用返回: %v", err)
	} else {
		fmt.Printf("数币下单 平台单号=%s 二维码=%s\n", eResp.CmbOrderID, eResp.QrCode)
	}

	// 含敏感字段时用 ExecuteEncrypted：先加密字段，再随请求发送数字信封。
	enc, _ := cmbpay.NewEncryptor()
	termCipher, _ := enc.Encrypt(`{"location":"+37.12/-121.213"}`)
	_ = termCipher // 将其填入请求体对应的加密字段后调用：
	// client.ExecuteEncrypted(ctx, cmbpay.PathXxx, reqWithCipher, enc, &out)
}
