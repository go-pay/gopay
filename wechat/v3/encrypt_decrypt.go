package wechat

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/util"
)

// 敏感信息加密，默认 PKCS1
func (c *ClientV3) V3EncryptText(publicKeyStr, text string) (cipherText string, err error) {

	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return util.NULL, errors.New("decode public key error")
	}

	var pubKey *rsa.PublicKey
	switch block.Type {
	case "PUBLIC KEY":
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return util.NULL, fmt.Errorf("ParsePKIXPublicKey err:%s", err.Error())
		}
		pKIXPublicKey, ok := pub.(*rsa.PublicKey)
		if !ok {
			return util.NULL, fmt.Errorf("断言ParsePKIXPublicKey异常 publicKeyStr:%s", publicKeyStr)
		}
		pubKey = pKIXPublicKey
	case "CERTIFICATE":
		pub, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return util.NULL, fmt.Errorf("ParseCertificate err:%s", err.Error())
		}
		certificatePubKey, ok := pub.PublicKey.(*rsa.PublicKey)
		if !ok {
			return util.NULL, fmt.Errorf("断言ParseCertificate异常 publicKeyStr:%s", publicKeyStr)
		}
		pubKey = certificatePubKey
	}

	cipherByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pubKey, []byte(text), nil)
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptOAEP：%w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// 敏感信息解密，默认 PKCS1
func (c *ClientV3) V3DecryptText(cipherText string) (text string, err error) {
	cipherByte, _ := base64.StdEncoding.DecodeString(cipherText)
	block, _ := pem.Decode(c.apiV3Key)
	if block == nil {
		return util.NULL, errors.New("pem.Decode：apiV3Key decode error")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509.ParsePKCS1PrivateKey：%w", err)
	}
	textByte, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherByte, nil)
	if err != nil {
		return "", fmt.Errorf("rsa.DecryptOAEP：%w", err)
	}
	return string(textByte), nil
}

// 敏感参数信息加密，默认 PKCS1
//	wxPkContent：微信平台证书内容
func V3EncryptText(text string, wxPkContent []byte) (cipherText string, err error) {
	block, _ := pem.Decode(wxPkContent)
	if block == nil {
		return util.NULL, errors.New("pem.Decode：rsaPublicKey decode error")
	}
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509.ParsePKCS1PublicKey：%w", err)
	}
	cipherByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pubKey, []byte(text), nil)
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptOAEP：%w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// 敏感参数信息解密，默认 PKCS1
//	apiV3Key：商户API证书字符串内容，商户平台获取
func V3DecryptText(cipherText string, apiV3Key []byte) (text string, err error) {
	cipherByte, _ := base64.StdEncoding.DecodeString(cipherText)
	block, _ := pem.Decode(apiV3Key)
	if block == nil {
		return util.NULL, errors.New("pem.Decode：rsaPrivateKey decode error")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509.ParsePKCS1PrivateKey：%w", err)
	}
	textByte, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherByte, nil)
	if err != nil {
		return "", fmt.Errorf("rsa.DecryptOAEP：%w", err)
	}
	return string(textByte), nil
}

// 解密 普通支付 回调中的加密信息
func V3DecryptNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密 普通退款 回调中的加密信息
func V3DecryptRefundNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptRefundResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptRefundResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密 合单支付 回调中的加密信息
func V3DecryptCombineNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptCombineResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptCombineResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密分账动账回调中的加密信息
func V3DecryptProfitShareNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptProfitShareResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptProfitShareResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密 支付分 回调中的加密信息
func V3DecryptScoreNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptScoreResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptScoreResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}
