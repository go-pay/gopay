// Package config 为各示例提供统一的客户端构建方式：优先读取环境变量，
// 未设置时回退到接口文档附录 1 的联调（UAT）示例参数，方便直接 go run 体验。
//
// 支持的环境变量：
//
//	CMB_ENV        prod 或 uat（默认 uat）
//	CMB_APPID      APP ID
//	CMB_APPSECRET  APP SECRET
//	CMB_MERID      商户号
//	CMB_PRIVKEY    商户 SM2 私钥（64 字符 HEX）
//	CMB_CMBPUBKEY  招行 SM2 公钥（Base64）
//	CMB_NOTIFYURL  异步通知回调地址
package config

import (
	"log"
	"os"

	"github.com/go-pay/gopay/cmbpay"
)

// 接口文档附录 1 提供的联调环境示例参数（仅用于本地体验，非真实生产凭证）。
const (
	// 应用ID
	demoAppID = "Your cmb APPID"
	// 应用密钥
	demoAppSecret = "Your cmb APPSecret"
	// 招行商户号
	demoMerID = "Your merchantId"
	// SM2私钥
	demoPrivKey = "Your SM2 Private Key with hex"
	// 招行公钥
	demoCMBPubKey = "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE6Q+fktsnY9OFP+LpSR5Udbxf5zHCFO0PmOKlFNTxDIGl8jsPbbB/9ET23NV+acSz4FEkzD74sW2iiNVHRLiKHg=="
)

// NotifyURL 返回回调地址，未设置环境变量时给出占位地址。
func NotifyURL() string {
	return env("CMB_NOTIFYURL", "https://your.domain/cmb/notify")
}

// MustClient 构建并返回一个 *cmbpay.Client，失败即退出。
func MustClient() *cmbpay.Client {
	host, billCheckHost := cmbpay.HostUAT, cmbpay.BillCheckHostUAT
	if env("CMB_ENV", "uat") == "prod" {
		host, billCheckHost = cmbpay.HostProd, cmbpay.BillCheckHostPRD
	}
	client, err := cmbpay.NewClient(cmbpay.Config{
		Host:               host,
		BillCheckHost:      billCheckHost,
		AppID:              env("CMB_APPID", demoAppID),
		AppSecret:          env("CMB_APPSECRET", demoAppSecret),
		MerID:              env("CMB_MERID", demoMerID),
		PrivateKeyHex:      env("CMB_PRIVKEY", demoPrivKey),
		CMBPublicKeyBase64: env("CMB_CMBPUBKEY", demoCMBPubKey),
	})
	if err != nil {
		log.Fatalf("初始化客户端失败: %v", err)
	}
	return client
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
