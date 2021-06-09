package aes

import (
	"encoding/base64"
	"testing"

	"github.com/go-pay/gopay/pkg/xlog"
)

var (
	secretKey = "GYBh3Rmey7nNzR/NpV0vAw=="
	iv        = "JR3unO2glQuMhUx3"
)

func TestAesCBCEncryptDecrypt(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("originData:", originData)
	encryptData, err := CBCEncryptData([]byte(originData), []byte(secretKey))
	if err != nil {
		xlog.Error("AesCBCEncryptToString:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := CBCDecryptData(encryptData, []byte(secretKey))
	if err != nil {
		xlog.Error("AesDecryptToBytes:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}

func TestAesCBCEncryptDecryptIv(t *testing.T) {
	originData := "https://www.fumm.cc"
	xlog.Debug("originData:", originData)
	encryptData, err := CBCEncryptIvData([]byte(originData), []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("CBCEncryptIvData:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := CBCDecryptIvData(encryptData, []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("CBCDecryptIvData:", err)
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
