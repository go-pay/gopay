package cmbpay

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/go-pay/crypto/sm2"
)

// cmbDemoPubKey 是接口文档附录 1 给出的招行联调环境公钥（公开信息）。
const cmbDemoPubKey = "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE6Q+fktsnY9OFP+LpSR5Udbxf5zHCFO0PmOKlFNTxDIGl8jsPbbB/9ET23NV+acSz4FEkzD74sW2iiNVHRLiKHg=="

// TestLoadPublicKeyBase64 招行下发的 Base64 ASN.1 公钥必须能被正确解析。
func TestLoadPublicKeyBase64(t *testing.T) {
	pub, err := LoadPublicKeyBase64(cmbDemoPubKey)
	if err != nil {
		t.Fatalf("LoadPublicKeyBase64 失败: %v", err)
	}
	if pub.X == nil || pub.Y == nil {
		t.Fatal("解析出的公钥坐标为空")
	}
	// 重新编码后应还原出同一段 Base64，确认与招行的 ASN.1 编码一致。
	back, err := sm2.MarshalPublicKeyBase64(pub)
	if err != nil {
		t.Fatalf("MarshalPublicKeyBase64 失败: %v", err)
	}
	if back != cmbDemoPubKey {
		t.Errorf("公钥编码往返不一致\n got: %s\nwant: %s", back, cmbDemoPubKey)
	}
	if _, err = LoadPublicKeyBase64("not-base64!!"); err == nil {
		t.Error("非法 Base64 公钥应返回错误")
	}
}

// TestLoadPrivateKeyHex 私钥必须为 64 字符 HEX。
func TestLoadPrivateKeyHex(t *testing.T) {
	priv, err := sm2.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey 失败: %v", err)
	}
	keyHex := sm2.MarshalPrivateKeyHex(priv)
	loaded, err := LoadPrivateKeyHex(keyHex)
	if err != nil {
		t.Fatalf("LoadPrivateKeyHex 失败: %v", err)
	}
	if loaded.D.Cmp(priv.D) != 0 {
		t.Error("加载出的私钥与原私钥不一致")
	}
	for _, bad := range []string{"", "zz", keyHex[:62], keyHex + "00"} {
		if _, err = LoadPrivateKeyHex(bad); err == nil {
			t.Errorf("非法私钥 %q 应返回错误", bad)
		}
	}
}

// TestSignVerifySM2 报文体加签与验签的往返（接口文档 2.4.1.1 / 2.4.1.2）。
func TestSignVerifySM2(t *testing.T) {
	priv, err := sm2.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey 失败: %v", err)
	}
	data := buildSignString(map[string]string{
		"biz_content": `{"merId":"123","orderId":"abc"}`,
		"encoding":    Encoding,
		"signMethod":  SignMethodSM2,
		"version":     Version,
		"sign":        "应被剔除",
	})
	if strings.Contains(data, "应被剔除") {
		t.Fatal("待签名字符串中不应包含 sign 字段")
	}

	signB64, err := signSM2(priv, data)
	if err != nil {
		t.Fatalf("signSM2 失败: %v", err)
	}
	if _, err = base64.StdEncoding.DecodeString(signB64); err != nil {
		t.Errorf("签名不是合法 Base64: %v", err)
	}
	if err = verifySM2(priv.Public(), data, signB64); err != nil {
		t.Errorf("verifySM2 失败: %v", err)
	}
	// 报文被篡改、签名被篡改都必须验签失败。
	if err = verifySM2(priv.Public(), data+"x", signB64); err == nil {
		t.Error("报文被篡改时应验签失败")
	}
	if err = verifySM2(priv.Public(), data, "AAAA"); err == nil {
		t.Error("非法签名应验签失败")
	}
}

// TestSignSM2Hex 对账接口 apisign 要求 HEX 编码的 R||S（64 字节）。
func TestSignSM2Hex(t *testing.T) {
	priv, err := sm2.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey 失败: %v", err)
	}
	sigHex, err := signSM2Hex(priv, "appid=a&secret=b&sign=c&timestamp=d")
	if err != nil {
		t.Fatalf("signSM2Hex 失败: %v", err)
	}
	if len(sigHex) != 128 {
		t.Errorf("apisign 长度 = %d, 期望 128（R||S 各 32 字节的 HEX）", len(sigHex))
	}
	if err = sm2.VerifyHex(priv.Public(), []byte("appid=a&secret=b&sign=c&timestamp=d"), sigHex, nil); err != nil {
		t.Errorf("apisign 验签失败: %v", err)
	}
}

// TestEncryptorRoundTrip 敏感字段 SM4 加解密与数字信封生成（接口文档 2.4.4）。
func TestEncryptorRoundTrip(t *testing.T) {
	enc, err := NewEncryptor()
	if err != nil {
		t.Fatalf("NewEncryptor 失败: %v", err)
	}
	const plain = "6225880123456789"
	cipherB64, err := enc.Encrypt(plain)
	if err != nil {
		t.Fatalf("Encrypt 失败: %v", err)
	}
	got, err := decryptSM4(enc.sm4Key, cipherB64)
	if err != nil {
		t.Fatalf("decryptSM4 失败: %v", err)
	}
	if string(got) != plain {
		t.Errorf("SM4 往返 = %q, 期望 %q", got, plain)
	}

	// 数字信封：用招行公钥加密对称密钥，密文须为合法 Base64 且每次不同。
	pub, err := LoadPublicKeyBase64(cmbDemoPubKey)
	if err != nil {
		t.Fatalf("LoadPublicKeyBase64 失败: %v", err)
	}
	env1, err := enc.envelope(pub)
	if err != nil {
		t.Fatalf("envelope 失败: %v", err)
	}
	if _, err = base64.StdEncoding.DecodeString(env1); err != nil {
		t.Errorf("数字信封不是合法 Base64: %v", err)
	}
	env2, err := enc.envelope(pub)
	if err != nil {
		t.Fatalf("envelope 失败: %v", err)
	}
	if env1 == env2 {
		t.Error("两次数字信封结果相同，SM2 加密随机数可能未生效")
	}
	if _, err = enc.envelope(nil); err == nil {
		t.Error("公钥为 nil 时应返回错误")
	}
}

func TestNewEncryptorWithKey(t *testing.T) {
	if _, err := NewEncryptorWithKey(make([]byte, 8)); err == nil {
		t.Error("密钥长度非 16 字节应返回错误")
	}
	key := []byte("0123456789abcdef")
	enc, err := NewEncryptorWithKey(key)
	if err != nil {
		t.Fatalf("NewEncryptorWithKey 失败: %v", err)
	}
	// 必须持有副本，调用方后续修改入参不应影响已创建的 Encryptor。
	key[0] = 'X'
	if enc.sm4Key[0] != '0' {
		t.Error("NewEncryptorWithKey 未复制密钥")
	}
}

// TestHeaderSign 报文头 apisign 为固定拼接串的 MD5（接口文档 2.4.3）。
func TestHeaderSign(t *testing.T) {
	// printf 'appid=app&secret=sec&sign=sig&timestamp=123' | md5
	const want = "ad1208145c71998d47ada12a4799d53e"
	got := headerSign("app", "sec", "sig", "123")
	if got != want {
		t.Errorf("headerSign = %s, 期望 %s", got, want)
	}
	if got != strings.ToLower(got) {
		t.Error("apisign 应为小写十六进制")
	}
	// 任一入参变化都必须改变结果。
	if headerSign("app", "sec", "sig", "124") == got {
		t.Error("timestamp 变化后 apisign 未变")
	}
}

// TestBuildSignString 待签名字符串按 key 的 ASCII 升序拼接，且剔除 sign。
func TestBuildSignString(t *testing.T) {
	got := buildSignString(map[string]string{
		"version":     "0.0.1",
		"biz_content": "{}",
		"encoding":    "UTF-8",
		"sign":        "xxx",
		"empty":       "",
	})
	const want = "biz_content={}&empty=&encoding=UTF-8&version=0.0.1"
	if got != want {
		t.Errorf("buildSignString = %q, 期望 %q", got, want)
	}
	if buildSignString(nil) != "" {
		t.Error("空参数应返回空串")
	}
}
