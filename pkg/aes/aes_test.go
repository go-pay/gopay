package aes

import (
	"encoding/base64"
	"testing"

	"github.com/go-pay/gopher/xlog"
)

var (
	secretKey = "JYRn4wbCy8KgVIZJaPhYTcTn2zixVC4Y"
	iv        = "JR3unO2glQuMhUx3"
)

func init() {
	xlog.Level = xlog.DebugLevel
}

func TestAesECBEncryptDecrypt(t *testing.T) {
	originData := "www.gopay.ink"
	xlog.Debug("originData:", originData)
	encryptData, err := ECBEncrypt([]byte(originData), []byte(secretKey))
	if err != nil {
		xlog.Error("AesCBCEncryptToString:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := ECBDecrypt(encryptData, []byte(secretKey))
	if err != nil {
		xlog.Error("AesDecryptToBytes:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}

func TestAesCBCEncryptDecrypt(t *testing.T) {
	originData := "www.gopay.ink"
	xlog.Debug("originData:", originData)
	encryptData, err := CBCEncrypt([]byte(originData), []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("CBCEncrypt:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := CBCDecrypt(encryptData, []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("CBCDecrypt:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}

func TestEncryptGCM(t *testing.T) {
	data := `我是要加密的数据`
	additional := "transaction"
	apiV3key := "Cj5xC9RXf0GFCKWeD9PyY1ZWLgionbvx"
	xlog.Debug("原始数据：", data)
	// 加密
	nonce, ciphertext, err := GCMEncrypt([]byte(data), []byte(additional), []byte(apiV3key))
	if err != nil {
		xlog.Error(err)
		return
	}
	encryptText := base64.StdEncoding.EncodeToString(ciphertext)
	xlog.Debug("加密后：", encryptText)
	xlog.Debug("nonce:", string(nonce))

	// 解密
	cipherBytes, _ := base64.StdEncoding.DecodeString(encryptText)
	decryptBytes, err := GCMDecrypt(cipherBytes, nonce, []byte(additional), []byte(apiV3key))
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("解密后：", string(decryptBytes))
}
