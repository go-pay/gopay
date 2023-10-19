package alipay

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
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

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/pkg/xrsa"
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

/*
Q：使用公钥证书签名方式下，为什么开放平台网关的响应报文需要携带支付宝公钥证书SN（alipay_cert_sn）？
**
A：开发者上传自己的应用公钥证书后，开放平台会为开发者应用自动签发支付宝公钥证书供开发者下载，用来对开放平台网关响应报文做验签。

但是支付宝公钥证书可能因证书到期或者变更CA签发机构等原因，可能会重新签发证书。在重新签发前，开放平台会在门户上提前提醒开发者支付宝应用公钥证书变更时间。

但为避免开发者因未能及时感知支付宝公钥证书变更而导致验签失败，开放平台提供了一种支付宝公钥证书无感知升级机制，具体流程如下：
1）开放平台网关在响应报文中会多返回支付宝公钥证书SN
2）开放平台网关提供根据SN下载对应支付宝公钥证书的API接口
3）开发者在验签过程中，先比较本地使用的支付宝公钥证书SN与开放平台网关响应中SN是否一致。若不一致，可调用支付宝公钥证书下载接口下载对应SN的支付宝公钥证书。
4）对下载的支付宝公钥证书执行证书链校验，若校验通过，则用该证书验签。

基于该机制可实现支付宝公钥证书变更时开发者无感知，当前开放平台提供的SDK已基于该机制实现对应功能。若开发者未通过SDK接入，须自行实现该功能。
*/

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
			return util.NULL, err
		}
	case []byte:
		certData = pathOrData
	default:
		return util.NULL, errors.New("certPathOrData 证书类型断言错误")
	}

	if block, _ := pem.Decode(certData); block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return util.NULL, err
		}
		name := cert.Issuer.String()
		serialNumber := cert.SerialNumber.String()
		h := md5.New()
		h.Write([]byte(name))
		h.Write([]byte(serialNumber))
		sn = hex.EncodeToString(h.Sum(nil))
	}
	if sn == util.NULL {
		return util.NULL, errors.New("failed to get sn,please check your cert")
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
			return util.NULL, err
		}
	case []byte:
		certData = pathOrData
	default:
		return util.NULL, errors.New("rootCertPathOrData 断言异常")
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
			if sn == util.NULL {
				sn += hex.EncodeToString(h.Sum(nil))
			} else {
				sn += "_" + hex.EncodeToString(h.Sum(nil))
			}
		}
	}
	if sn == util.NULL {
		return util.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// 获取支付宝参数签名
// bm：签名参数
// signType：签名类型，alipay.RSA 或 alipay.RSA2
// privateKey：应用私钥，支持PKCS1和PKCS8
func GetRsaSign(bm gopay.BodyMap, signType string, privateKey *rsa.PrivateKey) (sign string, err error) {
	var (
		h              hash.Hash
		hashs          crypto.Hash
		encryptedBytes []byte
	)

	switch signType {
	case RSA:
		h = sha1.New()
		hashs = crypto.SHA1
	case RSA2:
		h = sha256.New()
		hashs = crypto.SHA256
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}
	signParams := bm.EncodeAliPaySignParams()
	if _, err = h.Write([]byte(signParams)); err != nil {
		return
	}
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, privateKey, hashs, h.Sum(nil)); err != nil {
		return util.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

func (a *Client) getRsaSign(bm gopay.BodyMap, signType string) (sign string, err error) {
	var (
		h              hash.Hash
		hashs          crypto.Hash
		encryptedBytes []byte
	)

	switch signType {
	case RSA:
		h = sha1.New()
		hashs = crypto.SHA1
	case RSA2:
		h = sha256.New()
		hashs = crypto.SHA256
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}
	signParams := bm.EncodeAliPaySignParams()
	if a.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("Alipay_Request_SignStr: %s", signParams)
	}
	if _, err = h.Write([]byte(signParams)); err != nil {
		return
	}
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, a.privateKey, hashs, h.Sum(nil)); err != nil {
		return util.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

// =============================== 获取SignData ===============================

// 需注意的是，公钥签名模式和公钥证书签名模式的不同之处
// 验签文档：https://opendocs.alipay.com/open/200/106120
func (a *Client) getSignData(bs []byte, alipayCertSN string) (signData string, err error) {
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
		if alipayPublicKeyCert == util.NULL {
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

func (a *Client) autoVerifySignByCert(sign, signData string, signDataErr error) (err error) {
	if a.autoSign && a.aliPayPublicKey != nil {
		if a.DebugSwitch == gopay.DebugOn {
			xlog.Debugf("Alipay_SyncSignData: %s, Sign=[%s]", signData, sign)
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
	if alipayPublicKey == util.NULL || notifyBean == nil {
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
		if aliPayPublicKeyCert == util.NULL {
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
