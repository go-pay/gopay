package cmbpay

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
)

// LoadPrivateKeyHex 从 64 字节 HEX 字符串加载商户 SM2 私钥。
//
// 招行采用 SM2 标准秘钥格式，私钥为 32 字节字节流，转换为 HEX 即 64 字符
// （详见接口文档 2.4.1.3 国密秘钥标准规范）。例如联调私钥：
//
//	D5F2AFA24E6BA9071B54A8C9AD735F9A1DE9C4657FA386C09B592694BC118B38
func LoadPrivateKeyHex(hexKey string) (*sm2.PrivateKey, error) {
	raw, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 私钥 HEX 解析失败: %w", err)
	}
	if len(raw) != 32 {
		return nil, fmt.Errorf("cmbpay: 私钥长度应为 32 字节，实际 %d 字节", len(raw))
	}
	d := new(big.Int).SetBytes(raw)
	priv := new(sm2.PrivateKey)
	priv.Curve = sm2.P256Sm2()
	priv.D = d
	priv.PublicKey.Curve = priv.Curve
	priv.PublicKey.X, priv.PublicKey.Y = priv.Curve.ScalarBaseMult(d.Bytes())
	return priv, nil
}

// LoadPublicKeyBase64 从 Base64 编码的 ASN.1（ANS1）标准公钥加载招行 SM2 公钥。
//
// 招行公钥采用 base64 格式并符合 ASN.1 标准（详见接口文档 2.4.1.3），例如联调公钥：
//
//	MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE6Q+fktsnY9OFP+LpSR5Udbxf5zHCFO0Pm...
func LoadPublicKeyBase64(b64Key string) (*sm2.PublicKey, error) {
	der, err := base64.StdEncoding.DecodeString(b64Key)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 公钥 Base64 解析失败: %w", err)
	}
	pub, err := x509.ParseSm2PublicKey(der)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: 公钥 ASN.1 解析失败: %w", err)
	}
	return pub, nil
}

// sm2UserID 为国密局推荐的签名 USER_ID（详见接口文档 2.4.1.1）。
// var sm2UserID = []byte("1234567812345678")

// signSM2 使用商户私钥对待签名字符串做 SM2withSM3 裸签名（PKCS#1），
// 返回 Base64 编码后的签名值（详见接口文档 2.4.1.1）。
func signSM2(priv *sm2.PrivateKey, data string) (string, error) {
	// gmsm 的 Sign 使用国密标准默认 UID（1234567812345678），与招行要求一致。
	sig, err := priv.Sign(rand.Reader, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("cmbpay: SM2 加签失败: %w", err)
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

// signSM2Hex 使用商户私钥对待签名字符串做 SM2withSM3 签名，
// 返回 HEX 编码的 R||S 格式签名（64字节），用于对账接口的 apisign。
// 对应 Java 的 signHexBySm3WithSm2 方法。
func signSM2Hex(priv *sm2.PrivateKey, data string) (string, error) {
	// gmsm 的 Sign 返回 DER 格式的签名
	sig, err := priv.Sign(rand.Reader, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("cmbpay: SM2 加签失败: %w", err)
	}

	// 将 DER 签名转换为 R||S 格式（64字节）
	r, s, err := parseDERSignature(sig)
	if err != nil {
		return "", fmt.Errorf("cmbpay: 解析 DER 签名失败: %w", err)
	}

	// 拼接 R||S（各32字节）
	rBytes := formatTo32Bytes(r)
	sBytes := formatTo32Bytes(s)
	result := make([]byte, 64)
	copy(result[:32], rBytes)
	copy(result[32:], sBytes)

	return hex.EncodeToString(result), nil
}

// parseDERSignature 解析 DER 格式的 SM2 签名，返回 R 和 S 的 big.Int 值
func parseDERSignature(sig []byte) (*big.Int, *big.Int, error) {
	// DER 格式: 0x30 [total-length] 0x02 [r-length] [r] 0x02 [s-length] [s]
	if len(sig) < 8 || sig[0] != 0x30 {
		return nil, nil, errors.New("invalid DER signature format")
	}

	// 跳过 0x30 和总长度
	pos := 2
	if sig[pos] != 0x02 {
		return nil, nil, errors.New("invalid DER signature: missing R marker")
	}
	pos++

	// 读取 R
	rLen := int(sig[pos])
	pos++
	r := new(big.Int).SetBytes(sig[pos : pos+rLen])
	pos += rLen

	if sig[pos] != 0x02 {
		return nil, nil, errors.New("invalid DER signature: missing S marker")
	}
	pos++

	// 读取 S
	sLen := int(sig[pos])
	pos++
	s := new(big.Int).SetBytes(sig[pos : pos+sLen])

	return r, s, nil
}

// formatTo32Bytes 将 big.Int 格式化为32字节（左补零或右截取）
func formatTo32Bytes(n *big.Int) []byte {
	b := n.Bytes()
	if len(b) == 32 {
		return b
	}
	result := make([]byte, 32)
	if len(b) > 32 {
		// 右截取32字节
		copy(result, b[len(b)-32:])
	} else {
		// 左补零
		copy(result[32-len(b):], b)
	}
	return result
}

// verifySM2 使用招行公钥校验签名（详见接口文档 2.4.1.2）。
func verifySM2(pub *sm2.PublicKey, data, signB64 string) error {
	sig, err := base64.StdEncoding.DecodeString(signB64)
	if err != nil {
		return fmt.Errorf("cmbpay: 签名 Base64 解析失败: %w", err)
	}
	if !pub.Verify([]byte(data), sig) {
		return errors.New("cmbpay: 验签失败")
	}
	return nil
}

// encryptSM4 使用对称密钥对明文做 SM4-ECB 加密并 Base64 编码
// （详见接口文档 2.4.4 敏感信息加密 第 3 步）。
func encryptSM4(sm4Key, plain []byte) (string, error) {
	ct, err := sm4.Sm4Ecb(sm4Key, plain, true)
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
	pt, err := sm4.Sm4Ecb(sm4Key, ct, false)
	if err != nil {
		return nil, fmt.Errorf("cmbpay: SM4 解密失败: %w", err)
	}
	return pt, nil
}

// makeEnvelope 使用招行公钥对对称密钥做 SM2 加密（ASN.1 格式）并 Base64 编码，
// 生成数字信封（encryptKey 字段，详见接口文档 2.4.4 第 4 步）。
func makeEnvelope(pub *sm2.PublicKey, sm4Key []byte) (string, error) {
	enc, err := sm2.EncryptAsn1(pub, sm4Key, rand.Reader)
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
