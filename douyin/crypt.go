package douyin

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
)

// EncryptText 敏感字段加密（RSA/PKCS1v15）
// 使用最新有效的抖音支付平台证书公钥加密，需要先通过 SetPlatformCert 注册平台证书
func (c *Client) EncryptText(text string) (cipherText string, err error) {
	sn := c.NewestPlatformSerialNo()
	if sn == gopay.NULL {
		return gopay.NULL, errors.New("no platform cert found, please call SetPlatformCert() first")
	}
	pubKey, ok := c.getPlatformKey(sn)
	if !ok {
		return gopay.NULL, fmt.Errorf("platform cert of serial(%s) not found", sn)
	}
	cipherByte, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(text))
	if err != nil {
		return gopay.NULL, fmt.Errorf("rsa.EncryptPKCS1v15: %w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// DecryptText 敏感字段解密（使用商户私钥）
func (c *Client) DecryptText(cipherText string) (text string, err error) {
	cipherByte, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return gopay.NULL, fmt.Errorf("base64 decode failed: %w", err)
	}
	textByte, err := rsa.DecryptPKCS1v15(rand.Reader, c.privateKey, cipherByte)
	if err != nil {
		return gopay.NULL, fmt.Errorf("rsa.DecryptPKCS1v15: %w", err)
	}
	return string(textByte), nil
}

// EncryptText 敏感字段加密（外部工具版）
// pubKeyContent：抖音支付平台证书 PEM 内容
func EncryptText(text string, pubKeyContent []byte) (cipherText string, err error) {
	pubKey, err := xpem.DecodePublicKey(pubKeyContent)
	if err != nil {
		return gopay.NULL, err
	}
	cipherByte, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(text))
	if err != nil {
		return gopay.NULL, fmt.Errorf("rsa.EncryptPKCS1v15: %w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// DecryptText 敏感字段解密（外部工具版）
// privateKeyContent：商户 API 证书私钥 apiclient_key.pem 内容
func DecryptText(cipherText string, privateKeyContent []byte) (text string, err error) {
	priKey, err := xpem.DecodePrivateKey(privateKeyContent)
	if err != nil {
		return gopay.NULL, err
	}
	cipherByte, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return gopay.NULL, fmt.Errorf("base64 decode failed: %w", err)
	}
	textByte, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherByte)
	if err != nil {
		return gopay.NULL, fmt.Errorf("rsa.DecryptPKCS1v15: %w", err)
	}
	return string(textByte), nil
}

// DecryptNotifyCipherTextToBytes 回调密文 AES-256-GCM 解密到 []byte
// 使用 NewGCMWithNonceSize 兼容任意长度 nonce
func DecryptNotifyCipherTextToBytes(ciphertext, nonce, associatedData, apiKey string) (decrypt []byte, err error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("base64 decode ciphertext: %w", err)
	}
	block, err := aes.NewCipher([]byte(apiKey))
	if err != nil {
		return nil, fmt.Errorf("aes.NewCipher: %w", err)
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, fmt.Errorf("cipher.NewGCMWithNonceSize: %w", err)
	}
	decrypt, err = gcm.Open(nil, []byte(nonce), cipherBytes, []byte(associatedData))
	if err != nil {
		return nil, fmt.Errorf("gcm.Open: %w", err)
	}
	return decrypt, nil
}

// DecryptNotifyCipherTextToStruct 回调密文 AES-256-GCM 解密到结构体指针
func DecryptNotifyCipherTextToStruct(ciphertext, nonce, associatedData, apiKey string, objPtr any) (err error) {
	objValue := reflect.ValueOf(objPtr)
	if objValue.Kind() != reflect.Ptr {
		return errors.New("objPtr must be a pointer")
	}
	if objValue.Elem().Kind() != reflect.Struct {
		return errors.New("objPtr must point to a struct")
	}
	decrypt, err := DecryptNotifyCipherTextToBytes(ciphertext, nonce, associatedData, apiKey)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(decrypt, objPtr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s): %w", string(decrypt), err)
	}
	return nil
}
