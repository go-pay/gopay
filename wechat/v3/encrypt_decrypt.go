package wechat

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xpem"
)

// 敏感信息加密，默认使用最新的有效微信平台证书加密
func (c *ClientV3) V3EncryptText(text string) (cipherText string, err error) {
	if c.wxPublicKey == nil || c.WxSerialNo == "" {
		return util.NULL, errors.New("WxPublicKey or WxSerialNo is null")
	}
	cipherByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, c.wxPublicKey, []byte(text), nil)
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptOAEP：%w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// 敏感信息解密
func (c *ClientV3) V3DecryptText(cipherText string) (text string, err error) {
	cipherByte, _ := base64.StdEncoding.DecodeString(cipherText)
	textByte, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, c.privateKey, cipherByte, nil)
	if err != nil {
		return "", fmt.Errorf("rsa.DecryptOAEP：%w", err)
	}
	return string(textByte), nil
}

// 敏感参数信息加密
// wxPublicKeyContent：微信平台证书内容
func V3EncryptText(text string, wxPublicKeyContent []byte) (cipherText string, err error) {
	publicKey, err := xpem.DecodePublicKey(wxPublicKeyContent)
	if err != nil {
		return gopay.NULL, err
	}
	cipherByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, []byte(text), nil)
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptOAEP：%w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// 敏感参数信息解密
// privateKeyContent：私钥 apiclient_key.pem 读取后的字符串内容
func V3DecryptText(cipherText string, privateKeyContent []byte) (text string, err error) {
	privateKey, err := xpem.DecodePrivateKey(privateKeyContent)
	if err != nil {
		return gopay.NULL, err
	}
	cipherByte, _ := base64.StdEncoding.DecodeString(cipherText)
	textByte, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, cipherByte, nil)
	if err != nil {
		return "", fmt.Errorf("rsa.DecryptOAEP：%w", err)
	}
	return string(textByte), nil
}

// 解密 通用方法ToBytes对象
func V3DecryptNotifyCipherTextToBytes(ciphertext, nonce, additional, apiV3Key string) (decrypt []byte, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err = aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	return decrypt, nil
}

// 解密 普通支付 回调中的加密信息
func V3DecryptNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 服务商支付 回调中的加密信息
func V3DecryptPartnerNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptPartnerResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptPartnerResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 普通退款 回调中的加密信息
func V3DecryptRefundNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptRefundResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptRefundResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 服务商退款 回调中的加密信息
func V3DecryptPartnerRefundNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptPartnerRefundResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptPartnerRefundResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 合单支付 回调中的加密信息
func V3DecryptCombineNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptCombineResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptCombineResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密分账动账回调中的加密信息
func V3DecryptProfitShareNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptProfitShareResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptProfitShareResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 支付分 回调中的加密信息
func V3DecryptScoreNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptScoreResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptScoreResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密商家券回调中的加密信息
func V3DecryptBusifavorNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptBusifavorResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptBusifavorResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}
