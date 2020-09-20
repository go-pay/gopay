package alipay

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
	xaes "github.com/iGoogle-ink/gotil/aes"
	"github.com/iGoogle-ink/gotil/xhttp"
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

// 解析支付宝支付异步通知的参数到BodyMap
//	req：*http.Request
//	返回参数bm：Notify请求的参数
//	返回参数err：错误信息
//	文档：https://opendocs.alipay.com/open/203/105286
func ParseNotifyToBodyMap(req *http.Request) (bm gopay.BodyMap, err error) {
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	var form map[string][]string = req.Form
	bm = make(gopay.BodyMap, len(form)+1)
	for k, v := range form {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return
}

// 通过 url.Values 解析支付宝支付异步通知的参数到Struct
//	value：url.Values
//	返回参数notifyReq：Notify请求的参数
//	返回参数err：错误信息
//	文档：https://opendocs.alipay.com/open/203/105286
func ParseNotifyByURLValues(value url.Values) (bm gopay.BodyMap, err error) {
	bm = make(gopay.BodyMap, len(value)+1)
	for k, v := range value {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return
}

// Deprecated
// 解析支付宝支付异步通知的参数到Struct
//	req：*http.Request
//	返回参数notifyReq：Notify请求的参数
//	返回参数err：错误信息
//	文档：https://opendocs.alipay.com/open/203/105286
func ParseNotifyResult(req *http.Request) (notifyReq *NotifyRequest, err error) {
	notifyReq = new(NotifyRequest)
	if err = req.ParseForm(); err != nil {
		return
	}
	notifyReq.NotifyTime = req.Form.Get("notify_time")
	notifyReq.NotifyType = req.Form.Get("notify_type")
	notifyReq.NotifyId = req.Form.Get("notify_id")
	notifyReq.AppId = req.Form.Get("app_id")
	notifyReq.Charset = req.Form.Get("charset")
	notifyReq.Version = req.Form.Get("version")
	notifyReq.SignType = req.Form.Get("sign_type")
	notifyReq.Sign = req.Form.Get("sign")
	notifyReq.AuthAppId = req.Form.Get("auth_app_id")
	notifyReq.TradeNo = req.Form.Get("trade_no")
	notifyReq.OutTradeNo = req.Form.Get("out_trade_no")
	notifyReq.OutBizNo = req.Form.Get("out_biz_no")
	notifyReq.BuyerId = req.Form.Get("buyer_id")
	notifyReq.BuyerLogonId = req.Form.Get("buyer_logon_id")
	notifyReq.SellerId = req.Form.Get("seller_id")
	notifyReq.SellerEmail = req.Form.Get("seller_email")
	notifyReq.TradeStatus = req.Form.Get("trade_status")
	notifyReq.TotalAmount = req.Form.Get("total_amount")
	notifyReq.ReceiptAmount = req.Form.Get("receipt_amount")
	notifyReq.InvoiceAmount = req.Form.Get("invoice_amount")
	notifyReq.BuyerPayAmount = req.Form.Get("buyer_pay_amount")
	notifyReq.PointAmount = req.Form.Get("point_amount")
	notifyReq.RefundFee = req.Form.Get("refund_fee")
	notifyReq.Subject = req.Form.Get("subject")
	notifyReq.Body = req.Form.Get("body")
	notifyReq.GmtCreate = req.Form.Get("gmt_create")
	notifyReq.GmtPayment = req.Form.Get("gmt_payment")
	notifyReq.GmtRefund = req.Form.Get("gmt_refund")
	notifyReq.GmtClose = req.Form.Get("gmt_close")
	notifyReq.PassbackParams = req.Form.Get("passback_params")

	billList := req.Form.Get("fund_bill_list")
	if billList != gotil.NULL {
		bills := make([]*FundBillListInfo, 0)
		if err = json.Unmarshal([]byte(billList), &bills); err != nil {
			return nil, fmt.Errorf(`"fund_bill_list" xml.Unmarshal(%s)：%w`, billList, err)
		}
		notifyReq.FundBillList = bills
	} else {
		notifyReq.FundBillList = nil
	}

	detailList := req.Form.Get("voucher_detail_list")
	if detailList != gotil.NULL {
		details := make([]*VoucherDetailListInfo, 0)
		if err = json.Unmarshal([]byte(detailList), &details); err != nil {
			return nil, fmt.Errorf(`"voucher_detail_list" xml.Unmarshal(%s)：%w`, detailList, err)
		}
		notifyReq.VoucherDetailList = details
	} else {
		notifyReq.VoucherDetailList = nil
	}
	return
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

// VerifySyncSign 支付宝同步返回验签
//	注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
//	aliPayPublicKey：支付宝公钥
//	signData：待验签参数，aliRsp.SignData
//	sign：待验签sign，aliRsp.Sign
//	返回参数ok：是否验签通过
//	返回参数err：错误信息
//	验签文档：https://opendocs.alipay.com/open/200/106120
func VerifySyncSign(aliPayPublicKey, signData, sign string) (ok bool, err error) {
	// 支付宝公钥验签
	pKey := FormatPublicKey(aliPayPublicKey)
	if err = verifySign(signData, sign, RSA2, pKey); err != nil {
		return false, err
	}
	return true, nil
}

// VerifySign 支付宝异步通知验签
//	注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
//	aliPayPublicKey：支付宝公钥
//	bean：此参数为异步通知解析的结构体或BodyMap：notifyReq 或 bm，推荐通 BodyMap 验签
//	返回参数ok：是否验签通过
//	返回参数err：错误信息
//	验签文档：https://opendocs.alipay.com/open/200/106120
func VerifySign(aliPayPublicKey string, bean interface{}) (ok bool, err error) {
	if aliPayPublicKey == gotil.NULL {
		return false, errors.New("aliPayPublicKey is null")
	}
	if bean == nil {
		return false, errors.New("bean is nil")
	}
	var (
		bodySign     string
		bodySignType string
		signData     string
		bm           = make(gopay.BodyMap)
	)
	if reflect.ValueOf(bean).Kind() == reflect.Map {
		if bm, ok = bean.(gopay.BodyMap); ok {
			bodySign = bm.Get("sign")
			bodySignType = bm.Get("sign_type")
			bm.Remove("sign")
			bm.Remove("sign_type")
			signData = bm.EncodeAliPaySignParams()
		}
	} else {
		bs, err := json.Marshal(bean)
		if err != nil {
			return false, fmt.Errorf("json.Marshal：%w", err)
		}
		if err = json.Unmarshal(bs, &bm); err != nil {
			return false, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
		}
		bodySign = bm.Get("sign")
		bodySignType = bm.Get("sign_type")
		bm.Remove("sign")
		bm.Remove("sign_type")
		signData = bm.EncodeAliPaySignParams()
	}
	pKey := FormatPublicKey(aliPayPublicKey)
	if err = verifySign(signData, bodySign, bodySignType, pKey); err != nil {
		return false, err
	}
	return true, nil
}

// VerifySignWithCert 支付宝异步通知验签
//	注意：APP支付，手机网站支付，电脑网站支付 暂不支持同步返回验签
//	aliPayPublicKey：支付宝公钥存放路径 alipayCertPublicKey_RSA2.crt 或文件 buffer
//	bean：此参数为异步通知解析的结构体或BodyMap：notifyReq 或 bm，推荐通 BodyMap 验签
//	返回参数ok：是否验签通过
//	返回参数err：错误信息
//	验签文档：https://opendocs.alipay.com/open/200/106120
func VerifySignWithCert(aliPayPublicKey, bean interface{}) (ok bool, err error) {
	if bean == nil || aliPayPublicKey == nil {
		return false, errors.New("aliPayPublicKey or bean is nil")
	}
	switch aliPayPublicKey.(type) {
	case string:
		if aliPayPublicKey == gotil.NULL {
			return false, errors.New("aliPayPublicKeyPath is null")
		}
	case []byte:
	default:
		return false, errors.New("aliPayPublicKeyPath type assert error")
	}
	var (
		bodySign     string
		bodySignType string
		signData     string
		bm           = make(gopay.BodyMap)
	)
	if reflect.ValueOf(bean).Kind() == reflect.Map {
		if bm, ok = bean.(gopay.BodyMap); ok {
			bodySign = bm.Get("sign")
			bodySignType = bm.Get("sign_type")
			bm.Remove("sign")
			bm.Remove("sign_type")
			signData = bm.EncodeAliPaySignParams()
		}
	} else {
		bs, err := json.Marshal(bean)
		if err != nil {
			return false, fmt.Errorf("json.Marshal：%w", err)
		}
		if err = json.Unmarshal(bs, &bm); err != nil {
			return false, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
		}
		bodySign = bm.Get("sign")
		bodySignType = bm.Get("sign_type")
		bm.Remove("sign")
		bm.Remove("sign_type")
		signData = bm.EncodeAliPaySignParams()
	}
	if err = verifySignCert(signData, bodySign, bodySignType, aliPayPublicKey); err != nil {
		return false, err
	}
	return true, nil
}

func verifySign(signData, sign, signType, aliPayPublicKey string) (err error) {
	var (
		h         hash.Hash
		hashs     crypto.Hash
		block     *pem.Block
		pubKey    interface{}
		publicKey *rsa.PublicKey
		ok        bool
	)
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	if block, _ = pem.Decode([]byte(aliPayPublicKey)); block == nil {
		return errors.New("支付宝公钥Decode错误")
	}
	if pubKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return fmt.Errorf("x509.ParsePKIXPublicKey：%w", err)
	}
	if publicKey, ok = pubKey.(*rsa.PublicKey); !ok {
		return errors.New("支付宝公钥转换错误")
	}
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
	return rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
}

func verifySignCert(signData, sign, signType string, aliPayPublicKey interface{}) (err error) {
	var (
		h         hash.Hash
		hashs     crypto.Hash
		block     *pem.Block
		pubKey    *x509.Certificate
		publicKey *rsa.PublicKey
		ok        bool
		bytes     []byte
	)
	if v, ok := aliPayPublicKey.(string); ok {
		if bytes, err = ioutil.ReadFile(v); err != nil {
			return fmt.Errorf("支付宝公钥文件读取失败: %w", err)
		}
	} else {
		bytes, ok = aliPayPublicKey.([]byte)
		if !ok {
			return fmt.Errorf("支付宝公钥读取失败: %w", err)
		}
	}
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	if block, _ = pem.Decode(bytes); block == nil {
		return errors.New("支付宝公钥Decode错误")
	}
	if pubKey, err = x509.ParseCertificate(block.Bytes); err != nil {
		return fmt.Errorf("x509.ParseCertificate：%w", err)
	}
	if publicKey, ok = pubKey.PublicKey.(*rsa.PublicKey); !ok {
		return errors.New("支付宝公钥转换错误")
	}
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
	return rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
}

// FormatPrivateKey 格式化 普通应用秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

// FormatPublicKey 格式化 普通支付宝公钥
func FormatPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

// GetCertSN 获取证书序列号SN
//	certPathOrData.509证书文件路径(appCertPublicKey.crt、alipayCertPublicKey_RSA2.crt) 或证书 buffer
//	返回 sn：证书序列号(app_cert_sn、alipay_cert_sn)
//	返回 err：error 信息
func GetCertSN(certPathOrData interface{}) (sn string, err error) {
	var certData []byte
	switch certPathOrData.(type) {
	case string:
		certData, err = ioutil.ReadFile(certPathOrData.(string))
	case []byte:
		certData = certPathOrData.([]byte)
	default:
		return gotil.NULL, errors.New("certPathOrData 证书类型断言错误")
	}
	if err != nil {
		return gotil.NULL, err
	}

	if block, _ := pem.Decode(certData); block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return gotil.NULL, err
		}
		name := cert.Issuer.String()
		serialNumber := cert.SerialNumber.String()
		h := md5.New()
		h.Write([]byte(name))
		h.Write([]byte(serialNumber))
		sn = hex.EncodeToString(h.Sum(nil))
	}
	if sn == gotil.NULL {
		return gotil.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// GetRootCertSN 获取root证书序列号SN
//	rootCertPathOrData.509证书文件路径(alipayRootCert.crt) 或文件 buffer
//	返回 sn：证书序列号(alipay_root_cert_sn)
//	返回 err：error 信息
func GetRootCertSN(rootCertPathOrData interface{}) (sn string, err error) {
	var certData []byte
	var certEnd = `-----END CERTIFICATE-----`
	switch rootCertPathOrData.(type) {
	case string:
		certData, err = ioutil.ReadFile(rootCertPathOrData.(string))
	case []byte:
		certData = rootCertPathOrData.([]byte)
	default:
		return gotil.NULL, errors.New("rootCertPathOrData 断言异常")
	}
	if err != nil {
		return gotil.NULL, err
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
			if sn == gotil.NULL {
				sn += hex.EncodeToString(h.Sum(nil))
			} else {
				sn += "_" + hex.EncodeToString(h.Sum(nil))
			}
		}
	}
	if sn == gotil.NULL {
		return gotil.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// DecryptOpenDataToStruct 解密支付宝开放数据到 结构体
//	encryptedData:包括敏感数据在内的完整用户信息的加密数据
//	secretKey:AES密钥，支付宝管理平台配置
//	beanPtr:需要解析到的结构体指针
//	文档：https://opendocs.alipay.com/mini/introduce/aes
//	文档：https://opendocs.alipay.com/open/common/104567
func DecryptOpenDataToStruct(encryptedData, secretKey string, beanPtr interface{}) (err error) {
	if encryptedData == gotil.NULL || secretKey == gotil.NULL {
		return errors.New("encryptedData or secretKey is null")
	}
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入参数类型必须是以指针形式")
	}
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	var (
		block      cipher.Block
		blockMode  cipher.BlockMode
		originData []byte
	)
	aesKey, _ := base64.StdEncoding.DecodeString(secretKey)
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)
	if block, err = aes.NewCipher(aesKey); err != nil {
		return fmt.Errorf("aes.NewCipher：%w", err)
	}
	if len(secretData)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = xaes.PKCS5UnPadding(originData)
	}
	if err = json.Unmarshal(originData, beanPtr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s)：%w", string(originData), err)
	}
	return nil
}

// DecryptOpenDataToBodyMap 解密支付宝开放数据到 BodyMap
//	encryptedData:包括敏感数据在内的完整用户信息的加密数据
//	secretKey:AES密钥，支付宝管理平台配置
//	文档：https://opendocs.alipay.com/mini/introduce/aes
//	文档：https://opendocs.alipay.com/open/common/104567
func DecryptOpenDataToBodyMap(encryptedData, secretKey string) (bm gopay.BodyMap, err error) {
	if encryptedData == gotil.NULL || secretKey == gotil.NULL {
		return nil, errors.New("encryptedData or secretKey is null")
	}
	var (
		aesKey, originData []byte
		ivKey              = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		block              cipher.Block
		blockMode          cipher.BlockMode
	)
	aesKey, _ = base64.StdEncoding.DecodeString(secretKey)
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)
	if block, err = aes.NewCipher(aesKey); err != nil {
		return nil, fmt.Errorf("aes.NewCipher：%w", err)
	}
	if len(secretData)%len(aesKey) != 0 {
		return nil, errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = xaes.PKCS5UnPadding(originData)
	}
	bm = make(gopay.BodyMap)
	if err = json.Unmarshal(originData, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(originData), err)
	}
	return
}

// SystemOauthToken 换取授权访问令牌（默认使用utf-8，RSA2）
//	appId：应用ID
//	  t：支付宝私钥类型，alipay.PKCS1 或 alipay.PKCS8，默认 PKCS1
//	privateKey：应用私钥
//	grantType：值为 authorization_code 时，代表用code换取；值为 refresh_token 时，代表用refresh_token换取，传空默认code换取
//	codeOrToken：支付宝授权码或refresh_token
//	signType：签名方式 RSA 或 RSA2，默认 RSA2
//	文档：https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token
func SystemOauthToken(appId string, t PKCSType, privateKey, grantType, codeOrToken, signType string) (rsp *SystemOauthTokenResponse, err error) {
	var bs []byte
	bm := make(gopay.BodyMap)

	switch grantType {
	case "authorization_code":
		bm.Set("grant_type", "authorization_code")
		bm.Set("code", codeOrToken)
	case "refresh_token":
		bm.Set("grant_type", "refresh_token")
		bm.Set("refresh_token", codeOrToken)
	default:
		bm.Set("grant_type", "authorization_code")
		bm.Set("code", codeOrToken)
	}

	if bs, err = systemOauthToken(appId, t, privateKey, bm, "alipay.system.oauth.token", true, signType); err != nil {
		return
	}
	rsp = new(SystemOauthTokenResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if rsp.Response.AccessToken == "" {
		return nil, errors.New("access_token is NULL")
	}
	return
}

// systemOauthToken 向支付宝发送请求
func systemOauthToken(appId string, t PKCSType, privateKey string, bm gopay.BodyMap, method string, isProd bool, signType string) (bs []byte, err error) {
	bm.Set("app_id", appId)
	bm.Set("method", method)
	bm.Set("format", "JSON")
	bm.Set("charset", "utf-8")
	if signType == gotil.NULL {
		bm.Set("sign_type", RSA2)
	} else {
		bm.Set("sign_type", signType)
	}
	bm.Set("timestamp", time.Now().Format(gotil.TimeLayout))
	bm.Set("version", "1.0")
	var (
		sign    string
		baseUrl = baseUrlUtf8
	)
	if sign, err = GetRsaSign(bm, bm.Get("sign_type"), t, privateKey); err != nil {
		return nil, err
	}
	bm.Set("sign", sign)
	if !isProd {
		baseUrl = sandboxBaseUrlUtf8
	}
	_, bs, errs := xhttp.NewClient().Type(xhttp.TypeForm).Post(baseUrl).SendString(FormatURLParam(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return bs, nil
}

// monitor.heartbeat.syn(验签接口)
//	appId：应用ID
//	privateKey：应用私钥，支持PKCS1和PKCS8
//	signType：签名方式 alipay.RSA 或 alipay.RSA2，默认 RSA2
//	bizContent：验签时该参数不做任何处理，{任意值}，此参数具体看文档
//	文档：https://opendocs.alipay.com/apis/api_9/monitor.heartbeat.syn
func MonitorHeartbeatSyn(appId string, t PKCSType, privateKey, signType, bizContent string) (rsp *MonitorHeartbeatSynResponse, err error) {
	var bs []byte
	bm := make(gopay.BodyMap)
	bm.Set("biz_content", bizContent)
	bm.Set("app_id", appId)
	bm.Set("method", "monitor.heartbeat.syn")
	bm.Set("format", "JSON")
	bm.Set("charset", "utf-8")
	if signType == gotil.NULL {
		bm.Set("sign_type", RSA2)
	} else {
		bm.Set("sign_type", signType)
	}
	bm.Set("timestamp", time.Now().Format(gotil.TimeLayout))
	bm.Set("version", "1.0")

	sign, err := GetRsaSign(bm, bm.Get("sign_type"), t, privateKey)
	if err != nil {
		return nil, err
	}
	bm.Set("sign", sign)

	_, bs, errs := xhttp.NewClient().Type(xhttp.TypeForm).Post(baseUrlUtf8).SendString(FormatURLParam(bm)).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	fmt.Println("bsbsbs:", string(bs))
	rsp = new(MonitorHeartbeatSynResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
