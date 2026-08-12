package cmbpay

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/go-pay/crypto/sm2"
	"github.com/go-pay/crypto/sm4"
)

// LoadPrivateKeyHex 从 64 字符 HEX 字符串加载商户 SM2 私钥。
//
// 招行采用 SM2 标准秘钥格式，私钥为 32 字节字节流，转换为 HEX 即 64 字符
// （详见接口文档 2.4.1.3 国密秘钥标准规范）。
func LoadPrivateKeyHex(hexKey string) (*sm2.PrivateKey, error) {
	priv, err := sm2.ParsePrivateKeyHex(hexKey)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 商户私钥加载失败: %w", err)
	}
	return priv, nil
}

// LoadPublicKeyBase64 从 Base64 编码的 ASN.1（ANS1）标准公钥加载招行 SM2 公钥。
//
// 招行公钥采用 base64 格式并符合 ASN.1 标准，形如
// MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE...（详见接口文档 2.4.1.3）。
func LoadPublicKeyBase64(b64Key string) (*sm2.PublicKey, error) {
	pub, err := sm2.ParsePublicKeyBase64(b64Key)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 招行公钥加载失败: %w", err)
	}
	return pub, nil
}

// signSM2 使用商户私钥对待签名字符串做 SM2withSM3 裸签名，
// 返回 Base64 编码后的签名值（详见接口文档 2.4.1.1）。
//
// 签名摘要使用国密标准默认 USER_ID（1234567812345678），与招行要求一致
// （详见接口文档 2.4.1.1）。
func signSM2(priv *sm2.PrivateKey, data string) (string, error) {
	sig, err := sm2.Sign(priv, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("cmbpay: SM2 加签失败: %w", err)
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

// signSM2Hex 使用商户私钥对待签名字符串做 SM2withSM3 签名，
// 返回 HEX 编码的 R||S 格式签名（64 字节），用于对账接口的 apisign。
// 对应招行 Java Demo 的 signHexBySm3WithSm2 方法。
func signSM2Hex(priv *sm2.PrivateKey, data string) (string, error) {
	sigHex, err := sm2.SignHex(priv, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("cmbpay: SM2 加签失败: %w", err)
	}
	return sigHex, nil
}

// verifySM2 使用招行公钥校验签名（详见接口文档 2.4.1.2）。
func verifySM2(pub *sm2.PublicKey, data, signB64 string) error {
	sig, err := base64.StdEncoding.DecodeString(signB64)
	if err != nil {
		return fmt.Errorf("cmbpay: 签名 Base64 解析失败: %w", err)
	}
	if err = sm2.Verify(pub, []byte(data), sig, nil); err != nil {
		return fmt.Errorf("cmbpay: 验签失败: %w", err)
	}
	return nil
}

// encryptSM4 使用对称密钥对明文做 SM4-ECB 加密并 Base64 编码
// （详见接口文档 2.4.4 敏感信息加密 第 3 步）。
func encryptSM4(sm4Key, plain []byte) (string, error) {
	ct, err := sm4.ECBEncrypt(plain, sm4Key)
	if err != nil {
		return "", fmt.Errorf("cmbpay: SM4 加密失败: %w", err)
	}
	return base64.StdEncoding.EncodeToString(ct), nil
}

// decryptSM4 对 Base64 编码的 SM4-ECB 密文解密还原明文。
func decryptSM4(sm4Key []byte, cipherB64 string) ([]byte, error) {
	ct, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: SM4 密文 Base64 解析失败: %w", err)
	}
	pt, err := sm4.ECBDecrypt(ct, sm4Key)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: SM4 解密失败: %w", err)
	}
	return pt, nil
}

// makeEnvelope 使用招行公钥对对称密钥做 SM2 加密（ASN.1 格式）并 Base64 编码，
// 生成数字信封（encryptKey 字段，详见接口文档 2.4.4 第 4 步）。
func makeEnvelope(pub *sm2.PublicKey, sm4Key []byte) (string, error) {
	if pub == nil {
		return "", errors.New("cmbpay: 招行公钥为空，无法生成数字信封")
	}
	enc, err := sm2.Encrypt(pub, sm4Key)
	if err != nil {
		return "", fmt.Errorf("cmbpay: 数字信封 SM2 加密失败: %w", err)
	}
	return base64.StdEncoding.EncodeToString(enc), nil
}

// headerSign 计算报文头的 apisign（详见接口文档 2.4.3 APP ID 校验）。
//
// 规则：将 appid、secret、sign（报文体签名）、timestamp 按 KEY 首字母排序并以 &
// 连接为 appid=..&secret=..&sign=..&timestamp=..，再做 MD5，结果为小写十六进制。
func headerSign(appID, appSecret, bodySign, timestamp string) string {
	raw := fmt.Sprintf("appid=%s&secret=%s&sign=%s&timestamp=%s",
		appID, appSecret, bodySign, timestamp)
	sum := md5.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])
}
