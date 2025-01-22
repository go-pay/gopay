package wechat

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-pay/crypto/aes"
	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
)

// 敏感信息加密，默认使用最新的有效微信平台证书加密
func (c *ClientV3) V3EncryptText(text string) (cipherText string, err error) {
	if c.wxPublicKey == nil || c.WxSerialNo == "" {
		return gopay.NULL, errors.New("WxPublicKey or WxSerialNo is null")
	}
	cipherByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, c.wxPublicKey, []byte(text), nil)
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptOAEP: %w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// 敏感信息解密
func (c *ClientV3) V3DecryptText(cipherText string) (text string, err error) {
	cipherByte, _ := base64.StdEncoding.DecodeString(cipherText)
	textByte, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, c.privateKey, cipherByte, nil)
	if err != nil {
		return "", fmt.Errorf("rsa.DecryptOAEP: %w", err)
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
		return "", fmt.Errorf("rsa.EncryptOAEP: %w", err)
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
		return "", fmt.Errorf("rsa.DecryptOAEP: %w", err)
	}
	return string(textByte), nil
}

// 解密 统一数据 到 []byte
func V3DecryptNotifyCipherTextToBytes(ciphertext, nonce, additional, apiV3Key string) (decrypt []byte, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err = aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	return decrypt, nil
}

// 解密 统一数据 到指针结构体对象
func V3DecryptNotifyCipherTextToStruct(ciphertext, nonce, additional, apiV3Key string, objPtr any) (err error) {
	//验证参数类型
	objValue := reflect.ValueOf(objPtr)
	if objValue.Kind() != reflect.Ptr {
		return errors.New("传入objPtr 参数类型必须指针")
	}
	//验证 any 类型
	if objValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入 any 必须是结构体")
	}
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	if err = json.Unmarshal(decrypt, objPtr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %#v), err:%w", string(decrypt), objPtr, err)
	}
	return nil
}

// 解密 普通支付 回调中的加密信息
func V3DecryptPayNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptPayResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptPayResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 服务商支付 回调中的加密信息
func V3DecryptPartnerPayNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptPartnerPayResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptPartnerPayResult{}
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

// 解密 分账动账 回调中的加密信息
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

// 解密 支付分确认订单 回调中的加密信息
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

// 解密 支付分开启/解除授权服务 回调中的加密信息
func V3DecryptScorePermissionNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptScorePermissionResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptScorePermissionResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 领券事件 回调中的加密信息
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

// 解密 停车入场状态变更 回调中的加密信息
func V3DecryptParkingInNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptParkingInResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptParkingInResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 停车支付结果 回调中的加密信息
func V3DecryptParkingPayNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptParkingPayResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptParkingPayResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 代金券核销事件 回调中的加密信息
func V3DecryptCouponUseNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptCouponResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptCouponResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 用户发票抬头填写完成 回调中的加密信息
func V3DecryptInvoiceTitleNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptInvoiceTitleResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptInvoiceTitleResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 发票卡券作废/发票开具成功/发票冲红成功/发票插入用户卡包成功 回调中的加密信息
func V3DecryptInvoiceNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptInvoiceResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptInvoiceResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 服务商子商户处置记录 回调中的加密信息
func V3DecryptViolationNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptViolationResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptViolationResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 商家转账批次回调通知 回调中的加密信息
func V3DecryptTransferBatchNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptTransferBatchResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptTransferBatchResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}

// 解密 新版商家转账通知 回调中的加密信息
func V3DecryptTransferBillsNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptTransferBillsResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%w", err)
	}
	result = &V3DecryptTransferBillsResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%w", string(decrypt), err)
	}
	return result, nil
}
