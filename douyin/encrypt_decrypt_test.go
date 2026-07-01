package douyin

import (
	"testing"

	"github.com/go-pay/xlog"
)

var (
	// 抖音支付平台证书公钥（PEM 格式）
	// 用于包级 EncryptText 敏感字段加密
	platformPublicKey = ``

	// 商户 API 证书私钥（PEM 格式）
	// 用于包级 DecryptText 敏感字段解密
	merchantPrivateKey = ``
)

// TestClientEncryptDecryptText 使用 Client 上下文加解密
// 加密使用已注册的最新一张平台证书公钥；解密使用 NewClient 时传入的商户私钥
func TestClientEncryptDecryptText(t *testing.T) {
	text := "I love GoPay"

	cipherText, err := client.EncryptText(text)
	if err != nil {
		xlog.Errorf("client.EncryptText err: %+v", err)
		return
	}
	xlog.Debugf("encrypt text: %s", cipherText)

	originText, err := client.DecryptText(cipherText)
	if err != nil {
		xlog.Errorf("client.DecryptText err: %+v", err)
		return
	}
	xlog.Debugf("decrypt text: %s", originText)
}

// TestPkgEncryptDecryptText 使用包级函数进行加解密（不依赖 Client 实例）
func TestPkgEncryptDecryptText(t *testing.T) {
	text := "I love GoPay"

	cipherText, err := EncryptText(text, []byte(platformPublicKey))
	if err != nil {
		xlog.Errorf("EncryptText err: %+v", err)
		return
	}
	xlog.Debugf("encrypt text: %s", cipherText)

	originText, err := DecryptText(cipherText, []byte(merchantPrivateKey))
	if err != nil {
		xlog.Errorf("DecryptText err: %+v", err)
		return
	}
	xlog.Debugf("decrypt text: %s", originText)
}

// TestDecryptNotifyCipherText 演示回调 resource 密文的解密（AES-256-GCM）
// ciphertext / nonce / associated_data 来自回调 JSON 的 resource 字段
func TestDecryptNotifyCipherText(t *testing.T) {
	ciphertext := ""     // resource.ciphertext
	nonce := ""          // resource.nonce
	associatedData := "" // resource.associated_data

	bs, err := DecryptNotifyCipherTextToBytes(ciphertext, nonce, associatedData, ApiKey)
	if err != nil {
		xlog.Errorf("DecryptNotifyCipherTextToBytes err: %+v", err)
		return
	}
	xlog.Debugf("plain: %s", string(bs))
}
