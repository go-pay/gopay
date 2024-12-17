package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/util/convert"
)

// v3 鉴权请求 Authorization Header
func (a *ClientV3) authorization(method, uri string, bm gopay.BodyMap) (string, error) {
	var (
		jb        = ""
		timestamp = convert.Int64ToString(time.Now().UnixNano() / int64(time.Millisecond))
		nonceStr  = util.RandomString(32)
		// app_id=2014060600164699,app_cert_sn=xxx,nonce=5f9fba93-bbb2-40f0-b328-04d5ead3e131,timestamp=1667804301218
		authString = "app_id=" + a.AppId + ",app_cert_sn=" + a.AppCertSN + ",nonce=" + nonceStr + ",timestamp=" + timestamp
	)
	if a.AppCertSN == gopay.NULL {
		authString = "app_id=" + a.AppId + ",nonce=" + nonceStr + ",timestamp=" + timestamp
	}
	if bm != nil {
		jb = bm.JsonBody()
	}
	// ${authString}\n	步骤1中生成的认证串 authString。
	// ${httpMethod}\n	本次请求的 http 方法，例如 GET\POST\PUT 等。
	// ${httpReuqestUrl}\n   本次请求的 uri 信息，包括 queryString，不包括域名，例如 /v3/alipay/marketing/activity/ordervoucher/get?id=123。
	// ${httpRequestBody}\n	本次请求的 body 内容。当使用GET等请求时，body 为空，该值传入空字符串，即""。
	// ${appAuthToken}\n		应用授权令牌，和 header 参数中 alipay-app-auth-token 值保持一致。可选参数，不使用代调用模式时，不需要传入该字段。
	//
	// 示例：
	// app_id=2014060600164699,timestamp=1655869956477,nonce=eb4ade8f-8cfa-4ebf-a048-7eb52684ab32,expired_seconds=120
	// POST
	// /v3/alipay/marketing/activity/ordervoucher/create?auth_token=123
	// {"activity_name": "单品特价满10减1活动","publish_start_time": "2022-02-01 00:00:01"}
	//
	// body 空示例：
	// app_id=2014060600164699,timestamp=1655869956477,nonce=eb4ade8f-8cfa-4ebf-a048-7eb52684ab32,expired_seconds=120
	// GET
	// /v3/alipay/marketing/activity/ordervoucher?id=123
	//
	// 代调示例：
	// app_id=2014060600164699,timestamp=1655869956477,nonce=eb4ade8f-8cfa-4ebf-a048-7eb52684ab32,expired_seconds=120
	// GET
	// /v3/alipay/marketing/activity/ordervoucher?id=123
	// 202212BB_D64b2be8afd4b6c8468cf585bd05E50
	signStr := authString + "\n" + method + "\n" + uri + "\n" + jb + "\n"
	if a.AppAuthToken != "" {
		signStr += a.AppAuthToken + "\n"
	}
	if a.DebugSwitch == gopay.DebugOn {
		a.logger.Debugf("Alipay_V3_SignString:\n%s", signStr)
	}

	sign, err := a.rsaSign(signStr)
	if err != nil {
		return "", err
	}
	if a.DebugSwitch == gopay.DebugOn {
		a.logger.Debugf("Alipay_V3_Sign:\n%s", sign)
	}
	// authorization: ${签名算法} ${authString},sign=${signature}
	authorization := SignTypeRSA + " " + authString + ",sign=" + sign
	return authorization, nil
}

func (a *ClientV3) rsaSign(str string) (string, error) {
	if a.privateKey == nil {
		return "", errors.New("privateKey can't be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, a.privateKey, crypto.SHA256, h.Sum(nil))
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

func (a *ClientV3) autoVerifySignByCert(res *http.Response, body []byte) (err error) {
	if a.aliPayPublicKey != nil {
		ts := res.Header.Get(HeaderTimestamp)
		nonce := res.Header.Get(HeaderNonce)
		sign := res.Header.Get(HeaderSignature)
		if a.DebugSwitch == gopay.DebugOn {
			a.logger.Debugf("Alipay_VerifySignHeader: alipay-timestamp=[%s], alipay-nonce=[%s], alipay-signature=[%s]", ts, nonce, sign)
		}
		signData := ts + "\n" + nonce + "\n" + string(body) + "\n"

		signBytes, _ := base64.StdEncoding.DecodeString(sign)
		sum256 := sha256.Sum256([]byte(signData))
		if err = rsa.VerifyPKCS1v15(a.aliPayPublicKey, crypto.SHA256, sum256[:], signBytes); err != nil {
			return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
		}
	}
	return nil
}

// =============================== 通用底层验签方法 ===============================

func verifySign(signData, sign, alipayPublicKey string) (err error) {
	publicKey, err := xpem.DecodePublicKey([]byte(alipayPublicKey))
	if err != nil {
		return err
	}
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	h := sha256.New()
	h.Write([]byte(signData))
	if err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}

func verifySignCert(signData, sign string, alipayPublicKeyCert any) (err error) {
	var (
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

	h := sha256.New()
	h.Write([]byte(signData))
	if err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", gopay.VerifySignatureErr, err)
	}
	return nil
}
