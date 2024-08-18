package alipay

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/crypto/xrsa"
	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/util/convert"
)

// 允许进行 sn 提取的证书签名算法
var allowSignatureAlgorithm = map[string]bool{
	"MD2-RSA":       true,
	"MD5-RSA":       true,
	"SHA1-RSA":      true,
	"SHA256-RSA":    true,
	"SHA384-RSA":    true,
	"SHA512-RSA":    true,
	"SHA256-RSAPSS": true,
	"SHA384-RSAPSS": true,
	"SHA512-RSAPSS": true,
}

// GetCertSN 获取证书序列号SN
// certPathOrData x509证书文件路径(appPublicCert.crt、alipayPublicCert.crt) 或证书 buffer
// 返回 sn：证书序列号(app_cert_sn、alipay_cert_sn)
// 返回 err：error 信息
func GetCertSN(certPathOrData any) (sn string, err error) {
	var certData []byte
	switch pathOrData := certPathOrData.(type) {
	case string:
		certData, err = os.ReadFile(pathOrData)
		if err != nil {
			return gopay.NULL, err
		}
	case []byte:
		certData = pathOrData
	default:
		return gopay.NULL, errors.New("certPathOrData 证书类型断言错误")
	}

	if block, _ := pem.Decode(certData); block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return gopay.NULL, err
		}
		name := cert.Issuer.String()
		serialNumber := cert.SerialNumber.String()
		h := md5.New()
		h.Write([]byte(name))
		h.Write([]byte(serialNumber))
		sn = hex.EncodeToString(h.Sum(nil))
	}
	if sn == gopay.NULL {
		return gopay.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// GetRootCertSN 获取root证书序列号SN
// rootCertPathOrData x509证书文件路径(alipayRootCert.crt) 或文件 buffer
// 返回 sn：证书序列号(alipay_root_cert_sn)
// 返回 err：error 信息
func GetRootCertSN(rootCertPathOrData any) (sn string, err error) {
	var (
		certData []byte
		certEnd  = `-----END CERTIFICATE-----`
	)
	switch pathOrData := rootCertPathOrData.(type) {
	case string:
		certData, err = os.ReadFile(pathOrData)
		if err != nil {
			return gopay.NULL, err
		}
	case []byte:
		certData = pathOrData
	default:
		return gopay.NULL, errors.New("rootCertPathOrData 断言异常")
	}

	pems := strings.Split(string(certData), certEnd)
	for _, c := range pems {
		if block, _ := pem.Decode([]byte(c + certEnd)); block != nil {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				continue
			}
			if !allowSignatureAlgorithm[cert.SignatureAlgorithm.String()] {
				continue
			}
			name := cert.Issuer.String()
			serialNumber := cert.SerialNumber.String()
			h := md5.New()
			h.Write([]byte(name))
			h.Write([]byte(serialNumber))
			if sn == gopay.NULL {
				sn += hex.EncodeToString(h.Sum(nil))
			} else {
				sn += "_" + hex.EncodeToString(h.Sum(nil))
			}
		}
	}
	if sn == gopay.NULL {
		return gopay.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// v3 鉴权请求 authorization Header
func (c *ClientV3) authorization(method, path string, bm gopay.BodyMap) (string, error) {
	var (
		jb        = ""
		timestamp = convert.Int64ToString(time.Now().UnixNano() / int64(time.Millisecond))
		nonceStr  = util.RandomString(32)
		// app_id=2014060600164699,app_cert_sn=xxx,nonce=5f9fba93-bbb2-40f0-b328-04d5ead3e131,timestamp=1667804301218
		authString = "app_id=" + c.AppId + ",app_cert_sn=" + c.AppCertSN + ",nonce_str=" + nonceStr + ",timestamp=" + timestamp
	)
	if c.AppCertSN == gopay.NULL {
		authString = "app_id=" + c.AppId + ",nonce=" + nonceStr + ",timestamp=" + timestamp
	}
	if bm != nil {
		jb = bm.JsonBody()
	}
	//${authString}\n
	//${httpMethod}\n
	//${httpReuqestUrl}\n
	//${httpRequestBody}\n
	//${appAuthToken}\n
	signStr := authString + "\n" + method + "\n" + path + "\n" + jb + "\n"
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Alipay_V3_SignString:\n%s", signStr)
	}
	sign, err := c.rsaSign(signStr)
	if err != nil {
		return "", err
	}
	if c.DebugSwitch == gopay.DebugOn {
		c.logger.Debugf("Alipay_V3_Sign:\n%s", sign)
	}
	// authorization: ${签名算法} ${authString},sign=${signature}
	authorization := SignTypeRSA + " " + authString + ",sign=" + sign
	return authorization, nil
}

func (c *ClientV3) rsaSign(str string) (string, error) {
	if c.privateKey == nil {
		return "", errors.New("privateKey can't be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return gopay.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

// =============================== 获取SignData ===============================

// 需注意的是，公钥签名模式和公钥证书签名模式的不同之处
// 验签文档：https://opendocs.alipay.com/open/200/106120
func (a *ClientV3) getSignData(bs []byte, alipayCertSN string) (signData string, err error) {
	var (
		str        = string(bs)
		indexStart = strings.Index(str, `_response":`)
		indexEnd   int
	)
	indexStart = indexStart + 11
	bsLen := len(str)
	if alipayCertSN != "" {
		// 公钥证书模式
		if alipayCertSN != a.AliPayPublicCertSN {
			return gopay.NULL, fmt.Errorf("[%w], 当前使用的支付宝公钥证书SN[%s]与网关响应报文中的SN[%s]不匹配", gopay.CertNotMatchErr, a.AliPayPublicCertSN, alipayCertSN)
		}
		indexEnd = strings.Index(str, `,"alipay_cert_sn":`)
		if indexEnd > indexStart && bsLen > indexStart {
			signData = str[indexStart:indexEnd]
			return
		}
		return gopay.NULL, fmt.Errorf("[%w], value: %s", gopay.GetSignDataErr, str)
	}
	// 普通公钥模式
	indexEnd = strings.Index(str, `,"sign":`)
	if indexEnd > indexStart && bsLen > indexStart {
		signData = str[indexStart:indexEnd]
		return
	}
	return gopay.NULL, fmt.Errorf("[%w], value: %s", gopay.GetSignDataErr, str)
}

// =============================== 同步验签 ===============================

// VerifySyncSign 支付宝同步返回验签（公钥模式）
// 注意：APP支付，手机网站支付，电脑网站支付，身份认证开始认证 不支持同步返回验签
// aliPayPublicKey：支付宝平台获取的支付宝公钥
// signData：待验签参数，aliRsp.SignData
// sign：待验签sign，aliRsp.Sign
// 返回参数ok：是否验签通过
// 返回参数err：错误信息
// 验签文档：https://opendocs.alipay.com/open/200/106120
func VerifySyncSign(aliPayPublicKey, signData, sign string) (ok bool, err error) {
	// 支付宝公钥验签
	pKey := xrsa.FormatAlipayPublicKey(aliPayPublicKey)
	if err = verifySign(signData, sign, RSA2, pKey); err != nil {
		return false, err
	}
	return true, nil
}

// VerifySyncSignWithCert 支付宝同步返回验签（公钥证书模式）
// 注意：APP支付，手机网站支付，电脑网站支付，身份认证开始认证 不支持同步返回验签
// aliPayPublicKeyCert：支付宝公钥证书存放路径 alipayPublicCert.crt 或文件内容[]byte
// signData：待验签参数，aliRsp.SignData
// sign：待验签sign，aliRsp.Sign
// 返回参数ok：是否验签通过
// 返回参数err：错误信息
// 验签文档：https://opendocs.alipay.com/open/200/106120
func VerifySyncSignWithCert(alipayPublicKeyCert any, signData, sign string) (ok bool, err error) {
	switch alipayPublicKeyCert.(type) {
	case string:
		if alipayPublicKeyCert == gopay.NULL {
			return false, errors.New("aliPayPublicKeyPath is null")
		}
	case []byte:
	default:
		return false, errors.New("alipayPublicKeyCert type assert error")
	}
	if err = verifySignCert(signData, sign, RSA2, alipayPublicKeyCert); err != nil {
		return false, err
	}
	return true, nil
}

func (a *ClientV3) autoVerifySignByCert(sign, signData string, signDataErr error) (err error) {
	if a.autoSign && a.aliPayPublicKey != nil {
		if a.DebugSwitch == gopay.DebugOn {
			a.logger.Debugf("Alipay_SyncSignData: %s, Sign=[%s]", signData, sign)
		}
		// 只有证书验签时，才可能出现此error
		if signDataErr != nil {
			return signDataErr
		}

		signBytes, _ := base64.StdEncoding.DecodeString(sign)
		hashs := crypto.SHA256
		h := hashs.New()
		h.Write([]byte(signData))
		if err = rsa.VerifyPKCS1v15(a.aliPayPublicKey, hashs, h.Sum(nil), signBytes); err != nil {
			return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
		}
	}
	return nil
}

// =============================== 异步验签 ===============================

// VerifySign 支付宝异步通知验签（公钥模式）
// 注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
// alipayPublicKey：支付宝平台获取的支付宝公钥
// notifyBean：此参数为异步通知解析的结构体或BodyMap：notifyReq 或 bm，推荐通 BodyMap 验签
// 返回参数ok：是否验签通过
// 返回参数err：错误信息
// 验签文档：https://opendocs.alipay.com/open/200/106120
func VerifySign(alipayPublicKey string, notifyBean any) (ok bool, err error) {
	if alipayPublicKey == gopay.NULL || notifyBean == nil {
		return false, errors.New("alipayPublicKey or notifyBean is nil")
	}
	var (
		bodySign     string
		bodySignType string
		signData     string
		bm           = make(gopay.BodyMap)
	)
	if reflect.ValueOf(notifyBean).Kind() == reflect.Map {
		if bm, ok = notifyBean.(gopay.BodyMap); ok {
			bodySign = bm.GetString("sign")
			bodySignType = bm.GetString("sign_type")
			bm.Remove("sign")
			bm.Remove("sign_type")
			signData = bm.EncodeAliPaySignParams()
		}
	} else {
		bs, err := json.Marshal(notifyBean)
		if err != nil {
			return false, fmt.Errorf("json.Marshal：%w", err)
		}
		if err = json.Unmarshal(bs, &bm); err != nil {
			return false, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
		}
		bodySign = bm.GetString("sign")
		bodySignType = bm.GetString("sign_type")
		bm.Remove("sign")
		bm.Remove("sign_type")
		signData = bm.EncodeAliPaySignParams()
	}
	pKey := xrsa.FormatAlipayPublicKey(alipayPublicKey)
	if err = verifySign(signData, bodySign, bodySignType, pKey); err != nil {
		return false, err
	}
	return true, nil
}

// 支付宝异步通知验签（公钥证书模式）
// 注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
// aliPayPublicKeyCert：支付宝公钥证书存放路径 alipayPublicCert.crt 或文件内容[]byte
// notifyBean：此参数为异步通知解析的结构体或BodyMap：notifyReq 或 bm，推荐通 BodyMap 验签
// 返回参数ok：是否验签通过
// 返回参数err：错误信息
// 验签文档：https://opendocs.alipay.com/open/200/106120
func VerifySignWithCert(aliPayPublicKeyCert, notifyBean any) (ok bool, err error) {
	if notifyBean == nil || aliPayPublicKeyCert == nil {
		return false, errors.New("aliPayPublicKeyCert or notifyBean is nil")
	}
	switch aliPayPublicKeyCert.(type) {
	case string:
		if aliPayPublicKeyCert == gopay.NULL {
			return false, errors.New("aliPayPublicKeyPath is null")
		}
	case []byte:
	default:
		return false, errors.New("aliPayPublicKeyCert type assert error")
	}
	var bm gopay.BodyMap

	switch nb := notifyBean.(type) {
	case map[string]any:
		bm = make(gopay.BodyMap, len(nb))
		for key, val := range nb {
			bm[key] = val
		}
	case gopay.BodyMap:
		bm = nb
	default:
		bs, err := json.Marshal(notifyBean)
		if err != nil {
			return false, fmt.Errorf("json.Marshal：%w", err)
		}
		if err = json.Unmarshal(bs, &bm); err != nil {
			return false, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
		}
	}
	bodySign := bm.GetString("sign")
	bodySignType := bm.GetString("sign_type")
	bm.Remove("sign")
	bm.Remove("sign_type")
	signData := bm.EncodeAliPaySignParams()
	if err = verifySignCert(signData, bodySign, bodySignType, aliPayPublicKeyCert); err != nil {
		return false, err
	}
	return true, nil
}

// =============================== 通用底层验签方法 ===============================

func verifySign(signData, sign, signType, alipayPublicKey string) (err error) {
	var (
		h     hash.Hash
		hashs crypto.Hash
	)
	publicKey, err := xpem.DecodePublicKey([]byte(alipayPublicKey))
	if err != nil {
		return err
	}
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	switch signType {
	case RSA:
		hashs = crypto.SHA1
	case RSA2:
		hashs = crypto.SHA256
	default:
		hashs = crypto.SHA256
	}
	h = hashs.New()
	h.Write([]byte(signData))
	if err = rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}

func verifySignCert(signData, sign, signType string, alipayPublicKeyCert any) (err error) {
	var (
		h     hash.Hash
		hashs crypto.Hash
		bytes []byte
	)
	if v, ok := alipayPublicKeyCert.(string); ok {
		if bytes, err = os.ReadFile(v); err != nil {
			return fmt.Errorf("支付宝公钥文件读取失败: %w", err)
		}
	} else {
		bytes, ok = alipayPublicKeyCert.([]byte)
		if !ok {
			return fmt.Errorf("支付宝公钥读取失败: %w", err)
		}
	}
	publicKey, err := xpem.DecodePublicKey(bytes)
	if err != nil {
		return err
	}
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	switch signType {
	case RSA:
		hashs = crypto.SHA1
	case RSA2:
		hashs = crypto.SHA256
	default:
		hashs = crypto.SHA256
	}
	h = hashs.New()
	h.Write([]byte(signData))
	if err = rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}
